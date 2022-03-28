module github.com/WomenWhoGoTokyo/codelab/taskmanager

go 1.16

require cloud.google.com/go/datastore v1.6.0

// NOTE: Dependabotのalert対応
// https://wayne.cloud/go-mod-why/
replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
