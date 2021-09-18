# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
GO_VERSION=1.17.1
PRJDIR=$(CURDIR)/../
LOCALE_DIR=$(PRJDIR)/locale
ifndef GITHUB_ACTIONS
	# set environment variables on local machine or GitLab CI
	export GOROOT=$(realpath $(PRJDIR)/go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif

HTML_TEMPLATE_DIR=$(CURDIR)/theme/template/
EXTRA_DIR=$(CURDIR)/theme/extra/
SUPPORTED_LOCALES="en_US,zh_TW,vi_VN,fr_FR"
WEBSITE_DIR=$(CURDIR)/website
WEBSITE_ABOUT_DIR=$(WEBSITE_DIR)/about
SYMLINK_OPTIONS=-websiteDir=$(WEBSITE_DIR) -supportedLocales=$(SUPPORTED_LOCALES)
DIC_BUILD_OPTIONS=$(SYMLINK_OPTIONS) -htmlTemplateDir=$(HTML_TEMPLATE_DIR) -localeDir=$(LOCALE_DIR)

DATA_REPO_DIR=$(PRJDIR)/data
DICTIONARY_BOOTSTRAP_DATA_DIR=$(DATA_REPO_DIR)/dictionary
DICTIONARY_CONF_DIR=$(CURDIR)/config/

OFFLINE_LINUX=$(GOROOT)/bin/offline-linux
OFFLINE_WINDOWS=$(GOROOT)/bin/offline-windows.exe
OFFLINE_MAC=$(GOROOT)/bin/offline-darwin
GITLAB_PUBLIC_DIR=$(PRJDIR)/public/


# html must run before about_symlink. otherwise make symlink will fail
devserver: make-local-basic
	@# https://stackoverflow.com/a/5947779
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run $(PRJDIR)/devserver.go


make-common: dir js extra
make-local-basic: fmt make-common html about_symlink
make-local: fmt make-common html symlink
make-gitlab-basic: rmsite make-common htmlgitlab about_symlink
make-gitlab: rmsite make-common htmlgitlab symlink
make-dhamma-basic: rmsite make-common htmldhamma about_symlink cname-dhamma
make-dhamma: rmsite make-common htmldhamma symlink cname-dhamma
make-sutta-basic: rmsite make-common htmlsutta about_symlink cname-sutta
make-sutta: rmsite make-common htmlsutta symlink cname-sutta
rmsite:
	@echo "\033[92mRemove $(WEBSITE_DIR)\033[0m"
	rm -rf $(WEBSITE_DIR)
mvtopublic:
	mv $(WEBSITE_DIR) $(GITLAB_PUBLIC_DIR)


###############################
# Build Single Offline Binary #
###############################
clone-word-json:
	@echo "\033[92mClone Pali words JSON files ...\033[0m"
	[ -d $(WEBSITE_DIR)/json/ ] || git clone -b master https://github.com/siongui/xemaauj9k5qn34x88m4h.git $(WEBSITE_DIR)/json/

bulma-offline:
	@echo "\033[92mMaking offline bulma.css ...\033[0m"
	[ -f $(WEBSITE_DIR)/bulma.css ] || wget https://raw.githubusercontent.com/jgthms/bulma/0.9.0/css/bulma.css -O $(WEBSITE_DIR)/bulma.css
	find $(WEBSITE_DIR) -name "index.html" -exec sed -i -e 's/https:\/\/cdn.jsdelivr.net\/npm\/bulma@0.9.0\/css//' {} \;
	find $(WEBSITE_DIR) -name "index.html" -exec sed -i -e 's/ integrity="sha256-oSsWW\/ca1ll8OVu2t4EKkyb9lmH0Pkd8CXKXnYP4QvA=" crossorigin="anonymous"//' {} \;

offline: make-local-basic clone-word-json bulma-offline
	@echo "\033[92mBuild single offline website binary ...\033[0m"
	go run offline/embed.go -websiteDir=$(WEBSITE_DIR)
	# build binary for linux
	cd offline; go build -o $(OFFLINE_LINUX)
	# build binary for windows
	cd offline; GOOS=windows GOARCH=amd64 go build -o $(OFFLINE_WINDOWS)
	# build binary for Mac OS X
	cd offline; GOOS=darwin GOARCH=amd64 go build -o $(OFFLINE_MAC)

run-offline: offline
	$(OFFLINE_LINUX)

mvofflinetopublic:
	[ -d $(GITLAB_PUBLIC_DIR) ] || mkdir -p $(GITLAB_PUBLIC_DIR)
	mv $(OFFLINE_LINUX) $(GITLAB_PUBLIC_DIR)
	mv $(OFFLINE_WINDOWS) $(GITLAB_PUBLIC_DIR)
	mv $(OFFLINE_MAC) $(GITLAB_PUBLIC_DIR)
######################################
# End of Build Single Offline Binary #
######################################


##########################
# Local Deploy to GitHub #
##########################
# version control - How to reset a remote Git repository to remove all commits?
# https://stackoverflow.com/a/2006252
# git - Push local master commits to remote branch
# https://stackoverflow.com/a/3206144
printurl:
	@echo "\033[92mURL\033[0m": https://github.com/$(USERREPO)
	@echo "\033[92mHTTPS GIT\033[0m": https://github.com/$(USERREPO).git
local_deploy_to_github:
	USERREPO="$(USERREPO)" make printurl
	cd $(TMPDIR); git init
	cd $(TMPDIR); git add .
	cd $(TMPDIR); git commit -m "Initial commit"
	cd $(TMPDIR); git remote add origin https://github.com/$(USERREPO).git
	cd $(TMPDIR); git push --force --set-upstream origin master:gh-pages
	rm -rf $(TMPDIR)
	USERREPO="$(USERREPO)" make printurl
q-sutta:
	@USERREPO="siongui/dictionary.sutta.org" TMPDIR=$(WEBSITE_DIR) make local_deploy_to_github
q-dhamma:
	@USERREPO="siongui/dictionary.online-dhamma.net" TMPDIR=$(WEBSITE_DIR) make local_deploy_to_github
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
rb-sutta:
	@USERREPO="siongui/dictionary.sutta.org" USER=$(USER) make pagebuild
rb-dhamma:
	@USERREPO="siongui/dictionary.online-dhamma.net" USER=$(USER) make pagebuild
#################################
# End of Local Deploy to GitHub #
#################################

#########################################
# Travis CI Custom Deployment to GitHub #
#########################################
# https://docs.travis-ci.com/user/deployment/custom/
# https://stackoverflow.com/questions/18935539/authenticate-with-github-using-a-token
# How do I avoid the specification of the username and password at every git push?
# https://stackoverflow.com/a/48854773
# How to set up TravisCI for projects that push back to github
# https://gist.github.com/willprice/e07efd73fb7f13f917ea
# What is /dev/null 2>&1? https://stackoverflow.com/a/10508862
travis_deploy_to_github:
	cd $(TDDIR); git init
	# If "git all ." at once may cause 10min no output timeout
	# To prevent timeout, git add sub-dir and echo something one by one
	# https://stackoverflow.com/questions/13897945/wildcard-to-obtain-list-of-all-directories
	# https://stackoverflow.com/questions/17834582/run-make-in-each-subdirectory
	# http://kirste.userpage.fu-berlin.de/chemnet/use/info/make/make_8.html
	cd $(TDDIR); for subdir in $(subst $(TDDIR)/,,$(wildcard $(TDDIR)/*/.)); do \
		echo $$subdir ; \
		git add $$subdir ; \
	done
	cd $(TDDIR); git add .
	# --quient is to prevent exceeded max log length on Travis CI
	cd $(TDDIR); git commit -m "Initial commit" --quiet
	# > /dev/null 2>&1 is for security.
	cd $(TDDIR); git remote add origin https://siongui:$(GITHUB_TOKEN)@github.com/$(USERREPO).git > /dev/null 2>&1
	cd $(TDDIR); git push --force --set-upstream origin master:gh-pages
custom_sutta_deploy:
	@USERREPO="siongui/dictionary.sutta.org" TDDIR=$(WEBSITE_DIR) make travis_deploy_to_github
custom_dhamma_deploy:
	@USERREPO="siongui/dictionary.online-dhamma.net" TDDIR=$(WEBSITE_DIR) make travis_deploy_to_github
################################################
# End of Travis CI Custom Deployment to GitHub #
################################################


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

extra:
	@echo "\033[92mMoving extra items to website dir ...\033[0m"
	@cp -r $(EXTRA_DIR)/* $(WEBSITE_DIR)
ifeq ($(TARGET_WEBSITE), sutta)
	# allow search indexing only on dictionary.sutta.org
	mv $(WEBSITE_DIR)/robots-all.txt $(WEBSITE_DIR)/robots.txt
endif
ifeq ($(TARGET_WEBSITE), dhamma)
	mv $(WEBSITE_DIR)/manifest-dhamma.json $(WEBSITE_DIR)/manifest.json
endif
ifdef GITLAB_CI
	mv $(WEBSITE_DIR)/manifest-gitlab.json $(WEBSITE_DIR)/manifest.json
endif

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@#gopherjs build gopherjs/*.go -o $(WEBSITE_DIR)/pali.js
	@gopherjs build gopherjs/*.go -m -o $(WEBSITE_DIR)/pali.js

htmlsutta:
	@echo "\033[92mGenerating HTML for dictionary.sutta.org ...\033[0m"
	go run htmlspa.go -siteconf="$(DICTIONARY_CONF_DIR)/dictionary.sutta.org.json" $(DIC_BUILD_OPTIONS)

htmldhamma:
	@echo "\033[92mGenerating HTML for dictionary.online-dhamma.net ...\033[0m"
	go run htmlspa.go -siteconf="$(DICTIONARY_CONF_DIR)/dictionary.online-dhamma.net.json" $(DIC_BUILD_OPTIONS)

htmlgitlab:
	@echo "\033[92mGenerating HTML for siongui.gitlab.io/pali-dictionary ...\033[0m"
	go run htmlspa.go -siteconf="$(DICTIONARY_CONF_DIR)/siongui.gitlab.io-pali-dictionary.json" $(DIC_BUILD_OPTIONS)

html:
	@echo "\033[92mGenerating HTML ...\033[0m"
	@go run htmlspa.go -siteconf="$(DICTIONARY_CONF_DIR)/empty-siteurl.json" $(DIC_BUILD_OPTIONS)

nojekyll: dir
	@echo "\033[92mMaking symbolic links works on GitHub Pages ...\033[0m"
	@touch $(WEBSITE_DIR)/.nojekyll
about_symlink: nojekyll
	@echo "\033[92mMaking symbolic link for about page ...\033[0m"
	@cd $(WEBSITE_ABOUT_DIR); [ -f index.html ] || ln -s ../index.html index.html
symlink: about_symlink
	@echo "\033[92mMaking symbolic link for static website ...\033[0m"
	go run dicsetup.go -action=symlink $(SYMLINK_OPTIONS)

dir:
	@echo "\033[92mCreate website directory if not exists ...\033[0m"
	@[ -d $(WEBSITE_ABOUT_DIR) ] || mkdir -p $(WEBSITE_ABOUT_DIR)
###################################
# End of Build Dictionary Website #
###################################


####################
# gettext workflow #
####################
pot:
	@echo "\033[92mExtract translatable strings from html ...\033[0m"
	@# sed convert {{gettext "Pāḷi Dictionary"}} to {{gettext("Pāḷi Dictionary")}}
	@# https://siongui.github.io/2016/01/19/xgettext-extract-translatable-string-from-go-html-template/
	@sed "s/{{gettext \(".*"\)}}/{{gettext(\1)}}/g" `find $(HTML_TEMPLATE_DIR) -name "*.html"` | xgettext --no-wrap --language=c --from-code=UTF-8 --output=$(LOCALE_DIR)/messages.pot -
###########################
# End of gettext workflow #
###########################


###################################################
# Boostrap Website (only once, not daily routine) #
###################################################
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
install_gopherjs:
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
	@go fmt $(PRJDIR)/*.go
	@go fmt gopherjs/*.go
	@go fmt *.go
	@go fmt `find offline/ -name "*.go" | grep -v data.go`

clone_pali_data:
	@echo "\033[92mClone Pāli data Repo ...\033[0m"
	@[ -d $(DATA_REPO_DIR) ] || git clone https://github.com/siongui/data.git $(DATA_REPO_DIR) --depth=1

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
