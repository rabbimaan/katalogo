package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"os"
)

var config Config

type Config struct {
	Filename       string
	DoubleSided    bool
	Title          string
	FrontpageImage string
	Pages          []Page
}
type Page struct {
	Folder string
	Title  string
}

func main() {
	var err error

	// read config
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		die("failed to read config file.", err)
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		die("failed to parse config file.", err)
	}

	// init values
	// nothing yet.

	// create pdf
	pdf := gofpdf.New("P", "mm", "A5", "fonts")
	addFirstPage(pdf)
	for _, p := range config.Pages {
		addThemePages(pdf, p)
	}

	// write output
	err = pdf.OutputFileAndClose(config.Filename)
	if err != nil {
		die("failed to write output file.", err)
	}
	// ...
	// profit!
	fmt.Printf("file succesfully written.\n")
	fmt.Printf("press enter to continue.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func die(custommessage string, err error) {
	fmt.Printf("a fatal error occured: %s\n", custommessage)
	fmt.Printf("technical error message: %s\n", err)
	fmt.Printf("press enter to continue.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(1)
}
