package main

import (
	"strings"

	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/tipitaka"
)

func traverse(tree lib.Tree, indent int) {
	println(strings.Repeat(" ", indent) + tipitaka.TrimTreeText2(tree.Text))
	for _, subtree := range tree.SubTrees {
		traverse(subtree, indent+2)
	}
}

func SetupTipitakaUrl(tree lib.Tree) {
	traverse(tree, 0)
}
