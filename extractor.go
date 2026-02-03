package main

import (
	"fmt"
)

// 字幕抽出のインターフェース
type SubtitleExtractor interface {
	Extract(content string) (string, error)
	CanExtract(content string) bool
}

// 字幕形式を自動判別して抽出する
func DetectAndExtract(content string) (string, error) {
	extractors := []SubtitleExtractor{
		&ASSExtractor{},
		&SRTExtractor{},
	}

	for _, extractor := range extractors {
		if extractor.CanExtract(content) {
			return extractor.Extract(content)
		}
	}

	return "", fmt.Errorf("未対応の字幕形式です")
}