package iowriter

import (
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

// インターフェースの実装サンプル
func InterfaceExam() {
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()
}

// テキストへの出力
func OutputText() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write, err := file.Write([]byte("os.File example\n"))
	if err != nil {
		return
	}
	fmt.Println(write)
}

// 標準出力に出力する
func OutputStdout() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}

func StrBuilder() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}

// ネットワーク通信サンプル
// ascii.jpの80番ポートに接続し、HTTPのGETリクエストを送信する。
func NetDial() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}

// HttpServerと組み合わせて、http://localhost:8080/ にアクセスすると、
// http.ResponseWriter sample と表示される。簡易的なWebサーバー。
func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}

func HttpServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// 複数の出力先に同時に書き込む
func MultiWite() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")

}

func CompressAndSendFile() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.writer example\n")
	writer.Close()
}
