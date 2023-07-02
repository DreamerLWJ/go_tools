package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"strings"
)

func main() {
	inputJson := ""
	structName := ""
	pkgName := ""
	flag.StringVar(&inputJson, "j", "", "json body input")
	flag.StringVar(&structName, "n", "Data", "generate struct name")
	flag.StringVar(&pkgName, "p", "main", "generate struct package name")
	flag.Parse()

	if inputJson == "" {
		fmt.Println("inputJson not allow empty, using -j specific json body")
		return
	}

	generate, err := gojson.Generate(strings.NewReader(inputJson), gojson.ParseJson, structName, pkgName, []string{"json"}, false, true)
	if err != nil {
		fmt.Printf("generate err:%s\n", err)
		return
	}

	fmt.Println(generate)
}
