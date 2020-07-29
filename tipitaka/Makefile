# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
GO_VERSION=1.12.17
PRJDIR=$(CURDIR)/../
LOCALE_DIR=$(PRJDIR)/locale
ifndef TRAVIS
	# set environment variables on local machine or GitLab CI
	export GOROOT=$(realpath $(PRJDIR)/go)
	export GOPATH=$(realpath $(PRJDIR))
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

HTML_TEMPLATE_DIR=$(CURDIR)/theme/template/
EXTRA_DIR=$(CURDIR)/theme/extra/
WEBSITE_DIR=$(CURDIR)/website
WEBSITE_DATA_DIR=$(WEBSITE_DIR)/data
TPK_BUILD_OPTIONS=-websiteDir=$(WEBSITE_DIR) -htmlTemplateDir=$(HTML_TEMPLATE_DIR) -localeDir=$(LOCALE_DIR)
TIPITAKA_CONF_DIR=$(CURDIR)/config/
TIPITAKA_XML_DIR=$(WEBSITE_DIR)/tpkxml/


devserver: make-local
	@# https://stackoverflow.com/a/5947779
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run $(PRJDIR)/devserver.go


make-local: fmt dir extra js html
make-sutta: rmsite dir extra js htmlsutta cname-sutta
rmsite:
	@echo "\033[92mRemove $(WEBSITE_DIR)\033[0m"
	rm -rf $(WEBSITE_DIR)


##########################
# Build Tipiṭaka Website #
##########################
cname:
	@echo "\033[92mCreate CNAME for GitHub Pages custom domain ...\033[0m"
	echo "$(CNAME)" > $(WEBSITE_DIR)/CNAME
cname-sutta:
	CNAME=tipitaka.sutta.org make cname

extra:
	@echo "\033[92mMoving extra items to website dir ...\033[0m"
	@cp -r $(EXTRA_DIR)/* $(WEBSITE_DIR)

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@#gopherjs build gopherjs/*.go -o $(WEBSITE_DIR)/pali.js
	@gopherjs build gopherjs/*.go -m -o $(WEBSITE_DIR)/pali.js

htmlsutta:
	@echo "\033[92mGenerating HTML for dictionary.sutta.org ...\033[0m"
	go run htmlspa.go -siteconf="$(TIPITAKA_CONF_DIR)/tipitaka.sutta.org.json" $(TPK_BUILD_OPTIONS)
html:
	@echo "\033[92mGenerating HTML ...\033[0m"
	@go run htmlspa.go -siteconf="$(TIPITAKA_CONF_DIR)/empty-siteurl.json" $(TPK_BUILD_OPTIONS)

dir:
	@echo "\033[92mCreate website directory if not exists ...\033[0m"
	@[ -d $(WEBSITE_DIR) ] || mkdir -p $(WEBSITE_DIR)
#################################
# End of Build Tipiṭaka Website #
#################################


##############################
# Bootstrap Tipiṭaka Website #
##############################
clone_tpk_xml:
	@echo "\033[92mClone Tipiṭaka XML repo ...\033[0m"
	@[ -d $(TIPITAKA_XML_DIR) ] || git clone https://github.com/siongui/tipitaka-romn.git $(TIPITAKA_XML_DIR)

embed_data: fmt clone_tpk_xml
	@echo "\033[92mEmbed Tipiṭaka ToC(Table of Content) in Go code ...\033[0m"
	@[ -d $(WEBSITE_DATA_DIR) ] || mkdir -p $(WEBSITE_DATA_DIR)
	@go run embed.go -xmlDir=$(TIPITAKA_XML_DIR) -dataDir=$(WEBSITE_DATA_DIR)
#####################################
# End of Bootstrap Tipiṭaka Website #
#####################################


###################
# Install Library #
###################
install: lib_pali lib_gtmpl lib_gopherjs_i18n lib_godom lib_paliDataVFS lib_gopherjs lib_goef

lib_pali:
	@echo "\033[92mInstalling common lib used in this project ...\033[0m"
	go get -u github.com/siongui/gopalilib/util

lib_gtmpl:
	@echo "\033[92mInstalling Go html/template with gettext support ...\033[0m"
	go get -u github.com/siongui/gtmpl

lib_gopherjs_i18n:
	@echo "\033[92mInstalling GopherJS gettext library (online/client-side)...\033[0m"
	go get -u github.com/siongui/gopherjs-i18n

lib_godom:
	@echo "\033[92mInstalling dom manipulation library ...\033[0m"
	go get -u github.com/siongui/godom

lib_paliDataVFS:
	@echo "\033[92mInstalling VFS for fullstack Go ...\033[0m"
	go get -u github.com/siongui/paliDataVFS

lib_gopherjs:
	@echo "\033[92mInstalling GopherJS ...\033[0m"
	go get -u github.com/gopherjs/gopherjs

lib_goef:
	@echo "\033[92mInstalling GopherJS ...\033[0m"
	go get -u github.com/siongui/goef
##########################
# End of Install Library #
##########################


########
# Misc #
########
fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt $(PRJDIR)/*.go
	@go fmt *.go
	@go fmt gopherjs/*.go

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
	@mv go/ $(PRJDIR)