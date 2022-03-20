module github.com/WomenWhoGoTokyo/codelab/image-slender

go 1.14

require (
	github.com/jinzhu/configor v1.1.1
	github.com/pkg/errors v0.9.1
	golang.org/x/image v0.0.0-20200119044424-58c23975cae1
)

// NOTE: Dependabotのalert対応
// https://wayne.cloud/go-mod-why/
replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
