package main

import (
	"regexp"
	"unicode"
)

func IsKanji(r rune) bool {
	return unicode.In(r, unicode.Han)
}

func IsSymbol(r rune) bool {
	return unicode.IsSymbol(r) || unicode.IsPunct(r)
}

func IsDiscordMention(s string) bool {
	mentionreg := regexp.MustCompile(`^<@.*$`)
	return mentionreg.MatchString(s)
}

func IsDiscordEmoji(s string) bool {
	emojireg := regexp.MustCompile(`^<:[^:>]+:[0-9]+>$`)
	return emojireg.MatchString(s)
}

func IsAllowedCharacter(s string) bool {
	r := []rune(s)
	return IsKanji(r[0]) || IsSymbol(r[0]) || IsDiscordMention(s) || IsDiscordEmoji(s)
}
