#! /bin/bash

echo "Foo"

source readme.sh

echo "Code blocks from README.md are available both as functions and variables."
echo "I.e. here we print the content of the first example code block as a"
echo "variable using: echo \"\$_v_getting_started_0001\""

echo "---"
echo "$_v_getting_started_0001"
echo "---"

echo ""
echo "Next, we execute the first code block to make the content of the FOO and"
echo "BAR variables available to the test."
echo "We do this by running: _f_getting_started_0001"

echo "---"
_f_getting_started_0001
echo "---"

echo ""
echo "The variables FOO and BAR are now available:"
echo "FOO: $FOO"
echo "BAR: $BAR"

echo ""
echo "Finally, we can execute the second code block:"

echo "---"
_f_getting_started_0002
echo "---"
