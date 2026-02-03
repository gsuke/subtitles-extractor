package main

import (
	"fmt"
	"os"

	"golang.design/x/clipboard"
)

func main() {
	// コマンドライン引数のチェック
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	filename := os.Args[1]

	// ファイルを読み込む
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: ファイルを開けません: %v\n", err)
		os.Exit(1)
	}

	// 字幕形式を自動判別して抽出
	result, err := DetectAndExtract(string(content))
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: 処理に失敗しました: %v\n", err)
		os.Exit(1)
	}

	// 結果を出力
	fmt.Println(result)

	// クリップボードにコピー
	if err = clipboard.Init(); err != nil {
		panic(err)
	}
	clipboardMsg := `以下の字幕のストーリーを客観的に記述してください。
ただし、客観的に記述するものの、感情表現や登場人物の価値観を表す重要なセリフっぽいものはそのまま書いちゃってください。
また、時系列の改変や再構成はせず、展開をそのままなぞってください。

---

` + result
	clipboard.Write(clipboard.FmtText, []byte(clipboardMsg))
}

func printHelp() {
	help := `使い方:
  go run . <字幕ファイル>

例:
  go run . anime01.ass
  go run . anime01.srt

説明:
  字幕ファイル（ASS/SRT形式）からテキスト部分のみを抽出し、出力します。
  メタデータ（ASSの\posなど）は除去され、純粋なテキストが出力されます。
  加えて、クリップボードにコピーされます。`

	fmt.Println(help)
}
