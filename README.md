# Go bindings for Ethereum smart contracts

Generated using [abigen](https://geth.ethereum.org/docs/tools/abigen)

## How To Use

```go
package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/projectawakening/contracts-go/bindings/world"
	"log"
)

func main() {
	w, err := world.NewWorld(common.HexToAddress("0x0000"), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	w.EveworldBringOnline(...)
}
```

## How To Release

1. Wait for a new [world-chain-contracts](https://github.com/projectawakening/world-chain-contracts/releases) deployment
2. Make sure the new ABIs are available on abi-export.projectawakening.io
3. Create a new [GitHub Release](https://github.com/projectawakening/contracts-go/releases/new) here
4. Use the same Git tag as [world-chain-contracts](https://github.com/projectawakening/world-chain-contracts/releases) (e.g `v0.0.8`)
5. Publish Release
6. CI will generate and push the new Go bindings to `main`


## Generate Bindings Locally

1. clone this repo and `cd` into it
2. `go install github.com/ethereum/go-ethereum/cmd/abigen@latest`
3. `VERSION=v0.0.8 go generate` (replace the version if needed)