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

func markdownCodeBlocks(md []byte) []CodeBlock {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	var headings [10]string
	var currHeadingLevel = 0
	var id = 0
	var codeblocks []CodeBlock

	for _, c := range doc.GetChildren() {
		switch node := c.(type) {
		case *ast.Heading:
			hdr := node
			text := getCombinedText(c.GetChildren())
			headings[hdr.Level] = text
			currHeadingLevel = hdr.Level
			id = 0 // IDs reset within headings
		case *ast.CodeBlock:
			id++
			cblk := node
			blockHeading := strings.Join(headings[1:currHeadingLevel+1], "  ") + fmt.Sprintf("_%04d", id)
			blockHeading = strings.ReplaceAll(blockHeading, " ", "_")
			blockHeading = strings.ReplaceAll(blockHeading, "-", "_")
			blockHeading = strings.ToLower(blockHeading)

			codeblocks = append(codeblocks,
				CodeBlock{
					Code:    string(cblk.Literal),
					Heading: blockHeading,
					Info:    string(cblk.Info),
				})
		}
	}
	return codeblocks
}

func exportToBash(codeblocks []CodeBlock, out io.Writer) {
	for idx := range codeblocks {
		blk := &codeblocks[idx]
		funcName := "_f_" + blk.Heading
		varName := "_v_" + blk.Heading
		code := strings.TrimRight(blk.Code, "\n")
		fmt.Fprintf(out, "#######################\n")
		// As variable
		fmt.Fprintf(out, "read -r -d '' %v <<'EOF'\n", varName)
		fmt.Fprintf(out, "%v\n", code)
		fmt.Fprintf(out, "EOF\n\n")

		// As function
		fmt.Fprintf(out, "%v() {\n", funcName)
		codeIndented := strings.ReplaceAll(code, "\n", "\n  ")
		fmt.Fprintf(out, "  %v\n", codeIndented)
		fmt.Fprintf(out, "}\n")

		fmt.Fprintf(out, "\n")
	}
	fmt.Print("\n")
}

// getCombinedText returns a string that is a combination of Text and
// Code nodes. Useful for Markdown headers which contain command or
// other code literals like "The command `foo`"
func getCombinedText(nodes []ast.Node) string {
	var text string
	for _, c := range nodes {
		switch node := c.(type) {
		case *ast.Text:
			text += string(node.Literal)
		case *ast.Code:
			text += string(node.Literal)
		}
	}
	return text
}

func main() {
	md, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("*** Error reading stdin")
	}
	codeblocks := markdownCodeBlocks(md)
	exportToBash(codeblocks, os.Stdout)
}
