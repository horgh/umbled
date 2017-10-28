package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestAddError(t *testing.T) {
	tests := []struct {
		Format string
		Args   []interface{}
		Output *regexp.Regexp
	}{
		{
			"hi",
			nil,
			regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[-+Z]\d{2}:\d{2}: hi$`),
		},
		{
			"some error: %s",
			[]interface{}{fmt.Errorf("something went wrong")},
			regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[-+Z]\d{2}:\d{2}: some error: something went wrong$`),
		},
	}

	for _, test := range tests {
		s := &state{}
		s.addError(test.Format, test.Args...)

		if len(s.errors) != 1 {
			t.Errorf("len(s.errors) = %d, wanted %d", len(s.errors), 1)
			continue
		}

		if !test.Output.MatchString(s.errors[0]) {
			t.Errorf("addError(%s, %v) = %s, wanted %s", test.Format, test.Args,
				s.errors[0], test.Output)
		}
	}
}
