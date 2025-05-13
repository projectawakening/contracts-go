// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC2771World

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

// ERC2771WorldMetaData contains all meta data concerning the ERC2771World contract.
var ERC2771WorldMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"batchCall\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallData[]\",\"components\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchCallFrom\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallFromData[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"callFrom\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"creator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_handleNodeOffline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminSupportOrDirectOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"adminSupportOrDirectOwnerGates\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggression\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structAggressionParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"aggressor\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"victim\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"anchor\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"areGatesOnline\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assignItemToInventory\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assignOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bringOffline\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bringOnline\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canCrossTransferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canJump\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canTransferFromEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canTransferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canTransferToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configureDeployableAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureEntityRecordAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureEphemeralInteractAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureEphemeralInventoryAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureFuelAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureFuelEfficiency\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelEntityParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fuelEfficiency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureFuelParameters\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelParams\",\"type\":\"tuple\",\"internalType\":\"structFuelParams\",\"components\":[{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelBurnRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureGate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureInventoryAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureInventoryInteractAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureKillMailAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureLocationAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureNetworkNodeAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureOwnershipAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartAssemblyAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartCharacterAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartGateAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartStorageUnitAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartTurretAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureTurret\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"connectAssembly\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchor\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorGate\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"maxDistance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorNetworkNode\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelParams\",\"type\":\"tuple\",\"internalType\":\"structFuelParams\",\"components\":[{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelBurnRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"maxEnergyCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentProduction\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorStorageUnit\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorTurret\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndDepositEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structCreateInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndDepositInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structCreateInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAssembly\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createCharacter\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tribeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"entityRecordMetadata\",\"type\":\"tuple\",\"internalType\":\"structEntityMetadataParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createDeployable\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createMetadata\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordMetadata\",\"type\":\"tuple\",\"internalType\":\"structEntityMetadataParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRecord\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossTransferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromEphemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toEphemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"destroyDeployable\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentFuelConsumptionStatus\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"timeLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitsToConsume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualConsumptionRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEphemeralOwner\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEphemeralSmartObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getInventoryOwner\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNetworkNodeClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSmartCharacterClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSmartGateClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSmartStorageUnitClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSmartTurretClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inProximity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turretTarget\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAnyGateLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isClassScoped\",\"inputs\":[{\"name\":\"classId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEphemeralOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isGateLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerOfBothGates\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWithinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onAssemblyOffline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onAssemblyOnline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onNodeOffline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onlyAdminOrCallAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminOrClassScopedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminOrOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminOrOwnerSupported\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminOrScopeEnforcedCall\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminSupportedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminSupportedOwnerOrCall\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyCallAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyCallAccessOrDirectEphemeralOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyCallAccessWithScopeEnforced\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyClassScopedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyClassScopedOrCharAdminOrOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyDirectAdmin\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyDirectAdminOrCallAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyDirectEphemeralOwnerOrCall\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyDirectOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyEphemeralOwnerOrTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwnerOrEphemeralCrossTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwnerOrEphemeralTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwnerOrInventoryTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwnerWithAdminSupportAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlySmartAssemblyClassScopedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerNetworkNodeClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSmartCharacterClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSmartGateClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSmartStorageUnitClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSmartTurretClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeCharacter\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeItemFromInventory\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportKill\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killMailData\",\"type\":\"tuple\",\"internalType\":\"structKillMailData\",\"components\":[{\"name\":\"killerCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"victimCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lossType\",\"type\":\"uint8\",\"internalType\":\"enumKillMailLossType\"},{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"saveLocation\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAssemblyType\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCrossTransferToEphemeralAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDappURL\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDescription\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEphemeralCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFuelMaxCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFuelUnitVolume\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTransferFromEphemeralAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTransferToEphemeralAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTransferToInventoryAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBurn\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopBurn\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFromEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unanchor\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlinkGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAssemblyType\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTribeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tribeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldSlice\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLayout\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeySchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValueSchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initModule\",\"type\":\"address\",\"internalType\":\"contractIModule\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installRootModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"popFromDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"byteLengthToPop\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pushToDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dataToPush\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"systemFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespace\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerRootFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"worldFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"systemFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystem\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"contractSystem\"},{\"name\":\"publicAccess\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerTable\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"},{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"keyNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"fieldNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storeVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferBalanceToAddress\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferBalanceToNamespace\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"worldVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"HelloStore\",\"inputs\":[{\"name\":\"storeVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HelloWorld\",\"inputs\":[{\"name\":\"worldVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_DeleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SetRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"EncodedLengths\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Access_CannotTransferFromEphemeral\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrClassScoped\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrOwnerSupported\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupported\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupportedOrDirectOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupportedOrDirectOwnerGates\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupportedOwnerOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotClassScoped\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotClassScopedAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectAdminOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectEphemeralOwnerOrCanCrossTransferToEphemeral\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectOwnerOrCanTransferToEphemeral\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectOwnerOrCanTransferToInventory\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotEphemeralOwnerOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotEphemeralOwnerOrCallAccessWithEphemeralOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotOwnerWithAdminSupportAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Deployable_IncorrectState\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentState\",\"type\":\"uint8\",\"internalType\":\"enumState\"}]},{\"type\":\"error\",\"name\":\"Deployable_InvalidObjectOwner\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"smartObjectOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EncodedLengths_InvalidLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InsufficientCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"usedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidEphemeralOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidItemDepositQuantity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidItemObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidSmartObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidTenantId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_NonExistentEntityRecord\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FieldLayout_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"staticDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"computedStaticDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_StaticLengthDoesNotFitInAWord\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_StaticLengthIsNotZero\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_StaticLengthIsZero\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_TooManyDynamicFields\",\"inputs\":[{\"name\":\"numFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_TooManyFields\",\"inputs\":[{\"name\":\"numFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_BurnAlreadyStopped\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_BurnNotActive\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_ExceedsMaxCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalProjectedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InsufficientFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableFuel\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelAmount\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelBurnRate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelBurnRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelEfficiency\",\"inputs\":[{\"name\":\"fuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelEfficiency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelMaxCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelSmartObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelTypeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelUnitVolume\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_TypeMismatch\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentFuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newFuelSmartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_Ephemeral_InsufficientQuantity\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InsufficientQuantity\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InvalidInventory\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InvalidOperation\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InvalidQuantity\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_NonexistentItemRecord\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_NonexistentObject\",\"inputs\":[{\"name\":\"objectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_SingletonAlreadyAssigned\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentInventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_SingletonDirectlyOwned\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"directOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_ZeroQuantity\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InsufficientCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"usedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidItemDepositQuantity\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidItemObjectId\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidTenantId\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"Inventory_NonExistentEntityRecord\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"KillMail_AlreadyExists\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"KillMail_InvalidCharacterId\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Module_AlreadyInstalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Module_MissingDependency\",\"inputs\":[{\"name\":\"dependency\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Module_NonRootInstallNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Module_RootInstallNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NetworkNode_AlreadyExists\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_AssemblyAlreadyConnected\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_AssemblyNotConnected\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_DoesNotExist\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_InsufficientEnergy\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_NotOnline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Ownership_AlreadyOwned\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Ownership_InvalidAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Ownership_InvalidOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"invalidOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Ownership_InvalidSingleton\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Ownership_NonexistentObject\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Ownership_SingletonInInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Schema_InvalidLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Schema_StaticTypeAfterDynamicType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Slice_OutOfBounds\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_DoesNotExist\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_InvalidObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_InvalidTenantId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_InvalidTypeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_TypeCannotBeEmpty\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacterDoesNotExist\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_AlreadyCreated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_InvalidObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_InvalidTenantId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_InvalidTypeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateAlreadyLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateNotLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateNotOnline\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GatesNotOnline\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotWithtinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_SameSourceAndDestination\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_UndefinedClassId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Store_IndexOutOfBounds\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessedIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidBounds\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidFieldNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidKeyNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidSplice\",\"inputs\":[{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"fieldLength\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaDynamicLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaStaticLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_TableAlreadyExists\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_TableNotFound\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_AccessDenied\",\"inputs\":[{\"name\":\"resource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"World_CallbackNotAllowed\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_DelegationNotFound\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorAlreadyExists\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorNotFound\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"World_InterfaceNotSupported\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InvalidNamespace\",\"inputs\":[{\"name\":\"namespace\",\"type\":\"bytes14\",\"internalType\":\"bytes14\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceId\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceAlreadyExists\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceNotFound\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_SystemAlreadyExists\",\"inputs\":[{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_UnlimitedDelegationNotAllowed\",\"inputs\":[]}]",
}

// ERC2771WorldABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC2771WorldMetaData.ABI instead.
var ERC2771WorldABI = ERC2771WorldMetaData.ABI

// ERC2771World is an auto generated Go binding around an Ethereum contract.
type ERC2771World struct {
	ERC2771WorldCaller     // Read-only binding to the contract
	ERC2771WorldTransactor // Write-only binding to the contract
	ERC2771WorldFilterer   // Log filterer for contract events
}

// ERC2771WorldCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC2771WorldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771WorldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC2771WorldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771WorldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC2771WorldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771WorldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC2771WorldSession struct {
	Contract     *ERC2771World     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC2771WorldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC2771WorldCallerSession struct {
	Contract *ERC2771WorldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC2771WorldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC2771WorldTransactorSession struct {
	Contract     *ERC2771WorldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC2771WorldRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC2771WorldRaw struct {
	Contract *ERC2771World // Generic contract binding to access the raw methods on
}

// ERC2771WorldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC2771WorldCallerRaw struct {
	Contract *ERC2771WorldCaller // Generic read-only contract binding to access the raw methods on
}

// ERC2771WorldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC2771WorldTransactorRaw struct {
	Contract *ERC2771WorldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC2771World creates a new instance of ERC2771World, bound to a specific deployed contract.
func NewERC2771World(address common.Address, backend bind.ContractBackend) (*ERC2771World, error) {
	contract, err := bindERC2771World(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC2771World{ERC2771WorldCaller: ERC2771WorldCaller{contract: contract}, ERC2771WorldTransactor: ERC2771WorldTransactor{contract: contract}, ERC2771WorldFilterer: ERC2771WorldFilterer{contract: contract}}, nil
}

// NewERC2771WorldCaller creates a new read-only instance of ERC2771World, bound to a specific deployed contract.
func NewERC2771WorldCaller(address common.Address, caller bind.ContractCaller) (*ERC2771WorldCaller, error) {
	contract, err := bindERC2771World(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC2771WorldCaller{contract: contract}, nil
}

// NewERC2771WorldTransactor creates a new write-only instance of ERC2771World, bound to a specific deployed contract.
func NewERC2771WorldTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC2771WorldTransactor, error) {
	contract, err := bindERC2771World(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC2771WorldTransactor{contract: contract}, nil
}

// NewERC2771WorldFilterer creates a new log filterer instance of ERC2771World, bound to a specific deployed contract.
func NewERC2771WorldFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC2771WorldFilterer, error) {
	contract, err := bindERC2771World(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC2771WorldFilterer{contract: contract}, nil
}

// bindERC2771World binds a generic wrapper to an already deployed contract.
func bindERC2771World(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC2771WorldMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC2771World *ERC2771WorldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC2771World.Contract.ERC2771WorldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC2771World *ERC2771WorldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.Contract.ERC2771WorldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC2771World *ERC2771WorldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC2771World.Contract.ERC2771WorldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC2771World *ERC2771WorldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC2771World.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC2771World *ERC2771WorldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC2771World *ERC2771WorldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC2771World.Contract.contract.Transact(opts, method, params...)
}

// AdminSupportOrDirectOwner is a free data retrieval call binding the contract method 0xc6e0b1a1.
//
// Solidity: function adminSupportOrDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) AdminSupportOrDirectOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "adminSupportOrDirectOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// AdminSupportOrDirectOwner is a free data retrieval call binding the contract method 0xc6e0b1a1.
//
// Solidity: function adminSupportOrDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) AdminSupportOrDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.AdminSupportOrDirectOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// AdminSupportOrDirectOwner is a free data retrieval call binding the contract method 0xc6e0b1a1.
//
// Solidity: function adminSupportOrDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) AdminSupportOrDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.AdminSupportOrDirectOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// AdminSupportOrDirectOwnerGates is a free data retrieval call binding the contract method 0xa6f28a83.
//
// Solidity: function adminSupportOrDirectOwnerGates(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) AdminSupportOrDirectOwnerGates(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "adminSupportOrDirectOwnerGates", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// AdminSupportOrDirectOwnerGates is a free data retrieval call binding the contract method 0xa6f28a83.
//
// Solidity: function adminSupportOrDirectOwnerGates(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) AdminSupportOrDirectOwnerGates(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.AdminSupportOrDirectOwnerGates(&_ERC2771World.CallOpts, smartObjectId, data)
}

// AdminSupportOrDirectOwnerGates is a free data retrieval call binding the contract method 0xa6f28a83.
//
// Solidity: function adminSupportOrDirectOwnerGates(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) AdminSupportOrDirectOwnerGates(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.AdminSupportOrDirectOwnerGates(&_ERC2771World.CallOpts, smartObjectId, data)
}

// AreGatesOnline is a free data retrieval call binding the contract method 0x9b11b875.
//
// Solidity: function areGatesOnline(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) AreGatesOnline(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "areGatesOnline", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreGatesOnline is a free data retrieval call binding the contract method 0x9b11b875.
//
// Solidity: function areGatesOnline(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) AreGatesOnline(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771World.Contract.AreGatesOnline(&_ERC2771World.CallOpts, sourceGateId, destinationGateId)
}

// AreGatesOnline is a free data retrieval call binding the contract method 0x9b11b875.
//
// Solidity: function areGatesOnline(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) AreGatesOnline(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771World.Contract.AreGatesOnline(&_ERC2771World.CallOpts, sourceGateId, destinationGateId)
}

// CanCrossTransferToEphemeral is a free data retrieval call binding the contract method 0x95d62522.
//
// Solidity: function canCrossTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) CanCrossTransferToEphemeral(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "canCrossTransferToEphemeral", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCrossTransferToEphemeral is a free data retrieval call binding the contract method 0x95d62522.
//
// Solidity: function canCrossTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) CanCrossTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.CanCrossTransferToEphemeral(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// CanCrossTransferToEphemeral is a free data retrieval call binding the contract method 0x95d62522.
//
// Solidity: function canCrossTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) CanCrossTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.CanCrossTransferToEphemeral(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// CanTransferFromEphemeral is a free data retrieval call binding the contract method 0x609fb5b0.
//
// Solidity: function canTransferFromEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) CanTransferFromEphemeral(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "canTransferFromEphemeral", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransferFromEphemeral is a free data retrieval call binding the contract method 0x609fb5b0.
//
// Solidity: function canTransferFromEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) CanTransferFromEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.CanTransferFromEphemeral(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// CanTransferFromEphemeral is a free data retrieval call binding the contract method 0x609fb5b0.
//
// Solidity: function canTransferFromEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) CanTransferFromEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.CanTransferFromEphemeral(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// CanTransferToEphemeral is a free data retrieval call binding the contract method 0x88756560.
//
// Solidity: function canTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) CanTransferToEphemeral(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "canTransferToEphemeral", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransferToEphemeral is a free data retrieval call binding the contract method 0x88756560.
//
// Solidity: function canTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) CanTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.CanTransferToEphemeral(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// CanTransferToEphemeral is a free data retrieval call binding the contract method 0x88756560.
//
// Solidity: function canTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) CanTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.CanTransferToEphemeral(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// CanTransferToInventory is a free data retrieval call binding the contract method 0x0fcd8ed6.
//
// Solidity: function canTransferToInventory(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) CanTransferToInventory(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "canTransferToInventory", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransferToInventory is a free data retrieval call binding the contract method 0x0fcd8ed6.
//
// Solidity: function canTransferToInventory(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) CanTransferToInventory(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.CanTransferToInventory(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// CanTransferToInventory is a free data retrieval call binding the contract method 0x0fcd8ed6.
//
// Solidity: function canTransferToInventory(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) CanTransferToInventory(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.CanTransferToInventory(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_ERC2771World *ERC2771WorldCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_ERC2771World *ERC2771WorldSession) Creator() (common.Address, error) {
	return _ERC2771World.Contract.Creator(&_ERC2771World.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_ERC2771World *ERC2771WorldCallerSession) Creator() (common.Address, error) {
	return _ERC2771World.Contract.Creator(&_ERC2771World.CallOpts)
}

// GetCurrentFuelConsumptionStatus is a free data retrieval call binding the contract method 0x655d9035.
//
// Solidity: function getCurrentFuelConsumptionStatus(uint256 smartObjectId) view returns(uint256 timeLeft, uint256 unitsToConsume, uint256 actualConsumptionRateInSeconds, uint256 fuelAmount)
func (_ERC2771World *ERC2771WorldCaller) GetCurrentFuelConsumptionStatus(opts *bind.CallOpts, smartObjectId *big.Int) (struct {
	TimeLeft                       *big.Int
	UnitsToConsume                 *big.Int
	ActualConsumptionRateInSeconds *big.Int
	FuelAmount                     *big.Int
}, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getCurrentFuelConsumptionStatus", smartObjectId)

	outstruct := new(struct {
		TimeLeft                       *big.Int
		UnitsToConsume                 *big.Int
		ActualConsumptionRateInSeconds *big.Int
		FuelAmount                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TimeLeft = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnitsToConsume = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ActualConsumptionRateInSeconds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FuelAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentFuelConsumptionStatus is a free data retrieval call binding the contract method 0x655d9035.
//
// Solidity: function getCurrentFuelConsumptionStatus(uint256 smartObjectId) view returns(uint256 timeLeft, uint256 unitsToConsume, uint256 actualConsumptionRateInSeconds, uint256 fuelAmount)
func (_ERC2771World *ERC2771WorldSession) GetCurrentFuelConsumptionStatus(smartObjectId *big.Int) (struct {
	TimeLeft                       *big.Int
	UnitsToConsume                 *big.Int
	ActualConsumptionRateInSeconds *big.Int
	FuelAmount                     *big.Int
}, error) {
	return _ERC2771World.Contract.GetCurrentFuelConsumptionStatus(&_ERC2771World.CallOpts, smartObjectId)
}

// GetCurrentFuelConsumptionStatus is a free data retrieval call binding the contract method 0x655d9035.
//
// Solidity: function getCurrentFuelConsumptionStatus(uint256 smartObjectId) view returns(uint256 timeLeft, uint256 unitsToConsume, uint256 actualConsumptionRateInSeconds, uint256 fuelAmount)
func (_ERC2771World *ERC2771WorldCallerSession) GetCurrentFuelConsumptionStatus(smartObjectId *big.Int) (struct {
	TimeLeft                       *big.Int
	UnitsToConsume                 *big.Int
	ActualConsumptionRateInSeconds *big.Int
	FuelAmount                     *big.Int
}, error) {
	return _ERC2771World.Contract.GetCurrentFuelConsumptionStatus(&_ERC2771World.CallOpts, smartObjectId)
}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_ERC2771World *ERC2771WorldCaller) GetDynamicField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getDynamicField", tableId, keyTuple, dynamicFieldIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_ERC2771World *ERC2771WorldSession) GetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	return _ERC2771World.Contract.GetDynamicField(&_ERC2771World.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_ERC2771World *ERC2771WorldCallerSession) GetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	return _ERC2771World.Contract.GetDynamicField(&_ERC2771World.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_ERC2771World *ERC2771WorldCaller) GetDynamicFieldLength(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getDynamicFieldLength", tableId, keyTuple, dynamicFieldIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_ERC2771World *ERC2771WorldSession) GetDynamicFieldLength(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	return _ERC2771World.Contract.GetDynamicFieldLength(&_ERC2771World.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_ERC2771World *ERC2771WorldCallerSession) GetDynamicFieldLength(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	return _ERC2771World.Contract.GetDynamicFieldLength(&_ERC2771World.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_ERC2771World *ERC2771WorldCaller) GetDynamicFieldSlice(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getDynamicFieldSlice", tableId, keyTuple, dynamicFieldIndex, start, end)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_ERC2771World *ERC2771WorldSession) GetDynamicFieldSlice(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	return _ERC2771World.Contract.GetDynamicFieldSlice(&_ERC2771World.CallOpts, tableId, keyTuple, dynamicFieldIndex, start, end)
}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_ERC2771World *ERC2771WorldCallerSession) GetDynamicFieldSlice(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	return _ERC2771World.Contract.GetDynamicFieldSlice(&_ERC2771World.CallOpts, tableId, keyTuple, dynamicFieldIndex, start, end)
}

// GetEphemeralOwner is a free data retrieval call binding the contract method 0x2a53720d.
//
// Solidity: function getEphemeralOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771World *ERC2771WorldCaller) GetEphemeralOwner(opts *bind.CallOpts, inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getEphemeralOwner", inventoryObjectId, itemObjectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEphemeralOwner is a free data retrieval call binding the contract method 0x2a53720d.
//
// Solidity: function getEphemeralOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771World *ERC2771WorldSession) GetEphemeralOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _ERC2771World.Contract.GetEphemeralOwner(&_ERC2771World.CallOpts, inventoryObjectId, itemObjectId)
}

// GetEphemeralOwner is a free data retrieval call binding the contract method 0x2a53720d.
//
// Solidity: function getEphemeralOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771World *ERC2771WorldCallerSession) GetEphemeralOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _ERC2771World.Contract.GetEphemeralOwner(&_ERC2771World.CallOpts, inventoryObjectId, itemObjectId)
}

// GetEphemeralSmartObjectId is a free data retrieval call binding the contract method 0x83619510.
//
// Solidity: function getEphemeralSmartObjectId(uint256 smartObjectId, address ephemeralOwner) pure returns(uint256)
func (_ERC2771World *ERC2771WorldCaller) GetEphemeralSmartObjectId(opts *bind.CallOpts, smartObjectId *big.Int, ephemeralOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getEphemeralSmartObjectId", smartObjectId, ephemeralOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEphemeralSmartObjectId is a free data retrieval call binding the contract method 0x83619510.
//
// Solidity: function getEphemeralSmartObjectId(uint256 smartObjectId, address ephemeralOwner) pure returns(uint256)
func (_ERC2771World *ERC2771WorldSession) GetEphemeralSmartObjectId(smartObjectId *big.Int, ephemeralOwner common.Address) (*big.Int, error) {
	return _ERC2771World.Contract.GetEphemeralSmartObjectId(&_ERC2771World.CallOpts, smartObjectId, ephemeralOwner)
}

// GetEphemeralSmartObjectId is a free data retrieval call binding the contract method 0x83619510.
//
// Solidity: function getEphemeralSmartObjectId(uint256 smartObjectId, address ephemeralOwner) pure returns(uint256)
func (_ERC2771World *ERC2771WorldCallerSession) GetEphemeralSmartObjectId(smartObjectId *big.Int, ephemeralOwner common.Address) (*big.Int, error) {
	return _ERC2771World.Contract.GetEphemeralSmartObjectId(&_ERC2771World.CallOpts, smartObjectId, ephemeralOwner)
}

// GetField is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_ERC2771World *ERC2771WorldCaller) GetField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getField", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetField is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_ERC2771World *ERC2771WorldSession) GetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	return _ERC2771World.Contract.GetField(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetField is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_ERC2771World *ERC2771WorldCallerSession) GetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	return _ERC2771World.Contract.GetField(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetField0 is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_ERC2771World *ERC2771WorldCaller) GetField0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getField0", tableId, keyTuple, fieldIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetField0 is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_ERC2771World *ERC2771WorldSession) GetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	return _ERC2771World.Contract.GetField0(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetField0 is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_ERC2771World *ERC2771WorldCallerSession) GetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	return _ERC2771World.Contract.GetField0(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_ERC2771World *ERC2771WorldCaller) GetFieldLayout(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getFieldLayout", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_ERC2771World *ERC2771WorldSession) GetFieldLayout(tableId [32]byte) ([32]byte, error) {
	return _ERC2771World.Contract.GetFieldLayout(&_ERC2771World.CallOpts, tableId)
}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_ERC2771World *ERC2771WorldCallerSession) GetFieldLayout(tableId [32]byte) ([32]byte, error) {
	return _ERC2771World.Contract.GetFieldLayout(&_ERC2771World.CallOpts, tableId)
}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_ERC2771World *ERC2771WorldCaller) GetFieldLength(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getFieldLength", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_ERC2771World *ERC2771WorldSession) GetFieldLength(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	return _ERC2771World.Contract.GetFieldLength(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_ERC2771World *ERC2771WorldCallerSession) GetFieldLength(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	return _ERC2771World.Contract.GetFieldLength(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetFieldLength0 is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_ERC2771World *ERC2771WorldCaller) GetFieldLength0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getFieldLength0", tableId, keyTuple, fieldIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldLength0 is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_ERC2771World *ERC2771WorldSession) GetFieldLength0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	return _ERC2771World.Contract.GetFieldLength0(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetFieldLength0 is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_ERC2771World *ERC2771WorldCallerSession) GetFieldLength0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	return _ERC2771World.Contract.GetFieldLength0(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetInventoryOwner is a free data retrieval call binding the contract method 0x1904f95b.
//
// Solidity: function getInventoryOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771World *ERC2771WorldCaller) GetInventoryOwner(opts *bind.CallOpts, inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getInventoryOwner", inventoryObjectId, itemObjectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInventoryOwner is a free data retrieval call binding the contract method 0x1904f95b.
//
// Solidity: function getInventoryOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771World *ERC2771WorldSession) GetInventoryOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _ERC2771World.Contract.GetInventoryOwner(&_ERC2771World.CallOpts, inventoryObjectId, itemObjectId)
}

// GetInventoryOwner is a free data retrieval call binding the contract method 0x1904f95b.
//
// Solidity: function getInventoryOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771World *ERC2771WorldCallerSession) GetInventoryOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _ERC2771World.Contract.GetInventoryOwner(&_ERC2771World.CallOpts, inventoryObjectId, itemObjectId)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_ERC2771World *ERC2771WorldCaller) GetKeySchema(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getKeySchema", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_ERC2771World *ERC2771WorldSession) GetKeySchema(tableId [32]byte) ([32]byte, error) {
	return _ERC2771World.Contract.GetKeySchema(&_ERC2771World.CallOpts, tableId)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_ERC2771World *ERC2771WorldCallerSession) GetKeySchema(tableId [32]byte) ([32]byte, error) {
	return _ERC2771World.Contract.GetKeySchema(&_ERC2771World.CallOpts, tableId)
}

// GetNetworkNodeClassId is a free data retrieval call binding the contract method 0x520b91b9.
//
// Solidity: function getNetworkNodeClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCaller) GetNetworkNodeClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getNetworkNodeClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNetworkNodeClassId is a free data retrieval call binding the contract method 0x520b91b9.
//
// Solidity: function getNetworkNodeClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldSession) GetNetworkNodeClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetNetworkNodeClassId(&_ERC2771World.CallOpts)
}

// GetNetworkNodeClassId is a free data retrieval call binding the contract method 0x520b91b9.
//
// Solidity: function getNetworkNodeClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCallerSession) GetNetworkNodeClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetNetworkNodeClassId(&_ERC2771World.CallOpts)
}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771World *ERC2771WorldCaller) GetRecord(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getRecord", tableId, keyTuple, fieldLayout)

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
func (_ERC2771World *ERC2771WorldSession) GetRecord(tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _ERC2771World.Contract.GetRecord(&_ERC2771World.CallOpts, tableId, keyTuple, fieldLayout)
}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771World *ERC2771WorldCallerSession) GetRecord(tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _ERC2771World.Contract.GetRecord(&_ERC2771World.CallOpts, tableId, keyTuple, fieldLayout)
}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771World *ERC2771WorldCaller) GetRecord0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getRecord0", tableId, keyTuple)

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
func (_ERC2771World *ERC2771WorldSession) GetRecord0(tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _ERC2771World.Contract.GetRecord0(&_ERC2771World.CallOpts, tableId, keyTuple)
}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771World *ERC2771WorldCallerSession) GetRecord0(tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _ERC2771World.Contract.GetRecord0(&_ERC2771World.CallOpts, tableId, keyTuple)
}

// GetSmartCharacterClassId is a free data retrieval call binding the contract method 0x95b9e2d3.
//
// Solidity: function getSmartCharacterClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCaller) GetSmartCharacterClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getSmartCharacterClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmartCharacterClassId is a free data retrieval call binding the contract method 0x95b9e2d3.
//
// Solidity: function getSmartCharacterClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldSession) GetSmartCharacterClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetSmartCharacterClassId(&_ERC2771World.CallOpts)
}

// GetSmartCharacterClassId is a free data retrieval call binding the contract method 0x95b9e2d3.
//
// Solidity: function getSmartCharacterClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCallerSession) GetSmartCharacterClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetSmartCharacterClassId(&_ERC2771World.CallOpts)
}

// GetSmartGateClassId is a free data retrieval call binding the contract method 0xc7afc974.
//
// Solidity: function getSmartGateClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCaller) GetSmartGateClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getSmartGateClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmartGateClassId is a free data retrieval call binding the contract method 0xc7afc974.
//
// Solidity: function getSmartGateClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldSession) GetSmartGateClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetSmartGateClassId(&_ERC2771World.CallOpts)
}

// GetSmartGateClassId is a free data retrieval call binding the contract method 0xc7afc974.
//
// Solidity: function getSmartGateClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCallerSession) GetSmartGateClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetSmartGateClassId(&_ERC2771World.CallOpts)
}

// GetSmartStorageUnitClassId is a free data retrieval call binding the contract method 0xf3a7f1bb.
//
// Solidity: function getSmartStorageUnitClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCaller) GetSmartStorageUnitClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getSmartStorageUnitClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmartStorageUnitClassId is a free data retrieval call binding the contract method 0xf3a7f1bb.
//
// Solidity: function getSmartStorageUnitClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldSession) GetSmartStorageUnitClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetSmartStorageUnitClassId(&_ERC2771World.CallOpts)
}

// GetSmartStorageUnitClassId is a free data retrieval call binding the contract method 0xf3a7f1bb.
//
// Solidity: function getSmartStorageUnitClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCallerSession) GetSmartStorageUnitClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetSmartStorageUnitClassId(&_ERC2771World.CallOpts)
}

// GetSmartTurretClassId is a free data retrieval call binding the contract method 0x4967ae58.
//
// Solidity: function getSmartTurretClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCaller) GetSmartTurretClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getSmartTurretClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmartTurretClassId is a free data retrieval call binding the contract method 0x4967ae58.
//
// Solidity: function getSmartTurretClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldSession) GetSmartTurretClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetSmartTurretClassId(&_ERC2771World.CallOpts)
}

// GetSmartTurretClassId is a free data retrieval call binding the contract method 0x4967ae58.
//
// Solidity: function getSmartTurretClassId() view returns(uint256)
func (_ERC2771World *ERC2771WorldCallerSession) GetSmartTurretClassId() (*big.Int, error) {
	return _ERC2771World.Contract.GetSmartTurretClassId(&_ERC2771World.CallOpts)
}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_ERC2771World *ERC2771WorldCaller) GetStaticField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getStaticField", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_ERC2771World *ERC2771WorldSession) GetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	return _ERC2771World.Contract.GetStaticField(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_ERC2771World *ERC2771WorldCallerSession) GetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	return _ERC2771World.Contract.GetStaticField(&_ERC2771World.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_ERC2771World *ERC2771WorldCaller) GetValueSchema(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "getValueSchema", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_ERC2771World *ERC2771WorldSession) GetValueSchema(tableId [32]byte) ([32]byte, error) {
	return _ERC2771World.Contract.GetValueSchema(&_ERC2771World.CallOpts, tableId)
}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_ERC2771World *ERC2771WorldCallerSession) GetValueSchema(tableId [32]byte) ([32]byte, error) {
	return _ERC2771World.Contract.GetValueSchema(&_ERC2771World.CallOpts, tableId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) IsAdmin(opts *bind.CallOpts, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "isAdmin", caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) IsAdmin(caller common.Address) (bool, error) {
	return _ERC2771World.Contract.IsAdmin(&_ERC2771World.CallOpts, caller)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) IsAdmin(caller common.Address) (bool, error) {
	return _ERC2771World.Contract.IsAdmin(&_ERC2771World.CallOpts, caller)
}

// IsAnyGateLinked is a free data retrieval call binding the contract method 0x7222f910.
//
// Solidity: function isAnyGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) IsAnyGateLinked(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "isAnyGateLinked", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAnyGateLinked is a free data retrieval call binding the contract method 0x7222f910.
//
// Solidity: function isAnyGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) IsAnyGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771World.Contract.IsAnyGateLinked(&_ERC2771World.CallOpts, sourceGateId, destinationGateId)
}

// IsAnyGateLinked is a free data retrieval call binding the contract method 0x7222f910.
//
// Solidity: function isAnyGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) IsAnyGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771World.Contract.IsAnyGateLinked(&_ERC2771World.CallOpts, sourceGateId, destinationGateId)
}

// IsClassScoped is a free data retrieval call binding the contract method 0x590d5cdd.
//
// Solidity: function isClassScoped(uint256 classId, bytes32 systemId) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) IsClassScoped(opts *bind.CallOpts, classId *big.Int, systemId [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "isClassScoped", classId, systemId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClassScoped is a free data retrieval call binding the contract method 0x590d5cdd.
//
// Solidity: function isClassScoped(uint256 classId, bytes32 systemId) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) IsClassScoped(classId *big.Int, systemId [32]byte) (bool, error) {
	return _ERC2771World.Contract.IsClassScoped(&_ERC2771World.CallOpts, classId, systemId)
}

// IsClassScoped is a free data retrieval call binding the contract method 0x590d5cdd.
//
// Solidity: function isClassScoped(uint256 classId, bytes32 systemId) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) IsClassScoped(classId *big.Int, systemId [32]byte) (bool, error) {
	return _ERC2771World.Contract.IsClassScoped(&_ERC2771World.CallOpts, classId, systemId)
}

// IsEphemeralOwner is a free data retrieval call binding the contract method 0x9f218381.
//
// Solidity: function isEphemeralOwner(uint256 smartObjectId, address caller, bytes data) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) IsEphemeralOwner(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "isEphemeralOwner", smartObjectId, caller, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEphemeralOwner is a free data retrieval call binding the contract method 0x9f218381.
//
// Solidity: function isEphemeralOwner(uint256 smartObjectId, address caller, bytes data) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) IsEphemeralOwner(smartObjectId *big.Int, caller common.Address, data []byte) (bool, error) {
	return _ERC2771World.Contract.IsEphemeralOwner(&_ERC2771World.CallOpts, smartObjectId, caller, data)
}

// IsEphemeralOwner is a free data retrieval call binding the contract method 0x9f218381.
//
// Solidity: function isEphemeralOwner(uint256 smartObjectId, address caller, bytes data) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) IsEphemeralOwner(smartObjectId *big.Int, caller common.Address, data []byte) (bool, error) {
	return _ERC2771World.Contract.IsEphemeralOwner(&_ERC2771World.CallOpts, smartObjectId, caller, data)
}

// IsGateLinked is a free data retrieval call binding the contract method 0x2ab90d4d.
//
// Solidity: function isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) IsGateLinked(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "isGateLinked", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGateLinked is a free data retrieval call binding the contract method 0x2ab90d4d.
//
// Solidity: function isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) IsGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771World.Contract.IsGateLinked(&_ERC2771World.CallOpts, sourceGateId, destinationGateId)
}

// IsGateLinked is a free data retrieval call binding the contract method 0x2ab90d4d.
//
// Solidity: function isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) IsGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771World.Contract.IsGateLinked(&_ERC2771World.CallOpts, sourceGateId, destinationGateId)
}

// IsOwner is a free data retrieval call binding the contract method 0x5a5d096c.
//
// Solidity: function isOwner(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) IsOwner(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "isOwner", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x5a5d096c.
//
// Solidity: function isOwner(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) IsOwner(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.IsOwner(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// IsOwner is a free data retrieval call binding the contract method 0x5a5d096c.
//
// Solidity: function isOwner(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) IsOwner(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771World.Contract.IsOwner(&_ERC2771World.CallOpts, smartObjectId, caller)
}

// IsOwnerOfBothGates is a free data retrieval call binding the contract method 0xa134720c.
//
// Solidity: function isOwnerOfBothGates(address caller, bytes data) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) IsOwnerOfBothGates(opts *bind.CallOpts, caller common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "isOwnerOfBothGates", caller, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwnerOfBothGates is a free data retrieval call binding the contract method 0xa134720c.
//
// Solidity: function isOwnerOfBothGates(address caller, bytes data) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) IsOwnerOfBothGates(caller common.Address, data []byte) (bool, error) {
	return _ERC2771World.Contract.IsOwnerOfBothGates(&_ERC2771World.CallOpts, caller, data)
}

// IsOwnerOfBothGates is a free data retrieval call binding the contract method 0xa134720c.
//
// Solidity: function isOwnerOfBothGates(address caller, bytes data) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) IsOwnerOfBothGates(caller common.Address, data []byte) (bool, error) {
	return _ERC2771World.Contract.IsOwnerOfBothGates(&_ERC2771World.CallOpts, caller, data)
}

// IsWithinRange is a free data retrieval call binding the contract method 0x36dfd147.
//
// Solidity: function isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldCaller) IsWithinRange(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "isWithinRange", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithinRange is a free data retrieval call binding the contract method 0x36dfd147.
//
// Solidity: function isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldSession) IsWithinRange(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771World.Contract.IsWithinRange(&_ERC2771World.CallOpts, sourceGateId, destinationGateId)
}

// IsWithinRange is a free data retrieval call binding the contract method 0x36dfd147.
//
// Solidity: function isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771World *ERC2771WorldCallerSession) IsWithinRange(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771World.Contract.IsWithinRange(&_ERC2771World.CallOpts, sourceGateId, destinationGateId)
}

// OnlyAdminOrCallAccess is a free data retrieval call binding the contract method 0x1d5a7107.
//
// Solidity: function onlyAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyAdminOrCallAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyAdminOrCallAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrCallAccess is a free data retrieval call binding the contract method 0x1d5a7107.
//
// Solidity: function onlyAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrCallAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminOrCallAccess is a free data retrieval call binding the contract method 0x1d5a7107.
//
// Solidity: function onlyAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrCallAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminOrClassScopedAccess is a free data retrieval call binding the contract method 0x31098e24.
//
// Solidity: function onlyAdminOrClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyAdminOrClassScopedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyAdminOrClassScopedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrClassScopedAccess is a free data retrieval call binding the contract method 0x31098e24.
//
// Solidity: function onlyAdminOrClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyAdminOrClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrClassScopedAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminOrClassScopedAccess is a free data retrieval call binding the contract method 0x31098e24.
//
// Solidity: function onlyAdminOrClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyAdminOrClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrClassScopedAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminOrOwner is a free data retrieval call binding the contract method 0xd2df45d4.
//
// Solidity: function onlyAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyAdminOrOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyAdminOrOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrOwner is a free data retrieval call binding the contract method 0xd2df45d4.
//
// Solidity: function onlyAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminOrOwner is a free data retrieval call binding the contract method 0xd2df45d4.
//
// Solidity: function onlyAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminOrOwnerSupported is a free data retrieval call binding the contract method 0xf4682823.
//
// Solidity: function onlyAdminOrOwnerSupported(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyAdminOrOwnerSupported(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyAdminOrOwnerSupported", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrOwnerSupported is a free data retrieval call binding the contract method 0xf4682823.
//
// Solidity: function onlyAdminOrOwnerSupported(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyAdminOrOwnerSupported(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrOwnerSupported(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminOrOwnerSupported is a free data retrieval call binding the contract method 0xf4682823.
//
// Solidity: function onlyAdminOrOwnerSupported(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyAdminOrOwnerSupported(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrOwnerSupported(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminOrScopeEnforcedCall is a free data retrieval call binding the contract method 0xb65e6ac2.
//
// Solidity: function onlyAdminOrScopeEnforcedCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyAdminOrScopeEnforcedCall(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyAdminOrScopeEnforcedCall", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrScopeEnforcedCall is a free data retrieval call binding the contract method 0xb65e6ac2.
//
// Solidity: function onlyAdminOrScopeEnforcedCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyAdminOrScopeEnforcedCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrScopeEnforcedCall(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminOrScopeEnforcedCall is a free data retrieval call binding the contract method 0xb65e6ac2.
//
// Solidity: function onlyAdminOrScopeEnforcedCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyAdminOrScopeEnforcedCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminOrScopeEnforcedCall(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminSupportedAccess is a free data retrieval call binding the contract method 0x12fea21c.
//
// Solidity: function onlyAdminSupportedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyAdminSupportedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyAdminSupportedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminSupportedAccess is a free data retrieval call binding the contract method 0x12fea21c.
//
// Solidity: function onlyAdminSupportedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyAdminSupportedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminSupportedAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminSupportedAccess is a free data retrieval call binding the contract method 0x12fea21c.
//
// Solidity: function onlyAdminSupportedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyAdminSupportedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminSupportedAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminSupportedOwnerOrCall is a free data retrieval call binding the contract method 0x03e4c983.
//
// Solidity: function onlyAdminSupportedOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyAdminSupportedOwnerOrCall(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyAdminSupportedOwnerOrCall", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminSupportedOwnerOrCall is a free data retrieval call binding the contract method 0x03e4c983.
//
// Solidity: function onlyAdminSupportedOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyAdminSupportedOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminSupportedOwnerOrCall(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyAdminSupportedOwnerOrCall is a free data retrieval call binding the contract method 0x03e4c983.
//
// Solidity: function onlyAdminSupportedOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyAdminSupportedOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyAdminSupportedOwnerOrCall(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyCallAccess is a free data retrieval call binding the contract method 0xa58a5393.
//
// Solidity: function onlyCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyCallAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyCallAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyCallAccess is a free data retrieval call binding the contract method 0xa58a5393.
//
// Solidity: function onlyCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyCallAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyCallAccess is a free data retrieval call binding the contract method 0xa58a5393.
//
// Solidity: function onlyCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyCallAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyCallAccessOrDirectEphemeralOwner is a free data retrieval call binding the contract method 0xee906526.
//
// Solidity: function onlyCallAccessOrDirectEphemeralOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyCallAccessOrDirectEphemeralOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyCallAccessOrDirectEphemeralOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyCallAccessOrDirectEphemeralOwner is a free data retrieval call binding the contract method 0xee906526.
//
// Solidity: function onlyCallAccessOrDirectEphemeralOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyCallAccessOrDirectEphemeralOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyCallAccessOrDirectEphemeralOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyCallAccessOrDirectEphemeralOwner is a free data retrieval call binding the contract method 0xee906526.
//
// Solidity: function onlyCallAccessOrDirectEphemeralOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyCallAccessOrDirectEphemeralOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyCallAccessOrDirectEphemeralOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyCallAccessWithScopeEnforced is a free data retrieval call binding the contract method 0xdde2f623.
//
// Solidity: function onlyCallAccessWithScopeEnforced(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyCallAccessWithScopeEnforced(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyCallAccessWithScopeEnforced", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyCallAccessWithScopeEnforced is a free data retrieval call binding the contract method 0xdde2f623.
//
// Solidity: function onlyCallAccessWithScopeEnforced(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyCallAccessWithScopeEnforced(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyCallAccessWithScopeEnforced(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyCallAccessWithScopeEnforced is a free data retrieval call binding the contract method 0xdde2f623.
//
// Solidity: function onlyCallAccessWithScopeEnforced(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyCallAccessWithScopeEnforced(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyCallAccessWithScopeEnforced(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyClassScopedAccess is a free data retrieval call binding the contract method 0xf78da887.
//
// Solidity: function onlyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyClassScopedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyClassScopedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyClassScopedAccess is a free data retrieval call binding the contract method 0xf78da887.
//
// Solidity: function onlyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyClassScopedAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyClassScopedAccess is a free data retrieval call binding the contract method 0xf78da887.
//
// Solidity: function onlyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyClassScopedAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyClassScopedOrCharAdminOrOwner is a free data retrieval call binding the contract method 0xcc8636d8.
//
// Solidity: function onlyClassScopedOrCharAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyClassScopedOrCharAdminOrOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyClassScopedOrCharAdminOrOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyClassScopedOrCharAdminOrOwner is a free data retrieval call binding the contract method 0xcc8636d8.
//
// Solidity: function onlyClassScopedOrCharAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyClassScopedOrCharAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyClassScopedOrCharAdminOrOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyClassScopedOrCharAdminOrOwner is a free data retrieval call binding the contract method 0xcc8636d8.
//
// Solidity: function onlyClassScopedOrCharAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyClassScopedOrCharAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyClassScopedOrCharAdminOrOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyDirectAdmin is a free data retrieval call binding the contract method 0x542b659d.
//
// Solidity: function onlyDirectAdmin(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyDirectAdmin(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyDirectAdmin", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyDirectAdmin is a free data retrieval call binding the contract method 0x542b659d.
//
// Solidity: function onlyDirectAdmin(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyDirectAdmin(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyDirectAdmin(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyDirectAdmin is a free data retrieval call binding the contract method 0x542b659d.
//
// Solidity: function onlyDirectAdmin(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyDirectAdmin(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyDirectAdmin(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyDirectAdminOrCallAccess is a free data retrieval call binding the contract method 0xd7fe3c1c.
//
// Solidity: function onlyDirectAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyDirectAdminOrCallAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyDirectAdminOrCallAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyDirectAdminOrCallAccess is a free data retrieval call binding the contract method 0xd7fe3c1c.
//
// Solidity: function onlyDirectAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyDirectAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyDirectAdminOrCallAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyDirectAdminOrCallAccess is a free data retrieval call binding the contract method 0xd7fe3c1c.
//
// Solidity: function onlyDirectAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyDirectAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyDirectAdminOrCallAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyDirectEphemeralOwnerOrCall is a free data retrieval call binding the contract method 0xf10f61ca.
//
// Solidity: function onlyDirectEphemeralOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyDirectEphemeralOwnerOrCall(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyDirectEphemeralOwnerOrCall", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyDirectEphemeralOwnerOrCall is a free data retrieval call binding the contract method 0xf10f61ca.
//
// Solidity: function onlyDirectEphemeralOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyDirectEphemeralOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyDirectEphemeralOwnerOrCall(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyDirectEphemeralOwnerOrCall is a free data retrieval call binding the contract method 0xf10f61ca.
//
// Solidity: function onlyDirectEphemeralOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyDirectEphemeralOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyDirectEphemeralOwnerOrCall(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyDirectOwner is a free data retrieval call binding the contract method 0x3ac45a0c.
//
// Solidity: function onlyDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyDirectOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyDirectOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyDirectOwner is a free data retrieval call binding the contract method 0x3ac45a0c.
//
// Solidity: function onlyDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyDirectOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyDirectOwner is a free data retrieval call binding the contract method 0x3ac45a0c.
//
// Solidity: function onlyDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyDirectOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyEphemeralOwnerOrTransferRole is a free data retrieval call binding the contract method 0x5eee0cc8.
//
// Solidity: function onlyEphemeralOwnerOrTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyEphemeralOwnerOrTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyEphemeralOwnerOrTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyEphemeralOwnerOrTransferRole is a free data retrieval call binding the contract method 0x5eee0cc8.
//
// Solidity: function onlyEphemeralOwnerOrTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyEphemeralOwnerOrTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyEphemeralOwnerOrTransferRole(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyEphemeralOwnerOrTransferRole is a free data retrieval call binding the contract method 0x5eee0cc8.
//
// Solidity: function onlyEphemeralOwnerOrTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyEphemeralOwnerOrTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyEphemeralOwnerOrTransferRole(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwner is a free data retrieval call binding the contract method 0x4ded2d51.
//
// Solidity: function onlyOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwner is a free data retrieval call binding the contract method 0x4ded2d51.
//
// Solidity: function onlyOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwner is a free data retrieval call binding the contract method 0x4ded2d51.
//
// Solidity: function onlyOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwner(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrEphemeralCrossTransferRole is a free data retrieval call binding the contract method 0xd061b071.
//
// Solidity: function onlyOwnerOrEphemeralCrossTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyOwnerOrEphemeralCrossTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyOwnerOrEphemeralCrossTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwnerOrEphemeralCrossTransferRole is a free data retrieval call binding the contract method 0xd061b071.
//
// Solidity: function onlyOwnerOrEphemeralCrossTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyOwnerOrEphemeralCrossTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwnerOrEphemeralCrossTransferRole(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrEphemeralCrossTransferRole is a free data retrieval call binding the contract method 0xd061b071.
//
// Solidity: function onlyOwnerOrEphemeralCrossTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyOwnerOrEphemeralCrossTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwnerOrEphemeralCrossTransferRole(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrEphemeralTransferRole is a free data retrieval call binding the contract method 0x2227febd.
//
// Solidity: function onlyOwnerOrEphemeralTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyOwnerOrEphemeralTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyOwnerOrEphemeralTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwnerOrEphemeralTransferRole is a free data retrieval call binding the contract method 0x2227febd.
//
// Solidity: function onlyOwnerOrEphemeralTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyOwnerOrEphemeralTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwnerOrEphemeralTransferRole(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrEphemeralTransferRole is a free data retrieval call binding the contract method 0x2227febd.
//
// Solidity: function onlyOwnerOrEphemeralTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyOwnerOrEphemeralTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwnerOrEphemeralTransferRole(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrInventoryTransferRole is a free data retrieval call binding the contract method 0x992518f9.
//
// Solidity: function onlyOwnerOrInventoryTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyOwnerOrInventoryTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyOwnerOrInventoryTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwnerOrInventoryTransferRole is a free data retrieval call binding the contract method 0x992518f9.
//
// Solidity: function onlyOwnerOrInventoryTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyOwnerOrInventoryTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwnerOrInventoryTransferRole(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrInventoryTransferRole is a free data retrieval call binding the contract method 0x992518f9.
//
// Solidity: function onlyOwnerOrInventoryTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyOwnerOrInventoryTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwnerOrInventoryTransferRole(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwnerWithAdminSupportAccess is a free data retrieval call binding the contract method 0x036867ec.
//
// Solidity: function onlyOwnerWithAdminSupportAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlyOwnerWithAdminSupportAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlyOwnerWithAdminSupportAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwnerWithAdminSupportAccess is a free data retrieval call binding the contract method 0x036867ec.
//
// Solidity: function onlyOwnerWithAdminSupportAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlyOwnerWithAdminSupportAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwnerWithAdminSupportAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlyOwnerWithAdminSupportAccess is a free data retrieval call binding the contract method 0x036867ec.
//
// Solidity: function onlyOwnerWithAdminSupportAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlyOwnerWithAdminSupportAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlyOwnerWithAdminSupportAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlySmartAssemblyClassScopedAccess is a free data retrieval call binding the contract method 0xbccd5c2a.
//
// Solidity: function onlySmartAssemblyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCaller) OnlySmartAssemblyClassScopedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "onlySmartAssemblyClassScopedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlySmartAssemblyClassScopedAccess is a free data retrieval call binding the contract method 0xbccd5c2a.
//
// Solidity: function onlySmartAssemblyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldSession) OnlySmartAssemblyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlySmartAssemblyClassScopedAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// OnlySmartAssemblyClassScopedAccess is a free data retrieval call binding the contract method 0xbccd5c2a.
//
// Solidity: function onlySmartAssemblyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771World *ERC2771WorldCallerSession) OnlySmartAssemblyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771World.Contract.OnlySmartAssemblyClassScopedAccess(&_ERC2771World.CallOpts, smartObjectId, data)
}

// Owner is a free data retrieval call binding the contract method 0xa123c33e.
//
// Solidity: function owner(uint256 smartObjectId) view returns(address)
func (_ERC2771World *ERC2771WorldCaller) Owner(opts *bind.CallOpts, smartObjectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "owner", smartObjectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xa123c33e.
//
// Solidity: function owner(uint256 smartObjectId) view returns(address)
func (_ERC2771World *ERC2771WorldSession) Owner(smartObjectId *big.Int) (common.Address, error) {
	return _ERC2771World.Contract.Owner(&_ERC2771World.CallOpts, smartObjectId)
}

// Owner is a free data retrieval call binding the contract method 0xa123c33e.
//
// Solidity: function owner(uint256 smartObjectId) view returns(address)
func (_ERC2771World *ERC2771WorldCallerSession) Owner(smartObjectId *big.Int) (common.Address, error) {
	return _ERC2771World.Contract.Owner(&_ERC2771World.CallOpts, smartObjectId)
}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_ERC2771World *ERC2771WorldCaller) StoreVersion(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "storeVersion")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_ERC2771World *ERC2771WorldSession) StoreVersion() ([32]byte, error) {
	return _ERC2771World.Contract.StoreVersion(&_ERC2771World.CallOpts)
}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_ERC2771World *ERC2771WorldCallerSession) StoreVersion() ([32]byte, error) {
	return _ERC2771World.Contract.StoreVersion(&_ERC2771World.CallOpts)
}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_ERC2771World *ERC2771WorldCaller) WorldVersion(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "worldVersion")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_ERC2771World *ERC2771WorldSession) WorldVersion() ([32]byte, error) {
	return _ERC2771World.Contract.WorldVersion(&_ERC2771World.CallOpts)
}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_ERC2771World *ERC2771WorldCallerSession) WorldVersion() ([32]byte, error) {
	return _ERC2771World.Contract.WorldVersion(&_ERC2771World.CallOpts)
}

// HandleNodeOffline is a paid mutator transaction binding the contract method 0x5df931f9.
//
// Solidity: function _handleNodeOffline(uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactor) HandleNodeOffline(opts *bind.TransactOpts, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "_handleNodeOffline", networkNodeId)
}

// HandleNodeOffline is a paid mutator transaction binding the contract method 0x5df931f9.
//
// Solidity: function _handleNodeOffline(uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldSession) HandleNodeOffline(networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.HandleNodeOffline(&_ERC2771World.TransactOpts, networkNodeId)
}

// HandleNodeOffline is a paid mutator transaction binding the contract method 0x5df931f9.
//
// Solidity: function _handleNodeOffline(uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) HandleNodeOffline(networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.HandleNodeOffline(&_ERC2771World.TransactOpts, networkNodeId)
}

// Aggression is a paid mutator transaction binding the contract method 0x8ba2b5d8.
//
// Solidity: function aggression((uint256,((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[],(uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256)) params) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldTransactor) Aggression(opts *bind.TransactOpts, params AggressionParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "aggression", params)
}

// Aggression is a paid mutator transaction binding the contract method 0x8ba2b5d8.
//
// Solidity: function aggression((uint256,((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[],(uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256)) params) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldSession) Aggression(params AggressionParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.Aggression(&_ERC2771World.TransactOpts, params)
}

// Aggression is a paid mutator transaction binding the contract method 0x8ba2b5d8.
//
// Solidity: function aggression((uint256,((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[],(uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256)) params) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldTransactorSession) Aggression(params AggressionParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.Aggression(&_ERC2771World.TransactOpts, params)
}

// Anchor is a paid mutator transaction binding the contract method 0x4ccd479d.
//
// Solidity: function anchor(uint256 smartObjectId, address owner, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771World *ERC2771WorldTransactor) Anchor(opts *bind.TransactOpts, smartObjectId *big.Int, owner common.Address, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "anchor", smartObjectId, owner, locationData)
}

// Anchor is a paid mutator transaction binding the contract method 0x4ccd479d.
//
// Solidity: function anchor(uint256 smartObjectId, address owner, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771World *ERC2771WorldSession) Anchor(smartObjectId *big.Int, owner common.Address, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771World.Contract.Anchor(&_ERC2771World.TransactOpts, smartObjectId, owner, locationData)
}

// Anchor is a paid mutator transaction binding the contract method 0x4ccd479d.
//
// Solidity: function anchor(uint256 smartObjectId, address owner, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) Anchor(smartObjectId *big.Int, owner common.Address, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771World.Contract.Anchor(&_ERC2771World.TransactOpts, smartObjectId, owner, locationData)
}

// AssignItemToInventory is a paid mutator transaction binding the contract method 0xe2957099.
//
// Solidity: function assignItemToInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771World *ERC2771WorldTransactor) AssignItemToInventory(opts *bind.TransactOpts, inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "assignItemToInventory", inventoryObjectId, itemObjectId, quantity)
}

// AssignItemToInventory is a paid mutator transaction binding the contract method 0xe2957099.
//
// Solidity: function assignItemToInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771World *ERC2771WorldSession) AssignItemToInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.AssignItemToInventory(&_ERC2771World.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// AssignItemToInventory is a paid mutator transaction binding the contract method 0xe2957099.
//
// Solidity: function assignItemToInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) AssignItemToInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.AssignItemToInventory(&_ERC2771World.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// AssignOwner is a paid mutator transaction binding the contract method 0xe671644b.
//
// Solidity: function assignOwner(uint256 smartObjectId, address to) returns()
func (_ERC2771World *ERC2771WorldTransactor) AssignOwner(opts *bind.TransactOpts, smartObjectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "assignOwner", smartObjectId, to)
}

// AssignOwner is a paid mutator transaction binding the contract method 0xe671644b.
//
// Solidity: function assignOwner(uint256 smartObjectId, address to) returns()
func (_ERC2771World *ERC2771WorldSession) AssignOwner(smartObjectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.AssignOwner(&_ERC2771World.TransactOpts, smartObjectId, to)
}

// AssignOwner is a paid mutator transaction binding the contract method 0xe671644b.
//
// Solidity: function assignOwner(uint256 smartObjectId, address to) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) AssignOwner(smartObjectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.AssignOwner(&_ERC2771World.TransactOpts, smartObjectId, to)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771World *ERC2771WorldTransactor) BatchCall(opts *bind.TransactOpts, systemCalls []SystemCallData) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "batchCall", systemCalls)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771World *ERC2771WorldSession) BatchCall(systemCalls []SystemCallData) (*types.Transaction, error) {
	return _ERC2771World.Contract.BatchCall(&_ERC2771World.TransactOpts, systemCalls)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771World *ERC2771WorldTransactorSession) BatchCall(systemCalls []SystemCallData) (*types.Transaction, error) {
	return _ERC2771World.Contract.BatchCall(&_ERC2771World.TransactOpts, systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771World *ERC2771WorldTransactor) BatchCallFrom(opts *bind.TransactOpts, systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "batchCallFrom", systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771World *ERC2771WorldSession) BatchCallFrom(systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _ERC2771World.Contract.BatchCallFrom(&_ERC2771World.TransactOpts, systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771World *ERC2771WorldTransactorSession) BatchCallFrom(systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _ERC2771World.Contract.BatchCallFrom(&_ERC2771World.TransactOpts, systemCalls)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactor) BringOffline(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "bringOffline", smartObjectId)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldSession) BringOffline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.BringOffline(&_ERC2771World.TransactOpts, smartObjectId)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) BringOffline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.BringOffline(&_ERC2771World.TransactOpts, smartObjectId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactor) BringOnline(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "bringOnline", smartObjectId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldSession) BringOnline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.BringOnline(&_ERC2771World.TransactOpts, smartObjectId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) BringOnline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.BringOnline(&_ERC2771World.TransactOpts, smartObjectId)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771World *ERC2771WorldTransactor) Call(opts *bind.TransactOpts, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "call", systemId, callData)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771World *ERC2771WorldSession) Call(systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.Call(&_ERC2771World.TransactOpts, systemId, callData)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771World *ERC2771WorldTransactorSession) Call(systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.Call(&_ERC2771World.TransactOpts, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771World *ERC2771WorldTransactor) CallFrom(opts *bind.TransactOpts, delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "callFrom", delegator, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771World *ERC2771WorldSession) CallFrom(delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.CallFrom(&_ERC2771World.TransactOpts, delegator, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771World *ERC2771WorldTransactorSession) CallFrom(delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.CallFrom(&_ERC2771World.TransactOpts, delegator, systemId, callData)
}

// CanJump is a paid mutator transaction binding the contract method 0xc9cf8ac5.
//
// Solidity: function canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_ERC2771World *ERC2771WorldTransactor) CanJump(opts *bind.TransactOpts, characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "canJump", characterId, sourceGateId, destinationGateId)
}

// CanJump is a paid mutator transaction binding the contract method 0xc9cf8ac5.
//
// Solidity: function canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_ERC2771World *ERC2771WorldSession) CanJump(characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CanJump(&_ERC2771World.TransactOpts, characterId, sourceGateId, destinationGateId)
}

// CanJump is a paid mutator transaction binding the contract method 0xc9cf8ac5.
//
// Solidity: function canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_ERC2771World *ERC2771WorldTransactorSession) CanJump(characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CanJump(&_ERC2771World.TransactOpts, characterId, sourceGateId, destinationGateId)
}

// ConfigureDeployableAccess is a paid mutator transaction binding the contract method 0x03d05669.
//
// Solidity: function configureDeployableAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureDeployableAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureDeployableAccess")
}

// ConfigureDeployableAccess is a paid mutator transaction binding the contract method 0x03d05669.
//
// Solidity: function configureDeployableAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureDeployableAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureDeployableAccess(&_ERC2771World.TransactOpts)
}

// ConfigureDeployableAccess is a paid mutator transaction binding the contract method 0x03d05669.
//
// Solidity: function configureDeployableAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureDeployableAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureDeployableAccess(&_ERC2771World.TransactOpts)
}

// ConfigureEntityRecordAccess is a paid mutator transaction binding the contract method 0xc36c2700.
//
// Solidity: function configureEntityRecordAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureEntityRecordAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureEntityRecordAccess")
}

// ConfigureEntityRecordAccess is a paid mutator transaction binding the contract method 0xc36c2700.
//
// Solidity: function configureEntityRecordAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureEntityRecordAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureEntityRecordAccess(&_ERC2771World.TransactOpts)
}

// ConfigureEntityRecordAccess is a paid mutator transaction binding the contract method 0xc36c2700.
//
// Solidity: function configureEntityRecordAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureEntityRecordAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureEntityRecordAccess(&_ERC2771World.TransactOpts)
}

// ConfigureEphemeralInteractAccess is a paid mutator transaction binding the contract method 0x1a3be1ed.
//
// Solidity: function configureEphemeralInteractAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureEphemeralInteractAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureEphemeralInteractAccess")
}

// ConfigureEphemeralInteractAccess is a paid mutator transaction binding the contract method 0x1a3be1ed.
//
// Solidity: function configureEphemeralInteractAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureEphemeralInteractAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureEphemeralInteractAccess(&_ERC2771World.TransactOpts)
}

// ConfigureEphemeralInteractAccess is a paid mutator transaction binding the contract method 0x1a3be1ed.
//
// Solidity: function configureEphemeralInteractAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureEphemeralInteractAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureEphemeralInteractAccess(&_ERC2771World.TransactOpts)
}

// ConfigureEphemeralInventoryAccess is a paid mutator transaction binding the contract method 0xc8367523.
//
// Solidity: function configureEphemeralInventoryAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureEphemeralInventoryAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureEphemeralInventoryAccess")
}

// ConfigureEphemeralInventoryAccess is a paid mutator transaction binding the contract method 0xc8367523.
//
// Solidity: function configureEphemeralInventoryAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureEphemeralInventoryAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureEphemeralInventoryAccess(&_ERC2771World.TransactOpts)
}

// ConfigureEphemeralInventoryAccess is a paid mutator transaction binding the contract method 0xc8367523.
//
// Solidity: function configureEphemeralInventoryAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureEphemeralInventoryAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureEphemeralInventoryAccess(&_ERC2771World.TransactOpts)
}

// ConfigureFuelAccess is a paid mutator transaction binding the contract method 0xde941597.
//
// Solidity: function configureFuelAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureFuelAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureFuelAccess")
}

// ConfigureFuelAccess is a paid mutator transaction binding the contract method 0xde941597.
//
// Solidity: function configureFuelAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureFuelAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureFuelAccess(&_ERC2771World.TransactOpts)
}

// ConfigureFuelAccess is a paid mutator transaction binding the contract method 0xde941597.
//
// Solidity: function configureFuelAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureFuelAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureFuelAccess(&_ERC2771World.TransactOpts)
}

// ConfigureFuelEfficiency is a paid mutator transaction binding the contract method 0x05f9277b.
//
// Solidity: function configureFuelEfficiency(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) fuelEntityParams, uint256 fuelEfficiency) returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureFuelEfficiency(opts *bind.TransactOpts, smartObjectId *big.Int, fuelEntityParams EntityRecordParams, fuelEfficiency *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureFuelEfficiency", smartObjectId, fuelEntityParams, fuelEfficiency)
}

// ConfigureFuelEfficiency is a paid mutator transaction binding the contract method 0x05f9277b.
//
// Solidity: function configureFuelEfficiency(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) fuelEntityParams, uint256 fuelEfficiency) returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureFuelEfficiency(smartObjectId *big.Int, fuelEntityParams EntityRecordParams, fuelEfficiency *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureFuelEfficiency(&_ERC2771World.TransactOpts, smartObjectId, fuelEntityParams, fuelEfficiency)
}

// ConfigureFuelEfficiency is a paid mutator transaction binding the contract method 0x05f9277b.
//
// Solidity: function configureFuelEfficiency(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) fuelEntityParams, uint256 fuelEfficiency) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureFuelEfficiency(smartObjectId *big.Int, fuelEntityParams EntityRecordParams, fuelEfficiency *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureFuelEfficiency(&_ERC2771World.TransactOpts, smartObjectId, fuelEntityParams, fuelEfficiency)
}

// ConfigureFuelParameters is a paid mutator transaction binding the contract method 0xc3504c22.
//
// Solidity: function configureFuelParameters(uint256 smartObjectId, (uint256,uint256) fuelParams) returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureFuelParameters(opts *bind.TransactOpts, smartObjectId *big.Int, fuelParams FuelParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureFuelParameters", smartObjectId, fuelParams)
}

// ConfigureFuelParameters is a paid mutator transaction binding the contract method 0xc3504c22.
//
// Solidity: function configureFuelParameters(uint256 smartObjectId, (uint256,uint256) fuelParams) returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureFuelParameters(smartObjectId *big.Int, fuelParams FuelParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureFuelParameters(&_ERC2771World.TransactOpts, smartObjectId, fuelParams)
}

// ConfigureFuelParameters is a paid mutator transaction binding the contract method 0xc3504c22.
//
// Solidity: function configureFuelParameters(uint256 smartObjectId, (uint256,uint256) fuelParams) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureFuelParameters(smartObjectId *big.Int, fuelParams FuelParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureFuelParameters(&_ERC2771World.TransactOpts, smartObjectId, fuelParams)
}

// ConfigureGate is a paid mutator transaction binding the contract method 0xe6611746.
//
// Solidity: function configureGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureGate(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureGate", smartObjectId, systemId)
}

// ConfigureGate is a paid mutator transaction binding the contract method 0xe6611746.
//
// Solidity: function configureGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureGate(&_ERC2771World.TransactOpts, smartObjectId, systemId)
}

// ConfigureGate is a paid mutator transaction binding the contract method 0xe6611746.
//
// Solidity: function configureGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureGate(&_ERC2771World.TransactOpts, smartObjectId, systemId)
}

// ConfigureInventoryAccess is a paid mutator transaction binding the contract method 0xb47cb8f3.
//
// Solidity: function configureInventoryAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureInventoryAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureInventoryAccess")
}

// ConfigureInventoryAccess is a paid mutator transaction binding the contract method 0xb47cb8f3.
//
// Solidity: function configureInventoryAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureInventoryAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureInventoryAccess(&_ERC2771World.TransactOpts)
}

// ConfigureInventoryAccess is a paid mutator transaction binding the contract method 0xb47cb8f3.
//
// Solidity: function configureInventoryAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureInventoryAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureInventoryAccess(&_ERC2771World.TransactOpts)
}

// ConfigureInventoryInteractAccess is a paid mutator transaction binding the contract method 0x3884c463.
//
// Solidity: function configureInventoryInteractAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureInventoryInteractAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureInventoryInteractAccess")
}

// ConfigureInventoryInteractAccess is a paid mutator transaction binding the contract method 0x3884c463.
//
// Solidity: function configureInventoryInteractAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureInventoryInteractAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureInventoryInteractAccess(&_ERC2771World.TransactOpts)
}

// ConfigureInventoryInteractAccess is a paid mutator transaction binding the contract method 0x3884c463.
//
// Solidity: function configureInventoryInteractAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureInventoryInteractAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureInventoryInteractAccess(&_ERC2771World.TransactOpts)
}

// ConfigureKillMailAccess is a paid mutator transaction binding the contract method 0x752918af.
//
// Solidity: function configureKillMailAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureKillMailAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureKillMailAccess")
}

// ConfigureKillMailAccess is a paid mutator transaction binding the contract method 0x752918af.
//
// Solidity: function configureKillMailAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureKillMailAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureKillMailAccess(&_ERC2771World.TransactOpts)
}

// ConfigureKillMailAccess is a paid mutator transaction binding the contract method 0x752918af.
//
// Solidity: function configureKillMailAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureKillMailAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureKillMailAccess(&_ERC2771World.TransactOpts)
}

// ConfigureLocationAccess is a paid mutator transaction binding the contract method 0x520f13da.
//
// Solidity: function configureLocationAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureLocationAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureLocationAccess")
}

// ConfigureLocationAccess is a paid mutator transaction binding the contract method 0x520f13da.
//
// Solidity: function configureLocationAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureLocationAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureLocationAccess(&_ERC2771World.TransactOpts)
}

// ConfigureLocationAccess is a paid mutator transaction binding the contract method 0x520f13da.
//
// Solidity: function configureLocationAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureLocationAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureLocationAccess(&_ERC2771World.TransactOpts)
}

// ConfigureNetworkNodeAccess is a paid mutator transaction binding the contract method 0x38de6231.
//
// Solidity: function configureNetworkNodeAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureNetworkNodeAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureNetworkNodeAccess")
}

// ConfigureNetworkNodeAccess is a paid mutator transaction binding the contract method 0x38de6231.
//
// Solidity: function configureNetworkNodeAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureNetworkNodeAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureNetworkNodeAccess(&_ERC2771World.TransactOpts)
}

// ConfigureNetworkNodeAccess is a paid mutator transaction binding the contract method 0x38de6231.
//
// Solidity: function configureNetworkNodeAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureNetworkNodeAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureNetworkNodeAccess(&_ERC2771World.TransactOpts)
}

// ConfigureOwnershipAccess is a paid mutator transaction binding the contract method 0x4eee463a.
//
// Solidity: function configureOwnershipAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureOwnershipAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureOwnershipAccess")
}

// ConfigureOwnershipAccess is a paid mutator transaction binding the contract method 0x4eee463a.
//
// Solidity: function configureOwnershipAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureOwnershipAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureOwnershipAccess(&_ERC2771World.TransactOpts)
}

// ConfigureOwnershipAccess is a paid mutator transaction binding the contract method 0x4eee463a.
//
// Solidity: function configureOwnershipAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureOwnershipAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureOwnershipAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartAssemblyAccess is a paid mutator transaction binding the contract method 0xafa610cb.
//
// Solidity: function configureSmartAssemblyAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureSmartAssemblyAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureSmartAssemblyAccess")
}

// ConfigureSmartAssemblyAccess is a paid mutator transaction binding the contract method 0xafa610cb.
//
// Solidity: function configureSmartAssemblyAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureSmartAssemblyAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartAssemblyAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartAssemblyAccess is a paid mutator transaction binding the contract method 0xafa610cb.
//
// Solidity: function configureSmartAssemblyAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureSmartAssemblyAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartAssemblyAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartCharacterAccess is a paid mutator transaction binding the contract method 0xa2c4eac8.
//
// Solidity: function configureSmartCharacterAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureSmartCharacterAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureSmartCharacterAccess")
}

// ConfigureSmartCharacterAccess is a paid mutator transaction binding the contract method 0xa2c4eac8.
//
// Solidity: function configureSmartCharacterAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureSmartCharacterAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartCharacterAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartCharacterAccess is a paid mutator transaction binding the contract method 0xa2c4eac8.
//
// Solidity: function configureSmartCharacterAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureSmartCharacterAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartCharacterAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartGateAccess is a paid mutator transaction binding the contract method 0x6238f85c.
//
// Solidity: function configureSmartGateAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureSmartGateAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureSmartGateAccess")
}

// ConfigureSmartGateAccess is a paid mutator transaction binding the contract method 0x6238f85c.
//
// Solidity: function configureSmartGateAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureSmartGateAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartGateAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartGateAccess is a paid mutator transaction binding the contract method 0x6238f85c.
//
// Solidity: function configureSmartGateAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureSmartGateAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartGateAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartStorageUnitAccess is a paid mutator transaction binding the contract method 0x647dc438.
//
// Solidity: function configureSmartStorageUnitAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureSmartStorageUnitAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureSmartStorageUnitAccess")
}

// ConfigureSmartStorageUnitAccess is a paid mutator transaction binding the contract method 0x647dc438.
//
// Solidity: function configureSmartStorageUnitAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureSmartStorageUnitAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartStorageUnitAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartStorageUnitAccess is a paid mutator transaction binding the contract method 0x647dc438.
//
// Solidity: function configureSmartStorageUnitAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureSmartStorageUnitAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartStorageUnitAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartTurretAccess is a paid mutator transaction binding the contract method 0x232dba5b.
//
// Solidity: function configureSmartTurretAccess() returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureSmartTurretAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureSmartTurretAccess")
}

// ConfigureSmartTurretAccess is a paid mutator transaction binding the contract method 0x232dba5b.
//
// Solidity: function configureSmartTurretAccess() returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureSmartTurretAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartTurretAccess(&_ERC2771World.TransactOpts)
}

// ConfigureSmartTurretAccess is a paid mutator transaction binding the contract method 0x232dba5b.
//
// Solidity: function configureSmartTurretAccess() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureSmartTurretAccess() (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartTurretAccess(&_ERC2771World.TransactOpts)
}

// ConfigureTurret is a paid mutator transaction binding the contract method 0x4d7e3af1.
//
// Solidity: function configureTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureTurret(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureTurret", smartObjectId, systemId)
}

// ConfigureTurret is a paid mutator transaction binding the contract method 0x4d7e3af1.
//
// Solidity: function configureTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureTurret(&_ERC2771World.TransactOpts, smartObjectId, systemId)
}

// ConfigureTurret is a paid mutator transaction binding the contract method 0x4d7e3af1.
//
// Solidity: function configureTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureTurret(&_ERC2771World.TransactOpts, smartObjectId, systemId)
}

// ConnectAssembly is a paid mutator transaction binding the contract method 0xe68796c0.
//
// Solidity: function connectAssembly(uint256 networkNodeId, uint256 assemblyId) returns()
func (_ERC2771World *ERC2771WorldTransactor) ConnectAssembly(opts *bind.TransactOpts, networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "connectAssembly", networkNodeId, assemblyId)
}

// ConnectAssembly is a paid mutator transaction binding the contract method 0xe68796c0.
//
// Solidity: function connectAssembly(uint256 networkNodeId, uint256 assemblyId) returns()
func (_ERC2771World *ERC2771WorldSession) ConnectAssembly(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConnectAssembly(&_ERC2771World.TransactOpts, networkNodeId, assemblyId)
}

// ConnectAssembly is a paid mutator transaction binding the contract method 0xe68796c0.
//
// Solidity: function connectAssembly(uint256 networkNodeId, uint256 assemblyId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConnectAssembly(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConnectAssembly(&_ERC2771World.TransactOpts, networkNodeId, assemblyId)
}

// CreateAndAnchor is a paid mutator transaction binding the contract method 0x488ea387.
//
// Solidity: function createAndAnchor((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndAnchor(opts *bind.TransactOpts, params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndAnchor", params, networkNodeId)
}

// CreateAndAnchor is a paid mutator transaction binding the contract method 0x488ea387.
//
// Solidity: function createAndAnchor((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndAnchor(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchor(&_ERC2771World.TransactOpts, params, networkNodeId)
}

// CreateAndAnchor is a paid mutator transaction binding the contract method 0x488ea387.
//
// Solidity: function createAndAnchor((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndAnchor(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchor(&_ERC2771World.TransactOpts, params, networkNodeId)
}

// CreateAndAnchorGate is a paid mutator transaction binding the contract method 0xbb054d08.
//
// Solidity: function createAndAnchorGate((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 maxDistance, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndAnchorGate(opts *bind.TransactOpts, params CreateAndAnchorParams, maxDistance *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndAnchorGate", params, maxDistance, networkNodeId)
}

// CreateAndAnchorGate is a paid mutator transaction binding the contract method 0xbb054d08.
//
// Solidity: function createAndAnchorGate((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 maxDistance, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndAnchorGate(params CreateAndAnchorParams, maxDistance *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorGate(&_ERC2771World.TransactOpts, params, maxDistance, networkNodeId)
}

// CreateAndAnchorGate is a paid mutator transaction binding the contract method 0xbb054d08.
//
// Solidity: function createAndAnchorGate((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 maxDistance, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndAnchorGate(params CreateAndAnchorParams, maxDistance *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorGate(&_ERC2771World.TransactOpts, params, maxDistance, networkNodeId)
}

// CreateAndAnchorNetworkNode is a paid mutator transaction binding the contract method 0xf6b45b12.
//
// Solidity: function createAndAnchorNetworkNode((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, (uint256,uint256) fuelParams, uint256 maxEnergyCapacity, uint256 currentProduction) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndAnchorNetworkNode(opts *bind.TransactOpts, params CreateAndAnchorParams, fuelParams FuelParams, maxEnergyCapacity *big.Int, currentProduction *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndAnchorNetworkNode", params, fuelParams, maxEnergyCapacity, currentProduction)
}

// CreateAndAnchorNetworkNode is a paid mutator transaction binding the contract method 0xf6b45b12.
//
// Solidity: function createAndAnchorNetworkNode((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, (uint256,uint256) fuelParams, uint256 maxEnergyCapacity, uint256 currentProduction) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndAnchorNetworkNode(params CreateAndAnchorParams, fuelParams FuelParams, maxEnergyCapacity *big.Int, currentProduction *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorNetworkNode(&_ERC2771World.TransactOpts, params, fuelParams, maxEnergyCapacity, currentProduction)
}

// CreateAndAnchorNetworkNode is a paid mutator transaction binding the contract method 0xf6b45b12.
//
// Solidity: function createAndAnchorNetworkNode((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, (uint256,uint256) fuelParams, uint256 maxEnergyCapacity, uint256 currentProduction) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndAnchorNetworkNode(params CreateAndAnchorParams, fuelParams FuelParams, maxEnergyCapacity *big.Int, currentProduction *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorNetworkNode(&_ERC2771World.TransactOpts, params, fuelParams, maxEnergyCapacity, currentProduction)
}

// CreateAndAnchorStorageUnit is a paid mutator transaction binding the contract method 0xbf451df6.
//
// Solidity: function createAndAnchorStorageUnit((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 capacity, uint256 ephemeralCapacity, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndAnchorStorageUnit(opts *bind.TransactOpts, params CreateAndAnchorParams, capacity *big.Int, ephemeralCapacity *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndAnchorStorageUnit", params, capacity, ephemeralCapacity, networkNodeId)
}

// CreateAndAnchorStorageUnit is a paid mutator transaction binding the contract method 0xbf451df6.
//
// Solidity: function createAndAnchorStorageUnit((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 capacity, uint256 ephemeralCapacity, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndAnchorStorageUnit(params CreateAndAnchorParams, capacity *big.Int, ephemeralCapacity *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorStorageUnit(&_ERC2771World.TransactOpts, params, capacity, ephemeralCapacity, networkNodeId)
}

// CreateAndAnchorStorageUnit is a paid mutator transaction binding the contract method 0xbf451df6.
//
// Solidity: function createAndAnchorStorageUnit((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 capacity, uint256 ephemeralCapacity, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndAnchorStorageUnit(params CreateAndAnchorParams, capacity *big.Int, ephemeralCapacity *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorStorageUnit(&_ERC2771World.TransactOpts, params, capacity, ephemeralCapacity, networkNodeId)
}

// CreateAndAnchorTurret is a paid mutator transaction binding the contract method 0xb209af27.
//
// Solidity: function createAndAnchorTurret((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndAnchorTurret(opts *bind.TransactOpts, params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndAnchorTurret", params, networkNodeId)
}

// CreateAndAnchorTurret is a paid mutator transaction binding the contract method 0xb209af27.
//
// Solidity: function createAndAnchorTurret((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndAnchorTurret(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorTurret(&_ERC2771World.TransactOpts, params, networkNodeId)
}

// CreateAndAnchorTurret is a paid mutator transaction binding the contract method 0xb209af27.
//
// Solidity: function createAndAnchorTurret((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndAnchorTurret(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorTurret(&_ERC2771World.TransactOpts, params, networkNodeId)
}

// CreateAndDepositEphemeral is a paid mutator transaction binding the contract method 0x0f33de24.
//
// Solidity: function createAndDepositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndDepositEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndDepositEphemeral", smartObjectId, ephemeralOwner, items)
}

// CreateAndDepositEphemeral is a paid mutator transaction binding the contract method 0x0f33de24.
//
// Solidity: function createAndDepositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndDepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndDepositEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// CreateAndDepositEphemeral is a paid mutator transaction binding the contract method 0x0f33de24.
//
// Solidity: function createAndDepositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndDepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndDepositEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// CreateAndDepositInventory is a paid mutator transaction binding the contract method 0x087ad59f.
//
// Solidity: function createAndDepositInventory(uint256 smartObjectId, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndDepositInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndDepositInventory", smartObjectId, items)
}

// CreateAndDepositInventory is a paid mutator transaction binding the contract method 0x087ad59f.
//
// Solidity: function createAndDepositInventory(uint256 smartObjectId, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndDepositInventory(smartObjectId *big.Int, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndDepositInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// CreateAndDepositInventory is a paid mutator transaction binding the contract method 0x087ad59f.
//
// Solidity: function createAndDepositInventory(uint256 smartObjectId, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndDepositInventory(smartObjectId *big.Int, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndDepositInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// CreateAssembly is a paid mutator transaction binding the contract method 0xc9f4c7c2.
//
// Solidity: function createAssembly(uint256 smartObjectId, string assemblyType, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAssembly(opts *bind.TransactOpts, smartObjectId *big.Int, assemblyType string, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAssembly", smartObjectId, assemblyType, entityRecordParams)
}

// CreateAssembly is a paid mutator transaction binding the contract method 0xc9f4c7c2.
//
// Solidity: function createAssembly(uint256 smartObjectId, string assemblyType, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAssembly(smartObjectId *big.Int, assemblyType string, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAssembly(&_ERC2771World.TransactOpts, smartObjectId, assemblyType, entityRecordParams)
}

// CreateAssembly is a paid mutator transaction binding the contract method 0xc9f4c7c2.
//
// Solidity: function createAssembly(uint256 smartObjectId, string assemblyType, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAssembly(smartObjectId *big.Int, assemblyType string, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAssembly(&_ERC2771World.TransactOpts, smartObjectId, assemblyType, entityRecordParams)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x6d0030ca.
//
// Solidity: function createCharacter(uint256 smartObjectId, address owner, uint256 tribeId, (bytes32,uint256,uint256,uint256) entityRecordParams, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateCharacter(opts *bind.TransactOpts, smartObjectId *big.Int, owner common.Address, tribeId *big.Int, entityRecordParams EntityRecordParams, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createCharacter", smartObjectId, owner, tribeId, entityRecordParams, entityRecordMetadata)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x6d0030ca.
//
// Solidity: function createCharacter(uint256 smartObjectId, address owner, uint256 tribeId, (bytes32,uint256,uint256,uint256) entityRecordParams, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771World *ERC2771WorldSession) CreateCharacter(smartObjectId *big.Int, owner common.Address, tribeId *big.Int, entityRecordParams EntityRecordParams, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateCharacter(&_ERC2771World.TransactOpts, smartObjectId, owner, tribeId, entityRecordParams, entityRecordMetadata)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x6d0030ca.
//
// Solidity: function createCharacter(uint256 smartObjectId, address owner, uint256 tribeId, (bytes32,uint256,uint256,uint256) entityRecordParams, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateCharacter(smartObjectId *big.Int, owner common.Address, tribeId *big.Int, entityRecordParams EntityRecordParams, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateCharacter(&_ERC2771World.TransactOpts, smartObjectId, owner, tribeId, entityRecordParams, entityRecordMetadata)
}

// CreateDeployable is a paid mutator transaction binding the contract method 0xcadee21c.
//
// Solidity: function createDeployable(uint256 smartObjectId, address owner) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateDeployable(opts *bind.TransactOpts, smartObjectId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createDeployable", smartObjectId, owner)
}

// CreateDeployable is a paid mutator transaction binding the contract method 0xcadee21c.
//
// Solidity: function createDeployable(uint256 smartObjectId, address owner) returns()
func (_ERC2771World *ERC2771WorldSession) CreateDeployable(smartObjectId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateDeployable(&_ERC2771World.TransactOpts, smartObjectId, owner)
}

// CreateDeployable is a paid mutator transaction binding the contract method 0xcadee21c.
//
// Solidity: function createDeployable(uint256 smartObjectId, address owner) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateDeployable(smartObjectId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateDeployable(&_ERC2771World.TransactOpts, smartObjectId, owner)
}

// CreateMetadata is a paid mutator transaction binding the contract method 0x21823a2a.
//
// Solidity: function createMetadata(uint256 smartObjectId, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateMetadata(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createMetadata", smartObjectId, entityRecordMetadata)
}

// CreateMetadata is a paid mutator transaction binding the contract method 0x21823a2a.
//
// Solidity: function createMetadata(uint256 smartObjectId, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771World *ERC2771WorldSession) CreateMetadata(smartObjectId *big.Int, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateMetadata(&_ERC2771World.TransactOpts, smartObjectId, entityRecordMetadata)
}

// CreateMetadata is a paid mutator transaction binding the contract method 0x21823a2a.
//
// Solidity: function createMetadata(uint256 smartObjectId, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateMetadata(smartObjectId *big.Int, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateMetadata(&_ERC2771World.TransactOpts, smartObjectId, entityRecordMetadata)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xb4b891f7.
//
// Solidity: function createRecord(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateRecord(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createRecord", smartObjectId, entityRecordParams)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xb4b891f7.
//
// Solidity: function createRecord(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771World *ERC2771WorldSession) CreateRecord(smartObjectId *big.Int, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateRecord(&_ERC2771World.TransactOpts, smartObjectId, entityRecordParams)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xb4b891f7.
//
// Solidity: function createRecord(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateRecord(smartObjectId *big.Int, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateRecord(&_ERC2771World.TransactOpts, smartObjectId, entityRecordParams)
}

// CrossTransferToEphemeral is a paid mutator transaction binding the contract method 0xab2aea53.
//
// Solidity: function crossTransferToEphemeral(uint256 smartObjectId, address fromEphemeralOwner, address toEphemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) CrossTransferToEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, fromEphemeralOwner common.Address, toEphemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "crossTransferToEphemeral", smartObjectId, fromEphemeralOwner, toEphemeralOwner, items)
}

// CrossTransferToEphemeral is a paid mutator transaction binding the contract method 0xab2aea53.
//
// Solidity: function crossTransferToEphemeral(uint256 smartObjectId, address fromEphemeralOwner, address toEphemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) CrossTransferToEphemeral(smartObjectId *big.Int, fromEphemeralOwner common.Address, toEphemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CrossTransferToEphemeral(&_ERC2771World.TransactOpts, smartObjectId, fromEphemeralOwner, toEphemeralOwner, items)
}

// CrossTransferToEphemeral is a paid mutator transaction binding the contract method 0xab2aea53.
//
// Solidity: function crossTransferToEphemeral(uint256 smartObjectId, address fromEphemeralOwner, address toEphemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CrossTransferToEphemeral(smartObjectId *big.Int, fromEphemeralOwner common.Address, toEphemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.CrossTransferToEphemeral(&_ERC2771World.TransactOpts, smartObjectId, fromEphemeralOwner, toEphemeralOwner, items)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_ERC2771World *ERC2771WorldTransactor) DeleteRecord(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "deleteRecord", tableId, keyTuple)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_ERC2771World *ERC2771WorldSession) DeleteRecord(tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.DeleteRecord(&_ERC2771World.TransactOpts, tableId, keyTuple)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) DeleteRecord(tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.DeleteRecord(&_ERC2771World.TransactOpts, tableId, keyTuple)
}

// DepositEphemeral is a paid mutator transaction binding the contract method 0xf3688464.
//
// Solidity: function depositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) DepositEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "depositEphemeral", smartObjectId, ephemeralOwner, items)
}

// DepositEphemeral is a paid mutator transaction binding the contract method 0xf3688464.
//
// Solidity: function depositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) DepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// DepositEphemeral is a paid mutator transaction binding the contract method 0xf3688464.
//
// Solidity: function depositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) DepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// DepositFuel is a paid mutator transaction binding the contract method 0x752813db.
//
// Solidity: function depositFuel(uint256 smartObjectId, uint256 fuelSmartObjectId, uint256 fuelAmount) returns()
func (_ERC2771World *ERC2771WorldTransactor) DepositFuel(opts *bind.TransactOpts, smartObjectId *big.Int, fuelSmartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "depositFuel", smartObjectId, fuelSmartObjectId, fuelAmount)
}

// DepositFuel is a paid mutator transaction binding the contract method 0x752813db.
//
// Solidity: function depositFuel(uint256 smartObjectId, uint256 fuelSmartObjectId, uint256 fuelAmount) returns()
func (_ERC2771World *ERC2771WorldSession) DepositFuel(smartObjectId *big.Int, fuelSmartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositFuel(&_ERC2771World.TransactOpts, smartObjectId, fuelSmartObjectId, fuelAmount)
}

// DepositFuel is a paid mutator transaction binding the contract method 0x752813db.
//
// Solidity: function depositFuel(uint256 smartObjectId, uint256 fuelSmartObjectId, uint256 fuelAmount) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) DepositFuel(smartObjectId *big.Int, fuelSmartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositFuel(&_ERC2771World.TransactOpts, smartObjectId, fuelSmartObjectId, fuelAmount)
}

// DepositInventory is a paid mutator transaction binding the contract method 0x1931d4e5.
//
// Solidity: function depositInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) DepositInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "depositInventory", smartObjectId, items)
}

// DepositInventory is a paid mutator transaction binding the contract method 0x1931d4e5.
//
// Solidity: function depositInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) DepositInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// DepositInventory is a paid mutator transaction binding the contract method 0x1931d4e5.
//
// Solidity: function depositInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) DepositInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactor) DestroyDeployable(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "destroyDeployable", smartObjectId)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldSession) DestroyDeployable(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.DestroyDeployable(&_ERC2771World.TransactOpts, smartObjectId)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) DestroyDeployable(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.DestroyDeployable(&_ERC2771World.TransactOpts, smartObjectId)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771World *ERC2771WorldTransactor) GrantAccess(opts *bind.TransactOpts, resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "grantAccess", resourceId, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771World *ERC2771WorldSession) GrantAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.GrantAccess(&_ERC2771World.TransactOpts, resourceId, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) GrantAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.GrantAccess(&_ERC2771World.TransactOpts, resourceId, grantee)
}

// InProximity is a paid mutator transaction binding the contract method 0x5588377c.
//
// Solidity: function inProximity(uint256 smartObjectId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldTransactor) InProximity(opts *bind.TransactOpts, smartObjectId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "inProximity", smartObjectId, priorityQueue, turret, turretTarget)
}

// InProximity is a paid mutator transaction binding the contract method 0x5588377c.
//
// Solidity: function inProximity(uint256 smartObjectId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldSession) InProximity(smartObjectId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771World.Contract.InProximity(&_ERC2771World.TransactOpts, smartObjectId, priorityQueue, turret, turretTarget)
}

// InProximity is a paid mutator transaction binding the contract method 0x5588377c.
//
// Solidity: function inProximity(uint256 smartObjectId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldTransactorSession) InProximity(smartObjectId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771World.Contract.InProximity(&_ERC2771World.TransactOpts, smartObjectId, priorityQueue, turret, turretTarget)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_ERC2771World *ERC2771WorldTransactor) Initialize(opts *bind.TransactOpts, initModule common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "initialize", initModule)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_ERC2771World *ERC2771WorldSession) Initialize(initModule common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.Initialize(&_ERC2771World.TransactOpts, initModule)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) Initialize(initModule common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.Initialize(&_ERC2771World.TransactOpts, initModule)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_ERC2771World *ERC2771WorldTransactor) InstallModule(opts *bind.TransactOpts, module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "installModule", module, encodedArgs)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_ERC2771World *ERC2771WorldSession) InstallModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.InstallModule(&_ERC2771World.TransactOpts, module, encodedArgs)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) InstallModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.InstallModule(&_ERC2771World.TransactOpts, module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_ERC2771World *ERC2771WorldTransactor) InstallRootModule(opts *bind.TransactOpts, module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "installRootModule", module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_ERC2771World *ERC2771WorldSession) InstallRootModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.InstallRootModule(&_ERC2771World.TransactOpts, module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) InstallRootModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.InstallRootModule(&_ERC2771World.TransactOpts, module, encodedArgs)
}

// LinkGates is a paid mutator transaction binding the contract method 0x1caa06ce.
//
// Solidity: function linkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldTransactor) LinkGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "linkGates", sourceGateId, destinationGateId)
}

// LinkGates is a paid mutator transaction binding the contract method 0x1caa06ce.
//
// Solidity: function linkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldSession) LinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.LinkGates(&_ERC2771World.TransactOpts, sourceGateId, destinationGateId)
}

// LinkGates is a paid mutator transaction binding the contract method 0x1caa06ce.
//
// Solidity: function linkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) LinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.LinkGates(&_ERC2771World.TransactOpts, sourceGateId, destinationGateId)
}

// OnAssemblyOffline is a paid mutator transaction binding the contract method 0x725235c7.
//
// Solidity: function onAssemblyOffline(uint256 networkNodeId, uint256 assemblyId) returns()
func (_ERC2771World *ERC2771WorldTransactor) OnAssemblyOffline(opts *bind.TransactOpts, networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "onAssemblyOffline", networkNodeId, assemblyId)
}

// OnAssemblyOffline is a paid mutator transaction binding the contract method 0x725235c7.
//
// Solidity: function onAssemblyOffline(uint256 networkNodeId, uint256 assemblyId) returns()
func (_ERC2771World *ERC2771WorldSession) OnAssemblyOffline(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.OnAssemblyOffline(&_ERC2771World.TransactOpts, networkNodeId, assemblyId)
}

// OnAssemblyOffline is a paid mutator transaction binding the contract method 0x725235c7.
//
// Solidity: function onAssemblyOffline(uint256 networkNodeId, uint256 assemblyId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) OnAssemblyOffline(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.OnAssemblyOffline(&_ERC2771World.TransactOpts, networkNodeId, assemblyId)
}

// OnAssemblyOnline is a paid mutator transaction binding the contract method 0xe45f15fb.
//
// Solidity: function onAssemblyOnline(uint256 networkNodeId, uint256 assemblyId) returns()
func (_ERC2771World *ERC2771WorldTransactor) OnAssemblyOnline(opts *bind.TransactOpts, networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "onAssemblyOnline", networkNodeId, assemblyId)
}

// OnAssemblyOnline is a paid mutator transaction binding the contract method 0xe45f15fb.
//
// Solidity: function onAssemblyOnline(uint256 networkNodeId, uint256 assemblyId) returns()
func (_ERC2771World *ERC2771WorldSession) OnAssemblyOnline(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.OnAssemblyOnline(&_ERC2771World.TransactOpts, networkNodeId, assemblyId)
}

// OnAssemblyOnline is a paid mutator transaction binding the contract method 0xe45f15fb.
//
// Solidity: function onAssemblyOnline(uint256 networkNodeId, uint256 assemblyId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) OnAssemblyOnline(networkNodeId *big.Int, assemblyId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.OnAssemblyOnline(&_ERC2771World.TransactOpts, networkNodeId, assemblyId)
}

// OnNodeOffline is a paid mutator transaction binding the contract method 0x44eefc3f.
//
// Solidity: function onNodeOffline(uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactor) OnNodeOffline(opts *bind.TransactOpts, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "onNodeOffline", networkNodeId)
}

// OnNodeOffline is a paid mutator transaction binding the contract method 0x44eefc3f.
//
// Solidity: function onNodeOffline(uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldSession) OnNodeOffline(networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.OnNodeOffline(&_ERC2771World.TransactOpts, networkNodeId)
}

// OnNodeOffline is a paid mutator transaction binding the contract method 0x44eefc3f.
//
// Solidity: function onNodeOffline(uint256 networkNodeId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) OnNodeOffline(networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.OnNodeOffline(&_ERC2771World.TransactOpts, networkNodeId)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_ERC2771World *ERC2771WorldTransactor) PopFromDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "popFromDynamicField", tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_ERC2771World *ERC2771WorldSession) PopFromDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.PopFromDynamicField(&_ERC2771World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) PopFromDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.PopFromDynamicField(&_ERC2771World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_ERC2771World *ERC2771WorldTransactor) PushToDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "pushToDynamicField", tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_ERC2771World *ERC2771WorldSession) PushToDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.PushToDynamicField(&_ERC2771World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) PushToDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.PushToDynamicField(&_ERC2771World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterDelegation(opts *bind.TransactOpts, delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerDelegation", delegatee, delegationControlId, initCallData)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterDelegation(delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterDelegation(&_ERC2771World.TransactOpts, delegatee, delegationControlId, initCallData)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterDelegation(delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterDelegation(&_ERC2771World.TransactOpts, delegatee, delegationControlId, initCallData)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771World *ERC2771WorldTransactor) RegisterFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerFunctionSelector", systemId, systemFunctionSignature)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771World *ERC2771WorldSession) RegisterFunctionSelector(systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterFunctionSelector(&_ERC2771World.TransactOpts, systemId, systemFunctionSignature)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterFunctionSelector(systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterFunctionSelector(&_ERC2771World.TransactOpts, systemId, systemFunctionSignature)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterNamespace(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerNamespace", namespaceId)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterNamespace(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterNamespace(&_ERC2771World.TransactOpts, namespaceId)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterNamespace(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterNamespace(&_ERC2771World.TransactOpts, namespaceId)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterNamespaceDelegation(opts *bind.TransactOpts, namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerNamespaceDelegation", namespaceId, delegationControlId, initCallData)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterNamespaceDelegation(namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterNamespaceDelegation(&_ERC2771World.TransactOpts, namespaceId, delegationControlId, initCallData)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterNamespaceDelegation(namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterNamespaceDelegation(&_ERC2771World.TransactOpts, namespaceId, delegationControlId, initCallData)
}

// RegisterNetworkNodeClass is a paid mutator transaction binding the contract method 0x60ad4ad6.
//
// Solidity: function registerNetworkNodeClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterNetworkNodeClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerNetworkNodeClass", typeId, volume)
}

// RegisterNetworkNodeClass is a paid mutator transaction binding the contract method 0x60ad4ad6.
//
// Solidity: function registerNetworkNodeClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterNetworkNodeClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterNetworkNodeClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterNetworkNodeClass is a paid mutator transaction binding the contract method 0x60ad4ad6.
//
// Solidity: function registerNetworkNodeClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterNetworkNodeClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterNetworkNodeClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x6548a90a.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771World *ERC2771WorldTransactor) RegisterRootFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, worldFunctionSignature string, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerRootFunctionSelector", systemId, worldFunctionSignature, systemFunctionSignature)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x6548a90a.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771World *ERC2771WorldSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterRootFunctionSelector(&_ERC2771World.TransactOpts, systemId, worldFunctionSignature, systemFunctionSignature)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x6548a90a.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterRootFunctionSelector(&_ERC2771World.TransactOpts, systemId, worldFunctionSignature, systemFunctionSignature)
}

// RegisterSmartCharacterClass is a paid mutator transaction binding the contract method 0x3664e851.
//
// Solidity: function registerSmartCharacterClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterSmartCharacterClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerSmartCharacterClass", typeId, volume)
}

// RegisterSmartCharacterClass is a paid mutator transaction binding the contract method 0x3664e851.
//
// Solidity: function registerSmartCharacterClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterSmartCharacterClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSmartCharacterClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterSmartCharacterClass is a paid mutator transaction binding the contract method 0x3664e851.
//
// Solidity: function registerSmartCharacterClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterSmartCharacterClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSmartCharacterClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterSmartGateClass is a paid mutator transaction binding the contract method 0x63d71589.
//
// Solidity: function registerSmartGateClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterSmartGateClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerSmartGateClass", typeId, volume)
}

// RegisterSmartGateClass is a paid mutator transaction binding the contract method 0x63d71589.
//
// Solidity: function registerSmartGateClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterSmartGateClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSmartGateClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterSmartGateClass is a paid mutator transaction binding the contract method 0x63d71589.
//
// Solidity: function registerSmartGateClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterSmartGateClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSmartGateClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterSmartStorageUnitClass is a paid mutator transaction binding the contract method 0x1f5493d1.
//
// Solidity: function registerSmartStorageUnitClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterSmartStorageUnitClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerSmartStorageUnitClass", typeId, volume)
}

// RegisterSmartStorageUnitClass is a paid mutator transaction binding the contract method 0x1f5493d1.
//
// Solidity: function registerSmartStorageUnitClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterSmartStorageUnitClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSmartStorageUnitClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterSmartStorageUnitClass is a paid mutator transaction binding the contract method 0x1f5493d1.
//
// Solidity: function registerSmartStorageUnitClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterSmartStorageUnitClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSmartStorageUnitClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterSmartTurretClass is a paid mutator transaction binding the contract method 0x2cf57627.
//
// Solidity: function registerSmartTurretClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterSmartTurretClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerSmartTurretClass", typeId, volume)
}

// RegisterSmartTurretClass is a paid mutator transaction binding the contract method 0x2cf57627.
//
// Solidity: function registerSmartTurretClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterSmartTurretClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSmartTurretClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterSmartTurretClass is a paid mutator transaction binding the contract method 0x2cf57627.
//
// Solidity: function registerSmartTurretClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterSmartTurretClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSmartTurretClass(&_ERC2771World.TransactOpts, typeId, volume)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterStoreHook(opts *bind.TransactOpts, tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerStoreHook", tableId, hookAddress, enabledHooksBitmap)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterStoreHook(tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterStoreHook(&_ERC2771World.TransactOpts, tableId, hookAddress, enabledHooksBitmap)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterStoreHook(tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterStoreHook(&_ERC2771World.TransactOpts, tableId, hookAddress, enabledHooksBitmap)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterSystem(opts *bind.TransactOpts, systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerSystem", systemId, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterSystem(systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSystem(&_ERC2771World.TransactOpts, systemId, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterSystem(systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSystem(&_ERC2771World.TransactOpts, systemId, system, publicAccess)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterSystemHook(opts *bind.TransactOpts, systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerSystemHook", systemId, hookAddress, enabledHooksBitmap)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterSystemHook(systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSystemHook(&_ERC2771World.TransactOpts, systemId, hookAddress, enabledHooksBitmap)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterSystemHook(systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterSystemHook(&_ERC2771World.TransactOpts, systemId, hookAddress, enabledHooksBitmap)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterTable(opts *bind.TransactOpts, tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerTable", tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterTable(tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterTable(&_ERC2771World.TransactOpts, tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterTable(tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterTable(&_ERC2771World.TransactOpts, tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RemoveCharacter is a paid mutator transaction binding the contract method 0x834f6531.
//
// Solidity: function removeCharacter(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactor) RemoveCharacter(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "removeCharacter", smartObjectId)
}

// RemoveCharacter is a paid mutator transaction binding the contract method 0x834f6531.
//
// Solidity: function removeCharacter(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldSession) RemoveCharacter(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RemoveCharacter(&_ERC2771World.TransactOpts, smartObjectId)
}

// RemoveCharacter is a paid mutator transaction binding the contract method 0x834f6531.
//
// Solidity: function removeCharacter(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RemoveCharacter(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RemoveCharacter(&_ERC2771World.TransactOpts, smartObjectId)
}

// RemoveItemFromInventory is a paid mutator transaction binding the contract method 0x8a38b387.
//
// Solidity: function removeItemFromInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771World *ERC2771WorldTransactor) RemoveItemFromInventory(opts *bind.TransactOpts, inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "removeItemFromInventory", inventoryObjectId, itemObjectId, quantity)
}

// RemoveItemFromInventory is a paid mutator transaction binding the contract method 0x8a38b387.
//
// Solidity: function removeItemFromInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771World *ERC2771WorldSession) RemoveItemFromInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RemoveItemFromInventory(&_ERC2771World.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// RemoveItemFromInventory is a paid mutator transaction binding the contract method 0x8a38b387.
//
// Solidity: function removeItemFromInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RemoveItemFromInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RemoveItemFromInventory(&_ERC2771World.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x0058bbfd.
//
// Solidity: function removeOwner(uint256 smartObjectId, address from) returns()
func (_ERC2771World *ERC2771WorldTransactor) RemoveOwner(opts *bind.TransactOpts, smartObjectId *big.Int, from common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "removeOwner", smartObjectId, from)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x0058bbfd.
//
// Solidity: function removeOwner(uint256 smartObjectId, address from) returns()
func (_ERC2771World *ERC2771WorldSession) RemoveOwner(smartObjectId *big.Int, from common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.RemoveOwner(&_ERC2771World.TransactOpts, smartObjectId, from)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x0058bbfd.
//
// Solidity: function removeOwner(uint256 smartObjectId, address from) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RemoveOwner(smartObjectId *big.Int, from common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.RemoveOwner(&_ERC2771World.TransactOpts, smartObjectId, from)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_ERC2771World *ERC2771WorldTransactor) RenounceOwnership(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "renounceOwnership", namespaceId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_ERC2771World *ERC2771WorldSession) RenounceOwnership(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RenounceOwnership(&_ERC2771World.TransactOpts, namespaceId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RenounceOwnership(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RenounceOwnership(&_ERC2771World.TransactOpts, namespaceId)
}

// ReportKill is a paid mutator transaction binding the contract method 0xda3293e1.
//
// Solidity: function reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailData) returns()
func (_ERC2771World *ERC2771WorldTransactor) ReportKill(opts *bind.TransactOpts, killMailId *big.Int, killMailData KillMailData) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "reportKill", killMailId, killMailData)
}

// ReportKill is a paid mutator transaction binding the contract method 0xda3293e1.
//
// Solidity: function reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailData) returns()
func (_ERC2771World *ERC2771WorldSession) ReportKill(killMailId *big.Int, killMailData KillMailData) (*types.Transaction, error) {
	return _ERC2771World.Contract.ReportKill(&_ERC2771World.TransactOpts, killMailId, killMailData)
}

// ReportKill is a paid mutator transaction binding the contract method 0xda3293e1.
//
// Solidity: function reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailData) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ReportKill(killMailId *big.Int, killMailData KillMailData) (*types.Transaction, error) {
	return _ERC2771World.Contract.ReportKill(&_ERC2771World.TransactOpts, killMailId, killMailData)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771World *ERC2771WorldTransactor) RevokeAccess(opts *bind.TransactOpts, resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "revokeAccess", resourceId, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771World *ERC2771WorldSession) RevokeAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.RevokeAccess(&_ERC2771World.TransactOpts, resourceId, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RevokeAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.RevokeAccess(&_ERC2771World.TransactOpts, resourceId, grantee)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 smartObjectId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771World *ERC2771WorldTransactor) SaveLocation(opts *bind.TransactOpts, smartObjectId *big.Int, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "saveLocation", smartObjectId, locationData)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 smartObjectId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771World *ERC2771WorldSession) SaveLocation(smartObjectId *big.Int, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771World.Contract.SaveLocation(&_ERC2771World.TransactOpts, smartObjectId, locationData)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 smartObjectId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SaveLocation(smartObjectId *big.Int, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771World.Contract.SaveLocation(&_ERC2771World.TransactOpts, smartObjectId, locationData)
}

// SetAssemblyType is a paid mutator transaction binding the contract method 0xf9914db3.
//
// Solidity: function setAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetAssemblyType(opts *bind.TransactOpts, smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setAssemblyType", smartObjectId, assemblyType)
}

// SetAssemblyType is a paid mutator transaction binding the contract method 0xf9914db3.
//
// Solidity: function setAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771World *ERC2771WorldSession) SetAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetAssemblyType(&_ERC2771World.TransactOpts, smartObjectId, assemblyType)
}

// SetAssemblyType is a paid mutator transaction binding the contract method 0xf9914db3.
//
// Solidity: function setAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetAssemblyType(&_ERC2771World.TransactOpts, smartObjectId, assemblyType)
}

// SetCapacity is a paid mutator transaction binding the contract method 0xdee2b058.
//
// Solidity: function setCapacity(uint256 smartObjectId, uint256 capacity) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, capacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setCapacity", smartObjectId, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0xdee2b058.
//
// Solidity: function setCapacity(uint256 smartObjectId, uint256 capacity) returns()
func (_ERC2771World *ERC2771WorldSession) SetCapacity(smartObjectId *big.Int, capacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetCapacity(&_ERC2771World.TransactOpts, smartObjectId, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0xdee2b058.
//
// Solidity: function setCapacity(uint256 smartObjectId, uint256 capacity) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetCapacity(smartObjectId *big.Int, capacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetCapacity(&_ERC2771World.TransactOpts, smartObjectId, capacity)
}

// SetCrossTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0x5f3be290.
//
// Solidity: function setCrossTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetCrossTransferToEphemeralAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setCrossTransferToEphemeralAccess", smartObjectId, accessAddress, isAllowed)
}

// SetCrossTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0x5f3be290.
//
// Solidity: function setCrossTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldSession) SetCrossTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetCrossTransferToEphemeralAccess(&_ERC2771World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetCrossTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0x5f3be290.
//
// Solidity: function setCrossTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetCrossTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetCrossTransferToEphemeralAccess(&_ERC2771World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 smartObjectId, string dappURL) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetDappURL(opts *bind.TransactOpts, smartObjectId *big.Int, dappURL string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setDappURL", smartObjectId, dappURL)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 smartObjectId, string dappURL) returns()
func (_ERC2771World *ERC2771WorldSession) SetDappURL(smartObjectId *big.Int, dappURL string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDappURL(&_ERC2771World.TransactOpts, smartObjectId, dappURL)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 smartObjectId, string dappURL) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetDappURL(smartObjectId *big.Int, dappURL string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDappURL(&_ERC2771World.TransactOpts, smartObjectId, dappURL)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 smartObjectId, string description) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetDescription(opts *bind.TransactOpts, smartObjectId *big.Int, description string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setDescription", smartObjectId, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 smartObjectId, string description) returns()
func (_ERC2771World *ERC2771WorldSession) SetDescription(smartObjectId *big.Int, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDescription(&_ERC2771World.TransactOpts, smartObjectId, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 smartObjectId, string description) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetDescription(smartObjectId *big.Int, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDescription(&_ERC2771World.TransactOpts, smartObjectId, description)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setDynamicField", tableId, keyTuple, dynamicFieldIndex, data)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_ERC2771World *ERC2771WorldSession) SetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDynamicField(&_ERC2771World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, data)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDynamicField(&_ERC2771World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, data)
}

// SetEphemeralCapacity is a paid mutator transaction binding the contract method 0xbc8229dc.
//
// Solidity: function setEphemeralCapacity(uint256 smartObjectId, uint256 ephemeralCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetEphemeralCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setEphemeralCapacity", smartObjectId, ephemeralCapacity)
}

// SetEphemeralCapacity is a paid mutator transaction binding the contract method 0xbc8229dc.
//
// Solidity: function setEphemeralCapacity(uint256 smartObjectId, uint256 ephemeralCapacity) returns()
func (_ERC2771World *ERC2771WorldSession) SetEphemeralCapacity(smartObjectId *big.Int, ephemeralCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetEphemeralCapacity(&_ERC2771World.TransactOpts, smartObjectId, ephemeralCapacity)
}

// SetEphemeralCapacity is a paid mutator transaction binding the contract method 0xbc8229dc.
//
// Solidity: function setEphemeralCapacity(uint256 smartObjectId, uint256 ephemeralCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetEphemeralCapacity(smartObjectId *big.Int, ephemeralCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetEphemeralCapacity(&_ERC2771World.TransactOpts, smartObjectId, ephemeralCapacity)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setField", tableId, keyTuple, fieldIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_ERC2771World *ERC2771WorldSession) SetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetField(&_ERC2771World.TransactOpts, tableId, keyTuple, fieldIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetField(&_ERC2771World.TransactOpts, tableId, keyTuple, fieldIndex, data)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetField0(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setField0", tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771World *ERC2771WorldSession) SetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetField0(&_ERC2771World.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetField0(&_ERC2771World.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 smartObjectId, uint256 fuelMaxCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetFuelMaxCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setFuelMaxCapacity", smartObjectId, fuelMaxCapacity)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 smartObjectId, uint256 fuelMaxCapacity) returns()
func (_ERC2771World *ERC2771WorldSession) SetFuelMaxCapacity(smartObjectId *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetFuelMaxCapacity(&_ERC2771World.TransactOpts, smartObjectId, fuelMaxCapacity)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 smartObjectId, uint256 fuelMaxCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetFuelMaxCapacity(smartObjectId *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetFuelMaxCapacity(&_ERC2771World.TransactOpts, smartObjectId, fuelMaxCapacity)
}

// SetFuelUnitVolume is a paid mutator transaction binding the contract method 0x20fa3423.
//
// Solidity: function setFuelUnitVolume(uint256 smartObjectId, uint256 fuelUnitVolume) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetFuelUnitVolume(opts *bind.TransactOpts, smartObjectId *big.Int, fuelUnitVolume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setFuelUnitVolume", smartObjectId, fuelUnitVolume)
}

// SetFuelUnitVolume is a paid mutator transaction binding the contract method 0x20fa3423.
//
// Solidity: function setFuelUnitVolume(uint256 smartObjectId, uint256 fuelUnitVolume) returns()
func (_ERC2771World *ERC2771WorldSession) SetFuelUnitVolume(smartObjectId *big.Int, fuelUnitVolume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetFuelUnitVolume(&_ERC2771World.TransactOpts, smartObjectId, fuelUnitVolume)
}

// SetFuelUnitVolume is a paid mutator transaction binding the contract method 0x20fa3423.
//
// Solidity: function setFuelUnitVolume(uint256 smartObjectId, uint256 fuelUnitVolume) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetFuelUnitVolume(smartObjectId *big.Int, fuelUnitVolume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetFuelUnitVolume(&_ERC2771World.TransactOpts, smartObjectId, fuelUnitVolume)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 smartObjectId, string name) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetName(opts *bind.TransactOpts, smartObjectId *big.Int, name string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setName", smartObjectId, name)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 smartObjectId, string name) returns()
func (_ERC2771World *ERC2771WorldSession) SetName(smartObjectId *big.Int, name string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetName(&_ERC2771World.TransactOpts, smartObjectId, name)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 smartObjectId, string name) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetName(smartObjectId *big.Int, name string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetName(&_ERC2771World.TransactOpts, smartObjectId, name)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetRecord(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setRecord", tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_ERC2771World *ERC2771WorldSession) SetRecord(tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetRecord(&_ERC2771World.TransactOpts, tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetRecord(tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetRecord(&_ERC2771World.TransactOpts, tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetStaticField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setStaticField", tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771World *ERC2771WorldSession) SetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetStaticField(&_ERC2771World.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetStaticField(&_ERC2771World.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetTransferFromEphemeralAccess is a paid mutator transaction binding the contract method 0x8d0ead31.
//
// Solidity: function setTransferFromEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetTransferFromEphemeralAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setTransferFromEphemeralAccess", smartObjectId, accessAddress, isAllowed)
}

// SetTransferFromEphemeralAccess is a paid mutator transaction binding the contract method 0x8d0ead31.
//
// Solidity: function setTransferFromEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldSession) SetTransferFromEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetTransferFromEphemeralAccess(&_ERC2771World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferFromEphemeralAccess is a paid mutator transaction binding the contract method 0x8d0ead31.
//
// Solidity: function setTransferFromEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetTransferFromEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetTransferFromEphemeralAccess(&_ERC2771World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0xd6873420.
//
// Solidity: function setTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetTransferToEphemeralAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setTransferToEphemeralAccess", smartObjectId, accessAddress, isAllowed)
}

// SetTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0xd6873420.
//
// Solidity: function setTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldSession) SetTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetTransferToEphemeralAccess(&_ERC2771World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0xd6873420.
//
// Solidity: function setTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetTransferToEphemeralAccess(&_ERC2771World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferToInventoryAccess is a paid mutator transaction binding the contract method 0x0a231135.
//
// Solidity: function setTransferToInventoryAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetTransferToInventoryAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setTransferToInventoryAccess", smartObjectId, accessAddress, isAllowed)
}

// SetTransferToInventoryAccess is a paid mutator transaction binding the contract method 0x0a231135.
//
// Solidity: function setTransferToInventoryAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldSession) SetTransferToInventoryAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetTransferToInventoryAccess(&_ERC2771World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferToInventoryAccess is a paid mutator transaction binding the contract method 0x0a231135.
//
// Solidity: function setTransferToInventoryAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetTransferToInventoryAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetTransferToInventoryAccess(&_ERC2771World.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_ERC2771World *ERC2771WorldTransactor) SpliceDynamicData(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "spliceDynamicData", tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_ERC2771World *ERC2771WorldSession) SpliceDynamicData(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SpliceDynamicData(&_ERC2771World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SpliceDynamicData(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SpliceDynamicData(&_ERC2771World.TransactOpts, tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_ERC2771World *ERC2771WorldTransactor) SpliceStaticData(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "spliceStaticData", tableId, keyTuple, start, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_ERC2771World *ERC2771WorldSession) SpliceStaticData(tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SpliceStaticData(&_ERC2771World.TransactOpts, tableId, keyTuple, start, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SpliceStaticData(tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.SpliceStaticData(&_ERC2771World.TransactOpts, tableId, keyTuple, start, data)
}

// StartBurn is a paid mutator transaction binding the contract method 0x46d647ab.
//
// Solidity: function startBurn(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactor) StartBurn(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "startBurn", smartObjectId)
}

// StartBurn is a paid mutator transaction binding the contract method 0x46d647ab.
//
// Solidity: function startBurn(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldSession) StartBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.StartBurn(&_ERC2771World.TransactOpts, smartObjectId)
}

// StartBurn is a paid mutator transaction binding the contract method 0x46d647ab.
//
// Solidity: function startBurn(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) StartBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.StartBurn(&_ERC2771World.TransactOpts, smartObjectId)
}

// StopBurn is a paid mutator transaction binding the contract method 0xe2487ae0.
//
// Solidity: function stopBurn(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactor) StopBurn(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "stopBurn", smartObjectId)
}

// StopBurn is a paid mutator transaction binding the contract method 0xe2487ae0.
//
// Solidity: function stopBurn(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldSession) StopBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.StopBurn(&_ERC2771World.TransactOpts, smartObjectId)
}

// StopBurn is a paid mutator transaction binding the contract method 0xe2487ae0.
//
// Solidity: function stopBurn(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) StopBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.StopBurn(&_ERC2771World.TransactOpts, smartObjectId)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_ERC2771World *ERC2771WorldTransactor) TransferBalanceToAddress(opts *bind.TransactOpts, fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "transferBalanceToAddress", fromNamespaceId, toAddress, amount)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_ERC2771World *ERC2771WorldSession) TransferBalanceToAddress(fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferBalanceToAddress(&_ERC2771World.TransactOpts, fromNamespaceId, toAddress, amount)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) TransferBalanceToAddress(fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferBalanceToAddress(&_ERC2771World.TransactOpts, fromNamespaceId, toAddress, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_ERC2771World *ERC2771WorldTransactor) TransferBalanceToNamespace(opts *bind.TransactOpts, fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "transferBalanceToNamespace", fromNamespaceId, toNamespaceId, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_ERC2771World *ERC2771WorldSession) TransferBalanceToNamespace(fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferBalanceToNamespace(&_ERC2771World.TransactOpts, fromNamespaceId, toNamespaceId, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) TransferBalanceToNamespace(fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferBalanceToNamespace(&_ERC2771World.TransactOpts, fromNamespaceId, toNamespaceId, amount)
}

// TransferFromEphemeral is a paid mutator transaction binding the contract method 0x6f29e3b1.
//
// Solidity: function transferFromEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) TransferFromEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "transferFromEphemeral", smartObjectId, ephemeralOwner, items)
}

// TransferFromEphemeral is a paid mutator transaction binding the contract method 0x6f29e3b1.
//
// Solidity: function transferFromEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) TransferFromEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferFromEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// TransferFromEphemeral is a paid mutator transaction binding the contract method 0x6f29e3b1.
//
// Solidity: function transferFromEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) TransferFromEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferFromEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_ERC2771World *ERC2771WorldTransactor) TransferOwnership(opts *bind.TransactOpts, namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "transferOwnership", namespaceId, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_ERC2771World *ERC2771WorldSession) TransferOwnership(namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferOwnership(&_ERC2771World.TransactOpts, namespaceId, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) TransferOwnership(namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferOwnership(&_ERC2771World.TransactOpts, namespaceId, newOwner)
}

// TransferToEphemeral is a paid mutator transaction binding the contract method 0x63821267.
//
// Solidity: function transferToEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) TransferToEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "transferToEphemeral", smartObjectId, ephemeralOwner, items)
}

// TransferToEphemeral is a paid mutator transaction binding the contract method 0x63821267.
//
// Solidity: function transferToEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) TransferToEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferToEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// TransferToEphemeral is a paid mutator transaction binding the contract method 0x63821267.
//
// Solidity: function transferToEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) TransferToEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferToEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// TransferToInventory is a paid mutator transaction binding the contract method 0x4ce8f5ce.
//
// Solidity: function transferToInventory(uint256 smartObjectId, uint256 toObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) TransferToInventory(opts *bind.TransactOpts, smartObjectId *big.Int, toObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "transferToInventory", smartObjectId, toObjectId, items)
}

// TransferToInventory is a paid mutator transaction binding the contract method 0x4ce8f5ce.
//
// Solidity: function transferToInventory(uint256 smartObjectId, uint256 toObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) TransferToInventory(smartObjectId *big.Int, toObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferToInventory(&_ERC2771World.TransactOpts, smartObjectId, toObjectId, items)
}

// TransferToInventory is a paid mutator transaction binding the contract method 0x4ce8f5ce.
//
// Solidity: function transferToInventory(uint256 smartObjectId, uint256 toObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) TransferToInventory(smartObjectId *big.Int, toObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.TransferToInventory(&_ERC2771World.TransactOpts, smartObjectId, toObjectId, items)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactor) Unanchor(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "unanchor", smartObjectId)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldSession) Unanchor(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.Unanchor(&_ERC2771World.TransactOpts, smartObjectId)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) Unanchor(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.Unanchor(&_ERC2771World.TransactOpts, smartObjectId)
}

// UnlinkGates is a paid mutator transaction binding the contract method 0x6a57ebde.
//
// Solidity: function unlinkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldTransactor) UnlinkGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "unlinkGates", sourceGateId, destinationGateId)
}

// UnlinkGates is a paid mutator transaction binding the contract method 0x6a57ebde.
//
// Solidity: function unlinkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldSession) UnlinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnlinkGates(&_ERC2771World.TransactOpts, sourceGateId, destinationGateId)
}

// UnlinkGates is a paid mutator transaction binding the contract method 0x6a57ebde.
//
// Solidity: function unlinkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UnlinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnlinkGates(&_ERC2771World.TransactOpts, sourceGateId, destinationGateId)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_ERC2771World *ERC2771WorldTransactor) UnregisterDelegation(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "unregisterDelegation", delegatee)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_ERC2771World *ERC2771WorldSession) UnregisterDelegation(delegatee common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnregisterDelegation(&_ERC2771World.TransactOpts, delegatee)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UnregisterDelegation(delegatee common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnregisterDelegation(&_ERC2771World.TransactOpts, delegatee)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_ERC2771World *ERC2771WorldTransactor) UnregisterNamespaceDelegation(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "unregisterNamespaceDelegation", namespaceId)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_ERC2771World *ERC2771WorldSession) UnregisterNamespaceDelegation(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnregisterNamespaceDelegation(&_ERC2771World.TransactOpts, namespaceId)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UnregisterNamespaceDelegation(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnregisterNamespaceDelegation(&_ERC2771World.TransactOpts, namespaceId)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_ERC2771World *ERC2771WorldTransactor) UnregisterStoreHook(opts *bind.TransactOpts, tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "unregisterStoreHook", tableId, hookAddress)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_ERC2771World *ERC2771WorldSession) UnregisterStoreHook(tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnregisterStoreHook(&_ERC2771World.TransactOpts, tableId, hookAddress)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UnregisterStoreHook(tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnregisterStoreHook(&_ERC2771World.TransactOpts, tableId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_ERC2771World *ERC2771WorldTransactor) UnregisterSystemHook(opts *bind.TransactOpts, systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "unregisterSystemHook", systemId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_ERC2771World *ERC2771WorldSession) UnregisterSystemHook(systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnregisterSystemHook(&_ERC2771World.TransactOpts, systemId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UnregisterSystemHook(systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnregisterSystemHook(&_ERC2771World.TransactOpts, systemId, hookAddress)
}

// UpdateAssemblyType is a paid mutator transaction binding the contract method 0xbcb029b2.
//
// Solidity: function updateAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771World *ERC2771WorldTransactor) UpdateAssemblyType(opts *bind.TransactOpts, smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "updateAssemblyType", smartObjectId, assemblyType)
}

// UpdateAssemblyType is a paid mutator transaction binding the contract method 0xbcb029b2.
//
// Solidity: function updateAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771World *ERC2771WorldSession) UpdateAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateAssemblyType(&_ERC2771World.TransactOpts, smartObjectId, assemblyType)
}

// UpdateAssemblyType is a paid mutator transaction binding the contract method 0xbcb029b2.
//
// Solidity: function updateAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UpdateAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateAssemblyType(&_ERC2771World.TransactOpts, smartObjectId, assemblyType)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactor) UpdateFuel(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "updateFuel", smartObjectId)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldSession) UpdateFuel(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateFuel(&_ERC2771World.TransactOpts, smartObjectId)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 smartObjectId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UpdateFuel(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateFuel(&_ERC2771World.TransactOpts, smartObjectId)
}

// UpdateTribeId is a paid mutator transaction binding the contract method 0x16d51f3e.
//
// Solidity: function updateTribeId(uint256 smartObjectId, uint256 tribeId) returns()
func (_ERC2771World *ERC2771WorldTransactor) UpdateTribeId(opts *bind.TransactOpts, smartObjectId *big.Int, tribeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "updateTribeId", smartObjectId, tribeId)
}

// UpdateTribeId is a paid mutator transaction binding the contract method 0x16d51f3e.
//
// Solidity: function updateTribeId(uint256 smartObjectId, uint256 tribeId) returns()
func (_ERC2771World *ERC2771WorldSession) UpdateTribeId(smartObjectId *big.Int, tribeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateTribeId(&_ERC2771World.TransactOpts, smartObjectId, tribeId)
}

// UpdateTribeId is a paid mutator transaction binding the contract method 0x16d51f3e.
//
// Solidity: function updateTribeId(uint256 smartObjectId, uint256 tribeId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UpdateTribeId(smartObjectId *big.Int, tribeId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateTribeId(&_ERC2771World.TransactOpts, smartObjectId, tribeId)
}

// WithdrawEphemeral is a paid mutator transaction binding the contract method 0x690bcc03.
//
// Solidity: function withdrawEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) WithdrawEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "withdrawEphemeral", smartObjectId, ephemeralOwner, items)
}

// WithdrawEphemeral is a paid mutator transaction binding the contract method 0x690bcc03.
//
// Solidity: function withdrawEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) WithdrawEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// WithdrawEphemeral is a paid mutator transaction binding the contract method 0x690bcc03.
//
// Solidity: function withdrawEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) WithdrawEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawEphemeral(&_ERC2771World.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771World *ERC2771WorldTransactor) WithdrawFuel(opts *bind.TransactOpts, smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "withdrawFuel", smartObjectId, fuelAmount)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771World *ERC2771WorldSession) WithdrawFuel(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawFuel(&_ERC2771World.TransactOpts, smartObjectId, fuelAmount)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) WithdrawFuel(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawFuel(&_ERC2771World.TransactOpts, smartObjectId, fuelAmount)
}

// WithdrawInventory is a paid mutator transaction binding the contract method 0xec92b5a5.
//
// Solidity: function withdrawInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) WithdrawInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "withdrawInventory", smartObjectId, items)
}

// WithdrawInventory is a paid mutator transaction binding the contract method 0xec92b5a5.
//
// Solidity: function withdrawInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) WithdrawInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// WithdrawInventory is a paid mutator transaction binding the contract method 0xec92b5a5.
//
// Solidity: function withdrawInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) WithdrawInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// ERC2771WorldHelloStoreIterator is returned from FilterHelloStore and is used to iterate over the raw logs and unpacked data for HelloStore events raised by the ERC2771World contract.
type ERC2771WorldHelloStoreIterator struct {
	Event *ERC2771WorldHelloStore // Event containing the contract specifics and raw log

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
func (it *ERC2771WorldHelloStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771WorldHelloStore)
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
		it.Event = new(ERC2771WorldHelloStore)
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
func (it *ERC2771WorldHelloStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771WorldHelloStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771WorldHelloStore represents a HelloStore event raised by the ERC2771World contract.
type ERC2771WorldHelloStore struct {
	StoreVersion [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHelloStore is a free log retrieval operation binding the contract event 0xc7f5fdc8526b76f54916701bc910876243ffff2a40b0bb8d59eea8151c52c005.
//
// Solidity: event HelloStore(bytes32 indexed storeVersion)
func (_ERC2771World *ERC2771WorldFilterer) FilterHelloStore(opts *bind.FilterOpts, storeVersion [][32]byte) (*ERC2771WorldHelloStoreIterator, error) {

	var storeVersionRule []interface{}
	for _, storeVersionItem := range storeVersion {
		storeVersionRule = append(storeVersionRule, storeVersionItem)
	}

	logs, sub, err := _ERC2771World.contract.FilterLogs(opts, "HelloStore", storeVersionRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771WorldHelloStoreIterator{contract: _ERC2771World.contract, event: "HelloStore", logs: logs, sub: sub}, nil
}

// WatchHelloStore is a free log subscription operation binding the contract event 0xc7f5fdc8526b76f54916701bc910876243ffff2a40b0bb8d59eea8151c52c005.
//
// Solidity: event HelloStore(bytes32 indexed storeVersion)
func (_ERC2771World *ERC2771WorldFilterer) WatchHelloStore(opts *bind.WatchOpts, sink chan<- *ERC2771WorldHelloStore, storeVersion [][32]byte) (event.Subscription, error) {

	var storeVersionRule []interface{}
	for _, storeVersionItem := range storeVersion {
		storeVersionRule = append(storeVersionRule, storeVersionItem)
	}

	logs, sub, err := _ERC2771World.contract.WatchLogs(opts, "HelloStore", storeVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771WorldHelloStore)
				if err := _ERC2771World.contract.UnpackLog(event, "HelloStore", log); err != nil {
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
func (_ERC2771World *ERC2771WorldFilterer) ParseHelloStore(log types.Log) (*ERC2771WorldHelloStore, error) {
	event := new(ERC2771WorldHelloStore)
	if err := _ERC2771World.contract.UnpackLog(event, "HelloStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771WorldHelloWorldIterator is returned from FilterHelloWorld and is used to iterate over the raw logs and unpacked data for HelloWorld events raised by the ERC2771World contract.
type ERC2771WorldHelloWorldIterator struct {
	Event *ERC2771WorldHelloWorld // Event containing the contract specifics and raw log

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
func (it *ERC2771WorldHelloWorldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771WorldHelloWorld)
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
		it.Event = new(ERC2771WorldHelloWorld)
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
func (it *ERC2771WorldHelloWorldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771WorldHelloWorldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771WorldHelloWorld represents a HelloWorld event raised by the ERC2771World contract.
type ERC2771WorldHelloWorld struct {
	WorldVersion [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHelloWorld is a free log retrieval operation binding the contract event 0x7f8f36afe3fb61c459c1a54a60b8a477eab02cc58e49f547561a40906239cb82.
//
// Solidity: event HelloWorld(bytes32 indexed worldVersion)
func (_ERC2771World *ERC2771WorldFilterer) FilterHelloWorld(opts *bind.FilterOpts, worldVersion [][32]byte) (*ERC2771WorldHelloWorldIterator, error) {

	var worldVersionRule []interface{}
	for _, worldVersionItem := range worldVersion {
		worldVersionRule = append(worldVersionRule, worldVersionItem)
	}

	logs, sub, err := _ERC2771World.contract.FilterLogs(opts, "HelloWorld", worldVersionRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771WorldHelloWorldIterator{contract: _ERC2771World.contract, event: "HelloWorld", logs: logs, sub: sub}, nil
}

// WatchHelloWorld is a free log subscription operation binding the contract event 0x7f8f36afe3fb61c459c1a54a60b8a477eab02cc58e49f547561a40906239cb82.
//
// Solidity: event HelloWorld(bytes32 indexed worldVersion)
func (_ERC2771World *ERC2771WorldFilterer) WatchHelloWorld(opts *bind.WatchOpts, sink chan<- *ERC2771WorldHelloWorld, worldVersion [][32]byte) (event.Subscription, error) {

	var worldVersionRule []interface{}
	for _, worldVersionItem := range worldVersion {
		worldVersionRule = append(worldVersionRule, worldVersionItem)
	}

	logs, sub, err := _ERC2771World.contract.WatchLogs(opts, "HelloWorld", worldVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771WorldHelloWorld)
				if err := _ERC2771World.contract.UnpackLog(event, "HelloWorld", log); err != nil {
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
func (_ERC2771World *ERC2771WorldFilterer) ParseHelloWorld(log types.Log) (*ERC2771WorldHelloWorld, error) {
	event := new(ERC2771WorldHelloWorld)
	if err := _ERC2771World.contract.UnpackLog(event, "HelloWorld", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771WorldStoreDeleteRecordIterator is returned from FilterStoreDeleteRecord and is used to iterate over the raw logs and unpacked data for StoreDeleteRecord events raised by the ERC2771World contract.
type ERC2771WorldStoreDeleteRecordIterator struct {
	Event *ERC2771WorldStoreDeleteRecord // Event containing the contract specifics and raw log

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
func (it *ERC2771WorldStoreDeleteRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771WorldStoreDeleteRecord)
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
		it.Event = new(ERC2771WorldStoreDeleteRecord)
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
func (it *ERC2771WorldStoreDeleteRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771WorldStoreDeleteRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771WorldStoreDeleteRecord represents a StoreDeleteRecord event raised by the ERC2771World contract.
type ERC2771WorldStoreDeleteRecord struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreDeleteRecord is a free log retrieval operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_ERC2771World *ERC2771WorldFilterer) FilterStoreDeleteRecord(opts *bind.FilterOpts, tableId [][32]byte) (*ERC2771WorldStoreDeleteRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771World.contract.FilterLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771WorldStoreDeleteRecordIterator{contract: _ERC2771World.contract, event: "Store_DeleteRecord", logs: logs, sub: sub}, nil
}

// WatchStoreDeleteRecord is a free log subscription operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_ERC2771World *ERC2771WorldFilterer) WatchStoreDeleteRecord(opts *bind.WatchOpts, sink chan<- *ERC2771WorldStoreDeleteRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771World.contract.WatchLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771WorldStoreDeleteRecord)
				if err := _ERC2771World.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
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
func (_ERC2771World *ERC2771WorldFilterer) ParseStoreDeleteRecord(log types.Log) (*ERC2771WorldStoreDeleteRecord, error) {
	event := new(ERC2771WorldStoreDeleteRecord)
	if err := _ERC2771World.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771WorldStoreSetRecordIterator is returned from FilterStoreSetRecord and is used to iterate over the raw logs and unpacked data for StoreSetRecord events raised by the ERC2771World contract.
type ERC2771WorldStoreSetRecordIterator struct {
	Event *ERC2771WorldStoreSetRecord // Event containing the contract specifics and raw log

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
func (it *ERC2771WorldStoreSetRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771WorldStoreSetRecord)
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
		it.Event = new(ERC2771WorldStoreSetRecord)
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
func (it *ERC2771WorldStoreSetRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771WorldStoreSetRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771WorldStoreSetRecord represents a StoreSetRecord event raised by the ERC2771World contract.
type ERC2771WorldStoreSetRecord struct {
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
func (_ERC2771World *ERC2771WorldFilterer) FilterStoreSetRecord(opts *bind.FilterOpts, tableId [][32]byte) (*ERC2771WorldStoreSetRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771World.contract.FilterLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771WorldStoreSetRecordIterator{contract: _ERC2771World.contract, event: "Store_SetRecord", logs: logs, sub: sub}, nil
}

// WatchStoreSetRecord is a free log subscription operation binding the contract event 0x8dbb3a9672eebfd3773e72dd9c102393436816d832c7ba9e1e1ac8fcadcac7a9.
//
// Solidity: event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771World *ERC2771WorldFilterer) WatchStoreSetRecord(opts *bind.WatchOpts, sink chan<- *ERC2771WorldStoreSetRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771World.contract.WatchLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771WorldStoreSetRecord)
				if err := _ERC2771World.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
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
func (_ERC2771World *ERC2771WorldFilterer) ParseStoreSetRecord(log types.Log) (*ERC2771WorldStoreSetRecord, error) {
	event := new(ERC2771WorldStoreSetRecord)
	if err := _ERC2771World.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771WorldStoreSpliceDynamicDataIterator is returned from FilterStoreSpliceDynamicData and is used to iterate over the raw logs and unpacked data for StoreSpliceDynamicData events raised by the ERC2771World contract.
type ERC2771WorldStoreSpliceDynamicDataIterator struct {
	Event *ERC2771WorldStoreSpliceDynamicData // Event containing the contract specifics and raw log

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
func (it *ERC2771WorldStoreSpliceDynamicDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771WorldStoreSpliceDynamicData)
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
		it.Event = new(ERC2771WorldStoreSpliceDynamicData)
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
func (it *ERC2771WorldStoreSpliceDynamicDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771WorldStoreSpliceDynamicDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771WorldStoreSpliceDynamicData represents a StoreSpliceDynamicData event raised by the ERC2771World contract.
type ERC2771WorldStoreSpliceDynamicData struct {
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
func (_ERC2771World *ERC2771WorldFilterer) FilterStoreSpliceDynamicData(opts *bind.FilterOpts, tableId [][32]byte) (*ERC2771WorldStoreSpliceDynamicDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771World.contract.FilterLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771WorldStoreSpliceDynamicDataIterator{contract: _ERC2771World.contract, event: "Store_SpliceDynamicData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceDynamicData is a free log subscription operation binding the contract event 0xfe158a7adba34e256807c8a149028d3162918713c3838afc643ce9f96716ebfd.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
func (_ERC2771World *ERC2771WorldFilterer) WatchStoreSpliceDynamicData(opts *bind.WatchOpts, sink chan<- *ERC2771WorldStoreSpliceDynamicData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771World.contract.WatchLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771WorldStoreSpliceDynamicData)
				if err := _ERC2771World.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
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
func (_ERC2771World *ERC2771WorldFilterer) ParseStoreSpliceDynamicData(log types.Log) (*ERC2771WorldStoreSpliceDynamicData, error) {
	event := new(ERC2771WorldStoreSpliceDynamicData)
	if err := _ERC2771World.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771WorldStoreSpliceStaticDataIterator is returned from FilterStoreSpliceStaticData and is used to iterate over the raw logs and unpacked data for StoreSpliceStaticData events raised by the ERC2771World contract.
type ERC2771WorldStoreSpliceStaticDataIterator struct {
	Event *ERC2771WorldStoreSpliceStaticData // Event containing the contract specifics and raw log

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
func (it *ERC2771WorldStoreSpliceStaticDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771WorldStoreSpliceStaticData)
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
		it.Event = new(ERC2771WorldStoreSpliceStaticData)
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
func (it *ERC2771WorldStoreSpliceStaticDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771WorldStoreSpliceStaticDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771WorldStoreSpliceStaticData represents a StoreSpliceStaticData event raised by the ERC2771World contract.
type ERC2771WorldStoreSpliceStaticData struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Start    *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreSpliceStaticData is a free log retrieval operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_ERC2771World *ERC2771WorldFilterer) FilterStoreSpliceStaticData(opts *bind.FilterOpts, tableId [][32]byte) (*ERC2771WorldStoreSpliceStaticDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771World.contract.FilterLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771WorldStoreSpliceStaticDataIterator{contract: _ERC2771World.contract, event: "Store_SpliceStaticData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceStaticData is a free log subscription operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_ERC2771World *ERC2771WorldFilterer) WatchStoreSpliceStaticData(opts *bind.WatchOpts, sink chan<- *ERC2771WorldStoreSpliceStaticData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771World.contract.WatchLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771WorldStoreSpliceStaticData)
				if err := _ERC2771World.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
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
func (_ERC2771World *ERC2771WorldFilterer) ParseStoreSpliceStaticData(log types.Log) (*ERC2771WorldStoreSpliceStaticData, error) {
	event := new(ERC2771WorldStoreSpliceStaticData)
	if err := _ERC2771World.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
