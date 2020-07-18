# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
GO_VERSION=1.12.17
ifndef TRAVIS
	# set environment variables on local machine or GitLab CI
	export GOROOT=$(realpath ./go)
	export GOPATH=$(realpath .)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

WEBSITE_DIR=website
WEBSITE_JSON_DIR=$(WEBSITE_DIR)/json
WEBSITE_ABOUT_DIR=$(WEBSITE_DIR)/about
PRODUCTION_GITHUB_REPO=github.com/siongui/pali-dictionary
PRODUCTION_DIR=src/$(PRODUCTION_GITHUB_REPO)
LOCALE_DIR=locale

DATA_REPO_DIR=$(CURDIR)/data
DICTIONARY_DATA_DIR=$(DATA_REPO_DIR)/dictionary
DICTIONARY_CONF_DIR=dictionary/config/


# html must run before about_symlink. otherwise make symlink will fail
devserver: fmt dir html js about_symlink
	@# https://stackoverflow.com/a/5947779
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run devserver.go


make-gitlab-basic: rmsite dir html js about_symlink gitlab-specific
make-gitlab: rmsite dir html js symlink gitlab-specific
gitlab-specific:
	mv $(WEBSITE_DIR) public/
	echo -e 'User-agent: *\nDisallow: /' > public/robots.txt
make-dhamma-basic: rmsite dir htmldhamma js about_symlink cname-dhamma
make-dhamma: rmsite dir htmldhamma js symlink cname-dhamma
make-sutta-basic: rmsite dir htmlsutta js about_symlink cname-sutta
make-sutta: rmsite dir htmlsutta js symlink cname-sutta
rmsite:
	@echo "\033[92mRemove $(WEBSITE_DIR)\033[0m"
	rm -rf $(WEBSITE_DIR)


#################
# GitHub Deploy #
#################
printurl:
	@echo "\033[92mURL\033[0m": https://github.com/$(USERREPO)
	@echo "\033[92mHTTPS GIT\033[0m": https://github.com/$(USERREPO).git
TMPDIR=$(WEBSITE_DIR)
deploy:
	USERREPO="$(USERREPO)" make printurl
	#mv $(WEBSITE_DIR) $(TMPDIR)
	cd $(TMPDIR); git init
	cd $(TMPDIR); git add .
	cd $(TMPDIR); git commit -m "Initial commit"
	cd $(TMPDIR); git remote add origin https://github.com/$(USERREPO).git
	cd $(TMPDIR); git push --force --set-upstream origin master:gh-pages
	rm -rf $(TMPDIR)
	USERREPO="$(USERREPO)" make printurl
q-sutta:
	@USERREPO="siongui/dictionary.sutta.org" make deploy
q-dhamma:
	@USERREPO="siongui/dictionary.online-dhamma.net" make deploy
pagebuild:
	# Request a GitHub Pages build
	# https://docs.github.com/en/rest/reference/repos#pages
	# https://docs.github.com/en/rest/overview/other-authentication-methods
	@echo "\033[92mRequest a GitHub Pages build ...\033[0m"
	echo "\033[92m/repos/$(USERREPO)/pages/builds\033[0m"
	curl -u $(USER) https://api.github.com/user \
		-X POST \
		-H "Accept: application/vnd.github.v3+json" \
		https://api.github.com/repos/$(USERREPO)/pages/builds
qw-sutta:
	@USERREPO="siongui/dictionary.sutta.org" USER=$(USER) make pagebuild
qw-dhamma:
	@USERREPO="siongui/dictionary.online-dhamma.net" USER=$(USER) make pagebuild
########################
# End of GitHub Deploy #
########################

###############################
# Travis CI Custom Deployment #
###############################
# https://docs.travis-ci.com/user/deployment/custom/
# https://stackoverflow.com/questions/18935539/authenticate-with-github-using-a-token
# How do I avoid the specification of the username and password at every git push?
# https://stackoverflow.com/a/48854773
# How to set up TravisCI for projects that push back to github
# https://gist.github.com/willprice/e07efd73fb7f13f917ea
# What is /dev/null 2>&1? https://stackoverflow.com/a/10508862
travis_deploy:
	cd $(TDDIR); git init
	cd $(TDDIR); git add .
	cd $(TDDIR); git commit -m "Initial commit" --quiet
	cd $(TDDIR); git remote add origin https://siongui:$(GITHUB_TOKEN)@github.com/$(USERREPO).git > /dev/null 2>&1
	cd $(TDDIR); git push --force --set-upstream origin master:gh-pages
custom_github_pages_deploy:
	@USERREPO="siongui/dictionary.sutta.org" TDDIR="website" make travis_deploy
	@#USERREPO="siongui/dictionary.online-dhamma.net" TDDIR="website2" make deploy
######################################
# End of Travis CI Custom Deployment #
######################################


############################
# Build Dictionary Website #
############################
cname:
	@echo "\033[92mCreate CNAME for GitHub Pages custom domain ...\033[0m"
	echo "$(CNAME)" > $(WEBSITE_DIR)/CNAME
cname-dhamma:
	CNAME=dictionary.online-dhamma.net make cname
cname-sutta:
	CNAME=dictionary.sutta.org make cname

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@#gopherjs build gopherjs/*.go -o $(WEBSITE_DIR)/pali.js ; exit
	@gopherjs build gopherjs/*.go -m -o $(WEBSITE_DIR)/pali.js

htmlsutta:
	@echo "\033[92mGenerating HTML for dictionary.sutta.org ...\033[0m"
	go run dictionary/htmlspa.go -siteconf="$(DICTIONARY_CONF_DIR)/dictionary.sutta.org.json" -pathconf="$(DICTIONARY_CONF_DIR)/path-for-build.json"

htmldhamma:
	@echo "\033[92mGenerating HTML for dictionary.online-dhamma.net ...\033[0m"
	go run dictionary/htmlspa.go -siteconf="$(DICTIONARY_CONF_DIR)/dictionary.online-dhamma.net.json" -pathconf="$(DICTIONARY_CONF_DIR)/path-for-build.json"

html:
	@echo "\033[92mGenerating HTML ...\033[0m"
ifdef GITLAB_CI
	go run dictionary/htmlspa.go -siteconf="$(DICTIONARY_CONF_DIR)/siongui.gitlab.io-pali-dictionary.json" -pathconf="$(DICTIONARY_CONF_DIR)/path-for-build.json"
else
	@go run dictionary/htmlspa.go -siteconf="$(DICTIONARY_CONF_DIR)/empty-siteurl.json" -pathconf="$(DICTIONARY_CONF_DIR)/path-for-build.json"
endif

nojekyll: dir
	@echo "\033[92mMaking symbolic links works on GitHub Pages ...\033[0m"
	@touch $(WEBSITE_DIR)/.nojekyll
about_symlink: nojekyll
	@echo "\033[92mMaking symbolic link for about page ...\033[0m"
	@cd $(WEBSITE_ABOUT_DIR); [ -f index.html ] || ln -s ../index.html index.html
symlink: about_symlink
	@echo "\033[92mMaking symbolic link for static website ...\033[0m"
	go run dictionary/dicsetup.go -action=symlink -pathconf="$(DICTIONARY_CONF_DIR)/path-for-build.json"

dir:
	@echo "\033[92mCreate website directory if not exists ...\033[0m"
	@[ -d $(WEBSITE_JSON_DIR) ] || mkdir -p $(WEBSITE_JSON_DIR)
	@[ -d $(WEBSITE_ABOUT_DIR) ] || mkdir -p $(WEBSITE_ABOUT_DIR)
###################################
# End of Build Dictionary Website #
###################################


###################################################
# Boostrap Website (only once, not daily routine) #
###################################################
parsebooks: dir
	@echo "\033[92mParse Dictionary Books Information ...\033[0m"
	@go run dictionary/dicsetup.go -action=parsebooks

parsewords: dir
	@echo "\033[92mParse Dictionary Words ...\033[0m"
	@go run dictionary/dicsetup.go -action=parsewords

po2json:
	@echo "\033[92mConverting PO files to JSON (to be used in client-side/browser) ...\033[0m"
	@go run dictionary/dicsetup.go -action=po2json

succinct_trie:
	@echo "\033[92mBuilding Succinct Trie ...\033[0m"
	@go run dictionary/dicsetup.go -action=triebuild

twpo2cn:
	@echo "\033[92mConverting zh_TW PO files to zh_CN ...\033[0m"
	@#FIXME: go run setup/twpo2cn.go -tw=$(LOCALE_DIR)/zh_TW/LC_MESSAGES/messages.po -cn=$(LOCALE_DIR)/zh_CN/LC_MESSAGES/messages.po

po2mo:
	# not used now. keep here for references
	@echo "\033[92mmsgfmt PO to MO ...\033[0m"
	msgfmt $(LOCALE_DIR)/zh_TW/LC_MESSAGES/messages.po -o $(LOCALE_DIR)/zh_TW/LC_MESSAGES/messages.mo
	#@msgfmt $(LOCALE_DIR)/zh_CN/LC_MESSAGES/messages.po -o $(LOCALE_DIR)/zh_CN/LC_MESSAGES/messages.mo
	msgfmt $(LOCALE_DIR)/vi_VN/LC_MESSAGES/messages.po -o $(LOCALE_DIR)/vi_VN/LC_MESSAGES/messages.mo
	msgfmt $(LOCALE_DIR)/fr_FR/LC_MESSAGES/messages.po -o $(LOCALE_DIR)/fr_FR/LC_MESSAGES/messages.mo
###########################
# End of Boostrap Website #
###########################


###################
# Install Library #
###################
install: lib_pali lib_gtmpl lib_ime_pali lib_gopherjs_i18n lib_gopherjs_input_suggest lib_paliDataVFS lib_gopherjs

lib_pali:
	@echo "\033[92mInstalling common lib used in this project ...\033[0m"
	go get -u github.com/siongui/gopalilib/dicutil

lib_gtmpl:
	@echo "\033[92mInstalling Go html/template with gettext support ...\033[0m"
	go get -u github.com/siongui/gtmpl

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
##########################
# End of Install Library #
##########################


########
# Misc #
########
fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt dictionary/*.go
	@go fmt gopherjs/*.go
	@go fmt *.go

clone_pali_data:
	@echo "\033[92mClone Pāli data Repo ...\033[0m"
	@git clone https://github.com/siongui/data.git $(DATA_REPO_DIR) --depth=1

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
