package telegrammarkdown

import (
	"fmt"
	"strings"
)

func Combine(input ...styledText) styledText {
	return styledText{
		text:    getCombinedText(input...),
		escaped: true,
	}
}

func CombineWithSpace(input ...styledText) styledText {
	return styledText{
		text:    getCombinedTextWithSeparator(" ", input...),
		escaped: true,
	}
}

func CombineWithNewLine(input ...styledText) styledText {
	return styledText{
		text:    getCombinedTextWithSeparator(`\\n`, input...),
		escaped: true,
	}
}

func escape(input, charset string) string {
	replacementPairs := make([]string, 0, 2*len(charset))

	for _, ch := range charset {
		replacementPairs = append(replacementPairs, string(ch))
		replacementPairs = append(replacementPairs, `\`+string(ch))
	}
	replacer := strings.NewReplacer(replacementPairs...)

	return replacer.Replace(input)
}

func replaceSpaces(input string) string {
	replacements := []string{
		"     ", `\_`,
		"    ", `\_`,
		"   ", `\_`,
		"  ", `\_`,
		" ", `\_`,
		"\\-", `\_`,
		"-", `\_`,
		`\n`, `\_`,
	}
	replacer := strings.NewReplacer(replacements...)
	return replacer.Replace(input)
}

func getCombinedText(input ...styledText) string {
	return getCombinedTextWithSeparator("", input...)
}

func getCombinedTextWithSeparator(separator string, input ...styledText) string {
	var combined string
	for _, s := range input {
		s.escape()
		combined += s.text + separator
	}

	return strings.TrimSuffix(combined, separator)
}

func encloseText(input, prefix, suffix string) string {
	return fmt.Sprintf("%s%s%s", prefix, input, suffix)
}

func combineAndEnclose(closure string, input ...styledText) styledText {
	combined := getCombinedText(input...)
	return styledText{text: encloseText(combined, closure, closure), escaped: true}
}
