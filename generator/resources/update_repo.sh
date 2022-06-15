#!/bin/sh

binDir="/home/ubuntu/apps/generator"
remote="openresty"
run="ssh ${remote} "
repoDir="${binDir}/cleanarchitecture"
repoUrl="https://github.com/gustavohenrique/cleanarchitecture.git"

echo "Cloning ${repoUrl}..."
$run "rm -rf ${repoDir} 2>/dev/null && cd ${binDir} && git clone --depth 1 ${repoUrl}"
echo "Done."
