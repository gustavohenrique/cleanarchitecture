#!/bin/bash

dest="mocks"
rm ${dest}/* 2>/dev/null
for i in `/bin/ls -1 src/interfaces/*.go`; do
    out=`basename $i | sed 's,_interface,,g'`
    mockgen -source ${i} -destination ${dest}/mock_${out} -package ${dest}
    # Rename NewMockIAgentInterface to NewMockAgentInterface
    sed -i 's,MockI,Mock,g' ${dest}/mock_${out}
done
