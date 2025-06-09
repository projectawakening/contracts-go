package objectid

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Generator interface {
	Generate(id int64) (*big.Int, error)
	GetTenant() [32]byte
}

type objectIdGenerator struct {
	arguments abi.Arguments
	tenant    [32]byte
}

func NewGenerator(tenant string) (Generator, error) {
	bytes32, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bytes32 type: %w", err)
	}

	uint256, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create uint256 type: %w", err)
	}

	return &objectIdGenerator{
		arguments: abi.Arguments{{Type: bytes32}, {Type: uint256}},
		tenant:    [32]byte(crypto.Keccak256([]byte(tenant))),
	}, nil
}

func (o *objectIdGenerator) GetTenant() [32]byte {
	return o.tenant
}

func (o *objectIdGenerator) Generate(id int64) (*big.Int, error) {
	packedArguments, err := o.arguments.Pack(o.tenant, big.NewInt(id))
	if err != nil {
		return nil, fmt.Errorf("failed to pack arguments: %w", err)
	}

	return new(big.Int).SetBytes(crypto.Keccak256(packedArguments)), nil
}
