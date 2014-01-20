package swig

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type Filters struct {
	tokens []*matchedRule
}

func (f *Filters) addslashes(str string) []*matchedRule {
	//TODO(carlosc):iterateapply
	str = strings.Replace(str, `\`, `\\`, -1)
	str = strings.Replace(str, `'`, `\'`, -1)
	str = strings.Replace(str, `"`, `\"`, -1)
	return str
}

func (f *Filters) capitalize(str string) {
	//TODO(carlosc):iterateapply
	return strings.ToUpper(str[0:1]) + strings.ToLower(str[1:])
}

//TODO(carlosc): date
//TODO(carlosc): default

func (f *Filters) escape(str, langtype string) string {
	inp := str
	//TODO(carlosc):iterateapply

	out := ""
	switch langtype {
	case "js":
		inp = strings.Replace(`\\`, "\\u005C")
		for i := 0; i < len(inp); i++ {
			code, _, _, _ := strconv.UnquoteChar(inp[i], 0)
			if code < 32 {
				//Convert to hex and uppercase the hex
				//Prepend 0 if len < 2
				out += "\\u00" //+code
			} else {
				out += inp[i]
			}
		}
		out = strings.Replace(out, `&`, `\u0026`, -1)
		out = strings.Replace(out, `<`, `\u003C`, -1)
		out = strings.Replace(out, `>`, `\u003E`, -1)
		out = strings.Replace(out, `'`, `\u0027`, -1)
		out = strings.Replace(out, `"`, `\u0022`, -1)
		out = strings.Replace(out, `=`, `\u003D`, -1)
		out = strings.Replace(out, `-`, `\u002D`, -1)
		out = strings.Replace(out, `;`, `\u003B`, -1)
		return out
	default:
		inp = strings.Replace(inp, `&lt;`, `&amp;lt;`, -1)
		inp = strings.Replace(inp, `&gt;`, `&amp;gt;`, -1)
		inp = strings.Replace(inp, `&quot;`, `&amp;quot;`, -1)
		inp = strings.Replace(inp, `&#39;`, `&amp;#39;`, -1)
		inp = strings.Replace(inp, `<`, `&lt;`, -1)
		inp = strings.Replace(inp, `>`, `&gt;`, -1)
		inp = strings.Replace(inp, `"`, `&quot;`, -1)
		inp = strings.Replace(inp, `'`, `&#39;`, -1)
		return inp
	}
}

func (f *Filters) first(str interface{}) interface{} {
	return str[0]
}

//TODO(carlosc): groupBy
//TODO(carlosc): join
//TODO(carlosc): json

func (f *Filters) last(str interface{}) interface{} {
	return input[len(str)-1]
}

func (f *Filters) lower(str string) string {
	//TODO(carlosc):iterateapply
	return strings.ToLower(str)
}

func (f *Filters) raw(str string) string {
	return f.safe(str)
}

//TODO(carlosc): replace
//TODO(carlosc): reverse

func (f *Filters) safe(str string) string {
	return str
}

//TODO(carlosc): sort

func (f *Filters) striptags(str string) string {
	re := regexp.MustCompile(`?i:(<([^>]+)>)`)
	return re.ReplaceAllString(str, "")
}

//TODO(carlosc): title
//TODO(carlosc): uniq

func (f *Filters) upper(str string) string {
	//TODO(carlosc):iterateapply
	return strings.ToUpper(str)
}

func (f *Filters) urlEncode(str string) string {
	//TODO(carlosc):iterateapply
	return url.QueryEscape(str)
}

func (f *Filters) urlDescode(str string) string {
	//TODO(carlosc):iterateapply
	return url.QueryUnescape(str)
}
