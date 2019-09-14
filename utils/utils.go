package utils

import (

	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	bf"github.com/russross/blackfriday"
	"github.com/PuerkitoBio/goquery"
	."github.com/sourcegraph/syntaxhighlight"

)

func Md5(source string) string {
	checksum := md5.Sum([]byte(source))
	return fmt.Sprintf("%x", checksum)
}

func MarkdownToHtml(text string) template.HTML {
	content:=bf.MarkdownCommon([]byte(text))
	doc,_:=goquery.NewDocumentFromReader(bytes.NewReader(content))
	doc.Find("pre").Each(func(i int,selection *goquery.Selection){
		light,_:=AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
	})
	htmlString,_:=doc.Html()
	return template.HTML(htmlString)
}

