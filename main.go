package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.OpenFile("color-convert/advanced.css", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		os.Exit(1)
	}
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println("read file error:", err)
		os.Exit(1)
	}
	err = file.Close()
	if err != nil {
		fmt.Println("close file error:", err)
		os.Exit(1)
	}
	fmt.Println("%s\n", string(buffer[:n]))
	re := regexp.MustCompile(`#([0-9a-fA-F])+`)
	fmt.Println(re.ReplaceAllStringFunc(string(buffer[:n]), convert(re)))
}

func convert() func(string) string {
	return func(s string) string {
		return hex_to_dec(s)
	}
}

func hex_to_dec(s string) string {
	d := createValueMap()
	fs := make([]string, 6)
	st := strings.Split(s, "")
	sd := st[1:]
	var sp []string
	sp = normalizeHexString(s, sd, sp)

	fs = append(fs, "rgb(")
	fs = build_string(s, d, sp, fs)
	fs = append(fs, ")")
	return strings.Join(fs, "")
}

func normalizeHexString(s string, sd []string, sp []string) []string {
	if len(s) == 3 || len(s) == 4 {
		for _, x := range sd {
			sp = append(sp, x, x)
		}
	} else {
		sp = sd
	}
	return sp
}

func createValueMap() map[string]int {
	d := make(map[string]int)
	for i, x := range "0123456789abcdef" {
		d[string(x)] = i
	}
	return d
}

func build_string(s string, d map[string]int, sp []string, fs []string) []string {
	for i := 0; i < len(s)-1; i += 2 {
		color := strconv.Itoa((d[sp[i]] << 4) + d[sp[i+1]])
		if i > 5 {
			fs = append(fs, "/ "+color)
		} else if i == 5 && len(sp) == 8 {
			fs = append(fs, color)
		} else {
			fs = append(fs, color+" ")
		}
	}
	return fs
}
