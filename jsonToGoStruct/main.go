package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {
	inputJson := ""
	structName := ""
	pkgName := ""
	inputTags := "json"
	flag.StringVar(&inputJson, "j", "", "json body input")
	flag.StringVar(&structName, "n", "Data", "generate struct name")
	flag.StringVar(&pkgName, "p", "main", "generate struct package name")
	flag.StringVar(&inputTags, "t", "json", "generate struct inputTags")
	useSub := flag.Bool("sub", false, "generate sub struct")
	flag.Parse()

	if inputJson == "" {
		fmt.Println("inputJson not allow empty, using -j specific json body")
		return
	}

	tags := strings.Split(inputTags, ",")

	// distinct tag
	tagMap := make(map[string]struct{})
	for _, tag := range tags {
		tagMap[tag] = struct{}{}
	}
	tagMap["json"] = struct{}{}

	tags = tags[:0]
	for tag, _ := range tagMap {
		tags = append(tags, tag)
	}

	generate, err := gojson.Generate(strings.NewReader(inputJson), gojson.ParseJson, structName, pkgName, tags, *useSub, true)
	if err != nil {
		fmt.Printf("generate err:%s\n", err)
		return
	}

	fmt.Printf("output struct:\n\n")
	fmt.Println(string(generate))
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Need to copy the output to the clipboard?(y/other)")
	scanner.Scan()
	text := scanner.Text()
	if strings.ToLower(text) == "y" {
		err := clipboard.WriteAll(string(generate))
		if err != nil {
			fmt.Printf("copy to clipboard failed, err:%s", err)
			return
		}
	} else {

	}
	fmt.Println("copy to clipboard success")
}
