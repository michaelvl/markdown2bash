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
export IMAGE_VERSION=0.0.2
cat README.md | docker run --rm -i ghcr.io/michaelvl/markdown2bash:$(IMAGE_VERSION) > readme.sh
```

This will generate a shell script that e.g. includes the first code
block both as a function and a variable:

```bash
read -r -d '' _v_getting_started_0001 <<'EOF_5577006791947779410'
export FOO=hello
export BAR=world
EOF_5577006791947779410

_f_getting_started_0001() {
  export FOO=hello
  export BAR=world
}
```

Note how the function/variable names are constructed from the
lower-case version of the Headers (`getting_started`) in the markdown,
followed by an incremented id (`0001`) to handle multiple code blocks
per heading and function names start with a `_f_` prefix and variables
with `_v_`.

From these function and variable definitions, we can construct tests
that use the actual code from markdown documentation. See
[test/readme-example-test.sh](test/readme-example-test.sh) for an
example that test the code blocks above. See also the GitHub action
[e2e-tests.yaml](.github/workflows/e2e-tests.yaml) for how
`markdown2bash` is used to test the code blocks above.
