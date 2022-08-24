pkgname="e"

build:
	go build -o "${pkgname}" main.go

install:
	mv ./"${pkgname}" "${GOROOT}"/bin/"${pkgname}"
	rm -f /usr/bin/vi
	ln -s "${GOROOT}"/bin/"${pkgname}" /usr/bin/vi

visualize:
	go-callvis -focus "cmd" -group pkg,type -ignore "github.com/spf13,os,fmt,errors,path,strings,strconv"  ./...


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
