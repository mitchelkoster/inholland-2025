/*
Use the package manager to "go get github.com/fatih/color"
so that we can print using colors.

Note: You have to run "go mod init <projectname>"" first
*/
package main

import (
	"github.com/fatih/color"
)

func main() {
	color.Red("Hello world!")
}
