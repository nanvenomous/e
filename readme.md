# nvim Session runner
run nvim sessions from any subdirectory of your project

calling `e` is the equivalent of
```bash
nvim -S ./Session.nvim
```
but it will also search up the filetree for session files

# Installation

```
go install github.com/nanvenomous/e@latest
```

* default executable name in Makefile is
	> e
* default vim session file name in main.go is
	> Session.vim
