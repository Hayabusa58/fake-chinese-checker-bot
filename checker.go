package main

import (
	"regexp"
	"unicode"
)

// IsKanji は与えられた文字(rune)が漢字であるか判定し、結果を返す
func IsKanji(r rune) bool {
	if unicode.In(r, unicode.Han) {
		return true
	}
	return false
}

// IsSymbol は与えられた文字(rune)が記号であるか判定し、結果を返す
func IsSymbol(r rune) bool {
	if unicode.IsSymbol(r) || unicode.IsPunct(r) {
		return true
	}
	return false
}

// IsNewLine は与えられた文字(rune)が空白記号であるか判定し、結果を返す
func IsNewline(r rune) bool {
	if unicode.Is(unicode.Pattern_White_Space, r) {
		if r == '\n' || r == '\r' || r == '\u2028' || r == '\u2029' {
			return true
		}
	}
	return false
}

// IsDiscordMention は与えられた文字列(string)が Discord の mention であるか判定し、結果を返す
func IsDiscordMention(s string) bool {
	mentionreg := regexp.MustCompile(`^<@.*$`)
	return mentionreg.MatchString(s)
}

// IsDiscordEmoji は与えられた文字列(string)が Discord の emoji であるか判定し、結果を返す
func IsDiscordEmoji(s string) bool {
	emojireg := regexp.MustCompile(`^<:[^:>]+:[0-9]+>$`)
	return emojireg.MatchString(s)
}

func IsAllowedString(s string) bool {
	rs := []rune(s)
	// flag := false
	for _, r := range rs {
		// 漢字でない場合、判定除外する文字種であるか確認する
		if !(IsKanji(r)) {
			if IsSymbol(r) || IsNewline(r) {
				continue
			} else {
				return false
			}
		} else {
			continue
		}
	}
	return true
}
