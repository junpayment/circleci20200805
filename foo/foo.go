package foo

import (
	"errors"
	"os"
)

func Do() error {
	if len(os.Getenv("Bar")) != 0 {
		return nil
	}
	return errors.New("bar")
}
