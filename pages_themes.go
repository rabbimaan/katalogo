package main

import (
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
)

func addThemePages(pdf *gofpdf.Fpdf, page Page) {
	filelist, err := ioutil.ReadDir(page.Folder)
	if err != nil {
		die("failed to read theme folder.", err)
	}
	for i, fileinfo := range filelist {
		_ = i
		_ = fileinfo
	}
}
