package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// samplesフォルダ内の全サンプルについて変換結果を検証する
func TestDetectAndExtract_Samples(t *testing.T) {
	samplesDir := "samples"

	// 入力ファイルを検索（*-in.* パターン）
	entries, err := os.ReadDir(samplesDir)
	if err != nil {
		t.Fatalf("samplesフォルダの読み込みに失敗: %v", err)
	}

	for _, entry := range entries {
		name := entry.Name()

		// 入力ファイル（*-in.*）のみを対象
		if !strings.Contains(name, "-in.") {
			continue
		}

		t.Run(name, func(t *testing.T) {
			// 入力ファイルのパス
			inputPath := filepath.Join(samplesDir, name)

			// 期待出力ファイルのパスを生成（sample1-in.ass → sample1-out.txt）
			baseName := strings.Split(name, "-in.")[0]
			expectedPath := filepath.Join(samplesDir, baseName+"-out.txt")

			// 入力ファイルを読み込む
			inputContent, err := os.ReadFile(inputPath)
			if err != nil {
				t.Fatalf("入力ファイルの読み込みに失敗: %v", err)
			}

			// 期待出力ファイルを読み込む
			expectedContent, err := os.ReadFile(expectedPath)
			if err != nil {
				t.Fatalf("期待出力ファイルの読み込みに失敗: %v", err)
			}

			// 変換を実行
			result, err := DetectAndExtract(string(inputContent))
			if err != nil {
				t.Fatalf("変換に失敗: %v", err)
			}

			// 結果を比較（末尾の空白・改行を正規化して比較）
			expected := strings.TrimSpace(string(expectedContent))
			actual := strings.TrimSpace(result)

			if actual != expected {
				t.Errorf("変換結果が期待値と一致しません\n期待:\n%s\n\n実際:\n%s", expected, actual)
			}
		})
	}
}
