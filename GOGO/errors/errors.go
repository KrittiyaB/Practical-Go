package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type BusinessError struct {
	err error
}

func (e *BusinessError) Write() string {
	return fmt.Sprintf("Have err %s ", e.err)
}

func main() {
	errBiz := &BusinessError{errors.New("eeee")}
	fmt.Println(errBiz)
}
