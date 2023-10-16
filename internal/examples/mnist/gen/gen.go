package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	var qb bytes.Buffer
	fmt.Fprintln(&qb, `package main`)
	fmt.Fprintln(&qb, `func Get() []byte {`)
	fmt.Fprint(&qb, `return []byte("`)
	print(&qb, b)
	fmt.Fprint(&qb, `")}`)
	io.Copy(os.Stdout, &qb)
	/*
		if err = os.WriteFile("olwu.go", qb.Bytes(), 0644); err != nil {
			return
		}
			if string(b) == data {
				fmt.Println("ok")
			}
	*/
}

func print(dest *bytes.Buffer, data []byte) {
	for _, b := range data {
		if b == '\n' {
			dest.WriteString(`\n`)
			continue
		}
		if b == '\\' {
			dest.WriteString(`\\`)
			continue
		}
		if b == '"' {
			dest.WriteString(`\"`)
			continue
		}
		if (b >= 32 && b <= 126) || b == '\t' {
			dest.WriteByte(b)
			continue
		}
		fmt.Fprintf(dest, "\\x%02x", b)
	}
}
