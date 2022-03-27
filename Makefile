pkgname="e"

## build: compile bt executable in current directory from source
install:
	go build -o "${pkgname}" main.go
	mv ./"${pkgname}" "${GOROOT}"/bin/"${pkgname}"
	rm -f /usr/bin/vi
	ln -s "${GOROOT}"/bin/"${pkgname}" /usr/bin/vi


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
