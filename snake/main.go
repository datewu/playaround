package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type event struct{}

type snake struct {
	current string
	length  int
	x, y    int
}

var (
	a          [][]rune
	kick       = make(chan event)
	touched    = make(chan event, 1)
	onKeyboard = make(chan byte)
	logger     *log.Logger
	interMap   = map[string](chan event){
		"up":    make(chan event),
		"down":  make(chan event),
		"left":  make(chan event),
		"right": make(chan event),
		"pause": make(chan event),
	}
)

func main() {
	logFile, _ := os.Create("log.txt")
	defer logFile.Close()
	logger = log.New(logFile, "ex: ", log.LstdFlags|log.Lshortfile)

	go func() {
		for range kick {
			draw()
		}
	}()
	kick <- event{}

	go func() {
		run()
	}()

	s := snake{"right", 1, 0, 0}
	for {
		logger.Println("New round")
		s.move()
	}

}

func run() {
	for {
		d := <-onKeyboard
		switch d {
		case 'k', ' ':
			interMap["up"] <- event{}
		case 'j', '1':
			interMap["down"] <- event{}
		case 'h', '2':
			interMap["left"] <- event{}
		case 'l', '3':
			interMap["right"] <- event{}
		default:
			interMap["pause"] <- event{}

		}
	}

}

func (s *snake) move() {
	for {
		select {
		case <-interMap["up"]:
			s.current = "up"
			s.x--
		case <-interMap["down"]:
			s.current = "down"
			s.x++
		case <-interMap["left"]:
			s.current = "left"
			s.y--
		case <-interMap["right"]:
			s.current = "right"
			s.y++
		case <-interMap["pause"]:
			s.current = "pause"
		}
		if unTouch() {
			logger.Println("selected", s.x, s.y)
			selected(s.x, s.y)
			interMap[s.current] <- event{}
		}

	}

}

func unTouch() bool {
	select {
	case <-touched:
		return false
	default:
		return true
	}

}

func selected(i, j int) {
	a[i][j] ^= ' '
	kick <- event{}
	logger.Println("changing", i, j)
	a[i][j] ^= ' '

}

func draw() {
	print("\033[H\033[2J")
	fmt.Printf("%2c- ", ' ')
	for i := 0; i < len(a[0]); i++ {
		fmt.Printf("%3d", i+1)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println()
		fmt.Printf("%2d: ", i+1)
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%3c", a[i][j])
		}
	}
	time.Sleep(75 * time.Millisecond)
}

func init() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	go func() {
		var a = make([]byte, 1)
		for {
			os.Stdin.Read(make([]byte, 1))
			touched <- event{}
			onKeyboard <- a[0]
			logger.Println("keyboard", a[0])
		}
	}()

	for i := 0; i < 35; i++ {
		var row []rune
		for i := 0; i < 35; i++ {
			row = append(row, 'x')
		}
		a = append(a, row)
	}
}
