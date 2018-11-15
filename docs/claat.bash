#!/bin/bash

declare -a codelabs;
codelabs=(
    #"go-greeting"          "1TwbojRvp4r28KceQQHTlSdN4lsRzL9bqB0diS2Xm39I"
)

for ((i = 0; i < ${#codelabs[@]}; i+=2)) {
    echo "generate ${codelabs[i]}..."
    echo claat export -prefix="/codelab/" ${codelabs[i+1]}
    claat export -prefix="/codelab/" ${codelabs[i+1]}
}
