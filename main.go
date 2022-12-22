package main

import (
	"flag"
	"fmt"
	"golang.design/x/clipboard"
	"math/rand"
	"time"
)

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

const (
	defaultPassLen         = 48
	defaultCopyToClipboard = true
)

func main() {
	if err := clipboard.Init(); err != nil {
		panic(err)
	}

	var (
		passLen         = defaultPassLen
		copyToClipboard = defaultCopyToClipboard
	)

	l := flag.Int("l", defaultPassLen, "# len of password")
	c := flag.Bool("c", defaultCopyToClipboard, "# copy to clipboard: false or true")
	flag.Parse()

	if l != nil {
		passLen = *l
	}
	if c != nil {
		copyToClipboard = *c
	}

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, passLen)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	pass := string(b)

	if copyToClipboard {
		clipboard.Write(clipboard.FmtText, []byte(pass))
	}
	fmt.Printf("Success generate password with len: %v, copy to clipboard: %v\n", passLen, copyToClipboard)
	fmt.Println(pass)
}
