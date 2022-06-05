#!/bin/sh

appName=${1}
version=$(git describe --tags --always)
binaryName="${appName}-${version}"
remote="openresty"
scp="scp"
run="ssh ${remote} "
binDir="/home/ubuntu/apps/generator"
repoDir="${binDir}/cleanarchitecture"
repoUrl="https://github.com/gustavohenrique/cleanarchitecture.git"

echo "Copying ${binaryName}..."
$scp ${binaryName} ${remote}:${binDir}
$run "ls ${binDir}"
echo "Updating the binary file..."
$run "chmod +x ${binDir}/${binaryName}"
$run "mv ${binDir}/${appName} ${binDir}/${appName}-old" || echo ""
$run "cp ${binDir}/${binaryName} ${binDir}/${appName}"
echo "Ok. Restarting ${appName}..."
$run "systemctl --user restart ${appName}"
echo "Cloning ${repoUrl}..."
$run "rm -rf ${repoDir} 2>/dev/null && cd ${binDir} && git clone --depth 1 ${repoUrl}"
echo "Done."
