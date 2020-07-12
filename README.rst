=======================================
`Pāli Dictionary`_ and `Pāli Tipiṭaka`_
=======================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/paligo?status.svg
   :target: https://godoc.org/github.com/siongui/paligo

.. image:: https://travis-ci.org/siongui/paligo.svg?branch=master
    :target: https://travis-ci.org/siongui/paligo

.. image:: https://gitlab.com/siongui/pali-dictionary/badges/master/pipeline.svg
    :target: https://gitlab.com/siongui/pali-dictionary/-/commits/master

.. image:: https://goreportcard.com/badge/github.com/siongui/paligo
   :target: https://goreportcard.com/report/github.com/siongui/paligo

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/paligo/blob/master/UNLICENSE

Re-implementation of `Pāli Dictionary`_ and `Pāli Tipiṭaka`_ in Go_ programming
language.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.12.17`_
  - GopherJS_

Re-implementation of `Pāli Dictionary`_ is almost finished. `Pāli Tipiṭaka`_ not
yet.

Set Up Development Environment
++++++++++++++++++++++++++++++


1. Update Ubuntu and install packages for development:

   .. code-block:: bash

     $ sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade
     $ sudo apt-get install wget make git gcc g++ gettext


2. `git clone`_ the `pali repository`_ and `data repository`_:

   .. code-block:: bash

     # create a workspace in your home directory
     $ mkdir ~/dev
     # enter workspace
     $ cd ~/dev
     # git clone paligo repository
     $ git clone https://github.com/siongui/paligo.git --depth=1
     # or clone with full depth
     #$ git clone https://github.com/siongui/paligo.git
     # git clone data repository
     $ cd ~/dev/paligo
     $ make clone_pali_data


3. Install necessary packages:

   - Go_
   - gopalilib_
   - `go-online-pali-ime`_
   - `gopherjs-i18n`_
   - `gopherjs-input-suggest`_
   - paliDataVFS_
   -  GopherJS_

   |

   .. code-block:: bash

     $ cd ~/dev/paligo
     $ make download_go
     $ make install


4. Set up data of this project:

   .. code-block:: bash

     $ make po2mo
     $ make dir
     $ make html
     $ make scss
     $ make js


5. Run development server at http://localhost:8000/

   .. code-block:: bash

     $ make devserver


Deploy to GitHub Pages
++++++++++++++++++++++

See

- `.travis.yml <.travis.yml>`_
- `setup/dicsetup.go <setup/dicsetup.go>`_ (set SiteUrl)
- `GitHub Pages Deployment - Travis CI <https://docs.travis-ci.com/user/deployment/pages/>`_
- `Environment Variables - Travis CI <https://docs.travis-ci.com/user/environment-variables/>`_


Deploy to GitLab Pages
++++++++++++++++++++++

See `.gitlab-ci.yml <.gitlab-ci.yml>`_ and
`setup/dicsetup.go <setup/dicsetup.go>`_ (set SiteUrl)


Offline Data Processing (Optional)
++++++++++++++++++++++++++++++++++

.. code-block:: bash

  # optional: parse dictionary books
  $ make parsebooks

  $ make parsewords

  # optional: convert po files to json
  $ make po2json

  # optional: build succinct trie
  $ make succinct_trie

  # optional: create VFS (embed data in front-end Go code)
  #TODO: doc to build all pali words package using goef
  #TODO: doc to embed data except pali words


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `GitHub - siongui/pali: Pāḷi Tipiṭaka and Pāḷi Dictionaries <https://github.com/siongui/pali>`_

.. [2] `siongui/data: Data files for Pāḷi Tipiṭaka, Pāḷi Dictionaries, and external libraries <https://github.com/siongui/data>`_

.. [3] `old implementation of Pāli Dictionary <http://dictionary.sutta.org/>`_

.. [4] | Home Screen Icon on Android/iPhone & PWA support
       | `website icon on android home screen - Google search <https://www.google.com/search?q=website+icon+on+android+home+screen>`_
       | `Tutorial: Home Screen Icons | Responsive Web Design Training Tutorial | Webucator <https://www.webucator.com/tutorial/developing-mobile-websites/home-screen-icons.cfm>`_
       | `pwa manifest - Google search <https://www.google.com/search?q=pwa+manifest>`_
       | `WebPageTest - Website Performance and Optimization Test <https://www.webpagetest.org/>`_
       | `Microsoft and Google team up to make PWAs better in the Play Store | by Judah Gabriel Himango | PWABuilder | Jul, 2020 | Medium <https://medium.com/pwabuilder/microsoft-and-google-team-up-to-make-pwas-better-in-the-play-store-b59710e487>`_

.. [5] | Howto SPA on GitHub Pages
       | `Add single page application support for Github pages · Issue #408 · isaacs/github · GitHub <https://github.com/isaacs/github/issues/408>`_
       | `GitHub - rafgraph/spa-github-pages: Host single page apps with GitHub Pages <https://github.com/rafgraph/spa-github-pages>`_
       | `S(GH)PA: The Single-Page App Hack For GitHub Pages — Smashing Magazine <https://www.smashingmagazine.com/2016/08/sghpa-single-page-app-hack-github-pages/>`_
       | `GitHub - dmsnell/gh-pages-404-redirect: Can I use a custom 404 handler on GitHub pages to host a routed single-page app? <https://github.com/dmsnell/gh-pages-404-redirect>`_
       | `Redirect a GitHub Pages site with this HTTP hack | Opensource.com <https://opensource.com/article/19/7/permanently-redirect-github-pages>`_
       | `javascript - Is there a configuration in Github Pages that allows you to redirect everything to index.html for a Single Page App? - Stack Overflow <https://stackoverflow.com/questions/36296012/is-there-a-configuration-in-github-pages-that-allows-you-to-redirect-everything>`_

.. [6] | `github pages symbolic link - Google search <https://www.google.com/search?q=github+pages+symbolic+link>`_
       | `Pages: allow symlinks · Issue #553 · isaacs/github · GitHub <https://github.com/isaacs/github/issues/553>`_
       | `Added .nojekyll to workaround symlink issue in GitHub Pages. Ref: isaacs/github#553 · siongui/paligo@b9fe689 · GitHub <https://github.com/siongui/paligo/commit/b9fe689770d705743a29bd33a3c7583a5c81bec1>`_

.. [7] `Bulma: Free, open source, and modern CSS framework based on Flexbox <https://bulma.io/>`_

.. [8] | One Travis CI build deploy to two repository
       | `Github deployments are broken when deploying to multiple repositories · Issue #928 · travis-ci/dpl · GitHub <https://github.com/travis-ci/dpl/issues/928>`_
       | `Deploying to Multiple Providers - Deployment - Travis CI <https://docs.travis-ci.com/user/deployment#deploying-to-multiple-providers>`_


.. _Pāli Dictionary: https://siongui.github.io/pali-dictionary/
.. _Pāli Tipiṭaka: http://tipitaka.sutta.org/
.. _Go: https://golang.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.12.17: https://golang.org/dl/
.. _git clone: https://www.google.com/search?q=git+clone
.. _pali repository: https://github.com/siongui/pali
.. _data repository: https://github.com/siongui/data
.. _UNLICENSE: https://unlicense.org/
.. _GopherJS: http://www.gopherjs.org/
.. _go-online-pali-ime: https://github.com/siongui/go-online-input-method-pali
.. _gopherjs-i18n: https://github.com/siongui/gopherjs-i18n
.. _gopherjs-input-suggest: https://github.com/siongui/gopherjs-input-suggest
.. _gopalilib: https://github.com/siongui/gopalilib
.. _paliDataVFS: https://github.com/siongui/paliDataVFS
