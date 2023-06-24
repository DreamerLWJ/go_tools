package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
)

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	randomString := base64.URLEncoding.EncodeToString(bytes)
	return randomString[:length], nil
}

func main() {
	logHead := "go_tools_random_string:"
	var length int
	flag.IntVar(&length, "l", 0, "using -l Specifies the length of the generated string")
	flag.Parse()
	if length == 0 {
		fmt.Println(logHead + "the length of the generated string not allow 0 or empty, using -l specifies it")
		return
	}
	randomString, err := GenerateRandomString(length)
	if err != nil {
		fmt.Println(logHead+"Failed to generate random string:", err)
		return
	}

	fmt.Println("Random String: ", randomString)
}
