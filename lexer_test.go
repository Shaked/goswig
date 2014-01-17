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
			`|filter(`,
			&matchedRule{
				match:    `filter`,
				ruleType: TYPE_FILTER,
				length:   8,
			},
		},
		{
			`|filter`,
			&matchedRule{
				match:    `filter`,
				ruleType: TYPE_FILTEREMPTY,
				length:   7,
			},
		},
		{
			`function()`,
			&matchedRule{
				match:    `function`,
				ruleType: TYPE_FUNCTIONEMPTY,
				length:   10,
			},
		},
		{
			`function(param)`,
			&matchedRule{
				match:    `function`,
				ruleType: TYPE_FUNCTION,
				length:   9,
			},
		},
		{
			`(`,
			&matchedRule{
				match:    `(`,
				ruleType: TYPE_PARENOPEN,
				length:   1,
			},
		},
		{
			`)`,
			&matchedRule{
				match:    `)`,
				ruleType: TYPE_PARENCLOSE,
				length:   1,
			},
		},
		{
			`,`,
			&matchedRule{
				match:    `,`,
				ruleType: TYPE_COMMA,
				length:   1,
			},
		},
		{
			`and `,
			&matchedRule{
				match:    `&&`,
				ruleType: TYPE_LOGIC,
				length:   4,
			},
		},
		{
			`&& `,
			&matchedRule{
				match:    `&&`,
				ruleType: TYPE_LOGIC,
				length:   3,
			},
		},
		{
			`or `,
			&matchedRule{
				match:    `||`,
				ruleType: TYPE_LOGIC,
				length:   3,
			},
		},
		{
			`|| `,
			&matchedRule{
				match:    `||`,
				ruleType: TYPE_LOGIC,
				length:   3,
			},
		},
		{
			`===`,
			&matchedRule{
				match:    `===`,
				ruleType: TYPE_COMPARATOR,
				length:   3,
			},
		},
		{
			`==`,
			&matchedRule{
				match:    `==`,
				ruleType: TYPE_COMPARATOR,
				length:   2,
			},
		},
		{
			`!==`,
			&matchedRule{
				match:    `!==`,
				ruleType: TYPE_COMPARATOR,
				length:   3,
			},
		},
		{
			`!=`,
			&matchedRule{
				match:    `!=`,
				ruleType: TYPE_COMPARATOR,
				length:   2,
			},
		},
		{
			`<=`,
			&matchedRule{
				match:    `<=`,
				ruleType: TYPE_COMPARATOR,
				length:   2,
			},
		},
		{
			`<`,
			&matchedRule{
				match:    `<`,
				ruleType: TYPE_COMPARATOR,
				length:   1,
			},
		},
		{
			`>=`,
			&matchedRule{
				match:    `>=`,
				ruleType: TYPE_COMPARATOR,
				length:   2,
			},
		},
		{
			`>`,
			&matchedRule{
				match:    `>`,
				ruleType: TYPE_COMPARATOR,
				length:   1,
			},
		},
		{
			`in `,
			&matchedRule{
				match:    `in`,
				ruleType: TYPE_COMPARATOR,
				length:   3,
			},
		},
		{
			`gte `,
			&matchedRule{
				match:    `>=`,
				ruleType: TYPE_COMPARATOR,
				length:   4,
			},
		},
		{
			`gt `,
			&matchedRule{
				match:    `>`,
				ruleType: TYPE_COMPARATOR,
				length:   3,
			},
		},
		{
			`lte `,
			&matchedRule{
				match:    `<=`,
				ruleType: TYPE_COMPARATOR,
				length:   4,
			},
		},
		{
			`lt `,
			&matchedRule{
				match:    `<`,
				ruleType: TYPE_COMPARATOR,
				length:   3,
			},
		},
		{
			`=`,
			&matchedRule{
				match:    `=`,
				ruleType: TYPE_ASSIGNMENT,
				length:   1,
			},
		},
		{
			`+=`,
			&matchedRule{
				match:    `+=`,
				ruleType: TYPE_ASSIGNMENT,
				length:   2,
			},
		},
		{
			`-=`,
			&matchedRule{
				match:    `-=`,
				ruleType: TYPE_ASSIGNMENT,
				length:   2,
			},
		},
		{
			`*=`,
			&matchedRule{
				match:    `*=`,
				ruleType: TYPE_ASSIGNMENT,
				length:   2,
			},
		},
		{
			`/=`,
			&matchedRule{
				match:    `/=`,
				ruleType: TYPE_ASSIGNMENT,
				length:   2,
			},
		},
		{
			`! `,
			&matchedRule{
				match:    `!`,
				ruleType: TYPE_NOT,
				length:   2,
			},
		},
		{
			`not `,
			&matchedRule{
				match:    `!`,
				ruleType: TYPE_NOT,
				length:   4,
			},
		},
		// TODO TEST VAR
		{
			`[`,
			&matchedRule{
				match:    `[`,
				ruleType: TYPE_BRACKETOPEN,
				length:   1,
			},
		},
		{
			`]`,
			&matchedRule{
				match:    `]`,
				ruleType: TYPE_BRACKETCLOSE,
				length:   1,
			},
		},
		{
			`{`,
			&matchedRule{
				match:    `{`,
				ruleType: TYPE_CURLYOPEN,
				length:   1,
			},
		},
		{
			`:`,
			&matchedRule{
				match:    `:`,
				ruleType: TYPE_COLON,
				length:   1,
			},
		},
		{
			`}`,
			&matchedRule{
				match:    `}`,
				ruleType: TYPE_CURLYCLOSE,
				length:   1,
			},
		},
		{
			`10`,
			&matchedRule{
				match:    `10`,
				ruleType: TYPE_DOTKEY,
				length:   2,
			},
		},
		{
			`10.2`,
			&matchedRule{
				match:    `10.2`,
				ruleType: TYPE_DOTKEY,
				length:   4,
			},
		},
		{
			`-10.2`,
			&matchedRule{
				match:    `-10.2`,
				ruleType: TYPE_DOTKEY,
				length:   5,
			},
		},
		{
			`+`,
			&matchedRule{
				match:    `+`,
				ruleType: TYPE_OPERATOR,
				length:   1,
			},
		},
		{
			`-`,
			&matchedRule{
				match:    `-`,
				ruleType: TYPE_OPERATOR,
				length:   1,
			},
		},
		{
			`/`,
			&matchedRule{
				match:    `/`,
				ruleType: TYPE_OPERATOR,
				length:   1,
			},
		},
		{
			`*`,
			&matchedRule{
				match:    `*`,
				ruleType: TYPE_OPERATOR,
				length:   1,
			},
		},
		{
			`%`,
			&matchedRule{
				match:    `%`,
				ruleType: TYPE_OPERATOR,
				length:   1,
			},
		},
	}

	for _, test := range readertests {
		result := reader(test.in)
		if test.out.match != result.match || test.out.ruleType != result.ruleType || test.out.length != result.length {
			fmt.Printf("%#v\n", test.out)
			fmt.Printf("%#v\n", result)
			t.Error("Lexer failed for", test.in, "expected", test.out, " got ", result)
		}
	}
}
