package ast

import (
	"errors"
	"fmt"
)

func Debug(line, col int) string {
	return fmt.Sprintf("[%d:%d]", line, col)
}

func InvalidTokenErr(line, col int) error {
	return errors.New(fmt.Sprintf(Debug(line, col) + " Invalid token"))
}
