ifndef GOROOT
	export GOROOT=$(realpath ../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif

modinit:
	go mod init github.com/siongui/paligo

modtidy:
	#go list -m all
	go mod tidy
