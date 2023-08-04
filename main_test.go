package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	markdown := []byte("# Header1\n" +
		"## Header2\n" +
		"```bash\n" +
		"echo \"foo\"\n" +
		"```\n" +
		"```bash\n" +
		"echo \"bar\"\n" +
		"```")
	codeblocks := markdownCodeBlocks(markdown)
	if len(codeblocks) != 2 {
		t.Fatalf("Number of codeblocks mismatch, got %v", len(codeblocks))
	}
	if codeblocks[0].Heading != "header1__header2_0001" ||
		codeblocks[1].Heading != "header1__header2_0002" {
		t.Fatalf("Heading format error")
	}
}
