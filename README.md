# jimakuparse

字幕ファイルからテキスト部分のみを抽出するCLIスクリプト。

対応形式
- ASS
- SRT

## 使い方

```shell
go run . <字幕ファイル>
```

## サンプルデータ

`/samples` を参照。

```shell
# 実行例
go run . ./samples/sample1-in.ass
go run . ./samples/sample2-in.srt
```
