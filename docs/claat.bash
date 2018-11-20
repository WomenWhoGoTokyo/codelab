#!/bin/bash

declare -a codelabs;
codelabs=(
    "gocon-2018-autumn-setup" "1nBB9Sl-mnKE5htc_Fbx_-Q3VUY1YIfF_C3Wi_HFgzoo"
)

for ((i = 0; i < ${#codelabs[@]}; i+=2)) {
    echo "generate ${codelabs[i]}..."
    echo claat export -prefix="/codelab/" ${codelabs[i+1]}
    claat export -prefix="/codelab/" ${codelabs[i+1]}
}
