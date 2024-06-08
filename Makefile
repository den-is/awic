.PHONY: run linux_amd64 macos_amd64 macos_arm win_amd64 install

run:
	go run awic.go -i c5n.9xlarge

linux_amd64:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -ldflags "-w" -o bin/awic_linux_amd64

macos_amd64:
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build -ldflags "-w" -o bin/awic_macos_amd64

macos_arm:
	mkdir -p bin
	GOOS=darwin GOARCH=arm64 go build -ldflags "-w" -o bin/awic_macos_arm64

win_amd64:
	mkdir -p bin
	GOOS=windows GOARCH=amd64 go build -ldflags "-w" -o bin/awic_win_amd64
