.PHONY: all build run gotool clean help

LINUX_BINARY="vimtovscode"
WIN_BINARY="vimtovscode.exe"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${LINUX_BINARY}
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${WIN_BINARY}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${LINUX_BINARY} ] ; then rm ${LINUX_BINARY} ; fi
	@if [ -f ${WIN_BINARY} ] ; then rm ${WIN_BINARY} ; fi

help:
	@echo "make - go fmt and go vet, then complie file to binary"
	@echo "make build - compile"
	@echo "make run - go run ."
	@echo "make clean - remove binary "
	@echo "make gotool -run go 'fmt' and 'vet'"
