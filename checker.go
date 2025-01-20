package main

import (
	"regexp"
	"unicode"
)

func IsKanji(rs []rune) bool {
	flag := true
	for _, r := range rs {
		if !(unicode.In(r, unicode.Han)) {
			flag = false
		}
	}
	return flag
}

func IsSymbol(rs []rune) bool {
	flag := true
	for _, r := range rs {
		if !(unicode.IsSymbol(r) || unicode.IsPunct(r)) {
			flag = false
		}
	}
	return flag
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
	return IsKanji(r) || IsSymbol(r) || IsDiscordMention(s) || IsDiscordEmoji(s)
}
