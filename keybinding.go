package main

import (
	"bytes"
	"encoding/json"
)

//Keybinding | type of vscodevim keybinding
type Keybinding struct {
	Before []string `json:"before"`
	After  []string `json:"after"`
}

//KeybindingsOfAllModes | all type of vim modes
type KeybindingsOfAllModes struct {
	Normal   []Keybinding `json:"vim.normalModeKeyBindings"`
	Nnoremap []Keybinding `json:"vim.normalModeKeyBindingsNonRecursive"`
	Insert   []Keybinding `json:"vim.insertModeKeyBindings"`
	Inoremap []Keybinding `json:"vim.insertModeKeyBindingsNonRecursive"`
	Visual   []Keybinding `json:"vim.visualModeKeyBindings"`
	Vnoremap []Keybinding `json:"vim.visualModeKeyBindingsNonRecursive"`
}

// implement Stringer interface for KeybindingsOfAllModes
func (k KeybindingsOfAllModes) String() string {
	jsonData := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(jsonData)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "  ")
	jsonEncoder.Encode(k)
	return jsonData.String()
}
