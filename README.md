# Convert Markdown Code Blocks to Shell Scripts

Keeping documentation updated is tedious and error prone. Particularly
documentation that contain actual commands or e.g. configuration
data. General documentation can be high-level and thus remain correct
even when changes are being implemented, but actual commands or
configuration data cannot - it must be correct as documented.

This tool allows for **conversion of code blocks from Markdown
documentation into shell-scripts** that can be **exercised in automated
tests** and thus provide an **indication of incorrect documentation.**

# Getting Started

```bash
export FOO=hello
export BAR=world
```


```bash
echo "$FOO, $BAR"
```
