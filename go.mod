module github.com/sijad/entgql

go 1.18

replace entgo.io/ent => ../ent

replace entgo.io/contrib => ../ent-contrib

// replace github.com/sijad/entgql = > ../entgql

replace github.com/99designs/gqlgen => ../gqlgen

require (
	entgo.io/contrib v0.0.0-00010101000000-000000000000
	entgo.io/ent v0.11.0-0.20220502113020-4ac82f5bb3f0
	github.com/99designs/gqlgen v0.17.9
	github.com/go-openapi/inflect v0.19.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.5.0
)

require (
	ariga.io/atlas v0.6.2-0.20220819114704-2060066abac7 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/go-bindata/go-bindata v1.0.0 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vektah/gqlparser/v2 v2.4.4 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.13-0.20220804200503-81c7dc4e4efa // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
