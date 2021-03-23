package main

import (
	"path/filepath"
	"strings"

	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/tipitaka"
)

type ToCTree struct {
	lib.Tree
	ChildTrees []ToCTree
	UrlPath    string
}

func ActionToUrlPath(action string) string {
	noext := strings.TrimSuffix(action, filepath.Ext(action))
	return "/" + strings.Replace(noext, ".", "/", -1)
}

func CopyTreeToToCTree(t lib.Tree, toc *ToCTree) {
	toc.Text = t.Text
	toc.Src = t.Src
	toc.Action = t.Action

	/*
		if t.Text != "" && t.Src == "" && t.Action == "" {
			println(t.Text)
		}
	*/

	if subpath := tipitaka.TrimTreeText2(t.Text); subpath != "" {
		if subpath == "tipiṭaka (mūla)" {
			subpath = "canon"
		}
		toc.UrlPath = toc.UrlPath + "/" + subpath
	}

	for _, subtree := range t.SubTrees {
		st := ToCTree{UrlPath: toc.UrlPath}
		CopyTreeToToCTree(subtree, &st)
		toc.ChildTrees = append(toc.ChildTrees, st)
	}

	toc.UrlPath += "/"
	//println(toc.UrlPath)
}

func SetupTipitakaUrl(tree lib.Tree) {
	toctree := ToCTree{}
	CopyTreeToToCTree(tree, &toctree)
}
