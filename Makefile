
BIN_DIR=bin/
create_bin:
	mkdir -p ${BIN_DIR}

all: install

.PHONY: install
install: create_bin
	go build -o ${BIN_DIR}/tutup-botol

.PHONY: build-win
build-win:
	GOOS=windows go build -o ${BIN_DIR}/tutup-botol.exe

.PHONY: test
test:
	go test

.PHONY: clean
clean:
	rm -f ${BIN_DIR}/tutup-botol ${BIN_DIR}/tutup-botol.exe


