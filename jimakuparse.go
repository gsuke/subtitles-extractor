package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"golang.design/x/clipboard"
)

// ASS字幕ファイルからテキストを抽出する
func ProcessASSFile(input io.Reader) (string, error) {
	lines, err := extractDialogueLines(input)
	if err != nil {
		return "", err
	}

	var textParts []string
	for _, line := range lines {
		text := extractTextFromDialogue(line)
		if text != "" {
			textParts = append(textParts, text)
		}
	}

	return strings.Join(textParts, "\n"), nil
}

// "Dialogue:" で始まる行のみを抽出する
func extractDialogueLines(input io.Reader) ([]string, error) {
	var dialogueLines []string
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Dialogue: ") {
			dialogueLines = append(dialogueLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dialogueLines, nil
}

// Dialogue行からテキスト部分(Text列)を抽出する
func extractTextFromDialogue(line string) string {
	// 行の書式: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
	parts := strings.SplitN(line, ",", 10)
	if len(parts) < 10 {
		return ""
	}

	text := parts[9]

	// 抽出したText列から、不要なものを除去
	text = removeASSMetadata(text)
	text = replaceNewlineCode(text)
	text = strings.TrimSpace(text)

	return text
}

// ASS形式のメタデータ（{\pos(...)}など）を除去する
func removeASSMetadata(text string) string {
	// {\...} 形式のタグを除去
	re := regexp.MustCompile(`\{[^}]*\}`)
	return re.ReplaceAllString(text, "")
}

// \N を改行に置換する
func replaceNewlineCode(text string) string {
	return strings.ReplaceAll(text, `\N`, "\n")
}

func main() {
	// コマンドライン引数のチェック
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	filename := os.Args[1]

	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: ファイルを開けません: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// ASS字幕を処理
	result, err := ProcessASSFile(file)
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
	clipboard.Write(clipboard.FmtText, []byte(result))
}

func printHelp() {
	help := `使い方:
  go run jimakuparse.go <ASSファイル>

例:
  go run jimakuparse.go anime01.ass

説明:
  ASS字幕ファイルからテキスト部分のみを抽出し、出力します。
  メタデータ(\pos など)は除去され、純粋なテキストが出力されます。
  加えて、クリップボードにコピーされます。`

	fmt.Println(help)
}
