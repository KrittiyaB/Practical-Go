package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var err error
	originalError := &os.PathError{"read", "/a/b", errors.New("read error")}

	err = fmt.Errorf("context for error: %v", originalError)
	fmt.Printf("1 - error type: %T , value: %v\n", err, err)

	err = fmt.Errorf("wrapping context error: %w", originalError)
	fmt.Printf("2 - error type: %T , value: %v\n", err, err)

	errInner := errors.Unwrap(err)
	fmt.Printf("3 - error type: %T , value: %v\n", errInner, errInner)

	errInner2 := errors.Unwrap(errInner)
	fmt.Printf("4 - error type: %T , value: %v\n", errInner2, errInner2)

	errInner3 := errors.Unwrap(errInner2)
	fmt.Printf("5 - error type: %T , value: %v\n", errInner3, errInner3)

}
