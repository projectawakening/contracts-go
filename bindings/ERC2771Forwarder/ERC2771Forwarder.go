// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC2771Forwarder

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
	FuelTypeId            *big.Int
	FuelUnitVolume        *big.Int
	FuelMaxCapacity       *big.Int
	FuelBurnRateInSeconds *big.Int
	FuelAmount            *big.Int
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

// ERC2771ForwarderMetaData contains all meta data concerning the ERC2771Forwarder contract.
var ERC2771ForwarderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"batchCall\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallData[]\",\"components\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchCallFrom\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallFromData[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"callFrom\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"creator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminSupportOrDirectOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"adminSupportOrDirectOwnerGates\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggression\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structAggressionParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"aggressor\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"victim\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"anchor\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"areGatesOnline\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assignItemToInventory\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assignOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bringOffline\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bringOnline\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canCrossTransferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canJump\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canTransferFromEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canTransferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canTransferToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configureDeployableAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureEntityRecordAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureEphemeralInteractAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureEphemeralInventoryAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureFuelAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureFuelEfficiency\",\"inputs\":[{\"name\":\"fuelTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelEfficiency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureFuelParameters\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelParams\",\"type\":\"tuple\",\"internalType\":\"structFuelParams\",\"components\":[{\"name\":\"fuelTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelBurnRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureGate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureInventoryAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureInventoryInteractAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureKillMailAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureLocationAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureNetworkNodeAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureOwnershipAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartAssemblyAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartCharacterAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartGateAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartStorageUnitAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartTurretAccess\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureTurret\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"connectStructure\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"structureId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchor\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorGate\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"maxDistance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorNetworkNode\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelParams\",\"type\":\"tuple\",\"internalType\":\"structFuelParams\",\"components\":[{\"name\":\"fuelTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelBurnRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"maxEnergyCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentProduction\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorStorageUnit\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorTurret\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCreateAndAnchorParams\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndDepositEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structCreateInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndDepositInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structCreateInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAssembly\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createCharacter\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tribeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"entityRecordMetadata\",\"type\":\"tuple\",\"internalType\":\"structEntityMetadataParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createDeployable\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createMetadata\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordMetadata\",\"type\":\"tuple\",\"internalType\":\"structEntityMetadataParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRecord\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordParams\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordParams\",\"components\":[{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossTransferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromEphemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toEphemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"destroyDeployable\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentFuelConsumptionStatus\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"timeLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitsToConsume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualConsumptionRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEphemeralOwner\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEphemeralSmartObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getInventoryOwner\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNetworkNodeClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSmartCharacterClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSmartGateClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSmartStorageUnitClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSmartTurretClassId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleNodeOffline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"inProximity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turretTarget\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAnyGateLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isClassScoped\",\"inputs\":[{\"name\":\"classId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEphemeralOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isGateLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerOfBothGates\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWithinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onStructureOffline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"structureId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onStructureOnline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"structureId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onlyAdminOrCallAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminOrClassScopedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminOrOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminOrOwnerSupported\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminOrScopeEnforcedCall\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminSupportedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdminSupportedOwnerOrCall\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyCallAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyCallAccessOrDirectEphemeralOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyCallAccessWithScopeEnforced\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyClassScopedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyClassScopedOrCharAdminOrOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyDirectAdmin\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyDirectAdminOrCallAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyDirectEphemeralOwnerOrCall\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyDirectOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyEphemeralOwnerOrTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwnerOrEphemeralCrossTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwnerOrEphemeralTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwnerOrInventoryTransferRole\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyOwnerWithAdminSupportAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlySmartAssemblyClassScopedAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerNetworkNodeClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSmartCharacterClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSmartGateClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSmartStorageUnitClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSmartTurretClass\",\"inputs\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeCharacter\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeItemFromInventory\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportKill\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killMailData\",\"type\":\"tuple\",\"internalType\":\"structKillMailData\",\"components\":[{\"name\":\"killerCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"victimCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lossType\",\"type\":\"uint8\",\"internalType\":\"enumKillMailLossType\"},{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"saveLocation\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAssemblyType\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCrossTransferToEphemeralAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDappURL\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDescription\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEphemeralCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFuelAmount\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFuelMaxCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFuelUnitVolume\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTransferFromEphemeralAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTransferToEphemeralAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTransferToInventoryAccess\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBurn\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopBurn\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFromEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferToEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unanchor\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlinkGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAssemblyType\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assemblyType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTribeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tribeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawEphemeral\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItemParams[]\",\"components\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldSlice\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLayout\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeySchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValueSchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initModule\",\"type\":\"address\",\"internalType\":\"contractIModule\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installRootModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"popFromDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"byteLengthToPop\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pushToDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dataToPush\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"systemFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespace\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerRootFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"worldFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"systemFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystem\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"contractSystem\"},{\"name\":\"publicAccess\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerTable\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"},{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"keyNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"fieldNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storeVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferBalanceToAddress\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferBalanceToNamespace\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"worldVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"HelloStore\",\"inputs\":[{\"name\":\"storeVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HelloWorld\",\"inputs\":[{\"name\":\"worldVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_DeleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SetRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"EncodedLengths\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"EncodedLengths\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Access_CannotTransferFromEphemeral\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrClassScoped\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminOrOwnerSupported\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupported\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupportedOrDirectOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupportedOrDirectOwnerGates\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotAdminSupportedOwnerOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotClassScoped\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotClassScopedAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectAdminOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectEphemeralOwnerOrCanCrossTransferToEphemeral\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectOwnerOrCanTransferToEphemeral\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotDirectOwnerOrCanTransferToInventory\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotEphemeralOwnerOrCallAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotEphemeralOwnerOrCallAccessWithEphemeralOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Access_NotOwnerWithAdminSupportAccess\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Deployable_IncorrectState\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentState\",\"type\":\"uint8\",\"internalType\":\"enumState\"}]},{\"type\":\"error\",\"name\":\"Deployable_InvalidObjectOwner\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"smartObjectOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EncodedLengths_InvalidLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InsufficientCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"usedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidEphemeralOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidItemDepositQuantity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidItemObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidSmartObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_InvalidTenantId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EphemeralInventory_NonExistentEntityRecord\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FieldLayout_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"staticDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"computedStaticDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_StaticLengthDoesNotFitInAWord\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_StaticLengthIsNotZero\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_StaticLengthIsZero\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_TooManyDynamicFields\",\"inputs\":[{\"name\":\"numFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayout_TooManyFields\",\"inputs\":[{\"name\":\"numFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_BurnAlreadyStopped\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_BurnNotActive\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_ExceedsMaxCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalProjectedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InsufficientFuel\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableFuel\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelAmount\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelBurnRate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelBurnRateInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelEfficiency\",\"inputs\":[{\"name\":\"fuelTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelEfficiency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelMaxCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelTypeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Fuel_InvalidFuelUnitVolume\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_Ephemeral_InsufficientQuantity\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InsufficientQuantity\",\"inputs\":[{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InvalidInventory\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InvalidOperation\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_InvalidQuantity\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_NonexistentItemRecord\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_NonexistentObject\",\"inputs\":[{\"name\":\"objectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_SingletonAlreadyAssigned\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentInventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_SingletonDirectlyOwned\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"directOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InventoryOwnership_ZeroQuantity\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InsufficientCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"usedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidItemDepositQuantity\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidItemObjectId\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidTenantId\",\"inputs\":[{\"name\":\"itemObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"Inventory_NonExistentEntityRecord\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"KillMail_AlreadyExists\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"KillMail_InvalidCharacterId\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Module_AlreadyInstalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Module_MissingDependency\",\"inputs\":[{\"name\":\"dependency\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Module_NonRootInstallNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Module_RootInstallNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NetworkNode_AlreadyExists\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_DoesNotExist\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_InsufficientEnergy\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_NotOnline\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_StructureAlreadyConnected\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"structureId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NetworkNode_StructureNotConnected\",\"inputs\":[{\"name\":\"networkNodeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"structureId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Ownership_AlreadyOwned\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Ownership_InvalidAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Ownership_InvalidOwner\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"invalidOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Ownership_InvalidSingleton\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Ownership_NonexistentObject\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Ownership_SingletonInInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inventoryObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Schema_InvalidLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Schema_StaticTypeAfterDynamicType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Slice_OutOfBounds\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_DoesNotExist\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_InvalidObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_InvalidTenantId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_InvalidTypeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartAssembly_TypeCannotBeEmpty\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacterDoesNotExist\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_AlreadyCreated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_InvalidObjectId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_InvalidTenantId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tenantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_InvalidTypeId\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateAlreadyLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateNotLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateNotOnline\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GatesNotOnline\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotWithtinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_SameSourceAndDestination\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_UndefinedClassId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Store_IndexOutOfBounds\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessedIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidBounds\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidFieldNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidKeyNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidSplice\",\"inputs\":[{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"fieldLength\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaDynamicLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaStaticLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_TableAlreadyExists\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_TableNotFound\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_AccessDenied\",\"inputs\":[{\"name\":\"resource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"World_CallbackNotAllowed\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_DelegationNotFound\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorAlreadyExists\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorNotFound\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"World_InterfaceNotSupported\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InvalidNamespace\",\"inputs\":[{\"name\":\"namespace\",\"type\":\"bytes14\",\"internalType\":\"bytes14\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceId\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceAlreadyExists\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceNotFound\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_SystemAlreadyExists\",\"inputs\":[{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_UnlimitedDelegationNotAllowed\",\"inputs\":[]}]",
}

// ERC2771ForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC2771ForwarderMetaData.ABI instead.
var ERC2771ForwarderABI = ERC2771ForwarderMetaData.ABI

// ERC2771Forwarder is an auto generated Go binding around an Ethereum contract.
type ERC2771Forwarder struct {
	ERC2771ForwarderCaller     // Read-only binding to the contract
	ERC2771ForwarderTransactor // Write-only binding to the contract
	ERC2771ForwarderFilterer   // Log filterer for contract events
}

// ERC2771ForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC2771ForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771ForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC2771ForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771ForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC2771ForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771ForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC2771ForwarderSession struct {
	Contract     *ERC2771Forwarder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC2771ForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC2771ForwarderCallerSession struct {
	Contract *ERC2771ForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC2771ForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC2771ForwarderTransactorSession struct {
	Contract     *ERC2771ForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC2771ForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC2771ForwarderRaw struct {
	Contract *ERC2771Forwarder // Generic contract binding to access the raw methods on
}

// ERC2771ForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC2771ForwarderCallerRaw struct {
	Contract *ERC2771ForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// ERC2771ForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC2771ForwarderTransactorRaw struct {
	Contract *ERC2771ForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC2771Forwarder creates a new instance of ERC2771Forwarder, bound to a specific deployed contract.
func NewERC2771Forwarder(address common.Address, backend bind.ContractBackend) (*ERC2771Forwarder, error) {
	contract, err := bindERC2771Forwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC2771Forwarder{ERC2771ForwarderCaller: ERC2771ForwarderCaller{contract: contract}, ERC2771ForwarderTransactor: ERC2771ForwarderTransactor{contract: contract}, ERC2771ForwarderFilterer: ERC2771ForwarderFilterer{contract: contract}}, nil
}

// NewERC2771ForwarderCaller creates a new read-only instance of ERC2771Forwarder, bound to a specific deployed contract.
func NewERC2771ForwarderCaller(address common.Address, caller bind.ContractCaller) (*ERC2771ForwarderCaller, error) {
	contract, err := bindERC2771Forwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderCaller{contract: contract}, nil
}

// NewERC2771ForwarderTransactor creates a new write-only instance of ERC2771Forwarder, bound to a specific deployed contract.
func NewERC2771ForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC2771ForwarderTransactor, error) {
	contract, err := bindERC2771Forwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderTransactor{contract: contract}, nil
}

// NewERC2771ForwarderFilterer creates a new log filterer instance of ERC2771Forwarder, bound to a specific deployed contract.
func NewERC2771ForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC2771ForwarderFilterer, error) {
	contract, err := bindERC2771Forwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderFilterer{contract: contract}, nil
}

// bindERC2771Forwarder binds a generic wrapper to an already deployed contract.
func bindERC2771Forwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC2771ForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC2771Forwarder *ERC2771ForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC2771Forwarder.Contract.ERC2771ForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC2771Forwarder *ERC2771ForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ERC2771ForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC2771Forwarder *ERC2771ForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ERC2771ForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC2771Forwarder *ERC2771ForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC2771Forwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC2771Forwarder *ERC2771ForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC2771Forwarder *ERC2771ForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.contract.Transact(opts, method, params...)
}

// AdminSupportOrDirectOwner is a free data retrieval call binding the contract method 0xc6e0b1a1.
//
// Solidity: function adminSupportOrDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) AdminSupportOrDirectOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "adminSupportOrDirectOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// AdminSupportOrDirectOwner is a free data retrieval call binding the contract method 0xc6e0b1a1.
//
// Solidity: function adminSupportOrDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) AdminSupportOrDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.AdminSupportOrDirectOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// AdminSupportOrDirectOwner is a free data retrieval call binding the contract method 0xc6e0b1a1.
//
// Solidity: function adminSupportOrDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) AdminSupportOrDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.AdminSupportOrDirectOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// AdminSupportOrDirectOwnerGates is a free data retrieval call binding the contract method 0xa6f28a83.
//
// Solidity: function adminSupportOrDirectOwnerGates(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) AdminSupportOrDirectOwnerGates(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "adminSupportOrDirectOwnerGates", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// AdminSupportOrDirectOwnerGates is a free data retrieval call binding the contract method 0xa6f28a83.
//
// Solidity: function adminSupportOrDirectOwnerGates(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) AdminSupportOrDirectOwnerGates(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.AdminSupportOrDirectOwnerGates(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// AdminSupportOrDirectOwnerGates is a free data retrieval call binding the contract method 0xa6f28a83.
//
// Solidity: function adminSupportOrDirectOwnerGates(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) AdminSupportOrDirectOwnerGates(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.AdminSupportOrDirectOwnerGates(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// AreGatesOnline is a free data retrieval call binding the contract method 0x9b11b875.
//
// Solidity: function areGatesOnline(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) AreGatesOnline(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "areGatesOnline", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreGatesOnline is a free data retrieval call binding the contract method 0x9b11b875.
//
// Solidity: function areGatesOnline(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) AreGatesOnline(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.AreGatesOnline(&_ERC2771Forwarder.CallOpts, sourceGateId, destinationGateId)
}

// AreGatesOnline is a free data retrieval call binding the contract method 0x9b11b875.
//
// Solidity: function areGatesOnline(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) AreGatesOnline(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.AreGatesOnline(&_ERC2771Forwarder.CallOpts, sourceGateId, destinationGateId)
}

// CanCrossTransferToEphemeral is a free data retrieval call binding the contract method 0x95d62522.
//
// Solidity: function canCrossTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) CanCrossTransferToEphemeral(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "canCrossTransferToEphemeral", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCrossTransferToEphemeral is a free data retrieval call binding the contract method 0x95d62522.
//
// Solidity: function canCrossTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) CanCrossTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.CanCrossTransferToEphemeral(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// CanCrossTransferToEphemeral is a free data retrieval call binding the contract method 0x95d62522.
//
// Solidity: function canCrossTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) CanCrossTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.CanCrossTransferToEphemeral(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// CanTransferFromEphemeral is a free data retrieval call binding the contract method 0x609fb5b0.
//
// Solidity: function canTransferFromEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) CanTransferFromEphemeral(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "canTransferFromEphemeral", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransferFromEphemeral is a free data retrieval call binding the contract method 0x609fb5b0.
//
// Solidity: function canTransferFromEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) CanTransferFromEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.CanTransferFromEphemeral(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// CanTransferFromEphemeral is a free data retrieval call binding the contract method 0x609fb5b0.
//
// Solidity: function canTransferFromEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) CanTransferFromEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.CanTransferFromEphemeral(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// CanTransferToEphemeral is a free data retrieval call binding the contract method 0x88756560.
//
// Solidity: function canTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) CanTransferToEphemeral(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "canTransferToEphemeral", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransferToEphemeral is a free data retrieval call binding the contract method 0x88756560.
//
// Solidity: function canTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) CanTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.CanTransferToEphemeral(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// CanTransferToEphemeral is a free data retrieval call binding the contract method 0x88756560.
//
// Solidity: function canTransferToEphemeral(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) CanTransferToEphemeral(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.CanTransferToEphemeral(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// CanTransferToInventory is a free data retrieval call binding the contract method 0x0fcd8ed6.
//
// Solidity: function canTransferToInventory(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) CanTransferToInventory(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "canTransferToInventory", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransferToInventory is a free data retrieval call binding the contract method 0x0fcd8ed6.
//
// Solidity: function canTransferToInventory(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) CanTransferToInventory(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.CanTransferToInventory(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// CanTransferToInventory is a free data retrieval call binding the contract method 0x0fcd8ed6.
//
// Solidity: function canTransferToInventory(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) CanTransferToInventory(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.CanTransferToInventory(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderSession) Creator() (common.Address, error) {
	return _ERC2771Forwarder.Contract.Creator(&_ERC2771Forwarder.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) Creator() (common.Address, error) {
	return _ERC2771Forwarder.Contract.Creator(&_ERC2771Forwarder.CallOpts)
}

// GetCurrentFuelConsumptionStatus is a free data retrieval call binding the contract method 0x655d9035.
//
// Solidity: function getCurrentFuelConsumptionStatus(uint256 smartObjectId) view returns(uint256 timeLeft, uint256 unitsToConsume, uint256 actualConsumptionRateInSeconds, uint256 fuelAmount)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetCurrentFuelConsumptionStatus(opts *bind.CallOpts, smartObjectId *big.Int) (struct {
	TimeLeft                       *big.Int
	UnitsToConsume                 *big.Int
	ActualConsumptionRateInSeconds *big.Int
	FuelAmount                     *big.Int
}, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getCurrentFuelConsumptionStatus", smartObjectId)

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
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetCurrentFuelConsumptionStatus(smartObjectId *big.Int) (struct {
	TimeLeft                       *big.Int
	UnitsToConsume                 *big.Int
	ActualConsumptionRateInSeconds *big.Int
	FuelAmount                     *big.Int
}, error) {
	return _ERC2771Forwarder.Contract.GetCurrentFuelConsumptionStatus(&_ERC2771Forwarder.CallOpts, smartObjectId)
}

// GetCurrentFuelConsumptionStatus is a free data retrieval call binding the contract method 0x655d9035.
//
// Solidity: function getCurrentFuelConsumptionStatus(uint256 smartObjectId) view returns(uint256 timeLeft, uint256 unitsToConsume, uint256 actualConsumptionRateInSeconds, uint256 fuelAmount)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetCurrentFuelConsumptionStatus(smartObjectId *big.Int) (struct {
	TimeLeft                       *big.Int
	UnitsToConsume                 *big.Int
	ActualConsumptionRateInSeconds *big.Int
	FuelAmount                     *big.Int
}, error) {
	return _ERC2771Forwarder.Contract.GetCurrentFuelConsumptionStatus(&_ERC2771Forwarder.CallOpts, smartObjectId)
}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetDynamicField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getDynamicField", tableId, keyTuple, dynamicFieldIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	return _ERC2771Forwarder.Contract.GetDynamicField(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	return _ERC2771Forwarder.Contract.GetDynamicField(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetDynamicFieldLength(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getDynamicFieldLength", tableId, keyTuple, dynamicFieldIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetDynamicFieldLength(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetDynamicFieldLength(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetDynamicFieldLength(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetDynamicFieldLength(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetDynamicFieldSlice(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getDynamicFieldSlice", tableId, keyTuple, dynamicFieldIndex, start, end)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetDynamicFieldSlice(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	return _ERC2771Forwarder.Contract.GetDynamicFieldSlice(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, dynamicFieldIndex, start, end)
}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetDynamicFieldSlice(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	return _ERC2771Forwarder.Contract.GetDynamicFieldSlice(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, dynamicFieldIndex, start, end)
}

// GetEphemeralOwner is a free data retrieval call binding the contract method 0x2a53720d.
//
// Solidity: function getEphemeralOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetEphemeralOwner(opts *bind.CallOpts, inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getEphemeralOwner", inventoryObjectId, itemObjectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEphemeralOwner is a free data retrieval call binding the contract method 0x2a53720d.
//
// Solidity: function getEphemeralOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetEphemeralOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _ERC2771Forwarder.Contract.GetEphemeralOwner(&_ERC2771Forwarder.CallOpts, inventoryObjectId, itemObjectId)
}

// GetEphemeralOwner is a free data retrieval call binding the contract method 0x2a53720d.
//
// Solidity: function getEphemeralOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetEphemeralOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _ERC2771Forwarder.Contract.GetEphemeralOwner(&_ERC2771Forwarder.CallOpts, inventoryObjectId, itemObjectId)
}

// GetEphemeralSmartObjectId is a free data retrieval call binding the contract method 0x83619510.
//
// Solidity: function getEphemeralSmartObjectId(uint256 smartObjectId, address ephemeralOwner) pure returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetEphemeralSmartObjectId(opts *bind.CallOpts, smartObjectId *big.Int, ephemeralOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getEphemeralSmartObjectId", smartObjectId, ephemeralOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEphemeralSmartObjectId is a free data retrieval call binding the contract method 0x83619510.
//
// Solidity: function getEphemeralSmartObjectId(uint256 smartObjectId, address ephemeralOwner) pure returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetEphemeralSmartObjectId(smartObjectId *big.Int, ephemeralOwner common.Address) (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetEphemeralSmartObjectId(&_ERC2771Forwarder.CallOpts, smartObjectId, ephemeralOwner)
}

// GetEphemeralSmartObjectId is a free data retrieval call binding the contract method 0x83619510.
//
// Solidity: function getEphemeralSmartObjectId(uint256 smartObjectId, address ephemeralOwner) pure returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetEphemeralSmartObjectId(smartObjectId *big.Int, ephemeralOwner common.Address) (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetEphemeralSmartObjectId(&_ERC2771Forwarder.CallOpts, smartObjectId, ephemeralOwner)
}

// GetField is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getField", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetField is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	return _ERC2771Forwarder.Contract.GetField(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetField is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	return _ERC2771Forwarder.Contract.GetField(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetField0 is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetField0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getField0", tableId, keyTuple, fieldIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetField0 is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	return _ERC2771Forwarder.Contract.GetField0(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetField0 is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	return _ERC2771Forwarder.Contract.GetField0(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetFieldLayout(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getFieldLayout", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetFieldLayout(tableId [32]byte) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.GetFieldLayout(&_ERC2771Forwarder.CallOpts, tableId)
}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetFieldLayout(tableId [32]byte) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.GetFieldLayout(&_ERC2771Forwarder.CallOpts, tableId)
}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetFieldLength(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getFieldLength", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetFieldLength(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetFieldLength(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetFieldLength(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetFieldLength(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetFieldLength0 is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetFieldLength0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getFieldLength0", tableId, keyTuple, fieldIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldLength0 is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetFieldLength0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetFieldLength0(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetFieldLength0 is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetFieldLength0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetFieldLength0(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetInventoryOwner is a free data retrieval call binding the contract method 0x1904f95b.
//
// Solidity: function getInventoryOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetInventoryOwner(opts *bind.CallOpts, inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getInventoryOwner", inventoryObjectId, itemObjectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInventoryOwner is a free data retrieval call binding the contract method 0x1904f95b.
//
// Solidity: function getInventoryOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetInventoryOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _ERC2771Forwarder.Contract.GetInventoryOwner(&_ERC2771Forwarder.CallOpts, inventoryObjectId, itemObjectId)
}

// GetInventoryOwner is a free data retrieval call binding the contract method 0x1904f95b.
//
// Solidity: function getInventoryOwner(uint256 inventoryObjectId, uint256 itemObjectId) view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetInventoryOwner(inventoryObjectId *big.Int, itemObjectId *big.Int) (common.Address, error) {
	return _ERC2771Forwarder.Contract.GetInventoryOwner(&_ERC2771Forwarder.CallOpts, inventoryObjectId, itemObjectId)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetKeySchema(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getKeySchema", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetKeySchema(tableId [32]byte) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.GetKeySchema(&_ERC2771Forwarder.CallOpts, tableId)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetKeySchema(tableId [32]byte) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.GetKeySchema(&_ERC2771Forwarder.CallOpts, tableId)
}

// GetNetworkNodeClassId is a free data retrieval call binding the contract method 0x520b91b9.
//
// Solidity: function getNetworkNodeClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetNetworkNodeClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getNetworkNodeClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNetworkNodeClassId is a free data retrieval call binding the contract method 0x520b91b9.
//
// Solidity: function getNetworkNodeClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetNetworkNodeClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetNetworkNodeClassId(&_ERC2771Forwarder.CallOpts)
}

// GetNetworkNodeClassId is a free data retrieval call binding the contract method 0x520b91b9.
//
// Solidity: function getNetworkNodeClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetNetworkNodeClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetNetworkNodeClassId(&_ERC2771Forwarder.CallOpts)
}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetRecord(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getRecord", tableId, keyTuple, fieldLayout)

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
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetRecord(tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _ERC2771Forwarder.Contract.GetRecord(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldLayout)
}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetRecord(tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _ERC2771Forwarder.Contract.GetRecord(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldLayout)
}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetRecord0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getRecord0", tableId, keyTuple)

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
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetRecord0(tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _ERC2771Forwarder.Contract.GetRecord0(&_ERC2771Forwarder.CallOpts, tableId, keyTuple)
}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetRecord0(tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _ERC2771Forwarder.Contract.GetRecord0(&_ERC2771Forwarder.CallOpts, tableId, keyTuple)
}

// GetSmartCharacterClassId is a free data retrieval call binding the contract method 0x95b9e2d3.
//
// Solidity: function getSmartCharacterClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetSmartCharacterClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getSmartCharacterClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmartCharacterClassId is a free data retrieval call binding the contract method 0x95b9e2d3.
//
// Solidity: function getSmartCharacterClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetSmartCharacterClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetSmartCharacterClassId(&_ERC2771Forwarder.CallOpts)
}

// GetSmartCharacterClassId is a free data retrieval call binding the contract method 0x95b9e2d3.
//
// Solidity: function getSmartCharacterClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetSmartCharacterClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetSmartCharacterClassId(&_ERC2771Forwarder.CallOpts)
}

// GetSmartGateClassId is a free data retrieval call binding the contract method 0xc7afc974.
//
// Solidity: function getSmartGateClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetSmartGateClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getSmartGateClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmartGateClassId is a free data retrieval call binding the contract method 0xc7afc974.
//
// Solidity: function getSmartGateClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetSmartGateClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetSmartGateClassId(&_ERC2771Forwarder.CallOpts)
}

// GetSmartGateClassId is a free data retrieval call binding the contract method 0xc7afc974.
//
// Solidity: function getSmartGateClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetSmartGateClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetSmartGateClassId(&_ERC2771Forwarder.CallOpts)
}

// GetSmartStorageUnitClassId is a free data retrieval call binding the contract method 0xf3a7f1bb.
//
// Solidity: function getSmartStorageUnitClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetSmartStorageUnitClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getSmartStorageUnitClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmartStorageUnitClassId is a free data retrieval call binding the contract method 0xf3a7f1bb.
//
// Solidity: function getSmartStorageUnitClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetSmartStorageUnitClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetSmartStorageUnitClassId(&_ERC2771Forwarder.CallOpts)
}

// GetSmartStorageUnitClassId is a free data retrieval call binding the contract method 0xf3a7f1bb.
//
// Solidity: function getSmartStorageUnitClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetSmartStorageUnitClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetSmartStorageUnitClassId(&_ERC2771Forwarder.CallOpts)
}

// GetSmartTurretClassId is a free data retrieval call binding the contract method 0x4967ae58.
//
// Solidity: function getSmartTurretClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetSmartTurretClassId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getSmartTurretClassId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmartTurretClassId is a free data retrieval call binding the contract method 0x4967ae58.
//
// Solidity: function getSmartTurretClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetSmartTurretClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetSmartTurretClassId(&_ERC2771Forwarder.CallOpts)
}

// GetSmartTurretClassId is a free data retrieval call binding the contract method 0x4967ae58.
//
// Solidity: function getSmartTurretClassId() view returns(uint256)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetSmartTurretClassId() (*big.Int, error) {
	return _ERC2771Forwarder.Contract.GetSmartTurretClassId(&_ERC2771Forwarder.CallOpts)
}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetStaticField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getStaticField", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.GetStaticField(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.GetStaticField(&_ERC2771Forwarder.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) GetValueSchema(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "getValueSchema", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_ERC2771Forwarder *ERC2771ForwarderSession) GetValueSchema(tableId [32]byte) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.GetValueSchema(&_ERC2771Forwarder.CallOpts, tableId)
}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) GetValueSchema(tableId [32]byte) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.GetValueSchema(&_ERC2771Forwarder.CallOpts, tableId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) IsAdmin(opts *bind.CallOpts, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "isAdmin", caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) IsAdmin(caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.IsAdmin(&_ERC2771Forwarder.CallOpts, caller)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) IsAdmin(caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.IsAdmin(&_ERC2771Forwarder.CallOpts, caller)
}

// IsAnyGateLinked is a free data retrieval call binding the contract method 0x7222f910.
//
// Solidity: function isAnyGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) IsAnyGateLinked(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "isAnyGateLinked", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAnyGateLinked is a free data retrieval call binding the contract method 0x7222f910.
//
// Solidity: function isAnyGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) IsAnyGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.IsAnyGateLinked(&_ERC2771Forwarder.CallOpts, sourceGateId, destinationGateId)
}

// IsAnyGateLinked is a free data retrieval call binding the contract method 0x7222f910.
//
// Solidity: function isAnyGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) IsAnyGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.IsAnyGateLinked(&_ERC2771Forwarder.CallOpts, sourceGateId, destinationGateId)
}

// IsClassScoped is a free data retrieval call binding the contract method 0x590d5cdd.
//
// Solidity: function isClassScoped(uint256 classId, bytes32 systemId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) IsClassScoped(opts *bind.CallOpts, classId *big.Int, systemId [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "isClassScoped", classId, systemId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClassScoped is a free data retrieval call binding the contract method 0x590d5cdd.
//
// Solidity: function isClassScoped(uint256 classId, bytes32 systemId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) IsClassScoped(classId *big.Int, systemId [32]byte) (bool, error) {
	return _ERC2771Forwarder.Contract.IsClassScoped(&_ERC2771Forwarder.CallOpts, classId, systemId)
}

// IsClassScoped is a free data retrieval call binding the contract method 0x590d5cdd.
//
// Solidity: function isClassScoped(uint256 classId, bytes32 systemId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) IsClassScoped(classId *big.Int, systemId [32]byte) (bool, error) {
	return _ERC2771Forwarder.Contract.IsClassScoped(&_ERC2771Forwarder.CallOpts, classId, systemId)
}

// IsEphemeralOwner is a free data retrieval call binding the contract method 0x9f218381.
//
// Solidity: function isEphemeralOwner(uint256 smartObjectId, address caller, bytes data) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) IsEphemeralOwner(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "isEphemeralOwner", smartObjectId, caller, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEphemeralOwner is a free data retrieval call binding the contract method 0x9f218381.
//
// Solidity: function isEphemeralOwner(uint256 smartObjectId, address caller, bytes data) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) IsEphemeralOwner(smartObjectId *big.Int, caller common.Address, data []byte) (bool, error) {
	return _ERC2771Forwarder.Contract.IsEphemeralOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, caller, data)
}

// IsEphemeralOwner is a free data retrieval call binding the contract method 0x9f218381.
//
// Solidity: function isEphemeralOwner(uint256 smartObjectId, address caller, bytes data) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) IsEphemeralOwner(smartObjectId *big.Int, caller common.Address, data []byte) (bool, error) {
	return _ERC2771Forwarder.Contract.IsEphemeralOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, caller, data)
}

// IsGateLinked is a free data retrieval call binding the contract method 0x2ab90d4d.
//
// Solidity: function isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) IsGateLinked(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "isGateLinked", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGateLinked is a free data retrieval call binding the contract method 0x2ab90d4d.
//
// Solidity: function isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) IsGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.IsGateLinked(&_ERC2771Forwarder.CallOpts, sourceGateId, destinationGateId)
}

// IsGateLinked is a free data retrieval call binding the contract method 0x2ab90d4d.
//
// Solidity: function isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) IsGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.IsGateLinked(&_ERC2771Forwarder.CallOpts, sourceGateId, destinationGateId)
}

// IsOwner is a free data retrieval call binding the contract method 0x5a5d096c.
//
// Solidity: function isOwner(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) IsOwner(opts *bind.CallOpts, smartObjectId *big.Int, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "isOwner", smartObjectId, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x5a5d096c.
//
// Solidity: function isOwner(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) IsOwner(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.IsOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// IsOwner is a free data retrieval call binding the contract method 0x5a5d096c.
//
// Solidity: function isOwner(uint256 smartObjectId, address caller) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) IsOwner(smartObjectId *big.Int, caller common.Address) (bool, error) {
	return _ERC2771Forwarder.Contract.IsOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, caller)
}

// IsOwnerOfBothGates is a free data retrieval call binding the contract method 0xa134720c.
//
// Solidity: function isOwnerOfBothGates(address caller, bytes data) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) IsOwnerOfBothGates(opts *bind.CallOpts, caller common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "isOwnerOfBothGates", caller, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwnerOfBothGates is a free data retrieval call binding the contract method 0xa134720c.
//
// Solidity: function isOwnerOfBothGates(address caller, bytes data) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) IsOwnerOfBothGates(caller common.Address, data []byte) (bool, error) {
	return _ERC2771Forwarder.Contract.IsOwnerOfBothGates(&_ERC2771Forwarder.CallOpts, caller, data)
}

// IsOwnerOfBothGates is a free data retrieval call binding the contract method 0xa134720c.
//
// Solidity: function isOwnerOfBothGates(address caller, bytes data) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) IsOwnerOfBothGates(caller common.Address, data []byte) (bool, error) {
	return _ERC2771Forwarder.Contract.IsOwnerOfBothGates(&_ERC2771Forwarder.CallOpts, caller, data)
}

// IsWithinRange is a free data retrieval call binding the contract method 0x36dfd147.
//
// Solidity: function isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) IsWithinRange(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "isWithinRange", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithinRange is a free data retrieval call binding the contract method 0x36dfd147.
//
// Solidity: function isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) IsWithinRange(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.IsWithinRange(&_ERC2771Forwarder.CallOpts, sourceGateId, destinationGateId)
}

// IsWithinRange is a free data retrieval call binding the contract method 0x36dfd147.
//
// Solidity: function isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) IsWithinRange(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.IsWithinRange(&_ERC2771Forwarder.CallOpts, sourceGateId, destinationGateId)
}

// OnlyAdminOrCallAccess is a free data retrieval call binding the contract method 0x1d5a7107.
//
// Solidity: function onlyAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyAdminOrCallAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyAdminOrCallAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrCallAccess is a free data retrieval call binding the contract method 0x1d5a7107.
//
// Solidity: function onlyAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrCallAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminOrCallAccess is a free data retrieval call binding the contract method 0x1d5a7107.
//
// Solidity: function onlyAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrCallAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminOrClassScopedAccess is a free data retrieval call binding the contract method 0x31098e24.
//
// Solidity: function onlyAdminOrClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyAdminOrClassScopedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyAdminOrClassScopedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrClassScopedAccess is a free data retrieval call binding the contract method 0x31098e24.
//
// Solidity: function onlyAdminOrClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyAdminOrClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrClassScopedAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminOrClassScopedAccess is a free data retrieval call binding the contract method 0x31098e24.
//
// Solidity: function onlyAdminOrClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyAdminOrClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrClassScopedAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminOrOwner is a free data retrieval call binding the contract method 0xd2df45d4.
//
// Solidity: function onlyAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyAdminOrOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyAdminOrOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrOwner is a free data retrieval call binding the contract method 0xd2df45d4.
//
// Solidity: function onlyAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminOrOwner is a free data retrieval call binding the contract method 0xd2df45d4.
//
// Solidity: function onlyAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminOrOwnerSupported is a free data retrieval call binding the contract method 0xf4682823.
//
// Solidity: function onlyAdminOrOwnerSupported(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyAdminOrOwnerSupported(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyAdminOrOwnerSupported", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrOwnerSupported is a free data retrieval call binding the contract method 0xf4682823.
//
// Solidity: function onlyAdminOrOwnerSupported(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyAdminOrOwnerSupported(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrOwnerSupported(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminOrOwnerSupported is a free data retrieval call binding the contract method 0xf4682823.
//
// Solidity: function onlyAdminOrOwnerSupported(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyAdminOrOwnerSupported(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrOwnerSupported(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminOrScopeEnforcedCall is a free data retrieval call binding the contract method 0xb65e6ac2.
//
// Solidity: function onlyAdminOrScopeEnforcedCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyAdminOrScopeEnforcedCall(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyAdminOrScopeEnforcedCall", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminOrScopeEnforcedCall is a free data retrieval call binding the contract method 0xb65e6ac2.
//
// Solidity: function onlyAdminOrScopeEnforcedCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyAdminOrScopeEnforcedCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrScopeEnforcedCall(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminOrScopeEnforcedCall is a free data retrieval call binding the contract method 0xb65e6ac2.
//
// Solidity: function onlyAdminOrScopeEnforcedCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyAdminOrScopeEnforcedCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminOrScopeEnforcedCall(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminSupportedAccess is a free data retrieval call binding the contract method 0x12fea21c.
//
// Solidity: function onlyAdminSupportedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyAdminSupportedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyAdminSupportedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminSupportedAccess is a free data retrieval call binding the contract method 0x12fea21c.
//
// Solidity: function onlyAdminSupportedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyAdminSupportedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminSupportedAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminSupportedAccess is a free data retrieval call binding the contract method 0x12fea21c.
//
// Solidity: function onlyAdminSupportedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyAdminSupportedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminSupportedAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminSupportedOwnerOrCall is a free data retrieval call binding the contract method 0x03e4c983.
//
// Solidity: function onlyAdminSupportedOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyAdminSupportedOwnerOrCall(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyAdminSupportedOwnerOrCall", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyAdminSupportedOwnerOrCall is a free data retrieval call binding the contract method 0x03e4c983.
//
// Solidity: function onlyAdminSupportedOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyAdminSupportedOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminSupportedOwnerOrCall(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyAdminSupportedOwnerOrCall is a free data retrieval call binding the contract method 0x03e4c983.
//
// Solidity: function onlyAdminSupportedOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyAdminSupportedOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyAdminSupportedOwnerOrCall(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyCallAccess is a free data retrieval call binding the contract method 0xa58a5393.
//
// Solidity: function onlyCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyCallAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyCallAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyCallAccess is a free data retrieval call binding the contract method 0xa58a5393.
//
// Solidity: function onlyCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyCallAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyCallAccess is a free data retrieval call binding the contract method 0xa58a5393.
//
// Solidity: function onlyCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyCallAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyCallAccessOrDirectEphemeralOwner is a free data retrieval call binding the contract method 0xee906526.
//
// Solidity: function onlyCallAccessOrDirectEphemeralOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyCallAccessOrDirectEphemeralOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyCallAccessOrDirectEphemeralOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyCallAccessOrDirectEphemeralOwner is a free data retrieval call binding the contract method 0xee906526.
//
// Solidity: function onlyCallAccessOrDirectEphemeralOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyCallAccessOrDirectEphemeralOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyCallAccessOrDirectEphemeralOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyCallAccessOrDirectEphemeralOwner is a free data retrieval call binding the contract method 0xee906526.
//
// Solidity: function onlyCallAccessOrDirectEphemeralOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyCallAccessOrDirectEphemeralOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyCallAccessOrDirectEphemeralOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyCallAccessWithScopeEnforced is a free data retrieval call binding the contract method 0xdde2f623.
//
// Solidity: function onlyCallAccessWithScopeEnforced(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyCallAccessWithScopeEnforced(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyCallAccessWithScopeEnforced", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyCallAccessWithScopeEnforced is a free data retrieval call binding the contract method 0xdde2f623.
//
// Solidity: function onlyCallAccessWithScopeEnforced(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyCallAccessWithScopeEnforced(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyCallAccessWithScopeEnforced(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyCallAccessWithScopeEnforced is a free data retrieval call binding the contract method 0xdde2f623.
//
// Solidity: function onlyCallAccessWithScopeEnforced(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyCallAccessWithScopeEnforced(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyCallAccessWithScopeEnforced(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyClassScopedAccess is a free data retrieval call binding the contract method 0xf78da887.
//
// Solidity: function onlyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyClassScopedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyClassScopedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyClassScopedAccess is a free data retrieval call binding the contract method 0xf78da887.
//
// Solidity: function onlyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyClassScopedAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyClassScopedAccess is a free data retrieval call binding the contract method 0xf78da887.
//
// Solidity: function onlyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyClassScopedAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyClassScopedOrCharAdminOrOwner is a free data retrieval call binding the contract method 0xcc8636d8.
//
// Solidity: function onlyClassScopedOrCharAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyClassScopedOrCharAdminOrOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyClassScopedOrCharAdminOrOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyClassScopedOrCharAdminOrOwner is a free data retrieval call binding the contract method 0xcc8636d8.
//
// Solidity: function onlyClassScopedOrCharAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyClassScopedOrCharAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyClassScopedOrCharAdminOrOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyClassScopedOrCharAdminOrOwner is a free data retrieval call binding the contract method 0xcc8636d8.
//
// Solidity: function onlyClassScopedOrCharAdminOrOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyClassScopedOrCharAdminOrOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyClassScopedOrCharAdminOrOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyDirectAdmin is a free data retrieval call binding the contract method 0x542b659d.
//
// Solidity: function onlyDirectAdmin(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyDirectAdmin(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyDirectAdmin", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyDirectAdmin is a free data retrieval call binding the contract method 0x542b659d.
//
// Solidity: function onlyDirectAdmin(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyDirectAdmin(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyDirectAdmin(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyDirectAdmin is a free data retrieval call binding the contract method 0x542b659d.
//
// Solidity: function onlyDirectAdmin(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyDirectAdmin(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyDirectAdmin(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyDirectAdminOrCallAccess is a free data retrieval call binding the contract method 0xd7fe3c1c.
//
// Solidity: function onlyDirectAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyDirectAdminOrCallAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyDirectAdminOrCallAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyDirectAdminOrCallAccess is a free data retrieval call binding the contract method 0xd7fe3c1c.
//
// Solidity: function onlyDirectAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyDirectAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyDirectAdminOrCallAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyDirectAdminOrCallAccess is a free data retrieval call binding the contract method 0xd7fe3c1c.
//
// Solidity: function onlyDirectAdminOrCallAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyDirectAdminOrCallAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyDirectAdminOrCallAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyDirectEphemeralOwnerOrCall is a free data retrieval call binding the contract method 0xf10f61ca.
//
// Solidity: function onlyDirectEphemeralOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyDirectEphemeralOwnerOrCall(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyDirectEphemeralOwnerOrCall", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyDirectEphemeralOwnerOrCall is a free data retrieval call binding the contract method 0xf10f61ca.
//
// Solidity: function onlyDirectEphemeralOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyDirectEphemeralOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyDirectEphemeralOwnerOrCall(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyDirectEphemeralOwnerOrCall is a free data retrieval call binding the contract method 0xf10f61ca.
//
// Solidity: function onlyDirectEphemeralOwnerOrCall(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyDirectEphemeralOwnerOrCall(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyDirectEphemeralOwnerOrCall(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyDirectOwner is a free data retrieval call binding the contract method 0x3ac45a0c.
//
// Solidity: function onlyDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyDirectOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyDirectOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyDirectOwner is a free data retrieval call binding the contract method 0x3ac45a0c.
//
// Solidity: function onlyDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyDirectOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyDirectOwner is a free data retrieval call binding the contract method 0x3ac45a0c.
//
// Solidity: function onlyDirectOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyDirectOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyDirectOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyEphemeralOwnerOrTransferRole is a free data retrieval call binding the contract method 0x5eee0cc8.
//
// Solidity: function onlyEphemeralOwnerOrTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyEphemeralOwnerOrTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyEphemeralOwnerOrTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyEphemeralOwnerOrTransferRole is a free data retrieval call binding the contract method 0x5eee0cc8.
//
// Solidity: function onlyEphemeralOwnerOrTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyEphemeralOwnerOrTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyEphemeralOwnerOrTransferRole(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyEphemeralOwnerOrTransferRole is a free data retrieval call binding the contract method 0x5eee0cc8.
//
// Solidity: function onlyEphemeralOwnerOrTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyEphemeralOwnerOrTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyEphemeralOwnerOrTransferRole(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwner is a free data retrieval call binding the contract method 0x4ded2d51.
//
// Solidity: function onlyOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyOwner(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyOwner", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwner is a free data retrieval call binding the contract method 0x4ded2d51.
//
// Solidity: function onlyOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwner is a free data retrieval call binding the contract method 0x4ded2d51.
//
// Solidity: function onlyOwner(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyOwner(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwner(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrEphemeralCrossTransferRole is a free data retrieval call binding the contract method 0xd061b071.
//
// Solidity: function onlyOwnerOrEphemeralCrossTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyOwnerOrEphemeralCrossTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyOwnerOrEphemeralCrossTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwnerOrEphemeralCrossTransferRole is a free data retrieval call binding the contract method 0xd061b071.
//
// Solidity: function onlyOwnerOrEphemeralCrossTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyOwnerOrEphemeralCrossTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwnerOrEphemeralCrossTransferRole(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrEphemeralCrossTransferRole is a free data retrieval call binding the contract method 0xd061b071.
//
// Solidity: function onlyOwnerOrEphemeralCrossTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyOwnerOrEphemeralCrossTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwnerOrEphemeralCrossTransferRole(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrEphemeralTransferRole is a free data retrieval call binding the contract method 0x2227febd.
//
// Solidity: function onlyOwnerOrEphemeralTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyOwnerOrEphemeralTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyOwnerOrEphemeralTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwnerOrEphemeralTransferRole is a free data retrieval call binding the contract method 0x2227febd.
//
// Solidity: function onlyOwnerOrEphemeralTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyOwnerOrEphemeralTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwnerOrEphemeralTransferRole(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrEphemeralTransferRole is a free data retrieval call binding the contract method 0x2227febd.
//
// Solidity: function onlyOwnerOrEphemeralTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyOwnerOrEphemeralTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwnerOrEphemeralTransferRole(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrInventoryTransferRole is a free data retrieval call binding the contract method 0x992518f9.
//
// Solidity: function onlyOwnerOrInventoryTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyOwnerOrInventoryTransferRole(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyOwnerOrInventoryTransferRole", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwnerOrInventoryTransferRole is a free data retrieval call binding the contract method 0x992518f9.
//
// Solidity: function onlyOwnerOrInventoryTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyOwnerOrInventoryTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwnerOrInventoryTransferRole(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwnerOrInventoryTransferRole is a free data retrieval call binding the contract method 0x992518f9.
//
// Solidity: function onlyOwnerOrInventoryTransferRole(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyOwnerOrInventoryTransferRole(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwnerOrInventoryTransferRole(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwnerWithAdminSupportAccess is a free data retrieval call binding the contract method 0x036867ec.
//
// Solidity: function onlyOwnerWithAdminSupportAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlyOwnerWithAdminSupportAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlyOwnerWithAdminSupportAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlyOwnerWithAdminSupportAccess is a free data retrieval call binding the contract method 0x036867ec.
//
// Solidity: function onlyOwnerWithAdminSupportAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlyOwnerWithAdminSupportAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwnerWithAdminSupportAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlyOwnerWithAdminSupportAccess is a free data retrieval call binding the contract method 0x036867ec.
//
// Solidity: function onlyOwnerWithAdminSupportAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlyOwnerWithAdminSupportAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlyOwnerWithAdminSupportAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlySmartAssemblyClassScopedAccess is a free data retrieval call binding the contract method 0xbccd5c2a.
//
// Solidity: function onlySmartAssemblyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCaller) OnlySmartAssemblyClassScopedAccess(opts *bind.CallOpts, smartObjectId *big.Int, data []byte) error {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "onlySmartAssemblyClassScopedAccess", smartObjectId, data)

	if err != nil {
		return err
	}

	return err

}

// OnlySmartAssemblyClassScopedAccess is a free data retrieval call binding the contract method 0xbccd5c2a.
//
// Solidity: function onlySmartAssemblyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnlySmartAssemblyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlySmartAssemblyClassScopedAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// OnlySmartAssemblyClassScopedAccess is a free data retrieval call binding the contract method 0xbccd5c2a.
//
// Solidity: function onlySmartAssemblyClassScopedAccess(uint256 smartObjectId, bytes data) view returns()
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) OnlySmartAssemblyClassScopedAccess(smartObjectId *big.Int, data []byte) error {
	return _ERC2771Forwarder.Contract.OnlySmartAssemblyClassScopedAccess(&_ERC2771Forwarder.CallOpts, smartObjectId, data)
}

// Owner is a free data retrieval call binding the contract method 0xa123c33e.
//
// Solidity: function owner(uint256 smartObjectId) view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) Owner(opts *bind.CallOpts, smartObjectId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "owner", smartObjectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xa123c33e.
//
// Solidity: function owner(uint256 smartObjectId) view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderSession) Owner(smartObjectId *big.Int) (common.Address, error) {
	return _ERC2771Forwarder.Contract.Owner(&_ERC2771Forwarder.CallOpts, smartObjectId)
}

// Owner is a free data retrieval call binding the contract method 0xa123c33e.
//
// Solidity: function owner(uint256 smartObjectId) view returns(address)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) Owner(smartObjectId *big.Int) (common.Address, error) {
	return _ERC2771Forwarder.Contract.Owner(&_ERC2771Forwarder.CallOpts, smartObjectId)
}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) StoreVersion(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "storeVersion")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_ERC2771Forwarder *ERC2771ForwarderSession) StoreVersion() ([32]byte, error) {
	return _ERC2771Forwarder.Contract.StoreVersion(&_ERC2771Forwarder.CallOpts)
}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) StoreVersion() ([32]byte, error) {
	return _ERC2771Forwarder.Contract.StoreVersion(&_ERC2771Forwarder.CallOpts)
}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) WorldVersion(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "worldVersion")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_ERC2771Forwarder *ERC2771ForwarderSession) WorldVersion() ([32]byte, error) {
	return _ERC2771Forwarder.Contract.WorldVersion(&_ERC2771Forwarder.CallOpts)
}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) WorldVersion() ([32]byte, error) {
	return _ERC2771Forwarder.Contract.WorldVersion(&_ERC2771Forwarder.CallOpts)
}

// Aggression is a paid mutator transaction binding the contract method 0x8ba2b5d8.
//
// Solidity: function aggression((uint256,((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[],(uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256)) params) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) Aggression(opts *bind.TransactOpts, params AggressionParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "aggression", params)
}

// Aggression is a paid mutator transaction binding the contract method 0x8ba2b5d8.
//
// Solidity: function aggression((uint256,((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[],(uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256)) params) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771Forwarder *ERC2771ForwarderSession) Aggression(params AggressionParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Aggression(&_ERC2771Forwarder.TransactOpts, params)
}

// Aggression is a paid mutator transaction binding the contract method 0x8ba2b5d8.
//
// Solidity: function aggression((uint256,((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[],(uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256)) params) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) Aggression(params AggressionParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Aggression(&_ERC2771Forwarder.TransactOpts, params)
}

// Anchor is a paid mutator transaction binding the contract method 0x4ccd479d.
//
// Solidity: function anchor(uint256 smartObjectId, address owner, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) Anchor(opts *bind.TransactOpts, smartObjectId *big.Int, owner common.Address, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "anchor", smartObjectId, owner, locationData)
}

// Anchor is a paid mutator transaction binding the contract method 0x4ccd479d.
//
// Solidity: function anchor(uint256 smartObjectId, address owner, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) Anchor(smartObjectId *big.Int, owner common.Address, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Anchor(&_ERC2771Forwarder.TransactOpts, smartObjectId, owner, locationData)
}

// Anchor is a paid mutator transaction binding the contract method 0x4ccd479d.
//
// Solidity: function anchor(uint256 smartObjectId, address owner, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) Anchor(smartObjectId *big.Int, owner common.Address, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Anchor(&_ERC2771Forwarder.TransactOpts, smartObjectId, owner, locationData)
}

// AssignItemToInventory is a paid mutator transaction binding the contract method 0xe2957099.
//
// Solidity: function assignItemToInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) AssignItemToInventory(opts *bind.TransactOpts, inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "assignItemToInventory", inventoryObjectId, itemObjectId, quantity)
}

// AssignItemToInventory is a paid mutator transaction binding the contract method 0xe2957099.
//
// Solidity: function assignItemToInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) AssignItemToInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.AssignItemToInventory(&_ERC2771Forwarder.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// AssignItemToInventory is a paid mutator transaction binding the contract method 0xe2957099.
//
// Solidity: function assignItemToInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) AssignItemToInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.AssignItemToInventory(&_ERC2771Forwarder.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// AssignOwner is a paid mutator transaction binding the contract method 0xe671644b.
//
// Solidity: function assignOwner(uint256 smartObjectId, address to) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) AssignOwner(opts *bind.TransactOpts, smartObjectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "assignOwner", smartObjectId, to)
}

// AssignOwner is a paid mutator transaction binding the contract method 0xe671644b.
//
// Solidity: function assignOwner(uint256 smartObjectId, address to) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) AssignOwner(smartObjectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.AssignOwner(&_ERC2771Forwarder.TransactOpts, smartObjectId, to)
}

// AssignOwner is a paid mutator transaction binding the contract method 0xe671644b.
//
// Solidity: function assignOwner(uint256 smartObjectId, address to) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) AssignOwner(smartObjectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.AssignOwner(&_ERC2771Forwarder.TransactOpts, smartObjectId, to)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) BatchCall(opts *bind.TransactOpts, systemCalls []SystemCallData) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "batchCall", systemCalls)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771Forwarder *ERC2771ForwarderSession) BatchCall(systemCalls []SystemCallData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.BatchCall(&_ERC2771Forwarder.TransactOpts, systemCalls)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) BatchCall(systemCalls []SystemCallData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.BatchCall(&_ERC2771Forwarder.TransactOpts, systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) BatchCallFrom(opts *bind.TransactOpts, systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "batchCallFrom", systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771Forwarder *ERC2771ForwarderSession) BatchCallFrom(systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.BatchCallFrom(&_ERC2771Forwarder.TransactOpts, systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) BatchCallFrom(systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.BatchCallFrom(&_ERC2771Forwarder.TransactOpts, systemCalls)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) BringOffline(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "bringOffline", smartObjectId)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) BringOffline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.BringOffline(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) BringOffline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.BringOffline(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) BringOnline(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "bringOnline", smartObjectId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) BringOnline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.BringOnline(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) BringOnline(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.BringOnline(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) Call(opts *bind.TransactOpts, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "call", systemId, callData)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771Forwarder *ERC2771ForwarderSession) Call(systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Call(&_ERC2771Forwarder.TransactOpts, systemId, callData)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) Call(systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Call(&_ERC2771Forwarder.TransactOpts, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CallFrom(opts *bind.TransactOpts, delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "callFrom", delegator, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771Forwarder *ERC2771ForwarderSession) CallFrom(delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CallFrom(&_ERC2771Forwarder.TransactOpts, delegator, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CallFrom(delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CallFrom(&_ERC2771Forwarder.TransactOpts, delegator, systemId, callData)
}

// CanJump is a paid mutator transaction binding the contract method 0xc9cf8ac5.
//
// Solidity: function canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CanJump(opts *bind.TransactOpts, characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "canJump", characterId, sourceGateId, destinationGateId)
}

// CanJump is a paid mutator transaction binding the contract method 0xc9cf8ac5.
//
// Solidity: function canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) CanJump(characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CanJump(&_ERC2771Forwarder.TransactOpts, characterId, sourceGateId, destinationGateId)
}

// CanJump is a paid mutator transaction binding the contract method 0xc9cf8ac5.
//
// Solidity: function canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CanJump(characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CanJump(&_ERC2771Forwarder.TransactOpts, characterId, sourceGateId, destinationGateId)
}

// ConfigureDeployableAccess is a paid mutator transaction binding the contract method 0x03d05669.
//
// Solidity: function configureDeployableAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureDeployableAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureDeployableAccess")
}

// ConfigureDeployableAccess is a paid mutator transaction binding the contract method 0x03d05669.
//
// Solidity: function configureDeployableAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureDeployableAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureDeployableAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureDeployableAccess is a paid mutator transaction binding the contract method 0x03d05669.
//
// Solidity: function configureDeployableAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureDeployableAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureDeployableAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureEntityRecordAccess is a paid mutator transaction binding the contract method 0xc36c2700.
//
// Solidity: function configureEntityRecordAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureEntityRecordAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureEntityRecordAccess")
}

// ConfigureEntityRecordAccess is a paid mutator transaction binding the contract method 0xc36c2700.
//
// Solidity: function configureEntityRecordAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureEntityRecordAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureEntityRecordAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureEntityRecordAccess is a paid mutator transaction binding the contract method 0xc36c2700.
//
// Solidity: function configureEntityRecordAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureEntityRecordAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureEntityRecordAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureEphemeralInteractAccess is a paid mutator transaction binding the contract method 0x1a3be1ed.
//
// Solidity: function configureEphemeralInteractAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureEphemeralInteractAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureEphemeralInteractAccess")
}

// ConfigureEphemeralInteractAccess is a paid mutator transaction binding the contract method 0x1a3be1ed.
//
// Solidity: function configureEphemeralInteractAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureEphemeralInteractAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureEphemeralInteractAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureEphemeralInteractAccess is a paid mutator transaction binding the contract method 0x1a3be1ed.
//
// Solidity: function configureEphemeralInteractAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureEphemeralInteractAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureEphemeralInteractAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureEphemeralInventoryAccess is a paid mutator transaction binding the contract method 0xc8367523.
//
// Solidity: function configureEphemeralInventoryAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureEphemeralInventoryAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureEphemeralInventoryAccess")
}

// ConfigureEphemeralInventoryAccess is a paid mutator transaction binding the contract method 0xc8367523.
//
// Solidity: function configureEphemeralInventoryAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureEphemeralInventoryAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureEphemeralInventoryAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureEphemeralInventoryAccess is a paid mutator transaction binding the contract method 0xc8367523.
//
// Solidity: function configureEphemeralInventoryAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureEphemeralInventoryAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureEphemeralInventoryAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureFuelAccess is a paid mutator transaction binding the contract method 0xde941597.
//
// Solidity: function configureFuelAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureFuelAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureFuelAccess")
}

// ConfigureFuelAccess is a paid mutator transaction binding the contract method 0xde941597.
//
// Solidity: function configureFuelAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureFuelAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureFuelAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureFuelAccess is a paid mutator transaction binding the contract method 0xde941597.
//
// Solidity: function configureFuelAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureFuelAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureFuelAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureFuelEfficiency is a paid mutator transaction binding the contract method 0x98c79b88.
//
// Solidity: function configureFuelEfficiency(uint256 fuelTypeId, uint256 fuelEfficiency) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureFuelEfficiency(opts *bind.TransactOpts, fuelTypeId *big.Int, fuelEfficiency *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureFuelEfficiency", fuelTypeId, fuelEfficiency)
}

// ConfigureFuelEfficiency is a paid mutator transaction binding the contract method 0x98c79b88.
//
// Solidity: function configureFuelEfficiency(uint256 fuelTypeId, uint256 fuelEfficiency) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureFuelEfficiency(fuelTypeId *big.Int, fuelEfficiency *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureFuelEfficiency(&_ERC2771Forwarder.TransactOpts, fuelTypeId, fuelEfficiency)
}

// ConfigureFuelEfficiency is a paid mutator transaction binding the contract method 0x98c79b88.
//
// Solidity: function configureFuelEfficiency(uint256 fuelTypeId, uint256 fuelEfficiency) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureFuelEfficiency(fuelTypeId *big.Int, fuelEfficiency *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureFuelEfficiency(&_ERC2771Forwarder.TransactOpts, fuelTypeId, fuelEfficiency)
}

// ConfigureFuelParameters is a paid mutator transaction binding the contract method 0xdcd62506.
//
// Solidity: function configureFuelParameters(uint256 smartObjectId, (uint256,uint256,uint256,uint256,uint256) fuelParams) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureFuelParameters(opts *bind.TransactOpts, smartObjectId *big.Int, fuelParams FuelParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureFuelParameters", smartObjectId, fuelParams)
}

// ConfigureFuelParameters is a paid mutator transaction binding the contract method 0xdcd62506.
//
// Solidity: function configureFuelParameters(uint256 smartObjectId, (uint256,uint256,uint256,uint256,uint256) fuelParams) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureFuelParameters(smartObjectId *big.Int, fuelParams FuelParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureFuelParameters(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelParams)
}

// ConfigureFuelParameters is a paid mutator transaction binding the contract method 0xdcd62506.
//
// Solidity: function configureFuelParameters(uint256 smartObjectId, (uint256,uint256,uint256,uint256,uint256) fuelParams) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureFuelParameters(smartObjectId *big.Int, fuelParams FuelParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureFuelParameters(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelParams)
}

// ConfigureGate is a paid mutator transaction binding the contract method 0xe6611746.
//
// Solidity: function configureGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureGate(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureGate", smartObjectId, systemId)
}

// ConfigureGate is a paid mutator transaction binding the contract method 0xe6611746.
//
// Solidity: function configureGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureGate(&_ERC2771Forwarder.TransactOpts, smartObjectId, systemId)
}

// ConfigureGate is a paid mutator transaction binding the contract method 0xe6611746.
//
// Solidity: function configureGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureGate(&_ERC2771Forwarder.TransactOpts, smartObjectId, systemId)
}

// ConfigureInventoryAccess is a paid mutator transaction binding the contract method 0xb47cb8f3.
//
// Solidity: function configureInventoryAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureInventoryAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureInventoryAccess")
}

// ConfigureInventoryAccess is a paid mutator transaction binding the contract method 0xb47cb8f3.
//
// Solidity: function configureInventoryAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureInventoryAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureInventoryAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureInventoryAccess is a paid mutator transaction binding the contract method 0xb47cb8f3.
//
// Solidity: function configureInventoryAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureInventoryAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureInventoryAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureInventoryInteractAccess is a paid mutator transaction binding the contract method 0x3884c463.
//
// Solidity: function configureInventoryInteractAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureInventoryInteractAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureInventoryInteractAccess")
}

// ConfigureInventoryInteractAccess is a paid mutator transaction binding the contract method 0x3884c463.
//
// Solidity: function configureInventoryInteractAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureInventoryInteractAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureInventoryInteractAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureInventoryInteractAccess is a paid mutator transaction binding the contract method 0x3884c463.
//
// Solidity: function configureInventoryInteractAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureInventoryInteractAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureInventoryInteractAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureKillMailAccess is a paid mutator transaction binding the contract method 0x752918af.
//
// Solidity: function configureKillMailAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureKillMailAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureKillMailAccess")
}

// ConfigureKillMailAccess is a paid mutator transaction binding the contract method 0x752918af.
//
// Solidity: function configureKillMailAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureKillMailAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureKillMailAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureKillMailAccess is a paid mutator transaction binding the contract method 0x752918af.
//
// Solidity: function configureKillMailAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureKillMailAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureKillMailAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureLocationAccess is a paid mutator transaction binding the contract method 0x520f13da.
//
// Solidity: function configureLocationAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureLocationAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureLocationAccess")
}

// ConfigureLocationAccess is a paid mutator transaction binding the contract method 0x520f13da.
//
// Solidity: function configureLocationAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureLocationAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureLocationAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureLocationAccess is a paid mutator transaction binding the contract method 0x520f13da.
//
// Solidity: function configureLocationAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureLocationAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureLocationAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureNetworkNodeAccess is a paid mutator transaction binding the contract method 0x38de6231.
//
// Solidity: function configureNetworkNodeAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureNetworkNodeAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureNetworkNodeAccess")
}

// ConfigureNetworkNodeAccess is a paid mutator transaction binding the contract method 0x38de6231.
//
// Solidity: function configureNetworkNodeAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureNetworkNodeAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureNetworkNodeAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureNetworkNodeAccess is a paid mutator transaction binding the contract method 0x38de6231.
//
// Solidity: function configureNetworkNodeAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureNetworkNodeAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureNetworkNodeAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureOwnershipAccess is a paid mutator transaction binding the contract method 0x4eee463a.
//
// Solidity: function configureOwnershipAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureOwnershipAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureOwnershipAccess")
}

// ConfigureOwnershipAccess is a paid mutator transaction binding the contract method 0x4eee463a.
//
// Solidity: function configureOwnershipAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureOwnershipAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureOwnershipAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureOwnershipAccess is a paid mutator transaction binding the contract method 0x4eee463a.
//
// Solidity: function configureOwnershipAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureOwnershipAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureOwnershipAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartAssemblyAccess is a paid mutator transaction binding the contract method 0xafa610cb.
//
// Solidity: function configureSmartAssemblyAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureSmartAssemblyAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureSmartAssemblyAccess")
}

// ConfigureSmartAssemblyAccess is a paid mutator transaction binding the contract method 0xafa610cb.
//
// Solidity: function configureSmartAssemblyAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureSmartAssemblyAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartAssemblyAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartAssemblyAccess is a paid mutator transaction binding the contract method 0xafa610cb.
//
// Solidity: function configureSmartAssemblyAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureSmartAssemblyAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartAssemblyAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartCharacterAccess is a paid mutator transaction binding the contract method 0xa2c4eac8.
//
// Solidity: function configureSmartCharacterAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureSmartCharacterAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureSmartCharacterAccess")
}

// ConfigureSmartCharacterAccess is a paid mutator transaction binding the contract method 0xa2c4eac8.
//
// Solidity: function configureSmartCharacterAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureSmartCharacterAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartCharacterAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartCharacterAccess is a paid mutator transaction binding the contract method 0xa2c4eac8.
//
// Solidity: function configureSmartCharacterAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureSmartCharacterAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartCharacterAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartGateAccess is a paid mutator transaction binding the contract method 0x6238f85c.
//
// Solidity: function configureSmartGateAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureSmartGateAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureSmartGateAccess")
}

// ConfigureSmartGateAccess is a paid mutator transaction binding the contract method 0x6238f85c.
//
// Solidity: function configureSmartGateAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureSmartGateAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartGateAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartGateAccess is a paid mutator transaction binding the contract method 0x6238f85c.
//
// Solidity: function configureSmartGateAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureSmartGateAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartGateAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartStorageUnitAccess is a paid mutator transaction binding the contract method 0x647dc438.
//
// Solidity: function configureSmartStorageUnitAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureSmartStorageUnitAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureSmartStorageUnitAccess")
}

// ConfigureSmartStorageUnitAccess is a paid mutator transaction binding the contract method 0x647dc438.
//
// Solidity: function configureSmartStorageUnitAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureSmartStorageUnitAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartStorageUnitAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartStorageUnitAccess is a paid mutator transaction binding the contract method 0x647dc438.
//
// Solidity: function configureSmartStorageUnitAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureSmartStorageUnitAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartStorageUnitAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartTurretAccess is a paid mutator transaction binding the contract method 0x232dba5b.
//
// Solidity: function configureSmartTurretAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureSmartTurretAccess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureSmartTurretAccess")
}

// ConfigureSmartTurretAccess is a paid mutator transaction binding the contract method 0x232dba5b.
//
// Solidity: function configureSmartTurretAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureSmartTurretAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartTurretAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureSmartTurretAccess is a paid mutator transaction binding the contract method 0x232dba5b.
//
// Solidity: function configureSmartTurretAccess() returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureSmartTurretAccess() (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureSmartTurretAccess(&_ERC2771Forwarder.TransactOpts)
}

// ConfigureTurret is a paid mutator transaction binding the contract method 0x4d7e3af1.
//
// Solidity: function configureTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConfigureTurret(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "configureTurret", smartObjectId, systemId)
}

// ConfigureTurret is a paid mutator transaction binding the contract method 0x4d7e3af1.
//
// Solidity: function configureTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConfigureTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureTurret(&_ERC2771Forwarder.TransactOpts, smartObjectId, systemId)
}

// ConfigureTurret is a paid mutator transaction binding the contract method 0x4d7e3af1.
//
// Solidity: function configureTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConfigureTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConfigureTurret(&_ERC2771Forwarder.TransactOpts, smartObjectId, systemId)
}

// ConnectStructure is a paid mutator transaction binding the contract method 0x61e262e4.
//
// Solidity: function connectStructure(uint256 networkNodeId, uint256 structureId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ConnectStructure(opts *bind.TransactOpts, networkNodeId *big.Int, structureId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "connectStructure", networkNodeId, structureId)
}

// ConnectStructure is a paid mutator transaction binding the contract method 0x61e262e4.
//
// Solidity: function connectStructure(uint256 networkNodeId, uint256 structureId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ConnectStructure(networkNodeId *big.Int, structureId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConnectStructure(&_ERC2771Forwarder.TransactOpts, networkNodeId, structureId)
}

// ConnectStructure is a paid mutator transaction binding the contract method 0x61e262e4.
//
// Solidity: function connectStructure(uint256 networkNodeId, uint256 structureId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ConnectStructure(networkNodeId *big.Int, structureId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ConnectStructure(&_ERC2771Forwarder.TransactOpts, networkNodeId, structureId)
}

// CreateAndAnchor is a paid mutator transaction binding the contract method 0x488ea387.
//
// Solidity: function createAndAnchor((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateAndAnchor(opts *bind.TransactOpts, params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createAndAnchor", params, networkNodeId)
}

// CreateAndAnchor is a paid mutator transaction binding the contract method 0x488ea387.
//
// Solidity: function createAndAnchor((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateAndAnchor(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchor(&_ERC2771Forwarder.TransactOpts, params, networkNodeId)
}

// CreateAndAnchor is a paid mutator transaction binding the contract method 0x488ea387.
//
// Solidity: function createAndAnchor((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateAndAnchor(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchor(&_ERC2771Forwarder.TransactOpts, params, networkNodeId)
}

// CreateAndAnchorGate is a paid mutator transaction binding the contract method 0xbb054d08.
//
// Solidity: function createAndAnchorGate((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 maxDistance, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateAndAnchorGate(opts *bind.TransactOpts, params CreateAndAnchorParams, maxDistance *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createAndAnchorGate", params, maxDistance, networkNodeId)
}

// CreateAndAnchorGate is a paid mutator transaction binding the contract method 0xbb054d08.
//
// Solidity: function createAndAnchorGate((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 maxDistance, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateAndAnchorGate(params CreateAndAnchorParams, maxDistance *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchorGate(&_ERC2771Forwarder.TransactOpts, params, maxDistance, networkNodeId)
}

// CreateAndAnchorGate is a paid mutator transaction binding the contract method 0xbb054d08.
//
// Solidity: function createAndAnchorGate((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 maxDistance, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateAndAnchorGate(params CreateAndAnchorParams, maxDistance *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchorGate(&_ERC2771Forwarder.TransactOpts, params, maxDistance, networkNodeId)
}

// CreateAndAnchorNetworkNode is a paid mutator transaction binding the contract method 0x29c2970e.
//
// Solidity: function createAndAnchorNetworkNode((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, (uint256,uint256,uint256,uint256,uint256) fuelParams, uint256 maxEnergyCapacity, uint256 currentProduction) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateAndAnchorNetworkNode(opts *bind.TransactOpts, params CreateAndAnchorParams, fuelParams FuelParams, maxEnergyCapacity *big.Int, currentProduction *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createAndAnchorNetworkNode", params, fuelParams, maxEnergyCapacity, currentProduction)
}

// CreateAndAnchorNetworkNode is a paid mutator transaction binding the contract method 0x29c2970e.
//
// Solidity: function createAndAnchorNetworkNode((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, (uint256,uint256,uint256,uint256,uint256) fuelParams, uint256 maxEnergyCapacity, uint256 currentProduction) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateAndAnchorNetworkNode(params CreateAndAnchorParams, fuelParams FuelParams, maxEnergyCapacity *big.Int, currentProduction *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchorNetworkNode(&_ERC2771Forwarder.TransactOpts, params, fuelParams, maxEnergyCapacity, currentProduction)
}

// CreateAndAnchorNetworkNode is a paid mutator transaction binding the contract method 0x29c2970e.
//
// Solidity: function createAndAnchorNetworkNode((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, (uint256,uint256,uint256,uint256,uint256) fuelParams, uint256 maxEnergyCapacity, uint256 currentProduction) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateAndAnchorNetworkNode(params CreateAndAnchorParams, fuelParams FuelParams, maxEnergyCapacity *big.Int, currentProduction *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchorNetworkNode(&_ERC2771Forwarder.TransactOpts, params, fuelParams, maxEnergyCapacity, currentProduction)
}

// CreateAndAnchorStorageUnit is a paid mutator transaction binding the contract method 0xbf451df6.
//
// Solidity: function createAndAnchorStorageUnit((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 capacity, uint256 ephemeralCapacity, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateAndAnchorStorageUnit(opts *bind.TransactOpts, params CreateAndAnchorParams, capacity *big.Int, ephemeralCapacity *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createAndAnchorStorageUnit", params, capacity, ephemeralCapacity, networkNodeId)
}

// CreateAndAnchorStorageUnit is a paid mutator transaction binding the contract method 0xbf451df6.
//
// Solidity: function createAndAnchorStorageUnit((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 capacity, uint256 ephemeralCapacity, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateAndAnchorStorageUnit(params CreateAndAnchorParams, capacity *big.Int, ephemeralCapacity *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchorStorageUnit(&_ERC2771Forwarder.TransactOpts, params, capacity, ephemeralCapacity, networkNodeId)
}

// CreateAndAnchorStorageUnit is a paid mutator transaction binding the contract method 0xbf451df6.
//
// Solidity: function createAndAnchorStorageUnit((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 capacity, uint256 ephemeralCapacity, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateAndAnchorStorageUnit(params CreateAndAnchorParams, capacity *big.Int, ephemeralCapacity *big.Int, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchorStorageUnit(&_ERC2771Forwarder.TransactOpts, params, capacity, ephemeralCapacity, networkNodeId)
}

// CreateAndAnchorTurret is a paid mutator transaction binding the contract method 0xb209af27.
//
// Solidity: function createAndAnchorTurret((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateAndAnchorTurret(opts *bind.TransactOpts, params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createAndAnchorTurret", params, networkNodeId)
}

// CreateAndAnchorTurret is a paid mutator transaction binding the contract method 0xb209af27.
//
// Solidity: function createAndAnchorTurret((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateAndAnchorTurret(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchorTurret(&_ERC2771Forwarder.TransactOpts, params, networkNodeId)
}

// CreateAndAnchorTurret is a paid mutator transaction binding the contract method 0xb209af27.
//
// Solidity: function createAndAnchorTurret((uint256,string,(bytes32,uint256,uint256,uint256),address,(uint256,uint256,uint256,uint256)) params, uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateAndAnchorTurret(params CreateAndAnchorParams, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndAnchorTurret(&_ERC2771Forwarder.TransactOpts, params, networkNodeId)
}

// CreateAndDepositEphemeral is a paid mutator transaction binding the contract method 0x0f33de24.
//
// Solidity: function createAndDepositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateAndDepositEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createAndDepositEphemeral", smartObjectId, ephemeralOwner, items)
}

// CreateAndDepositEphemeral is a paid mutator transaction binding the contract method 0x0f33de24.
//
// Solidity: function createAndDepositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateAndDepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndDepositEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// CreateAndDepositEphemeral is a paid mutator transaction binding the contract method 0x0f33de24.
//
// Solidity: function createAndDepositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateAndDepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndDepositEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// CreateAndDepositInventory is a paid mutator transaction binding the contract method 0x087ad59f.
//
// Solidity: function createAndDepositInventory(uint256 smartObjectId, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateAndDepositInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createAndDepositInventory", smartObjectId, items)
}

// CreateAndDepositInventory is a paid mutator transaction binding the contract method 0x087ad59f.
//
// Solidity: function createAndDepositInventory(uint256 smartObjectId, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateAndDepositInventory(smartObjectId *big.Int, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndDepositInventory(&_ERC2771Forwarder.TransactOpts, smartObjectId, items)
}

// CreateAndDepositInventory is a paid mutator transaction binding the contract method 0x087ad59f.
//
// Solidity: function createAndDepositInventory(uint256 smartObjectId, (uint256,bytes32,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateAndDepositInventory(smartObjectId *big.Int, items []CreateInventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAndDepositInventory(&_ERC2771Forwarder.TransactOpts, smartObjectId, items)
}

// CreateAssembly is a paid mutator transaction binding the contract method 0xc9f4c7c2.
//
// Solidity: function createAssembly(uint256 smartObjectId, string assemblyType, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateAssembly(opts *bind.TransactOpts, smartObjectId *big.Int, assemblyType string, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createAssembly", smartObjectId, assemblyType, entityRecordParams)
}

// CreateAssembly is a paid mutator transaction binding the contract method 0xc9f4c7c2.
//
// Solidity: function createAssembly(uint256 smartObjectId, string assemblyType, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateAssembly(smartObjectId *big.Int, assemblyType string, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAssembly(&_ERC2771Forwarder.TransactOpts, smartObjectId, assemblyType, entityRecordParams)
}

// CreateAssembly is a paid mutator transaction binding the contract method 0xc9f4c7c2.
//
// Solidity: function createAssembly(uint256 smartObjectId, string assemblyType, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateAssembly(smartObjectId *big.Int, assemblyType string, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateAssembly(&_ERC2771Forwarder.TransactOpts, smartObjectId, assemblyType, entityRecordParams)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x6d0030ca.
//
// Solidity: function createCharacter(uint256 smartObjectId, address owner, uint256 tribeId, (bytes32,uint256,uint256,uint256) entityRecordParams, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateCharacter(opts *bind.TransactOpts, smartObjectId *big.Int, owner common.Address, tribeId *big.Int, entityRecordParams EntityRecordParams, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createCharacter", smartObjectId, owner, tribeId, entityRecordParams, entityRecordMetadata)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x6d0030ca.
//
// Solidity: function createCharacter(uint256 smartObjectId, address owner, uint256 tribeId, (bytes32,uint256,uint256,uint256) entityRecordParams, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateCharacter(smartObjectId *big.Int, owner common.Address, tribeId *big.Int, entityRecordParams EntityRecordParams, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateCharacter(&_ERC2771Forwarder.TransactOpts, smartObjectId, owner, tribeId, entityRecordParams, entityRecordMetadata)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x6d0030ca.
//
// Solidity: function createCharacter(uint256 smartObjectId, address owner, uint256 tribeId, (bytes32,uint256,uint256,uint256) entityRecordParams, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateCharacter(smartObjectId *big.Int, owner common.Address, tribeId *big.Int, entityRecordParams EntityRecordParams, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateCharacter(&_ERC2771Forwarder.TransactOpts, smartObjectId, owner, tribeId, entityRecordParams, entityRecordMetadata)
}

// CreateDeployable is a paid mutator transaction binding the contract method 0xcadee21c.
//
// Solidity: function createDeployable(uint256 smartObjectId, address owner) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateDeployable(opts *bind.TransactOpts, smartObjectId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createDeployable", smartObjectId, owner)
}

// CreateDeployable is a paid mutator transaction binding the contract method 0xcadee21c.
//
// Solidity: function createDeployable(uint256 smartObjectId, address owner) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateDeployable(smartObjectId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateDeployable(&_ERC2771Forwarder.TransactOpts, smartObjectId, owner)
}

// CreateDeployable is a paid mutator transaction binding the contract method 0xcadee21c.
//
// Solidity: function createDeployable(uint256 smartObjectId, address owner) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateDeployable(smartObjectId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateDeployable(&_ERC2771Forwarder.TransactOpts, smartObjectId, owner)
}

// CreateMetadata is a paid mutator transaction binding the contract method 0x21823a2a.
//
// Solidity: function createMetadata(uint256 smartObjectId, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateMetadata(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createMetadata", smartObjectId, entityRecordMetadata)
}

// CreateMetadata is a paid mutator transaction binding the contract method 0x21823a2a.
//
// Solidity: function createMetadata(uint256 smartObjectId, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateMetadata(smartObjectId *big.Int, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateMetadata(&_ERC2771Forwarder.TransactOpts, smartObjectId, entityRecordMetadata)
}

// CreateMetadata is a paid mutator transaction binding the contract method 0x21823a2a.
//
// Solidity: function createMetadata(uint256 smartObjectId, (string,string,string) entityRecordMetadata) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateMetadata(smartObjectId *big.Int, entityRecordMetadata EntityMetadataParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateMetadata(&_ERC2771Forwarder.TransactOpts, smartObjectId, entityRecordMetadata)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xb4b891f7.
//
// Solidity: function createRecord(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CreateRecord(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "createRecord", smartObjectId, entityRecordParams)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xb4b891f7.
//
// Solidity: function createRecord(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CreateRecord(smartObjectId *big.Int, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateRecord(&_ERC2771Forwarder.TransactOpts, smartObjectId, entityRecordParams)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xb4b891f7.
//
// Solidity: function createRecord(uint256 smartObjectId, (bytes32,uint256,uint256,uint256) entityRecordParams) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CreateRecord(smartObjectId *big.Int, entityRecordParams EntityRecordParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CreateRecord(&_ERC2771Forwarder.TransactOpts, smartObjectId, entityRecordParams)
}

// CrossTransferToEphemeral is a paid mutator transaction binding the contract method 0xab2aea53.
//
// Solidity: function crossTransferToEphemeral(uint256 smartObjectId, address fromEphemeralOwner, address toEphemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) CrossTransferToEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, fromEphemeralOwner common.Address, toEphemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "crossTransferToEphemeral", smartObjectId, fromEphemeralOwner, toEphemeralOwner, items)
}

// CrossTransferToEphemeral is a paid mutator transaction binding the contract method 0xab2aea53.
//
// Solidity: function crossTransferToEphemeral(uint256 smartObjectId, address fromEphemeralOwner, address toEphemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) CrossTransferToEphemeral(smartObjectId *big.Int, fromEphemeralOwner common.Address, toEphemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CrossTransferToEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, fromEphemeralOwner, toEphemeralOwner, items)
}

// CrossTransferToEphemeral is a paid mutator transaction binding the contract method 0xab2aea53.
//
// Solidity: function crossTransferToEphemeral(uint256 smartObjectId, address fromEphemeralOwner, address toEphemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) CrossTransferToEphemeral(smartObjectId *big.Int, fromEphemeralOwner common.Address, toEphemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.CrossTransferToEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, fromEphemeralOwner, toEphemeralOwner, items)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) DeleteRecord(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "deleteRecord", tableId, keyTuple)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) DeleteRecord(tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DeleteRecord(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) DeleteRecord(tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DeleteRecord(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple)
}

// DepositEphemeral is a paid mutator transaction binding the contract method 0xf3688464.
//
// Solidity: function depositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) DepositEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "depositEphemeral", smartObjectId, ephemeralOwner, items)
}

// DepositEphemeral is a paid mutator transaction binding the contract method 0xf3688464.
//
// Solidity: function depositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) DepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DepositEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// DepositEphemeral is a paid mutator transaction binding the contract method 0xf3688464.
//
// Solidity: function depositEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) DepositEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DepositEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// DepositFuel is a paid mutator transaction binding the contract method 0xc3e9a45f.
//
// Solidity: function depositFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) DepositFuel(opts *bind.TransactOpts, smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "depositFuel", smartObjectId, fuelAmount)
}

// DepositFuel is a paid mutator transaction binding the contract method 0xc3e9a45f.
//
// Solidity: function depositFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) DepositFuel(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DepositFuel(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelAmount)
}

// DepositFuel is a paid mutator transaction binding the contract method 0xc3e9a45f.
//
// Solidity: function depositFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) DepositFuel(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DepositFuel(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelAmount)
}

// DepositInventory is a paid mutator transaction binding the contract method 0x1931d4e5.
//
// Solidity: function depositInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) DepositInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "depositInventory", smartObjectId, items)
}

// DepositInventory is a paid mutator transaction binding the contract method 0x1931d4e5.
//
// Solidity: function depositInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) DepositInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DepositInventory(&_ERC2771Forwarder.TransactOpts, smartObjectId, items)
}

// DepositInventory is a paid mutator transaction binding the contract method 0x1931d4e5.
//
// Solidity: function depositInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) DepositInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DepositInventory(&_ERC2771Forwarder.TransactOpts, smartObjectId, items)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) DestroyDeployable(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "destroyDeployable", smartObjectId)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) DestroyDeployable(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DestroyDeployable(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) DestroyDeployable(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.DestroyDeployable(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) GrantAccess(opts *bind.TransactOpts, resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "grantAccess", resourceId, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) GrantAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.GrantAccess(&_ERC2771Forwarder.TransactOpts, resourceId, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) GrantAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.GrantAccess(&_ERC2771Forwarder.TransactOpts, resourceId, grantee)
}

// HandleNodeOffline is a paid mutator transaction binding the contract method 0x9cf7293a.
//
// Solidity: function handleNodeOffline(uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) HandleNodeOffline(opts *bind.TransactOpts, networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "handleNodeOffline", networkNodeId)
}

// HandleNodeOffline is a paid mutator transaction binding the contract method 0x9cf7293a.
//
// Solidity: function handleNodeOffline(uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) HandleNodeOffline(networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.HandleNodeOffline(&_ERC2771Forwarder.TransactOpts, networkNodeId)
}

// HandleNodeOffline is a paid mutator transaction binding the contract method 0x9cf7293a.
//
// Solidity: function handleNodeOffline(uint256 networkNodeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) HandleNodeOffline(networkNodeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.HandleNodeOffline(&_ERC2771Forwarder.TransactOpts, networkNodeId)
}

// InProximity is a paid mutator transaction binding the contract method 0x5588377c.
//
// Solidity: function inProximity(uint256 smartObjectId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) InProximity(opts *bind.TransactOpts, smartObjectId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "inProximity", smartObjectId, priorityQueue, turret, turretTarget)
}

// InProximity is a paid mutator transaction binding the contract method 0x5588377c.
//
// Solidity: function inProximity(uint256 smartObjectId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771Forwarder *ERC2771ForwarderSession) InProximity(smartObjectId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.InProximity(&_ERC2771Forwarder.TransactOpts, smartObjectId, priorityQueue, turret, turretTarget)
}

// InProximity is a paid mutator transaction binding the contract method 0x5588377c.
//
// Solidity: function inProximity(uint256 smartObjectId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) InProximity(smartObjectId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.InProximity(&_ERC2771Forwarder.TransactOpts, smartObjectId, priorityQueue, turret, turretTarget)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) Initialize(opts *bind.TransactOpts, initModule common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "initialize", initModule)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) Initialize(initModule common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Initialize(&_ERC2771Forwarder.TransactOpts, initModule)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) Initialize(initModule common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Initialize(&_ERC2771Forwarder.TransactOpts, initModule)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) InstallModule(opts *bind.TransactOpts, module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "installModule", module, encodedArgs)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) InstallModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.InstallModule(&_ERC2771Forwarder.TransactOpts, module, encodedArgs)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) InstallModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.InstallModule(&_ERC2771Forwarder.TransactOpts, module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) InstallRootModule(opts *bind.TransactOpts, module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "installRootModule", module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) InstallRootModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.InstallRootModule(&_ERC2771Forwarder.TransactOpts, module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) InstallRootModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.InstallRootModule(&_ERC2771Forwarder.TransactOpts, module, encodedArgs)
}

// LinkGates is a paid mutator transaction binding the contract method 0x1caa06ce.
//
// Solidity: function linkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) LinkGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "linkGates", sourceGateId, destinationGateId)
}

// LinkGates is a paid mutator transaction binding the contract method 0x1caa06ce.
//
// Solidity: function linkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) LinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.LinkGates(&_ERC2771Forwarder.TransactOpts, sourceGateId, destinationGateId)
}

// LinkGates is a paid mutator transaction binding the contract method 0x1caa06ce.
//
// Solidity: function linkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) LinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.LinkGates(&_ERC2771Forwarder.TransactOpts, sourceGateId, destinationGateId)
}

// OnStructureOffline is a paid mutator transaction binding the contract method 0x4196f625.
//
// Solidity: function onStructureOffline(uint256 networkNodeId, uint256 structureId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) OnStructureOffline(opts *bind.TransactOpts, networkNodeId *big.Int, structureId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "onStructureOffline", networkNodeId, structureId)
}

// OnStructureOffline is a paid mutator transaction binding the contract method 0x4196f625.
//
// Solidity: function onStructureOffline(uint256 networkNodeId, uint256 structureId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnStructureOffline(networkNodeId *big.Int, structureId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.OnStructureOffline(&_ERC2771Forwarder.TransactOpts, networkNodeId, structureId)
}

// OnStructureOffline is a paid mutator transaction binding the contract method 0x4196f625.
//
// Solidity: function onStructureOffline(uint256 networkNodeId, uint256 structureId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) OnStructureOffline(networkNodeId *big.Int, structureId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.OnStructureOffline(&_ERC2771Forwarder.TransactOpts, networkNodeId, structureId)
}

// OnStructureOnline is a paid mutator transaction binding the contract method 0xfb03f98d.
//
// Solidity: function onStructureOnline(uint256 networkNodeId, uint256 structureId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) OnStructureOnline(opts *bind.TransactOpts, networkNodeId *big.Int, structureId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "onStructureOnline", networkNodeId, structureId)
}

// OnStructureOnline is a paid mutator transaction binding the contract method 0xfb03f98d.
//
// Solidity: function onStructureOnline(uint256 networkNodeId, uint256 structureId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) OnStructureOnline(networkNodeId *big.Int, structureId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.OnStructureOnline(&_ERC2771Forwarder.TransactOpts, networkNodeId, structureId)
}

// OnStructureOnline is a paid mutator transaction binding the contract method 0xfb03f98d.
//
// Solidity: function onStructureOnline(uint256 networkNodeId, uint256 structureId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) OnStructureOnline(networkNodeId *big.Int, structureId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.OnStructureOnline(&_ERC2771Forwarder.TransactOpts, networkNodeId, structureId)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) PopFromDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "popFromDynamicField", tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) PopFromDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.PopFromDynamicField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) PopFromDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.PopFromDynamicField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) PushToDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "pushToDynamicField", tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) PushToDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.PushToDynamicField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) PushToDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.PushToDynamicField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterDelegation(opts *bind.TransactOpts, delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerDelegation", delegatee, delegationControlId, initCallData)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterDelegation(delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterDelegation(&_ERC2771Forwarder.TransactOpts, delegatee, delegationControlId, initCallData)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterDelegation(delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterDelegation(&_ERC2771Forwarder.TransactOpts, delegatee, delegationControlId, initCallData)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerFunctionSelector", systemId, systemFunctionSignature)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterFunctionSelector(systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterFunctionSelector(&_ERC2771Forwarder.TransactOpts, systemId, systemFunctionSignature)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterFunctionSelector(systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterFunctionSelector(&_ERC2771Forwarder.TransactOpts, systemId, systemFunctionSignature)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterNamespace(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerNamespace", namespaceId)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterNamespace(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterNamespace(&_ERC2771Forwarder.TransactOpts, namespaceId)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterNamespace(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterNamespace(&_ERC2771Forwarder.TransactOpts, namespaceId)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterNamespaceDelegation(opts *bind.TransactOpts, namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerNamespaceDelegation", namespaceId, delegationControlId, initCallData)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterNamespaceDelegation(namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterNamespaceDelegation(&_ERC2771Forwarder.TransactOpts, namespaceId, delegationControlId, initCallData)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterNamespaceDelegation(namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterNamespaceDelegation(&_ERC2771Forwarder.TransactOpts, namespaceId, delegationControlId, initCallData)
}

// RegisterNetworkNodeClass is a paid mutator transaction binding the contract method 0x60ad4ad6.
//
// Solidity: function registerNetworkNodeClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterNetworkNodeClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerNetworkNodeClass", typeId, volume)
}

// RegisterNetworkNodeClass is a paid mutator transaction binding the contract method 0x60ad4ad6.
//
// Solidity: function registerNetworkNodeClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterNetworkNodeClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterNetworkNodeClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterNetworkNodeClass is a paid mutator transaction binding the contract method 0x60ad4ad6.
//
// Solidity: function registerNetworkNodeClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterNetworkNodeClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterNetworkNodeClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x6548a90a.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterRootFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, worldFunctionSignature string, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerRootFunctionSelector", systemId, worldFunctionSignature, systemFunctionSignature)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x6548a90a.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterRootFunctionSelector(&_ERC2771Forwarder.TransactOpts, systemId, worldFunctionSignature, systemFunctionSignature)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x6548a90a.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSignature string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterRootFunctionSelector(&_ERC2771Forwarder.TransactOpts, systemId, worldFunctionSignature, systemFunctionSignature)
}

// RegisterSmartCharacterClass is a paid mutator transaction binding the contract method 0x3664e851.
//
// Solidity: function registerSmartCharacterClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterSmartCharacterClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerSmartCharacterClass", typeId, volume)
}

// RegisterSmartCharacterClass is a paid mutator transaction binding the contract method 0x3664e851.
//
// Solidity: function registerSmartCharacterClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterSmartCharacterClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSmartCharacterClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterSmartCharacterClass is a paid mutator transaction binding the contract method 0x3664e851.
//
// Solidity: function registerSmartCharacterClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterSmartCharacterClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSmartCharacterClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterSmartGateClass is a paid mutator transaction binding the contract method 0x63d71589.
//
// Solidity: function registerSmartGateClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterSmartGateClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerSmartGateClass", typeId, volume)
}

// RegisterSmartGateClass is a paid mutator transaction binding the contract method 0x63d71589.
//
// Solidity: function registerSmartGateClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterSmartGateClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSmartGateClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterSmartGateClass is a paid mutator transaction binding the contract method 0x63d71589.
//
// Solidity: function registerSmartGateClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterSmartGateClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSmartGateClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterSmartStorageUnitClass is a paid mutator transaction binding the contract method 0x1f5493d1.
//
// Solidity: function registerSmartStorageUnitClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterSmartStorageUnitClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerSmartStorageUnitClass", typeId, volume)
}

// RegisterSmartStorageUnitClass is a paid mutator transaction binding the contract method 0x1f5493d1.
//
// Solidity: function registerSmartStorageUnitClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterSmartStorageUnitClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSmartStorageUnitClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterSmartStorageUnitClass is a paid mutator transaction binding the contract method 0x1f5493d1.
//
// Solidity: function registerSmartStorageUnitClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterSmartStorageUnitClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSmartStorageUnitClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterSmartTurretClass is a paid mutator transaction binding the contract method 0x2cf57627.
//
// Solidity: function registerSmartTurretClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterSmartTurretClass(opts *bind.TransactOpts, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerSmartTurretClass", typeId, volume)
}

// RegisterSmartTurretClass is a paid mutator transaction binding the contract method 0x2cf57627.
//
// Solidity: function registerSmartTurretClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterSmartTurretClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSmartTurretClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterSmartTurretClass is a paid mutator transaction binding the contract method 0x2cf57627.
//
// Solidity: function registerSmartTurretClass(uint256 typeId, uint256 volume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterSmartTurretClass(typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSmartTurretClass(&_ERC2771Forwarder.TransactOpts, typeId, volume)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterStoreHook(opts *bind.TransactOpts, tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerStoreHook", tableId, hookAddress, enabledHooksBitmap)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterStoreHook(tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterStoreHook(&_ERC2771Forwarder.TransactOpts, tableId, hookAddress, enabledHooksBitmap)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterStoreHook(tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterStoreHook(&_ERC2771Forwarder.TransactOpts, tableId, hookAddress, enabledHooksBitmap)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterSystem(opts *bind.TransactOpts, systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerSystem", systemId, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterSystem(systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSystem(&_ERC2771Forwarder.TransactOpts, systemId, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterSystem(systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSystem(&_ERC2771Forwarder.TransactOpts, systemId, system, publicAccess)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterSystemHook(opts *bind.TransactOpts, systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerSystemHook", systemId, hookAddress, enabledHooksBitmap)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterSystemHook(systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSystemHook(&_ERC2771Forwarder.TransactOpts, systemId, hookAddress, enabledHooksBitmap)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterSystemHook(systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterSystemHook(&_ERC2771Forwarder.TransactOpts, systemId, hookAddress, enabledHooksBitmap)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RegisterTable(opts *bind.TransactOpts, tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "registerTable", tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RegisterTable(tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterTable(&_ERC2771Forwarder.TransactOpts, tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RegisterTable(tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RegisterTable(&_ERC2771Forwarder.TransactOpts, tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RemoveCharacter is a paid mutator transaction binding the contract method 0x834f6531.
//
// Solidity: function removeCharacter(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RemoveCharacter(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "removeCharacter", smartObjectId)
}

// RemoveCharacter is a paid mutator transaction binding the contract method 0x834f6531.
//
// Solidity: function removeCharacter(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RemoveCharacter(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RemoveCharacter(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// RemoveCharacter is a paid mutator transaction binding the contract method 0x834f6531.
//
// Solidity: function removeCharacter(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RemoveCharacter(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RemoveCharacter(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// RemoveItemFromInventory is a paid mutator transaction binding the contract method 0x8a38b387.
//
// Solidity: function removeItemFromInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RemoveItemFromInventory(opts *bind.TransactOpts, inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "removeItemFromInventory", inventoryObjectId, itemObjectId, quantity)
}

// RemoveItemFromInventory is a paid mutator transaction binding the contract method 0x8a38b387.
//
// Solidity: function removeItemFromInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RemoveItemFromInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RemoveItemFromInventory(&_ERC2771Forwarder.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// RemoveItemFromInventory is a paid mutator transaction binding the contract method 0x8a38b387.
//
// Solidity: function removeItemFromInventory(uint256 inventoryObjectId, uint256 itemObjectId, uint256 quantity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RemoveItemFromInventory(inventoryObjectId *big.Int, itemObjectId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RemoveItemFromInventory(&_ERC2771Forwarder.TransactOpts, inventoryObjectId, itemObjectId, quantity)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x0058bbfd.
//
// Solidity: function removeOwner(uint256 smartObjectId, address from) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RemoveOwner(opts *bind.TransactOpts, smartObjectId *big.Int, from common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "removeOwner", smartObjectId, from)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x0058bbfd.
//
// Solidity: function removeOwner(uint256 smartObjectId, address from) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RemoveOwner(smartObjectId *big.Int, from common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RemoveOwner(&_ERC2771Forwarder.TransactOpts, smartObjectId, from)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x0058bbfd.
//
// Solidity: function removeOwner(uint256 smartObjectId, address from) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RemoveOwner(smartObjectId *big.Int, from common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RemoveOwner(&_ERC2771Forwarder.TransactOpts, smartObjectId, from)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RenounceOwnership(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "renounceOwnership", namespaceId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RenounceOwnership(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RenounceOwnership(&_ERC2771Forwarder.TransactOpts, namespaceId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RenounceOwnership(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RenounceOwnership(&_ERC2771Forwarder.TransactOpts, namespaceId)
}

// ReportKill is a paid mutator transaction binding the contract method 0xda3293e1.
//
// Solidity: function reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ReportKill(opts *bind.TransactOpts, killMailId *big.Int, killMailData KillMailData) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "reportKill", killMailId, killMailData)
}

// ReportKill is a paid mutator transaction binding the contract method 0xda3293e1.
//
// Solidity: function reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ReportKill(killMailId *big.Int, killMailData KillMailData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ReportKill(&_ERC2771Forwarder.TransactOpts, killMailId, killMailData)
}

// ReportKill is a paid mutator transaction binding the contract method 0xda3293e1.
//
// Solidity: function reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ReportKill(killMailId *big.Int, killMailData KillMailData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ReportKill(&_ERC2771Forwarder.TransactOpts, killMailId, killMailData)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) RevokeAccess(opts *bind.TransactOpts, resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "revokeAccess", resourceId, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) RevokeAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RevokeAccess(&_ERC2771Forwarder.TransactOpts, resourceId, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) RevokeAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.RevokeAccess(&_ERC2771Forwarder.TransactOpts, resourceId, grantee)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 smartObjectId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SaveLocation(opts *bind.TransactOpts, smartObjectId *big.Int, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "saveLocation", smartObjectId, locationData)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 smartObjectId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SaveLocation(smartObjectId *big.Int, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SaveLocation(&_ERC2771Forwarder.TransactOpts, smartObjectId, locationData)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 smartObjectId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SaveLocation(smartObjectId *big.Int, locationData LocationData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SaveLocation(&_ERC2771Forwarder.TransactOpts, smartObjectId, locationData)
}

// SetAssemblyType is a paid mutator transaction binding the contract method 0xf9914db3.
//
// Solidity: function setAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetAssemblyType(opts *bind.TransactOpts, smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setAssemblyType", smartObjectId, assemblyType)
}

// SetAssemblyType is a paid mutator transaction binding the contract method 0xf9914db3.
//
// Solidity: function setAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetAssemblyType(&_ERC2771Forwarder.TransactOpts, smartObjectId, assemblyType)
}

// SetAssemblyType is a paid mutator transaction binding the contract method 0xf9914db3.
//
// Solidity: function setAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetAssemblyType(&_ERC2771Forwarder.TransactOpts, smartObjectId, assemblyType)
}

// SetCapacity is a paid mutator transaction binding the contract method 0xdee2b058.
//
// Solidity: function setCapacity(uint256 smartObjectId, uint256 capacity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, capacity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setCapacity", smartObjectId, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0xdee2b058.
//
// Solidity: function setCapacity(uint256 smartObjectId, uint256 capacity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetCapacity(smartObjectId *big.Int, capacity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetCapacity(&_ERC2771Forwarder.TransactOpts, smartObjectId, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0xdee2b058.
//
// Solidity: function setCapacity(uint256 smartObjectId, uint256 capacity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetCapacity(smartObjectId *big.Int, capacity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetCapacity(&_ERC2771Forwarder.TransactOpts, smartObjectId, capacity)
}

// SetCrossTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0x5f3be290.
//
// Solidity: function setCrossTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetCrossTransferToEphemeralAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setCrossTransferToEphemeralAccess", smartObjectId, accessAddress, isAllowed)
}

// SetCrossTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0x5f3be290.
//
// Solidity: function setCrossTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetCrossTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetCrossTransferToEphemeralAccess(&_ERC2771Forwarder.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetCrossTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0x5f3be290.
//
// Solidity: function setCrossTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetCrossTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetCrossTransferToEphemeralAccess(&_ERC2771Forwarder.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 smartObjectId, string dappURL) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetDappURL(opts *bind.TransactOpts, smartObjectId *big.Int, dappURL string) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setDappURL", smartObjectId, dappURL)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 smartObjectId, string dappURL) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetDappURL(smartObjectId *big.Int, dappURL string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetDappURL(&_ERC2771Forwarder.TransactOpts, smartObjectId, dappURL)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 smartObjectId, string dappURL) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetDappURL(smartObjectId *big.Int, dappURL string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetDappURL(&_ERC2771Forwarder.TransactOpts, smartObjectId, dappURL)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 smartObjectId, string description) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetDescription(opts *bind.TransactOpts, smartObjectId *big.Int, description string) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setDescription", smartObjectId, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 smartObjectId, string description) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetDescription(smartObjectId *big.Int, description string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetDescription(&_ERC2771Forwarder.TransactOpts, smartObjectId, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 smartObjectId, string description) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetDescription(smartObjectId *big.Int, description string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetDescription(&_ERC2771Forwarder.TransactOpts, smartObjectId, description)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setDynamicField", tableId, keyTuple, dynamicFieldIndex, data)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetDynamicField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, dynamicFieldIndex, data)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetDynamicField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, dynamicFieldIndex, data)
}

// SetEphemeralCapacity is a paid mutator transaction binding the contract method 0xbc8229dc.
//
// Solidity: function setEphemeralCapacity(uint256 smartObjectId, uint256 ephemeralCapacity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetEphemeralCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setEphemeralCapacity", smartObjectId, ephemeralCapacity)
}

// SetEphemeralCapacity is a paid mutator transaction binding the contract method 0xbc8229dc.
//
// Solidity: function setEphemeralCapacity(uint256 smartObjectId, uint256 ephemeralCapacity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetEphemeralCapacity(smartObjectId *big.Int, ephemeralCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetEphemeralCapacity(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralCapacity)
}

// SetEphemeralCapacity is a paid mutator transaction binding the contract method 0xbc8229dc.
//
// Solidity: function setEphemeralCapacity(uint256 smartObjectId, uint256 ephemeralCapacity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetEphemeralCapacity(smartObjectId *big.Int, ephemeralCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetEphemeralCapacity(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralCapacity)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setField", tableId, keyTuple, fieldIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, fieldIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, fieldIndex, data)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetField0(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setField0", tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetField0(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetField0(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetFuelAmount is a paid mutator transaction binding the contract method 0xdd9879ba.
//
// Solidity: function setFuelAmount(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetFuelAmount(opts *bind.TransactOpts, smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setFuelAmount", smartObjectId, fuelAmount)
}

// SetFuelAmount is a paid mutator transaction binding the contract method 0xdd9879ba.
//
// Solidity: function setFuelAmount(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetFuelAmount(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetFuelAmount(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelAmount)
}

// SetFuelAmount is a paid mutator transaction binding the contract method 0xdd9879ba.
//
// Solidity: function setFuelAmount(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetFuelAmount(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetFuelAmount(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelAmount)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 smartObjectId, uint256 fuelMaxCapacity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetFuelMaxCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setFuelMaxCapacity", smartObjectId, fuelMaxCapacity)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 smartObjectId, uint256 fuelMaxCapacity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetFuelMaxCapacity(smartObjectId *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetFuelMaxCapacity(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelMaxCapacity)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 smartObjectId, uint256 fuelMaxCapacity) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetFuelMaxCapacity(smartObjectId *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetFuelMaxCapacity(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelMaxCapacity)
}

// SetFuelUnitVolume is a paid mutator transaction binding the contract method 0x20fa3423.
//
// Solidity: function setFuelUnitVolume(uint256 smartObjectId, uint256 fuelUnitVolume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetFuelUnitVolume(opts *bind.TransactOpts, smartObjectId *big.Int, fuelUnitVolume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setFuelUnitVolume", smartObjectId, fuelUnitVolume)
}

// SetFuelUnitVolume is a paid mutator transaction binding the contract method 0x20fa3423.
//
// Solidity: function setFuelUnitVolume(uint256 smartObjectId, uint256 fuelUnitVolume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetFuelUnitVolume(smartObjectId *big.Int, fuelUnitVolume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetFuelUnitVolume(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelUnitVolume)
}

// SetFuelUnitVolume is a paid mutator transaction binding the contract method 0x20fa3423.
//
// Solidity: function setFuelUnitVolume(uint256 smartObjectId, uint256 fuelUnitVolume) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetFuelUnitVolume(smartObjectId *big.Int, fuelUnitVolume *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetFuelUnitVolume(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelUnitVolume)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 smartObjectId, string name) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetName(opts *bind.TransactOpts, smartObjectId *big.Int, name string) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setName", smartObjectId, name)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 smartObjectId, string name) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetName(smartObjectId *big.Int, name string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetName(&_ERC2771Forwarder.TransactOpts, smartObjectId, name)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 smartObjectId, string name) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetName(smartObjectId *big.Int, name string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetName(&_ERC2771Forwarder.TransactOpts, smartObjectId, name)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetRecord(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setRecord", tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetRecord(tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetRecord(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetRecord(tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetRecord(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetStaticField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setStaticField", tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetStaticField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetStaticField(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetTransferFromEphemeralAccess is a paid mutator transaction binding the contract method 0x8d0ead31.
//
// Solidity: function setTransferFromEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetTransferFromEphemeralAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setTransferFromEphemeralAccess", smartObjectId, accessAddress, isAllowed)
}

// SetTransferFromEphemeralAccess is a paid mutator transaction binding the contract method 0x8d0ead31.
//
// Solidity: function setTransferFromEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetTransferFromEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetTransferFromEphemeralAccess(&_ERC2771Forwarder.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferFromEphemeralAccess is a paid mutator transaction binding the contract method 0x8d0ead31.
//
// Solidity: function setTransferFromEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetTransferFromEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetTransferFromEphemeralAccess(&_ERC2771Forwarder.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0xd6873420.
//
// Solidity: function setTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetTransferToEphemeralAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setTransferToEphemeralAccess", smartObjectId, accessAddress, isAllowed)
}

// SetTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0xd6873420.
//
// Solidity: function setTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetTransferToEphemeralAccess(&_ERC2771Forwarder.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferToEphemeralAccess is a paid mutator transaction binding the contract method 0xd6873420.
//
// Solidity: function setTransferToEphemeralAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetTransferToEphemeralAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetTransferToEphemeralAccess(&_ERC2771Forwarder.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferToInventoryAccess is a paid mutator transaction binding the contract method 0x0a231135.
//
// Solidity: function setTransferToInventoryAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SetTransferToInventoryAccess(opts *bind.TransactOpts, smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "setTransferToInventoryAccess", smartObjectId, accessAddress, isAllowed)
}

// SetTransferToInventoryAccess is a paid mutator transaction binding the contract method 0x0a231135.
//
// Solidity: function setTransferToInventoryAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SetTransferToInventoryAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetTransferToInventoryAccess(&_ERC2771Forwarder.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SetTransferToInventoryAccess is a paid mutator transaction binding the contract method 0x0a231135.
//
// Solidity: function setTransferToInventoryAccess(uint256 smartObjectId, address accessAddress, bool isAllowed) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SetTransferToInventoryAccess(smartObjectId *big.Int, accessAddress common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SetTransferToInventoryAccess(&_ERC2771Forwarder.TransactOpts, smartObjectId, accessAddress, isAllowed)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SpliceDynamicData(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "spliceDynamicData", tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SpliceDynamicData(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SpliceDynamicData(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SpliceDynamicData(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SpliceDynamicData(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) SpliceStaticData(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "spliceStaticData", tableId, keyTuple, start, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) SpliceStaticData(tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SpliceStaticData(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, start, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) SpliceStaticData(tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.SpliceStaticData(&_ERC2771Forwarder.TransactOpts, tableId, keyTuple, start, data)
}

// StartBurn is a paid mutator transaction binding the contract method 0x46d647ab.
//
// Solidity: function startBurn(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) StartBurn(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "startBurn", smartObjectId)
}

// StartBurn is a paid mutator transaction binding the contract method 0x46d647ab.
//
// Solidity: function startBurn(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) StartBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.StartBurn(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// StartBurn is a paid mutator transaction binding the contract method 0x46d647ab.
//
// Solidity: function startBurn(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) StartBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.StartBurn(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// StopBurn is a paid mutator transaction binding the contract method 0xe2487ae0.
//
// Solidity: function stopBurn(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) StopBurn(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "stopBurn", smartObjectId)
}

// StopBurn is a paid mutator transaction binding the contract method 0xe2487ae0.
//
// Solidity: function stopBurn(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) StopBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.StopBurn(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// StopBurn is a paid mutator transaction binding the contract method 0xe2487ae0.
//
// Solidity: function stopBurn(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) StopBurn(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.StopBurn(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) TransferBalanceToAddress(opts *bind.TransactOpts, fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "transferBalanceToAddress", fromNamespaceId, toAddress, amount)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) TransferBalanceToAddress(fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferBalanceToAddress(&_ERC2771Forwarder.TransactOpts, fromNamespaceId, toAddress, amount)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) TransferBalanceToAddress(fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferBalanceToAddress(&_ERC2771Forwarder.TransactOpts, fromNamespaceId, toAddress, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) TransferBalanceToNamespace(opts *bind.TransactOpts, fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "transferBalanceToNamespace", fromNamespaceId, toNamespaceId, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) TransferBalanceToNamespace(fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferBalanceToNamespace(&_ERC2771Forwarder.TransactOpts, fromNamespaceId, toNamespaceId, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) TransferBalanceToNamespace(fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferBalanceToNamespace(&_ERC2771Forwarder.TransactOpts, fromNamespaceId, toNamespaceId, amount)
}

// TransferFromEphemeral is a paid mutator transaction binding the contract method 0x6f29e3b1.
//
// Solidity: function transferFromEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) TransferFromEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "transferFromEphemeral", smartObjectId, ephemeralOwner, items)
}

// TransferFromEphemeral is a paid mutator transaction binding the contract method 0x6f29e3b1.
//
// Solidity: function transferFromEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) TransferFromEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferFromEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// TransferFromEphemeral is a paid mutator transaction binding the contract method 0x6f29e3b1.
//
// Solidity: function transferFromEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) TransferFromEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferFromEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) TransferOwnership(opts *bind.TransactOpts, namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "transferOwnership", namespaceId, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) TransferOwnership(namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferOwnership(&_ERC2771Forwarder.TransactOpts, namespaceId, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) TransferOwnership(namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferOwnership(&_ERC2771Forwarder.TransactOpts, namespaceId, newOwner)
}

// TransferToEphemeral is a paid mutator transaction binding the contract method 0x63821267.
//
// Solidity: function transferToEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) TransferToEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "transferToEphemeral", smartObjectId, ephemeralOwner, items)
}

// TransferToEphemeral is a paid mutator transaction binding the contract method 0x63821267.
//
// Solidity: function transferToEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) TransferToEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferToEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// TransferToEphemeral is a paid mutator transaction binding the contract method 0x63821267.
//
// Solidity: function transferToEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) TransferToEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferToEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// TransferToInventory is a paid mutator transaction binding the contract method 0x4ce8f5ce.
//
// Solidity: function transferToInventory(uint256 smartObjectId, uint256 toObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) TransferToInventory(opts *bind.TransactOpts, smartObjectId *big.Int, toObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "transferToInventory", smartObjectId, toObjectId, items)
}

// TransferToInventory is a paid mutator transaction binding the contract method 0x4ce8f5ce.
//
// Solidity: function transferToInventory(uint256 smartObjectId, uint256 toObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) TransferToInventory(smartObjectId *big.Int, toObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferToInventory(&_ERC2771Forwarder.TransactOpts, smartObjectId, toObjectId, items)
}

// TransferToInventory is a paid mutator transaction binding the contract method 0x4ce8f5ce.
//
// Solidity: function transferToInventory(uint256 smartObjectId, uint256 toObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) TransferToInventory(smartObjectId *big.Int, toObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.TransferToInventory(&_ERC2771Forwarder.TransactOpts, smartObjectId, toObjectId, items)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) Unanchor(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "unanchor", smartObjectId)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) Unanchor(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Unanchor(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) Unanchor(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Unanchor(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// UnlinkGates is a paid mutator transaction binding the contract method 0x6a57ebde.
//
// Solidity: function unlinkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) UnlinkGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "unlinkGates", sourceGateId, destinationGateId)
}

// UnlinkGates is a paid mutator transaction binding the contract method 0x6a57ebde.
//
// Solidity: function unlinkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) UnlinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnlinkGates(&_ERC2771Forwarder.TransactOpts, sourceGateId, destinationGateId)
}

// UnlinkGates is a paid mutator transaction binding the contract method 0x6a57ebde.
//
// Solidity: function unlinkGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) UnlinkGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnlinkGates(&_ERC2771Forwarder.TransactOpts, sourceGateId, destinationGateId)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) UnregisterDelegation(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "unregisterDelegation", delegatee)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) UnregisterDelegation(delegatee common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnregisterDelegation(&_ERC2771Forwarder.TransactOpts, delegatee)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) UnregisterDelegation(delegatee common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnregisterDelegation(&_ERC2771Forwarder.TransactOpts, delegatee)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) UnregisterNamespaceDelegation(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "unregisterNamespaceDelegation", namespaceId)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) UnregisterNamespaceDelegation(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnregisterNamespaceDelegation(&_ERC2771Forwarder.TransactOpts, namespaceId)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) UnregisterNamespaceDelegation(namespaceId [32]byte) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnregisterNamespaceDelegation(&_ERC2771Forwarder.TransactOpts, namespaceId)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) UnregisterStoreHook(opts *bind.TransactOpts, tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "unregisterStoreHook", tableId, hookAddress)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) UnregisterStoreHook(tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnregisterStoreHook(&_ERC2771Forwarder.TransactOpts, tableId, hookAddress)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) UnregisterStoreHook(tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnregisterStoreHook(&_ERC2771Forwarder.TransactOpts, tableId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) UnregisterSystemHook(opts *bind.TransactOpts, systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "unregisterSystemHook", systemId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) UnregisterSystemHook(systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnregisterSystemHook(&_ERC2771Forwarder.TransactOpts, systemId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) UnregisterSystemHook(systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UnregisterSystemHook(&_ERC2771Forwarder.TransactOpts, systemId, hookAddress)
}

// UpdateAssemblyType is a paid mutator transaction binding the contract method 0xbcb029b2.
//
// Solidity: function updateAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) UpdateAssemblyType(opts *bind.TransactOpts, smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "updateAssemblyType", smartObjectId, assemblyType)
}

// UpdateAssemblyType is a paid mutator transaction binding the contract method 0xbcb029b2.
//
// Solidity: function updateAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) UpdateAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UpdateAssemblyType(&_ERC2771Forwarder.TransactOpts, smartObjectId, assemblyType)
}

// UpdateAssemblyType is a paid mutator transaction binding the contract method 0xbcb029b2.
//
// Solidity: function updateAssemblyType(uint256 smartObjectId, string assemblyType) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) UpdateAssemblyType(smartObjectId *big.Int, assemblyType string) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UpdateAssemblyType(&_ERC2771Forwarder.TransactOpts, smartObjectId, assemblyType)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) UpdateFuel(opts *bind.TransactOpts, smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "updateFuel", smartObjectId)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) UpdateFuel(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UpdateFuel(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 smartObjectId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) UpdateFuel(smartObjectId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UpdateFuel(&_ERC2771Forwarder.TransactOpts, smartObjectId)
}

// UpdateTribeId is a paid mutator transaction binding the contract method 0x16d51f3e.
//
// Solidity: function updateTribeId(uint256 smartObjectId, uint256 tribeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) UpdateTribeId(opts *bind.TransactOpts, smartObjectId *big.Int, tribeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "updateTribeId", smartObjectId, tribeId)
}

// UpdateTribeId is a paid mutator transaction binding the contract method 0x16d51f3e.
//
// Solidity: function updateTribeId(uint256 smartObjectId, uint256 tribeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) UpdateTribeId(smartObjectId *big.Int, tribeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UpdateTribeId(&_ERC2771Forwarder.TransactOpts, smartObjectId, tribeId)
}

// UpdateTribeId is a paid mutator transaction binding the contract method 0x16d51f3e.
//
// Solidity: function updateTribeId(uint256 smartObjectId, uint256 tribeId) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) UpdateTribeId(smartObjectId *big.Int, tribeId *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.UpdateTribeId(&_ERC2771Forwarder.TransactOpts, smartObjectId, tribeId)
}

// WithdrawEphemeral is a paid mutator transaction binding the contract method 0x690bcc03.
//
// Solidity: function withdrawEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) WithdrawEphemeral(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "withdrawEphemeral", smartObjectId, ephemeralOwner, items)
}

// WithdrawEphemeral is a paid mutator transaction binding the contract method 0x690bcc03.
//
// Solidity: function withdrawEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) WithdrawEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.WithdrawEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// WithdrawEphemeral is a paid mutator transaction binding the contract method 0x690bcc03.
//
// Solidity: function withdrawEphemeral(uint256 smartObjectId, address ephemeralOwner, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) WithdrawEphemeral(smartObjectId *big.Int, ephemeralOwner common.Address, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.WithdrawEphemeral(&_ERC2771Forwarder.TransactOpts, smartObjectId, ephemeralOwner, items)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) WithdrawFuel(opts *bind.TransactOpts, smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "withdrawFuel", smartObjectId, fuelAmount)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) WithdrawFuel(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.WithdrawFuel(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelAmount)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 smartObjectId, uint256 fuelAmount) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) WithdrawFuel(smartObjectId *big.Int, fuelAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.WithdrawFuel(&_ERC2771Forwarder.TransactOpts, smartObjectId, fuelAmount)
}

// WithdrawInventory is a paid mutator transaction binding the contract method 0xec92b5a5.
//
// Solidity: function withdrawInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) WithdrawInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "withdrawInventory", smartObjectId, items)
}

// WithdrawInventory is a paid mutator transaction binding the contract method 0xec92b5a5.
//
// Solidity: function withdrawInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) WithdrawInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.WithdrawInventory(&_ERC2771Forwarder.TransactOpts, smartObjectId, items)
}

// WithdrawInventory is a paid mutator transaction binding the contract method 0xec92b5a5.
//
// Solidity: function withdrawInventory(uint256 smartObjectId, (uint256,uint256)[] items) returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) WithdrawInventory(smartObjectId *big.Int, items []InventoryItemParams) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.WithdrawInventory(&_ERC2771Forwarder.TransactOpts, smartObjectId, items)
}

// ERC2771ForwarderHelloStoreIterator is returned from FilterHelloStore and is used to iterate over the raw logs and unpacked data for HelloStore events raised by the ERC2771Forwarder contract.
type ERC2771ForwarderHelloStoreIterator struct {
	Event *ERC2771ForwarderHelloStore // Event containing the contract specifics and raw log

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
func (it *ERC2771ForwarderHelloStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771ForwarderHelloStore)
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
		it.Event = new(ERC2771ForwarderHelloStore)
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
func (it *ERC2771ForwarderHelloStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771ForwarderHelloStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771ForwarderHelloStore represents a HelloStore event raised by the ERC2771Forwarder contract.
type ERC2771ForwarderHelloStore struct {
	StoreVersion [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHelloStore is a free log retrieval operation binding the contract event 0xc7f5fdc8526b76f54916701bc910876243ffff2a40b0bb8d59eea8151c52c005.
//
// Solidity: event HelloStore(bytes32 indexed storeVersion)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) FilterHelloStore(opts *bind.FilterOpts, storeVersion [][32]byte) (*ERC2771ForwarderHelloStoreIterator, error) {

	var storeVersionRule []interface{}
	for _, storeVersionItem := range storeVersion {
		storeVersionRule = append(storeVersionRule, storeVersionItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.FilterLogs(opts, "HelloStore", storeVersionRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderHelloStoreIterator{contract: _ERC2771Forwarder.contract, event: "HelloStore", logs: logs, sub: sub}, nil
}

// WatchHelloStore is a free log subscription operation binding the contract event 0xc7f5fdc8526b76f54916701bc910876243ffff2a40b0bb8d59eea8151c52c005.
//
// Solidity: event HelloStore(bytes32 indexed storeVersion)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) WatchHelloStore(opts *bind.WatchOpts, sink chan<- *ERC2771ForwarderHelloStore, storeVersion [][32]byte) (event.Subscription, error) {

	var storeVersionRule []interface{}
	for _, storeVersionItem := range storeVersion {
		storeVersionRule = append(storeVersionRule, storeVersionItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.WatchLogs(opts, "HelloStore", storeVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771ForwarderHelloStore)
				if err := _ERC2771Forwarder.contract.UnpackLog(event, "HelloStore", log); err != nil {
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
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) ParseHelloStore(log types.Log) (*ERC2771ForwarderHelloStore, error) {
	event := new(ERC2771ForwarderHelloStore)
	if err := _ERC2771Forwarder.contract.UnpackLog(event, "HelloStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771ForwarderHelloWorldIterator is returned from FilterHelloWorld and is used to iterate over the raw logs and unpacked data for HelloWorld events raised by the ERC2771Forwarder contract.
type ERC2771ForwarderHelloWorldIterator struct {
	Event *ERC2771ForwarderHelloWorld // Event containing the contract specifics and raw log

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
func (it *ERC2771ForwarderHelloWorldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771ForwarderHelloWorld)
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
		it.Event = new(ERC2771ForwarderHelloWorld)
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
func (it *ERC2771ForwarderHelloWorldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771ForwarderHelloWorldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771ForwarderHelloWorld represents a HelloWorld event raised by the ERC2771Forwarder contract.
type ERC2771ForwarderHelloWorld struct {
	WorldVersion [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHelloWorld is a free log retrieval operation binding the contract event 0x7f8f36afe3fb61c459c1a54a60b8a477eab02cc58e49f547561a40906239cb82.
//
// Solidity: event HelloWorld(bytes32 indexed worldVersion)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) FilterHelloWorld(opts *bind.FilterOpts, worldVersion [][32]byte) (*ERC2771ForwarderHelloWorldIterator, error) {

	var worldVersionRule []interface{}
	for _, worldVersionItem := range worldVersion {
		worldVersionRule = append(worldVersionRule, worldVersionItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.FilterLogs(opts, "HelloWorld", worldVersionRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderHelloWorldIterator{contract: _ERC2771Forwarder.contract, event: "HelloWorld", logs: logs, sub: sub}, nil
}

// WatchHelloWorld is a free log subscription operation binding the contract event 0x7f8f36afe3fb61c459c1a54a60b8a477eab02cc58e49f547561a40906239cb82.
//
// Solidity: event HelloWorld(bytes32 indexed worldVersion)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) WatchHelloWorld(opts *bind.WatchOpts, sink chan<- *ERC2771ForwarderHelloWorld, worldVersion [][32]byte) (event.Subscription, error) {

	var worldVersionRule []interface{}
	for _, worldVersionItem := range worldVersion {
		worldVersionRule = append(worldVersionRule, worldVersionItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.WatchLogs(opts, "HelloWorld", worldVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771ForwarderHelloWorld)
				if err := _ERC2771Forwarder.contract.UnpackLog(event, "HelloWorld", log); err != nil {
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
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) ParseHelloWorld(log types.Log) (*ERC2771ForwarderHelloWorld, error) {
	event := new(ERC2771ForwarderHelloWorld)
	if err := _ERC2771Forwarder.contract.UnpackLog(event, "HelloWorld", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771ForwarderStoreDeleteRecordIterator is returned from FilterStoreDeleteRecord and is used to iterate over the raw logs and unpacked data for StoreDeleteRecord events raised by the ERC2771Forwarder contract.
type ERC2771ForwarderStoreDeleteRecordIterator struct {
	Event *ERC2771ForwarderStoreDeleteRecord // Event containing the contract specifics and raw log

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
func (it *ERC2771ForwarderStoreDeleteRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771ForwarderStoreDeleteRecord)
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
		it.Event = new(ERC2771ForwarderStoreDeleteRecord)
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
func (it *ERC2771ForwarderStoreDeleteRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771ForwarderStoreDeleteRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771ForwarderStoreDeleteRecord represents a StoreDeleteRecord event raised by the ERC2771Forwarder contract.
type ERC2771ForwarderStoreDeleteRecord struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreDeleteRecord is a free log retrieval operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) FilterStoreDeleteRecord(opts *bind.FilterOpts, tableId [][32]byte) (*ERC2771ForwarderStoreDeleteRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.FilterLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderStoreDeleteRecordIterator{contract: _ERC2771Forwarder.contract, event: "Store_DeleteRecord", logs: logs, sub: sub}, nil
}

// WatchStoreDeleteRecord is a free log subscription operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) WatchStoreDeleteRecord(opts *bind.WatchOpts, sink chan<- *ERC2771ForwarderStoreDeleteRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.WatchLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771ForwarderStoreDeleteRecord)
				if err := _ERC2771Forwarder.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
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
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) ParseStoreDeleteRecord(log types.Log) (*ERC2771ForwarderStoreDeleteRecord, error) {
	event := new(ERC2771ForwarderStoreDeleteRecord)
	if err := _ERC2771Forwarder.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771ForwarderStoreSetRecordIterator is returned from FilterStoreSetRecord and is used to iterate over the raw logs and unpacked data for StoreSetRecord events raised by the ERC2771Forwarder contract.
type ERC2771ForwarderStoreSetRecordIterator struct {
	Event *ERC2771ForwarderStoreSetRecord // Event containing the contract specifics and raw log

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
func (it *ERC2771ForwarderStoreSetRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771ForwarderStoreSetRecord)
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
		it.Event = new(ERC2771ForwarderStoreSetRecord)
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
func (it *ERC2771ForwarderStoreSetRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771ForwarderStoreSetRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771ForwarderStoreSetRecord represents a StoreSetRecord event raised by the ERC2771Forwarder contract.
type ERC2771ForwarderStoreSetRecord struct {
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
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) FilterStoreSetRecord(opts *bind.FilterOpts, tableId [][32]byte) (*ERC2771ForwarderStoreSetRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.FilterLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderStoreSetRecordIterator{contract: _ERC2771Forwarder.contract, event: "Store_SetRecord", logs: logs, sub: sub}, nil
}

// WatchStoreSetRecord is a free log subscription operation binding the contract event 0x8dbb3a9672eebfd3773e72dd9c102393436816d832c7ba9e1e1ac8fcadcac7a9.
//
// Solidity: event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) WatchStoreSetRecord(opts *bind.WatchOpts, sink chan<- *ERC2771ForwarderStoreSetRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.WatchLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771ForwarderStoreSetRecord)
				if err := _ERC2771Forwarder.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
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
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) ParseStoreSetRecord(log types.Log) (*ERC2771ForwarderStoreSetRecord, error) {
	event := new(ERC2771ForwarderStoreSetRecord)
	if err := _ERC2771Forwarder.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771ForwarderStoreSpliceDynamicDataIterator is returned from FilterStoreSpliceDynamicData and is used to iterate over the raw logs and unpacked data for StoreSpliceDynamicData events raised by the ERC2771Forwarder contract.
type ERC2771ForwarderStoreSpliceDynamicDataIterator struct {
	Event *ERC2771ForwarderStoreSpliceDynamicData // Event containing the contract specifics and raw log

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
func (it *ERC2771ForwarderStoreSpliceDynamicDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771ForwarderStoreSpliceDynamicData)
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
		it.Event = new(ERC2771ForwarderStoreSpliceDynamicData)
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
func (it *ERC2771ForwarderStoreSpliceDynamicDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771ForwarderStoreSpliceDynamicDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771ForwarderStoreSpliceDynamicData represents a StoreSpliceDynamicData event raised by the ERC2771Forwarder contract.
type ERC2771ForwarderStoreSpliceDynamicData struct {
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
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) FilterStoreSpliceDynamicData(opts *bind.FilterOpts, tableId [][32]byte) (*ERC2771ForwarderStoreSpliceDynamicDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.FilterLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderStoreSpliceDynamicDataIterator{contract: _ERC2771Forwarder.contract, event: "Store_SpliceDynamicData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceDynamicData is a free log subscription operation binding the contract event 0xfe158a7adba34e256807c8a149028d3162918713c3838afc643ce9f96716ebfd.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) WatchStoreSpliceDynamicData(opts *bind.WatchOpts, sink chan<- *ERC2771ForwarderStoreSpliceDynamicData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.WatchLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771ForwarderStoreSpliceDynamicData)
				if err := _ERC2771Forwarder.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
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
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) ParseStoreSpliceDynamicData(log types.Log) (*ERC2771ForwarderStoreSpliceDynamicData, error) {
	event := new(ERC2771ForwarderStoreSpliceDynamicData)
	if err := _ERC2771Forwarder.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771ForwarderStoreSpliceStaticDataIterator is returned from FilterStoreSpliceStaticData and is used to iterate over the raw logs and unpacked data for StoreSpliceStaticData events raised by the ERC2771Forwarder contract.
type ERC2771ForwarderStoreSpliceStaticDataIterator struct {
	Event *ERC2771ForwarderStoreSpliceStaticData // Event containing the contract specifics and raw log

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
func (it *ERC2771ForwarderStoreSpliceStaticDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771ForwarderStoreSpliceStaticData)
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
		it.Event = new(ERC2771ForwarderStoreSpliceStaticData)
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
func (it *ERC2771ForwarderStoreSpliceStaticDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771ForwarderStoreSpliceStaticDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771ForwarderStoreSpliceStaticData represents a StoreSpliceStaticData event raised by the ERC2771Forwarder contract.
type ERC2771ForwarderStoreSpliceStaticData struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Start    *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreSpliceStaticData is a free log retrieval operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) FilterStoreSpliceStaticData(opts *bind.FilterOpts, tableId [][32]byte) (*ERC2771ForwarderStoreSpliceStaticDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.FilterLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderStoreSpliceStaticDataIterator{contract: _ERC2771Forwarder.contract, event: "Store_SpliceStaticData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceStaticData is a free log subscription operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) WatchStoreSpliceStaticData(opts *bind.WatchOpts, sink chan<- *ERC2771ForwarderStoreSpliceStaticData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.WatchLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771ForwarderStoreSpliceStaticData)
				if err := _ERC2771Forwarder.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
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
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) ParseStoreSpliceStaticData(log types.Log) (*ERC2771ForwarderStoreSpliceStaticData, error) {
	event := new(ERC2771ForwarderStoreSpliceStaticData)
	if err := _ERC2771Forwarder.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
