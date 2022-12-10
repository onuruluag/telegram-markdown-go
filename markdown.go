// https://core.telegram.org/bots/api#markdownv2-style

package telegrammarkdown

import (
	"fmt"
)

const escapes = `_*[]()~>#+-=|{}.!'` + "`"

type styledText struct {
	text    string
	escaped bool
}

func (s styledText) Equals(input string) bool {
	return s.text == input
}

func (s styledText) String() string {
	return s.text
}

func (s *styledText) escape() {
	if s.escaped {
		return
	}

	s.text = escape(s.text, escapes)
	s.escaped = true
}

func Text(input string) styledText {
	styled := styledText{text: input}
	styled.escape()
	return styled
}

func Space() styledText {
	styled := styledText{
		text:    " ",
		escaped: true,
	}
	return styled
}

func NewLine() styledText {
	styled := styledText{
		text:    `\\n`,
		escaped: true,
	}
	return styled
}

func Bold(input ...styledText) styledText {
	return combineAndEnclose("*", input...)
}

func BoldText(input string) styledText {
	return Bold(Text(input))
}

func Italic(input ...styledText) styledText {
	return combineAndEnclose("_", input...)
}

func ItalicText(input string) styledText {
	return Italic(Text(input))
}

func Underline(input ...styledText) styledText {
	return combineAndEnclose("__", input...)
}

func UnderlineText(input string) styledText {
	return Underline(Text(input))
}

func Strikethrough(input ...styledText) styledText {
	return combineAndEnclose("~", input...)
}

func StrikethroughText(input string) styledText {
	return Strikethrough(Text(input))
}

func Spoiler(input ...styledText) styledText {
	return combineAndEnclose("||", input...)
}

func SpoilerText(input string) styledText {
	return Spoiler(Text(input))
}

func InlineURL(text, url string) styledText {
	styled := styledText{text: fmt.Sprintf("[%s](%s)",
		escape(text, escapes),
		escape(url, `)\`),
	),
		escaped: true,
	}
	return styled
}

func MentionUser(username string) styledText {
	return styledText{text: fmt.Sprintf("@%s", username), escaped: true}
}

func InlineMentionUser(text, userId string) styledText {
	return InlineURL(text, fmt.Sprintf("tg://user?id=%s", userId))
}

func Hashtag(input string) styledText {
	return styledText{text: `\#` + replaceSpaces(escape(input, escapes)), escaped: true}
}

func InlineFixWidth(input string) styledText {
	return styledText{text: encloseText(input, "`", "`"), escaped: true}
}

func Preformatted(input string) styledText {
	return CodeBlock("", input)
}

func CodeBlock(language, input string) styledText {
	return styledText{
		text:    encloseText(input, "```"+language+"\n", "```\n"),
		escaped: true,
	}
}
