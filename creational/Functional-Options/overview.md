# Functional Options Pattern

## 概要

Functional Options パターンとは、構造体の初期化時に**オプション引数を柔軟に渡す**ための Go イディオムです。
Rob Pike が提唱し、Go のライブラリ設計で広く使われています。

---

## 問題（適用前）

```go
func NewServer(host string, port int, timeout int, maxConn int) *Server { ... }
```

このような設計には以下の問題があります。

- 引数が増えると呼び出し側が煩雑になる
- 一部だけ指定したいとき、ゼロ値を渡す必要がある
- 引数の順番を間違えてもコンパイルエラーにならない

---

## 解決策（適用後）

**「オプションを関数として表現する」** のが Functional Options パターンです。

### ① Option 型を定義する

```go
type Option func(*Server)
```

### ② オプションを返す関数を定義する

```go
func WithPort(port int) Option {
    return func(s *Server) {
        s.port = port
    }
}

func WithTimeout(timeout int) Option {
    return func(s *Server) {
        s.timeout = timeout
    }
}
```

### ③ コンストラクタで可変長引数として受け取る

```go
func NewServer(host string, opts ...Option) *Server {
    s := &Server{host: host, port: 8080, timeout: 30} // デフォルト値
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

### ④ 呼び出し側はシンプルに

```go
// ポートだけ変えたい
s1 := NewServer("localhost", WithPort(9090))

// タイムアウトも変えたい
s2 := NewServer("localhost", WithPort(9090), WithTimeout(60))

// デフォルトのままでいい
s3 := NewServer("localhost")
```

---

## まとめ

| 観点 | 適用前 | 適用後 |
|------|--------|--------|
| 引数の数 | 固定（全部渡す必要あり） | 必要なものだけ渡せる |
| デフォルト値 | 呼び出し側が意識する | コンストラクタ内で管理 |
| 拡張性 | 引数追加で既存コードが壊れる | 既存コードに影響なし |
| 可読性 | 引数の意味が不明確 | `WithXxx` で意図が明確 |

---

## ポイント

- `Option` は `func(*Server)` 型のエイリアス
- オプションは**クロージャ**として値をキャプチャする
- デフォルト値はコンストラクタ内にまとめて書ける
- 新しいオプションを追加しても、既存の呼び出しコードは変更不要
