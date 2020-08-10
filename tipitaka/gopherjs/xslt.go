package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/libfrontend"
)

var xsltProcessor *XSLTProcessor

func GetXSLTProcessor() *XSLTProcessor {
	return xsltProcessor
}

func SetupXSLTProcessor() {
	xsltProcessor = NewXSLTProcessor()

	// Load the xsl file using synchronous (third param is set to false) XMLHttpRequest
	myXMLHTTPRequest := NewXMLHttpRequest()
	myXMLHTTPRequest.Open("GET", libfrontend.GetXslUrl(), false)
	myXMLHTTPRequest.Send()

	xslStylesheet := myXMLHTTPRequest.ResponseXML()

	// Finally import the .xsl
	xsltProcessor.ImportStylesheet(xslStylesheet)
}
