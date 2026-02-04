# Subtitles Extractor

字幕ファイルからテキスト部分のみを抽出するCLIスクリプト。

対応形式
- ASS
- SRT

## 使い方

```shell
go run . [オプション] <字幕ファイル...>

# 単一ファイル（標準出力 + クリップボード出力）
go run . anime01.ass

# 複数ファイル（出力先フォルダ指定）
go run . *.ass -o ./outdir
go run . ep01.srt ep02.srt ep03.srt -o ./outdir
```

単一ファイルを指定した場合は、クリップボードに、LLMに渡すためのプロンプト込みで出力されます。
そのままLLMに貼り付けて与えることで、字幕の要約を得ることができます。

## サンプルデータ

`/samples` を参照。

## テスト

```shell
# 自動テスト
go test -v

# 手動テスト
go run . # エラー(ヘルプ)
go run . ./samples/sample1-in.ass # 標準出力 + クリップボード出力
go run . ./samples/sample1-in.ass -o outdir # 単一ファイルのフォルダ出力
go run . ./samples/sample1-in.ass ./samples/sample2-in.srt -o outdir # 単一ファイルのフォルダ出力
```

## 開発

新しい字幕形式に対応する場合は、 `extractor.go` の `SubtitlesExtractor` インターフェースを実装し、`DetectAndExtract` にそれを追加してください。
