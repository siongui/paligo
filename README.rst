=======================================
`Pāli Dictionary`_ and `Pāli Tipiṭaka`_
=======================================

.. image:: https://travis-ci.org/siongui/paligo.svg?branch=master
    :target: https://travis-ci.org/siongui/paligo

.. image:: https://gitlab.com/siongui/pali-dictionary/badges/master/pipeline.svg
    :target: https://gitlab.com/siongui/pali-dictionary/-/commits/master

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


1. `git clone`_ the `pali repository`_ and `data repository`_:

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


2. Install necessary packages:

   - Go_
   - gopalilib_
   - `go-libsass`_
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


3. Set up data of this project:

   .. code-block:: bash

     $ make po2mo
     $ make dir
     $ make html
     $ make scss
     $ make js


4. Run development server at http://localhost:8000/

   .. code-block:: bash

     $ make devserver


Deploy to GitHub Pages
++++++++++++++++++++++

See

- `.travis.yml <.travis.yml>`_
- `GitHub Pages Deployment - Travis CI <https://docs.travis-ci.com/user/deployment/pages/>`_
- `Environment Variables - Travis CI <https://docs.travis-ci.com/user/environment-variables/>`_


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
.. _go-libsass: https://github.com/wellington/go-libsass
.. _go-online-pali-ime: https://github.com/siongui/go-online-input-method-pali
.. _gopherjs-i18n: https://github.com/siongui/gopherjs-i18n
.. _gopherjs-input-suggest: https://github.com/siongui/gopherjs-input-suggest
.. _gopalilib: https://github.com/siongui/gopalilib
.. _paliDataVFS: https://github.com/siongui/paliDataVFS
