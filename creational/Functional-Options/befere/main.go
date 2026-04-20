package main

import "fmt"

// Server はサーバーの設定を保持する構造体
type Server struct {
	host    string
	port    int
	timeout int
	maxConn int
}

// NewServer はサーバーを生成する
// 問題点：引数が多く、何が何の値か分かりにくい
//
//	一部だけ設定したくても全引数を渡す必要がある
//	引数を増やすと既存の呼び出し箇所を全部修正しないといけない
func NewServer(host string, port int, timeout int, maxConn int) *Server {
	return &Server{
		host:    host,
		port:    port,
		timeout: timeout,
		maxConn: maxConn,
	}
}

func (s *Server) PrintConfig() {
	fmt.Printf("host=%s port=%d timeout=%d maxConn=%d\n",
		s.host, s.port, s.timeout, s.maxConn)
}

func main() {
	// タイムアウトとmaxConnはデフォルトでいいのに、
	// ゼロ値か適当な値を渡すしかない
	s1 := NewServer("localhost", 8080, 30, 100)
	s1.PrintConfig()

	// maxConnだけデフォルトにしたいが、0を渡すと意図が不明確
	s2 := NewServer("localhost", 9090, 60, 0)
	s2.PrintConfig()
}
