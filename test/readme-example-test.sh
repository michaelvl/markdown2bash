#! /bin/bash

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

echo ""
echo "since README.md also contain the expected output of running the second code"
echo "block, we can compare the result with the expected output in the third code block"

echo "---"
_f_getting_started_0002 > out.txt

outdata=`cat out.txt`
if [ "$_v_getting_started_0003" = "$outdata" ]; then
    echo "Variable \$_v_getting_started_0003 matches output from \$_f_getting_started_0002"
else
    echo "Variable \$_v_getting_started_0003 DOES NOT match output from \$_f_getting_started_0002 ($_v_getting_started_0003 vs $outdata)"
fi
