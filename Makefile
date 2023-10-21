ifndef GOROOT
	export GOROOT=$(realpath $(CURDIR)/go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif

GO_VERSION=1.18.10

modinit:
	go mod init github.com/siongui/paligo

modtidy:
	#go list -m all
	go mod tidy

libupdate:
	go get -u github.com/siongui/gopalilib

install_gopherjs:
	@echo "\033[92mInstalling GopherJS ...\033[0m"
	go install github.com/gopherjs/gopherjs@v1.18.0-beta3

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@[ -f go$(GO_VERSION).linux-amd64.tar.gz ] || wget https://dl.google.com/go/go$(GO_VERSION).linux-amd64.tar.gz
	@tar -xvzf go$(GO_VERSION).linux-amd64.tar.gz
	@#rm go$(GO_VERSION).linux-amd64.tar.gz
