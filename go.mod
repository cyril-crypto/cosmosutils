module github.com/calvinlauyh/cosmosutils

go 1.15

require (
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/cosmos-sdk v0.45.3
	github.com/cosmos/ibc-go v1.2.2
	github.com/cosmos/ibc-go/v2 v2.0.2 // indirect
	github.com/crypto-org-chain/chain-main/v3 v3.3.3
	github.com/crypto-org-chain/cronos v0.6.5
	github.com/peggyjv/gravity-bridge/module/v2 v2.0.0-20220420162017-838c0d25e974
	github.com/stretchr/testify v1.7.1
	github.com/tharsis/ethermint v0.9.0
	github.com/thatisuday/commando v1.0.4
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/peggyjv/gravity-bridge/module/v2 => github.com/crypto-org-chain/gravity-bridge/module/v2 v2.0.0-20220620100333-f9256032790b
)
