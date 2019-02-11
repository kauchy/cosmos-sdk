module github.com/cosmos/cosmos-sdk

require (
	github.com/ZondaX/hid-go v0.4.0
	github.com/ZondaX/ledger-go v0.4.0
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.0.0-20181130015935-7d2daa5bfef2
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8
	github.com/davecgh/go-spew v1.1.1
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-kit/kit v0.6.0
	github.com/go-logfmt/logfmt v0.4.0
	github.com/go-stack/stack v1.8.0
	github.com/gogo/protobuf v1.1.1
	github.com/golang/protobuf v1.2.0
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/gorilla/context v1.1.1
	github.com/gorilla/mux v1.6.2
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/hcl v1.0.0
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/jmhodges/levigo v0.0.0-20161115193449-c42d9e0ca023
	github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
	github.com/magiconair/properties v1.8.0
	github.com/mattn/go-isatty v0.0.4
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/mitchellh/mapstructure v1.1.2
	github.com/otiai10/copy v0.0.0-20180813032824-7e9a647135a1
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v0.9.2
	github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910
	github.com/prometheus/common v0.0.0-20181126121408-4724e9255275
	github.com/prometheus/procfs v0.0.0-20181204211112-1dc9a6cbc91a
	github.com/rakyll/statik v0.1.4
	github.com/rcrowley/go-metrics v0.0.0-20180503174638-e2704e165165
	github.com/rs/cors v1.6.0
	github.com/spf13/afero v1.1.2
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/jwalterweatherman v1.0.0
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.0.3
	github.com/stretchr/testify v1.2.2
	github.com/syndtr/goleveldb v0.0.0-20180708030551-c4c61651e9e3
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/iavl v0.12.0
	github.com/tendermint/tendermint v0.30.0-rc0
	github.com/zondax/ledger-cosmos-go v0.9.2
	golang.org/x/crypto v0.0.0-20181106171534-e4dc69e5b2fd
	golang.org/x/sys v0.0.0-20190114130336-2be517255631 // indirect
	google.golang.org/grpc v1.17.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace golang.org/x/crypto v0.0.0-20181106171534-e4dc69e5b2fd => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
