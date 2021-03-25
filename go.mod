module github.com/sijad/entgql

go 1.15

replace entgo.io/ent => ../ent

replace entgo.io/contrib => ../ent-contrib

// replace github.com/sijad/entgql = > ../entgql

replace github.com/99designs/gqlgen => ../gqlgen

require (
	entgo.io/contrib v0.0.0-00010101000000-000000000000
	entgo.io/ent v0.7.0
	github.com/99designs/gqlgen v0.13.0
	github.com/go-openapi/inflect v0.19.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/joncalhoun/pipe v0.0.0-20170510025636-72505674a733 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
)
