module github.com/thalesfu/CK2Commands

go 1.19

require (
	github.com/fsnotify/fsnotify v1.7.0
	github.com/thalesfu/ck2nebula v0.0.0-20240124143540-d6fb273d2f4f
	github.com/thalesfu/nebulagolang v0.0.0-20240125053636-96c493d57eef
	github.com/thalesfu/paradoxtools v0.0.0-20230223131212-ea15646baa51
)

replace (
	github.com/thalesfu/ck2nebula => ../ck2nebula
	github.com/thalesfu/nebulagolang => ../nebulagolang
	github.com/thalesfu/paradoxtools => ../paradoxtools
)

require (
	github.com/samber/lo v1.39.0 // indirect
	github.com/vesoft-inc/fbthrift v0.0.0-20230214024353-fa2f34755b28 // indirect
	github.com/vesoft-inc/nebula-go/v3 v3.7.0 // indirect
	golang.org/x/exp v0.0.0-20240409090435-93d18d7e34b8 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
