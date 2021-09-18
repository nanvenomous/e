pkgname="e"

## build: compile bt executable in current directory from source
build:
	go build -o "${pkgname}" main.go
	sudo cp ./"${pkgname}" "${GOROOT}"/bin/"${pkgname}"


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
