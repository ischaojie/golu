package main

import (
	"encoding/base64"
	"flag"
	"fmt"
)

var (
	h bool

	e bool
	d bool
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&e, "e", false, "base64 encode")
	flag.BoolVar(&d, "d", false, "base64 decode")
}

func main() {
	flag.Parse()
	msg := flag.Arg(0)

	if h {
		fmt.Println("base64 tool")
		fmt.Println("Usage: base64 [-d] [-e] <some str>")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	if e && d == false || e==false && d==false{
		encoded := base64.StdEncoding.EncodeToString([]byte(msg))
		fmt.Println(encoded)
	}
	if d && e == false{
		decoded, err := base64.StdEncoding.DecodeString(msg)
		if err != nil {
			fmt.Println("decode error:", err)
			return
		}
		fmt.Println(string(decoded))
	}

}
