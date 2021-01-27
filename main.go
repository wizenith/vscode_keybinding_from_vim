package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ProcessDistrubutionKeybindingModes split and process different type of vim mode
func ProcessDistrubutionKeybindingModes(matchArr []string, matchList *[]Keybinding) {

	// fmt.Println("match_arr: ", strings.Join(matchArr, "|"))
	beforeGroup := SplitVimFormat(matchArr[1])
	afterGroup := SplitVimFormat(matchArr[2])
	// fmt.Println("before_group", beforeGroup)
	// fmt.Println("after_group", afterGroup)
	*matchList = append(*matchList, Keybinding{Before: beforeGroup, After: afterGroup})

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

	file, err := os.Open("lab.vim")
	if err != nil {
		log.Fatal("failed to open")
	}

	scanner := bufio.NewScanner(file)

	var matchList KeybindingsOfAllModes
	for scanner.Scan() {

		currentLine := strings.Fields(scanner.Text())
		matchArr := strings.SplitN(strings.Join(currentLine, " "), " ", 3)

		if matchArr[0] == "imap" {
			ProcessDistrubutionKeybindingModes(matchArr, &matchList.Insert)
		}
		if matchArr[0] == "inoremap" {
			ProcessDistrubutionKeybindingModes(matchArr, &matchList.Inoremap)
		}

		if matchArr[0] == "nmap" {
			ProcessDistrubutionKeybindingModes(matchArr, &matchList.Normal)
		}

		if matchArr[0] == "nnoremap" {
			ProcessDistrubutionKeybindingModes(matchArr, &matchList.Nnoremap)
		}

		if matchArr[0] == "vmap" {
			ProcessDistrubutionKeybindingModes(matchArr, &matchList.Visual)
		}

		if matchArr[0] == "vnoremap" {
			ProcessDistrubutionKeybindingModes(matchArr, &matchList.Vnoremap)
		}
	}

	fmt.Println("JSON：", matchList)

	file.Close()

	ioutil.WriteFile("keybindingOfVscodeVim.json", []byte(matchList.String()), os.ModePerm)

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
