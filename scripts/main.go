package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kotaroooo0/gojaconv/jaconv"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	nicknames, err := getNicknames(inPath)

	if err != nil {
		panic(err)
	}

	nicknameJPENList := convertToNicknameJPENList(nicknames)

	if len(*nicknameJPENList) == 0 {
		panic("no nickname was found")
	}

	if err = saveJson(outPath, nicknameJPENList); err != nil {
		panic(err)
	}

	fmt.Printf("%v Kawaii Japanese Nicknames were exported!! 🎉\n", len(*nicknames))
}

func getNicknames(inPath string) (*[]string, error) {
	var nicknames []string
	f, err := os.Open(inPath)
	if err != nil {
		return &nicknames, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nickname := scanner.Text()
		if nickname == "" {
			continue
		}
		nicknames = append(nicknames, nickname)
	}
	return &nicknames, nil
}

func convertToNicknameJPENList(nicknames *[]string) *[][]string {
	var result [][]string
	for _, nickname := range *nicknames {
		hebon := jaconv.ToHebon(nickname)
		result = append(result, []string{nickname, hebon})
	}
	return &result
}

func saveJson(outPath string, jpenList *[][]string) error {
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	str, err := json.MarshalIndent(*jpenList, "", "  ")
	if err != nil {
		return err
	}

	data := []byte(str)
	_, err = f.Write(data)
	return err
}
