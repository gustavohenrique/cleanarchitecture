#!/bin/bash

dest="mocks"
sources=( "src/domain/ports/*.go" "src/infrastructure/datastores/db/*.go" )
rm ${dest}/* 2>/dev/null
for pattern in "${sources[@]}"; do
    for i in `/bin/ls -1 ${pattern}`; do
        out=`basename $i | sed 's,_interface,,g'`
        mockgen -source ${i} -destination ${dest}/mock_${out} -package ${dest}
    done
done
