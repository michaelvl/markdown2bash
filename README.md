# Convert Markdown Code Blocks to Shell Scripts

Keeping documentation updated is tedious and error prone. Particularly
documentation that contain actual commands or e.g. configuration
data. General documentation can be high-level and thus remain correct
even when changes are being implemented, but actual commands or
configuration data cannot - it must be correct as documented.

This tool allows for **extraction of code blocks from Markdown
documentation into shell-scripts** that can be **exercised in
automated tests** and thus provide an **indication of documentation
correctness.**

# Getting Started

Imagine the following example with three code blocks. One that sets
environment variables:

```bash
export FOO=hello
export BAR=world
```

One that prints an output based on the variables:

```bash
echo "$FOO, $BAR"
```

And finally the expected output:

```
hello, world
```

We can automate testing of this documentation using the
`markdown2bash` tool. First extract code blocks:

```bash
export IMAGE_VERSION=0.0.1
cat README.md | docker run --rm -i ghcr.io/local/markdown2bash:$(IMAGE_VERSION) > readme.sh
```
