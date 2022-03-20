module github.com/WomenWhoGoTokyo/codelab/guestbook

go 1.12

require cloud.google.com/go/datastore v1.5.0

// NOTE: Dependabotのalert対応
// https://wayne.cloud/go-mod-why/
replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
