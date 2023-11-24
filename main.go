package main

import (
	"flag"
	"fmt"
	"golang.design/x/clipboard"
	"log"
	"math/rand"
	"strconv"
)

var (
	symbolsSpec = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+-=")
	symbols     = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

func main() {
	if err := clipboard.Init(); err != nil {
		panic(err)
	}

	var (
		passLen         = 64
		copyToClipboard = true
		specSymbols     = false

		err error
	)

	l := flag.Int("l", passLen, "# len of password")
	c := flag.String("c", fmt.Sprint(copyToClipboard), "# copy to clipboard: false or true")
	s := flag.String("s", fmt.Sprint(specSymbols), "# spec symbols: false or true")
	flag.Parse()

	if l != nil {
		passLen = *l
	}
	if s != nil {
		specSymbols, err = strconv.ParseBool(*s)
		if err != nil {
			log.Fatalf("cannot parse param 's' - without spec symbols, err: %v", err)
		}
	}
	if c != nil {
		copyToClipboard, err = strconv.ParseBool(*c)
		if err != nil {
			log.Fatalf("cannot parse param 'c' - copy to clipboard, err: %v", err)
		}
	}

	b := make([]rune, 0, passLen)
	if specSymbols {
		for i := 0; i < passLen; i++ {
			b = append(b, symbolsSpec[rand.Intn(len(symbolsSpec))])
		}
	} else {
		for i := 0; i < passLen; i++ {
			b = append(b, symbols[rand.Intn(len(symbols))])
		}
	}
	pass := string(b)

	if copyToClipboard {
		clipboard.Write(clipboard.FmtText, []byte(pass))
	}
	fmt.Printf("Success generate password with len: %v; spec symbols: %v; copy to clipboard: %v\n", passLen, specSymbols, copyToClipboard)
	fmt.Println(pass)
}
