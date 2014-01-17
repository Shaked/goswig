package swig

import (
	"regexp"
	"strings"
)

const (
	TYPE_WHITESPACE = iota
	TYPE_STRING
	TYPE_FILTER
	TYPE_FILTEREMPTY
	TYPE_FUNCTION
	TYPE_FUNCTIONEMPTY
	TYPE_PARENOPEN
	TYPE_PARENCLOSE
	TYPE_COMMA
	TYPE_VAR
	TYPE_NUMBER
	TYPE_OPERATOR
	TYPE_BRACKETOPEN
	TYPE_BRACKETCLOSE
	TYPE_DOTKEY
	TYPE_ARRAYOPEN
	TYPE_ARRAYCLOSE
	TYPE_CURLYOPEN
	TYPE_CURLYCLOSE
	TYPE_COLON
	TYPE_COMPARATOR
	TYPE_LOGIC
	TYPE_NOT
	TYPE_BOOL
	TYPE_ASSIGNMENT
	TYPE_METHODOPEN
	TYPE_METHODEND
)
const TYPE_UNKNOWN = 100

type parseRule struct {
	ruleType int
	regex    []*regexp.Regexp
	idx      int
	replace  map[string]string
}

var rules = []parseRule{
	&parseRule{
		TYPE_WHITESPACE,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\s+/`),
		},
		0,
	},
	&parseRule{
		TYPE_STRING,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^""/`),
			regexp.MustCompile(`/^".*?[^\\]"/`),
			regexp.MustCompile(`/^''/`),
			regexp.MustCompile(`/^'.*?[^\\]'/`),
		},
		0,
	},
	&parseRule{
		TYPE_FILTER,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\|\s*(\w+)\(/`),
		},
		1,
	},
	&parseRule{
		TYPE_FILTEREMPTY,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\|\s*(\w+)/`),
		},
		1,
	},
	&parseRule{
		TYPE_FUNCTIONEMPTY,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\s*(\w+)\(\)/`),
		},
		1,
	},
	&parseRule{
		TYPE_FUNCTION,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\s*(\w+)\(/`),
		},
		1,
	},
	&parseRule{
		TYPE_PARENOPEN,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\(/`),
		},
		0,
	},
	&parseRule{
		TYPE_PARENCLOSE,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\)/`),
		},
		0,
	},
	&parseRule{
		TYPE_COMMA,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^,/`),
		},
		0,
	},
	&parseRule{
		TYPE_LOGIC,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^(&&|\|\|)\s*/`),
			regexp.MustCompile(`/^(and|or)\s+/`),
		},
		1,
		map[string]string{
			"and": "&&",
			"or":  "||",
		},
	},
	&parseRule{
		TYPE_COMPARATOR,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^(===|==|\!==|\!=|<=|<|>=|>|in\s|gte\s|gt\s|lte\s|lt\s)\s*/`),
		},
		1,
		map[string]string{
			"gte": ">=",
			"gt":  ">",
			"lte": "<=",
			"lt":  "<",
		},
	},
	&parseRule{
		TYPE_ASSIGNMENT,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^(=|\+=|-=|\*=|\/=)/`),
		},
	},
	&parseRule{
		TYPE_NOT,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\!\s*/`),
			regexp.MustCompile(`/^not\s+/`),
		},
		0,
		map[string]string{
			"not": "!",
		},
	},
	&parseRule{
		TYPE_BOOL,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^(true|false)\s+/`),
			regexp.MustCompile(`/^(true|false)$/`),
		},
		1,
	},
	&parseRule{
		TYPE_VAR,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^[a-zA-Z_$]\w*((\.\w*)+)?/`),
			regexp.MustCompile(`/^[a-zA-Z_$]\w*/`),
		},
	},
	&parseRule{
		TYPE_BRACKETOPEN,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\[/`),
		},
	},
	&parseRule{
		TYPE_BRACKETCLOSE,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\]/`),
		},
	},
	&parseRule{
		TYPE_CURLYOPEN,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\{/`),
		},
	},
	&parseRule{
		TYPE_COLON,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\:/`),
		},
	},
	&parseRule{
		TYPE_CURLYCLOSE,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\}/`),
		},
	},
	&parseRule{
		TYPE_DOTKEY,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^\.(\w+)/`),
		},
		1,
	},
	&parseRule{
		TYPE_DOTKEY,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^[+\-]?\d+(\.\d+)?/`),
		},
	},
	&parseRule{
		TYPE_OPERATOR,
		[]*regexp.Regexp{
			regexp.MustCompile(`/^(\+|\-|\/|\*|%)/`),
		},
	},
}

type matchedRule struct {
	match    string
	ruleType int
	length   int
}

func reader(str string) *matchedRule {
	for rule := range rules {
		for regex := range rule.regex {
			match := regex.FindAllString(str)
			if nil == match {
				continue
			}

			normalized := strings.Trim(match[rule.idx], " ")
			//TODO
			//normalized = (rule.hasOwnProperty('replace') && rule.replace.hasOwnProperty(normalized)) ? rule.replace[normalized] : normalized;

			return &matchedRule{
				match:    normalized,
				ruleType: rule.ruleType,
				length:   len(match[0]),
			}
		}
	}
	return &matchedRule{
		match:    str,
		ruleType: TYPE_UNKNOWN,
		length:   len(str),
	}
}

func read(str string) []*matchedRule {
	offset := 0
	for offset < len(str) {
		substr := str[offset:]
		match := reader(substr)
		offset += len(match.length)
		tokens = append(tokens, match)
	}
	return tokens
}
