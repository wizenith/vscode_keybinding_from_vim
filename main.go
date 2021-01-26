package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Keybinding | type of vscodevim keybinding
type Keybinding struct {
	Before []string `json:"before"`
	After  []string `json:"after"`
}

type KeybindingsOfAllModes struct {
	Normal   []Keybinding `json:"vim.normalModeKeyBindings"`
	Nnoremap []Keybinding `json:"vim.normalModeKeyBindingsNonRecursive"`
	Insert   []Keybinding `json:"vim.insertModeKeyBindings"`
	Inoremap []Keybinding `json:"vim.insertModeKeyBindingsNonRecursive"`
	Visual   []Keybinding `json:"vim.visualModeKeyBindings"`
	Vnoremap []Keybinding `json:"vim.visualModeKeyBindingsNonRecursive"`
}

func print(arr []string) {

	fmt.Println(strings.Join(arr, " | "))

}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	return bytes.TrimRight(data, "\r")
}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {

		return 0, nil, nil
	}
	if data[0] == '<' {

		right := bytes.IndexByte(data, '>')

		if right == -1 {

			return 1, data[:1], nil
		}
		return right + 1, data[:right+1], nil
	}

	if i := bytes.IndexByte(data, '\n'); i >= 0 {

		return i + 1, dropCR(data[0:i]), nil
	}

	return 1, data[:1], nil
}

func main() {

	file, err := os.Open("init.vim")
	if err != nil {
		log.Fatal("failed to open")
	}

	scanner := bufio.NewScanner(file)

	var matchList []KeybindingsOfAllModes
	for scanner.Scan() {

		currentLine := strings.Fields(scanner.Text())
		matchArr := strings.SplitN(strings.Join(currentLine, " "), " ", 3)

		if matchArr[0] == "inoremap" || matchArr[0] == "imap" {
			fmt.Println("match_arr: ", strings.Join(matchArr, "|"))
			beforeGroup := SplitVimFormat(matchArr[1])
			afterGroup := SplitVimFormat(matchArr[2])
			fmt.Println("before_group", beforeGroup)
			fmt.Println("after_group", afterGroup)
			matchList = append(matchList, KeybindingsOfAllModes{Insert{Before: beforeGroup, After: afterGroup}})
		}

	}

	jsonData := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(jsonData)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "  ")
	jsonEncoder.Encode(matchList)
	fmt.Println("JSON：", jsonData.String())

	file.Close()

	ioutil.WriteFile("keybindingOfVscodeVim.json", []byte(jsonData.String()), os.ModePerm)

}

// SplitVimFormat | Analyize special key and split into the proper format of vscodevim keybinding
func SplitVimFormat(text string) []string {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(splitFunc)
	var chars []string
	for scanner.Scan() {
		chars = append(chars, scanner.Text())
	}
	return chars
}
