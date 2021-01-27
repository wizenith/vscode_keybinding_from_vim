package main

import (
	"bufio"
	"bytes"
	"flag"
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
	var filepath string
	flag.StringVar(&filepath, "filepath", "", "your init.vim filepath")
	flag.Parse()

	if len(filepath) == 0 {
		// fmt.Fprintf(os.Stderr, "You must specify the filepath your init.vim")
		// fmt.Errorf("You must specify the filepat of your init.vim: %v", os.Stderr)
		// return
		filepath = "lab.vim"
	}

	fmt.Printf("Your vim config filepath was: %s\n", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("failed to open")
	}

	scanner := bufio.NewScanner(file)

	var matchList KeybindingsOfAllModes

	modeMap := map[string]*[]Keybinding{
		"imap":     &matchList.Insert,
		"inoremap": &matchList.Inoremap,
		"nmap":     &matchList.Normal,
		"nnoremap": &matchList.Nnoremap,
		"vmap":     &matchList.Visual,
		"vnoremap": &matchList.Vnoremap,
	}

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		currentLine := strings.Fields(scanner.Text())
		matchArr := strings.SplitN(strings.Join(currentLine, " "), " ", 3)

		if strings.Contains(matchArr[1], "silent") && strings.Contains(matchArr[1], "expr") {
			continue
		}
		if mode, ok := modeMap[matchArr[0]]; ok {
			ProcessDistrubutionKeybindingModes(matchArr, mode)
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
