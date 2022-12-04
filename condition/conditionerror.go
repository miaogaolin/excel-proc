package condition

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// rewrite error
type ConditionError struct {
	antlr.DefaultErrorListener
	// antlr.DefaultErrorStrategy
	Errors []error
}

func NewConditionError() *ConditionError {
	return &ConditionError{}
}
func (d *ConditionError) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.Errors = append(d.Errors, fmt.Errorf("column:%d, %s", column, msg))
}
