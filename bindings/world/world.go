// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package world

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AggressionParams is an auto generated low-level Go binding around an user-defined struct.
type AggressionParams struct {
	SmartObjectId *big.Int
	PriorityQueue []TargetPriority
	Turret        Turret
	Aggressor     SmartTurretTarget
	Victim        SmartTurretTarget
}

// CreateAndAnchorParams is an auto generated low-level Go binding around an user-defined struct.
type CreateAndAnchorParams struct {
	SmartObjectId      *big.Int
	AssemblyType       string
	EntityRecordParams EntityRecordParams
	Owner              common.Address
	LocationData       LocationData
}

// CreateInventoryItemParams is an auto generated low-level Go binding around an user-defined struct.
type CreateInventoryItemParams struct {
	SmartObjectId *big.Int
	TenantId      [32]byte
	ItemId        *big.Int
	TypeId        *big.Int
	Volume        *big.Int
	Quantity      *big.Int
}

// EntityMetadataParams is an auto generated low-level Go binding around an user-defined struct.
type EntityMetadataParams struct {
	Name        string
	DappURL     string
	Description string
}

// EntityRecordParams is an auto generated low-level Go binding around an user-defined struct.
type EntityRecordParams struct {
	TenantId [32]byte
	TypeId   *big.Int
	ItemId   *big.Int
	Volume   *big.Int
}

// FuelParams is an auto generated low-level Go binding around an user-defined struct.
type FuelParams struct {
	FuelMaxCapacity       *big.Int
	FuelBurnRateInSeconds *big.Int
}

// InventoryItemParams is an auto generated low-level Go binding around an user-defined struct.
type InventoryItemParams struct {
	SmartObjectId *big.Int
	Quantity      *big.Int
}

// KillMailData is an auto generated low-level Go binding around an user-defined struct.
type KillMailData struct {
	KillerCharacterId *big.Int
	VictimCharacterId *big.Int
	LossType          uint8
	SolarSystemId     *big.Int
	KillTimestamp     *big.Int
}

// LocationData is an auto generated low-level Go binding around an user-defined struct.
type LocationData struct {
	SolarSystemId *big.Int
	X             *big.Int
	Y             *big.Int
	Z             *big.Int
}

// SmartTurretTarget is an auto generated low-level Go binding around an user-defined struct.
type SmartTurretTarget struct {
	ShipId      *big.Int
	ShipTypeId  *big.Int
	CharacterId *big.Int
	HpRatio     *big.Int
	ShieldRatio *big.Int
	ArmorRatio  *big.Int
}

// SystemCallData is an auto generated low-level Go binding around an user-defined struct.
type SystemCallData struct {
	SystemId [32]byte
	CallData []byte
}

// SystemCallFromData is an auto generated low-level Go binding around an user-defined struct.
type SystemCallFromData struct {
	From     common.Address
	SystemId [32]byte
	CallData []byte
}

// TargetPriority is an auto generated low-level Go binding around an user-defined struct.
type TargetPriority struct {
	Target SmartTurretTarget
	Weight *big.Int
}

// Turret is an auto generated low-level Go binding around an user-defined struct.
type Turret struct {
	WeaponTypeId *big.Int
	AmmoTypeId   *big.Int
	ChargesLeft  *big.Int
}

// WorldMetaData contains all meta data concerning the World contract.
var WorldMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"batchCall\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallData[]\",\"components\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchCallFrom\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallFromData[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"callFrom\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"creator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier___handleNodeOffline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__adminSupportOrDirectOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__adminSupportOrDirectOwnerGates\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__aggression\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structAggressionParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"aggressor\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"victim\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__anchor\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__areGatesOnline\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__assignItemToInventory\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__assignOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__bringOffline\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__bringOnline\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__canCrossTransferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__canJump\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__canTransferFromEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__canTransferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__canTransferToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__configureDeployableAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureEntityRecordAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureEphemeralInteractAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureEphemeralInventoryAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureFuelAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureFuelEfficiency\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelEntityParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fuelEfficiency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureFuelParameters\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelParams\",\"type\":\"tuple\",\"internalType\":\"structFuelParams\",\"components\":[{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelBurnRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureGate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureInventoryAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureInventoryInteractAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureKillMailAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureLocationAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureNetworkNodeAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureOwnershipAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureSmartAssemblyAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureSmartCharacterAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureSmartGateAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureSmartStorageUnitAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureSmartTurretAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__configureTurret\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__connectAssembly\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createAndAnchor\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createAndAnchorGate\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"maxDistance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createAndAnchorNetworkNode\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelParams\",\"type\":\"tuple\",\"internalType\":\"structFuelParams\",\"components\":[{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelBurnRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"maxEnergyCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentProduction\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createAndAnchorStorageUnit\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createAndAnchorTurret\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createAndDepositEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structCreateInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createAndDepositInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structCreateInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createAssembly\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createCharacter\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tribeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"entityRecordMetadata\",\"type\":\"tuple\",\"internalType\":\"structEntityMetadataParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createDeployable\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createMetadata\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordMetadata\",\"type\":\"tuple\",\"internalType\":\"structEntityMetadataParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__createRecord\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__crossTransferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromEphemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toEphemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__depositEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__depositFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__depositInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__destroyDeployable\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__disconnectAssembly\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__disconnectNetworkNode\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__getCurrentFuelConsumptionStatus\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"elapsedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitsToConsume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualConsumptionRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__getDeployableClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__getEphemeralOwner\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__getEphemeralSmartObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"evefrontier__getInventoryOwner\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__getNetworkNodeClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__getSmartCharacterClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__inProximity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turretTarget\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__isAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__isAnyGateLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__isClassScoped\",\"inputs\":[{\"name\":\"classId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__isEphemeralOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__isGateLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__isOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__isOwnerOfBothGates\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__isWithinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__linkGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyAdminOrCallAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyAdminOrClassScopedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyAdminOrOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyAdminOrOwnerSupported\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyAdminOrScopeEnforcedCall\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyAdminSupportedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyAdminSupportedOwnerOrCall\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyCallAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyCallAccessOrDirectEphemeralOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyCallAccessWithScopeEnforced\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyClassScopedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyClassScopedOrCharAdminOrOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyDirectAdmin\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyDirectAdminOrCallAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyDirectEphemeralOwnerOrCall\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyDirectOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyEphemeralOwnerOrTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyOwnerOrEphemeralCrossTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyOwnerOrEphemeralTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyOwnerOrInventoryTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlyOwnerWithAdminSupportAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__onlySmartAssemblyClassScopedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__owner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evefrontier__registerNetworkNodeClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__registerSmartAssemblies\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__registerSmartCharacterClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__releaseAssemblyEnergy\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__releaseNetworkNodeEnergy\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__removeCharacter\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__removeItemFromInventory\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__removeOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__reportKill\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killMailData\",\"type\":\"tuple\",\"internalType\":\"structKillMailData\",\"components\":[{\"name\":\"killerCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"victimCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lossType\",\"type\":\"uint8\",\"internalType\":\"enumKillMailLossType\"},{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__reserveAssemblyEnergy\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__reserveNetworkNodeEnergy\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__saveLocation\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setAssemblyType\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setCrossTransferToEphemeralAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setDappURL\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setDescription\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setEphemeralCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setFuelMaxCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setName\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setTransferFromEphemeralAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setTransferToEphemeralAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__setTransferToInventoryAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__startBurn\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__stopBurn\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__transferFromEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__transferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__transferToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__unanchor\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__unlinkGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__updateAssemblyType\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__updateEnergyHistory\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__updateFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__updateTribeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tribeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__withdrawEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__withdrawFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evefrontier__withdrawInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldSlice\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLayout\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeySchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValueSchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initModule\",\"type\":\"address\",\"internalType\":\"contractIModule\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installRootModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"popFromDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"byteLengthToPop\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pushToDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dataToPush\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"systemFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespace\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerRootFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"worldFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"systemFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystem\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"contractSystem\"},{\"name\":\"publicAccess\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerTable\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"},{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"keyNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"fieldNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storeVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferBalanceToAddress\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferBalanceToNamespace\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"worldVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"HelloStore\",\"inputs\":[{\"name\":\"storeVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HelloWorld\",\"inputs\":[{\"name\":\"worldVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_DeleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SetRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"EncodedLengths\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Access_CannotTransferFromEphemeral\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrClassScoped\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrOwnerSupported\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupported\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupportedOrDirectOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupportedOrDirectOwnerGates\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupportedOwnerOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotClassScoped\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotClassScopedAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectAdminOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectEphemeralOwnerOrCanCrossTransferToEphemeral\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectOwnerOrCanTransferToEphemeral\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectOwnerOrCanTransferToInventory\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotEphemeralOwnerOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotEphemeralOwnerOrCallAccessWithEphemeralOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotOwnerWithAdminSupportAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Deployable_IncorrectState\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentState\",\"type\":\"uint8\",\"internalType\":\"enumState\"}]},{\"type\":\"error\",\"name\":\"Deployable_InvalidObjectOwner\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"smartObjectOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EncodedLengths_InvalidLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InsufficientCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"usedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidEphemeralOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidItemDepositQuantity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidItemObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidSmartObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidTenantId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_NonExistentEntityRecord\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FieldLayout_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"staticDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"computedStaticDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_StaticLengthDoesNotFitInAWord\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_StaticLengthIsNotZero\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_StaticLengthIsZero\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_TooManyDynamicFields\",\"inputs\":[{\"name\":\"numFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_TooManyFields\",\"inputs\":[{\"name\":\"numFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_BurnAlreadyStopped\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_BurnNotActive\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_ExceedsMaxCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalProjectedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InsufficientFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableFuel\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelAmount\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelBurnRate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelBurnRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelEfficiency\",\"inputs\":[{\"name\":\"fuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelEfficiency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelMaxCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelSmartObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelTypeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelUnitVolume\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_TypeMismatch\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentFuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newFuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_Ephemeral_InsufficientQuantity\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InsufficientQuantity\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InvalidInventory\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InvalidOperation\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InvalidQuantity\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_NonexistentItemRecord\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_NonexistentObject\",\"inputs\":[{\"name\":\"objectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_SingletonAlreadyAssigned\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentInventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_SingletonDirectlyOwned\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"directOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_ZeroQuantity\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InsufficientCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"usedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidItemDepositQuantity\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidItemObjectId\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidTenantId\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"Inventory_NonExistentEntityRecord\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"KillMail_AlreadyExists\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"KillMail_InvalidCharacterId\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Module_AlreadyInstalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Module_MissingDependency\",\"inputs\":[{\"name\":\"dependency\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Module_NonRootInstallNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Module_RootInstallNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NetworkNode_AlreadyExists\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_AssemblyAlreadyConnected\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_AssemblyNotConnected\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_DoesNotExist\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_InsufficientEnergy\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_NotOnline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Ownership_AlreadyOwned\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Ownership_InvalidAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Ownership_InvalidOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"invalidOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Ownership_InvalidSingleton\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Ownership_NonexistentObject\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Ownership_SingletonInInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Schema_InvalidLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Schema_StaticTypeAfterDynamicType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Slice_OutOfBounds\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_DoesNotExist\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_InvalidObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_InvalidTenantId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_InvalidTypeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_TypeCannotBeEmpty\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacterDoesNotExist\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_AlreadyCreated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_InvalidObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_InvalidTenantId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_InvalidTypeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateAlreadyLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateNotLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateNotOnline\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GatesNotOnline\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotWithtinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_SameSourceAndDestination\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_UndefinedClassId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Store_IndexOutOfBounds\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessedIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidBounds\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidFieldNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidKeyNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidSplice\",\"inputs\":[{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"fieldLength\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaDynamicLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaStaticLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_TableAlreadyExists\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_TableNotFound\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_AccessDenied\",\"inputs\":[{\"name\":\"resource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"World_CallbackNotAllowed\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_DelegationNotFound\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorAlreadyExists\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorNotFound\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"World_InterfaceNotSupported\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InvalidNamespace\",\"inputs\":[{\"name\":\"namespace\",\"type\":\"bytes14\",\"internalType\":\"bytes14\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceId\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceAlreadyExists\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceNotFound\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_SystemAlreadyExists\",\"inputs\":[{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_UnlimitedDelegationNotAllowed\",\"inputs\":[]}]",
}

// WorldABI is the input ABI used to generate the binding from.
// Deprecated: Use WorldMetaData.ABI instead.
var WorldABI = WorldMetaData.ABI

// World is an auto generated Go binding around an Ethereum contract.
type World struct {
	WorldCaller     // Read-only binding to the contract
	WorldTransactor // Write-only binding to the contract
	WorldFilterer   // Log filterer for contract events
}

// WorldCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorldSession struct {
	Contract     *World            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorldCallerSession struct {
	Contract *WorldCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WorldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorldTransactorSession struct {
	Contract     *WorldTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorldRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorldRaw struct {
	Contract *World // Generic contract binding to access the raw methods on
}

// WorldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorldCallerRaw struct {
	Contract *WorldCaller // Generic read-only contract binding to access the raw methods on
}

// WorldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorldTransactorRaw struct {
	Contract *WorldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorld creates a new instance of World, bound to a specific deployed contract.
func NewWorld(address common.Address, backend bind.ContractBackend) (*World, error) {
	contract, err := bindWorld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &World{WorldCaller: WorldCaller{contract: contract}, WorldTransactor: WorldTransactor{contract: contract}, WorldFilterer: WorldFilterer{contract: contract}}, nil
}

// NewWorldCaller creates a new read-only instance of World, bound to a specific deployed contract.
func NewWorldCaller(address common.Address, caller bind.ContractCaller) (*WorldCaller, error) {
	contract, err := bindWorld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorldCaller{contract: contract}, nil
}

// NewWorldTransactor creates a new write-only instance of World, bound to a specific deployed contract.
func NewWorldTransactor(address common.Address, transactor bind.ContractTransactor) (*WorldTransactor, error) {
	contract, err := bindWorld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorldTransactor{contract: contract}, nil
}

// NewWorldFilterer creates a new log filterer instance of World, bound to a specific deployed contract.
func NewWorldFilterer(address common.Address, filterer bind.ContractFilterer) (*WorldFilterer, error) {
	contract, err := bindWorld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorldFilterer{contract: contract}, nil
}

// bindWorld binds a generic wrapper to an already deployed contract.
func bindWorld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorldMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_World *WorldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _World.Contract.WorldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_World *WorldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.Contract.WorldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_World *WorldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _World.Contract.WorldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_World *WorldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _World.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_World *WorldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_World *WorldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _World.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_World *WorldCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_World *WorldSession) Creator() (common.Address, error) {
	return _World.Contract.Creator(&_World.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_World *WorldCallerSession) Creator() (common.Address, error) {
	return _World.Contract.Creator(&_World.CallOpts)
}

// EvefrontierAdminSupportOrDirectOwner is a free data retrieval call binding the contract method 0x746be375.
//
// Solidity: function evefrontier__adminSupportOrDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierAdminSupportOrDirectOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__adminSupportOrDirectOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierAdminSupportOrDirectOwner is a free data retrieval call binding the contract method 0x746be375.
//
// Solidity: function evefrontier__adminSupportOrDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierAdminSupportOrDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierAdminSupportOrDirectOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierAdminSupportOrDirectOwner is a free data retrieval call binding the contract method 0x746be375.
//
// Solidity: function evefrontier__adminSupportOrDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierAdminSupportOrDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierAdminSupportOrDirectOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierAdminSupportOrDirectOwnerGates is a free data retrieval call binding the contract method 0xed8913bd.
//
// Solidity: function evefrontier__adminSupportOrDirectOwnerGates(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierAdminSupportOrDirectOwnerGates(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__adminSupportOrDirectOwnerGates", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierAdminSupportOrDirectOwnerGates is a free data retrieval call binding the contract method 0xed8913bd.
//
// Solidity: function evefrontier__adminSupportOrDirectOwnerGates(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierAdminSupportOrDirectOwnerGates(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierAdminSupportOrDirectOwnerGates(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierAdminSupportOrDirectOwnerGates is a free data retrieval call binding the contract method 0xed8913bd.
//
// Solidity: function evefrontier__adminSupportOrDirectOwnerGates(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierAdminSupportOrDirectOwnerGates(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierAdminSupportOrDirectOwnerGates(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierAreGatesOnline is a free data retrieval call binding the contract method 0xad1780b3.
//
// Solidity: function evefrontier__areGatesOnline(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCaller) EvefrontierAreGatesOnline(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__areGatesOnline", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierAreGatesOnline is a free data retrieval call binding the contract method 0xad1780b3.
//
// Solidity: function evefrontier__areGatesOnline(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldSession) EvefrontierAreGatesOnline(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EvefrontierAreGatesOnline(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EvefrontierAreGatesOnline is a free data retrieval call binding the contract method 0xad1780b3.
//
// Solidity: function evefrontier__areGatesOnline(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCallerSession) EvefrontierAreGatesOnline(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EvefrontierAreGatesOnline(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EvefrontierCanCrossTransferToEphemeral is a free data retrieval call binding the contract method 0xa0fc493b.
//
// Solidity: function evefrontier__canCrossTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCaller) EvefrontierCanCrossTransferToEphemeral(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__canCrossTransferToEphemeral", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierCanCrossTransferToEphemeral is a free data retrieval call binding the contract method 0xa0fc493b.
//
// Solidity: function evefrontier__canCrossTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldSession) EvefrontierCanCrossTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierCanCrossTransferToEphemeral(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierCanCrossTransferToEphemeral is a free data retrieval call binding the contract method 0xa0fc493b.
//
// Solidity: function evefrontier__canCrossTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCallerSession) EvefrontierCanCrossTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierCanCrossTransferToEphemeral(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierCanTransferFromEphemeral is a free data retrieval call binding the contract method 0xea0aa7e7.
//
// Solidity: function evefrontier__canTransferFromEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCaller) EvefrontierCanTransferFromEphemeral(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__canTransferFromEphemeral", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierCanTransferFromEphemeral is a free data retrieval call binding the contract method 0xea0aa7e7.
//
// Solidity: function evefrontier__canTransferFromEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldSession) EvefrontierCanTransferFromEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierCanTransferFromEphemeral(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierCanTransferFromEphemeral is a free data retrieval call binding the contract method 0xea0aa7e7.
//
// Solidity: function evefrontier__canTransferFromEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCallerSession) EvefrontierCanTransferFromEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierCanTransferFromEphemeral(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierCanTransferToEphemeral is a free data retrieval call binding the contract method 0x971cd049.
//
// Solidity: function evefrontier__canTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCaller) EvefrontierCanTransferToEphemeral(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__canTransferToEphemeral", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierCanTransferToEphemeral is a free data retrieval call binding the contract method 0x971cd049.
//
// Solidity: function evefrontier__canTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldSession) EvefrontierCanTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierCanTransferToEphemeral(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierCanTransferToEphemeral is a free data retrieval call binding the contract method 0x971cd049.
//
// Solidity: function evefrontier__canTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCallerSession) EvefrontierCanTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierCanTransferToEphemeral(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierCanTransferToInventory is a free data retrieval call binding the contract method 0x11bc02d9.
//
// Solidity: function evefrontier__canTransferToInventory(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCaller) EvefrontierCanTransferToInventory(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__canTransferToInventory", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierCanTransferToInventory is a free data retrieval call binding the contract method 0x11bc02d9.
//
// Solidity: function evefrontier__canTransferToInventory(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldSession) EvefrontierCanTransferToInventory(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierCanTransferToInventory(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierCanTransferToInventory is a free data retrieval call binding the contract method 0x11bc02d9.
//
// Solidity: function evefrontier__canTransferToInventory(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCallerSession) EvefrontierCanTransferToInventory(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierCanTransferToInventory(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierGetCurrentFuelConsumptionStatus is a free data retrieval call binding the contract method 0x711b5650.
//
// Solidity: function evefrontier__getCurrentFuelConsumptionStatus(uint256 smartObjectId) view returns(uint256 elapsedTime, uint256 unitsToConsume, uint256 actualConsumptionRateInSeconds, uint256 fuelAmount)
func (_World *WorldCaller) EvefrontierGetCurrentFuelConsumptionStatus(opts *bind.CallOpts, smartObjectId *big.Int) (struct {
	ElapsedTime                    *big.Int
	UnitsToConsume                 *big.Int
	ActualConsumptionRateInSeconds *big.Int
	FuelAmount                     *big.Int
}, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__getCurrentFuelConsumptionStatus", smartObjectId)

	outstruct := new(struct {
		ElapsedTime                    *big.Int
		UnitsToConsume                 *big.Int
		ActualConsumptionRateInSeconds *big.Int
		FuelAmount                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ElapsedTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnitsToConsume = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ActualConsumptionRateInSeconds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FuelAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EvefrontierGetCurrentFuelConsumptionStatus is a free data retrieval call binding the contract method 0x711b5650.
//
// Solidity: function evefrontier__getCurrentFuelConsumptionStatus(uint256 smartObjectId) view returns(uint256 elapsedTime, uint256 unitsToConsume, uint256 actualConsumptionRateInSeconds, uint256 fuelAmount)
func (_World *WorldSession) EvefrontierGetCurrentFuelConsumptionStatus(smartObjectId *big.Int) (struct {
	ElapsedTime                    *big.Int
	UnitsToConsume                 *big.Int
	ActualConsumptionRateInSeconds *big.Int
	FuelAmount                     *big.Int
}, error) {
	return _World.Contract.EvefrontierGetCurrentFuelConsumptionStatus(&_World.CallOpts, smartObjectId)
}

// EvefrontierGetCurrentFuelConsumptionStatus is a free data retrieval call binding the contract method 0x711b5650.
//
// Solidity: function evefrontier__getCurrentFuelConsumptionStatus(uint256 smartObjectId) view returns(uint256 elapsedTime, uint256 unitsToConsume, uint256 actualConsumptionRateInSeconds, uint256 fuelAmount)
func (_World *WorldCallerSession) EvefrontierGetCurrentFuelConsumptionStatus(smartObjectId *big.Int) (struct {
	ElapsedTime                    *big.Int
	UnitsToConsume                 *big.Int
	ActualConsumptionRateInSeconds *big.Int
	FuelAmount                     *big.Int
}, error) {
	return _World.Contract.EvefrontierGetCurrentFuelConsumptionStatus(&_World.CallOpts, smartObjectId)
}

// EvefrontierGetDeployableClassId is a free data retrieval call binding the contract method 0xda5351c6.
//
// Solidity: function evefrontier__getDeployableClassId() view returns(uint256)
func (_World *WorldCaller) EvefrontierGetDeployableClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__getDeployableClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EvefrontierGetDeployableClassId is a free data retrieval call binding the contract method 0xda5351c6.
//
// Solidity: function evefrontier__getDeployableClassId() view returns(uint256)
func (_World *WorldSession) EvefrontierGetDeployableClassId() (*big.Int, error) {
	return _World.Contract.EvefrontierGetDeployableClassId(&_World.CallOpts)
}

// EvefrontierGetDeployableClassId is a free data retrieval call binding the contract method 0xda5351c6.
//
// Solidity: function evefrontier__getDeployableClassId() view returns(uint256)
func (_World *WorldCallerSession) EvefrontierGetDeployableClassId() (*big.Int, error) {
	return _World.Contract.EvefrontierGetDeployableClassId(&_World.CallOpts)
}

// EvefrontierGetEphemeralOwner is a free data retrieval call binding the contract method 0xbc4c917d.
//
// Solidity: function evefrontier__getEphemeralOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_World *WorldCaller) EvefrontierGetEphemeralOwner(opts *bind.CallOpts, inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__getEphemeralOwner", inventoryObjectId, itemObjectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EvefrontierGetEphemeralOwner is a free data retrieval call binding the contract method 0xbc4c917d.
//
// Solidity: function evefrontier__getEphemeralOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_World *WorldSession) EvefrontierGetEphemeralOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _World.Contract.EvefrontierGetEphemeralOwner(&_World.CallOpts, inventoryObjectId, itemObjectId)
}

// EvefrontierGetEphemeralOwner is a free data retrieval call binding the contract method 0xbc4c917d.
//
// Solidity: function evefrontier__getEphemeralOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_World *WorldCallerSession) EvefrontierGetEphemeralOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _World.Contract.EvefrontierGetEphemeralOwner(&_World.CallOpts, inventoryObjectId, itemObjectId)
}

// EvefrontierGetEphemeralSmartObjectId is a free data retrieval call binding the contract method 0x9da4f12e.
//
// Solidity: function evefrontier__getEphemeralSmartObjectId(uint256 smartObjectId, address ephemeralOwner) pure returns(uint256)
func (_World *WorldCaller) EvefrontierGetEphemeralSmartObjectId(opts *bind.CallOpts, smartObjectId *big.Int, ephemeralOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__getEphemeralSmartObjectId", smartObjectId, ephemeralOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EvefrontierGetEphemeralSmartObjectId is a free data retrieval call binding the contract method 0x9da4f12e.
//
// Solidity: function evefrontier__getEphemeralSmartObjectId(uint256 smartObjectId, address ephemeralOwner) pure returns(uint256)
func (_World *WorldSession) EvefrontierGetEphemeralSmartObjectId(smartObjectId *big.Int, ephemeralOwner common.Address) (*big.Int, error) {
	return _World.Contract.EvefrontierGetEphemeralSmartObjectId(&_World.CallOpts, smartObjectId, ephemeralOwner)
}

// EvefrontierGetEphemeralSmartObjectId is a free data retrieval call binding the contract method 0x9da4f12e.
//
// Solidity: function evefrontier__getEphemeralSmartObjectId(uint256 smartObjectId, address ephemeralOwner) pure returns(uint256)
func (_World *WorldCallerSession) EvefrontierGetEphemeralSmartObjectId(smartObjectId *big.Int, ephemeralOwner common.Address) (*big.Int, error) {
	return _World.Contract.EvefrontierGetEphemeralSmartObjectId(&_World.CallOpts, smartObjectId, ephemeralOwner)
}

// EvefrontierGetInventoryOwner is a free data retrieval call binding the contract method 0x3b570bc4.
//
// Solidity: function evefrontier__getInventoryOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_World *WorldCaller) EvefrontierGetInventoryOwner(opts *bind.CallOpts, inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__getInventoryOwner", inventoryObjectId, itemObjectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EvefrontierGetInventoryOwner is a free data retrieval call binding the contract method 0x3b570bc4.
//
// Solidity: function evefrontier__getInventoryOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_World *WorldSession) EvefrontierGetInventoryOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _World.Contract.EvefrontierGetInventoryOwner(&_World.CallOpts, inventoryObjectId, itemObjectId)
}

// EvefrontierGetInventoryOwner is a free data retrieval call binding the contract method 0x3b570bc4.
//
// Solidity: function evefrontier__getInventoryOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_World *WorldCallerSession) EvefrontierGetInventoryOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _World.Contract.EvefrontierGetInventoryOwner(&_World.CallOpts, inventoryObjectId, itemObjectId)
}

// EvefrontierGetNetworkNodeClassId is a free data retrieval call binding the contract method 0xaab350f3.
//
// Solidity: function evefrontier__getNetworkNodeClassId() view returns(uint256)
func (_World *WorldCaller) EvefrontierGetNetworkNodeClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__getNetworkNodeClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EvefrontierGetNetworkNodeClassId is a free data retrieval call binding the contract method 0xaab350f3.
//
// Solidity: function evefrontier__getNetworkNodeClassId() view returns(uint256)
func (_World *WorldSession) EvefrontierGetNetworkNodeClassId() (*big.Int, error) {
	return _World.Contract.EvefrontierGetNetworkNodeClassId(&_World.CallOpts)
}

// EvefrontierGetNetworkNodeClassId is a free data retrieval call binding the contract method 0xaab350f3.
//
// Solidity: function evefrontier__getNetworkNodeClassId() view returns(uint256)
func (_World *WorldCallerSession) EvefrontierGetNetworkNodeClassId() (*big.Int, error) {
	return _World.Contract.EvefrontierGetNetworkNodeClassId(&_World.CallOpts)
}

// EvefrontierGetSmartCharacterClassId is a free data retrieval call binding the contract method 0x8e29e41f.
//
// Solidity: function evefrontier__getSmartCharacterClassId() view returns(uint256)
func (_World *WorldCaller) EvefrontierGetSmartCharacterClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__getSmartCharacterClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EvefrontierGetSmartCharacterClassId is a free data retrieval call binding the contract method 0x8e29e41f.
//
// Solidity: function evefrontier__getSmartCharacterClassId() view returns(uint256)
func (_World *WorldSession) EvefrontierGetSmartCharacterClassId() (*big.Int, error) {
	return _World.Contract.EvefrontierGetSmartCharacterClassId(&_World.CallOpts)
}

// EvefrontierGetSmartCharacterClassId is a free data retrieval call binding the contract method 0x8e29e41f.
//
// Solidity: function evefrontier__getSmartCharacterClassId() view returns(uint256)
func (_World *WorldCallerSession) EvefrontierGetSmartCharacterClassId() (*big.Int, error) {
	return _World.Contract.EvefrontierGetSmartCharacterClassId(&_World.CallOpts)
}

// EvefrontierIsAdmin is a free data retrieval call binding the contract method 0x20c979b3.
//
// Solidity: function evefrontier__isAdmin(address caller) view returns(bool)
func (_World *WorldCaller) EvefrontierIsAdmin(opts *bind.CallOpts, caller common.Address) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__isAdmin", caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierIsAdmin is a free data retrieval call binding the contract method 0x20c979b3.
//
// Solidity: function evefrontier__isAdmin(address caller) view returns(bool)
func (_World *WorldSession) EvefrontierIsAdmin(caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierIsAdmin(&_World.CallOpts, caller)
}

// EvefrontierIsAdmin is a free data retrieval call binding the contract method 0x20c979b3.
//
// Solidity: function evefrontier__isAdmin(address caller) view returns(bool)
func (_World *WorldCallerSession) EvefrontierIsAdmin(caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierIsAdmin(&_World.CallOpts, caller)
}

// EvefrontierIsAnyGateLinked is a free data retrieval call binding the contract method 0x3ef67471.
//
// Solidity: function evefrontier__isAnyGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCaller) EvefrontierIsAnyGateLinked(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__isAnyGateLinked", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierIsAnyGateLinked is a free data retrieval call binding the contract method 0x3ef67471.
//
// Solidity: function evefrontier__isAnyGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldSession) EvefrontierIsAnyGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EvefrontierIsAnyGateLinked(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EvefrontierIsAnyGateLinked is a free data retrieval call binding the contract method 0x3ef67471.
//
// Solidity: function evefrontier__isAnyGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCallerSession) EvefrontierIsAnyGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EvefrontierIsAnyGateLinked(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EvefrontierIsClassScoped is a free data retrieval call binding the contract method 0x7c5f6493.
//
// Solidity: function evefrontier__isClassScoped(uint256 classId, bytes32 systemId) view returns(bool)
func (_World *WorldCaller) EvefrontierIsClassScoped(opts *bind.CallOpts, classId *big.Int, systemId [32]byte) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__isClassScoped", classId, systemId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierIsClassScoped is a free data retrieval call binding the contract method 0x7c5f6493.
//
// Solidity: function evefrontier__isClassScoped(uint256 classId, bytes32 systemId) view returns(bool)
func (_World *WorldSession) EvefrontierIsClassScoped(classId *big.Int, systemId [32]byte) (bool, error) {
	return _World.Contract.EvefrontierIsClassScoped(&_World.CallOpts, classId, systemId)
}

// EvefrontierIsClassScoped is a free data retrieval call binding the contract method 0x7c5f6493.
//
// Solidity: function evefrontier__isClassScoped(uint256 classId, bytes32 systemId) view returns(bool)
func (_World *WorldCallerSession) EvefrontierIsClassScoped(classId *big.Int, systemId [32]byte) (bool, error) {
	return _World.Contract.EvefrontierIsClassScoped(&_World.CallOpts, classId, systemId)
}

// EvefrontierIsEphemeralOwner is a free data retrieval call binding the contract method 0x4ab4920e.
//
// Solidity: function evefrontier__isEphemeralOwner(uint256 smartObjectId, address caller, bytes data) view returns(bool)
func (_World *WorldCaller) EvefrontierIsEphemeralOwner(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__isEphemeralOwner", smartObjectId, caller, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierIsEphemeralOwner is a free data retrieval call binding the contract method 0x4ab4920e.
//
// Solidity: function evefrontier__isEphemeralOwner(uint256 smartObjectId, address caller, bytes data) view returns(bool)
func (_World *WorldSession) EvefrontierIsEphemeralOwner(smartObjectId *big.Int, caller common.Address, data []byte) (bool, error) {
	return _World.Contract.EvefrontierIsEphemeralOwner(&_World.CallOpts, smartObjectId, caller, data)
}

// EvefrontierIsEphemeralOwner is a free data retrieval call binding the contract method 0x4ab4920e.
//
// Solidity: function evefrontier__isEphemeralOwner(uint256 smartObjectId, address caller, bytes data) view returns(bool)
func (_World *WorldCallerSession) EvefrontierIsEphemeralOwner(smartObjectId *big.Int, caller common.Address, data []byte) (bool, error) {
	return _World.Contract.EvefrontierIsEphemeralOwner(&_World.CallOpts, smartObjectId, caller, data)
}

// EvefrontierIsGateLinked is a free data retrieval call binding the contract method 0xd9af843a.
//
// Solidity: function evefrontier__isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCaller) EvefrontierIsGateLinked(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__isGateLinked", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierIsGateLinked is a free data retrieval call binding the contract method 0xd9af843a.
//
// Solidity: function evefrontier__isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldSession) EvefrontierIsGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EvefrontierIsGateLinked(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EvefrontierIsGateLinked is a free data retrieval call binding the contract method 0xd9af843a.
//
// Solidity: function evefrontier__isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCallerSession) EvefrontierIsGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EvefrontierIsGateLinked(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EvefrontierIsOwner is a free data retrieval call binding the contract method 0xe714222b.
//
// Solidity: function evefrontier__isOwner(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCaller) EvefrontierIsOwner(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__isOwner", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierIsOwner is a free data retrieval call binding the contract method 0xe714222b.
//
// Solidity: function evefrontier__isOwner(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldSession) EvefrontierIsOwner(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierIsOwner(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierIsOwner is a free data retrieval call binding the contract method 0xe714222b.
//
// Solidity: function evefrontier__isOwner(uint256 smartObjectId, address caller) view returns(bool)
func (_World *WorldCallerSession) EvefrontierIsOwner(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _World.Contract.EvefrontierIsOwner(&_World.CallOpts, smartObjectId, caller)
}

// EvefrontierIsOwnerOfBothGates is a free data retrieval call binding the contract method 0x2e37052f.
//
// Solidity: function evefrontier__isOwnerOfBothGates(address caller, bytes data) view returns(bool)
func (_World *WorldCaller) EvefrontierIsOwnerOfBothGates(opts *bind.CallOpts, caller common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__isOwnerOfBothGates", caller, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierIsOwnerOfBothGates is a free data retrieval call binding the contract method 0x2e37052f.
//
// Solidity: function evefrontier__isOwnerOfBothGates(address caller, bytes data) view returns(bool)
func (_World *WorldSession) EvefrontierIsOwnerOfBothGates(caller common.Address, data []byte) (bool, error) {
	return _World.Contract.EvefrontierIsOwnerOfBothGates(&_World.CallOpts, caller, data)
}

// EvefrontierIsOwnerOfBothGates is a free data retrieval call binding the contract method 0x2e37052f.
//
// Solidity: function evefrontier__isOwnerOfBothGates(address caller, bytes data) view returns(bool)
func (_World *WorldCallerSession) EvefrontierIsOwnerOfBothGates(caller common.Address, data []byte) (bool, error) {
	return _World.Contract.EvefrontierIsOwnerOfBothGates(&_World.CallOpts, caller, data)
}

// EvefrontierIsWithinRange is a free data retrieval call binding the contract method 0x3bf821c8.
//
// Solidity: function evefrontier__isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCaller) EvefrontierIsWithinRange(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__isWithinRange", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvefrontierIsWithinRange is a free data retrieval call binding the contract method 0x3bf821c8.
//
// Solidity: function evefrontier__isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldSession) EvefrontierIsWithinRange(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EvefrontierIsWithinRange(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EvefrontierIsWithinRange is a free data retrieval call binding the contract method 0x3bf821c8.
//
// Solidity: function evefrontier__isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCallerSession) EvefrontierIsWithinRange(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EvefrontierIsWithinRange(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EvefrontierOnlyAdminOrCallAccess is a free data retrieval call binding the contract method 0x5adad401.
//
// Solidity: function evefrontier__onlyAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyAdminOrCallAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyAdminOrCallAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyAdminOrCallAccess is a free data retrieval call binding the contract method 0x5adad401.
//
// Solidity: function evefrontier__onlyAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrCallAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminOrCallAccess is a free data retrieval call binding the contract method 0x5adad401.
//
// Solidity: function evefrontier__onlyAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrCallAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminOrClassScopedAccess is a free data retrieval call binding the contract method 0xa2b35df2.
//
// Solidity: function evefrontier__onlyAdminOrClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyAdminOrClassScopedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyAdminOrClassScopedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyAdminOrClassScopedAccess is a free data retrieval call binding the contract method 0xa2b35df2.
//
// Solidity: function evefrontier__onlyAdminOrClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyAdminOrClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrClassScopedAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminOrClassScopedAccess is a free data retrieval call binding the contract method 0xa2b35df2.
//
// Solidity: function evefrontier__onlyAdminOrClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyAdminOrClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrClassScopedAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminOrOwner is a free data retrieval call binding the contract method 0xf0aa3f76.
//
// Solidity: function evefrontier__onlyAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyAdminOrOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyAdminOrOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyAdminOrOwner is a free data retrieval call binding the contract method 0xf0aa3f76.
//
// Solidity: function evefrontier__onlyAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminOrOwner is a free data retrieval call binding the contract method 0xf0aa3f76.
//
// Solidity: function evefrontier__onlyAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminOrOwnerSupported is a free data retrieval call binding the contract method 0xaa789b46.
//
// Solidity: function evefrontier__onlyAdminOrOwnerSupported(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyAdminOrOwnerSupported(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyAdminOrOwnerSupported", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyAdminOrOwnerSupported is a free data retrieval call binding the contract method 0xaa789b46.
//
// Solidity: function evefrontier__onlyAdminOrOwnerSupported(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyAdminOrOwnerSupported(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrOwnerSupported(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminOrOwnerSupported is a free data retrieval call binding the contract method 0xaa789b46.
//
// Solidity: function evefrontier__onlyAdminOrOwnerSupported(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyAdminOrOwnerSupported(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrOwnerSupported(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminOrScopeEnforcedCall is a free data retrieval call binding the contract method 0x3f29c115.
//
// Solidity: function evefrontier__onlyAdminOrScopeEnforcedCall(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyAdminOrScopeEnforcedCall(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyAdminOrScopeEnforcedCall", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyAdminOrScopeEnforcedCall is a free data retrieval call binding the contract method 0x3f29c115.
//
// Solidity: function evefrontier__onlyAdminOrScopeEnforcedCall(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyAdminOrScopeEnforcedCall(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrScopeEnforcedCall(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminOrScopeEnforcedCall is a free data retrieval call binding the contract method 0x3f29c115.
//
// Solidity: function evefrontier__onlyAdminOrScopeEnforcedCall(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyAdminOrScopeEnforcedCall(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminOrScopeEnforcedCall(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminSupportedAccess is a free data retrieval call binding the contract method 0xcd039547.
//
// Solidity: function evefrontier__onlyAdminSupportedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyAdminSupportedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyAdminSupportedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyAdminSupportedAccess is a free data retrieval call binding the contract method 0xcd039547.
//
// Solidity: function evefrontier__onlyAdminSupportedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyAdminSupportedAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminSupportedAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminSupportedAccess is a free data retrieval call binding the contract method 0xcd039547.
//
// Solidity: function evefrontier__onlyAdminSupportedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyAdminSupportedAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminSupportedAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminSupportedOwnerOrCall is a free data retrieval call binding the contract method 0xf5d6c0db.
//
// Solidity: function evefrontier__onlyAdminSupportedOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyAdminSupportedOwnerOrCall(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyAdminSupportedOwnerOrCall", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyAdminSupportedOwnerOrCall is a free data retrieval call binding the contract method 0xf5d6c0db.
//
// Solidity: function evefrontier__onlyAdminSupportedOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyAdminSupportedOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminSupportedOwnerOrCall(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyAdminSupportedOwnerOrCall is a free data retrieval call binding the contract method 0xf5d6c0db.
//
// Solidity: function evefrontier__onlyAdminSupportedOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyAdminSupportedOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyAdminSupportedOwnerOrCall(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyCallAccess is a free data retrieval call binding the contract method 0x2ed8aeb2.
//
// Solidity: function evefrontier__onlyCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyCallAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyCallAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyCallAccess is a free data retrieval call binding the contract method 0x2ed8aeb2.
//
// Solidity: function evefrontier__onlyCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyCallAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyCallAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyCallAccess is a free data retrieval call binding the contract method 0x2ed8aeb2.
//
// Solidity: function evefrontier__onlyCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyCallAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyCallAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyCallAccessOrDirectEphemeralOwner is a free data retrieval call binding the contract method 0x7f243ee3.
//
// Solidity: function evefrontier__onlyCallAccessOrDirectEphemeralOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyCallAccessOrDirectEphemeralOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyCallAccessOrDirectEphemeralOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyCallAccessOrDirectEphemeralOwner is a free data retrieval call binding the contract method 0x7f243ee3.
//
// Solidity: function evefrontier__onlyCallAccessOrDirectEphemeralOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyCallAccessOrDirectEphemeralOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyCallAccessOrDirectEphemeralOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyCallAccessOrDirectEphemeralOwner is a free data retrieval call binding the contract method 0x7f243ee3.
//
// Solidity: function evefrontier__onlyCallAccessOrDirectEphemeralOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyCallAccessOrDirectEphemeralOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyCallAccessOrDirectEphemeralOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyCallAccessWithScopeEnforced is a free data retrieval call binding the contract method 0xaed734ed.
//
// Solidity: function evefrontier__onlyCallAccessWithScopeEnforced(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyCallAccessWithScopeEnforced(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyCallAccessWithScopeEnforced", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyCallAccessWithScopeEnforced is a free data retrieval call binding the contract method 0xaed734ed.
//
// Solidity: function evefrontier__onlyCallAccessWithScopeEnforced(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyCallAccessWithScopeEnforced(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyCallAccessWithScopeEnforced(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyCallAccessWithScopeEnforced is a free data retrieval call binding the contract method 0xaed734ed.
//
// Solidity: function evefrontier__onlyCallAccessWithScopeEnforced(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyCallAccessWithScopeEnforced(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyCallAccessWithScopeEnforced(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyClassScopedAccess is a free data retrieval call binding the contract method 0x19be9e2a.
//
// Solidity: function evefrontier__onlyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyClassScopedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyClassScopedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyClassScopedAccess is a free data retrieval call binding the contract method 0x19be9e2a.
//
// Solidity: function evefrontier__onlyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyClassScopedAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyClassScopedAccess is a free data retrieval call binding the contract method 0x19be9e2a.
//
// Solidity: function evefrontier__onlyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyClassScopedAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyClassScopedOrCharAdminOrOwner is a free data retrieval call binding the contract method 0x850b497e.
//
// Solidity: function evefrontier__onlyClassScopedOrCharAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyClassScopedOrCharAdminOrOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyClassScopedOrCharAdminOrOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyClassScopedOrCharAdminOrOwner is a free data retrieval call binding the contract method 0x850b497e.
//
// Solidity: function evefrontier__onlyClassScopedOrCharAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyClassScopedOrCharAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyClassScopedOrCharAdminOrOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyClassScopedOrCharAdminOrOwner is a free data retrieval call binding the contract method 0x850b497e.
//
// Solidity: function evefrontier__onlyClassScopedOrCharAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyClassScopedOrCharAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyClassScopedOrCharAdminOrOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyDirectAdmin is a free data retrieval call binding the contract method 0x1525fb7a.
//
// Solidity: function evefrontier__onlyDirectAdmin(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyDirectAdmin(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyDirectAdmin", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyDirectAdmin is a free data retrieval call binding the contract method 0x1525fb7a.
//
// Solidity: function evefrontier__onlyDirectAdmin(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyDirectAdmin(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyDirectAdmin(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyDirectAdmin is a free data retrieval call binding the contract method 0x1525fb7a.
//
// Solidity: function evefrontier__onlyDirectAdmin(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyDirectAdmin(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyDirectAdmin(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyDirectAdminOrCallAccess is a free data retrieval call binding the contract method 0x930f4b0f.
//
// Solidity: function evefrontier__onlyDirectAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyDirectAdminOrCallAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyDirectAdminOrCallAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyDirectAdminOrCallAccess is a free data retrieval call binding the contract method 0x930f4b0f.
//
// Solidity: function evefrontier__onlyDirectAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyDirectAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyDirectAdminOrCallAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyDirectAdminOrCallAccess is a free data retrieval call binding the contract method 0x930f4b0f.
//
// Solidity: function evefrontier__onlyDirectAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyDirectAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyDirectAdminOrCallAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyDirectEphemeralOwnerOrCall is a free data retrieval call binding the contract method 0x9d334dd6.
//
// Solidity: function evefrontier__onlyDirectEphemeralOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyDirectEphemeralOwnerOrCall(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyDirectEphemeralOwnerOrCall", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyDirectEphemeralOwnerOrCall is a free data retrieval call binding the contract method 0x9d334dd6.
//
// Solidity: function evefrontier__onlyDirectEphemeralOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyDirectEphemeralOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyDirectEphemeralOwnerOrCall(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyDirectEphemeralOwnerOrCall is a free data retrieval call binding the contract method 0x9d334dd6.
//
// Solidity: function evefrontier__onlyDirectEphemeralOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyDirectEphemeralOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyDirectEphemeralOwnerOrCall(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyDirectOwner is a free data retrieval call binding the contract method 0x2f0f9300.
//
// Solidity: function evefrontier__onlyDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyDirectOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyDirectOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyDirectOwner is a free data retrieval call binding the contract method 0x2f0f9300.
//
// Solidity: function evefrontier__onlyDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyDirectOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyDirectOwner is a free data retrieval call binding the contract method 0x2f0f9300.
//
// Solidity: function evefrontier__onlyDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyDirectOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyEphemeralOwnerOrTransferRole is a free data retrieval call binding the contract method 0xa8a60d87.
//
// Solidity: function evefrontier__onlyEphemeralOwnerOrTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyEphemeralOwnerOrTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyEphemeralOwnerOrTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyEphemeralOwnerOrTransferRole is a free data retrieval call binding the contract method 0xa8a60d87.
//
// Solidity: function evefrontier__onlyEphemeralOwnerOrTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyEphemeralOwnerOrTransferRole(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyEphemeralOwnerOrTransferRole(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyEphemeralOwnerOrTransferRole is a free data retrieval call binding the contract method 0xa8a60d87.
//
// Solidity: function evefrontier__onlyEphemeralOwnerOrTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyEphemeralOwnerOrTransferRole(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyEphemeralOwnerOrTransferRole(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwner is a free data retrieval call binding the contract method 0x791f0c7e.
//
// Solidity: function evefrontier__onlyOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyOwner is a free data retrieval call binding the contract method 0x791f0c7e.
//
// Solidity: function evefrontier__onlyOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwner is a free data retrieval call binding the contract method 0x791f0c7e.
//
// Solidity: function evefrontier__onlyOwner(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyOwner(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwner(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwnerOrEphemeralCrossTransferRole is a free data retrieval call binding the contract method 0x2e640ef2.
//
// Solidity: function evefrontier__onlyOwnerOrEphemeralCrossTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyOwnerOrEphemeralCrossTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyOwnerOrEphemeralCrossTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyOwnerOrEphemeralCrossTransferRole is a free data retrieval call binding the contract method 0x2e640ef2.
//
// Solidity: function evefrontier__onlyOwnerOrEphemeralCrossTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyOwnerOrEphemeralCrossTransferRole(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwnerOrEphemeralCrossTransferRole(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwnerOrEphemeralCrossTransferRole is a free data retrieval call binding the contract method 0x2e640ef2.
//
// Solidity: function evefrontier__onlyOwnerOrEphemeralCrossTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyOwnerOrEphemeralCrossTransferRole(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwnerOrEphemeralCrossTransferRole(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwnerOrEphemeralTransferRole is a free data retrieval call binding the contract method 0x6ce81a90.
//
// Solidity: function evefrontier__onlyOwnerOrEphemeralTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyOwnerOrEphemeralTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyOwnerOrEphemeralTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyOwnerOrEphemeralTransferRole is a free data retrieval call binding the contract method 0x6ce81a90.
//
// Solidity: function evefrontier__onlyOwnerOrEphemeralTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyOwnerOrEphemeralTransferRole(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwnerOrEphemeralTransferRole(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwnerOrEphemeralTransferRole is a free data retrieval call binding the contract method 0x6ce81a90.
//
// Solidity: function evefrontier__onlyOwnerOrEphemeralTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyOwnerOrEphemeralTransferRole(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwnerOrEphemeralTransferRole(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwnerOrInventoryTransferRole is a free data retrieval call binding the contract method 0x389ee55a.
//
// Solidity: function evefrontier__onlyOwnerOrInventoryTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyOwnerOrInventoryTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyOwnerOrInventoryTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyOwnerOrInventoryTransferRole is a free data retrieval call binding the contract method 0x389ee55a.
//
// Solidity: function evefrontier__onlyOwnerOrInventoryTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyOwnerOrInventoryTransferRole(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwnerOrInventoryTransferRole(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwnerOrInventoryTransferRole is a free data retrieval call binding the contract method 0x389ee55a.
//
// Solidity: function evefrontier__onlyOwnerOrInventoryTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyOwnerOrInventoryTransferRole(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwnerOrInventoryTransferRole(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwnerWithAdminSupportAccess is a free data retrieval call binding the contract method 0xed894f4e.
//
// Solidity: function evefrontier__onlyOwnerWithAdminSupportAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlyOwnerWithAdminSupportAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlyOwnerWithAdminSupportAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlyOwnerWithAdminSupportAccess is a free data retrieval call binding the contract method 0xed894f4e.
//
// Solidity: function evefrontier__onlyOwnerWithAdminSupportAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlyOwnerWithAdminSupportAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwnerWithAdminSupportAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlyOwnerWithAdminSupportAccess is a free data retrieval call binding the contract method 0xed894f4e.
//
// Solidity: function evefrontier__onlyOwnerWithAdminSupportAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlyOwnerWithAdminSupportAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlyOwnerWithAdminSupportAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlySmartAssemblyClassScopedAccess is a free data retrieval call binding the contract method 0x62201c94.
//
// Solidity: function evefrontier__onlySmartAssemblyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCaller) EvefrontierOnlySmartAssemblyClassScopedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__onlySmartAssemblyClassScopedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// EvefrontierOnlySmartAssemblyClassScopedAccess is a free data retrieval call binding the contract method 0x62201c94.
//
// Solidity: function evefrontier__onlySmartAssemblyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldSession) EvefrontierOnlySmartAssemblyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlySmartAssemblyClassScopedAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOnlySmartAssemblyClassScopedAccess is a free data retrieval call binding the contract method 0x62201c94.
//
// Solidity: function evefrontier__onlySmartAssemblyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_World *WorldCallerSession) EvefrontierOnlySmartAssemblyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _World.Contract.EvefrontierOnlySmartAssemblyClassScopedAccess(&_World.CallOpts, smartObjectId, data)
}

// EvefrontierOwner is a free data retrieval call binding the contract method 0xb79c9133.
//
// Solidity: function evefrontier__owner(uint256 smartObjectId) view returns(address)
func (_World *WorldCaller) EvefrontierOwner(opts *bind.CallOpts, smartObjectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "evefrontier__owner", smartObjectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EvefrontierOwner is a free data retrieval call binding the contract method 0xb79c9133.
//
// Solidity: function evefrontier__owner(uint256 smartObjectId) view returns(address)
func (_World *WorldSession) EvefrontierOwner(smartObjectId *big.Int) (common.Address, error) {
	return _World.Contract.EvefrontierOwner(&_World.CallOpts, smartObjectId)
}

// EvefrontierOwner is a free data retrieval call binding the contract method 0xb79c9133.
//
// Solidity: function evefrontier__owner(uint256 smartObjectId) view returns(address)
func (_World *WorldCallerSession) EvefrontierOwner(smartObjectId *big.Int) (common.Address, error) {
	return _World.Contract.EvefrontierOwner(&_World.CallOpts, smartObjectId)
}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_World *WorldCaller) GetDynamicField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getDynamicField", tableId, keyTuple, dynamicFieldIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_World *WorldSession) GetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	return _World.Contract.GetDynamicField(&_World.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_World *WorldCallerSession) GetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	return _World.Contract.GetDynamicField(&_World.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_World *WorldCaller) GetDynamicFieldLength(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getDynamicFieldLength", tableId, keyTuple, dynamicFieldIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_World *WorldSession) GetDynamicFieldLength(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	return _World.Contract.GetDynamicFieldLength(&_World.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_World *WorldCallerSession) GetDynamicFieldLength(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	return _World.Contract.GetDynamicFieldLength(&_World.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_World *WorldCaller) GetDynamicFieldSlice(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getDynamicFieldSlice", tableId, keyTuple, dynamicFieldIndex, start, end)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_World *WorldSession) GetDynamicFieldSlice(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	return _World.Contract.GetDynamicFieldSlice(&_World.CallOpts, tableId, keyTuple, dynamicFieldIndex, start, end)
}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_World *WorldCallerSession) GetDynamicFieldSlice(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	return _World.Contract.GetDynamicFieldSlice(&_World.CallOpts, tableId, keyTuple, dynamicFieldIndex, start, end)
}

// GetField is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_World *WorldCaller) GetField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getField", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetField is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_World *WorldSession) GetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	return _World.Contract.GetField(&_World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetField is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_World *WorldCallerSession) GetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	return _World.Contract.GetField(&_World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetField0 is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_World *WorldCaller) GetField0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getField0", tableId, keyTuple, fieldIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetField0 is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_World *WorldSession) GetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	return _World.Contract.GetField0(&_World.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetField0 is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_World *WorldCallerSession) GetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	return _World.Contract.GetField0(&_World.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_World *WorldCaller) GetFieldLayout(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getFieldLayout", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_World *WorldSession) GetFieldLayout(tableId [32]byte) ([32]byte, error) {
	return _World.Contract.GetFieldLayout(&_World.CallOpts, tableId)
}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_World *WorldCallerSession) GetFieldLayout(tableId [32]byte) ([32]byte, error) {
	return _World.Contract.GetFieldLayout(&_World.CallOpts, tableId)
}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_World *WorldCaller) GetFieldLength(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getFieldLength", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_World *WorldSession) GetFieldLength(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	return _World.Contract.GetFieldLength(&_World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_World *WorldCallerSession) GetFieldLength(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	return _World.Contract.GetFieldLength(&_World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetFieldLength0 is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_World *WorldCaller) GetFieldLength0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getFieldLength0", tableId, keyTuple, fieldIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldLength0 is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_World *WorldSession) GetFieldLength0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	return _World.Contract.GetFieldLength0(&_World.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetFieldLength0 is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_World *WorldCallerSession) GetFieldLength0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	return _World.Contract.GetFieldLength0(&_World.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_World *WorldCaller) GetKeySchema(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getKeySchema", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_World *WorldSession) GetKeySchema(tableId [32]byte) ([32]byte, error) {
	return _World.Contract.GetKeySchema(&_World.CallOpts, tableId)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_World *WorldCallerSession) GetKeySchema(tableId [32]byte) ([32]byte, error) {
	return _World.Contract.GetKeySchema(&_World.CallOpts, tableId)
}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_World *WorldCaller) GetRecord(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getRecord", tableId, keyTuple, fieldLayout)

	outstruct := new(struct {
		StaticData     []byte
		EncodedLengths [32]byte
		DynamicData    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StaticData = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.EncodedLengths = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.DynamicData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_World *WorldSession) GetRecord(tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _World.Contract.GetRecord(&_World.CallOpts, tableId, keyTuple, fieldLayout)
}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_World *WorldCallerSession) GetRecord(tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _World.Contract.GetRecord(&_World.CallOpts, tableId, keyTuple, fieldLayout)
}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_World *WorldCaller) GetRecord0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getRecord0", tableId, keyTuple)

	outstruct := new(struct {
		StaticData     []byte
		EncodedLengths [32]byte
		DynamicData    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StaticData = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.EncodedLengths = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.DynamicData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_World *WorldSession) GetRecord0(tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _World.Contract.GetRecord0(&_World.CallOpts, tableId, keyTuple)
}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_World *WorldCallerSession) GetRecord0(tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _World.Contract.GetRecord0(&_World.CallOpts, tableId, keyTuple)
}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_World *WorldCaller) GetStaticField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getStaticField", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_World *WorldSession) GetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	return _World.Contract.GetStaticField(&_World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_World *WorldCallerSession) GetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	return _World.Contract.GetStaticField(&_World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_World *WorldCaller) GetValueSchema(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "getValueSchema", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_World *WorldSession) GetValueSchema(tableId [32]byte) ([32]byte, error) {
	return _World.Contract.GetValueSchema(&_World.CallOpts, tableId)
}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_World *WorldCallerSession) GetValueSchema(tableId [32]byte) ([32]byte, error) {
	return _World.Contract.GetValueSchema(&_World.CallOpts, tableId)
}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_World *WorldCaller) StoreVersion(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "storeVersion")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_World *WorldSession) StoreVersion() ([32]byte, error) {
	return _World.Contract.StoreVersion(&_World.CallOpts)
}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_World *WorldCallerSession) StoreVersion() ([32]byte, error) {
	return _World.Contract.StoreVersion(&_World.CallOpts)
}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_World *WorldCaller) WorldVersion(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "worldVersion")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_World *WorldSession) WorldVersion() ([32]byte, error) {
	return _World.Contract.WorldVersion(&_World.CallOpts)
}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_World *WorldCallerSession) WorldVersion() ([32]byte, error) {
	return _World.Contract.WorldVersion(&_World.CallOpts)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_World *WorldTransactor) BatchCall(opts *bind.TransactOpts, systemCalls []SystemCallData) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "batchCall", systemCalls)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_World *WorldSession) BatchCall(systemCalls []SystemCallData) (*types.Transaction, error) {
	return _World.Contract.BatchCall(&_World.TransactOpts, systemCalls)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_World *WorldTransactorSession) BatchCall(systemCalls []SystemCallData) (*types.Transaction, error) {
	return _World.Contract.BatchCall(&_World.TransactOpts, systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_World *WorldTransactor) BatchCallFrom(opts *bind.TransactOpts, systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "batchCallFrom", systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_World *WorldSession) BatchCallFrom(systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _World.Contract.BatchCallFrom(&_World.TransactOpts, systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_World *WorldTransactorSession) BatchCallFrom(systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _World.Contract.BatchCallFrom(&_World.TransactOpts, systemCalls)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_World *WorldTransactor) Call(opts *bind.TransactOpts, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "call", systemId, callData)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_World *WorldSession) Call(systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _World.Contract.Call(&_World.TransactOpts, systemId, callData)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_World *WorldTransactorSession) Call(systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _World.Contract.Call(&_World.TransactOpts, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_World *WorldTransactor) CallFrom(opts *bind.TransactOpts, delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "callFrom", delegator, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_World *WorldSession) CallFrom(delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _World.Contract.CallFrom(&_World.TransactOpts, delegator, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_World *WorldTransactorSession) CallFrom(delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _World.Contract.CallFrom(&_World.TransactOpts, delegator, systemId, callData)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_World *WorldTransactor) DeleteRecord(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "deleteRecord", tableId, keyTuple)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_World *WorldSession) DeleteRecord(tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _World.Contract.DeleteRecord(&_World.TransactOpts, tableId, keyTuple)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_World *WorldTransactorSession) DeleteRecord(tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _World.Contract.DeleteRecord(&_World.TransactOpts, tableId, keyTuple)
}

// EvefrontierHandleNodeOffline is a paid mutator transaction binding the contract method 0x38d71916.
//
// Solidity: function evefrontier___handleNodeOffline(uint256 networkNodeId) returns()
func (_World *WorldTransactor) EvefrontierHandleNodeOffline(opts *bind.TransactOpts, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier___handleNodeOffline", networkNodeId)
}

// EvefrontierHandleNodeOffline is a paid mutator transaction binding the contract method 0x38d71916.
//
// Solidity: function evefrontier___handleNodeOffline(uint256 networkNodeId) returns()
func (_World *WorldSession) EvefrontierHandleNodeOffline(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierHandleNodeOffline(&_World.TransactOpts, networkNodeId)
}

// EvefrontierHandleNodeOffline is a paid mutator transaction binding the contract method 0x38d71916.
//
// Solidity: function evefrontier___handleNodeOffline(uint256 networkNodeId) returns()
func (_World *WorldTransactorSession) EvefrontierHandleNodeOffline(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierHandleNodeOffline(&_World.TransactOpts, networkNodeId)
}

// EvefrontierAggression is a paid mutator transaction binding the contract method 0x32a4e99b.
//
// Solidity: function evefrontier__aggression((uint256,((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[],(uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256)) params) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldTransactor) EvefrontierAggression(opts *bind.TransactOpts, params AggressionParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__aggression", params)
}

// EvefrontierAggression is a paid mutator transaction binding the contract method 0x32a4e99b.
//
// Solidity: function evefrontier__aggression((uint256,((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[],(uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256)) params) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldSession) EvefrontierAggression(params AggressionParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierAggression(&_World.TransactOpts, params)
}

// EvefrontierAggression is a paid mutator transaction binding the contract method 0x32a4e99b.
//
// Solidity: function evefrontier__aggression((uint256,((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[],(uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256)) params) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldTransactorSession) EvefrontierAggression(params AggressionParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierAggression(&_World.TransactOpts, params)
}

// EvefrontierAnchor is a paid mutator transaction binding the contract method 0xb43e2f20.
//
// Solidity: function evefrontier__anchor(uint256 smartObjectId, address owner, (uint256,uint256,uint256,uint256) locationData) returns()
func (_World *WorldTransactor) EvefrontierAnchor(opts *bind.TransactOpts, smartObjectId *big.Int, owner common.Address, locationData LocationData) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__anchor", smartObjectId, owner, locationData)
}

// EvefrontierAnchor is a paid mutator transaction binding the contract method 0xb43e2f20.
//
// Solidity: function evefrontier__anchor(uint256 smartObjectId, address owner, (uint256,uint256,uint256,uint256) locationData) returns()
func (_World *WorldSession) EvefrontierAnchor(smartObjectId *big.Int, owner common.Address, locationData LocationData) (*types.Transaction, error) {
	return _World.Contract.EvefrontierAnchor(&_World.TransactOpts, smartObjectId, owner, locationData)
}

// EvefrontierAnchor is a paid mutator transaction binding the contract method 0xb43e2f20.
//
// Solidity: function evefrontier__anchor(uint256 smartObjectId, address owner, (uint256,uint256,uint256,uint256) locationData) returns()
func (_World *WorldTransactorSession) EvefrontierAnchor(smartObjectId *big.Int, owner common.Address, locationData LocationData) (*types.Transaction, error) {
	return _World.Contract.EvefrontierAnchor(&_World.TransactOpts, smartObjectId, owner, locationData)
}

// EvefrontierAssignItemToInventory is a paid mutator transaction binding the contract method 0x523ce2d5.
//
// Solidity: function evefrontier__assignItemToInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_World *WorldTransactor) EvefrontierAssignItemToInventory(opts *bind.TransactOpts, inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__assignItemToInventory", inventoryObjectId, itemObjectId, quantity)
}

// EvefrontierAssignItemToInventory is a paid mutator transaction binding the contract method 0x523ce2d5.
//
// Solidity: function evefrontier__assignItemToInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_World *WorldSession) EvefrontierAssignItemToInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierAssignItemToInventory(&_World.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// EvefrontierAssignItemToInventory is a paid mutator transaction binding the contract method 0x523ce2d5.
//
// Solidity: function evefrontier__assignItemToInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_World *WorldTransactorSession) EvefrontierAssignItemToInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierAssignItemToInventory(&_World.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// EvefrontierAssignOwner is a paid mutator transaction binding the contract method 0xaccb037a.
//
// Solidity: function evefrontier__assignOwner(uint256 smartObjectId, address to) returns()
func (_World *WorldTransactor) EvefrontierAssignOwner(opts *bind.TransactOpts, smartObjectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__assignOwner", smartObjectId, to)
}

// EvefrontierAssignOwner is a paid mutator transaction binding the contract method 0xaccb037a.
//
// Solidity: function evefrontier__assignOwner(uint256 smartObjectId, address to) returns()
func (_World *WorldSession) EvefrontierAssignOwner(smartObjectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _World.Contract.EvefrontierAssignOwner(&_World.TransactOpts, smartObjectId, to)
}

// EvefrontierAssignOwner is a paid mutator transaction binding the contract method 0xaccb037a.
//
// Solidity: function evefrontier__assignOwner(uint256 smartObjectId, address to) returns()
func (_World *WorldTransactorSession) EvefrontierAssignOwner(smartObjectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _World.Contract.EvefrontierAssignOwner(&_World.TransactOpts, smartObjectId, to)
}

// EvefrontierBringOffline is a paid mutator transaction binding the contract method 0x918b0ba0.
//
// Solidity: function evefrontier__bringOffline(uint256 smartObjectId) returns()
func (_World *WorldTransactor) EvefrontierBringOffline(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__bringOffline", smartObjectId)
}

// EvefrontierBringOffline is a paid mutator transaction binding the contract method 0x918b0ba0.
//
// Solidity: function evefrontier__bringOffline(uint256 smartObjectId) returns()
func (_World *WorldSession) EvefrontierBringOffline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierBringOffline(&_World.TransactOpts, smartObjectId)
}

// EvefrontierBringOffline is a paid mutator transaction binding the contract method 0x918b0ba0.
//
// Solidity: function evefrontier__bringOffline(uint256 smartObjectId) returns()
func (_World *WorldTransactorSession) EvefrontierBringOffline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierBringOffline(&_World.TransactOpts, smartObjectId)
}

// EvefrontierBringOnline is a paid mutator transaction binding the contract method 0xe47344f8.
//
// Solidity: function evefrontier__bringOnline(uint256 smartObjectId) returns()
func (_World *WorldTransactor) EvefrontierBringOnline(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__bringOnline", smartObjectId)
}

// EvefrontierBringOnline is a paid mutator transaction binding the contract method 0xe47344f8.
//
// Solidity: function evefrontier__bringOnline(uint256 smartObjectId) returns()
func (_World *WorldSession) EvefrontierBringOnline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierBringOnline(&_World.TransactOpts, smartObjectId)
}

// EvefrontierBringOnline is a paid mutator transaction binding the contract method 0xe47344f8.
//
// Solidity: function evefrontier__bringOnline(uint256 smartObjectId) returns()
func (_World *WorldTransactorSession) EvefrontierBringOnline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierBringOnline(&_World.TransactOpts, smartObjectId)
}

// EvefrontierCanJump is a paid mutator transaction binding the contract method 0x287ca681.
//
// Solidity: function evefrontier__canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_World *WorldTransactor) EvefrontierCanJump(opts *bind.TransactOpts, characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__canJump", characterId, sourceGateId, destinationGateId)
}

// EvefrontierCanJump is a paid mutator transaction binding the contract method 0x287ca681.
//
// Solidity: function evefrontier__canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_World *WorldSession) EvefrontierCanJump(characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCanJump(&_World.TransactOpts, characterId, sourceGateId, destinationGateId)
}

// EvefrontierCanJump is a paid mutator transaction binding the contract method 0x287ca681.
//
// Solidity: function evefrontier__canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_World *WorldTransactorSession) EvefrontierCanJump(characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCanJump(&_World.TransactOpts, characterId, sourceGateId, destinationGateId)
}

// EvefrontierConfigureDeployableAccess is a paid mutator transaction binding the contract method 0x783d1442.
//
// Solidity: function evefrontier__configureDeployableAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureDeployableAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureDeployableAccess")
}

// EvefrontierConfigureDeployableAccess is a paid mutator transaction binding the contract method 0x783d1442.
//
// Solidity: function evefrontier__configureDeployableAccess() returns()
func (_World *WorldSession) EvefrontierConfigureDeployableAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureDeployableAccess(&_World.TransactOpts)
}

// EvefrontierConfigureDeployableAccess is a paid mutator transaction binding the contract method 0x783d1442.
//
// Solidity: function evefrontier__configureDeployableAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureDeployableAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureDeployableAccess(&_World.TransactOpts)
}

// EvefrontierConfigureEntityRecordAccess is a paid mutator transaction binding the contract method 0xbb82285c.
//
// Solidity: function evefrontier__configureEntityRecordAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureEntityRecordAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureEntityRecordAccess")
}

// EvefrontierConfigureEntityRecordAccess is a paid mutator transaction binding the contract method 0xbb82285c.
//
// Solidity: function evefrontier__configureEntityRecordAccess() returns()
func (_World *WorldSession) EvefrontierConfigureEntityRecordAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureEntityRecordAccess(&_World.TransactOpts)
}

// EvefrontierConfigureEntityRecordAccess is a paid mutator transaction binding the contract method 0xbb82285c.
//
// Solidity: function evefrontier__configureEntityRecordAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureEntityRecordAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureEntityRecordAccess(&_World.TransactOpts)
}

// EvefrontierConfigureEphemeralInteractAccess is a paid mutator transaction binding the contract method 0x98a66514.
//
// Solidity: function evefrontier__configureEphemeralInteractAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureEphemeralInteractAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureEphemeralInteractAccess")
}

// EvefrontierConfigureEphemeralInteractAccess is a paid mutator transaction binding the contract method 0x98a66514.
//
// Solidity: function evefrontier__configureEphemeralInteractAccess() returns()
func (_World *WorldSession) EvefrontierConfigureEphemeralInteractAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureEphemeralInteractAccess(&_World.TransactOpts)
}

// EvefrontierConfigureEphemeralInteractAccess is a paid mutator transaction binding the contract method 0x98a66514.
//
// Solidity: function evefrontier__configureEphemeralInteractAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureEphemeralInteractAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureEphemeralInteractAccess(&_World.TransactOpts)
}

// EvefrontierConfigureEphemeralInventoryAccess is a paid mutator transaction binding the contract method 0x7beb2dc5.
//
// Solidity: function evefrontier__configureEphemeralInventoryAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureEphemeralInventoryAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureEphemeralInventoryAccess")
}

// EvefrontierConfigureEphemeralInventoryAccess is a paid mutator transaction binding the contract method 0x7beb2dc5.
//
// Solidity: function evefrontier__configureEphemeralInventoryAccess() returns()
func (_World *WorldSession) EvefrontierConfigureEphemeralInventoryAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureEphemeralInventoryAccess(&_World.TransactOpts)
}

// EvefrontierConfigureEphemeralInventoryAccess is a paid mutator transaction binding the contract method 0x7beb2dc5.
//
// Solidity: function evefrontier__configureEphemeralInventoryAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureEphemeralInventoryAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureEphemeralInventoryAccess(&_World.TransactOpts)
}

// EvefrontierConfigureFuelAccess is a paid mutator transaction binding the contract method 0x27341a1d.
//
// Solidity: function evefrontier__configureFuelAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureFuelAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureFuelAccess")
}

// EvefrontierConfigureFuelAccess is a paid mutator transaction binding the contract method 0x27341a1d.
//
// Solidity: function evefrontier__configureFuelAccess() returns()
func (_World *WorldSession) EvefrontierConfigureFuelAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureFuelAccess(&_World.TransactOpts)
}

// EvefrontierConfigureFuelAccess is a paid mutator transaction binding the contract method 0x27341a1d.
//
// Solidity: function evefrontier__configureFuelAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureFuelAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureFuelAccess(&_World.TransactOpts)
}

// EvefrontierConfigureFuelEfficiency is a paid mutator transaction binding the contract method 0xe4136a37.
//
// Solidity: function evefrontier__configureFuelEfficiency(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) fuelEntityParams, uint256 fuelEfficiency) returns()
func (_World *WorldTransactor) EvefrontierConfigureFuelEfficiency(opts *bind.TransactOpts, smartObjectId *big.Int, fuelEntityParams EntityRecordParams, fuelEfficiency *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureFuelEfficiency", smartObjectId, fuelEntityParams, fuelEfficiency)
}

// EvefrontierConfigureFuelEfficiency is a paid mutator transaction binding the contract method 0xe4136a37.
//
// Solidity: function evefrontier__configureFuelEfficiency(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) fuelEntityParams, uint256 fuelEfficiency) returns()
func (_World *WorldSession) EvefrontierConfigureFuelEfficiency(smartObjectId *big.Int, fuelEntityParams EntityRecordParams, fuelEfficiency *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureFuelEfficiency(&_World.TransactOpts, smartObjectId, fuelEntityParams, fuelEfficiency)
}

// EvefrontierConfigureFuelEfficiency is a paid mutator transaction binding the contract method 0xe4136a37.
//
// Solidity: function evefrontier__configureFuelEfficiency(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) fuelEntityParams, uint256 fuelEfficiency) returns()
func (_World *WorldTransactorSession) EvefrontierConfigureFuelEfficiency(smartObjectId *big.Int, fuelEntityParams EntityRecordParams, fuelEfficiency *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureFuelEfficiency(&_World.TransactOpts, smartObjectId, fuelEntityParams, fuelEfficiency)
}

// EvefrontierConfigureFuelParameters is a paid mutator transaction binding the contract method 0x3f107c4c.
//
// Solidity: function evefrontier__configureFuelParameters(uint256 smartObjectId, (uint256,uint256) fuelParams) returns()
func (_World *WorldTransactor) EvefrontierConfigureFuelParameters(opts *bind.TransactOpts, smartObjectId *big.Int, fuelParams FuelParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureFuelParameters", smartObjectId, fuelParams)
}

// EvefrontierConfigureFuelParameters is a paid mutator transaction binding the contract method 0x3f107c4c.
//
// Solidity: function evefrontier__configureFuelParameters(uint256 smartObjectId, (uint256,uint256) fuelParams) returns()
func (_World *WorldSession) EvefrontierConfigureFuelParameters(smartObjectId *big.Int, fuelParams FuelParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureFuelParameters(&_World.TransactOpts, smartObjectId, fuelParams)
}

// EvefrontierConfigureFuelParameters is a paid mutator transaction binding the contract method 0x3f107c4c.
//
// Solidity: function evefrontier__configureFuelParameters(uint256 smartObjectId, (uint256,uint256) fuelParams) returns()
func (_World *WorldTransactorSession) EvefrontierConfigureFuelParameters(smartObjectId *big.Int, fuelParams FuelParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureFuelParameters(&_World.TransactOpts, smartObjectId, fuelParams)
}

// EvefrontierConfigureGate is a paid mutator transaction binding the contract method 0x4646184b.
//
// Solidity: function evefrontier__configureGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldTransactor) EvefrontierConfigureGate(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureGate", smartObjectId, systemId)
}

// EvefrontierConfigureGate is a paid mutator transaction binding the contract method 0x4646184b.
//
// Solidity: function evefrontier__configureGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldSession) EvefrontierConfigureGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureGate(&_World.TransactOpts, smartObjectId, systemId)
}

// EvefrontierConfigureGate is a paid mutator transaction binding the contract method 0x4646184b.
//
// Solidity: function evefrontier__configureGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldTransactorSession) EvefrontierConfigureGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureGate(&_World.TransactOpts, smartObjectId, systemId)
}

// EvefrontierConfigureInventoryAccess is a paid mutator transaction binding the contract method 0x4d4bdfda.
//
// Solidity: function evefrontier__configureInventoryAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureInventoryAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureInventoryAccess")
}

// EvefrontierConfigureInventoryAccess is a paid mutator transaction binding the contract method 0x4d4bdfda.
//
// Solidity: function evefrontier__configureInventoryAccess() returns()
func (_World *WorldSession) EvefrontierConfigureInventoryAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureInventoryAccess(&_World.TransactOpts)
}

// EvefrontierConfigureInventoryAccess is a paid mutator transaction binding the contract method 0x4d4bdfda.
//
// Solidity: function evefrontier__configureInventoryAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureInventoryAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureInventoryAccess(&_World.TransactOpts)
}

// EvefrontierConfigureInventoryInteractAccess is a paid mutator transaction binding the contract method 0x8c22b4f6.
//
// Solidity: function evefrontier__configureInventoryInteractAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureInventoryInteractAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureInventoryInteractAccess")
}

// EvefrontierConfigureInventoryInteractAccess is a paid mutator transaction binding the contract method 0x8c22b4f6.
//
// Solidity: function evefrontier__configureInventoryInteractAccess() returns()
func (_World *WorldSession) EvefrontierConfigureInventoryInteractAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureInventoryInteractAccess(&_World.TransactOpts)
}

// EvefrontierConfigureInventoryInteractAccess is a paid mutator transaction binding the contract method 0x8c22b4f6.
//
// Solidity: function evefrontier__configureInventoryInteractAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureInventoryInteractAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureInventoryInteractAccess(&_World.TransactOpts)
}

// EvefrontierConfigureKillMailAccess is a paid mutator transaction binding the contract method 0xe4216edb.
//
// Solidity: function evefrontier__configureKillMailAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureKillMailAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureKillMailAccess")
}

// EvefrontierConfigureKillMailAccess is a paid mutator transaction binding the contract method 0xe4216edb.
//
// Solidity: function evefrontier__configureKillMailAccess() returns()
func (_World *WorldSession) EvefrontierConfigureKillMailAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureKillMailAccess(&_World.TransactOpts)
}

// EvefrontierConfigureKillMailAccess is a paid mutator transaction binding the contract method 0xe4216edb.
//
// Solidity: function evefrontier__configureKillMailAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureKillMailAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureKillMailAccess(&_World.TransactOpts)
}

// EvefrontierConfigureLocationAccess is a paid mutator transaction binding the contract method 0x17cd377c.
//
// Solidity: function evefrontier__configureLocationAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureLocationAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureLocationAccess")
}

// EvefrontierConfigureLocationAccess is a paid mutator transaction binding the contract method 0x17cd377c.
//
// Solidity: function evefrontier__configureLocationAccess() returns()
func (_World *WorldSession) EvefrontierConfigureLocationAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureLocationAccess(&_World.TransactOpts)
}

// EvefrontierConfigureLocationAccess is a paid mutator transaction binding the contract method 0x17cd377c.
//
// Solidity: function evefrontier__configureLocationAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureLocationAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureLocationAccess(&_World.TransactOpts)
}

// EvefrontierConfigureNetworkNodeAccess is a paid mutator transaction binding the contract method 0x7915fd00.
//
// Solidity: function evefrontier__configureNetworkNodeAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureNetworkNodeAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureNetworkNodeAccess")
}

// EvefrontierConfigureNetworkNodeAccess is a paid mutator transaction binding the contract method 0x7915fd00.
//
// Solidity: function evefrontier__configureNetworkNodeAccess() returns()
func (_World *WorldSession) EvefrontierConfigureNetworkNodeAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureNetworkNodeAccess(&_World.TransactOpts)
}

// EvefrontierConfigureNetworkNodeAccess is a paid mutator transaction binding the contract method 0x7915fd00.
//
// Solidity: function evefrontier__configureNetworkNodeAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureNetworkNodeAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureNetworkNodeAccess(&_World.TransactOpts)
}

// EvefrontierConfigureOwnershipAccess is a paid mutator transaction binding the contract method 0xc197291b.
//
// Solidity: function evefrontier__configureOwnershipAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureOwnershipAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureOwnershipAccess")
}

// EvefrontierConfigureOwnershipAccess is a paid mutator transaction binding the contract method 0xc197291b.
//
// Solidity: function evefrontier__configureOwnershipAccess() returns()
func (_World *WorldSession) EvefrontierConfigureOwnershipAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureOwnershipAccess(&_World.TransactOpts)
}

// EvefrontierConfigureOwnershipAccess is a paid mutator transaction binding the contract method 0xc197291b.
//
// Solidity: function evefrontier__configureOwnershipAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureOwnershipAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureOwnershipAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartAssemblyAccess is a paid mutator transaction binding the contract method 0xbbe175c4.
//
// Solidity: function evefrontier__configureSmartAssemblyAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureSmartAssemblyAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureSmartAssemblyAccess")
}

// EvefrontierConfigureSmartAssemblyAccess is a paid mutator transaction binding the contract method 0xbbe175c4.
//
// Solidity: function evefrontier__configureSmartAssemblyAccess() returns()
func (_World *WorldSession) EvefrontierConfigureSmartAssemblyAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartAssemblyAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartAssemblyAccess is a paid mutator transaction binding the contract method 0xbbe175c4.
//
// Solidity: function evefrontier__configureSmartAssemblyAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureSmartAssemblyAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartAssemblyAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartCharacterAccess is a paid mutator transaction binding the contract method 0x195b4a20.
//
// Solidity: function evefrontier__configureSmartCharacterAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureSmartCharacterAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureSmartCharacterAccess")
}

// EvefrontierConfigureSmartCharacterAccess is a paid mutator transaction binding the contract method 0x195b4a20.
//
// Solidity: function evefrontier__configureSmartCharacterAccess() returns()
func (_World *WorldSession) EvefrontierConfigureSmartCharacterAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartCharacterAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartCharacterAccess is a paid mutator transaction binding the contract method 0x195b4a20.
//
// Solidity: function evefrontier__configureSmartCharacterAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureSmartCharacterAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartCharacterAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartGateAccess is a paid mutator transaction binding the contract method 0x68b34f22.
//
// Solidity: function evefrontier__configureSmartGateAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureSmartGateAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureSmartGateAccess")
}

// EvefrontierConfigureSmartGateAccess is a paid mutator transaction binding the contract method 0x68b34f22.
//
// Solidity: function evefrontier__configureSmartGateAccess() returns()
func (_World *WorldSession) EvefrontierConfigureSmartGateAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartGateAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartGateAccess is a paid mutator transaction binding the contract method 0x68b34f22.
//
// Solidity: function evefrontier__configureSmartGateAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureSmartGateAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartGateAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartStorageUnitAccess is a paid mutator transaction binding the contract method 0x17c4abce.
//
// Solidity: function evefrontier__configureSmartStorageUnitAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureSmartStorageUnitAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureSmartStorageUnitAccess")
}

// EvefrontierConfigureSmartStorageUnitAccess is a paid mutator transaction binding the contract method 0x17c4abce.
//
// Solidity: function evefrontier__configureSmartStorageUnitAccess() returns()
func (_World *WorldSession) EvefrontierConfigureSmartStorageUnitAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartStorageUnitAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartStorageUnitAccess is a paid mutator transaction binding the contract method 0x17c4abce.
//
// Solidity: function evefrontier__configureSmartStorageUnitAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureSmartStorageUnitAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartStorageUnitAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartTurretAccess is a paid mutator transaction binding the contract method 0xca4a3a60.
//
// Solidity: function evefrontier__configureSmartTurretAccess() returns()
func (_World *WorldTransactor) EvefrontierConfigureSmartTurretAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureSmartTurretAccess")
}

// EvefrontierConfigureSmartTurretAccess is a paid mutator transaction binding the contract method 0xca4a3a60.
//
// Solidity: function evefrontier__configureSmartTurretAccess() returns()
func (_World *WorldSession) EvefrontierConfigureSmartTurretAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartTurretAccess(&_World.TransactOpts)
}

// EvefrontierConfigureSmartTurretAccess is a paid mutator transaction binding the contract method 0xca4a3a60.
//
// Solidity: function evefrontier__configureSmartTurretAccess() returns()
func (_World *WorldTransactorSession) EvefrontierConfigureSmartTurretAccess() (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureSmartTurretAccess(&_World.TransactOpts)
}

// EvefrontierConfigureTurret is a paid mutator transaction binding the contract method 0x25a7ae12.
//
// Solidity: function evefrontier__configureTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldTransactor) EvefrontierConfigureTurret(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__configureTurret", smartObjectId, systemId)
}

// EvefrontierConfigureTurret is a paid mutator transaction binding the contract method 0x25a7ae12.
//
// Solidity: function evefrontier__configureTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldSession) EvefrontierConfigureTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureTurret(&_World.TransactOpts, smartObjectId, systemId)
}

// EvefrontierConfigureTurret is a paid mutator transaction binding the contract method 0x25a7ae12.
//
// Solidity: function evefrontier__configureTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldTransactorSession) EvefrontierConfigureTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConfigureTurret(&_World.TransactOpts, smartObjectId, systemId)
}

// EvefrontierConnectAssembly is a paid mutator transaction binding the contract method 0xcc026f2b.
//
// Solidity: function evefrontier__connectAssembly(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldTransactor) EvefrontierConnectAssembly(opts *bind.TransactOpts, networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__connectAssembly", networkNodeId, assemblyId)
}

// EvefrontierConnectAssembly is a paid mutator transaction binding the contract method 0xcc026f2b.
//
// Solidity: function evefrontier__connectAssembly(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldSession) EvefrontierConnectAssembly(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConnectAssembly(&_World.TransactOpts, networkNodeId, assemblyId)
}

// EvefrontierConnectAssembly is a paid mutator transaction binding the contract method 0xcc026f2b.
//
// Solidity: function evefrontier__connectAssembly(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldTransactorSession) EvefrontierConnectAssembly(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierConnectAssembly(&_World.TransactOpts, networkNodeId, assemblyId)
}

// EvefrontierCreateAndAnchor is a paid mutator transaction binding the contract method 0x62b05b35.
//
// Solidity: function evefrontier__createAndAnchor((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_World *WorldTransactor) EvefrontierCreateAndAnchor(opts *bind.TransactOpts, params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createAndAnchor", params, networkNodeId)
}

// EvefrontierCreateAndAnchor is a paid mutator transaction binding the contract method 0x62b05b35.
//
// Solidity: function evefrontier__createAndAnchor((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_World *WorldSession) EvefrontierCreateAndAnchor(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchor(&_World.TransactOpts, params, networkNodeId)
}

// EvefrontierCreateAndAnchor is a paid mutator transaction binding the contract method 0x62b05b35.
//
// Solidity: function evefrontier__createAndAnchor((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_World *WorldTransactorSession) EvefrontierCreateAndAnchor(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchor(&_World.TransactOpts, params, networkNodeId)
}

// EvefrontierCreateAndAnchorGate is a paid mutator transaction binding the contract method 0x02b43f77.
//
// Solidity: function evefrontier__createAndAnchorGate((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 maxDistance, uint256 networkNodeId) returns()
func (_World *WorldTransactor) EvefrontierCreateAndAnchorGate(opts *bind.TransactOpts, params CreateAndAnchorParams, maxDistance *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createAndAnchorGate", params, maxDistance, networkNodeId)
}

// EvefrontierCreateAndAnchorGate is a paid mutator transaction binding the contract method 0x02b43f77.
//
// Solidity: function evefrontier__createAndAnchorGate((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 maxDistance, uint256 networkNodeId) returns()
func (_World *WorldSession) EvefrontierCreateAndAnchorGate(params CreateAndAnchorParams, maxDistance *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchorGate(&_World.TransactOpts, params, maxDistance, networkNodeId)
}

// EvefrontierCreateAndAnchorGate is a paid mutator transaction binding the contract method 0x02b43f77.
//
// Solidity: function evefrontier__createAndAnchorGate((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 maxDistance, uint256 networkNodeId) returns()
func (_World *WorldTransactorSession) EvefrontierCreateAndAnchorGate(params CreateAndAnchorParams, maxDistance *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchorGate(&_World.TransactOpts, params, maxDistance, networkNodeId)
}

// EvefrontierCreateAndAnchorNetworkNode is a paid mutator transaction binding the contract method 0x311c6cec.
//
// Solidity: function evefrontier__createAndAnchorNetworkNode((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, (uint256,uint256) fuelParams, uint256 maxEnergyCapacity, uint256 currentProduction) returns()
func (_World *WorldTransactor) EvefrontierCreateAndAnchorNetworkNode(opts *bind.TransactOpts, params CreateAndAnchorParams, fuelParams FuelParams, maxEnergyCapacity *big.Int, currentProduction *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createAndAnchorNetworkNode", params, fuelParams, maxEnergyCapacity, currentProduction)
}

// EvefrontierCreateAndAnchorNetworkNode is a paid mutator transaction binding the contract method 0x311c6cec.
//
// Solidity: function evefrontier__createAndAnchorNetworkNode((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, (uint256,uint256) fuelParams, uint256 maxEnergyCapacity, uint256 currentProduction) returns()
func (_World *WorldSession) EvefrontierCreateAndAnchorNetworkNode(params CreateAndAnchorParams, fuelParams FuelParams, maxEnergyCapacity *big.Int, currentProduction *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchorNetworkNode(&_World.TransactOpts, params, fuelParams, maxEnergyCapacity, currentProduction)
}

// EvefrontierCreateAndAnchorNetworkNode is a paid mutator transaction binding the contract method 0x311c6cec.
//
// Solidity: function evefrontier__createAndAnchorNetworkNode((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, (uint256,uint256) fuelParams, uint256 maxEnergyCapacity, uint256 currentProduction) returns()
func (_World *WorldTransactorSession) EvefrontierCreateAndAnchorNetworkNode(params CreateAndAnchorParams, fuelParams FuelParams, maxEnergyCapacity *big.Int, currentProduction *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchorNetworkNode(&_World.TransactOpts, params, fuelParams, maxEnergyCapacity, currentProduction)
}

// EvefrontierCreateAndAnchorStorageUnit is a paid mutator transaction binding the contract method 0x4dde8471.
//
// Solidity: function evefrontier__createAndAnchorStorageUnit((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 capacity, uint256 ephemeralCapacity, uint256 networkNodeId) returns()
func (_World *WorldTransactor) EvefrontierCreateAndAnchorStorageUnit(opts *bind.TransactOpts, params CreateAndAnchorParams, capacity *big.Int, ephemeralCapacity *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createAndAnchorStorageUnit", params, capacity, ephemeralCapacity, networkNodeId)
}

// EvefrontierCreateAndAnchorStorageUnit is a paid mutator transaction binding the contract method 0x4dde8471.
//
// Solidity: function evefrontier__createAndAnchorStorageUnit((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 capacity, uint256 ephemeralCapacity, uint256 networkNodeId) returns()
func (_World *WorldSession) EvefrontierCreateAndAnchorStorageUnit(params CreateAndAnchorParams, capacity *big.Int, ephemeralCapacity *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchorStorageUnit(&_World.TransactOpts, params, capacity, ephemeralCapacity, networkNodeId)
}

// EvefrontierCreateAndAnchorStorageUnit is a paid mutator transaction binding the contract method 0x4dde8471.
//
// Solidity: function evefrontier__createAndAnchorStorageUnit((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 capacity, uint256 ephemeralCapacity, uint256 networkNodeId) returns()
func (_World *WorldTransactorSession) EvefrontierCreateAndAnchorStorageUnit(params CreateAndAnchorParams, capacity *big.Int, ephemeralCapacity *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchorStorageUnit(&_World.TransactOpts, params, capacity, ephemeralCapacity, networkNodeId)
}

// EvefrontierCreateAndAnchorTurret is a paid mutator transaction binding the contract method 0xed8fa3f4.
//
// Solidity: function evefrontier__createAndAnchorTurret((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_World *WorldTransactor) EvefrontierCreateAndAnchorTurret(opts *bind.TransactOpts, params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createAndAnchorTurret", params, networkNodeId)
}

// EvefrontierCreateAndAnchorTurret is a paid mutator transaction binding the contract method 0xed8fa3f4.
//
// Solidity: function evefrontier__createAndAnchorTurret((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_World *WorldSession) EvefrontierCreateAndAnchorTurret(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchorTurret(&_World.TransactOpts, params, networkNodeId)
}

// EvefrontierCreateAndAnchorTurret is a paid mutator transaction binding the contract method 0xed8fa3f4.
//
// Solidity: function evefrontier__createAndAnchorTurret((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_World *WorldTransactorSession) EvefrontierCreateAndAnchorTurret(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndAnchorTurret(&_World.TransactOpts, params, networkNodeId)
}

// EvefrontierCreateAndDepositEphemeral is a paid mutator transaction binding the contract method 0x5b6122a8.
//
// Solidity: function evefrontier__createAndDepositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierCreateAndDepositEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createAndDepositEphemeral", smartObjectId, ephemeralOwner, items)
}

// EvefrontierCreateAndDepositEphemeral is a paid mutator transaction binding the contract method 0x5b6122a8.
//
// Solidity: function evefrontier__createAndDepositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierCreateAndDepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndDepositEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierCreateAndDepositEphemeral is a paid mutator transaction binding the contract method 0x5b6122a8.
//
// Solidity: function evefrontier__createAndDepositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierCreateAndDepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndDepositEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierCreateAndDepositInventory is a paid mutator transaction binding the contract method 0xff4fe5fa.
//
// Solidity: function evefrontier__createAndDepositInventory(uint256 smartObjectId, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierCreateAndDepositInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createAndDepositInventory", smartObjectId, items)
}

// EvefrontierCreateAndDepositInventory is a paid mutator transaction binding the contract method 0xff4fe5fa.
//
// Solidity: function evefrontier__createAndDepositInventory(uint256 smartObjectId, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierCreateAndDepositInventory(smartObjectId *big.Int, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndDepositInventory(&_World.TransactOpts, smartObjectId, items)
}

// EvefrontierCreateAndDepositInventory is a paid mutator transaction binding the contract method 0xff4fe5fa.
//
// Solidity: function evefrontier__createAndDepositInventory(uint256 smartObjectId, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierCreateAndDepositInventory(smartObjectId *big.Int, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAndDepositInventory(&_World.TransactOpts, smartObjectId, items)
}

// EvefrontierCreateAssembly is a paid mutator transaction binding the contract method 0x3a057ea4.
//
// Solidity: function evefrontier__createAssembly(uint256 smartObjectId, string assemblyType, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_World *WorldTransactor) EvefrontierCreateAssembly(opts *bind.TransactOpts, smartObjectId *big.Int, assemblyType string, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createAssembly", smartObjectId, assemblyType, entityRecordParams)
}

// EvefrontierCreateAssembly is a paid mutator transaction binding the contract method 0x3a057ea4.
//
// Solidity: function evefrontier__createAssembly(uint256 smartObjectId, string assemblyType, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_World *WorldSession) EvefrontierCreateAssembly(smartObjectId *big.Int, assemblyType string, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAssembly(&_World.TransactOpts, smartObjectId, assemblyType, entityRecordParams)
}

// EvefrontierCreateAssembly is a paid mutator transaction binding the contract method 0x3a057ea4.
//
// Solidity: function evefrontier__createAssembly(uint256 smartObjectId, string assemblyType, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_World *WorldTransactorSession) EvefrontierCreateAssembly(smartObjectId *big.Int, assemblyType string, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateAssembly(&_World.TransactOpts, smartObjectId, assemblyType, entityRecordParams)
}

// EvefrontierCreateCharacter is a paid mutator transaction binding the contract method 0x6279d7c6.
//
// Solidity: function evefrontier__createCharacter(uint256 smartObjectId, address owner, uint256 tribeId, (bytes32,uint256,uint256,uint256) entityRecordParams, (string,string,string) entityRecordMetadata) returns()
func (_World *WorldTransactor) EvefrontierCreateCharacter(opts *bind.TransactOpts, smartObjectId *big.Int, owner common.Address, tribeId *big.Int, entityRecordParams EntityRecordParams, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createCharacter", smartObjectId, owner, tribeId, entityRecordParams, entityRecordMetadata)
}

// EvefrontierCreateCharacter is a paid mutator transaction binding the contract method 0x6279d7c6.
//
// Solidity: function evefrontier__createCharacter(uint256 smartObjectId, address owner, uint256 tribeId, (bytes32,uint256,uint256,uint256) entityRecordParams, (string,string,string) entityRecordMetadata) returns()
func (_World *WorldSession) EvefrontierCreateCharacter(smartObjectId *big.Int, owner common.Address, tribeId *big.Int, entityRecordParams EntityRecordParams, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateCharacter(&_World.TransactOpts, smartObjectId, owner, tribeId, entityRecordParams, entityRecordMetadata)
}

// EvefrontierCreateCharacter is a paid mutator transaction binding the contract method 0x6279d7c6.
//
// Solidity: function evefrontier__createCharacter(uint256 smartObjectId, address owner, uint256 tribeId, (bytes32,uint256,uint256,uint256) entityRecordParams, (string,string,string) entityRecordMetadata) returns()
func (_World *WorldTransactorSession) EvefrontierCreateCharacter(smartObjectId *big.Int, owner common.Address, tribeId *big.Int, entityRecordParams EntityRecordParams, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateCharacter(&_World.TransactOpts, smartObjectId, owner, tribeId, entityRecordParams, entityRecordMetadata)
}

// EvefrontierCreateDeployable is a paid mutator transaction binding the contract method 0x4c866525.
//
// Solidity: function evefrontier__createDeployable(uint256 smartObjectId, address owner) returns()
func (_World *WorldTransactor) EvefrontierCreateDeployable(opts *bind.TransactOpts, smartObjectId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createDeployable", smartObjectId, owner)
}

// EvefrontierCreateDeployable is a paid mutator transaction binding the contract method 0x4c866525.
//
// Solidity: function evefrontier__createDeployable(uint256 smartObjectId, address owner) returns()
func (_World *WorldSession) EvefrontierCreateDeployable(smartObjectId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateDeployable(&_World.TransactOpts, smartObjectId, owner)
}

// EvefrontierCreateDeployable is a paid mutator transaction binding the contract method 0x4c866525.
//
// Solidity: function evefrontier__createDeployable(uint256 smartObjectId, address owner) returns()
func (_World *WorldTransactorSession) EvefrontierCreateDeployable(smartObjectId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateDeployable(&_World.TransactOpts, smartObjectId, owner)
}

// EvefrontierCreateMetadata is a paid mutator transaction binding the contract method 0x35b7b9c2.
//
// Solidity: function evefrontier__createMetadata(uint256 smartObjectId, (string,string,string) entityRecordMetadata) returns()
func (_World *WorldTransactor) EvefrontierCreateMetadata(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createMetadata", smartObjectId, entityRecordMetadata)
}

// EvefrontierCreateMetadata is a paid mutator transaction binding the contract method 0x35b7b9c2.
//
// Solidity: function evefrontier__createMetadata(uint256 smartObjectId, (string,string,string) entityRecordMetadata) returns()
func (_World *WorldSession) EvefrontierCreateMetadata(smartObjectId *big.Int, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateMetadata(&_World.TransactOpts, smartObjectId, entityRecordMetadata)
}

// EvefrontierCreateMetadata is a paid mutator transaction binding the contract method 0x35b7b9c2.
//
// Solidity: function evefrontier__createMetadata(uint256 smartObjectId, (string,string,string) entityRecordMetadata) returns()
func (_World *WorldTransactorSession) EvefrontierCreateMetadata(smartObjectId *big.Int, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateMetadata(&_World.TransactOpts, smartObjectId, entityRecordMetadata)
}

// EvefrontierCreateRecord is a paid mutator transaction binding the contract method 0xa474cc97.
//
// Solidity: function evefrontier__createRecord(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_World *WorldTransactor) EvefrontierCreateRecord(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__createRecord", smartObjectId, entityRecordParams)
}

// EvefrontierCreateRecord is a paid mutator transaction binding the contract method 0xa474cc97.
//
// Solidity: function evefrontier__createRecord(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_World *WorldSession) EvefrontierCreateRecord(smartObjectId *big.Int, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateRecord(&_World.TransactOpts, smartObjectId, entityRecordParams)
}

// EvefrontierCreateRecord is a paid mutator transaction binding the contract method 0xa474cc97.
//
// Solidity: function evefrontier__createRecord(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_World *WorldTransactorSession) EvefrontierCreateRecord(smartObjectId *big.Int, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCreateRecord(&_World.TransactOpts, smartObjectId, entityRecordParams)
}

// EvefrontierCrossTransferToEphemeral is a paid mutator transaction binding the contract method 0x8127ba2a.
//
// Solidity: function evefrontier__crossTransferToEphemeral(uint256 smartObjectId, address fromEphemeralOwner, address toEphemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierCrossTransferToEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, fromEphemeralOwner common.Address, toEphemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__crossTransferToEphemeral", smartObjectId, fromEphemeralOwner, toEphemeralOwner, items)
}

// EvefrontierCrossTransferToEphemeral is a paid mutator transaction binding the contract method 0x8127ba2a.
//
// Solidity: function evefrontier__crossTransferToEphemeral(uint256 smartObjectId, address fromEphemeralOwner, address toEphemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierCrossTransferToEphemeral(smartObjectId *big.Int, fromEphemeralOwner common.Address, toEphemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCrossTransferToEphemeral(&_World.TransactOpts, smartObjectId, fromEphemeralOwner, toEphemeralOwner, items)
}

// EvefrontierCrossTransferToEphemeral is a paid mutator transaction binding the contract method 0x8127ba2a.
//
// Solidity: function evefrontier__crossTransferToEphemeral(uint256 smartObjectId, address fromEphemeralOwner, address toEphemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierCrossTransferToEphemeral(smartObjectId *big.Int, fromEphemeralOwner common.Address, toEphemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierCrossTransferToEphemeral(&_World.TransactOpts, smartObjectId, fromEphemeralOwner, toEphemeralOwner, items)
}

// EvefrontierDepositEphemeral is a paid mutator transaction binding the contract method 0x01721642.
//
// Solidity: function evefrontier__depositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierDepositEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__depositEphemeral", smartObjectId, ephemeralOwner, items)
}

// EvefrontierDepositEphemeral is a paid mutator transaction binding the contract method 0x01721642.
//
// Solidity: function evefrontier__depositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierDepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDepositEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierDepositEphemeral is a paid mutator transaction binding the contract method 0x01721642.
//
// Solidity: function evefrontier__depositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierDepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDepositEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierDepositFuel is a paid mutator transaction binding the contract method 0x7e0ead20.
//
// Solidity: function evefrontier__depositFuel(uint256 smartObjectId, uint256 fuelSmartObjectId, uint256 fuelAmount) returns()
func (_World *WorldTransactor) EvefrontierDepositFuel(opts *bind.TransactOpts, smartObjectId *big.Int, fuelSmartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__depositFuel", smartObjectId, fuelSmartObjectId, fuelAmount)
}

// EvefrontierDepositFuel is a paid mutator transaction binding the contract method 0x7e0ead20.
//
// Solidity: function evefrontier__depositFuel(uint256 smartObjectId, uint256 fuelSmartObjectId, uint256 fuelAmount) returns()
func (_World *WorldSession) EvefrontierDepositFuel(smartObjectId *big.Int, fuelSmartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDepositFuel(&_World.TransactOpts, smartObjectId, fuelSmartObjectId, fuelAmount)
}

// EvefrontierDepositFuel is a paid mutator transaction binding the contract method 0x7e0ead20.
//
// Solidity: function evefrontier__depositFuel(uint256 smartObjectId, uint256 fuelSmartObjectId, uint256 fuelAmount) returns()
func (_World *WorldTransactorSession) EvefrontierDepositFuel(smartObjectId *big.Int, fuelSmartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDepositFuel(&_World.TransactOpts, smartObjectId, fuelSmartObjectId, fuelAmount)
}

// EvefrontierDepositInventory is a paid mutator transaction binding the contract method 0x1a3c8327.
//
// Solidity: function evefrontier__depositInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierDepositInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__depositInventory", smartObjectId, items)
}

// EvefrontierDepositInventory is a paid mutator transaction binding the contract method 0x1a3c8327.
//
// Solidity: function evefrontier__depositInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierDepositInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDepositInventory(&_World.TransactOpts, smartObjectId, items)
}

// EvefrontierDepositInventory is a paid mutator transaction binding the contract method 0x1a3c8327.
//
// Solidity: function evefrontier__depositInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierDepositInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDepositInventory(&_World.TransactOpts, smartObjectId, items)
}

// EvefrontierDestroyDeployable is a paid mutator transaction binding the contract method 0xa1e758dd.
//
// Solidity: function evefrontier__destroyDeployable(uint256 smartObjectId) returns()
func (_World *WorldTransactor) EvefrontierDestroyDeployable(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__destroyDeployable", smartObjectId)
}

// EvefrontierDestroyDeployable is a paid mutator transaction binding the contract method 0xa1e758dd.
//
// Solidity: function evefrontier__destroyDeployable(uint256 smartObjectId) returns()
func (_World *WorldSession) EvefrontierDestroyDeployable(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDestroyDeployable(&_World.TransactOpts, smartObjectId)
}

// EvefrontierDestroyDeployable is a paid mutator transaction binding the contract method 0xa1e758dd.
//
// Solidity: function evefrontier__destroyDeployable(uint256 smartObjectId) returns()
func (_World *WorldTransactorSession) EvefrontierDestroyDeployable(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDestroyDeployable(&_World.TransactOpts, smartObjectId)
}

// EvefrontierDisconnectAssembly is a paid mutator transaction binding the contract method 0xeb4910d9.
//
// Solidity: function evefrontier__disconnectAssembly(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldTransactor) EvefrontierDisconnectAssembly(opts *bind.TransactOpts, networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__disconnectAssembly", networkNodeId, assemblyId)
}

// EvefrontierDisconnectAssembly is a paid mutator transaction binding the contract method 0xeb4910d9.
//
// Solidity: function evefrontier__disconnectAssembly(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldSession) EvefrontierDisconnectAssembly(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDisconnectAssembly(&_World.TransactOpts, networkNodeId, assemblyId)
}

// EvefrontierDisconnectAssembly is a paid mutator transaction binding the contract method 0xeb4910d9.
//
// Solidity: function evefrontier__disconnectAssembly(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldTransactorSession) EvefrontierDisconnectAssembly(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDisconnectAssembly(&_World.TransactOpts, networkNodeId, assemblyId)
}

// EvefrontierDisconnectNetworkNode is a paid mutator transaction binding the contract method 0x0abc00b2.
//
// Solidity: function evefrontier__disconnectNetworkNode(uint256 networkNodeId) returns()
func (_World *WorldTransactor) EvefrontierDisconnectNetworkNode(opts *bind.TransactOpts, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__disconnectNetworkNode", networkNodeId)
}

// EvefrontierDisconnectNetworkNode is a paid mutator transaction binding the contract method 0x0abc00b2.
//
// Solidity: function evefrontier__disconnectNetworkNode(uint256 networkNodeId) returns()
func (_World *WorldSession) EvefrontierDisconnectNetworkNode(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDisconnectNetworkNode(&_World.TransactOpts, networkNodeId)
}

// EvefrontierDisconnectNetworkNode is a paid mutator transaction binding the contract method 0x0abc00b2.
//
// Solidity: function evefrontier__disconnectNetworkNode(uint256 networkNodeId) returns()
func (_World *WorldTransactorSession) EvefrontierDisconnectNetworkNode(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierDisconnectNetworkNode(&_World.TransactOpts, networkNodeId)
}

// EvefrontierInProximity is a paid mutator transaction binding the contract method 0xd515b309.
//
// Solidity: function evefrontier__inProximity(uint256 smartObjectId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldTransactor) EvefrontierInProximity(opts *bind.TransactOpts, smartObjectId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__inProximity", smartObjectId, priorityQueue, turret, turretTarget)
}

// EvefrontierInProximity is a paid mutator transaction binding the contract method 0xd515b309.
//
// Solidity: function evefrontier__inProximity(uint256 smartObjectId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldSession) EvefrontierInProximity(smartObjectId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _World.Contract.EvefrontierInProximity(&_World.TransactOpts, smartObjectId, priorityQueue, turret, turretTarget)
}

// EvefrontierInProximity is a paid mutator transaction binding the contract method 0xd515b309.
//
// Solidity: function evefrontier__inProximity(uint256 smartObjectId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldTransactorSession) EvefrontierInProximity(smartObjectId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _World.Contract.EvefrontierInProximity(&_World.TransactOpts, smartObjectId, priorityQueue, turret, turretTarget)
}

// EvefrontierLinkGates is a paid mutator transaction binding the contract method 0x90ae5fe8.
//
// Solidity: function evefrontier__linkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldTransactor) EvefrontierLinkGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__linkGates", sourceGateId, destinationGateId)
}

// EvefrontierLinkGates is a paid mutator transaction binding the contract method 0x90ae5fe8.
//
// Solidity: function evefrontier__linkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldSession) EvefrontierLinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierLinkGates(&_World.TransactOpts, sourceGateId, destinationGateId)
}

// EvefrontierLinkGates is a paid mutator transaction binding the contract method 0x90ae5fe8.
//
// Solidity: function evefrontier__linkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldTransactorSession) EvefrontierLinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierLinkGates(&_World.TransactOpts, sourceGateId, destinationGateId)
}

// EvefrontierRegisterNetworkNodeClass is a paid mutator transaction binding the contract method 0x53c464d3.
//
// Solidity: function evefrontier__registerNetworkNodeClass(uint256 typeId, uint256 volume) returns()
func (_World *WorldTransactor) EvefrontierRegisterNetworkNodeClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__registerNetworkNodeClass", typeId, volume)
}

// EvefrontierRegisterNetworkNodeClass is a paid mutator transaction binding the contract method 0x53c464d3.
//
// Solidity: function evefrontier__registerNetworkNodeClass(uint256 typeId, uint256 volume) returns()
func (_World *WorldSession) EvefrontierRegisterNetworkNodeClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRegisterNetworkNodeClass(&_World.TransactOpts, typeId, volume)
}

// EvefrontierRegisterNetworkNodeClass is a paid mutator transaction binding the contract method 0x53c464d3.
//
// Solidity: function evefrontier__registerNetworkNodeClass(uint256 typeId, uint256 volume) returns()
func (_World *WorldTransactorSession) EvefrontierRegisterNetworkNodeClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRegisterNetworkNodeClass(&_World.TransactOpts, typeId, volume)
}

// EvefrontierRegisterSmartAssemblies is a paid mutator transaction binding the contract method 0x9356789a.
//
// Solidity: function evefrontier__registerSmartAssemblies(uint256 typeId, uint256 volume) returns()
func (_World *WorldTransactor) EvefrontierRegisterSmartAssemblies(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__registerSmartAssemblies", typeId, volume)
}

// EvefrontierRegisterSmartAssemblies is a paid mutator transaction binding the contract method 0x9356789a.
//
// Solidity: function evefrontier__registerSmartAssemblies(uint256 typeId, uint256 volume) returns()
func (_World *WorldSession) EvefrontierRegisterSmartAssemblies(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRegisterSmartAssemblies(&_World.TransactOpts, typeId, volume)
}

// EvefrontierRegisterSmartAssemblies is a paid mutator transaction binding the contract method 0x9356789a.
//
// Solidity: function evefrontier__registerSmartAssemblies(uint256 typeId, uint256 volume) returns()
func (_World *WorldTransactorSession) EvefrontierRegisterSmartAssemblies(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRegisterSmartAssemblies(&_World.TransactOpts, typeId, volume)
}

// EvefrontierRegisterSmartCharacterClass is a paid mutator transaction binding the contract method 0x8f7b1549.
//
// Solidity: function evefrontier__registerSmartCharacterClass(uint256 typeId, uint256 volume) returns()
func (_World *WorldTransactor) EvefrontierRegisterSmartCharacterClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__registerSmartCharacterClass", typeId, volume)
}

// EvefrontierRegisterSmartCharacterClass is a paid mutator transaction binding the contract method 0x8f7b1549.
//
// Solidity: function evefrontier__registerSmartCharacterClass(uint256 typeId, uint256 volume) returns()
func (_World *WorldSession) EvefrontierRegisterSmartCharacterClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRegisterSmartCharacterClass(&_World.TransactOpts, typeId, volume)
}

// EvefrontierRegisterSmartCharacterClass is a paid mutator transaction binding the contract method 0x8f7b1549.
//
// Solidity: function evefrontier__registerSmartCharacterClass(uint256 typeId, uint256 volume) returns()
func (_World *WorldTransactorSession) EvefrontierRegisterSmartCharacterClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRegisterSmartCharacterClass(&_World.TransactOpts, typeId, volume)
}

// EvefrontierReleaseAssemblyEnergy is a paid mutator transaction binding the contract method 0xc61c081e.
//
// Solidity: function evefrontier__releaseAssemblyEnergy(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldTransactor) EvefrontierReleaseAssemblyEnergy(opts *bind.TransactOpts, networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__releaseAssemblyEnergy", networkNodeId, assemblyId)
}

// EvefrontierReleaseAssemblyEnergy is a paid mutator transaction binding the contract method 0xc61c081e.
//
// Solidity: function evefrontier__releaseAssemblyEnergy(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldSession) EvefrontierReleaseAssemblyEnergy(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReleaseAssemblyEnergy(&_World.TransactOpts, networkNodeId, assemblyId)
}

// EvefrontierReleaseAssemblyEnergy is a paid mutator transaction binding the contract method 0xc61c081e.
//
// Solidity: function evefrontier__releaseAssemblyEnergy(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldTransactorSession) EvefrontierReleaseAssemblyEnergy(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReleaseAssemblyEnergy(&_World.TransactOpts, networkNodeId, assemblyId)
}

// EvefrontierReleaseNetworkNodeEnergy is a paid mutator transaction binding the contract method 0x819471bc.
//
// Solidity: function evefrontier__releaseNetworkNodeEnergy(uint256 networkNodeId) returns()
func (_World *WorldTransactor) EvefrontierReleaseNetworkNodeEnergy(opts *bind.TransactOpts, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__releaseNetworkNodeEnergy", networkNodeId)
}

// EvefrontierReleaseNetworkNodeEnergy is a paid mutator transaction binding the contract method 0x819471bc.
//
// Solidity: function evefrontier__releaseNetworkNodeEnergy(uint256 networkNodeId) returns()
func (_World *WorldSession) EvefrontierReleaseNetworkNodeEnergy(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReleaseNetworkNodeEnergy(&_World.TransactOpts, networkNodeId)
}

// EvefrontierReleaseNetworkNodeEnergy is a paid mutator transaction binding the contract method 0x819471bc.
//
// Solidity: function evefrontier__releaseNetworkNodeEnergy(uint256 networkNodeId) returns()
func (_World *WorldTransactorSession) EvefrontierReleaseNetworkNodeEnergy(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReleaseNetworkNodeEnergy(&_World.TransactOpts, networkNodeId)
}

// EvefrontierRemoveCharacter is a paid mutator transaction binding the contract method 0xe4ada7c4.
//
// Solidity: function evefrontier__removeCharacter(uint256 smartObjectId) returns()
func (_World *WorldTransactor) EvefrontierRemoveCharacter(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__removeCharacter", smartObjectId)
}

// EvefrontierRemoveCharacter is a paid mutator transaction binding the contract method 0xe4ada7c4.
//
// Solidity: function evefrontier__removeCharacter(uint256 smartObjectId) returns()
func (_World *WorldSession) EvefrontierRemoveCharacter(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRemoveCharacter(&_World.TransactOpts, smartObjectId)
}

// EvefrontierRemoveCharacter is a paid mutator transaction binding the contract method 0xe4ada7c4.
//
// Solidity: function evefrontier__removeCharacter(uint256 smartObjectId) returns()
func (_World *WorldTransactorSession) EvefrontierRemoveCharacter(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRemoveCharacter(&_World.TransactOpts, smartObjectId)
}

// EvefrontierRemoveItemFromInventory is a paid mutator transaction binding the contract method 0xf86ebd47.
//
// Solidity: function evefrontier__removeItemFromInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_World *WorldTransactor) EvefrontierRemoveItemFromInventory(opts *bind.TransactOpts, inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__removeItemFromInventory", inventoryObjectId, itemObjectId, quantity)
}

// EvefrontierRemoveItemFromInventory is a paid mutator transaction binding the contract method 0xf86ebd47.
//
// Solidity: function evefrontier__removeItemFromInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_World *WorldSession) EvefrontierRemoveItemFromInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRemoveItemFromInventory(&_World.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// EvefrontierRemoveItemFromInventory is a paid mutator transaction binding the contract method 0xf86ebd47.
//
// Solidity: function evefrontier__removeItemFromInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_World *WorldTransactorSession) EvefrontierRemoveItemFromInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRemoveItemFromInventory(&_World.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// EvefrontierRemoveOwner is a paid mutator transaction binding the contract method 0x2c65361b.
//
// Solidity: function evefrontier__removeOwner(uint256 smartObjectId, address from) returns()
func (_World *WorldTransactor) EvefrontierRemoveOwner(opts *bind.TransactOpts, smartObjectId *big.Int, from common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__removeOwner", smartObjectId, from)
}

// EvefrontierRemoveOwner is a paid mutator transaction binding the contract method 0x2c65361b.
//
// Solidity: function evefrontier__removeOwner(uint256 smartObjectId, address from) returns()
func (_World *WorldSession) EvefrontierRemoveOwner(smartObjectId *big.Int, from common.Address) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRemoveOwner(&_World.TransactOpts, smartObjectId, from)
}

// EvefrontierRemoveOwner is a paid mutator transaction binding the contract method 0x2c65361b.
//
// Solidity: function evefrontier__removeOwner(uint256 smartObjectId, address from) returns()
func (_World *WorldTransactorSession) EvefrontierRemoveOwner(smartObjectId *big.Int, from common.Address) (*types.Transaction, error) {
	return _World.Contract.EvefrontierRemoveOwner(&_World.TransactOpts, smartObjectId, from)
}

// EvefrontierReportKill is a paid mutator transaction binding the contract method 0x4f204e3f.
//
// Solidity: function evefrontier__reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailData) returns()
func (_World *WorldTransactor) EvefrontierReportKill(opts *bind.TransactOpts, killMailId *big.Int, killMailData KillMailData) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__reportKill", killMailId, killMailData)
}

// EvefrontierReportKill is a paid mutator transaction binding the contract method 0x4f204e3f.
//
// Solidity: function evefrontier__reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailData) returns()
func (_World *WorldSession) EvefrontierReportKill(killMailId *big.Int, killMailData KillMailData) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReportKill(&_World.TransactOpts, killMailId, killMailData)
}

// EvefrontierReportKill is a paid mutator transaction binding the contract method 0x4f204e3f.
//
// Solidity: function evefrontier__reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailData) returns()
func (_World *WorldTransactorSession) EvefrontierReportKill(killMailId *big.Int, killMailData KillMailData) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReportKill(&_World.TransactOpts, killMailId, killMailData)
}

// EvefrontierReserveAssemblyEnergy is a paid mutator transaction binding the contract method 0x542f9f6c.
//
// Solidity: function evefrontier__reserveAssemblyEnergy(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldTransactor) EvefrontierReserveAssemblyEnergy(opts *bind.TransactOpts, networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__reserveAssemblyEnergy", networkNodeId, assemblyId)
}

// EvefrontierReserveAssemblyEnergy is a paid mutator transaction binding the contract method 0x542f9f6c.
//
// Solidity: function evefrontier__reserveAssemblyEnergy(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldSession) EvefrontierReserveAssemblyEnergy(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReserveAssemblyEnergy(&_World.TransactOpts, networkNodeId, assemblyId)
}

// EvefrontierReserveAssemblyEnergy is a paid mutator transaction binding the contract method 0x542f9f6c.
//
// Solidity: function evefrontier__reserveAssemblyEnergy(uint256 networkNodeId, uint256 assemblyId) returns()
func (_World *WorldTransactorSession) EvefrontierReserveAssemblyEnergy(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReserveAssemblyEnergy(&_World.TransactOpts, networkNodeId, assemblyId)
}

// EvefrontierReserveNetworkNodeEnergy is a paid mutator transaction binding the contract method 0xbf38b9fa.
//
// Solidity: function evefrontier__reserveNetworkNodeEnergy(uint256 networkNodeId) returns()
func (_World *WorldTransactor) EvefrontierReserveNetworkNodeEnergy(opts *bind.TransactOpts, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__reserveNetworkNodeEnergy", networkNodeId)
}

// EvefrontierReserveNetworkNodeEnergy is a paid mutator transaction binding the contract method 0xbf38b9fa.
//
// Solidity: function evefrontier__reserveNetworkNodeEnergy(uint256 networkNodeId) returns()
func (_World *WorldSession) EvefrontierReserveNetworkNodeEnergy(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReserveNetworkNodeEnergy(&_World.TransactOpts, networkNodeId)
}

// EvefrontierReserveNetworkNodeEnergy is a paid mutator transaction binding the contract method 0xbf38b9fa.
//
// Solidity: function evefrontier__reserveNetworkNodeEnergy(uint256 networkNodeId) returns()
func (_World *WorldTransactorSession) EvefrontierReserveNetworkNodeEnergy(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierReserveNetworkNodeEnergy(&_World.TransactOpts, networkNodeId)
}

// EvefrontierSaveLocation is a paid mutator transaction binding the contract method 0x1470877e.
//
// Solidity: function evefrontier__saveLocation(uint256 smartObjectId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_World *WorldTransactor) EvefrontierSaveLocation(opts *bind.TransactOpts, smartObjectId *big.Int, locationData LocationData) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__saveLocation", smartObjectId, locationData)
}

// EvefrontierSaveLocation is a paid mutator transaction binding the contract method 0x1470877e.
//
// Solidity: function evefrontier__saveLocation(uint256 smartObjectId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_World *WorldSession) EvefrontierSaveLocation(smartObjectId *big.Int, locationData LocationData) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSaveLocation(&_World.TransactOpts, smartObjectId, locationData)
}

// EvefrontierSaveLocation is a paid mutator transaction binding the contract method 0x1470877e.
//
// Solidity: function evefrontier__saveLocation(uint256 smartObjectId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_World *WorldTransactorSession) EvefrontierSaveLocation(smartObjectId *big.Int, locationData LocationData) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSaveLocation(&_World.TransactOpts, smartObjectId, locationData)
}

// EvefrontierSetAssemblyType is a paid mutator transaction binding the contract method 0x0a1d5e46.
//
// Solidity: function evefrontier__setAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_World *WorldTransactor) EvefrontierSetAssemblyType(opts *bind.TransactOpts, smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setAssemblyType", smartObjectId, assemblyType)
}

// EvefrontierSetAssemblyType is a paid mutator transaction binding the contract method 0x0a1d5e46.
//
// Solidity: function evefrontier__setAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_World *WorldSession) EvefrontierSetAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetAssemblyType(&_World.TransactOpts, smartObjectId, assemblyType)
}

// EvefrontierSetAssemblyType is a paid mutator transaction binding the contract method 0x0a1d5e46.
//
// Solidity: function evefrontier__setAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_World *WorldTransactorSession) EvefrontierSetAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetAssemblyType(&_World.TransactOpts, smartObjectId, assemblyType)
}

// EvefrontierSetCapacity is a paid mutator transaction binding the contract method 0xf82cdbd0.
//
// Solidity: function evefrontier__setCapacity(uint256 smartObjectId, uint256 capacity) returns()
func (_World *WorldTransactor) EvefrontierSetCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, capacity *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setCapacity", smartObjectId, capacity)
}

// EvefrontierSetCapacity is a paid mutator transaction binding the contract method 0xf82cdbd0.
//
// Solidity: function evefrontier__setCapacity(uint256 smartObjectId, uint256 capacity) returns()
func (_World *WorldSession) EvefrontierSetCapacity(smartObjectId *big.Int, capacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetCapacity(&_World.TransactOpts, smartObjectId, capacity)
}

// EvefrontierSetCapacity is a paid mutator transaction binding the contract method 0xf82cdbd0.
//
// Solidity: function evefrontier__setCapacity(uint256 smartObjectId, uint256 capacity) returns()
func (_World *WorldTransactorSession) EvefrontierSetCapacity(smartObjectId *big.Int, capacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetCapacity(&_World.TransactOpts, smartObjectId, capacity)
}

// EvefrontierSetCrossTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0x272e071f.
//
// Solidity: function evefrontier__setCrossTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldTransactor) EvefrontierSetCrossTransferToEphemeralAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setCrossTransferToEphemeralAccess", smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetCrossTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0x272e071f.
//
// Solidity: function evefrontier__setCrossTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldSession) EvefrontierSetCrossTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetCrossTransferToEphemeralAccess(&_World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetCrossTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0x272e071f.
//
// Solidity: function evefrontier__setCrossTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldTransactorSession) EvefrontierSetCrossTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetCrossTransferToEphemeralAccess(&_World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetDappURL is a paid mutator transaction binding the contract method 0x9066625a.
//
// Solidity: function evefrontier__setDappURL(uint256 smartObjectId, string dappURL) returns()
func (_World *WorldTransactor) EvefrontierSetDappURL(opts *bind.TransactOpts, smartObjectId *big.Int, dappURL string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setDappURL", smartObjectId, dappURL)
}

// EvefrontierSetDappURL is a paid mutator transaction binding the contract method 0x9066625a.
//
// Solidity: function evefrontier__setDappURL(uint256 smartObjectId, string dappURL) returns()
func (_World *WorldSession) EvefrontierSetDappURL(smartObjectId *big.Int, dappURL string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetDappURL(&_World.TransactOpts, smartObjectId, dappURL)
}

// EvefrontierSetDappURL is a paid mutator transaction binding the contract method 0x9066625a.
//
// Solidity: function evefrontier__setDappURL(uint256 smartObjectId, string dappURL) returns()
func (_World *WorldTransactorSession) EvefrontierSetDappURL(smartObjectId *big.Int, dappURL string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetDappURL(&_World.TransactOpts, smartObjectId, dappURL)
}

// EvefrontierSetDescription is a paid mutator transaction binding the contract method 0x8587cadc.
//
// Solidity: function evefrontier__setDescription(uint256 smartObjectId, string description) returns()
func (_World *WorldTransactor) EvefrontierSetDescription(opts *bind.TransactOpts, smartObjectId *big.Int, description string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setDescription", smartObjectId, description)
}

// EvefrontierSetDescription is a paid mutator transaction binding the contract method 0x8587cadc.
//
// Solidity: function evefrontier__setDescription(uint256 smartObjectId, string description) returns()
func (_World *WorldSession) EvefrontierSetDescription(smartObjectId *big.Int, description string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetDescription(&_World.TransactOpts, smartObjectId, description)
}

// EvefrontierSetDescription is a paid mutator transaction binding the contract method 0x8587cadc.
//
// Solidity: function evefrontier__setDescription(uint256 smartObjectId, string description) returns()
func (_World *WorldTransactorSession) EvefrontierSetDescription(smartObjectId *big.Int, description string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetDescription(&_World.TransactOpts, smartObjectId, description)
}

// EvefrontierSetEphemeralCapacity is a paid mutator transaction binding the contract method 0x92ceb04a.
//
// Solidity: function evefrontier__setEphemeralCapacity(uint256 smartObjectId, uint256 ephemeralCapacity) returns()
func (_World *WorldTransactor) EvefrontierSetEphemeralCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralCapacity *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setEphemeralCapacity", smartObjectId, ephemeralCapacity)
}

// EvefrontierSetEphemeralCapacity is a paid mutator transaction binding the contract method 0x92ceb04a.
//
// Solidity: function evefrontier__setEphemeralCapacity(uint256 smartObjectId, uint256 ephemeralCapacity) returns()
func (_World *WorldSession) EvefrontierSetEphemeralCapacity(smartObjectId *big.Int, ephemeralCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetEphemeralCapacity(&_World.TransactOpts, smartObjectId, ephemeralCapacity)
}

// EvefrontierSetEphemeralCapacity is a paid mutator transaction binding the contract method 0x92ceb04a.
//
// Solidity: function evefrontier__setEphemeralCapacity(uint256 smartObjectId, uint256 ephemeralCapacity) returns()
func (_World *WorldTransactorSession) EvefrontierSetEphemeralCapacity(smartObjectId *big.Int, ephemeralCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetEphemeralCapacity(&_World.TransactOpts, smartObjectId, ephemeralCapacity)
}

// EvefrontierSetFuelMaxCapacity is a paid mutator transaction binding the contract method 0x57cf8783.
//
// Solidity: function evefrontier__setFuelMaxCapacity(uint256 smartObjectId, uint256 fuelMaxCapacity) returns()
func (_World *WorldTransactor) EvefrontierSetFuelMaxCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setFuelMaxCapacity", smartObjectId, fuelMaxCapacity)
}

// EvefrontierSetFuelMaxCapacity is a paid mutator transaction binding the contract method 0x57cf8783.
//
// Solidity: function evefrontier__setFuelMaxCapacity(uint256 smartObjectId, uint256 fuelMaxCapacity) returns()
func (_World *WorldSession) EvefrontierSetFuelMaxCapacity(smartObjectId *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetFuelMaxCapacity(&_World.TransactOpts, smartObjectId, fuelMaxCapacity)
}

// EvefrontierSetFuelMaxCapacity is a paid mutator transaction binding the contract method 0x57cf8783.
//
// Solidity: function evefrontier__setFuelMaxCapacity(uint256 smartObjectId, uint256 fuelMaxCapacity) returns()
func (_World *WorldTransactorSession) EvefrontierSetFuelMaxCapacity(smartObjectId *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetFuelMaxCapacity(&_World.TransactOpts, smartObjectId, fuelMaxCapacity)
}

// EvefrontierSetName is a paid mutator transaction binding the contract method 0xc2b27db2.
//
// Solidity: function evefrontier__setName(uint256 smartObjectId, string name) returns()
func (_World *WorldTransactor) EvefrontierSetName(opts *bind.TransactOpts, smartObjectId *big.Int, name string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setName", smartObjectId, name)
}

// EvefrontierSetName is a paid mutator transaction binding the contract method 0xc2b27db2.
//
// Solidity: function evefrontier__setName(uint256 smartObjectId, string name) returns()
func (_World *WorldSession) EvefrontierSetName(smartObjectId *big.Int, name string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetName(&_World.TransactOpts, smartObjectId, name)
}

// EvefrontierSetName is a paid mutator transaction binding the contract method 0xc2b27db2.
//
// Solidity: function evefrontier__setName(uint256 smartObjectId, string name) returns()
func (_World *WorldTransactorSession) EvefrontierSetName(smartObjectId *big.Int, name string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetName(&_World.TransactOpts, smartObjectId, name)
}

// EvefrontierSetTransferFromEphemeralAccess is a paid mutator transaction binding the contract method 0x2487dc35.
//
// Solidity: function evefrontier__setTransferFromEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldTransactor) EvefrontierSetTransferFromEphemeralAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setTransferFromEphemeralAccess", smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetTransferFromEphemeralAccess is a paid mutator transaction binding the contract method 0x2487dc35.
//
// Solidity: function evefrontier__setTransferFromEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldSession) EvefrontierSetTransferFromEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetTransferFromEphemeralAccess(&_World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetTransferFromEphemeralAccess is a paid mutator transaction binding the contract method 0x2487dc35.
//
// Solidity: function evefrontier__setTransferFromEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldTransactorSession) EvefrontierSetTransferFromEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetTransferFromEphemeralAccess(&_World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0xa5f8de86.
//
// Solidity: function evefrontier__setTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldTransactor) EvefrontierSetTransferToEphemeralAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setTransferToEphemeralAccess", smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0xa5f8de86.
//
// Solidity: function evefrontier__setTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldSession) EvefrontierSetTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetTransferToEphemeralAccess(&_World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0xa5f8de86.
//
// Solidity: function evefrontier__setTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldTransactorSession) EvefrontierSetTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetTransferToEphemeralAccess(&_World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetTransferToInventoryAccess is a paid mutator transaction binding the contract method 0xf96d7230.
//
// Solidity: function evefrontier__setTransferToInventoryAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldTransactor) EvefrontierSetTransferToInventoryAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__setTransferToInventoryAccess", smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetTransferToInventoryAccess is a paid mutator transaction binding the contract method 0xf96d7230.
//
// Solidity: function evefrontier__setTransferToInventoryAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldSession) EvefrontierSetTransferToInventoryAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetTransferToInventoryAccess(&_World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// EvefrontierSetTransferToInventoryAccess is a paid mutator transaction binding the contract method 0xf96d7230.
//
// Solidity: function evefrontier__setTransferToInventoryAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_World *WorldTransactorSession) EvefrontierSetTransferToInventoryAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _World.Contract.EvefrontierSetTransferToInventoryAccess(&_World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// EvefrontierStartBurn is a paid mutator transaction binding the contract method 0xb141a03b.
//
// Solidity: function evefrontier__startBurn(uint256 smartObjectId) returns()
func (_World *WorldTransactor) EvefrontierStartBurn(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__startBurn", smartObjectId)
}

// EvefrontierStartBurn is a paid mutator transaction binding the contract method 0xb141a03b.
//
// Solidity: function evefrontier__startBurn(uint256 smartObjectId) returns()
func (_World *WorldSession) EvefrontierStartBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierStartBurn(&_World.TransactOpts, smartObjectId)
}

// EvefrontierStartBurn is a paid mutator transaction binding the contract method 0xb141a03b.
//
// Solidity: function evefrontier__startBurn(uint256 smartObjectId) returns()
func (_World *WorldTransactorSession) EvefrontierStartBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierStartBurn(&_World.TransactOpts, smartObjectId)
}

// EvefrontierStopBurn is a paid mutator transaction binding the contract method 0xee5a29c6.
//
// Solidity: function evefrontier__stopBurn(uint256 smartObjectId) returns()
func (_World *WorldTransactor) EvefrontierStopBurn(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__stopBurn", smartObjectId)
}

// EvefrontierStopBurn is a paid mutator transaction binding the contract method 0xee5a29c6.
//
// Solidity: function evefrontier__stopBurn(uint256 smartObjectId) returns()
func (_World *WorldSession) EvefrontierStopBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierStopBurn(&_World.TransactOpts, smartObjectId)
}

// EvefrontierStopBurn is a paid mutator transaction binding the contract method 0xee5a29c6.
//
// Solidity: function evefrontier__stopBurn(uint256 smartObjectId) returns()
func (_World *WorldTransactorSession) EvefrontierStopBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierStopBurn(&_World.TransactOpts, smartObjectId)
}

// EvefrontierTransferFromEphemeral is a paid mutator transaction binding the contract method 0xde72908c.
//
// Solidity: function evefrontier__transferFromEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierTransferFromEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__transferFromEphemeral", smartObjectId, ephemeralOwner, items)
}

// EvefrontierTransferFromEphemeral is a paid mutator transaction binding the contract method 0xde72908c.
//
// Solidity: function evefrontier__transferFromEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierTransferFromEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierTransferFromEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierTransferFromEphemeral is a paid mutator transaction binding the contract method 0xde72908c.
//
// Solidity: function evefrontier__transferFromEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierTransferFromEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierTransferFromEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierTransferToEphemeral is a paid mutator transaction binding the contract method 0xa8430a5d.
//
// Solidity: function evefrontier__transferToEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierTransferToEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__transferToEphemeral", smartObjectId, ephemeralOwner, items)
}

// EvefrontierTransferToEphemeral is a paid mutator transaction binding the contract method 0xa8430a5d.
//
// Solidity: function evefrontier__transferToEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierTransferToEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierTransferToEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierTransferToEphemeral is a paid mutator transaction binding the contract method 0xa8430a5d.
//
// Solidity: function evefrontier__transferToEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierTransferToEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierTransferToEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierTransferToInventory is a paid mutator transaction binding the contract method 0xad524293.
//
// Solidity: function evefrontier__transferToInventory(uint256 smartObjectId, uint256 toObjectId, (uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierTransferToInventory(opts *bind.TransactOpts, smartObjectId *big.Int, toObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__transferToInventory", smartObjectId, toObjectId, items)
}

// EvefrontierTransferToInventory is a paid mutator transaction binding the contract method 0xad524293.
//
// Solidity: function evefrontier__transferToInventory(uint256 smartObjectId, uint256 toObjectId, (uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierTransferToInventory(smartObjectId *big.Int, toObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierTransferToInventory(&_World.TransactOpts, smartObjectId, toObjectId, items)
}

// EvefrontierTransferToInventory is a paid mutator transaction binding the contract method 0xad524293.
//
// Solidity: function evefrontier__transferToInventory(uint256 smartObjectId, uint256 toObjectId, (uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierTransferToInventory(smartObjectId *big.Int, toObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierTransferToInventory(&_World.TransactOpts, smartObjectId, toObjectId, items)
}

// EvefrontierUnanchor is a paid mutator transaction binding the contract method 0x9abf3ec9.
//
// Solidity: function evefrontier__unanchor(uint256 smartObjectId) returns()
func (_World *WorldTransactor) EvefrontierUnanchor(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__unanchor", smartObjectId)
}

// EvefrontierUnanchor is a paid mutator transaction binding the contract method 0x9abf3ec9.
//
// Solidity: function evefrontier__unanchor(uint256 smartObjectId) returns()
func (_World *WorldSession) EvefrontierUnanchor(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUnanchor(&_World.TransactOpts, smartObjectId)
}

// EvefrontierUnanchor is a paid mutator transaction binding the contract method 0x9abf3ec9.
//
// Solidity: function evefrontier__unanchor(uint256 smartObjectId) returns()
func (_World *WorldTransactorSession) EvefrontierUnanchor(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUnanchor(&_World.TransactOpts, smartObjectId)
}

// EvefrontierUnlinkGates is a paid mutator transaction binding the contract method 0x1a6f6bb1.
//
// Solidity: function evefrontier__unlinkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldTransactor) EvefrontierUnlinkGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__unlinkGates", sourceGateId, destinationGateId)
}

// EvefrontierUnlinkGates is a paid mutator transaction binding the contract method 0x1a6f6bb1.
//
// Solidity: function evefrontier__unlinkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldSession) EvefrontierUnlinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUnlinkGates(&_World.TransactOpts, sourceGateId, destinationGateId)
}

// EvefrontierUnlinkGates is a paid mutator transaction binding the contract method 0x1a6f6bb1.
//
// Solidity: function evefrontier__unlinkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldTransactorSession) EvefrontierUnlinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUnlinkGates(&_World.TransactOpts, sourceGateId, destinationGateId)
}

// EvefrontierUpdateAssemblyType is a paid mutator transaction binding the contract method 0x5921cdd6.
//
// Solidity: function evefrontier__updateAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_World *WorldTransactor) EvefrontierUpdateAssemblyType(opts *bind.TransactOpts, smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__updateAssemblyType", smartObjectId, assemblyType)
}

// EvefrontierUpdateAssemblyType is a paid mutator transaction binding the contract method 0x5921cdd6.
//
// Solidity: function evefrontier__updateAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_World *WorldSession) EvefrontierUpdateAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUpdateAssemblyType(&_World.TransactOpts, smartObjectId, assemblyType)
}

// EvefrontierUpdateAssemblyType is a paid mutator transaction binding the contract method 0x5921cdd6.
//
// Solidity: function evefrontier__updateAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_World *WorldTransactorSession) EvefrontierUpdateAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUpdateAssemblyType(&_World.TransactOpts, smartObjectId, assemblyType)
}

// EvefrontierUpdateEnergyHistory is a paid mutator transaction binding the contract method 0xf18500a1.
//
// Solidity: function evefrontier__updateEnergyHistory(uint256 networkNodeId) returns()
func (_World *WorldTransactor) EvefrontierUpdateEnergyHistory(opts *bind.TransactOpts, networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__updateEnergyHistory", networkNodeId)
}

// EvefrontierUpdateEnergyHistory is a paid mutator transaction binding the contract method 0xf18500a1.
//
// Solidity: function evefrontier__updateEnergyHistory(uint256 networkNodeId) returns()
func (_World *WorldSession) EvefrontierUpdateEnergyHistory(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUpdateEnergyHistory(&_World.TransactOpts, networkNodeId)
}

// EvefrontierUpdateEnergyHistory is a paid mutator transaction binding the contract method 0xf18500a1.
//
// Solidity: function evefrontier__updateEnergyHistory(uint256 networkNodeId) returns()
func (_World *WorldTransactorSession) EvefrontierUpdateEnergyHistory(networkNodeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUpdateEnergyHistory(&_World.TransactOpts, networkNodeId)
}

// EvefrontierUpdateFuel is a paid mutator transaction binding the contract method 0x75679e2e.
//
// Solidity: function evefrontier__updateFuel(uint256 smartObjectId) returns()
func (_World *WorldTransactor) EvefrontierUpdateFuel(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__updateFuel", smartObjectId)
}

// EvefrontierUpdateFuel is a paid mutator transaction binding the contract method 0x75679e2e.
//
// Solidity: function evefrontier__updateFuel(uint256 smartObjectId) returns()
func (_World *WorldSession) EvefrontierUpdateFuel(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUpdateFuel(&_World.TransactOpts, smartObjectId)
}

// EvefrontierUpdateFuel is a paid mutator transaction binding the contract method 0x75679e2e.
//
// Solidity: function evefrontier__updateFuel(uint256 smartObjectId) returns()
func (_World *WorldTransactorSession) EvefrontierUpdateFuel(smartObjectId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUpdateFuel(&_World.TransactOpts, smartObjectId)
}

// EvefrontierUpdateTribeId is a paid mutator transaction binding the contract method 0x7402be76.
//
// Solidity: function evefrontier__updateTribeId(uint256 smartObjectId, uint256 tribeId) returns()
func (_World *WorldTransactor) EvefrontierUpdateTribeId(opts *bind.TransactOpts, smartObjectId *big.Int, tribeId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__updateTribeId", smartObjectId, tribeId)
}

// EvefrontierUpdateTribeId is a paid mutator transaction binding the contract method 0x7402be76.
//
// Solidity: function evefrontier__updateTribeId(uint256 smartObjectId, uint256 tribeId) returns()
func (_World *WorldSession) EvefrontierUpdateTribeId(smartObjectId *big.Int, tribeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUpdateTribeId(&_World.TransactOpts, smartObjectId, tribeId)
}

// EvefrontierUpdateTribeId is a paid mutator transaction binding the contract method 0x7402be76.
//
// Solidity: function evefrontier__updateTribeId(uint256 smartObjectId, uint256 tribeId) returns()
func (_World *WorldTransactorSession) EvefrontierUpdateTribeId(smartObjectId *big.Int, tribeId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierUpdateTribeId(&_World.TransactOpts, smartObjectId, tribeId)
}

// EvefrontierWithdrawEphemeral is a paid mutator transaction binding the contract method 0xe792638f.
//
// Solidity: function evefrontier__withdrawEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierWithdrawEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__withdrawEphemeral", smartObjectId, ephemeralOwner, items)
}

// EvefrontierWithdrawEphemeral is a paid mutator transaction binding the contract method 0xe792638f.
//
// Solidity: function evefrontier__withdrawEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierWithdrawEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierWithdrawEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierWithdrawEphemeral is a paid mutator transaction binding the contract method 0xe792638f.
//
// Solidity: function evefrontier__withdrawEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierWithdrawEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierWithdrawEphemeral(&_World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// EvefrontierWithdrawFuel is a paid mutator transaction binding the contract method 0xd7edb90a.
//
// Solidity: function evefrontier__withdrawFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_World *WorldTransactor) EvefrontierWithdrawFuel(opts *bind.TransactOpts, smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__withdrawFuel", smartObjectId, fuelAmount)
}

// EvefrontierWithdrawFuel is a paid mutator transaction binding the contract method 0xd7edb90a.
//
// Solidity: function evefrontier__withdrawFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_World *WorldSession) EvefrontierWithdrawFuel(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierWithdrawFuel(&_World.TransactOpts, smartObjectId, fuelAmount)
}

// EvefrontierWithdrawFuel is a paid mutator transaction binding the contract method 0xd7edb90a.
//
// Solidity: function evefrontier__withdrawFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_World *WorldTransactorSession) EvefrontierWithdrawFuel(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _World.Contract.EvefrontierWithdrawFuel(&_World.TransactOpts, smartObjectId, fuelAmount)
}

// EvefrontierWithdrawInventory is a paid mutator transaction binding the contract method 0xd8a2f475.
//
// Solidity: function evefrontier__withdrawInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EvefrontierWithdrawInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "evefrontier__withdrawInventory", smartObjectId, items)
}

// EvefrontierWithdrawInventory is a paid mutator transaction binding the contract method 0xd8a2f475.
//
// Solidity: function evefrontier__withdrawInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_World *WorldSession) EvefrontierWithdrawInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierWithdrawInventory(&_World.TransactOpts, smartObjectId, items)
}

// EvefrontierWithdrawInventory is a paid mutator transaction binding the contract method 0xd8a2f475.
//
// Solidity: function evefrontier__withdrawInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EvefrontierWithdrawInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _World.Contract.EvefrontierWithdrawInventory(&_World.TransactOpts, smartObjectId, items)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_World *WorldTransactor) GrantAccess(opts *bind.TransactOpts, resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "grantAccess", resourceId, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_World *WorldSession) GrantAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _World.Contract.GrantAccess(&_World.TransactOpts, resourceId, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_World *WorldTransactorSession) GrantAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _World.Contract.GrantAccess(&_World.TransactOpts, resourceId, grantee)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_World *WorldTransactor) Initialize(opts *bind.TransactOpts, initModule common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "initialize", initModule)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_World *WorldSession) Initialize(initModule common.Address) (*types.Transaction, error) {
	return _World.Contract.Initialize(&_World.TransactOpts, initModule)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_World *WorldTransactorSession) Initialize(initModule common.Address) (*types.Transaction, error) {
	return _World.Contract.Initialize(&_World.TransactOpts, initModule)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_World *WorldTransactor) InstallModule(opts *bind.TransactOpts, module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "installModule", module, encodedArgs)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_World *WorldSession) InstallModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _World.Contract.InstallModule(&_World.TransactOpts, module, encodedArgs)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_World *WorldTransactorSession) InstallModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _World.Contract.InstallModule(&_World.TransactOpts, module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_World *WorldTransactor) InstallRootModule(opts *bind.TransactOpts, module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "installRootModule", module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_World *WorldSession) InstallRootModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _World.Contract.InstallRootModule(&_World.TransactOpts, module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_World *WorldTransactorSession) InstallRootModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _World.Contract.InstallRootModule(&_World.TransactOpts, module, encodedArgs)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_World *WorldTransactor) PopFromDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "popFromDynamicField", tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_World *WorldSession) PopFromDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _World.Contract.PopFromDynamicField(&_World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_World *WorldTransactorSession) PopFromDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _World.Contract.PopFromDynamicField(&_World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_World *WorldTransactor) PushToDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "pushToDynamicField", tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_World *WorldSession) PushToDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _World.Contract.PushToDynamicField(&_World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_World *WorldTransactorSession) PushToDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _World.Contract.PushToDynamicField(&_World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_World *WorldTransactor) RegisterDelegation(opts *bind.TransactOpts, delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerDelegation", delegatee, delegationControlId, initCallData)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_World *WorldSession) RegisterDelegation(delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _World.Contract.RegisterDelegation(&_World.TransactOpts, delegatee, delegationControlId, initCallData)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_World *WorldTransactorSession) RegisterDelegation(delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _World.Contract.RegisterDelegation(&_World.TransactOpts, delegatee, delegationControlId, initCallData)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_World *WorldTransactor) RegisterFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerFunctionSelector", systemId, systemFunctionSignature)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_World *WorldSession) RegisterFunctionSelector(systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _World.Contract.RegisterFunctionSelector(&_World.TransactOpts, systemId, systemFunctionSignature)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_World *WorldTransactorSession) RegisterFunctionSelector(systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _World.Contract.RegisterFunctionSelector(&_World.TransactOpts, systemId, systemFunctionSignature)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_World *WorldTransactor) RegisterNamespace(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerNamespace", namespaceId)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_World *WorldSession) RegisterNamespace(namespaceId [32]byte) (*types.Transaction, error) {
	return _World.Contract.RegisterNamespace(&_World.TransactOpts, namespaceId)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_World *WorldTransactorSession) RegisterNamespace(namespaceId [32]byte) (*types.Transaction, error) {
	return _World.Contract.RegisterNamespace(&_World.TransactOpts, namespaceId)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_World *WorldTransactor) RegisterNamespaceDelegation(opts *bind.TransactOpts, namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerNamespaceDelegation", namespaceId, delegationControlId, initCallData)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_World *WorldSession) RegisterNamespaceDelegation(namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _World.Contract.RegisterNamespaceDelegation(&_World.TransactOpts, namespaceId, delegationControlId, initCallData)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_World *WorldTransactorSession) RegisterNamespaceDelegation(namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _World.Contract.RegisterNamespaceDelegation(&_World.TransactOpts, namespaceId, delegationControlId, initCallData)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x6548a90a.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_World *WorldTransactor) RegisterRootFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, worldFunctionSignature string, systemFunctionSignature string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerRootFunctionSelector", systemId, worldFunctionSignature, systemFunctionSignature)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x6548a90a.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_World *WorldSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSignature string) (*types.Transaction, error) {
	return _World.Contract.RegisterRootFunctionSelector(&_World.TransactOpts, systemId, worldFunctionSignature, systemFunctionSignature)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x6548a90a.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_World *WorldTransactorSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSignature string) (*types.Transaction, error) {
	return _World.Contract.RegisterRootFunctionSelector(&_World.TransactOpts, systemId, worldFunctionSignature, systemFunctionSignature)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_World *WorldTransactor) RegisterStoreHook(opts *bind.TransactOpts, tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerStoreHook", tableId, hookAddress, enabledHooksBitmap)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_World *WorldSession) RegisterStoreHook(tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _World.Contract.RegisterStoreHook(&_World.TransactOpts, tableId, hookAddress, enabledHooksBitmap)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_World *WorldTransactorSession) RegisterStoreHook(tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _World.Contract.RegisterStoreHook(&_World.TransactOpts, tableId, hookAddress, enabledHooksBitmap)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_World *WorldTransactor) RegisterSystem(opts *bind.TransactOpts, systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerSystem", systemId, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_World *WorldSession) RegisterSystem(systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _World.Contract.RegisterSystem(&_World.TransactOpts, systemId, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_World *WorldTransactorSession) RegisterSystem(systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _World.Contract.RegisterSystem(&_World.TransactOpts, systemId, system, publicAccess)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_World *WorldTransactor) RegisterSystemHook(opts *bind.TransactOpts, systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerSystemHook", systemId, hookAddress, enabledHooksBitmap)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_World *WorldSession) RegisterSystemHook(systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _World.Contract.RegisterSystemHook(&_World.TransactOpts, systemId, hookAddress, enabledHooksBitmap)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_World *WorldTransactorSession) RegisterSystemHook(systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _World.Contract.RegisterSystemHook(&_World.TransactOpts, systemId, hookAddress, enabledHooksBitmap)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_World *WorldTransactor) RegisterTable(opts *bind.TransactOpts, tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerTable", tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_World *WorldSession) RegisterTable(tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _World.Contract.RegisterTable(&_World.TransactOpts, tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_World *WorldTransactorSession) RegisterTable(tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _World.Contract.RegisterTable(&_World.TransactOpts, tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_World *WorldTransactor) RenounceOwnership(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "renounceOwnership", namespaceId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_World *WorldSession) RenounceOwnership(namespaceId [32]byte) (*types.Transaction, error) {
	return _World.Contract.RenounceOwnership(&_World.TransactOpts, namespaceId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_World *WorldTransactorSession) RenounceOwnership(namespaceId [32]byte) (*types.Transaction, error) {
	return _World.Contract.RenounceOwnership(&_World.TransactOpts, namespaceId)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_World *WorldTransactor) RevokeAccess(opts *bind.TransactOpts, resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "revokeAccess", resourceId, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_World *WorldSession) RevokeAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _World.Contract.RevokeAccess(&_World.TransactOpts, resourceId, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_World *WorldTransactorSession) RevokeAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _World.Contract.RevokeAccess(&_World.TransactOpts, resourceId, grantee)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_World *WorldTransactor) SetDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "setDynamicField", tableId, keyTuple, dynamicFieldIndex, data)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_World *WorldSession) SetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _World.Contract.SetDynamicField(&_World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, data)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_World *WorldTransactorSession) SetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _World.Contract.SetDynamicField(&_World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_World *WorldTransactor) SetField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "setField", tableId, keyTuple, fieldIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_World *WorldSession) SetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _World.Contract.SetField(&_World.TransactOpts, tableId, keyTuple, fieldIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_World *WorldTransactorSession) SetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _World.Contract.SetField(&_World.TransactOpts, tableId, keyTuple, fieldIndex, data)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_World *WorldTransactor) SetField0(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "setField0", tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_World *WorldSession) SetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _World.Contract.SetField0(&_World.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_World *WorldTransactorSession) SetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _World.Contract.SetField0(&_World.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_World *WorldTransactor) SetRecord(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "setRecord", tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_World *WorldSession) SetRecord(tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _World.Contract.SetRecord(&_World.TransactOpts, tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_World *WorldTransactorSession) SetRecord(tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _World.Contract.SetRecord(&_World.TransactOpts, tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_World *WorldTransactor) SetStaticField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "setStaticField", tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_World *WorldSession) SetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _World.Contract.SetStaticField(&_World.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_World *WorldTransactorSession) SetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _World.Contract.SetStaticField(&_World.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_World *WorldTransactor) SpliceDynamicData(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "spliceDynamicData", tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_World *WorldSession) SpliceDynamicData(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _World.Contract.SpliceDynamicData(&_World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_World *WorldTransactorSession) SpliceDynamicData(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _World.Contract.SpliceDynamicData(&_World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_World *WorldTransactor) SpliceStaticData(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "spliceStaticData", tableId, keyTuple, start, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_World *WorldSession) SpliceStaticData(tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _World.Contract.SpliceStaticData(&_World.TransactOpts, tableId, keyTuple, start, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_World *WorldTransactorSession) SpliceStaticData(tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _World.Contract.SpliceStaticData(&_World.TransactOpts, tableId, keyTuple, start, data)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_World *WorldTransactor) TransferBalanceToAddress(opts *bind.TransactOpts, fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "transferBalanceToAddress", fromNamespaceId, toAddress, amount)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_World *WorldSession) TransferBalanceToAddress(fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _World.Contract.TransferBalanceToAddress(&_World.TransactOpts, fromNamespaceId, toAddress, amount)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_World *WorldTransactorSession) TransferBalanceToAddress(fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _World.Contract.TransferBalanceToAddress(&_World.TransactOpts, fromNamespaceId, toAddress, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_World *WorldTransactor) TransferBalanceToNamespace(opts *bind.TransactOpts, fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "transferBalanceToNamespace", fromNamespaceId, toNamespaceId, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_World *WorldSession) TransferBalanceToNamespace(fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _World.Contract.TransferBalanceToNamespace(&_World.TransactOpts, fromNamespaceId, toNamespaceId, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_World *WorldTransactorSession) TransferBalanceToNamespace(fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _World.Contract.TransferBalanceToNamespace(&_World.TransactOpts, fromNamespaceId, toNamespaceId, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_World *WorldTransactor) TransferOwnership(opts *bind.TransactOpts, namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "transferOwnership", namespaceId, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_World *WorldSession) TransferOwnership(namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _World.Contract.TransferOwnership(&_World.TransactOpts, namespaceId, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_World *WorldTransactorSession) TransferOwnership(namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _World.Contract.TransferOwnership(&_World.TransactOpts, namespaceId, newOwner)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_World *WorldTransactor) UnregisterDelegation(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "unregisterDelegation", delegatee)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_World *WorldSession) UnregisterDelegation(delegatee common.Address) (*types.Transaction, error) {
	return _World.Contract.UnregisterDelegation(&_World.TransactOpts, delegatee)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_World *WorldTransactorSession) UnregisterDelegation(delegatee common.Address) (*types.Transaction, error) {
	return _World.Contract.UnregisterDelegation(&_World.TransactOpts, delegatee)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_World *WorldTransactor) UnregisterNamespaceDelegation(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "unregisterNamespaceDelegation", namespaceId)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_World *WorldSession) UnregisterNamespaceDelegation(namespaceId [32]byte) (*types.Transaction, error) {
	return _World.Contract.UnregisterNamespaceDelegation(&_World.TransactOpts, namespaceId)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_World *WorldTransactorSession) UnregisterNamespaceDelegation(namespaceId [32]byte) (*types.Transaction, error) {
	return _World.Contract.UnregisterNamespaceDelegation(&_World.TransactOpts, namespaceId)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_World *WorldTransactor) UnregisterStoreHook(opts *bind.TransactOpts, tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "unregisterStoreHook", tableId, hookAddress)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_World *WorldSession) UnregisterStoreHook(tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _World.Contract.UnregisterStoreHook(&_World.TransactOpts, tableId, hookAddress)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_World *WorldTransactorSession) UnregisterStoreHook(tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _World.Contract.UnregisterStoreHook(&_World.TransactOpts, tableId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_World *WorldTransactor) UnregisterSystemHook(opts *bind.TransactOpts, systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "unregisterSystemHook", systemId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_World *WorldSession) UnregisterSystemHook(systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _World.Contract.UnregisterSystemHook(&_World.TransactOpts, systemId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_World *WorldTransactorSession) UnregisterSystemHook(systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _World.Contract.UnregisterSystemHook(&_World.TransactOpts, systemId, hookAddress)
}

// WorldHelloStoreIterator is returned from FilterHelloStore and is used to iterate over the raw logs and unpacked data for HelloStore events raised by the World contract.
type WorldHelloStoreIterator struct {
	Event *WorldHelloStore // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorldHelloStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldHelloStore)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorldHelloStore)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorldHelloStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldHelloStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldHelloStore represents a HelloStore event raised by the World contract.
type WorldHelloStore struct {
	StoreVersion [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHelloStore is a free log retrieval operation binding the contract event 0xc7f5fdc8526b76f54916701bc910876243ffff2a40b0bb8d59eea8151c52c005.
//
// Solidity: event HelloStore(bytes32 indexed storeVersion)
func (_World *WorldFilterer) FilterHelloStore(opts *bind.FilterOpts, storeVersion [][32]byte) (*WorldHelloStoreIterator, error) {

	var storeVersionRule []interface{}
	for _, storeVersionItem := range storeVersion {
		storeVersionRule = append(storeVersionRule, storeVersionItem)
	}

	logs, sub, err := _World.contract.FilterLogs(opts, "HelloStore", storeVersionRule)
	if err != nil {
		return nil, err
	}
	return &WorldHelloStoreIterator{contract: _World.contract, event: "HelloStore", logs: logs, sub: sub}, nil
}

// WatchHelloStore is a free log subscription operation binding the contract event 0xc7f5fdc8526b76f54916701bc910876243ffff2a40b0bb8d59eea8151c52c005.
//
// Solidity: event HelloStore(bytes32 indexed storeVersion)
func (_World *WorldFilterer) WatchHelloStore(opts *bind.WatchOpts, sink chan<- *WorldHelloStore, storeVersion [][32]byte) (event.Subscription, error) {

	var storeVersionRule []interface{}
	for _, storeVersionItem := range storeVersion {
		storeVersionRule = append(storeVersionRule, storeVersionItem)
	}

	logs, sub, err := _World.contract.WatchLogs(opts, "HelloStore", storeVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldHelloStore)
				if err := _World.contract.UnpackLog(event, "HelloStore", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHelloStore is a log parse operation binding the contract event 0xc7f5fdc8526b76f54916701bc910876243ffff2a40b0bb8d59eea8151c52c005.
//
// Solidity: event HelloStore(bytes32 indexed storeVersion)
func (_World *WorldFilterer) ParseHelloStore(log types.Log) (*WorldHelloStore, error) {
	event := new(WorldHelloStore)
	if err := _World.contract.UnpackLog(event, "HelloStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldHelloWorldIterator is returned from FilterHelloWorld and is used to iterate over the raw logs and unpacked data for HelloWorld events raised by the World contract.
type WorldHelloWorldIterator struct {
	Event *WorldHelloWorld // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorldHelloWorldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldHelloWorld)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorldHelloWorld)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorldHelloWorldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldHelloWorldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldHelloWorld represents a HelloWorld event raised by the World contract.
type WorldHelloWorld struct {
	WorldVersion [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHelloWorld is a free log retrieval operation binding the contract event 0x7f8f36afe3fb61c459c1a54a60b8a477eab02cc58e49f547561a40906239cb82.
//
// Solidity: event HelloWorld(bytes32 indexed worldVersion)
func (_World *WorldFilterer) FilterHelloWorld(opts *bind.FilterOpts, worldVersion [][32]byte) (*WorldHelloWorldIterator, error) {

	var worldVersionRule []interface{}
	for _, worldVersionItem := range worldVersion {
		worldVersionRule = append(worldVersionRule, worldVersionItem)
	}

	logs, sub, err := _World.contract.FilterLogs(opts, "HelloWorld", worldVersionRule)
	if err != nil {
		return nil, err
	}
	return &WorldHelloWorldIterator{contract: _World.contract, event: "HelloWorld", logs: logs, sub: sub}, nil
}

// WatchHelloWorld is a free log subscription operation binding the contract event 0x7f8f36afe3fb61c459c1a54a60b8a477eab02cc58e49f547561a40906239cb82.
//
// Solidity: event HelloWorld(bytes32 indexed worldVersion)
func (_World *WorldFilterer) WatchHelloWorld(opts *bind.WatchOpts, sink chan<- *WorldHelloWorld, worldVersion [][32]byte) (event.Subscription, error) {

	var worldVersionRule []interface{}
	for _, worldVersionItem := range worldVersion {
		worldVersionRule = append(worldVersionRule, worldVersionItem)
	}

	logs, sub, err := _World.contract.WatchLogs(opts, "HelloWorld", worldVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldHelloWorld)
				if err := _World.contract.UnpackLog(event, "HelloWorld", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHelloWorld is a log parse operation binding the contract event 0x7f8f36afe3fb61c459c1a54a60b8a477eab02cc58e49f547561a40906239cb82.
//
// Solidity: event HelloWorld(bytes32 indexed worldVersion)
func (_World *WorldFilterer) ParseHelloWorld(log types.Log) (*WorldHelloWorld, error) {
	event := new(WorldHelloWorld)
	if err := _World.contract.UnpackLog(event, "HelloWorld", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldStoreDeleteRecordIterator is returned from FilterStoreDeleteRecord and is used to iterate over the raw logs and unpacked data for StoreDeleteRecord events raised by the World contract.
type WorldStoreDeleteRecordIterator struct {
	Event *WorldStoreDeleteRecord // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorldStoreDeleteRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldStoreDeleteRecord)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorldStoreDeleteRecord)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorldStoreDeleteRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldStoreDeleteRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldStoreDeleteRecord represents a StoreDeleteRecord event raised by the World contract.
type WorldStoreDeleteRecord struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreDeleteRecord is a free log retrieval operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_World *WorldFilterer) FilterStoreDeleteRecord(opts *bind.FilterOpts, tableId [][32]byte) (*WorldStoreDeleteRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _World.contract.FilterLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &WorldStoreDeleteRecordIterator{contract: _World.contract, event: "Store_DeleteRecord", logs: logs, sub: sub}, nil
}

// WatchStoreDeleteRecord is a free log subscription operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_World *WorldFilterer) WatchStoreDeleteRecord(opts *bind.WatchOpts, sink chan<- *WorldStoreDeleteRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _World.contract.WatchLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldStoreDeleteRecord)
				if err := _World.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStoreDeleteRecord is a log parse operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_World *WorldFilterer) ParseStoreDeleteRecord(log types.Log) (*WorldStoreDeleteRecord, error) {
	event := new(WorldStoreDeleteRecord)
	if err := _World.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldStoreSetRecordIterator is returned from FilterStoreSetRecord and is used to iterate over the raw logs and unpacked data for StoreSetRecord events raised by the World contract.
type WorldStoreSetRecordIterator struct {
	Event *WorldStoreSetRecord // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorldStoreSetRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldStoreSetRecord)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorldStoreSetRecord)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorldStoreSetRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldStoreSetRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldStoreSetRecord represents a StoreSetRecord event raised by the World contract.
type WorldStoreSetRecord struct {
	TableId        [32]byte
	KeyTuple       [][32]byte
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStoreSetRecord is a free log retrieval operation binding the contract event 0x8dbb3a9672eebfd3773e72dd9c102393436816d832c7ba9e1e1ac8fcadcac7a9.
//
// Solidity: event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_World *WorldFilterer) FilterStoreSetRecord(opts *bind.FilterOpts, tableId [][32]byte) (*WorldStoreSetRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _World.contract.FilterLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &WorldStoreSetRecordIterator{contract: _World.contract, event: "Store_SetRecord", logs: logs, sub: sub}, nil
}

// WatchStoreSetRecord is a free log subscription operation binding the contract event 0x8dbb3a9672eebfd3773e72dd9c102393436816d832c7ba9e1e1ac8fcadcac7a9.
//
// Solidity: event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_World *WorldFilterer) WatchStoreSetRecord(opts *bind.WatchOpts, sink chan<- *WorldStoreSetRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _World.contract.WatchLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldStoreSetRecord)
				if err := _World.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStoreSetRecord is a log parse operation binding the contract event 0x8dbb3a9672eebfd3773e72dd9c102393436816d832c7ba9e1e1ac8fcadcac7a9.
//
// Solidity: event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_World *WorldFilterer) ParseStoreSetRecord(log types.Log) (*WorldStoreSetRecord, error) {
	event := new(WorldStoreSetRecord)
	if err := _World.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldStoreSpliceDynamicDataIterator is returned from FilterStoreSpliceDynamicData and is used to iterate over the raw logs and unpacked data for StoreSpliceDynamicData events raised by the World contract.
type WorldStoreSpliceDynamicDataIterator struct {
	Event *WorldStoreSpliceDynamicData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorldStoreSpliceDynamicDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldStoreSpliceDynamicData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorldStoreSpliceDynamicData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorldStoreSpliceDynamicDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldStoreSpliceDynamicDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldStoreSpliceDynamicData represents a StoreSpliceDynamicData event raised by the World contract.
type WorldStoreSpliceDynamicData struct {
	TableId           [32]byte
	KeyTuple          [][32]byte
	DynamicFieldIndex uint8
	Start             *big.Int
	DeleteCount       *big.Int
	EncodedLengths    [32]byte
	Data              []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStoreSpliceDynamicData is a free log retrieval operation binding the contract event 0xfe158a7adba34e256807c8a149028d3162918713c3838afc643ce9f96716ebfd.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
func (_World *WorldFilterer) FilterStoreSpliceDynamicData(opts *bind.FilterOpts, tableId [][32]byte) (*WorldStoreSpliceDynamicDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _World.contract.FilterLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &WorldStoreSpliceDynamicDataIterator{contract: _World.contract, event: "Store_SpliceDynamicData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceDynamicData is a free log subscription operation binding the contract event 0xfe158a7adba34e256807c8a149028d3162918713c3838afc643ce9f96716ebfd.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
func (_World *WorldFilterer) WatchStoreSpliceDynamicData(opts *bind.WatchOpts, sink chan<- *WorldStoreSpliceDynamicData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _World.contract.WatchLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldStoreSpliceDynamicData)
				if err := _World.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStoreSpliceDynamicData is a log parse operation binding the contract event 0xfe158a7adba34e256807c8a149028d3162918713c3838afc643ce9f96716ebfd.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
func (_World *WorldFilterer) ParseStoreSpliceDynamicData(log types.Log) (*WorldStoreSpliceDynamicData, error) {
	event := new(WorldStoreSpliceDynamicData)
	if err := _World.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldStoreSpliceStaticDataIterator is returned from FilterStoreSpliceStaticData and is used to iterate over the raw logs and unpacked data for StoreSpliceStaticData events raised by the World contract.
type WorldStoreSpliceStaticDataIterator struct {
	Event *WorldStoreSpliceStaticData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorldStoreSpliceStaticDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldStoreSpliceStaticData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorldStoreSpliceStaticData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorldStoreSpliceStaticDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldStoreSpliceStaticDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldStoreSpliceStaticData represents a StoreSpliceStaticData event raised by the World contract.
type WorldStoreSpliceStaticData struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Start    *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreSpliceStaticData is a free log retrieval operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_World *WorldFilterer) FilterStoreSpliceStaticData(opts *bind.FilterOpts, tableId [][32]byte) (*WorldStoreSpliceStaticDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _World.contract.FilterLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &WorldStoreSpliceStaticDataIterator{contract: _World.contract, event: "Store_SpliceStaticData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceStaticData is a free log subscription operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_World *WorldFilterer) WatchStoreSpliceStaticData(opts *bind.WatchOpts, sink chan<- *WorldStoreSpliceStaticData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _World.contract.WatchLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldStoreSpliceStaticData)
				if err := _World.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStoreSpliceStaticData is a log parse operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_World *WorldFilterer) ParseStoreSpliceStaticData(log types.Log) (*WorldStoreSpliceStaticData, error) {
	event := new(WorldStoreSpliceStaticData)
	if err := _World.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
