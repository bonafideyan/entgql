module github.com/sijad/entgql

go 1.22.5

replace entgo.io/ent => ../ent

replace entgo.io/contrib => ../ent-contrib

// replace github.com/sijad/entgql = > ../entgql

replace github.com/99designs/gqlgen => ../gqlgen

require (
	entgo.io/contrib v0.0.0-00010101000000-000000000000
	entgo.io/ent v0.13.2-0.20240717044502-34158f2c129b
	github.com/99designs/gqlgen v0.17.48
	github.com/go-openapi/inflect v0.19.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.6.1
)

require (
	ariga.io/atlas v0.25.1-0.20240717145915-af51d3945208 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/go-bindata/go-bindata v1.0.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vektah/gqlparser/v2 v2.5.12 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	golang.org/x/exp v0.0.0-20221230185412-738e83a70c30 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/tools v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
