package runner

import (
	"context"
	"errors"
	"os"
	"os/signal"
)

// Runner runs a sert of tasks within a given timeout and can be
// shutdown on an operating system interrupt.
type Runner struct {
	// interrupt channel reports a singal from the os.
	interrupt chan os.Signal
	complete  chan error
	c         context.Context
	tasks     []func(int)
}

// ErrTimeout is returned when a value is received on the timeout.
var ErrTimeout = errors.New("reveived timeout")

// ErrInterrupt is returned when an event from the os is received.
var ErrInterrupt = errors.New("received interrupt")

// New returns a new ready-to-use Runner.
func New(ctx context.Context) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		c:         ctx,
	}
}

// Add attaches tasks to the Runner. A task is a function that takes a int ID.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events.
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.c.Done():
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
