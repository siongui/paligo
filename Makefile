PRJDIR=$(CURDIR)
ifndef GITHUB_ACTIONS
ifndef GITLAB_CI
	export GOROOT=$(realpath $(PRJDIR)/go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif
endif

modinit:
	go mod init github.com/siongui/paligo

modtidy:
	#go list -m all
	go mod tidy
