package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

//Keybinding | type of vscodevim keybinding
type Keybinding struct {
	Before []string `json:"before"`
	After  []string `json:"after"`
}

// type AllKeybindings struct {
// 	Items []Keybinding
// }

// func (kb *AllKeybindings) AddItem(item Keybinding) {
// 	kb.Items = append(kb.Items, item)
// }

func print(arr []string) {
	// print 的時候用 ▲ 做分隔
	fmt.Println(strings.Join(arr, " | "))

}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	return bytes.TrimRight(data, "\r")
}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		// 遇到空字串，那就回傳空字串
		return 0, nil, nil
	}
	if data[0] == '<' {
		// 遇到開頭是 <，就把整個 <...> 當作同一串切下來
		right := bytes.IndexByte(data, '>')
		// fmt.Println("right:", right)
		if right == -1 {
			// fmt.Println("-1, exist")
			return 1, data[:1], nil
		}
		return right + 1, data[:right+1], nil
	}

	// return newline
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}

	// 否則就只切一個字元下來
	return 1, data[:1], nil
}

func main() {
	// read file and get each lines
	// analyze and assign arrays
	// export as json format in array
	file, err := os.Open("lab.vim")
	if err != nil {
		log.Fatal("failed to open")
	}

	scanner := bufio.NewScanner(file)
	// scanner.Split(strings.SplitN(scanner.Text(), " ", 3))
	// scanner.Split(bufio.ScanLines)
	// var text []string

	var matchList []Keybinding
	for scanner.Scan() {
		// fmt.Println("scanner.Text():", scanner.Text())
		currentLine := strings.Fields(scanner.Text())
		matchArr := strings.SplitN(strings.Join(currentLine, " "), " ", 3)
		// fmt.Println("arr", strings.Join(arr, "|"))
		// fmt.Println("arr", match_arr[0])
		// if strings.Contains(match_arr[0], "inoremap") || strings.Contains(match_arr[0], "imap") {
		if matchArr[0] == "inoremap" || matchArr[0] == "imap" {
			fmt.Println("match_arr: ", strings.Join(matchArr, "|"))
			// if !strings.HasPrefix(match_arr[0], "\"") {
			// 	fmt.Println("line:", match_arr)
			// fmt.Println("before group:", strings.Split(match_arr[1], ""))
			beforeGroup := SplitVimFormat(matchArr[1])
			afterGroup := SplitVimFormat(matchArr[2])
			fmt.Println("before_group", beforeGroup)
			fmt.Println("after_group", afterGroup)
			// fmt.Println("after group:", match_arr[2])
			matchList = append(matchList, Keybinding{Before: beforeGroup, After: afterGroup})
			// }
		}
		// arr := strings.SplitN(scanner.Text(), " ", 3)
		// fmt.Println("arr:", strings.Join(arr, "|"))
		// text = append(text, scanner.Text())

	}

	// fmt.Println("matchList", matchList)

	// jsonData, err := json.MarshalIndent(matchList, "", "  ")
	// if err != nil {
	// 	log.Println(err)
	// }

	// jsonData := bytes.NewBuffer([]byte{})
	// jsonEncoder := json.NewEncoder(jsonData)
	// jsonEncoder.SetEscapeHTML(false)
	// jsonEncoder.SetIndent("", "  ")
	// jsonEncoder.Encode(matchList)
	// fmt.Println("JSON：", jsonData.String())

	// fmt.Println("JSON:", string(jsonData))

	// for _, eachLine := range text {
	// 	// if
	// 	arr := strings.SplitN(eachLine, " ", 3)
	// 	// fmt.Println("arr:", arr)
	// 	fmt.Println("eachLine:", strings.Join(arr, "|"))
	// }
	file.Close()

	// fileOfVscodeVimKeybinding, err := os.Create("data.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer fileOfVscodeVimKeybinding.Close()
	// ioutil.WriteFile("keybindingOfVscodeVim.json", []byte(jsonData.String()), os.ModePerm)

	// for _, eachLine := range scanner.Text() {
	// 	typeOf := reflect.TypeOf(eachLine).Kind()
	// 	fmt.Println("typeOf:", typeOf)
	// 	// words := strings.Fields(eachLine)
	// 	fmt.Println("eachLine: ", eachLine)
	// 	// split_group := strings.SplitN(eachLine, " ", 3)
	// 	// fmt.Println("split_gropu", split_group)
	// 	// print(split_group)
	// }
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

// SplitLab | Split into three parts and
// func SplitLab() {
// 	// 用 SplitN 切成三段
// 	// input := "imap <leader>;, <C-o>$, <C-a><leader> <CR> <CR>"
// 	// print(strings.SplitN(input, " ", 3))

// 	// 用自訂的 splitFunc 切成一個一個 char
// 	input2 := "<C-o>$, <C-a><leader> <CR> <CR>"
// 	scanner := bufio.NewScanner(strings.NewReader(input2))
// 	scanner.Split(splitFunc)

// 	// 切完之後 append 到 chars、輸出
// 	var chars []string
// 	for scanner.Scan() {
// 		chars = append(chars, scanner.Text())
// 	}
// 	print(chars)
// }
//tset
