<!--
parent:
  order: false
-->

<div align="center">
  <h1> Cascadia </h1>
</div>

Cascadia is a scalable, high-throughput Proof-of-Stake blockchain that is fully compatible and
interoperable with Ethereum. It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.

**Note**: Requires [Go 1.18+](https://golang.org/dl/)

## Installation

For prerequisites and detailed build instructions please read the [Installation]
instructions. Once the dependencies are installed, run:

```bash
make install
```

## Build binary

Before running `make install` in the repo, you need open a file `go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.45.9/types/coin.go` and modify the line 764 as follows. reDnmString = `[a-zA-Z][a-zA-Z0-9/-]{2,127}` -> reDnmString = `[a-zA-Z][a-zA-Z0-9/-]{1,127}`






