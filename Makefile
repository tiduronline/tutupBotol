.PHONY: install
install:
	go build -o tutup-botol

.PHONY: test
test:
	go test

.PHONY: clean
clean:
	rm -f tutup-botol

