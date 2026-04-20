package main

import "fmt"

// Server はサーバーの設定を保持する構造体
type Server struct {
	host    string
	port    int
	timeout int
	maxConn int
}

// Option はServerの設定を変更する関数型
type Option func(*Server)

// WithPort はポートを設定するオプション
func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// WithTimeout はタイムアウトを設定するオプション
func WithTimeout(timeout int) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// WithMaxConn は最大接続数を設定するオプション
func WithMaxConn(maxConn int) Option {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

// NewServer はサーバーを生成する
// デフォルト値をここで一元管理し、optsで上書きする
func NewServer(host string, opts ...Option) *Server {
	s := &Server{
		host:    host,
		port:    8080, // デフォルト
		timeout: 30,   // デフォルト
		maxConn: 100,  // デフォルト
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) PrintConfig() {
	fmt.Printf("host=%s port=%d timeout=%d maxConn=%d\n",
		s.host, s.port, s.timeout, s.maxConn)
}

func main() {
	// デフォルト値のまま使う（オプション不要）
	s1 := NewServer("localhost")
	s1.PrintConfig()

	// ポートだけ変えたい
	s2 := NewServer("localhost", WithPort(9090))
	s2.PrintConfig()

	// 複数のオプションを組み合わせる
	s3 := NewServer("localhost",
		WithPort(9090),
		WithTimeout(60),
		WithMaxConn(200),
	)
	s3.PrintConfig()
}
