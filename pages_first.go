package main

import "github.com/jung-kurt/gofpdf"

func addFirstPage(pdf *gofpdf.Fpdf) {
	pdf.AddPage()

}
