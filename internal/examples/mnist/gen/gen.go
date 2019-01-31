package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("test_data_set_0/output_0.pb")
	if err != nil {
		panic(err)
	}
	var qb bytes.Buffer
	fmt.Fprint(&qb, `data := "`)
	print(&qb, b)
	fmt.Fprint(&qb, `"`)
	if err = ioutil.WriteFile("olwu.go", qb.Bytes(), 0644); err != nil {
		return
	}
	/*
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
