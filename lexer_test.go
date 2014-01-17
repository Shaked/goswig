package swig

import (
	"fmt"
	"testing"
)

func TestReader(t *testing.T) {
	readertests := []struct {
		in  string
		out *matchedRule
	}{
		{
			` `,
			&matchedRule{
				match:    ``,
				ruleType: TYPE_WHITESPACE,
				length:   1,
			},
		},
		{
			`"string"`,
			&matchedRule{
				match:    `"string"`,
				ruleType: TYPE_STRING,
				length:   8,
			},
		},
		{
			`|filter()`,
			&matchedRule{
				match:    `filter`,
				ruleType: TYPE_FILTER,
				length:   8,
			},
		},
	}

	for _, test := range readertests {
		result := reader(test.in)
		if test.out.match != result.match || test.out.ruleType != result.ruleType || test.out.length != result.length {
			fmt.Printf("%#v\n", test.out)
			fmt.Printf("%#v\n", result)
			t.Error("Lexer failed for ", test.in, " expected ", test.out, " got ", result)
		}
	}
}
