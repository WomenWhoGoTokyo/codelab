module functions/taskmanagement

go 1.16

require (
	cloud.google.com/go/datastore v1.1.0
	github.com/pkg/errors v0.9.1
)

// NOTE: Dependabotのalert対応
// https://wayne.cloud/go-mod-why/
replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
