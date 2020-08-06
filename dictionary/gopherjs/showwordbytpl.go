package main

import (
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dicmgr"
)

func showWordByTemplate(wi lib.BookIdWordExps) {
	mainContent.RemoveAllChildNodes()

	mainContent.Set("innerHTML", dicmgr.GetWordDefinitionHtml(wi, getSetting(), navigatorLanguages))
}
