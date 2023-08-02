package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"

)

type CodeBlock struct {
	Info    string
	Heading string
	Code    string
}

func markdownToBash(out io.Writer, md []byte) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	var headings [10]string
	var currHeadingLevel = 0
	var id = 0

	for _, c := range doc.GetChildren() {
		id += 1
		switch c.(type) {
		case *ast.Heading:
			hdr := c.(*ast.Heading)
			text := getCombinedText(c.GetChildren())
			headings[hdr.Level] = text
			currHeadingLevel = hdr.Level
			id = 0 // IDs reset within headings
		case *ast.CodeBlock:
			cblk := c.(*ast.CodeBlock)
			//fmt.Printf("CodeBlock1  %+v\n", cblk)
			//fmt.Printf("CodeBlock2  %v\n", string(cblk.Info))
			//fmt.Printf("CodeBlock3  \n%v\n", string(cblk.Literal))
			blockHeading := strings.Join(headings[1:currHeadingLevel], "  ") + fmt.Sprintf("_%04d", id)
			blockHeading = strings.ReplaceAll(blockHeading, " ", "_")
			blockHeading = strings.ReplaceAll(blockHeading, "-", "_")
			blockHeading = strings.ToLower(blockHeading)
			funcName := "_f_" + blockHeading
			varName := "_v_" + blockHeading
			fmt.Fprintf(out, "#######################\n")
			// As variable
			fmt.Fprintf(out, "read -r -d '' %v <<'EOF'\n", varName)
			fmt.Fprintf(out, "%v\n", string(cblk.Literal))
			fmt.Fprintf(out, "EOF\n")

			// As function
			fmt.Fprintf(out, "%v() {\n", funcName)
			fmt.Fprintf(out, "  echo \"----------------------\"\n")
			fmt.Fprintf(out, "  echo \"### Executing: %v\"\n", funcName)
			fmt.Fprintf(out, "  echo \"$%v\"\n", varName)
			fmt.Fprintf(out, "  %v\n", strings.ReplaceAll(string(cblk.Literal), "\n", "\n  "))
			fmt.Fprintf(out, "}\n")

			fmt.Fprintf(out, "\n")
		}
	}
	fmt.Print("\n")
}

// getCombinedText returns a string that is a combination of Text and
// Code nodes. Useful for Markdown headers which contain command or
// other code literals like "The command `foo`"
func getCombinedText(nodes []ast.Node) string {
	var text string
	for _, c := range nodes {
		switch c.(type) {
		case *ast.Text:
			txt := c.(*ast.Text)
			text += string(txt.Literal)
		case *ast.Code:
			cblk := c.(*ast.Code)
			text += string(cblk.Literal)
		}
	}
	return text
}

func main() {
	md, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("*** Error reading stdin")
	}
	markdownToBash(os.Stdout, md)
}
