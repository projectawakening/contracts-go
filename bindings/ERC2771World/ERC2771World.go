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

// Coord is an auto generated low-level Go binding around an user-defined struct.
type Coord struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
}

// EntityRecordData is an auto generated low-level Go binding around an user-defined struct.
type EntityRecordData struct {
	TypeId *big.Int
	ItemId *big.Int
	Volume *big.Int
}

// EntityRecordOffchainTableData is an auto generated low-level Go binding around an user-defined struct.
type EntityRecordOffchainTableData struct {
	Name        string
	DappURL     string
	Description string
}

// InventoryItem is an auto generated low-level Go binding around an user-defined struct.
type InventoryItem struct {
	InventoryItemId *big.Int
	Owner           common.Address
	ItemId          *big.Int
	TypeId          *big.Int
	Volume          *big.Int
	Quantity        *big.Int
}

// KillMailTableData is an auto generated low-level Go binding around an user-defined struct.
type KillMailTableData struct {
	KillerCharacterId *big.Int
	VictimCharacterId *big.Int
	LossType          uint8
	SolarSystemId     *big.Int
	KillTimestamp     *big.Int
}

// LocationTableData is an auto generated low-level Go binding around an user-defined struct.
type LocationTableData struct {
	SolarSystemId *big.Int
	X             *big.Int
	Y             *big.Int
	Z             *big.Int
}

// SmartObjectData is an auto generated low-level Go binding around an user-defined struct.
type SmartObjectData struct {
	Owner    common.Address
	TokenURI string
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

// StaticDataGlobalTableData is an auto generated low-level Go binding around an user-defined struct.
type StaticDataGlobalTableData struct {
	Name    string
	Symbol  string
	BaseURI string
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

// WorldPosition is an auto generated low-level Go binding around an user-defined struct.
type WorldPosition struct {
	SolarSystemId *big.Int
	Position      Coord
}

// ERC2771WorldMetaData contains all meta data concerning the ERC2771World contract.
var ERC2771WorldMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"batchCall\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallData[]\",\"components\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchCallFrom\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallFromData[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"callFrom\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"creator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggression\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"turretOwnerCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"aggressor\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"victim\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"anchor\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationTableData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bringOffline\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bringOnline\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canJump\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureInteractionHandler\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interactionParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartGate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureSmartTurret\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorSmartGate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordData\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"worldPosition\",\"type\":\"tuple\",\"internalType\":\"structWorldPosition\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"position\",\"type\":\"tuple\",\"internalType\":\"structCoord\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDistance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorSmartStorageUnit\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordData\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"worldPosition\",\"type\":\"tuple\",\"internalType\":\"structWorldPosition\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"position\",\"type\":\"tuple\",\"internalType\":\"structCoord\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralStorageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorSmartTurret\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordData\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"worldPosition\",\"type\":\"tuple\",\"internalType\":\"structWorldPosition\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"position\",\"type\":\"tuple\",\"internalType\":\"structCoord\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndDepositItemsToEphemeralInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndDepositItemsToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createCharacter\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"corpId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecord\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"entityRecordOffchain\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordOffchainTableData\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"tokenCid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEntityRecord\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEntityRecordOffchain\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentFuelAmount\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentFuelAmountInWei\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositToEphemeralInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"destroyDeployable\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ephemeralToInventoryTransfer\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"globalPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"globalResume\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"inProximity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"turretOwnerCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turretTarget\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"inventoryToEphemeralTransfer\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outItems\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"inventoryToEphemeralTransferWithParam\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outItems\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isGateLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWithinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkSmartGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDeployable\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"fuelUnitVolumeInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacityInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDeployableToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerERC721Token\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportKill\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killMailTableData\",\"type\":\"tuple\",\"internalType\":\"structKillMailTableData\",\"components\":[{\"name\":\"killerCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"victimCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lossType\",\"type\":\"uint8\",\"internalType\":\"enumKillMailLossType\"},{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"saveLocation\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"location\",\"type\":\"tuple\",\"internalType\":\"structLocationTableData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAccessEnforcement\",\"inputs\":[{\"name\":\"target\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isEnforced\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAccessListByRole\",\"inputs\":[{\"name\":\"accessRoleId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accessList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseURI\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"baseURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCharClassId\",\"inputs\":[{\"name\":\"classId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCid\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDappURL\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDeployableMetadata\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDescription\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEntityMetadata\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEphemeralInventoryCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralStorageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFuelConsumptionPerMinute\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFuelMaxCapacity\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacityInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInventoryCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadata\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structStaticDataGlobalTableData\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"baseURI\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSSUClassId\",\"inputs\":[{\"name\":\"classId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSmartAssemblyType\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"smartAssemblyType\",\"type\":\"uint8\",\"internalType\":\"enumSmartAssemblyType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unanchor\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlinkSmartGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateCorpId\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"corpId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFromEphemeralInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFromInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldSlice\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLayout\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeySchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValueSchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialMsgSender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initModule\",\"type\":\"address\",\"internalType\":\"contractIModule\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installRootModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"popFromDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"byteLengthToPop\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pushToDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dataToPush\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"systemFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespace\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerRootFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"worldFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"systemFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystem\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"contractSystem\"},{\"name\":\"publicAccess\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerTable\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"},{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"keyNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"fieldNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storeVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferBalanceToAddress\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferBalanceToNamespace\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"worldVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"HelloStore\",\"inputs\":[{\"name\":\"storeVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HelloWorld\",\"inputs\":[{\"name\":\"worldVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_DeleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SetRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"PackedCounter\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"SmartGate_GateAlreadyLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateNotLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotWithtinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_SameSourceAndDestination\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_UndefinedClassId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartStorageUnitERC721AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartTurret_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartTurret_UndefinedClassId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Store_IndexOutOfBounds\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessedIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidBounds\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidFieldNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidKeyNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidSplice\",\"inputs\":[{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"fieldLength\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaDynamicLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaStaticLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_TableAlreadyExists\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_TableNotFound\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_AccessDenied\",\"inputs\":[{\"name\":\"resource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"World_CallbackNotAllowed\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_DelegationNotFound\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_DirectCallToSystemForbidden\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorAlreadyExists\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorNotFound\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"World_InterfaceNotSupported\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InvalidNamespace\",\"inputs\":[{\"name\":\"namespace\",\"type\":\"bytes14\",\"internalType\":\"bytes14\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceId\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceAlreadyExists\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceNotFound\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_SystemAlreadyExists\",\"inputs\":[{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_UnlimitedDelegationNotAllowed\",\"inputs\":[]}]",
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

// CurrentFuelAmount is a free data retrieval call binding the contract method 0x22d3581d.
//
// Solidity: function currentFuelAmount(uint256 entityId) view returns(uint256 amount)
func (_ERC2771World *ERC2771WorldCaller) CurrentFuelAmount(opts *bind.CallOpts, entityId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "currentFuelAmount", entityId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFuelAmount is a free data retrieval call binding the contract method 0x22d3581d.
//
// Solidity: function currentFuelAmount(uint256 entityId) view returns(uint256 amount)
func (_ERC2771World *ERC2771WorldSession) CurrentFuelAmount(entityId *big.Int) (*big.Int, error) {
	return _ERC2771World.Contract.CurrentFuelAmount(&_ERC2771World.CallOpts, entityId)
}

// CurrentFuelAmount is a free data retrieval call binding the contract method 0x22d3581d.
//
// Solidity: function currentFuelAmount(uint256 entityId) view returns(uint256 amount)
func (_ERC2771World *ERC2771WorldCallerSession) CurrentFuelAmount(entityId *big.Int) (*big.Int, error) {
	return _ERC2771World.Contract.CurrentFuelAmount(&_ERC2771World.CallOpts, entityId)
}

// CurrentFuelAmountInWei is a free data retrieval call binding the contract method 0x40bbcc11.
//
// Solidity: function currentFuelAmountInWei(uint256 entityId) view returns(uint256 amount)
func (_ERC2771World *ERC2771WorldCaller) CurrentFuelAmountInWei(opts *bind.CallOpts, entityId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "currentFuelAmountInWei", entityId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFuelAmountInWei is a free data retrieval call binding the contract method 0x40bbcc11.
//
// Solidity: function currentFuelAmountInWei(uint256 entityId) view returns(uint256 amount)
func (_ERC2771World *ERC2771WorldSession) CurrentFuelAmountInWei(entityId *big.Int) (*big.Int, error) {
	return _ERC2771World.Contract.CurrentFuelAmountInWei(&_ERC2771World.CallOpts, entityId)
}

// CurrentFuelAmountInWei is a free data retrieval call binding the contract method 0x40bbcc11.
//
// Solidity: function currentFuelAmountInWei(uint256 entityId) view returns(uint256 amount)
func (_ERC2771World *ERC2771WorldCallerSession) CurrentFuelAmountInWei(entityId *big.Int) (*big.Int, error) {
	return _ERC2771World.Contract.CurrentFuelAmountInWei(&_ERC2771World.CallOpts, entityId)
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

// InitialMsgSender is a free data retrieval call binding the contract method 0x1e71d54c.
//
// Solidity: function initialMsgSender() view returns(address)
func (_ERC2771World *ERC2771WorldCaller) InitialMsgSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC2771World.contract.Call(opts, &out, "initialMsgSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InitialMsgSender is a free data retrieval call binding the contract method 0x1e71d54c.
//
// Solidity: function initialMsgSender() view returns(address)
func (_ERC2771World *ERC2771WorldSession) InitialMsgSender() (common.Address, error) {
	return _ERC2771World.Contract.InitialMsgSender(&_ERC2771World.CallOpts)
}

// InitialMsgSender is a free data retrieval call binding the contract method 0x1e71d54c.
//
// Solidity: function initialMsgSender() view returns(address)
func (_ERC2771World *ERC2771WorldCallerSession) InitialMsgSender() (common.Address, error) {
	return _ERC2771World.Contract.InitialMsgSender(&_ERC2771World.CallOpts)
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

// Aggression is a paid mutator transaction binding the contract method 0xb9782236.
//
// Solidity: function aggression(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) aggressor, (uint256,uint256,uint256,uint256,uint256,uint256) victim) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldTransactor) Aggression(opts *bind.TransactOpts, smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, aggressor SmartTurretTarget, victim SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "aggression", smartObjectId, turretOwnerCharacterId, priorityQueue, turret, aggressor, victim)
}

// Aggression is a paid mutator transaction binding the contract method 0xb9782236.
//
// Solidity: function aggression(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) aggressor, (uint256,uint256,uint256,uint256,uint256,uint256) victim) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldSession) Aggression(smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, aggressor SmartTurretTarget, victim SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771World.Contract.Aggression(&_ERC2771World.TransactOpts, smartObjectId, turretOwnerCharacterId, priorityQueue, turret, aggressor, victim)
}

// Aggression is a paid mutator transaction binding the contract method 0xb9782236.
//
// Solidity: function aggression(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) aggressor, (uint256,uint256,uint256,uint256,uint256,uint256) victim) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldTransactorSession) Aggression(smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, aggressor SmartTurretTarget, victim SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771World.Contract.Aggression(&_ERC2771World.TransactOpts, smartObjectId, turretOwnerCharacterId, priorityQueue, turret, aggressor, victim)
}

// Anchor is a paid mutator transaction binding the contract method 0xc0395eab.
//
// Solidity: function anchor(uint256 entityId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771World *ERC2771WorldTransactor) Anchor(opts *bind.TransactOpts, entityId *big.Int, locationData LocationTableData) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "anchor", entityId, locationData)
}

// Anchor is a paid mutator transaction binding the contract method 0xc0395eab.
//
// Solidity: function anchor(uint256 entityId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771World *ERC2771WorldSession) Anchor(entityId *big.Int, locationData LocationTableData) (*types.Transaction, error) {
	return _ERC2771World.Contract.Anchor(&_ERC2771World.TransactOpts, entityId, locationData)
}

// Anchor is a paid mutator transaction binding the contract method 0xc0395eab.
//
// Solidity: function anchor(uint256 entityId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) Anchor(entityId *big.Int, locationData LocationTableData) (*types.Transaction, error) {
	return _ERC2771World.Contract.Anchor(&_ERC2771World.TransactOpts, entityId, locationData)
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
// Solidity: function bringOffline(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactor) BringOffline(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "bringOffline", entityId)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldSession) BringOffline(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.BringOffline(&_ERC2771World.TransactOpts, entityId)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) BringOffline(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.BringOffline(&_ERC2771World.TransactOpts, entityId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactor) BringOnline(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "bringOnline", entityId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldSession) BringOnline(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.BringOnline(&_ERC2771World.TransactOpts, entityId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) BringOnline(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.BringOnline(&_ERC2771World.TransactOpts, entityId)
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

// ConfigureInteractionHandler is a paid mutator transaction binding the contract method 0xb81db02a.
//
// Solidity: function configureInteractionHandler(uint256 smartObjectId, bytes interactionParams) returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureInteractionHandler(opts *bind.TransactOpts, smartObjectId *big.Int, interactionParams []byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureInteractionHandler", smartObjectId, interactionParams)
}

// ConfigureInteractionHandler is a paid mutator transaction binding the contract method 0xb81db02a.
//
// Solidity: function configureInteractionHandler(uint256 smartObjectId, bytes interactionParams) returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureInteractionHandler(smartObjectId *big.Int, interactionParams []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureInteractionHandler(&_ERC2771World.TransactOpts, smartObjectId, interactionParams)
}

// ConfigureInteractionHandler is a paid mutator transaction binding the contract method 0xb81db02a.
//
// Solidity: function configureInteractionHandler(uint256 smartObjectId, bytes interactionParams) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureInteractionHandler(smartObjectId *big.Int, interactionParams []byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureInteractionHandler(&_ERC2771World.TransactOpts, smartObjectId, interactionParams)
}

// ConfigureSmartGate is a paid mutator transaction binding the contract method 0xd2f7fd7a.
//
// Solidity: function configureSmartGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureSmartGate(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureSmartGate", smartObjectId, systemId)
}

// ConfigureSmartGate is a paid mutator transaction binding the contract method 0xd2f7fd7a.
//
// Solidity: function configureSmartGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureSmartGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartGate(&_ERC2771World.TransactOpts, smartObjectId, systemId)
}

// ConfigureSmartGate is a paid mutator transaction binding the contract method 0xd2f7fd7a.
//
// Solidity: function configureSmartGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureSmartGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartGate(&_ERC2771World.TransactOpts, smartObjectId, systemId)
}

// ConfigureSmartTurret is a paid mutator transaction binding the contract method 0x83340514.
//
// Solidity: function configureSmartTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldTransactor) ConfigureSmartTurret(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "configureSmartTurret", smartObjectId, systemId)
}

// ConfigureSmartTurret is a paid mutator transaction binding the contract method 0x83340514.
//
// Solidity: function configureSmartTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldSession) ConfigureSmartTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartTurret(&_ERC2771World.TransactOpts, smartObjectId, systemId)
}

// ConfigureSmartTurret is a paid mutator transaction binding the contract method 0x83340514.
//
// Solidity: function configureSmartTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ConfigureSmartTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.ConfigureSmartTurret(&_ERC2771World.TransactOpts, smartObjectId, systemId)
}

// CreateAndAnchorSmartGate is a paid mutator transaction binding the contract method 0xb6fbbf35.
//
// Solidity: function createAndAnchorSmartGate(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 maxDistance) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndAnchorSmartGate(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, maxDistance *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndAnchorSmartGate", smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, maxDistance)
}

// CreateAndAnchorSmartGate is a paid mutator transaction binding the contract method 0xb6fbbf35.
//
// Solidity: function createAndAnchorSmartGate(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 maxDistance) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndAnchorSmartGate(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, maxDistance *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorSmartGate(&_ERC2771World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, maxDistance)
}

// CreateAndAnchorSmartGate is a paid mutator transaction binding the contract method 0xb6fbbf35.
//
// Solidity: function createAndAnchorSmartGate(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 maxDistance) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndAnchorSmartGate(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, maxDistance *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorSmartGate(&_ERC2771World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, maxDistance)
}

// CreateAndAnchorSmartStorageUnit is a paid mutator transaction binding the contract method 0x1bd1eb3e.
//
// Solidity: function createAndAnchorSmartStorageUnit(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 storageCapacity, uint256 ephemeralStorageCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndAnchorSmartStorageUnit(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, storageCapacity *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndAnchorSmartStorageUnit", smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, storageCapacity, ephemeralStorageCapacity)
}

// CreateAndAnchorSmartStorageUnit is a paid mutator transaction binding the contract method 0x1bd1eb3e.
//
// Solidity: function createAndAnchorSmartStorageUnit(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 storageCapacity, uint256 ephemeralStorageCapacity) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndAnchorSmartStorageUnit(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, storageCapacity *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorSmartStorageUnit(&_ERC2771World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, storageCapacity, ephemeralStorageCapacity)
}

// CreateAndAnchorSmartStorageUnit is a paid mutator transaction binding the contract method 0x1bd1eb3e.
//
// Solidity: function createAndAnchorSmartStorageUnit(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 storageCapacity, uint256 ephemeralStorageCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndAnchorSmartStorageUnit(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, storageCapacity *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorSmartStorageUnit(&_ERC2771World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, storageCapacity, ephemeralStorageCapacity)
}

// CreateAndAnchorSmartTurret is a paid mutator transaction binding the contract method 0xa9ceb1a1.
//
// Solidity: function createAndAnchorSmartTurret(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndAnchorSmartTurret(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndAnchorSmartTurret", smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity)
}

// CreateAndAnchorSmartTurret is a paid mutator transaction binding the contract method 0xa9ceb1a1.
//
// Solidity: function createAndAnchorSmartTurret(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndAnchorSmartTurret(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorSmartTurret(&_ERC2771World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity)
}

// CreateAndAnchorSmartTurret is a paid mutator transaction binding the contract method 0xa9ceb1a1.
//
// Solidity: function createAndAnchorSmartTurret(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndAnchorSmartTurret(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndAnchorSmartTurret(&_ERC2771World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity)
}

// CreateAndDepositItemsToEphemeralInventory is a paid mutator transaction binding the contract method 0xe111c1a6.
//
// Solidity: function createAndDepositItemsToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndDepositItemsToEphemeralInventory(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndDepositItemsToEphemeralInventory", smartObjectId, ephemeralInventoryOwner, items)
}

// CreateAndDepositItemsToEphemeralInventory is a paid mutator transaction binding the contract method 0xe111c1a6.
//
// Solidity: function createAndDepositItemsToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndDepositItemsToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndDepositItemsToEphemeralInventory(&_ERC2771World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// CreateAndDepositItemsToEphemeralInventory is a paid mutator transaction binding the contract method 0xe111c1a6.
//
// Solidity: function createAndDepositItemsToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndDepositItemsToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndDepositItemsToEphemeralInventory(&_ERC2771World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// CreateAndDepositItemsToInventory is a paid mutator transaction binding the contract method 0x9e19e7aa.
//
// Solidity: function createAndDepositItemsToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateAndDepositItemsToInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createAndDepositItemsToInventory", smartObjectId, items)
}

// CreateAndDepositItemsToInventory is a paid mutator transaction binding the contract method 0x9e19e7aa.
//
// Solidity: function createAndDepositItemsToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) CreateAndDepositItemsToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndDepositItemsToInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// CreateAndDepositItemsToInventory is a paid mutator transaction binding the contract method 0x9e19e7aa.
//
// Solidity: function createAndDepositItemsToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateAndDepositItemsToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateAndDepositItemsToInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x036bb5d5.
//
// Solidity: function createCharacter(uint256 characterId, address characterAddress, uint256 corpId, (uint256,uint256,uint256) entityRecord, (string,string,string) entityRecordOffchain, string tokenCid) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateCharacter(opts *bind.TransactOpts, characterId *big.Int, characterAddress common.Address, corpId *big.Int, entityRecord EntityRecordData, entityRecordOffchain EntityRecordOffchainTableData, tokenCid string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createCharacter", characterId, characterAddress, corpId, entityRecord, entityRecordOffchain, tokenCid)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x036bb5d5.
//
// Solidity: function createCharacter(uint256 characterId, address characterAddress, uint256 corpId, (uint256,uint256,uint256) entityRecord, (string,string,string) entityRecordOffchain, string tokenCid) returns()
func (_ERC2771World *ERC2771WorldSession) CreateCharacter(characterId *big.Int, characterAddress common.Address, corpId *big.Int, entityRecord EntityRecordData, entityRecordOffchain EntityRecordOffchainTableData, tokenCid string) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateCharacter(&_ERC2771World.TransactOpts, characterId, characterAddress, corpId, entityRecord, entityRecordOffchain, tokenCid)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x036bb5d5.
//
// Solidity: function createCharacter(uint256 characterId, address characterAddress, uint256 corpId, (uint256,uint256,uint256) entityRecord, (string,string,string) entityRecordOffchain, string tokenCid) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateCharacter(characterId *big.Int, characterAddress common.Address, corpId *big.Int, entityRecord EntityRecordData, entityRecordOffchain EntityRecordOffchainTableData, tokenCid string) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateCharacter(&_ERC2771World.TransactOpts, characterId, characterAddress, corpId, entityRecord, entityRecordOffchain, tokenCid)
}

// CreateEntityRecord is a paid mutator transaction binding the contract method 0x2c63f58f.
//
// Solidity: function createEntityRecord(uint256 entityId, uint256 itemId, uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateEntityRecord(opts *bind.TransactOpts, entityId *big.Int, itemId *big.Int, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createEntityRecord", entityId, itemId, typeId, volume)
}

// CreateEntityRecord is a paid mutator transaction binding the contract method 0x2c63f58f.
//
// Solidity: function createEntityRecord(uint256 entityId, uint256 itemId, uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldSession) CreateEntityRecord(entityId *big.Int, itemId *big.Int, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateEntityRecord(&_ERC2771World.TransactOpts, entityId, itemId, typeId, volume)
}

// CreateEntityRecord is a paid mutator transaction binding the contract method 0x2c63f58f.
//
// Solidity: function createEntityRecord(uint256 entityId, uint256 itemId, uint256 typeId, uint256 volume) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateEntityRecord(entityId *big.Int, itemId *big.Int, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateEntityRecord(&_ERC2771World.TransactOpts, entityId, itemId, typeId, volume)
}

// CreateEntityRecordOffchain is a paid mutator transaction binding the contract method 0xbf662710.
//
// Solidity: function createEntityRecordOffchain(uint256 entityId, string name, string dappURL, string description) returns()
func (_ERC2771World *ERC2771WorldTransactor) CreateEntityRecordOffchain(opts *bind.TransactOpts, entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "createEntityRecordOffchain", entityId, name, dappURL, description)
}

// CreateEntityRecordOffchain is a paid mutator transaction binding the contract method 0xbf662710.
//
// Solidity: function createEntityRecordOffchain(uint256 entityId, string name, string dappURL, string description) returns()
func (_ERC2771World *ERC2771WorldSession) CreateEntityRecordOffchain(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateEntityRecordOffchain(&_ERC2771World.TransactOpts, entityId, name, dappURL, description)
}

// CreateEntityRecordOffchain is a paid mutator transaction binding the contract method 0xbf662710.
//
// Solidity: function createEntityRecordOffchain(uint256 entityId, string name, string dappURL, string description) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) CreateEntityRecordOffchain(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.CreateEntityRecordOffchain(&_ERC2771World.TransactOpts, entityId, name, dappURL, description)
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

// DepositFuel is a paid mutator transaction binding the contract method 0xc3e9a45f.
//
// Solidity: function depositFuel(uint256 entityId, uint256 unitAmount) returns()
func (_ERC2771World *ERC2771WorldTransactor) DepositFuel(opts *bind.TransactOpts, entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "depositFuel", entityId, unitAmount)
}

// DepositFuel is a paid mutator transaction binding the contract method 0xc3e9a45f.
//
// Solidity: function depositFuel(uint256 entityId, uint256 unitAmount) returns()
func (_ERC2771World *ERC2771WorldSession) DepositFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositFuel(&_ERC2771World.TransactOpts, entityId, unitAmount)
}

// DepositFuel is a paid mutator transaction binding the contract method 0xc3e9a45f.
//
// Solidity: function depositFuel(uint256 entityId, uint256 unitAmount) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) DepositFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositFuel(&_ERC2771World.TransactOpts, entityId, unitAmount)
}

// DepositToEphemeralInventory is a paid mutator transaction binding the contract method 0xaff42af2.
//
// Solidity: function depositToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) DepositToEphemeralInventory(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "depositToEphemeralInventory", smartObjectId, ephemeralInventoryOwner, items)
}

// DepositToEphemeralInventory is a paid mutator transaction binding the contract method 0xaff42af2.
//
// Solidity: function depositToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) DepositToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositToEphemeralInventory(&_ERC2771World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// DepositToEphemeralInventory is a paid mutator transaction binding the contract method 0xaff42af2.
//
// Solidity: function depositToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) DepositToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositToEphemeralInventory(&_ERC2771World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// DepositToInventory is a paid mutator transaction binding the contract method 0x15e306bb.
//
// Solidity: function depositToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) DepositToInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "depositToInventory", smartObjectId, items)
}

// DepositToInventory is a paid mutator transaction binding the contract method 0x15e306bb.
//
// Solidity: function depositToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) DepositToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositToInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// DepositToInventory is a paid mutator transaction binding the contract method 0x15e306bb.
//
// Solidity: function depositToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) DepositToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.DepositToInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactor) DestroyDeployable(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "destroyDeployable", entityId)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldSession) DestroyDeployable(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.DestroyDeployable(&_ERC2771World.TransactOpts, entityId)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) DestroyDeployable(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.DestroyDeployable(&_ERC2771World.TransactOpts, entityId)
}

// EphemeralToInventoryTransfer is a paid mutator transaction binding the contract method 0x6f89f6c3.
//
// Solidity: function ephemeralToInventoryTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) EphemeralToInventoryTransfer(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "ephemeralToInventoryTransfer", smartObjectId, items)
}

// EphemeralToInventoryTransfer is a paid mutator transaction binding the contract method 0x6f89f6c3.
//
// Solidity: function ephemeralToInventoryTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) EphemeralToInventoryTransfer(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.EphemeralToInventoryTransfer(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// EphemeralToInventoryTransfer is a paid mutator transaction binding the contract method 0x6f89f6c3.
//
// Solidity: function ephemeralToInventoryTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) EphemeralToInventoryTransfer(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.EphemeralToInventoryTransfer(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// GlobalPause is a paid mutator transaction binding the contract method 0xf12d54d8.
//
// Solidity: function globalPause() returns()
func (_ERC2771World *ERC2771WorldTransactor) GlobalPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "globalPause")
}

// GlobalPause is a paid mutator transaction binding the contract method 0xf12d54d8.
//
// Solidity: function globalPause() returns()
func (_ERC2771World *ERC2771WorldSession) GlobalPause() (*types.Transaction, error) {
	return _ERC2771World.Contract.GlobalPause(&_ERC2771World.TransactOpts)
}

// GlobalPause is a paid mutator transaction binding the contract method 0xf12d54d8.
//
// Solidity: function globalPause() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) GlobalPause() (*types.Transaction, error) {
	return _ERC2771World.Contract.GlobalPause(&_ERC2771World.TransactOpts)
}

// GlobalResume is a paid mutator transaction binding the contract method 0x59c7d378.
//
// Solidity: function globalResume() returns()
func (_ERC2771World *ERC2771WorldTransactor) GlobalResume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "globalResume")
}

// GlobalResume is a paid mutator transaction binding the contract method 0x59c7d378.
//
// Solidity: function globalResume() returns()
func (_ERC2771World *ERC2771WorldSession) GlobalResume() (*types.Transaction, error) {
	return _ERC2771World.Contract.GlobalResume(&_ERC2771World.TransactOpts)
}

// GlobalResume is a paid mutator transaction binding the contract method 0x59c7d378.
//
// Solidity: function globalResume() returns()
func (_ERC2771World *ERC2771WorldTransactorSession) GlobalResume() (*types.Transaction, error) {
	return _ERC2771World.Contract.GlobalResume(&_ERC2771World.TransactOpts)
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

// InProximity is a paid mutator transaction binding the contract method 0xea49f0a8.
//
// Solidity: function inProximity(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldTransactor) InProximity(opts *bind.TransactOpts, smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "inProximity", smartObjectId, turretOwnerCharacterId, priorityQueue, turret, turretTarget)
}

// InProximity is a paid mutator transaction binding the contract method 0xea49f0a8.
//
// Solidity: function inProximity(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldSession) InProximity(smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771World.Contract.InProximity(&_ERC2771World.TransactOpts, smartObjectId, turretOwnerCharacterId, priorityQueue, turret, turretTarget)
}

// InProximity is a paid mutator transaction binding the contract method 0xea49f0a8.
//
// Solidity: function inProximity(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_ERC2771World *ERC2771WorldTransactorSession) InProximity(smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _ERC2771World.Contract.InProximity(&_ERC2771World.TransactOpts, smartObjectId, turretOwnerCharacterId, priorityQueue, turret, turretTarget)
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

// InventoryToEphemeralTransfer is a paid mutator transaction binding the contract method 0x181c3771.
//
// Solidity: function inventoryToEphemeralTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_ERC2771World *ERC2771WorldTransactor) InventoryToEphemeralTransfer(opts *bind.TransactOpts, smartObjectId *big.Int, outItems []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "inventoryToEphemeralTransfer", smartObjectId, outItems)
}

// InventoryToEphemeralTransfer is a paid mutator transaction binding the contract method 0x181c3771.
//
// Solidity: function inventoryToEphemeralTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_ERC2771World *ERC2771WorldSession) InventoryToEphemeralTransfer(smartObjectId *big.Int, outItems []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.InventoryToEphemeralTransfer(&_ERC2771World.TransactOpts, smartObjectId, outItems)
}

// InventoryToEphemeralTransfer is a paid mutator transaction binding the contract method 0x181c3771.
//
// Solidity: function inventoryToEphemeralTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) InventoryToEphemeralTransfer(smartObjectId *big.Int, outItems []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.InventoryToEphemeralTransfer(&_ERC2771World.TransactOpts, smartObjectId, outItems)
}

// InventoryToEphemeralTransferWithParam is a paid mutator transaction binding the contract method 0x870642b6.
//
// Solidity: function inventoryToEphemeralTransferWithParam(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_ERC2771World *ERC2771WorldTransactor) InventoryToEphemeralTransferWithParam(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, outItems []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "inventoryToEphemeralTransferWithParam", smartObjectId, ephemeralInventoryOwner, outItems)
}

// InventoryToEphemeralTransferWithParam is a paid mutator transaction binding the contract method 0x870642b6.
//
// Solidity: function inventoryToEphemeralTransferWithParam(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_ERC2771World *ERC2771WorldSession) InventoryToEphemeralTransferWithParam(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, outItems []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.InventoryToEphemeralTransferWithParam(&_ERC2771World.TransactOpts, smartObjectId, ephemeralInventoryOwner, outItems)
}

// InventoryToEphemeralTransferWithParam is a paid mutator transaction binding the contract method 0x870642b6.
//
// Solidity: function inventoryToEphemeralTransferWithParam(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) InventoryToEphemeralTransferWithParam(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, outItems []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.InventoryToEphemeralTransferWithParam(&_ERC2771World.TransactOpts, smartObjectId, ephemeralInventoryOwner, outItems)
}

// LinkSmartGates is a paid mutator transaction binding the contract method 0xd40adbfb.
//
// Solidity: function linkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldTransactor) LinkSmartGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "linkSmartGates", sourceGateId, destinationGateId)
}

// LinkSmartGates is a paid mutator transaction binding the contract method 0xd40adbfb.
//
// Solidity: function linkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldSession) LinkSmartGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.LinkSmartGates(&_ERC2771World.TransactOpts, sourceGateId, destinationGateId)
}

// LinkSmartGates is a paid mutator transaction binding the contract method 0xd40adbfb.
//
// Solidity: function linkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) LinkSmartGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.LinkSmartGates(&_ERC2771World.TransactOpts, sourceGateId, destinationGateId)
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

// RegisterDeployable is a paid mutator transaction binding the contract method 0x922bd0ca.
//
// Solidity: function registerDeployable(uint256 entityId, (address,string) smartObjectData, uint256 fuelUnitVolumeInWei, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacityInWei) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterDeployable(opts *bind.TransactOpts, entityId *big.Int, smartObjectData SmartObjectData, fuelUnitVolumeInWei *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacityInWei *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerDeployable", entityId, smartObjectData, fuelUnitVolumeInWei, fuelConsumptionIntervalInSeconds, fuelMaxCapacityInWei)
}

// RegisterDeployable is a paid mutator transaction binding the contract method 0x922bd0ca.
//
// Solidity: function registerDeployable(uint256 entityId, (address,string) smartObjectData, uint256 fuelUnitVolumeInWei, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacityInWei) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterDeployable(entityId *big.Int, smartObjectData SmartObjectData, fuelUnitVolumeInWei *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacityInWei *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterDeployable(&_ERC2771World.TransactOpts, entityId, smartObjectData, fuelUnitVolumeInWei, fuelConsumptionIntervalInSeconds, fuelMaxCapacityInWei)
}

// RegisterDeployable is a paid mutator transaction binding the contract method 0x922bd0ca.
//
// Solidity: function registerDeployable(uint256 entityId, (address,string) smartObjectData, uint256 fuelUnitVolumeInWei, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacityInWei) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterDeployable(entityId *big.Int, smartObjectData SmartObjectData, fuelUnitVolumeInWei *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacityInWei *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterDeployable(&_ERC2771World.TransactOpts, entityId, smartObjectData, fuelUnitVolumeInWei, fuelConsumptionIntervalInSeconds, fuelMaxCapacityInWei)
}

// RegisterDeployableToken is a paid mutator transaction binding the contract method 0x7b0d0e3c.
//
// Solidity: function registerDeployableToken(address tokenAddress) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterDeployableToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerDeployableToken", tokenAddress)
}

// RegisterDeployableToken is a paid mutator transaction binding the contract method 0x7b0d0e3c.
//
// Solidity: function registerDeployableToken(address tokenAddress) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterDeployableToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterDeployableToken(&_ERC2771World.TransactOpts, tokenAddress)
}

// RegisterDeployableToken is a paid mutator transaction binding the contract method 0x7b0d0e3c.
//
// Solidity: function registerDeployableToken(address tokenAddress) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterDeployableToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterDeployableToken(&_ERC2771World.TransactOpts, tokenAddress)
}

// RegisterERC721Token is a paid mutator transaction binding the contract method 0x2c3309d1.
//
// Solidity: function registerERC721Token(address tokenAddress) returns()
func (_ERC2771World *ERC2771WorldTransactor) RegisterERC721Token(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerERC721Token", tokenAddress)
}

// RegisterERC721Token is a paid mutator transaction binding the contract method 0x2c3309d1.
//
// Solidity: function registerERC721Token(address tokenAddress) returns()
func (_ERC2771World *ERC2771WorldSession) RegisterERC721Token(tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterERC721Token(&_ERC2771World.TransactOpts, tokenAddress)
}

// RegisterERC721Token is a paid mutator transaction binding the contract method 0x2c3309d1.
//
// Solidity: function registerERC721Token(address tokenAddress) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterERC721Token(tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterERC721Token(&_ERC2771World.TransactOpts, tokenAddress)
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

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x742d6118.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, bytes4 systemFunctionSelector) returns(bytes4 worldFunctionSelector)
func (_ERC2771World *ERC2771WorldTransactor) RegisterRootFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, worldFunctionSignature string, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "registerRootFunctionSelector", systemId, worldFunctionSignature, systemFunctionSelector)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x742d6118.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, bytes4 systemFunctionSelector) returns(bytes4 worldFunctionSelector)
func (_ERC2771World *ERC2771WorldSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterRootFunctionSelector(&_ERC2771World.TransactOpts, systemId, worldFunctionSignature, systemFunctionSelector)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x742d6118.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, bytes4 systemFunctionSelector) returns(bytes4 worldFunctionSelector)
func (_ERC2771World *ERC2771WorldTransactorSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _ERC2771World.Contract.RegisterRootFunctionSelector(&_ERC2771World.TransactOpts, systemId, worldFunctionSignature, systemFunctionSelector)
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
// Solidity: function reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailTableData) returns()
func (_ERC2771World *ERC2771WorldTransactor) ReportKill(opts *bind.TransactOpts, killMailId *big.Int, killMailTableData KillMailTableData) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "reportKill", killMailId, killMailTableData)
}

// ReportKill is a paid mutator transaction binding the contract method 0xda3293e1.
//
// Solidity: function reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailTableData) returns()
func (_ERC2771World *ERC2771WorldSession) ReportKill(killMailId *big.Int, killMailTableData KillMailTableData) (*types.Transaction, error) {
	return _ERC2771World.Contract.ReportKill(&_ERC2771World.TransactOpts, killMailId, killMailTableData)
}

// ReportKill is a paid mutator transaction binding the contract method 0xda3293e1.
//
// Solidity: function reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailTableData) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) ReportKill(killMailId *big.Int, killMailTableData KillMailTableData) (*types.Transaction, error) {
	return _ERC2771World.Contract.ReportKill(&_ERC2771World.TransactOpts, killMailId, killMailTableData)
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
// Solidity: function saveLocation(uint256 entityId, (uint256,uint256,uint256,uint256) location) returns()
func (_ERC2771World *ERC2771WorldTransactor) SaveLocation(opts *bind.TransactOpts, entityId *big.Int, location LocationTableData) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "saveLocation", entityId, location)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 entityId, (uint256,uint256,uint256,uint256) location) returns()
func (_ERC2771World *ERC2771WorldSession) SaveLocation(entityId *big.Int, location LocationTableData) (*types.Transaction, error) {
	return _ERC2771World.Contract.SaveLocation(&_ERC2771World.TransactOpts, entityId, location)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 entityId, (uint256,uint256,uint256,uint256) location) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SaveLocation(entityId *big.Int, location LocationTableData) (*types.Transaction, error) {
	return _ERC2771World.Contract.SaveLocation(&_ERC2771World.TransactOpts, entityId, location)
}

// SetAccessEnforcement is a paid mutator transaction binding the contract method 0x572b9514.
//
// Solidity: function setAccessEnforcement(bytes32 target, bool isEnforced) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetAccessEnforcement(opts *bind.TransactOpts, target [32]byte, isEnforced bool) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setAccessEnforcement", target, isEnforced)
}

// SetAccessEnforcement is a paid mutator transaction binding the contract method 0x572b9514.
//
// Solidity: function setAccessEnforcement(bytes32 target, bool isEnforced) returns()
func (_ERC2771World *ERC2771WorldSession) SetAccessEnforcement(target [32]byte, isEnforced bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetAccessEnforcement(&_ERC2771World.TransactOpts, target, isEnforced)
}

// SetAccessEnforcement is a paid mutator transaction binding the contract method 0x572b9514.
//
// Solidity: function setAccessEnforcement(bytes32 target, bool isEnforced) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetAccessEnforcement(target [32]byte, isEnforced bool) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetAccessEnforcement(&_ERC2771World.TransactOpts, target, isEnforced)
}

// SetAccessListByRole is a paid mutator transaction binding the contract method 0x9991d221.
//
// Solidity: function setAccessListByRole(bytes32 accessRoleId, address[] accessList) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetAccessListByRole(opts *bind.TransactOpts, accessRoleId [32]byte, accessList []common.Address) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setAccessListByRole", accessRoleId, accessList)
}

// SetAccessListByRole is a paid mutator transaction binding the contract method 0x9991d221.
//
// Solidity: function setAccessListByRole(bytes32 accessRoleId, address[] accessList) returns()
func (_ERC2771World *ERC2771WorldSession) SetAccessListByRole(accessRoleId [32]byte, accessList []common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetAccessListByRole(&_ERC2771World.TransactOpts, accessRoleId, accessList)
}

// SetAccessListByRole is a paid mutator transaction binding the contract method 0x9991d221.
//
// Solidity: function setAccessListByRole(bytes32 accessRoleId, address[] accessList) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetAccessListByRole(accessRoleId [32]byte, accessList []common.Address) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetAccessListByRole(&_ERC2771World.TransactOpts, accessRoleId, accessList)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x8bf3d594.
//
// Solidity: function setBaseURI(bytes32 systemId, string baseURI) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetBaseURI(opts *bind.TransactOpts, systemId [32]byte, baseURI string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setBaseURI", systemId, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x8bf3d594.
//
// Solidity: function setBaseURI(bytes32 systemId, string baseURI) returns()
func (_ERC2771World *ERC2771WorldSession) SetBaseURI(systemId [32]byte, baseURI string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetBaseURI(&_ERC2771World.TransactOpts, systemId, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x8bf3d594.
//
// Solidity: function setBaseURI(bytes32 systemId, string baseURI) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetBaseURI(systemId [32]byte, baseURI string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetBaseURI(&_ERC2771World.TransactOpts, systemId, baseURI)
}

// SetCharClassId is a paid mutator transaction binding the contract method 0xf83a5887.
//
// Solidity: function setCharClassId(uint256 classId) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetCharClassId(opts *bind.TransactOpts, classId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setCharClassId", classId)
}

// SetCharClassId is a paid mutator transaction binding the contract method 0xf83a5887.
//
// Solidity: function setCharClassId(uint256 classId) returns()
func (_ERC2771World *ERC2771WorldSession) SetCharClassId(classId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetCharClassId(&_ERC2771World.TransactOpts, classId)
}

// SetCharClassId is a paid mutator transaction binding the contract method 0xf83a5887.
//
// Solidity: function setCharClassId(uint256 classId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetCharClassId(classId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetCharClassId(&_ERC2771World.TransactOpts, classId)
}

// SetCid is a paid mutator transaction binding the contract method 0x1cc5fe59.
//
// Solidity: function setCid(uint256 entityId, string cid) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetCid(opts *bind.TransactOpts, entityId *big.Int, cid string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setCid", entityId, cid)
}

// SetCid is a paid mutator transaction binding the contract method 0x1cc5fe59.
//
// Solidity: function setCid(uint256 entityId, string cid) returns()
func (_ERC2771World *ERC2771WorldSession) SetCid(entityId *big.Int, cid string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetCid(&_ERC2771World.TransactOpts, entityId, cid)
}

// SetCid is a paid mutator transaction binding the contract method 0x1cc5fe59.
//
// Solidity: function setCid(uint256 entityId, string cid) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetCid(entityId *big.Int, cid string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetCid(&_ERC2771World.TransactOpts, entityId, cid)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 entityId, string dappURL) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetDappURL(opts *bind.TransactOpts, entityId *big.Int, dappURL string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setDappURL", entityId, dappURL)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 entityId, string dappURL) returns()
func (_ERC2771World *ERC2771WorldSession) SetDappURL(entityId *big.Int, dappURL string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDappURL(&_ERC2771World.TransactOpts, entityId, dappURL)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 entityId, string dappURL) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetDappURL(entityId *big.Int, dappURL string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDappURL(&_ERC2771World.TransactOpts, entityId, dappURL)
}

// SetDeployableMetadata is a paid mutator transaction binding the contract method 0x8418f4cf.
//
// Solidity: function setDeployableMetadata(uint256 smartObjectId, string name, string dappURL, string description) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetDeployableMetadata(opts *bind.TransactOpts, smartObjectId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setDeployableMetadata", smartObjectId, name, dappURL, description)
}

// SetDeployableMetadata is a paid mutator transaction binding the contract method 0x8418f4cf.
//
// Solidity: function setDeployableMetadata(uint256 smartObjectId, string name, string dappURL, string description) returns()
func (_ERC2771World *ERC2771WorldSession) SetDeployableMetadata(smartObjectId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDeployableMetadata(&_ERC2771World.TransactOpts, smartObjectId, name, dappURL, description)
}

// SetDeployableMetadata is a paid mutator transaction binding the contract method 0x8418f4cf.
//
// Solidity: function setDeployableMetadata(uint256 smartObjectId, string name, string dappURL, string description) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetDeployableMetadata(smartObjectId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDeployableMetadata(&_ERC2771World.TransactOpts, smartObjectId, name, dappURL, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 entityId, string description) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetDescription(opts *bind.TransactOpts, entityId *big.Int, description string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setDescription", entityId, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 entityId, string description) returns()
func (_ERC2771World *ERC2771WorldSession) SetDescription(entityId *big.Int, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDescription(&_ERC2771World.TransactOpts, entityId, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 entityId, string description) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetDescription(entityId *big.Int, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetDescription(&_ERC2771World.TransactOpts, entityId, description)
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

// SetEntityMetadata is a paid mutator transaction binding the contract method 0x5f9c496a.
//
// Solidity: function setEntityMetadata(uint256 entityId, string name, string dappURL, string description) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetEntityMetadata(opts *bind.TransactOpts, entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setEntityMetadata", entityId, name, dappURL, description)
}

// SetEntityMetadata is a paid mutator transaction binding the contract method 0x5f9c496a.
//
// Solidity: function setEntityMetadata(uint256 entityId, string name, string dappURL, string description) returns()
func (_ERC2771World *ERC2771WorldSession) SetEntityMetadata(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetEntityMetadata(&_ERC2771World.TransactOpts, entityId, name, dappURL, description)
}

// SetEntityMetadata is a paid mutator transaction binding the contract method 0x5f9c496a.
//
// Solidity: function setEntityMetadata(uint256 entityId, string name, string dappURL, string description) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetEntityMetadata(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetEntityMetadata(&_ERC2771World.TransactOpts, entityId, name, dappURL, description)
}

// SetEphemeralInventoryCapacity is a paid mutator transaction binding the contract method 0x20305602.
//
// Solidity: function setEphemeralInventoryCapacity(uint256 smartObjectId, uint256 ephemeralStorageCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetEphemeralInventoryCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setEphemeralInventoryCapacity", smartObjectId, ephemeralStorageCapacity)
}

// SetEphemeralInventoryCapacity is a paid mutator transaction binding the contract method 0x20305602.
//
// Solidity: function setEphemeralInventoryCapacity(uint256 smartObjectId, uint256 ephemeralStorageCapacity) returns()
func (_ERC2771World *ERC2771WorldSession) SetEphemeralInventoryCapacity(smartObjectId *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetEphemeralInventoryCapacity(&_ERC2771World.TransactOpts, smartObjectId, ephemeralStorageCapacity)
}

// SetEphemeralInventoryCapacity is a paid mutator transaction binding the contract method 0x20305602.
//
// Solidity: function setEphemeralInventoryCapacity(uint256 smartObjectId, uint256 ephemeralStorageCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetEphemeralInventoryCapacity(smartObjectId *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetEphemeralInventoryCapacity(&_ERC2771World.TransactOpts, smartObjectId, ephemeralStorageCapacity)
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

// SetFuelConsumptionPerMinute is a paid mutator transaction binding the contract method 0xeb37e8f0.
//
// Solidity: function setFuelConsumptionPerMinute(uint256 entityId, uint256 fuelConsumptionIntervalInSeconds) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetFuelConsumptionPerMinute(opts *bind.TransactOpts, entityId *big.Int, fuelConsumptionIntervalInSeconds *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setFuelConsumptionPerMinute", entityId, fuelConsumptionIntervalInSeconds)
}

// SetFuelConsumptionPerMinute is a paid mutator transaction binding the contract method 0xeb37e8f0.
//
// Solidity: function setFuelConsumptionPerMinute(uint256 entityId, uint256 fuelConsumptionIntervalInSeconds) returns()
func (_ERC2771World *ERC2771WorldSession) SetFuelConsumptionPerMinute(entityId *big.Int, fuelConsumptionIntervalInSeconds *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetFuelConsumptionPerMinute(&_ERC2771World.TransactOpts, entityId, fuelConsumptionIntervalInSeconds)
}

// SetFuelConsumptionPerMinute is a paid mutator transaction binding the contract method 0xeb37e8f0.
//
// Solidity: function setFuelConsumptionPerMinute(uint256 entityId, uint256 fuelConsumptionIntervalInSeconds) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetFuelConsumptionPerMinute(entityId *big.Int, fuelConsumptionIntervalInSeconds *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetFuelConsumptionPerMinute(&_ERC2771World.TransactOpts, entityId, fuelConsumptionIntervalInSeconds)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 entityId, uint256 capacityInWei) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetFuelMaxCapacity(opts *bind.TransactOpts, entityId *big.Int, capacityInWei *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setFuelMaxCapacity", entityId, capacityInWei)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 entityId, uint256 capacityInWei) returns()
func (_ERC2771World *ERC2771WorldSession) SetFuelMaxCapacity(entityId *big.Int, capacityInWei *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetFuelMaxCapacity(&_ERC2771World.TransactOpts, entityId, capacityInWei)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 entityId, uint256 capacityInWei) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetFuelMaxCapacity(entityId *big.Int, capacityInWei *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetFuelMaxCapacity(&_ERC2771World.TransactOpts, entityId, capacityInWei)
}

// SetInventoryCapacity is a paid mutator transaction binding the contract method 0x21b01b1d.
//
// Solidity: function setInventoryCapacity(uint256 smartObjectId, uint256 storageCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetInventoryCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, storageCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setInventoryCapacity", smartObjectId, storageCapacity)
}

// SetInventoryCapacity is a paid mutator transaction binding the contract method 0x21b01b1d.
//
// Solidity: function setInventoryCapacity(uint256 smartObjectId, uint256 storageCapacity) returns()
func (_ERC2771World *ERC2771WorldSession) SetInventoryCapacity(smartObjectId *big.Int, storageCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetInventoryCapacity(&_ERC2771World.TransactOpts, smartObjectId, storageCapacity)
}

// SetInventoryCapacity is a paid mutator transaction binding the contract method 0x21b01b1d.
//
// Solidity: function setInventoryCapacity(uint256 smartObjectId, uint256 storageCapacity) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetInventoryCapacity(smartObjectId *big.Int, storageCapacity *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetInventoryCapacity(&_ERC2771World.TransactOpts, smartObjectId, storageCapacity)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x7eaf1400.
//
// Solidity: function setMetadata(bytes32 systemId, (string,string,string) data) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetMetadata(opts *bind.TransactOpts, systemId [32]byte, data StaticDataGlobalTableData) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setMetadata", systemId, data)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x7eaf1400.
//
// Solidity: function setMetadata(bytes32 systemId, (string,string,string) data) returns()
func (_ERC2771World *ERC2771WorldSession) SetMetadata(systemId [32]byte, data StaticDataGlobalTableData) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetMetadata(&_ERC2771World.TransactOpts, systemId, data)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x7eaf1400.
//
// Solidity: function setMetadata(bytes32 systemId, (string,string,string) data) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetMetadata(systemId [32]byte, data StaticDataGlobalTableData) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetMetadata(&_ERC2771World.TransactOpts, systemId, data)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 entityId, string name) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetName(opts *bind.TransactOpts, entityId *big.Int, name string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setName", entityId, name)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 entityId, string name) returns()
func (_ERC2771World *ERC2771WorldSession) SetName(entityId *big.Int, name string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetName(&_ERC2771World.TransactOpts, entityId, name)
}

// SetName is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 entityId, string name) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetName(entityId *big.Int, name string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetName(&_ERC2771World.TransactOpts, entityId, name)
}

// SetName0 is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 systemId, string name) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetName0(opts *bind.TransactOpts, systemId [32]byte, name string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setName0", systemId, name)
}

// SetName0 is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 systemId, string name) returns()
func (_ERC2771World *ERC2771WorldSession) SetName0(systemId [32]byte, name string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetName0(&_ERC2771World.TransactOpts, systemId, name)
}

// SetName0 is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 systemId, string name) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetName0(systemId [32]byte, name string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetName0(&_ERC2771World.TransactOpts, systemId, name)
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

// SetSSUClassId is a paid mutator transaction binding the contract method 0xe2afb03b.
//
// Solidity: function setSSUClassId(uint256 classId) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetSSUClassId(opts *bind.TransactOpts, classId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setSSUClassId", classId)
}

// SetSSUClassId is a paid mutator transaction binding the contract method 0xe2afb03b.
//
// Solidity: function setSSUClassId(uint256 classId) returns()
func (_ERC2771World *ERC2771WorldSession) SetSSUClassId(classId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetSSUClassId(&_ERC2771World.TransactOpts, classId)
}

// SetSSUClassId is a paid mutator transaction binding the contract method 0xe2afb03b.
//
// Solidity: function setSSUClassId(uint256 classId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetSSUClassId(classId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetSSUClassId(&_ERC2771World.TransactOpts, classId)
}

// SetSmartAssemblyType is a paid mutator transaction binding the contract method 0xff6f1664.
//
// Solidity: function setSmartAssemblyType(uint256 entityId, uint8 smartAssemblyType) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetSmartAssemblyType(opts *bind.TransactOpts, entityId *big.Int, smartAssemblyType uint8) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setSmartAssemblyType", entityId, smartAssemblyType)
}

// SetSmartAssemblyType is a paid mutator transaction binding the contract method 0xff6f1664.
//
// Solidity: function setSmartAssemblyType(uint256 entityId, uint8 smartAssemblyType) returns()
func (_ERC2771World *ERC2771WorldSession) SetSmartAssemblyType(entityId *big.Int, smartAssemblyType uint8) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetSmartAssemblyType(&_ERC2771World.TransactOpts, entityId, smartAssemblyType)
}

// SetSmartAssemblyType is a paid mutator transaction binding the contract method 0xff6f1664.
//
// Solidity: function setSmartAssemblyType(uint256 entityId, uint8 smartAssemblyType) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetSmartAssemblyType(entityId *big.Int, smartAssemblyType uint8) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetSmartAssemblyType(&_ERC2771World.TransactOpts, entityId, smartAssemblyType)
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

// SetSymbol is a paid mutator transaction binding the contract method 0x0b1cb716.
//
// Solidity: function setSymbol(bytes32 systemId, string symbol) returns()
func (_ERC2771World *ERC2771WorldTransactor) SetSymbol(opts *bind.TransactOpts, systemId [32]byte, symbol string) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "setSymbol", systemId, symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0x0b1cb716.
//
// Solidity: function setSymbol(bytes32 systemId, string symbol) returns()
func (_ERC2771World *ERC2771WorldSession) SetSymbol(systemId [32]byte, symbol string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetSymbol(&_ERC2771World.TransactOpts, systemId, symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0x0b1cb716.
//
// Solidity: function setSymbol(bytes32 systemId, string symbol) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) SetSymbol(systemId [32]byte, symbol string) (*types.Transaction, error) {
	return _ERC2771World.Contract.SetSymbol(&_ERC2771World.TransactOpts, systemId, symbol)
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

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactor) Unanchor(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "unanchor", entityId)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldSession) Unanchor(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.Unanchor(&_ERC2771World.TransactOpts, entityId)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) Unanchor(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.Unanchor(&_ERC2771World.TransactOpts, entityId)
}

// UnlinkSmartGates is a paid mutator transaction binding the contract method 0xa51dc713.
//
// Solidity: function unlinkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldTransactor) UnlinkSmartGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "unlinkSmartGates", sourceGateId, destinationGateId)
}

// UnlinkSmartGates is a paid mutator transaction binding the contract method 0xa51dc713.
//
// Solidity: function unlinkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldSession) UnlinkSmartGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnlinkSmartGates(&_ERC2771World.TransactOpts, sourceGateId, destinationGateId)
}

// UnlinkSmartGates is a paid mutator transaction binding the contract method 0xa51dc713.
//
// Solidity: function unlinkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UnlinkSmartGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UnlinkSmartGates(&_ERC2771World.TransactOpts, sourceGateId, destinationGateId)
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

// UpdateCorpId is a paid mutator transaction binding the contract method 0x325675d4.
//
// Solidity: function updateCorpId(uint256 characterId, uint256 corpId) returns()
func (_ERC2771World *ERC2771WorldTransactor) UpdateCorpId(opts *bind.TransactOpts, characterId *big.Int, corpId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "updateCorpId", characterId, corpId)
}

// UpdateCorpId is a paid mutator transaction binding the contract method 0x325675d4.
//
// Solidity: function updateCorpId(uint256 characterId, uint256 corpId) returns()
func (_ERC2771World *ERC2771WorldSession) UpdateCorpId(characterId *big.Int, corpId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateCorpId(&_ERC2771World.TransactOpts, characterId, corpId)
}

// UpdateCorpId is a paid mutator transaction binding the contract method 0x325675d4.
//
// Solidity: function updateCorpId(uint256 characterId, uint256 corpId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UpdateCorpId(characterId *big.Int, corpId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateCorpId(&_ERC2771World.TransactOpts, characterId, corpId)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactor) UpdateFuel(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "updateFuel", entityId)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldSession) UpdateFuel(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateFuel(&_ERC2771World.TransactOpts, entityId)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 entityId) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) UpdateFuel(entityId *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.UpdateFuel(&_ERC2771World.TransactOpts, entityId)
}

// WithdrawFromEphemeralInventory is a paid mutator transaction binding the contract method 0x2e098e36.
//
// Solidity: function withdrawFromEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) WithdrawFromEphemeralInventory(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "withdrawFromEphemeralInventory", smartObjectId, ephemeralInventoryOwner, items)
}

// WithdrawFromEphemeralInventory is a paid mutator transaction binding the contract method 0x2e098e36.
//
// Solidity: function withdrawFromEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) WithdrawFromEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawFromEphemeralInventory(&_ERC2771World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// WithdrawFromEphemeralInventory is a paid mutator transaction binding the contract method 0x2e098e36.
//
// Solidity: function withdrawFromEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) WithdrawFromEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawFromEphemeralInventory(&_ERC2771World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// WithdrawFromInventory is a paid mutator transaction binding the contract method 0x8f7513ff.
//
// Solidity: function withdrawFromInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactor) WithdrawFromInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "withdrawFromInventory", smartObjectId, items)
}

// WithdrawFromInventory is a paid mutator transaction binding the contract method 0x8f7513ff.
//
// Solidity: function withdrawFromInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldSession) WithdrawFromInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawFromInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// WithdrawFromInventory is a paid mutator transaction binding the contract method 0x8f7513ff.
//
// Solidity: function withdrawFromInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) WithdrawFromInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawFromInventory(&_ERC2771World.TransactOpts, smartObjectId, items)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 entityId, uint256 unitAmount) returns()
func (_ERC2771World *ERC2771WorldTransactor) WithdrawFuel(opts *bind.TransactOpts, entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.contract.Transact(opts, "withdrawFuel", entityId, unitAmount)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 entityId, uint256 unitAmount) returns()
func (_ERC2771World *ERC2771WorldSession) WithdrawFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawFuel(&_ERC2771World.TransactOpts, entityId, unitAmount)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 entityId, uint256 unitAmount) returns()
func (_ERC2771World *ERC2771WorldTransactorSession) WithdrawFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _ERC2771World.Contract.WithdrawFuel(&_ERC2771World.TransactOpts, entityId, unitAmount)
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
	TableId        [32]byte
	KeyTuple       [][32]byte
	Start          *big.Int
	DeleteCount    *big.Int
	EncodedLengths [32]byte
	Data           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStoreSpliceDynamicData is a free log retrieval operation binding the contract event 0xaa63765a776145e5e6492f471ae097dfed11cd57a61bc2679dd43180422385b4.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
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

// WatchStoreSpliceDynamicData is a free log subscription operation binding the contract event 0xaa63765a776145e5e6492f471ae097dfed11cd57a61bc2679dd43180422385b4.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
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

// ParseStoreSpliceDynamicData is a log parse operation binding the contract event 0xaa63765a776145e5e6492f471ae097dfed11cd57a61bc2679dd43180422385b4.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
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
