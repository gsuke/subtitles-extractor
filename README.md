# jimakuparse

ASS字幕ファイルからテキスト部分のみを抽出するツール

## 使い方

### 初回セットアップ

```
go mod tidy
```

### 実行

```
go run jimakuparse.go <ASSファイル>
```

### サンプル

入力
```
[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
Dialogue: 0,0:00:05.00,0:00:15.00,Default,,0,0,0,,{\pos(300,1000)}こんにちは\N
Dialogue: 0,0:00:25.00,0:00:35.00,Default,,0,0,0,,{\pos(300,1000)}今日はいい天気ですね\N
```

出力
```
こんにちは
今日はいい天気ですね
```

## Tips

```
以下の字幕のストーリーを客観的に記述してください。
ただし一部、感情や、登場人物の価値観を表す重要なセリフなどは抜粋してください。

---

```

