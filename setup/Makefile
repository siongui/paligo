# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef TRAVIS
	export GOROOT=$(realpath ../../go)
	export GOPATH=$(realpath ..)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif


WEBSITE_DIR=../website
WEBSITE_JSON_DIR=$(WEBSITE_DIR)/json

DATA_REPO_DIR=$(CURDIR)/data
DICTIONARY_DATA_DIR=$(DATA_REPO_DIR)/dictionary


default: fmt vfs_test

parsebooks: dir
	@echo "\033[92mParse Dictionary Books Information ...\033[0m"
	@go run vfssetup.go -action=parsebooks

parsewords: dir
	@echo "\033[92mParse Dictionary Words ...\033[0m"
	@go run vfssetup.go -action=parsewords

vfs:
	@echo "\033[92mBuilding Pali Words vfs...\033[0m"
	@go run vfssetup.go -action=buildvfs

vfs_clean:
	@echo "\033[92mRemoving Pali Words vfs...\033[0m"
	@rm -rf ../src/github.com/siongui/gopaliwordvfs/

vfs_test:
	@echo "\033[92mTestinging Pali Words vfs...\033[0m"
	@go run testvfs.go

dir:
	@echo "\033[92mCreate website directory if not exists ...\033[0m"
	@[ -d $(WEBSITE_JSON_DIR) ] || mkdir -p $(WEBSITE_JSON_DIR)

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go

data_clone:
	@echo "\033[92mClone Pāli data Repo ...\033[0m"
	@git clone https://github.com/siongui/data.git $(DATA_REPO_DIR) --depth=1

lib_pali:
	@echo "\033[92mInstalling common lib used in this project ...\033[0m"
	go get -u github.com/siongui/gopalilib/dicutil

lib_goef:
	@echo "\033[92mInstalling library for embedding data in Go code...\033[0m"
	go get -u github.com/siongui/goef

clean:
	@echo "\033[92mClean Repo ...\033[0m"
	@#cd ..; rm -rf bin pkg src data website
	cd ..; rm -rf bin pkg src website

install: lib_pali lib_goef
