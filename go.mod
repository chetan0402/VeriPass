module github.com/chetan0402/veripass

go 1.25.5

require (
	connectrpc.com/connect v1.18.1
	entgo.io/ent v0.14.4
	github.com/coreos/go-oidc/v3 v3.17.0
	github.com/dexidp/dex/api/v2 v2.4.0
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.7.5
	golang.org/x/net v0.41.0
	golang.org/x/oauth2 v0.30.0
	google.golang.org/grpc v1.75.0
	google.golang.org/protobuf v1.36.8
)

require github.com/google/go-cmp v0.7.0 // indirect

require (
	ariga.io/atlas v0.31.1-0.20250212144724-069be8033e83 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/go-jose/go-jose/v4 v4.1.3 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	github.com/zclconf/go-cty-yaml v1.1.0 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
)

ignore (
	node_modules
	web/app/node_modules
)
