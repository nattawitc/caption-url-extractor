package main

import (
	"regexp"
	"strings"
	"syscall/js"
)

func processText(this js.Value, p []js.Value) interface{} {
	if len(p) < 1 {
		return js.Null()
	}
	str := p[0].String()
	// Split the string by spaces, you can adjust the delimiter or logic as needed
	/* fmt.Println(str) */
	/* fmt.Println(sanitize(str)) */
	/* fmt.Println(strings.Join(process(str), "\n")) */
	words := process(str)
	// Join words with a newline separator for demonstration
	result := strings.Join(words, "\n")
	return js.ValueOf(result)
}

func main() {
	/* raw := `https://www  */
	/* .p ixiv.net/en/art	w orks/122493699 */
	/* something something https://www.pixiv .net/en/artworks/1165049 07` */
	/**/
	c := make(chan struct{}, 0)

	js.Global().Set("processText", js.FuncOf(processText))

	<-c
}

type regexpBase struct {
	prefix  string
	baseURL string
	path    string
}

var (
	s              = regexpBuilder()
	matcher        = regexp.MustCompile(s)
	removeSpaceReg = regexp.MustCompile(`\s+`)
	// (http(s?)://)?(www.|m.)
	defaultPrefix = `(h\s*t\s*t\s*p\s*(s\s*)?:\s*/\s*/\s*)?(w\s*w\s*w\s*\.\s*|m\s*\.\s*)?`
)

var urlBuilders = []regexpBase{
	regexpBase{
		baseURL: "pixiv.net",
		path:    `en\/(artworks|users)\/\d+`,
	},
	regexpBase{
		baseURL: "twitter.com",
		path:    `[a-zA-Z0-9_]+(\/status\/\d+)?`,
	},
	regexpBase{
		baseURL: "x.com",
		path:    `[a-zA-Z0-9_]+(\/status\/\d+)?`,
	},
	regexpBase{
		baseURL: "dlsite.com",
		path:    `[a-zA-Z0-9_-]+(\/(([a-zA-Z0-9_-]+)|=))*(\.html)?`,
	},
}

func process(raw string) []string {
	raw = sanitize(raw)

	/* return matcher.FindAllString(raw, -1) */
	all := matcher.FindAllString(raw, -1)
	for i := range all {
		all[i] = removeSpaceReg.ReplaceAllString(all[i], "")
	}

	return all
}

func sanitize(s string) string {
	return strings.ReplaceAll(s, "(.)", ".")
	/* return removeSpaceReg.ReplaceAllString(s, "") */
}

func regexpBuilder() string {
	matches := []string{}
	for _, urlBuilder := range urlBuilders {
		if urlBuilder.prefix == "" {
			urlBuilder.prefix = defaultPrefix
		}
		match := urlBuilder.prefix + spacer(urlBuilder.baseURL) + `\/` + urlBuilder.path
		matches = append(matches, match)
	}

	return "(" + strings.Join(matches, ")|(") + ")"
}

func spacer(str string) string {
	runes := []rune(str)
	strs := make([]string, len(runes))
	for i := range runes {
		strs[i] = regexp.QuoteMeta(string(runes[i]))
	}

	return strings.Join(strs, `\s*`)
}
