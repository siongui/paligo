package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
)

func traverseTreeviewData(tree lib.Tree, actionFunc func(string)) *Object {
	if len(tree.SubTrees) > 0 {
		div := Document.CreateElement("div")
		div.ClassList().Add("item")

		sign := Document.CreateElement("span")
		sign.SetInnerHTML("+")

		span := Document.CreateElement("span")
		span.ClassList().Add("treeNode")
		span.SetInnerHTML(tree.Text)

		div.AppendChild(sign)
		div.AppendChild(span)

		childrenContainer := Document.CreateElement("div")
		childrenContainer.ClassList().Add("childrenContainer")
		for _, child := range tree.SubTrees {
			childrenContainer.AppendChild(traverseTreeviewData(child, actionFunc))
		}
		childrenContainer.Style().SetDisplay("none")

		span.AddEventListener("click", func(e Event) {
			if childrenContainer.Style().Display() == "none" {
				childrenContainer.Style().SetDisplay("")
				sign.SetInnerHTML("-")
			} else {
				childrenContainer.Style().SetDisplay("none")
				sign.SetInnerHTML("+")
			}
		})

		all := Document.CreateElement("div")
		all.AppendChild(div)
		all.AppendChild(childrenContainer)

		return all
	} else {
		div := Document.CreateElement("div")
		div.ClassList().Add("item")

		span := Document.CreateElement("span")
		span.ClassList().Add("treeNode")
		span.SetInnerHTML(tree.Text)
		span.AddEventListener("click", func(e Event) {
			actionFunc(tree.Action)
		})

		div.AppendChild(span)
		return div
	}
}

func appendCSSToHeadElement() {
	css := `.item {
	  margin-bottom: 3px;
	  padding-bottom: 3px;
	  border-bottom: 1px solid #E0E0E0;
	}

	.item:hover {
	  background-color: #F0F8FF;
	}

	.treeNode:hover {
	  cursor: pointer;
	  color: blue;
	}

	.childrenContainer {
	  margin-left: .4em;
	  padding-left: .4em;
	  border-left: 1px dotted blue;
	}`
	s := Document.CreateElement("style")
	s.SetInnerHTML(css)
	// insert style of treeview at the end of head element
	Document.QuerySelector("head").AppendChild(s)
}

func NewTreeview(id string, root lib.Tree, actionFunc func(string)) {
	appendCSSToHeadElement()
	treeviewContainer := Document.GetElementById(id)

	for _, child := range root.SubTrees {
		tree := traverseTreeviewData(child, actionFunc)
		treeviewContainer.AppendChild(tree)
	}
}
