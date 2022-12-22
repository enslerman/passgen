package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"math/rand"
	"os"
	"strconv"
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

	args := os.Args

	switch len(args) {
	case 2:
		c, err := strconv.ParseBool(args[1])
		if err == nil {
			copyToClipboard = c
		}
		fallthrough
	case 1:
		i, err := strconv.ParseInt(args[0], 10, 64)
		if err == nil {
			passLen = int(i)
		}
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
