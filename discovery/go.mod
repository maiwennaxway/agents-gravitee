module github.com/maiwennaxway/agents-gravitee/discovery

go 1.21.6

require (
	github.com/Axway/agent-sdk v1.1.89
	github.com/maiwennaxway/agents-gravitee/client v0.0.0-20240702074428-f1dd1a36750e
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/akavel/rsrc v0.10.2 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/dlclark/regexp2 v1.11.0 // indirect
	github.com/dop251/goja v0.0.0-20240627195025-eb1f15ee67d2 // indirect
	github.com/dop251/goja_nodejs v0.0.0-20240418154818-2aae10d4cbcf // indirect
	github.com/elastic/beats/v7 v7.17.22 // indirect
	github.com/elastic/elastic-agent-client/v7 v7.13.0 // indirect
	github.com/elastic/elastic-agent-libs v0.9.13 // indirect
	github.com/elastic/go-licenser v0.4.2 // indirect
	github.com/elastic/go-sysinfo v1.14.0 // indirect
	github.com/elastic/go-ucfg v0.8.8 // indirect
	github.com/elastic/go-windows v1.0.1 // indirect
	github.com/emersion/go-sasl v0.0.0-20231106173351-e73c9f7bad43 // indirect
	github.com/emersion/go-smtp v0.21.2 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.4+incompatible // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/pprof v0.0.0-20240625030939-27f56978b8b0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/ianlancetaylor/demangle v0.0.0-20240312041847-bd984b5ce465 // indirect
	github.com/jcchavezs/porto v0.6.0 // indirect
	github.com/joeshaw/multierror v0.0.0-20140124173710-69b34d4ec901 // indirect
	github.com/josephspurrier/goversioninfo v1.4.0 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/santhosh-tekuri/jsonschema v1.2.4 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	go.elastic.co/apm v1.15.0 // indirect
	go.elastic.co/apm/module/apmhttp v1.15.0 // indirect
	go.elastic.co/ecszap v1.0.2 // indirect
	go.elastic.co/fastjson v1.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
	golang.org/x/tools/go/vcs v0.1.0-deprecated // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	howett.net/plist v1.0.1 // indirect
)

replace (
	github.com/Shopify/sarama => github.com/elastic/sarama v1.19.1-0.20210823122811-11c3ef800752
	github.com/containerd/containerd => github.com/containerd/containerd v1.5.18
	github.com/docker/docker => github.com/docker/engine v0.0.0-20191113042239-ea84732a7725
	github.com/getkin/kin-openapi => github.com/getkin/kin-openapi v0.67.0
	github.com/maiwennaxway/agents-gravitee/client => ../client
)
