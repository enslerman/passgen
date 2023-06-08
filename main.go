package main

import (
	"flag"
	"fmt"
	"math/rand"

	"golang.design/x/clipboard"
)

var (
	symbolsSpec = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+-=")
	symbols     = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

const (
	defaultPassLen         = 64
	defaultCopyToClipboard = true
)

func main() {
	if err := clipboard.Init(); err != nil {
		panic(err)
	}

	var (
		passLen         = defaultPassLen
		copyToClipboard = defaultCopyToClipboard

		withoutSpecSymbols bool
	)

	l := flag.Int("l", defaultPassLen, "# len of password")
	c := flag.Bool("c", defaultCopyToClipboard, "# copy to clipboard: false or true")
	s := flag.Bool("s", withoutSpecSymbols, "# without spec symbols: false or true")
	flag.Parse()

	if l != nil {
		passLen = *l
	}
	if c != nil {
		copyToClipboard = *c
	}

	b := make([]rune, passLen)
	if s {
		for i := range b {
			b[i] = symbols[rand.Intn(len(symbolsSpec))]
		}
	} else {
		for i := range b {
			b[i] = symbols[rand.Intn(len(symbols))]
		}
	}
	pass := string(b)

	if copyToClipboard {
		clipboard.Write(clipboard.FmtText, []byte(pass))
	}
	fmt.Printf("Success generate password with len: %v, copy to clipboard: %v\n", passLen, copyToClipboard)
	fmt.Println(pass)
}
