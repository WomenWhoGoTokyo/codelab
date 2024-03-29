#!/bin/bash

declare -a codelabs;
codelabs=(
    "gocon-2018-autumn-setup"   "1nBB9Sl-mnKE5htc_Fbx_-Q3VUY1YIfF_C3Wi_HFgzoo"
    "gocon-2019-spring-setup"   "1SBE8eGYPdnZVuMffOA0t4OT_6sxndTjNqmEOv-d61mA"
    "gocon-2019-autumn-setup"   "1y9GR6RntgL7R-Ll3R_LGVy-c_lsibk5-txsbu2R5xQI"
    "google-cloud-shell-go-solo" "19rSvwntPRo0eZzZa44HBeJ2j0OX15vqh6CewrGatEF4"
    "google-app-engine-go"      "1AjOewYLEtWy-x8RGQtxQe1vMCTKia15CK3iKczmfoQM"
    "google-cloud-functions-go" "1uwGLNnjLfFky0mvQuAKu7CIVTt6jO2rIMD4AkxrRUj4"
    "tutorial-concurrency-go" "1iFi9Dq1Aw-iH-cJ8JqyoaQdrOgVDwk41Or3p1Xf93lM"
    "play-with-number" "1XGFY4kS12qH0nT3hC-dFB-Nce-HfA896GY19muS-Czk"
    "gopher-amigurumi-ja" "1hSsUCgcBIyPFY0z07T2IXuiRJkrsjMe3IQc2xaWJHYw"
    "gopher-amigurumi-en" "1CmQBaleilToO5uvoVtQFaG6t2abpsJHslep3uKBUe3Q"
)

for ((i = 0; i < ${#codelabs[@]}; i+=2)) {
    echo "generate ${codelabs[i]}..."
    echo claat export -prefix="https://storage.googleapis.com" ${codelabs[i+1]}
    claat export -prefix="https://storage.googleapis.com" ${codelabs[i+1]}
}
