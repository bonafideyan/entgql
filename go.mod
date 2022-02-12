module github.com/sijad/entgql

go 1.17

replace entgo.io/ent => ../ent

replace entgo.io/contrib => ../ent-contrib

// replace github.com/sijad/entgql = > ../entgql

replace github.com/99designs/gqlgen => ../gqlgen

require (
	entgo.io/contrib v0.0.0-00010101000000-000000000000
	entgo.io/ent v0.9.1
	github.com/99designs/gqlgen v0.13.0
	github.com/go-openapi/inflect v0.19.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
)

require (
	github.com/agnivade/levenshtein v1.1.0 // indirect
	github.com/go-bindata/go-bindata v1.0.1-0.20190711162640-ee3c2418e368 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vektah/gqlparser/v2 v2.1.0 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20211020064051-0ec99a608a1b // indirect
	golang.org/x/tools v0.1.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
