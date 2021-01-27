## The real working translator for vscodevim keybinding
#### Since I've tried to use the setting which vscodevim has provided that you can import your init.vim, but it turned out to be failed on converting all my keybinding from the insert mode. Therefore, I created this little tool for myself. 

#### convert your keybindings from vim into the format of vscodevim keybinding 

![demo1](https://i.imgur.com/yQz3IGO.png)

![demo2](https://i.imgur.com/Tjl7ctM.gif)
OS Platform Support for Windows, Linux and OSX
### **Usage as below:**
### Go
First of all, clone this repo: <br>
```git clone https://github.com/wizenith/vscode_keybinding_from_vim.git``` <br>
switch into this directory and then: <br>
```go run . --filepath [specify the path of your vim keybinding config file]``` <br>
### Windows
```vimtovscodekeybinding.exe -f [specify the path of your vim keybinding config file] ```<br>
**or**<br>
```vimtovscodekeybinding.exe --filepath [specify the path of your vim keybinding config file]```<br>
### Linux
```vimtovscodekeybinding -f [specify the path of your vim keybinding config file] ```<br>
**or**<br>
```vimtovscodekeybinding --filepath [specify the path of your vim keybinding config file]```<br>
### OSX
```vimtovscodekeybinding_osx -f [specify the path of your vim keybinding config file] ```<br>
**or**<br>
```vimtovscodekeybinding_osx --filepath [specify the path of your vim keybinding config file]```<br>

## you can download the binary from [this page](https://github.com/wizenith/vscode_keybinding_from_vim/releases) 



