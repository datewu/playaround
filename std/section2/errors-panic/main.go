package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	errBadThing   = errors.New("somethig bad")
	errWorseThing = errors.New("worse thing")
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered from:", r)
		}
	}()

	errorsSwitch()
	customErrors() // should come before panic
}

func errorsSwitch() {
	//err := doSomethingBad()
	err := doSomethingWorse()
	if err != nil {
		switch err {
		case errBadThing:
			log.Println("Uh, oh:", err)
		case errWorseThing:
			panic("Abandon ship!")
		}
	}
}

func doSomethingBad() error {
	return errBadThing
}

func doSomethingWorse() error {
	return errWorseThing
}

func customErrors() {
	// success, err := doRiskyManeuver(riskLevelLow)
	// success, err := doRiskyManeuver(riskLevelMedium)
	success, err := doRiskyManeuver(riskLevelHigh)
	if err != nil {
		switch err.(type) {
		case *badError:
			fmt.Println("it's bad Captain:", err, err.(*badError).additionContext)
		case *worseError:
			fmt.Println("It's really bad Captain:", err)
			if err.(*worseError).warpDriveUnstable {
				fmt.Println("The warp drive is unstable Captain! What do we do?")
			}
		}
	}
	if success {
		fmt.Println("We made it port, Captain!")
	}

}

type badError struct {
	message         string
	additionContext string
}

func (e *badError) Error() string {
	return e.message
}

type worseError struct {
	message           string
	warpDriveUnstable bool
}

func (e *worseError) Error() string {
	return e.message
}

type riskLevel int

const (
	riskLevelLow riskLevel = iota
	riskLevelMedium
	riskLevelHigh
)

func doRiskyManeuver(level riskLevel) (bool, error) {
	switch level {
	case riskLevelMedium:
		return false, &badError{
			message:         "We're running on empty.",
			additionContext: "Not enough dillitium crystals",
		}
	case riskLevelHigh:
		return false, &worseError{
			message:           "Plasma chamber's melting.",
			warpDriveUnstable: true,
		}
	default:
		return true, nil
	}
}
