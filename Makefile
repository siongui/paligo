# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH

# Complex conditions check in Makefile
# https://stackoverflow.com/a/5586785
ifndef_any_of = $(filter undefined,$(foreach v,$(1),$(origin $(v))))
ifdef_any_of = $(filter-out undefined,$(foreach v,$(1),$(origin $(v))))

GO_VERSION=1.12.17
ifndef TRAVIS
	export GOROOT=$(realpath ./go)
	export GOPATH=$(realpath .)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

SCSS_DIR=theme/styling
SCSS_PATH=$(SCSS_DIR)/style.scss
WEBSITE_DIR=website
WEBSITE_JSON_DIR=$(WEBSITE_DIR)/json
PRODUCTION_GITHUB_REPO=github.com/siongui/pali-dictionary
PRODUCTION_DIR=src/$(PRODUCTION_GITHUB_REPO)
LOCALE_DIR=locale

DATA_REPO_DIR=$(CURDIR)/data
DICTIONARY_DATA_DIR=$(DATA_REPO_DIR)/dictionary


devserver: fmt html js scss
	@# http://stackoverflow.com/a/5947779
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run server.go

cname:
	@echo "\033[92mCreate CNAME for GitHub Pages custom domain ...\033[0m"
	@echo "dictionary.online-dhamma.net" > $(WEBSITE_DIR)/CNAME

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
# ifdef TRAVIS || GITLAB_CI
ifneq ($(call ifdef_any_of,TRAVIS GITLAB_CI),)
	@gopherjs build gopherjs/*.go -m -o $(WEBSITE_DIR)/pali.js
else
	@gopherjs build gopherjs/*.go -o $(WEBSITE_DIR)/pali.js
endif

scss:
	@echo "\033[92mGenerating CSS ...\033[0m"
	@go run setup/scss.go -scsspath=$(SCSS_PATH) -scssdir=$(SCSS_DIR) -csspath=$(WEBSITE_DIR)/style.css

html:
	@echo "\033[92mGenerating HTML ...\033[0m"
	@# Google Search: shell stdout to file
# ifdef TRAVIS || GITLAB_CI
ifneq ($(call ifdef_any_of,TRAVIS GITLAB_CI),)
	@go run setup/dicsetup.go -action=html > $(WEBSITE_DIR)/index.html
else
	@go run setup/dicsetup.go -action=html -isdev=true > $(WEBSITE_DIR)/index.html
endif


parsebooks: dir
	@echo "\033[92mParse Dictionary Books Information ...\033[0m"
	@go run setup/dicsetup.go -action=parsebooks

parsewords: dir
	@echo "\033[92mParse Dictionary Words ...\033[0m"
	@go run setup/dicsetup.go -action=parsewords

po2json:
	@echo "\033[92mConverting PO files to JSON (to be used in client-side/browser) ...\033[0m"
	@go run setup/dicsetup.go -action=po2json

succinct_trie:
	@echo "\033[92mBuilding Succinct Trie ...\033[0m"
	@go run setup/dicsetup.go -action=triebuild

symlink:
	@echo "\033[92mMaking symbolic link for static website ...\033[0m"
	@echo "" > $(WEBSITE_DIR)/.nojekyll
	@go run setup/dicsetup.go -action=symlink

dir:
	@echo "\033[92mCreate website directory if not exists ...\033[0m"
	@[ -d $(WEBSITE_JSON_DIR) ] || mkdir -p $(WEBSITE_JSON_DIR)

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt setup/*.go
	@go fmt gopherjs/*.go
	@go fmt server.go

clone_pali_data:
	@echo "\033[92mClone Pāli data Repo ...\033[0m"
	@git clone https://github.com/siongui/data.git $(DATA_REPO_DIR) --depth=1


install: lib_pali lib_go_libsass lib_ime_pali lib_gopherjs_i18n lib_gopherjs_input_suggest lib_paliDataVFS lib_gopherjs install_gettext

lib_pali:
	@echo "\033[92mInstalling common lib used in this project ...\033[0m"
	go get -u github.com/siongui/gopalilib/dicutil

lib_go_libsass: lib_normalize_css
	@echo "\033[92mInstalling libsass and its Go binding ...\033[0m"
	go get -u github.com/wellington/go-libsass

lib_normalize_css:
	@echo "\033[92mInstalling Normalize.css ...\033[0m"
	@[ -e $(SCSS_DIR)/_normalize801.scss ] || wget -O $(SCSS_DIR)/_normalize801.scss https://necolas.github.io/normalize.css/8.0.1/normalize.css

lib_ime_pali:
	@echo "\033[92mInstalling Online Go Pāli IME ...\033[0m"
	go get -u github.com/siongui/go-online-input-method-pali

lib_gopherjs_i18n:
	@echo "\033[92mInstalling GopherJS gettext library (online/client-side)...\033[0m"
	go get -u github.com/siongui/gopherjs-i18n

lib_paliDataVFS:
	@echo "\033[92mInstalling VFS for fullstack Go ...\033[0m"
	go get -u github.com/siongui/paliDataVFS

lib_gopherjs_input_suggest:
	@echo "\033[92mInstalling GopherJS input suggest library ...\033[0m"
	go get -u github.com/siongui/gopherjs-input-suggest

lib_gopherjs:
	@echo "\033[92mInstalling GopherJS ...\033[0m"
	go get -u github.com/gopherjs/gopherjs

install_gettext:
	@echo "\033[92mInstalling GNU Internationalization utilities ...\033[0m"
	sudo apt-get -y install gettext

twpo2cn:
	@echo "\033[92mConverting zh_TW PO files to zh_CN ...\033[0m"
	@go run setup/twpo2cn.go -tw=$(LOCALE_DIR)/zh_TW/LC_MESSAGES/messages.po -cn=$(LOCALE_DIR)/zh_CN/LC_MESSAGES/messages.po

po2mo:
	@echo "\033[92mmsgfmt PO to MO ...\033[0m"
	msgfmt $(LOCALE_DIR)/zh_TW/LC_MESSAGES/messages.po -o $(LOCALE_DIR)/zh_TW/LC_MESSAGES/messages.mo
	#@msgfmt $(LOCALE_DIR)/zh_CN/LC_MESSAGES/messages.po -o $(LOCALE_DIR)/zh_CN/LC_MESSAGES/messages.mo
	msgfmt $(LOCALE_DIR)/vi_VN/LC_MESSAGES/messages.po -o $(LOCALE_DIR)/vi_VN/LC_MESSAGES/messages.mo
	msgfmt $(LOCALE_DIR)/fr_FR/LC_MESSAGES/messages.po -o $(LOCALE_DIR)/fr_FR/LC_MESSAGES/messages.mo

clean:
	@echo "\033[92mClean Repo ...\033[0m"
	@#rm -rf bin pkg src data $(WEBSITE_DIR)
	rm -rf bin pkg src $(WEBSITE_DIR)

update_ubuntu:
	@echo "\033[92mUpdating Ubuntu ...\033[0m"
	@sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@#wget https://storage.googleapis.com/golang/go$(GO_VERSION).linux-amd64.tar.gz
	@wget https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz
	@tar -xvzf go$(GO_VERSION).linux-amd64.tar.gz
	@rm go$(GO_VERSION).linux-amd64.tar.gz
