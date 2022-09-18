#!/bin/bash

dest="mocks"
rm ${dest}/* 2>/dev/null
for i in `/bin/ls -1 src/interfaces/*.go`; do
    out=`basename $i | sed 's,_interface,,g'`
    mockgen -source ${i} -destination ${dest}/mock_${out} -package ${dest}
    # Remove the I preffix from the struct name. ex.: NewMockIAuthInterface to NewMockAuthInterface
    if [ $(uname -s) = "Darwin" ]; then
        sed -i '' 's,MockI,Mock,g' ${dest}/mock_${out}
    else
        sed -i 's,MockI,Mock,g' ${dest}/mock_${out}
    fi
done
