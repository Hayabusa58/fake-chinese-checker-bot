package main

import (
	"testing"
)

func TestIsAllowedString(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected bool
	}{
		{"全て漢字", "全漢字", true},
		{"全て漢字、改行あり", "全漢字\n改行", true},
		{"漢字と記号", "漢字！？（）。、", true},
		{"メンション", "@h4y4bus4 あいうえお", false},
		{"絵文字", ":teikoku:", false},
		{"記号のみ", "！？（）。、", true},
		{"全てひらがな", "あいうえお", false},
		{"全てかたかな", "アイウエオ", false},
		{"ひらがな+漢字", "あいうえお漢字", false},
		{"漢字+ひらがな", "漢字あいうえお", false},
		{"漢字+カタカナ", "漢字アイウエオ", false},
		{"アルファベット+漢字", "abcdEF漢字", false},
		{"漢字+アルファベット", "漢字abcdEF", false},
		{"漢字+アルファベット", "漢字abcdEF", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsAllowedString(tt.content)
			if result != tt.expected {
				t.Errorf("IsAllowedString(%s) = %v; want %v", tt.content, result, tt.expected)
			}
		})
	}

}
