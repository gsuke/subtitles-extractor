# 開発ルール for CLAUDE

## コメント文の形式

関数などにコメントをつけるとき、自分自身の名前は不要。

NG
```go
// CanExtract はASS形式かどうかを判定する
func (e *ASSExtractor) CanExtract(content string) bool {
}
```

OK
```go
// ASS形式かどうかを判定する
func (e *ASSExtractor) CanExtract(content string) bool {
}
```
