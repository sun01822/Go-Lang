package main

import (
	"fmt"
	"strings"
)

// Formatter interface
type Formatter interface {
	Format([]string) string
}

// CSVFormatter struct
type CSVFormatter struct{}


// Format method for CSVFormatter
func (c CSVFormatter) Format(data []string) string {
	return strings.Join(data, ",")
}

// HTMLFormatter struct
type HTMLFormatter struct{}

// Format method for HTMLFormatter
func (h HTMLFormatter) Format(data []string) string {
	return "<ul>\n   <li>" + strings.Join(data, "</li>\n   <li>") + "</li>\n</ul>"
}

// TextFormat struct
func TestFormat(f Formatter, data []string) {
	result := f.Format(data)
	fmt.Println(result)
}

func main(){
	//Test CSV Formatter
	csvFormatter := CSVFormatter{}  // instantiate CSVFormatter
	TestFormat(csvFormatter, []string{"one", "two", "three"})  // call TestFormat with CSVFormatter
	fmt.Println()

	//Test HTML Formatter
	htmlFormatter := HTMLFormatter{}  // instantiate HTMLFormatter
	TestFormat(htmlFormatter, []string{"one", "two", "three"})  // call TestFormat with HTMLFormatter
}