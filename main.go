package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.OpenFile("color-convert/simple.css", os.O_RDWR, 0666)
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
	fmt.Println("read file length:", n)
	fmt.Printf("%s\n", buffer[:n])
	re := regexp.MustCompile(`#([0-9a-fA-F])+`)
	fmt.Println(re.ReplaceAllStringFunc(string(buffer[:n]), convert()))
}

func convert() func(string) string {
	return func(s string) string {
		return hex_to_dec(s)
	}
}

func hex_to_dec(s string) string {
	d := make(map[string]int)
	for i, x := range "0123456789abcdef" {
		d[string(x)] = i
	}
	fs := make([]string, 6)
	sp := strings.Split(s, "")
	fs = append(fs, "rgb(")
	for i := 1; i < len(s)-1; i += 2 {
		color := strconv.Itoa((d[sp[i]] << 4) + d[sp[i+1]])
		if i == 5 {
			fs = append(fs, color)
		} else {
			fs = append(fs, color+" ")
		}
	}
	fs = append(fs, ")")
	return strings.Join(fs, "")
}
