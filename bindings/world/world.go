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

// WorldMetaData contains all meta data concerning the World contract.
var WorldMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"batchCall\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallData[]\",\"components\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchCallFrom\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallFromData[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"callFrom\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"creator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__aggression\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"turretOwnerCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"aggressor\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"victim\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__anchor\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationTableData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__bringOffline\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__bringOnline\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__canJump\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__configureInteractionHandler\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interactionParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__configureSmartGate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__configureSmartTurret\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__createAndAnchorSmartGate\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordData\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"worldPosition\",\"type\":\"tuple\",\"internalType\":\"structWorldPosition\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"position\",\"type\":\"tuple\",\"internalType\":\"structCoord\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDistance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__createAndAnchorSmartStorageUnit\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordData\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"worldPosition\",\"type\":\"tuple\",\"internalType\":\"structWorldPosition\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"position\",\"type\":\"tuple\",\"internalType\":\"structCoord\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralStorageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__createAndAnchorSmartTurret\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordData\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"worldPosition\",\"type\":\"tuple\",\"internalType\":\"structWorldPosition\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"position\",\"type\":\"tuple\",\"internalType\":\"structCoord\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__createAndDepositItemsToEphemeralInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__createAndDepositItemsToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__createCharacter\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"corpId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecord\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"entityRecordOffchain\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordOffchainTableData\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"tokenCid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__createEntityRecord\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__createEntityRecordOffchain\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__currentFuelAmount\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eveworld__currentFuelAmountInWei\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eveworld__depositFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__depositToEphemeralInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__depositToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__destroyDeployable\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__ephemeralToInventoryTransfer\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__globalPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__globalResume\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__inProximity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"turretOwnerCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turret\",\"type\":\"tuple\",\"internalType\":\"structTurret\",\"components\":[{\"name\":\"weaponTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ammoTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargesLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"turretTarget\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"updatedPriorityQueue\",\"type\":\"tuple[]\",\"internalType\":\"structTargetPriority[]\",\"components\":[{\"name\":\"target\",\"type\":\"tuple\",\"internalType\":\"structSmartTurretTarget\",\"components\":[{\"name\":\"shipId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shipTypeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hpRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shieldRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"armorRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__inventoryToEphemeralTransfer\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outItems\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__inventoryToEphemeralTransferWithParam\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outItems\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__isGateLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eveworld__isWithinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eveworld__linkSmartGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__registerDeployable\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"fuelUnitVolumeInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacityInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__registerDeployableToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__registerERC721Token\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__reportKill\",\"inputs\":[{\"name\":\"killMailId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killMailTableData\",\"type\":\"tuple\",\"internalType\":\"structKillMailTableData\",\"components\":[{\"name\":\"killerCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"victimCharacterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lossType\",\"type\":\"uint8\",\"internalType\":\"enumKillMailLossType\"},{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"killTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__saveLocation\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"location\",\"type\":\"tuple\",\"internalType\":\"structLocationTableData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setAccessEnforcement\",\"inputs\":[{\"name\":\"target\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isEnforced\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setAccessListByRole\",\"inputs\":[{\"name\":\"accessRoleId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accessList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setBaseURI\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"baseURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setCharClassId\",\"inputs\":[{\"name\":\"classId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setCid\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setDappURL\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setDeployableMetadata\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setDescription\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setEntityMetadata\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setEphemeralInventoryCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralStorageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setFuelConsumptionPerMinute\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionIntervalInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setFuelMaxCapacity\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacityInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setInventoryCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setMetadata\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structStaticDataGlobalTableData\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"baseURI\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setName\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setName\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setSSUClassId\",\"inputs\":[{\"name\":\"classId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setSmartAssemblyType\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"smartAssemblyType\",\"type\":\"uint8\",\"internalType\":\"enumSmartAssemblyType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__setSymbol\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__unanchor\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__unlinkSmartGates\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__updateCorpId\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"corpId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__updateFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__withdrawFromEphemeralInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__withdrawFromInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eveworld__withdrawFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldSlice\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLayout\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeySchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValueSchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialMsgSender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initModule\",\"type\":\"address\",\"internalType\":\"contractIModule\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installRootModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"popFromDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"byteLengthToPop\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pushToDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dataToPush\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"systemFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespace\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerRootFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"worldFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"systemFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystem\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"contractSystem\"},{\"name\":\"publicAccess\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerTable\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"},{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"keyNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"fieldNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storeVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferBalanceToAddress\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferBalanceToNamespace\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"worldVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"HelloStore\",\"inputs\":[{\"name\":\"storeVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HelloWorld\",\"inputs\":[{\"name\":\"worldVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_DeleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SetRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"PackedCounter\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"SmartGate_GateAlreadyLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_GateNotLinked\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_NotWithtinRange\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_SameSourceAndDestination\",\"inputs\":[{\"name\":\"sourceGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationGateId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartGate_UndefinedClassId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartStorageUnitERC721AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartTurret_NotConfigured\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartTurret_UndefinedClassId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Store_IndexOutOfBounds\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessedIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidBounds\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidFieldNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidKeyNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidSplice\",\"inputs\":[{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"fieldLength\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaDynamicLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaStaticLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_TableAlreadyExists\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_TableNotFound\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_AccessDenied\",\"inputs\":[{\"name\":\"resource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"World_CallbackNotAllowed\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_DelegationNotFound\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_DirectCallToSystemForbidden\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorAlreadyExists\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorNotFound\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"World_InterfaceNotSupported\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InvalidNamespace\",\"inputs\":[{\"name\":\"namespace\",\"type\":\"bytes14\",\"internalType\":\"bytes14\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceId\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceAlreadyExists\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceNotFound\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_SystemAlreadyExists\",\"inputs\":[{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_UnlimitedDelegationNotAllowed\",\"inputs\":[]}]",
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

// EveworldCurrentFuelAmount is a free data retrieval call binding the contract method 0xfe055b2f.
//
// Solidity: function eveworld__currentFuelAmount(uint256 entityId) view returns(uint256 amount)
func (_World *WorldCaller) EveworldCurrentFuelAmount(opts *bind.CallOpts, entityId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "eveworld__currentFuelAmount", entityId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EveworldCurrentFuelAmount is a free data retrieval call binding the contract method 0xfe055b2f.
//
// Solidity: function eveworld__currentFuelAmount(uint256 entityId) view returns(uint256 amount)
func (_World *WorldSession) EveworldCurrentFuelAmount(entityId *big.Int) (*big.Int, error) {
	return _World.Contract.EveworldCurrentFuelAmount(&_World.CallOpts, entityId)
}

// EveworldCurrentFuelAmount is a free data retrieval call binding the contract method 0xfe055b2f.
//
// Solidity: function eveworld__currentFuelAmount(uint256 entityId) view returns(uint256 amount)
func (_World *WorldCallerSession) EveworldCurrentFuelAmount(entityId *big.Int) (*big.Int, error) {
	return _World.Contract.EveworldCurrentFuelAmount(&_World.CallOpts, entityId)
}

// EveworldCurrentFuelAmountInWei is a free data retrieval call binding the contract method 0x9335ecac.
//
// Solidity: function eveworld__currentFuelAmountInWei(uint256 entityId) view returns(uint256 amount)
func (_World *WorldCaller) EveworldCurrentFuelAmountInWei(opts *bind.CallOpts, entityId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "eveworld__currentFuelAmountInWei", entityId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EveworldCurrentFuelAmountInWei is a free data retrieval call binding the contract method 0x9335ecac.
//
// Solidity: function eveworld__currentFuelAmountInWei(uint256 entityId) view returns(uint256 amount)
func (_World *WorldSession) EveworldCurrentFuelAmountInWei(entityId *big.Int) (*big.Int, error) {
	return _World.Contract.EveworldCurrentFuelAmountInWei(&_World.CallOpts, entityId)
}

// EveworldCurrentFuelAmountInWei is a free data retrieval call binding the contract method 0x9335ecac.
//
// Solidity: function eveworld__currentFuelAmountInWei(uint256 entityId) view returns(uint256 amount)
func (_World *WorldCallerSession) EveworldCurrentFuelAmountInWei(entityId *big.Int) (*big.Int, error) {
	return _World.Contract.EveworldCurrentFuelAmountInWei(&_World.CallOpts, entityId)
}

// EveworldIsGateLinked is a free data retrieval call binding the contract method 0xc3dc12f2.
//
// Solidity: function eveworld__isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCaller) EveworldIsGateLinked(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "eveworld__isGateLinked", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EveworldIsGateLinked is a free data retrieval call binding the contract method 0xc3dc12f2.
//
// Solidity: function eveworld__isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldSession) EveworldIsGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EveworldIsGateLinked(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EveworldIsGateLinked is a free data retrieval call binding the contract method 0xc3dc12f2.
//
// Solidity: function eveworld__isGateLinked(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCallerSession) EveworldIsGateLinked(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EveworldIsGateLinked(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EveworldIsWithinRange is a free data retrieval call binding the contract method 0xe16fff52.
//
// Solidity: function eveworld__isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCaller) EveworldIsWithinRange(opts *bind.CallOpts, sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "eveworld__isWithinRange", sourceGateId, destinationGateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EveworldIsWithinRange is a free data retrieval call binding the contract method 0xe16fff52.
//
// Solidity: function eveworld__isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldSession) EveworldIsWithinRange(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EveworldIsWithinRange(&_World.CallOpts, sourceGateId, destinationGateId)
}

// EveworldIsWithinRange is a free data retrieval call binding the contract method 0xe16fff52.
//
// Solidity: function eveworld__isWithinRange(uint256 sourceGateId, uint256 destinationGateId) view returns(bool)
func (_World *WorldCallerSession) EveworldIsWithinRange(sourceGateId *big.Int, destinationGateId *big.Int) (bool, error) {
	return _World.Contract.EveworldIsWithinRange(&_World.CallOpts, sourceGateId, destinationGateId)
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

// InitialMsgSender is a free data retrieval call binding the contract method 0x1e71d54c.
//
// Solidity: function initialMsgSender() view returns(address)
func (_World *WorldCaller) InitialMsgSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _World.contract.Call(opts, &out, "initialMsgSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InitialMsgSender is a free data retrieval call binding the contract method 0x1e71d54c.
//
// Solidity: function initialMsgSender() view returns(address)
func (_World *WorldSession) InitialMsgSender() (common.Address, error) {
	return _World.Contract.InitialMsgSender(&_World.CallOpts)
}

// InitialMsgSender is a free data retrieval call binding the contract method 0x1e71d54c.
//
// Solidity: function initialMsgSender() view returns(address)
func (_World *WorldCallerSession) InitialMsgSender() (common.Address, error) {
	return _World.Contract.InitialMsgSender(&_World.CallOpts)
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

// EveworldAggression is a paid mutator transaction binding the contract method 0xcbabbe17.
//
// Solidity: function eveworld__aggression(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) aggressor, (uint256,uint256,uint256,uint256,uint256,uint256) victim) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldTransactor) EveworldAggression(opts *bind.TransactOpts, smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, aggressor SmartTurretTarget, victim SmartTurretTarget) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__aggression", smartObjectId, turretOwnerCharacterId, priorityQueue, turret, aggressor, victim)
}

// EveworldAggression is a paid mutator transaction binding the contract method 0xcbabbe17.
//
// Solidity: function eveworld__aggression(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) aggressor, (uint256,uint256,uint256,uint256,uint256,uint256) victim) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldSession) EveworldAggression(smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, aggressor SmartTurretTarget, victim SmartTurretTarget) (*types.Transaction, error) {
	return _World.Contract.EveworldAggression(&_World.TransactOpts, smartObjectId, turretOwnerCharacterId, priorityQueue, turret, aggressor, victim)
}

// EveworldAggression is a paid mutator transaction binding the contract method 0xcbabbe17.
//
// Solidity: function eveworld__aggression(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) aggressor, (uint256,uint256,uint256,uint256,uint256,uint256) victim) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldTransactorSession) EveworldAggression(smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, aggressor SmartTurretTarget, victim SmartTurretTarget) (*types.Transaction, error) {
	return _World.Contract.EveworldAggression(&_World.TransactOpts, smartObjectId, turretOwnerCharacterId, priorityQueue, turret, aggressor, victim)
}

// EveworldAnchor is a paid mutator transaction binding the contract method 0x708c3bf5.
//
// Solidity: function eveworld__anchor(uint256 entityId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_World *WorldTransactor) EveworldAnchor(opts *bind.TransactOpts, entityId *big.Int, locationData LocationTableData) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__anchor", entityId, locationData)
}

// EveworldAnchor is a paid mutator transaction binding the contract method 0x708c3bf5.
//
// Solidity: function eveworld__anchor(uint256 entityId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_World *WorldSession) EveworldAnchor(entityId *big.Int, locationData LocationTableData) (*types.Transaction, error) {
	return _World.Contract.EveworldAnchor(&_World.TransactOpts, entityId, locationData)
}

// EveworldAnchor is a paid mutator transaction binding the contract method 0x708c3bf5.
//
// Solidity: function eveworld__anchor(uint256 entityId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_World *WorldTransactorSession) EveworldAnchor(entityId *big.Int, locationData LocationTableData) (*types.Transaction, error) {
	return _World.Contract.EveworldAnchor(&_World.TransactOpts, entityId, locationData)
}

// EveworldBringOffline is a paid mutator transaction binding the contract method 0x9d147e19.
//
// Solidity: function eveworld__bringOffline(uint256 entityId) returns()
func (_World *WorldTransactor) EveworldBringOffline(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__bringOffline", entityId)
}

// EveworldBringOffline is a paid mutator transaction binding the contract method 0x9d147e19.
//
// Solidity: function eveworld__bringOffline(uint256 entityId) returns()
func (_World *WorldSession) EveworldBringOffline(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldBringOffline(&_World.TransactOpts, entityId)
}

// EveworldBringOffline is a paid mutator transaction binding the contract method 0x9d147e19.
//
// Solidity: function eveworld__bringOffline(uint256 entityId) returns()
func (_World *WorldTransactorSession) EveworldBringOffline(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldBringOffline(&_World.TransactOpts, entityId)
}

// EveworldBringOnline is a paid mutator transaction binding the contract method 0x2a4606fb.
//
// Solidity: function eveworld__bringOnline(uint256 entityId) returns()
func (_World *WorldTransactor) EveworldBringOnline(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__bringOnline", entityId)
}

// EveworldBringOnline is a paid mutator transaction binding the contract method 0x2a4606fb.
//
// Solidity: function eveworld__bringOnline(uint256 entityId) returns()
func (_World *WorldSession) EveworldBringOnline(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldBringOnline(&_World.TransactOpts, entityId)
}

// EveworldBringOnline is a paid mutator transaction binding the contract method 0x2a4606fb.
//
// Solidity: function eveworld__bringOnline(uint256 entityId) returns()
func (_World *WorldTransactorSession) EveworldBringOnline(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldBringOnline(&_World.TransactOpts, entityId)
}

// EveworldCanJump is a paid mutator transaction binding the contract method 0xacb54c49.
//
// Solidity: function eveworld__canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_World *WorldTransactor) EveworldCanJump(opts *bind.TransactOpts, characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__canJump", characterId, sourceGateId, destinationGateId)
}

// EveworldCanJump is a paid mutator transaction binding the contract method 0xacb54c49.
//
// Solidity: function eveworld__canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_World *WorldSession) EveworldCanJump(characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCanJump(&_World.TransactOpts, characterId, sourceGateId, destinationGateId)
}

// EveworldCanJump is a paid mutator transaction binding the contract method 0xacb54c49.
//
// Solidity: function eveworld__canJump(uint256 characterId, uint256 sourceGateId, uint256 destinationGateId) returns(bool)
func (_World *WorldTransactorSession) EveworldCanJump(characterId *big.Int, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCanJump(&_World.TransactOpts, characterId, sourceGateId, destinationGateId)
}

// EveworldConfigureInteractionHandler is a paid mutator transaction binding the contract method 0x443782f5.
//
// Solidity: function eveworld__configureInteractionHandler(uint256 smartObjectId, bytes interactionParams) returns()
func (_World *WorldTransactor) EveworldConfigureInteractionHandler(opts *bind.TransactOpts, smartObjectId *big.Int, interactionParams []byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__configureInteractionHandler", smartObjectId, interactionParams)
}

// EveworldConfigureInteractionHandler is a paid mutator transaction binding the contract method 0x443782f5.
//
// Solidity: function eveworld__configureInteractionHandler(uint256 smartObjectId, bytes interactionParams) returns()
func (_World *WorldSession) EveworldConfigureInteractionHandler(smartObjectId *big.Int, interactionParams []byte) (*types.Transaction, error) {
	return _World.Contract.EveworldConfigureInteractionHandler(&_World.TransactOpts, smartObjectId, interactionParams)
}

// EveworldConfigureInteractionHandler is a paid mutator transaction binding the contract method 0x443782f5.
//
// Solidity: function eveworld__configureInteractionHandler(uint256 smartObjectId, bytes interactionParams) returns()
func (_World *WorldTransactorSession) EveworldConfigureInteractionHandler(smartObjectId *big.Int, interactionParams []byte) (*types.Transaction, error) {
	return _World.Contract.EveworldConfigureInteractionHandler(&_World.TransactOpts, smartObjectId, interactionParams)
}

// EveworldConfigureSmartGate is a paid mutator transaction binding the contract method 0xefb17cbc.
//
// Solidity: function eveworld__configureSmartGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldTransactor) EveworldConfigureSmartGate(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__configureSmartGate", smartObjectId, systemId)
}

// EveworldConfigureSmartGate is a paid mutator transaction binding the contract method 0xefb17cbc.
//
// Solidity: function eveworld__configureSmartGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldSession) EveworldConfigureSmartGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.Contract.EveworldConfigureSmartGate(&_World.TransactOpts, smartObjectId, systemId)
}

// EveworldConfigureSmartGate is a paid mutator transaction binding the contract method 0xefb17cbc.
//
// Solidity: function eveworld__configureSmartGate(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldTransactorSession) EveworldConfigureSmartGate(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.Contract.EveworldConfigureSmartGate(&_World.TransactOpts, smartObjectId, systemId)
}

// EveworldConfigureSmartTurret is a paid mutator transaction binding the contract method 0xc0b938af.
//
// Solidity: function eveworld__configureSmartTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldTransactor) EveworldConfigureSmartTurret(opts *bind.TransactOpts, smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__configureSmartTurret", smartObjectId, systemId)
}

// EveworldConfigureSmartTurret is a paid mutator transaction binding the contract method 0xc0b938af.
//
// Solidity: function eveworld__configureSmartTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldSession) EveworldConfigureSmartTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.Contract.EveworldConfigureSmartTurret(&_World.TransactOpts, smartObjectId, systemId)
}

// EveworldConfigureSmartTurret is a paid mutator transaction binding the contract method 0xc0b938af.
//
// Solidity: function eveworld__configureSmartTurret(uint256 smartObjectId, bytes32 systemId) returns()
func (_World *WorldTransactorSession) EveworldConfigureSmartTurret(smartObjectId *big.Int, systemId [32]byte) (*types.Transaction, error) {
	return _World.Contract.EveworldConfigureSmartTurret(&_World.TransactOpts, smartObjectId, systemId)
}

// EveworldCreateAndAnchorSmartGate is a paid mutator transaction binding the contract method 0x22414d41.
//
// Solidity: function eveworld__createAndAnchorSmartGate(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 maxDistance) returns()
func (_World *WorldTransactor) EveworldCreateAndAnchorSmartGate(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, maxDistance *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__createAndAnchorSmartGate", smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, maxDistance)
}

// EveworldCreateAndAnchorSmartGate is a paid mutator transaction binding the contract method 0x22414d41.
//
// Solidity: function eveworld__createAndAnchorSmartGate(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 maxDistance) returns()
func (_World *WorldSession) EveworldCreateAndAnchorSmartGate(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, maxDistance *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndAnchorSmartGate(&_World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, maxDistance)
}

// EveworldCreateAndAnchorSmartGate is a paid mutator transaction binding the contract method 0x22414d41.
//
// Solidity: function eveworld__createAndAnchorSmartGate(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 maxDistance) returns()
func (_World *WorldTransactorSession) EveworldCreateAndAnchorSmartGate(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, maxDistance *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndAnchorSmartGate(&_World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, maxDistance)
}

// EveworldCreateAndAnchorSmartStorageUnit is a paid mutator transaction binding the contract method 0xa1353e41.
//
// Solidity: function eveworld__createAndAnchorSmartStorageUnit(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 storageCapacity, uint256 ephemeralStorageCapacity) returns()
func (_World *WorldTransactor) EveworldCreateAndAnchorSmartStorageUnit(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, storageCapacity *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__createAndAnchorSmartStorageUnit", smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, storageCapacity, ephemeralStorageCapacity)
}

// EveworldCreateAndAnchorSmartStorageUnit is a paid mutator transaction binding the contract method 0xa1353e41.
//
// Solidity: function eveworld__createAndAnchorSmartStorageUnit(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 storageCapacity, uint256 ephemeralStorageCapacity) returns()
func (_World *WorldSession) EveworldCreateAndAnchorSmartStorageUnit(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, storageCapacity *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndAnchorSmartStorageUnit(&_World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, storageCapacity, ephemeralStorageCapacity)
}

// EveworldCreateAndAnchorSmartStorageUnit is a paid mutator transaction binding the contract method 0xa1353e41.
//
// Solidity: function eveworld__createAndAnchorSmartStorageUnit(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity, uint256 storageCapacity, uint256 ephemeralStorageCapacity) returns()
func (_World *WorldTransactorSession) EveworldCreateAndAnchorSmartStorageUnit(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int, storageCapacity *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndAnchorSmartStorageUnit(&_World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity, storageCapacity, ephemeralStorageCapacity)
}

// EveworldCreateAndAnchorSmartTurret is a paid mutator transaction binding the contract method 0xc22f9932.
//
// Solidity: function eveworld__createAndAnchorSmartTurret(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity) returns()
func (_World *WorldTransactor) EveworldCreateAndAnchorSmartTurret(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__createAndAnchorSmartTurret", smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity)
}

// EveworldCreateAndAnchorSmartTurret is a paid mutator transaction binding the contract method 0xc22f9932.
//
// Solidity: function eveworld__createAndAnchorSmartTurret(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity) returns()
func (_World *WorldSession) EveworldCreateAndAnchorSmartTurret(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndAnchorSmartTurret(&_World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity)
}

// EveworldCreateAndAnchorSmartTurret is a paid mutator transaction binding the contract method 0xc22f9932.
//
// Solidity: function eveworld__createAndAnchorSmartTurret(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacity) returns()
func (_World *WorldTransactorSession) EveworldCreateAndAnchorSmartTurret(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndAnchorSmartTurret(&_World.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionIntervalInSeconds, fuelMaxCapacity)
}

// EveworldCreateAndDepositItemsToEphemeralInventory is a paid mutator transaction binding the contract method 0x88633ae4.
//
// Solidity: function eveworld__createAndDepositItemsToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EveworldCreateAndDepositItemsToEphemeralInventory(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__createAndDepositItemsToEphemeralInventory", smartObjectId, ephemeralInventoryOwner, items)
}

// EveworldCreateAndDepositItemsToEphemeralInventory is a paid mutator transaction binding the contract method 0x88633ae4.
//
// Solidity: function eveworld__createAndDepositItemsToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldSession) EveworldCreateAndDepositItemsToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndDepositItemsToEphemeralInventory(&_World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// EveworldCreateAndDepositItemsToEphemeralInventory is a paid mutator transaction binding the contract method 0x88633ae4.
//
// Solidity: function eveworld__createAndDepositItemsToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EveworldCreateAndDepositItemsToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndDepositItemsToEphemeralInventory(&_World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// EveworldCreateAndDepositItemsToInventory is a paid mutator transaction binding the contract method 0xd6f49960.
//
// Solidity: function eveworld__createAndDepositItemsToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EveworldCreateAndDepositItemsToInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__createAndDepositItemsToInventory", smartObjectId, items)
}

// EveworldCreateAndDepositItemsToInventory is a paid mutator transaction binding the contract method 0xd6f49960.
//
// Solidity: function eveworld__createAndDepositItemsToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldSession) EveworldCreateAndDepositItemsToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndDepositItemsToInventory(&_World.TransactOpts, smartObjectId, items)
}

// EveworldCreateAndDepositItemsToInventory is a paid mutator transaction binding the contract method 0xd6f49960.
//
// Solidity: function eveworld__createAndDepositItemsToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EveworldCreateAndDepositItemsToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateAndDepositItemsToInventory(&_World.TransactOpts, smartObjectId, items)
}

// EveworldCreateCharacter is a paid mutator transaction binding the contract method 0x552a848c.
//
// Solidity: function eveworld__createCharacter(uint256 characterId, address characterAddress, uint256 corpId, (uint256,uint256,uint256) entityRecord, (string,string,string) entityRecordOffchain, string tokenCid) returns()
func (_World *WorldTransactor) EveworldCreateCharacter(opts *bind.TransactOpts, characterId *big.Int, characterAddress common.Address, corpId *big.Int, entityRecord EntityRecordData, entityRecordOffchain EntityRecordOffchainTableData, tokenCid string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__createCharacter", characterId, characterAddress, corpId, entityRecord, entityRecordOffchain, tokenCid)
}

// EveworldCreateCharacter is a paid mutator transaction binding the contract method 0x552a848c.
//
// Solidity: function eveworld__createCharacter(uint256 characterId, address characterAddress, uint256 corpId, (uint256,uint256,uint256) entityRecord, (string,string,string) entityRecordOffchain, string tokenCid) returns()
func (_World *WorldSession) EveworldCreateCharacter(characterId *big.Int, characterAddress common.Address, corpId *big.Int, entityRecord EntityRecordData, entityRecordOffchain EntityRecordOffchainTableData, tokenCid string) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateCharacter(&_World.TransactOpts, characterId, characterAddress, corpId, entityRecord, entityRecordOffchain, tokenCid)
}

// EveworldCreateCharacter is a paid mutator transaction binding the contract method 0x552a848c.
//
// Solidity: function eveworld__createCharacter(uint256 characterId, address characterAddress, uint256 corpId, (uint256,uint256,uint256) entityRecord, (string,string,string) entityRecordOffchain, string tokenCid) returns()
func (_World *WorldTransactorSession) EveworldCreateCharacter(characterId *big.Int, characterAddress common.Address, corpId *big.Int, entityRecord EntityRecordData, entityRecordOffchain EntityRecordOffchainTableData, tokenCid string) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateCharacter(&_World.TransactOpts, characterId, characterAddress, corpId, entityRecord, entityRecordOffchain, tokenCid)
}

// EveworldCreateEntityRecord is a paid mutator transaction binding the contract method 0x3d3bf32e.
//
// Solidity: function eveworld__createEntityRecord(uint256 entityId, uint256 itemId, uint256 typeId, uint256 volume) returns()
func (_World *WorldTransactor) EveworldCreateEntityRecord(opts *bind.TransactOpts, entityId *big.Int, itemId *big.Int, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__createEntityRecord", entityId, itemId, typeId, volume)
}

// EveworldCreateEntityRecord is a paid mutator transaction binding the contract method 0x3d3bf32e.
//
// Solidity: function eveworld__createEntityRecord(uint256 entityId, uint256 itemId, uint256 typeId, uint256 volume) returns()
func (_World *WorldSession) EveworldCreateEntityRecord(entityId *big.Int, itemId *big.Int, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateEntityRecord(&_World.TransactOpts, entityId, itemId, typeId, volume)
}

// EveworldCreateEntityRecord is a paid mutator transaction binding the contract method 0x3d3bf32e.
//
// Solidity: function eveworld__createEntityRecord(uint256 entityId, uint256 itemId, uint256 typeId, uint256 volume) returns()
func (_World *WorldTransactorSession) EveworldCreateEntityRecord(entityId *big.Int, itemId *big.Int, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateEntityRecord(&_World.TransactOpts, entityId, itemId, typeId, volume)
}

// EveworldCreateEntityRecordOffchain is a paid mutator transaction binding the contract method 0x31fa27a3.
//
// Solidity: function eveworld__createEntityRecordOffchain(uint256 entityId, string name, string dappURL, string description) returns()
func (_World *WorldTransactor) EveworldCreateEntityRecordOffchain(opts *bind.TransactOpts, entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__createEntityRecordOffchain", entityId, name, dappURL, description)
}

// EveworldCreateEntityRecordOffchain is a paid mutator transaction binding the contract method 0x31fa27a3.
//
// Solidity: function eveworld__createEntityRecordOffchain(uint256 entityId, string name, string dappURL, string description) returns()
func (_World *WorldSession) EveworldCreateEntityRecordOffchain(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateEntityRecordOffchain(&_World.TransactOpts, entityId, name, dappURL, description)
}

// EveworldCreateEntityRecordOffchain is a paid mutator transaction binding the contract method 0x31fa27a3.
//
// Solidity: function eveworld__createEntityRecordOffchain(uint256 entityId, string name, string dappURL, string description) returns()
func (_World *WorldTransactorSession) EveworldCreateEntityRecordOffchain(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _World.Contract.EveworldCreateEntityRecordOffchain(&_World.TransactOpts, entityId, name, dappURL, description)
}

// EveworldDepositFuel is a paid mutator transaction binding the contract method 0x23cbd201.
//
// Solidity: function eveworld__depositFuel(uint256 entityId, uint256 unitAmount) returns()
func (_World *WorldTransactor) EveworldDepositFuel(opts *bind.TransactOpts, entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__depositFuel", entityId, unitAmount)
}

// EveworldDepositFuel is a paid mutator transaction binding the contract method 0x23cbd201.
//
// Solidity: function eveworld__depositFuel(uint256 entityId, uint256 unitAmount) returns()
func (_World *WorldSession) EveworldDepositFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldDepositFuel(&_World.TransactOpts, entityId, unitAmount)
}

// EveworldDepositFuel is a paid mutator transaction binding the contract method 0x23cbd201.
//
// Solidity: function eveworld__depositFuel(uint256 entityId, uint256 unitAmount) returns()
func (_World *WorldTransactorSession) EveworldDepositFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldDepositFuel(&_World.TransactOpts, entityId, unitAmount)
}

// EveworldDepositToEphemeralInventory is a paid mutator transaction binding the contract method 0x90d1f28e.
//
// Solidity: function eveworld__depositToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EveworldDepositToEphemeralInventory(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__depositToEphemeralInventory", smartObjectId, ephemeralInventoryOwner, items)
}

// EveworldDepositToEphemeralInventory is a paid mutator transaction binding the contract method 0x90d1f28e.
//
// Solidity: function eveworld__depositToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldSession) EveworldDepositToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldDepositToEphemeralInventory(&_World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// EveworldDepositToEphemeralInventory is a paid mutator transaction binding the contract method 0x90d1f28e.
//
// Solidity: function eveworld__depositToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EveworldDepositToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldDepositToEphemeralInventory(&_World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// EveworldDepositToInventory is a paid mutator transaction binding the contract method 0x37b52bd6.
//
// Solidity: function eveworld__depositToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EveworldDepositToInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__depositToInventory", smartObjectId, items)
}

// EveworldDepositToInventory is a paid mutator transaction binding the contract method 0x37b52bd6.
//
// Solidity: function eveworld__depositToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldSession) EveworldDepositToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldDepositToInventory(&_World.TransactOpts, smartObjectId, items)
}

// EveworldDepositToInventory is a paid mutator transaction binding the contract method 0x37b52bd6.
//
// Solidity: function eveworld__depositToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EveworldDepositToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldDepositToInventory(&_World.TransactOpts, smartObjectId, items)
}

// EveworldDestroyDeployable is a paid mutator transaction binding the contract method 0xad47cec7.
//
// Solidity: function eveworld__destroyDeployable(uint256 entityId) returns()
func (_World *WorldTransactor) EveworldDestroyDeployable(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__destroyDeployable", entityId)
}

// EveworldDestroyDeployable is a paid mutator transaction binding the contract method 0xad47cec7.
//
// Solidity: function eveworld__destroyDeployable(uint256 entityId) returns()
func (_World *WorldSession) EveworldDestroyDeployable(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldDestroyDeployable(&_World.TransactOpts, entityId)
}

// EveworldDestroyDeployable is a paid mutator transaction binding the contract method 0xad47cec7.
//
// Solidity: function eveworld__destroyDeployable(uint256 entityId) returns()
func (_World *WorldTransactorSession) EveworldDestroyDeployable(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldDestroyDeployable(&_World.TransactOpts, entityId)
}

// EveworldEphemeralToInventoryTransfer is a paid mutator transaction binding the contract method 0xd8e15827.
//
// Solidity: function eveworld__ephemeralToInventoryTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EveworldEphemeralToInventoryTransfer(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__ephemeralToInventoryTransfer", smartObjectId, items)
}

// EveworldEphemeralToInventoryTransfer is a paid mutator transaction binding the contract method 0xd8e15827.
//
// Solidity: function eveworld__ephemeralToInventoryTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldSession) EveworldEphemeralToInventoryTransfer(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldEphemeralToInventoryTransfer(&_World.TransactOpts, smartObjectId, items)
}

// EveworldEphemeralToInventoryTransfer is a paid mutator transaction binding the contract method 0xd8e15827.
//
// Solidity: function eveworld__ephemeralToInventoryTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EveworldEphemeralToInventoryTransfer(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldEphemeralToInventoryTransfer(&_World.TransactOpts, smartObjectId, items)
}

// EveworldGlobalPause is a paid mutator transaction binding the contract method 0xf023982a.
//
// Solidity: function eveworld__globalPause() returns()
func (_World *WorldTransactor) EveworldGlobalPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__globalPause")
}

// EveworldGlobalPause is a paid mutator transaction binding the contract method 0xf023982a.
//
// Solidity: function eveworld__globalPause() returns()
func (_World *WorldSession) EveworldGlobalPause() (*types.Transaction, error) {
	return _World.Contract.EveworldGlobalPause(&_World.TransactOpts)
}

// EveworldGlobalPause is a paid mutator transaction binding the contract method 0xf023982a.
//
// Solidity: function eveworld__globalPause() returns()
func (_World *WorldTransactorSession) EveworldGlobalPause() (*types.Transaction, error) {
	return _World.Contract.EveworldGlobalPause(&_World.TransactOpts)
}

// EveworldGlobalResume is a paid mutator transaction binding the contract method 0x03a630f1.
//
// Solidity: function eveworld__globalResume() returns()
func (_World *WorldTransactor) EveworldGlobalResume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__globalResume")
}

// EveworldGlobalResume is a paid mutator transaction binding the contract method 0x03a630f1.
//
// Solidity: function eveworld__globalResume() returns()
func (_World *WorldSession) EveworldGlobalResume() (*types.Transaction, error) {
	return _World.Contract.EveworldGlobalResume(&_World.TransactOpts)
}

// EveworldGlobalResume is a paid mutator transaction binding the contract method 0x03a630f1.
//
// Solidity: function eveworld__globalResume() returns()
func (_World *WorldTransactorSession) EveworldGlobalResume() (*types.Transaction, error) {
	return _World.Contract.EveworldGlobalResume(&_World.TransactOpts)
}

// EveworldInProximity is a paid mutator transaction binding the contract method 0x238d7ca6.
//
// Solidity: function eveworld__inProximity(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldTransactor) EveworldInProximity(opts *bind.TransactOpts, smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__inProximity", smartObjectId, turretOwnerCharacterId, priorityQueue, turret, turretTarget)
}

// EveworldInProximity is a paid mutator transaction binding the contract method 0x238d7ca6.
//
// Solidity: function eveworld__inProximity(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldSession) EveworldInProximity(smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _World.Contract.EveworldInProximity(&_World.TransactOpts, smartObjectId, turretOwnerCharacterId, priorityQueue, turret, turretTarget)
}

// EveworldInProximity is a paid mutator transaction binding the contract method 0x238d7ca6.
//
// Solidity: function eveworld__inProximity(uint256 smartObjectId, uint256 turretOwnerCharacterId, ((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] priorityQueue, (uint256,uint256,uint256) turret, (uint256,uint256,uint256,uint256,uint256,uint256) turretTarget) returns(((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[] updatedPriorityQueue)
func (_World *WorldTransactorSession) EveworldInProximity(smartObjectId *big.Int, turretOwnerCharacterId *big.Int, priorityQueue []TargetPriority, turret Turret, turretTarget SmartTurretTarget) (*types.Transaction, error) {
	return _World.Contract.EveworldInProximity(&_World.TransactOpts, smartObjectId, turretOwnerCharacterId, priorityQueue, turret, turretTarget)
}

// EveworldInventoryToEphemeralTransfer is a paid mutator transaction binding the contract method 0x8f5d3942.
//
// Solidity: function eveworld__inventoryToEphemeralTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_World *WorldTransactor) EveworldInventoryToEphemeralTransfer(opts *bind.TransactOpts, smartObjectId *big.Int, outItems []InventoryItem) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__inventoryToEphemeralTransfer", smartObjectId, outItems)
}

// EveworldInventoryToEphemeralTransfer is a paid mutator transaction binding the contract method 0x8f5d3942.
//
// Solidity: function eveworld__inventoryToEphemeralTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_World *WorldSession) EveworldInventoryToEphemeralTransfer(smartObjectId *big.Int, outItems []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldInventoryToEphemeralTransfer(&_World.TransactOpts, smartObjectId, outItems)
}

// EveworldInventoryToEphemeralTransfer is a paid mutator transaction binding the contract method 0x8f5d3942.
//
// Solidity: function eveworld__inventoryToEphemeralTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_World *WorldTransactorSession) EveworldInventoryToEphemeralTransfer(smartObjectId *big.Int, outItems []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldInventoryToEphemeralTransfer(&_World.TransactOpts, smartObjectId, outItems)
}

// EveworldInventoryToEphemeralTransferWithParam is a paid mutator transaction binding the contract method 0x96d6a57b.
//
// Solidity: function eveworld__inventoryToEphemeralTransferWithParam(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_World *WorldTransactor) EveworldInventoryToEphemeralTransferWithParam(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, outItems []InventoryItem) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__inventoryToEphemeralTransferWithParam", smartObjectId, ephemeralInventoryOwner, outItems)
}

// EveworldInventoryToEphemeralTransferWithParam is a paid mutator transaction binding the contract method 0x96d6a57b.
//
// Solidity: function eveworld__inventoryToEphemeralTransferWithParam(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_World *WorldSession) EveworldInventoryToEphemeralTransferWithParam(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, outItems []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldInventoryToEphemeralTransferWithParam(&_World.TransactOpts, smartObjectId, ephemeralInventoryOwner, outItems)
}

// EveworldInventoryToEphemeralTransferWithParam is a paid mutator transaction binding the contract method 0x96d6a57b.
//
// Solidity: function eveworld__inventoryToEphemeralTransferWithParam(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_World *WorldTransactorSession) EveworldInventoryToEphemeralTransferWithParam(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, outItems []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldInventoryToEphemeralTransferWithParam(&_World.TransactOpts, smartObjectId, ephemeralInventoryOwner, outItems)
}

// EveworldLinkSmartGates is a paid mutator transaction binding the contract method 0x5ada755f.
//
// Solidity: function eveworld__linkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldTransactor) EveworldLinkSmartGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__linkSmartGates", sourceGateId, destinationGateId)
}

// EveworldLinkSmartGates is a paid mutator transaction binding the contract method 0x5ada755f.
//
// Solidity: function eveworld__linkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldSession) EveworldLinkSmartGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldLinkSmartGates(&_World.TransactOpts, sourceGateId, destinationGateId)
}

// EveworldLinkSmartGates is a paid mutator transaction binding the contract method 0x5ada755f.
//
// Solidity: function eveworld__linkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldTransactorSession) EveworldLinkSmartGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldLinkSmartGates(&_World.TransactOpts, sourceGateId, destinationGateId)
}

// EveworldRegisterDeployable is a paid mutator transaction binding the contract method 0x2ebaa73a.
//
// Solidity: function eveworld__registerDeployable(uint256 entityId, (address,string) smartObjectData, uint256 fuelUnitVolumeInWei, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacityInWei) returns()
func (_World *WorldTransactor) EveworldRegisterDeployable(opts *bind.TransactOpts, entityId *big.Int, smartObjectData SmartObjectData, fuelUnitVolumeInWei *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacityInWei *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__registerDeployable", entityId, smartObjectData, fuelUnitVolumeInWei, fuelConsumptionIntervalInSeconds, fuelMaxCapacityInWei)
}

// EveworldRegisterDeployable is a paid mutator transaction binding the contract method 0x2ebaa73a.
//
// Solidity: function eveworld__registerDeployable(uint256 entityId, (address,string) smartObjectData, uint256 fuelUnitVolumeInWei, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacityInWei) returns()
func (_World *WorldSession) EveworldRegisterDeployable(entityId *big.Int, smartObjectData SmartObjectData, fuelUnitVolumeInWei *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacityInWei *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldRegisterDeployable(&_World.TransactOpts, entityId, smartObjectData, fuelUnitVolumeInWei, fuelConsumptionIntervalInSeconds, fuelMaxCapacityInWei)
}

// EveworldRegisterDeployable is a paid mutator transaction binding the contract method 0x2ebaa73a.
//
// Solidity: function eveworld__registerDeployable(uint256 entityId, (address,string) smartObjectData, uint256 fuelUnitVolumeInWei, uint256 fuelConsumptionIntervalInSeconds, uint256 fuelMaxCapacityInWei) returns()
func (_World *WorldTransactorSession) EveworldRegisterDeployable(entityId *big.Int, smartObjectData SmartObjectData, fuelUnitVolumeInWei *big.Int, fuelConsumptionIntervalInSeconds *big.Int, fuelMaxCapacityInWei *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldRegisterDeployable(&_World.TransactOpts, entityId, smartObjectData, fuelUnitVolumeInWei, fuelConsumptionIntervalInSeconds, fuelMaxCapacityInWei)
}

// EveworldRegisterDeployableToken is a paid mutator transaction binding the contract method 0xd9ba88ef.
//
// Solidity: function eveworld__registerDeployableToken(address tokenAddress) returns()
func (_World *WorldTransactor) EveworldRegisterDeployableToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__registerDeployableToken", tokenAddress)
}

// EveworldRegisterDeployableToken is a paid mutator transaction binding the contract method 0xd9ba88ef.
//
// Solidity: function eveworld__registerDeployableToken(address tokenAddress) returns()
func (_World *WorldSession) EveworldRegisterDeployableToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _World.Contract.EveworldRegisterDeployableToken(&_World.TransactOpts, tokenAddress)
}

// EveworldRegisterDeployableToken is a paid mutator transaction binding the contract method 0xd9ba88ef.
//
// Solidity: function eveworld__registerDeployableToken(address tokenAddress) returns()
func (_World *WorldTransactorSession) EveworldRegisterDeployableToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _World.Contract.EveworldRegisterDeployableToken(&_World.TransactOpts, tokenAddress)
}

// EveworldRegisterERC721Token is a paid mutator transaction binding the contract method 0xb94de5b5.
//
// Solidity: function eveworld__registerERC721Token(address tokenAddress) returns()
func (_World *WorldTransactor) EveworldRegisterERC721Token(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__registerERC721Token", tokenAddress)
}

// EveworldRegisterERC721Token is a paid mutator transaction binding the contract method 0xb94de5b5.
//
// Solidity: function eveworld__registerERC721Token(address tokenAddress) returns()
func (_World *WorldSession) EveworldRegisterERC721Token(tokenAddress common.Address) (*types.Transaction, error) {
	return _World.Contract.EveworldRegisterERC721Token(&_World.TransactOpts, tokenAddress)
}

// EveworldRegisterERC721Token is a paid mutator transaction binding the contract method 0xb94de5b5.
//
// Solidity: function eveworld__registerERC721Token(address tokenAddress) returns()
func (_World *WorldTransactorSession) EveworldRegisterERC721Token(tokenAddress common.Address) (*types.Transaction, error) {
	return _World.Contract.EveworldRegisterERC721Token(&_World.TransactOpts, tokenAddress)
}

// EveworldReportKill is a paid mutator transaction binding the contract method 0xe6efbb55.
//
// Solidity: function eveworld__reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailTableData) returns()
func (_World *WorldTransactor) EveworldReportKill(opts *bind.TransactOpts, killMailId *big.Int, killMailTableData KillMailTableData) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__reportKill", killMailId, killMailTableData)
}

// EveworldReportKill is a paid mutator transaction binding the contract method 0xe6efbb55.
//
// Solidity: function eveworld__reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailTableData) returns()
func (_World *WorldSession) EveworldReportKill(killMailId *big.Int, killMailTableData KillMailTableData) (*types.Transaction, error) {
	return _World.Contract.EveworldReportKill(&_World.TransactOpts, killMailId, killMailTableData)
}

// EveworldReportKill is a paid mutator transaction binding the contract method 0xe6efbb55.
//
// Solidity: function eveworld__reportKill(uint256 killMailId, (uint256,uint256,uint8,uint256,uint256) killMailTableData) returns()
func (_World *WorldTransactorSession) EveworldReportKill(killMailId *big.Int, killMailTableData KillMailTableData) (*types.Transaction, error) {
	return _World.Contract.EveworldReportKill(&_World.TransactOpts, killMailId, killMailTableData)
}

// EveworldSaveLocation is a paid mutator transaction binding the contract method 0x452934c5.
//
// Solidity: function eveworld__saveLocation(uint256 entityId, (uint256,uint256,uint256,uint256) location) returns()
func (_World *WorldTransactor) EveworldSaveLocation(opts *bind.TransactOpts, entityId *big.Int, location LocationTableData) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__saveLocation", entityId, location)
}

// EveworldSaveLocation is a paid mutator transaction binding the contract method 0x452934c5.
//
// Solidity: function eveworld__saveLocation(uint256 entityId, (uint256,uint256,uint256,uint256) location) returns()
func (_World *WorldSession) EveworldSaveLocation(entityId *big.Int, location LocationTableData) (*types.Transaction, error) {
	return _World.Contract.EveworldSaveLocation(&_World.TransactOpts, entityId, location)
}

// EveworldSaveLocation is a paid mutator transaction binding the contract method 0x452934c5.
//
// Solidity: function eveworld__saveLocation(uint256 entityId, (uint256,uint256,uint256,uint256) location) returns()
func (_World *WorldTransactorSession) EveworldSaveLocation(entityId *big.Int, location LocationTableData) (*types.Transaction, error) {
	return _World.Contract.EveworldSaveLocation(&_World.TransactOpts, entityId, location)
}

// EveworldSetAccessEnforcement is a paid mutator transaction binding the contract method 0xc108f384.
//
// Solidity: function eveworld__setAccessEnforcement(bytes32 target, bool isEnforced) returns()
func (_World *WorldTransactor) EveworldSetAccessEnforcement(opts *bind.TransactOpts, target [32]byte, isEnforced bool) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setAccessEnforcement", target, isEnforced)
}

// EveworldSetAccessEnforcement is a paid mutator transaction binding the contract method 0xc108f384.
//
// Solidity: function eveworld__setAccessEnforcement(bytes32 target, bool isEnforced) returns()
func (_World *WorldSession) EveworldSetAccessEnforcement(target [32]byte, isEnforced bool) (*types.Transaction, error) {
	return _World.Contract.EveworldSetAccessEnforcement(&_World.TransactOpts, target, isEnforced)
}

// EveworldSetAccessEnforcement is a paid mutator transaction binding the contract method 0xc108f384.
//
// Solidity: function eveworld__setAccessEnforcement(bytes32 target, bool isEnforced) returns()
func (_World *WorldTransactorSession) EveworldSetAccessEnforcement(target [32]byte, isEnforced bool) (*types.Transaction, error) {
	return _World.Contract.EveworldSetAccessEnforcement(&_World.TransactOpts, target, isEnforced)
}

// EveworldSetAccessListByRole is a paid mutator transaction binding the contract method 0x73acabe9.
//
// Solidity: function eveworld__setAccessListByRole(bytes32 accessRoleId, address[] accessList) returns()
func (_World *WorldTransactor) EveworldSetAccessListByRole(opts *bind.TransactOpts, accessRoleId [32]byte, accessList []common.Address) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setAccessListByRole", accessRoleId, accessList)
}

// EveworldSetAccessListByRole is a paid mutator transaction binding the contract method 0x73acabe9.
//
// Solidity: function eveworld__setAccessListByRole(bytes32 accessRoleId, address[] accessList) returns()
func (_World *WorldSession) EveworldSetAccessListByRole(accessRoleId [32]byte, accessList []common.Address) (*types.Transaction, error) {
	return _World.Contract.EveworldSetAccessListByRole(&_World.TransactOpts, accessRoleId, accessList)
}

// EveworldSetAccessListByRole is a paid mutator transaction binding the contract method 0x73acabe9.
//
// Solidity: function eveworld__setAccessListByRole(bytes32 accessRoleId, address[] accessList) returns()
func (_World *WorldTransactorSession) EveworldSetAccessListByRole(accessRoleId [32]byte, accessList []common.Address) (*types.Transaction, error) {
	return _World.Contract.EveworldSetAccessListByRole(&_World.TransactOpts, accessRoleId, accessList)
}

// EveworldSetBaseURI is a paid mutator transaction binding the contract method 0xdfb18ad6.
//
// Solidity: function eveworld__setBaseURI(bytes32 systemId, string baseURI) returns()
func (_World *WorldTransactor) EveworldSetBaseURI(opts *bind.TransactOpts, systemId [32]byte, baseURI string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setBaseURI", systemId, baseURI)
}

// EveworldSetBaseURI is a paid mutator transaction binding the contract method 0xdfb18ad6.
//
// Solidity: function eveworld__setBaseURI(bytes32 systemId, string baseURI) returns()
func (_World *WorldSession) EveworldSetBaseURI(systemId [32]byte, baseURI string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetBaseURI(&_World.TransactOpts, systemId, baseURI)
}

// EveworldSetBaseURI is a paid mutator transaction binding the contract method 0xdfb18ad6.
//
// Solidity: function eveworld__setBaseURI(bytes32 systemId, string baseURI) returns()
func (_World *WorldTransactorSession) EveworldSetBaseURI(systemId [32]byte, baseURI string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetBaseURI(&_World.TransactOpts, systemId, baseURI)
}

// EveworldSetCharClassId is a paid mutator transaction binding the contract method 0x217fd119.
//
// Solidity: function eveworld__setCharClassId(uint256 classId) returns()
func (_World *WorldTransactor) EveworldSetCharClassId(opts *bind.TransactOpts, classId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setCharClassId", classId)
}

// EveworldSetCharClassId is a paid mutator transaction binding the contract method 0x217fd119.
//
// Solidity: function eveworld__setCharClassId(uint256 classId) returns()
func (_World *WorldSession) EveworldSetCharClassId(classId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetCharClassId(&_World.TransactOpts, classId)
}

// EveworldSetCharClassId is a paid mutator transaction binding the contract method 0x217fd119.
//
// Solidity: function eveworld__setCharClassId(uint256 classId) returns()
func (_World *WorldTransactorSession) EveworldSetCharClassId(classId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetCharClassId(&_World.TransactOpts, classId)
}

// EveworldSetCid is a paid mutator transaction binding the contract method 0x0b124c11.
//
// Solidity: function eveworld__setCid(uint256 entityId, string cid) returns()
func (_World *WorldTransactor) EveworldSetCid(opts *bind.TransactOpts, entityId *big.Int, cid string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setCid", entityId, cid)
}

// EveworldSetCid is a paid mutator transaction binding the contract method 0x0b124c11.
//
// Solidity: function eveworld__setCid(uint256 entityId, string cid) returns()
func (_World *WorldSession) EveworldSetCid(entityId *big.Int, cid string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetCid(&_World.TransactOpts, entityId, cid)
}

// EveworldSetCid is a paid mutator transaction binding the contract method 0x0b124c11.
//
// Solidity: function eveworld__setCid(uint256 entityId, string cid) returns()
func (_World *WorldTransactorSession) EveworldSetCid(entityId *big.Int, cid string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetCid(&_World.TransactOpts, entityId, cid)
}

// EveworldSetDappURL is a paid mutator transaction binding the contract method 0x589db2f0.
//
// Solidity: function eveworld__setDappURL(uint256 entityId, string dappURL) returns()
func (_World *WorldTransactor) EveworldSetDappURL(opts *bind.TransactOpts, entityId *big.Int, dappURL string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setDappURL", entityId, dappURL)
}

// EveworldSetDappURL is a paid mutator transaction binding the contract method 0x589db2f0.
//
// Solidity: function eveworld__setDappURL(uint256 entityId, string dappURL) returns()
func (_World *WorldSession) EveworldSetDappURL(entityId *big.Int, dappURL string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetDappURL(&_World.TransactOpts, entityId, dappURL)
}

// EveworldSetDappURL is a paid mutator transaction binding the contract method 0x589db2f0.
//
// Solidity: function eveworld__setDappURL(uint256 entityId, string dappURL) returns()
func (_World *WorldTransactorSession) EveworldSetDappURL(entityId *big.Int, dappURL string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetDappURL(&_World.TransactOpts, entityId, dappURL)
}

// EveworldSetDeployableMetadata is a paid mutator transaction binding the contract method 0x121aefc7.
//
// Solidity: function eveworld__setDeployableMetadata(uint256 smartObjectId, string name, string dappURL, string description) returns()
func (_World *WorldTransactor) EveworldSetDeployableMetadata(opts *bind.TransactOpts, smartObjectId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setDeployableMetadata", smartObjectId, name, dappURL, description)
}

// EveworldSetDeployableMetadata is a paid mutator transaction binding the contract method 0x121aefc7.
//
// Solidity: function eveworld__setDeployableMetadata(uint256 smartObjectId, string name, string dappURL, string description) returns()
func (_World *WorldSession) EveworldSetDeployableMetadata(smartObjectId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetDeployableMetadata(&_World.TransactOpts, smartObjectId, name, dappURL, description)
}

// EveworldSetDeployableMetadata is a paid mutator transaction binding the contract method 0x121aefc7.
//
// Solidity: function eveworld__setDeployableMetadata(uint256 smartObjectId, string name, string dappURL, string description) returns()
func (_World *WorldTransactorSession) EveworldSetDeployableMetadata(smartObjectId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetDeployableMetadata(&_World.TransactOpts, smartObjectId, name, dappURL, description)
}

// EveworldSetDescription is a paid mutator transaction binding the contract method 0xbb7218b2.
//
// Solidity: function eveworld__setDescription(uint256 entityId, string description) returns()
func (_World *WorldTransactor) EveworldSetDescription(opts *bind.TransactOpts, entityId *big.Int, description string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setDescription", entityId, description)
}

// EveworldSetDescription is a paid mutator transaction binding the contract method 0xbb7218b2.
//
// Solidity: function eveworld__setDescription(uint256 entityId, string description) returns()
func (_World *WorldSession) EveworldSetDescription(entityId *big.Int, description string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetDescription(&_World.TransactOpts, entityId, description)
}

// EveworldSetDescription is a paid mutator transaction binding the contract method 0xbb7218b2.
//
// Solidity: function eveworld__setDescription(uint256 entityId, string description) returns()
func (_World *WorldTransactorSession) EveworldSetDescription(entityId *big.Int, description string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetDescription(&_World.TransactOpts, entityId, description)
}

// EveworldSetEntityMetadata is a paid mutator transaction binding the contract method 0xc493045b.
//
// Solidity: function eveworld__setEntityMetadata(uint256 entityId, string name, string dappURL, string description) returns()
func (_World *WorldTransactor) EveworldSetEntityMetadata(opts *bind.TransactOpts, entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setEntityMetadata", entityId, name, dappURL, description)
}

// EveworldSetEntityMetadata is a paid mutator transaction binding the contract method 0xc493045b.
//
// Solidity: function eveworld__setEntityMetadata(uint256 entityId, string name, string dappURL, string description) returns()
func (_World *WorldSession) EveworldSetEntityMetadata(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetEntityMetadata(&_World.TransactOpts, entityId, name, dappURL, description)
}

// EveworldSetEntityMetadata is a paid mutator transaction binding the contract method 0xc493045b.
//
// Solidity: function eveworld__setEntityMetadata(uint256 entityId, string name, string dappURL, string description) returns()
func (_World *WorldTransactorSession) EveworldSetEntityMetadata(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetEntityMetadata(&_World.TransactOpts, entityId, name, dappURL, description)
}

// EveworldSetEphemeralInventoryCapacity is a paid mutator transaction binding the contract method 0xba4034b6.
//
// Solidity: function eveworld__setEphemeralInventoryCapacity(uint256 smartObjectId, uint256 ephemeralStorageCapacity) returns()
func (_World *WorldTransactor) EveworldSetEphemeralInventoryCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setEphemeralInventoryCapacity", smartObjectId, ephemeralStorageCapacity)
}

// EveworldSetEphemeralInventoryCapacity is a paid mutator transaction binding the contract method 0xba4034b6.
//
// Solidity: function eveworld__setEphemeralInventoryCapacity(uint256 smartObjectId, uint256 ephemeralStorageCapacity) returns()
func (_World *WorldSession) EveworldSetEphemeralInventoryCapacity(smartObjectId *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetEphemeralInventoryCapacity(&_World.TransactOpts, smartObjectId, ephemeralStorageCapacity)
}

// EveworldSetEphemeralInventoryCapacity is a paid mutator transaction binding the contract method 0xba4034b6.
//
// Solidity: function eveworld__setEphemeralInventoryCapacity(uint256 smartObjectId, uint256 ephemeralStorageCapacity) returns()
func (_World *WorldTransactorSession) EveworldSetEphemeralInventoryCapacity(smartObjectId *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetEphemeralInventoryCapacity(&_World.TransactOpts, smartObjectId, ephemeralStorageCapacity)
}

// EveworldSetFuelConsumptionPerMinute is a paid mutator transaction binding the contract method 0x906c8d2d.
//
// Solidity: function eveworld__setFuelConsumptionPerMinute(uint256 entityId, uint256 fuelConsumptionIntervalInSeconds) returns()
func (_World *WorldTransactor) EveworldSetFuelConsumptionPerMinute(opts *bind.TransactOpts, entityId *big.Int, fuelConsumptionIntervalInSeconds *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setFuelConsumptionPerMinute", entityId, fuelConsumptionIntervalInSeconds)
}

// EveworldSetFuelConsumptionPerMinute is a paid mutator transaction binding the contract method 0x906c8d2d.
//
// Solidity: function eveworld__setFuelConsumptionPerMinute(uint256 entityId, uint256 fuelConsumptionIntervalInSeconds) returns()
func (_World *WorldSession) EveworldSetFuelConsumptionPerMinute(entityId *big.Int, fuelConsumptionIntervalInSeconds *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetFuelConsumptionPerMinute(&_World.TransactOpts, entityId, fuelConsumptionIntervalInSeconds)
}

// EveworldSetFuelConsumptionPerMinute is a paid mutator transaction binding the contract method 0x906c8d2d.
//
// Solidity: function eveworld__setFuelConsumptionPerMinute(uint256 entityId, uint256 fuelConsumptionIntervalInSeconds) returns()
func (_World *WorldTransactorSession) EveworldSetFuelConsumptionPerMinute(entityId *big.Int, fuelConsumptionIntervalInSeconds *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetFuelConsumptionPerMinute(&_World.TransactOpts, entityId, fuelConsumptionIntervalInSeconds)
}

// EveworldSetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xcd80c3ae.
//
// Solidity: function eveworld__setFuelMaxCapacity(uint256 entityId, uint256 capacityInWei) returns()
func (_World *WorldTransactor) EveworldSetFuelMaxCapacity(opts *bind.TransactOpts, entityId *big.Int, capacityInWei *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setFuelMaxCapacity", entityId, capacityInWei)
}

// EveworldSetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xcd80c3ae.
//
// Solidity: function eveworld__setFuelMaxCapacity(uint256 entityId, uint256 capacityInWei) returns()
func (_World *WorldSession) EveworldSetFuelMaxCapacity(entityId *big.Int, capacityInWei *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetFuelMaxCapacity(&_World.TransactOpts, entityId, capacityInWei)
}

// EveworldSetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xcd80c3ae.
//
// Solidity: function eveworld__setFuelMaxCapacity(uint256 entityId, uint256 capacityInWei) returns()
func (_World *WorldTransactorSession) EveworldSetFuelMaxCapacity(entityId *big.Int, capacityInWei *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetFuelMaxCapacity(&_World.TransactOpts, entityId, capacityInWei)
}

// EveworldSetInventoryCapacity is a paid mutator transaction binding the contract method 0xce754284.
//
// Solidity: function eveworld__setInventoryCapacity(uint256 smartObjectId, uint256 storageCapacity) returns()
func (_World *WorldTransactor) EveworldSetInventoryCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, storageCapacity *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setInventoryCapacity", smartObjectId, storageCapacity)
}

// EveworldSetInventoryCapacity is a paid mutator transaction binding the contract method 0xce754284.
//
// Solidity: function eveworld__setInventoryCapacity(uint256 smartObjectId, uint256 storageCapacity) returns()
func (_World *WorldSession) EveworldSetInventoryCapacity(smartObjectId *big.Int, storageCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetInventoryCapacity(&_World.TransactOpts, smartObjectId, storageCapacity)
}

// EveworldSetInventoryCapacity is a paid mutator transaction binding the contract method 0xce754284.
//
// Solidity: function eveworld__setInventoryCapacity(uint256 smartObjectId, uint256 storageCapacity) returns()
func (_World *WorldTransactorSession) EveworldSetInventoryCapacity(smartObjectId *big.Int, storageCapacity *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetInventoryCapacity(&_World.TransactOpts, smartObjectId, storageCapacity)
}

// EveworldSetMetadata is a paid mutator transaction binding the contract method 0xddc0ed7f.
//
// Solidity: function eveworld__setMetadata(bytes32 systemId, (string,string,string) data) returns()
func (_World *WorldTransactor) EveworldSetMetadata(opts *bind.TransactOpts, systemId [32]byte, data StaticDataGlobalTableData) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setMetadata", systemId, data)
}

// EveworldSetMetadata is a paid mutator transaction binding the contract method 0xddc0ed7f.
//
// Solidity: function eveworld__setMetadata(bytes32 systemId, (string,string,string) data) returns()
func (_World *WorldSession) EveworldSetMetadata(systemId [32]byte, data StaticDataGlobalTableData) (*types.Transaction, error) {
	return _World.Contract.EveworldSetMetadata(&_World.TransactOpts, systemId, data)
}

// EveworldSetMetadata is a paid mutator transaction binding the contract method 0xddc0ed7f.
//
// Solidity: function eveworld__setMetadata(bytes32 systemId, (string,string,string) data) returns()
func (_World *WorldTransactorSession) EveworldSetMetadata(systemId [32]byte, data StaticDataGlobalTableData) (*types.Transaction, error) {
	return _World.Contract.EveworldSetMetadata(&_World.TransactOpts, systemId, data)
}

// EveworldSetName is a paid mutator transaction binding the contract method 0x37453fd9.
//
// Solidity: function eveworld__setName(uint256 entityId, string name) returns()
func (_World *WorldTransactor) EveworldSetName(opts *bind.TransactOpts, entityId *big.Int, name string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setName", entityId, name)
}

// EveworldSetName is a paid mutator transaction binding the contract method 0x37453fd9.
//
// Solidity: function eveworld__setName(uint256 entityId, string name) returns()
func (_World *WorldSession) EveworldSetName(entityId *big.Int, name string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetName(&_World.TransactOpts, entityId, name)
}

// EveworldSetName is a paid mutator transaction binding the contract method 0x37453fd9.
//
// Solidity: function eveworld__setName(uint256 entityId, string name) returns()
func (_World *WorldTransactorSession) EveworldSetName(entityId *big.Int, name string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetName(&_World.TransactOpts, entityId, name)
}

// EveworldSetName0 is a paid mutator transaction binding the contract method 0x648bb531.
//
// Solidity: function eveworld__setName(bytes32 systemId, string name) returns()
func (_World *WorldTransactor) EveworldSetName0(opts *bind.TransactOpts, systemId [32]byte, name string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setName0", systemId, name)
}

// EveworldSetName0 is a paid mutator transaction binding the contract method 0x648bb531.
//
// Solidity: function eveworld__setName(bytes32 systemId, string name) returns()
func (_World *WorldSession) EveworldSetName0(systemId [32]byte, name string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetName0(&_World.TransactOpts, systemId, name)
}

// EveworldSetName0 is a paid mutator transaction binding the contract method 0x648bb531.
//
// Solidity: function eveworld__setName(bytes32 systemId, string name) returns()
func (_World *WorldTransactorSession) EveworldSetName0(systemId [32]byte, name string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetName0(&_World.TransactOpts, systemId, name)
}

// EveworldSetSSUClassId is a paid mutator transaction binding the contract method 0xddd96344.
//
// Solidity: function eveworld__setSSUClassId(uint256 classId) returns()
func (_World *WorldTransactor) EveworldSetSSUClassId(opts *bind.TransactOpts, classId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setSSUClassId", classId)
}

// EveworldSetSSUClassId is a paid mutator transaction binding the contract method 0xddd96344.
//
// Solidity: function eveworld__setSSUClassId(uint256 classId) returns()
func (_World *WorldSession) EveworldSetSSUClassId(classId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetSSUClassId(&_World.TransactOpts, classId)
}

// EveworldSetSSUClassId is a paid mutator transaction binding the contract method 0xddd96344.
//
// Solidity: function eveworld__setSSUClassId(uint256 classId) returns()
func (_World *WorldTransactorSession) EveworldSetSSUClassId(classId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldSetSSUClassId(&_World.TransactOpts, classId)
}

// EveworldSetSmartAssemblyType is a paid mutator transaction binding the contract method 0xe8bb13da.
//
// Solidity: function eveworld__setSmartAssemblyType(uint256 entityId, uint8 smartAssemblyType) returns()
func (_World *WorldTransactor) EveworldSetSmartAssemblyType(opts *bind.TransactOpts, entityId *big.Int, smartAssemblyType uint8) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setSmartAssemblyType", entityId, smartAssemblyType)
}

// EveworldSetSmartAssemblyType is a paid mutator transaction binding the contract method 0xe8bb13da.
//
// Solidity: function eveworld__setSmartAssemblyType(uint256 entityId, uint8 smartAssemblyType) returns()
func (_World *WorldSession) EveworldSetSmartAssemblyType(entityId *big.Int, smartAssemblyType uint8) (*types.Transaction, error) {
	return _World.Contract.EveworldSetSmartAssemblyType(&_World.TransactOpts, entityId, smartAssemblyType)
}

// EveworldSetSmartAssemblyType is a paid mutator transaction binding the contract method 0xe8bb13da.
//
// Solidity: function eveworld__setSmartAssemblyType(uint256 entityId, uint8 smartAssemblyType) returns()
func (_World *WorldTransactorSession) EveworldSetSmartAssemblyType(entityId *big.Int, smartAssemblyType uint8) (*types.Transaction, error) {
	return _World.Contract.EveworldSetSmartAssemblyType(&_World.TransactOpts, entityId, smartAssemblyType)
}

// EveworldSetSymbol is a paid mutator transaction binding the contract method 0x08e7ccb8.
//
// Solidity: function eveworld__setSymbol(bytes32 systemId, string symbol) returns()
func (_World *WorldTransactor) EveworldSetSymbol(opts *bind.TransactOpts, systemId [32]byte, symbol string) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__setSymbol", systemId, symbol)
}

// EveworldSetSymbol is a paid mutator transaction binding the contract method 0x08e7ccb8.
//
// Solidity: function eveworld__setSymbol(bytes32 systemId, string symbol) returns()
func (_World *WorldSession) EveworldSetSymbol(systemId [32]byte, symbol string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetSymbol(&_World.TransactOpts, systemId, symbol)
}

// EveworldSetSymbol is a paid mutator transaction binding the contract method 0x08e7ccb8.
//
// Solidity: function eveworld__setSymbol(bytes32 systemId, string symbol) returns()
func (_World *WorldTransactorSession) EveworldSetSymbol(systemId [32]byte, symbol string) (*types.Transaction, error) {
	return _World.Contract.EveworldSetSymbol(&_World.TransactOpts, systemId, symbol)
}

// EveworldUnanchor is a paid mutator transaction binding the contract method 0x694e5e2c.
//
// Solidity: function eveworld__unanchor(uint256 entityId) returns()
func (_World *WorldTransactor) EveworldUnanchor(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__unanchor", entityId)
}

// EveworldUnanchor is a paid mutator transaction binding the contract method 0x694e5e2c.
//
// Solidity: function eveworld__unanchor(uint256 entityId) returns()
func (_World *WorldSession) EveworldUnanchor(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldUnanchor(&_World.TransactOpts, entityId)
}

// EveworldUnanchor is a paid mutator transaction binding the contract method 0x694e5e2c.
//
// Solidity: function eveworld__unanchor(uint256 entityId) returns()
func (_World *WorldTransactorSession) EveworldUnanchor(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldUnanchor(&_World.TransactOpts, entityId)
}

// EveworldUnlinkSmartGates is a paid mutator transaction binding the contract method 0xce0bf1e6.
//
// Solidity: function eveworld__unlinkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldTransactor) EveworldUnlinkSmartGates(opts *bind.TransactOpts, sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__unlinkSmartGates", sourceGateId, destinationGateId)
}

// EveworldUnlinkSmartGates is a paid mutator transaction binding the contract method 0xce0bf1e6.
//
// Solidity: function eveworld__unlinkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldSession) EveworldUnlinkSmartGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldUnlinkSmartGates(&_World.TransactOpts, sourceGateId, destinationGateId)
}

// EveworldUnlinkSmartGates is a paid mutator transaction binding the contract method 0xce0bf1e6.
//
// Solidity: function eveworld__unlinkSmartGates(uint256 sourceGateId, uint256 destinationGateId) returns()
func (_World *WorldTransactorSession) EveworldUnlinkSmartGates(sourceGateId *big.Int, destinationGateId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldUnlinkSmartGates(&_World.TransactOpts, sourceGateId, destinationGateId)
}

// EveworldUpdateCorpId is a paid mutator transaction binding the contract method 0x715464e7.
//
// Solidity: function eveworld__updateCorpId(uint256 characterId, uint256 corpId) returns()
func (_World *WorldTransactor) EveworldUpdateCorpId(opts *bind.TransactOpts, characterId *big.Int, corpId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__updateCorpId", characterId, corpId)
}

// EveworldUpdateCorpId is a paid mutator transaction binding the contract method 0x715464e7.
//
// Solidity: function eveworld__updateCorpId(uint256 characterId, uint256 corpId) returns()
func (_World *WorldSession) EveworldUpdateCorpId(characterId *big.Int, corpId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldUpdateCorpId(&_World.TransactOpts, characterId, corpId)
}

// EveworldUpdateCorpId is a paid mutator transaction binding the contract method 0x715464e7.
//
// Solidity: function eveworld__updateCorpId(uint256 characterId, uint256 corpId) returns()
func (_World *WorldTransactorSession) EveworldUpdateCorpId(characterId *big.Int, corpId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldUpdateCorpId(&_World.TransactOpts, characterId, corpId)
}

// EveworldUpdateFuel is a paid mutator transaction binding the contract method 0x77eb0abf.
//
// Solidity: function eveworld__updateFuel(uint256 entityId) returns()
func (_World *WorldTransactor) EveworldUpdateFuel(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__updateFuel", entityId)
}

// EveworldUpdateFuel is a paid mutator transaction binding the contract method 0x77eb0abf.
//
// Solidity: function eveworld__updateFuel(uint256 entityId) returns()
func (_World *WorldSession) EveworldUpdateFuel(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldUpdateFuel(&_World.TransactOpts, entityId)
}

// EveworldUpdateFuel is a paid mutator transaction binding the contract method 0x77eb0abf.
//
// Solidity: function eveworld__updateFuel(uint256 entityId) returns()
func (_World *WorldTransactorSession) EveworldUpdateFuel(entityId *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldUpdateFuel(&_World.TransactOpts, entityId)
}

// EveworldWithdrawFromEphemeralInventory is a paid mutator transaction binding the contract method 0x54160975.
//
// Solidity: function eveworld__withdrawFromEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EveworldWithdrawFromEphemeralInventory(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__withdrawFromEphemeralInventory", smartObjectId, ephemeralInventoryOwner, items)
}

// EveworldWithdrawFromEphemeralInventory is a paid mutator transaction binding the contract method 0x54160975.
//
// Solidity: function eveworld__withdrawFromEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldSession) EveworldWithdrawFromEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldWithdrawFromEphemeralInventory(&_World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// EveworldWithdrawFromEphemeralInventory is a paid mutator transaction binding the contract method 0x54160975.
//
// Solidity: function eveworld__withdrawFromEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EveworldWithdrawFromEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldWithdrawFromEphemeralInventory(&_World.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// EveworldWithdrawFromInventory is a paid mutator transaction binding the contract method 0x6e7e1184.
//
// Solidity: function eveworld__withdrawFromInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactor) EveworldWithdrawFromInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__withdrawFromInventory", smartObjectId, items)
}

// EveworldWithdrawFromInventory is a paid mutator transaction binding the contract method 0x6e7e1184.
//
// Solidity: function eveworld__withdrawFromInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldSession) EveworldWithdrawFromInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldWithdrawFromInventory(&_World.TransactOpts, smartObjectId, items)
}

// EveworldWithdrawFromInventory is a paid mutator transaction binding the contract method 0x6e7e1184.
//
// Solidity: function eveworld__withdrawFromInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_World *WorldTransactorSession) EveworldWithdrawFromInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _World.Contract.EveworldWithdrawFromInventory(&_World.TransactOpts, smartObjectId, items)
}

// EveworldWithdrawFuel is a paid mutator transaction binding the contract method 0x365d7949.
//
// Solidity: function eveworld__withdrawFuel(uint256 entityId, uint256 unitAmount) returns()
func (_World *WorldTransactor) EveworldWithdrawFuel(opts *bind.TransactOpts, entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "eveworld__withdrawFuel", entityId, unitAmount)
}

// EveworldWithdrawFuel is a paid mutator transaction binding the contract method 0x365d7949.
//
// Solidity: function eveworld__withdrawFuel(uint256 entityId, uint256 unitAmount) returns()
func (_World *WorldSession) EveworldWithdrawFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldWithdrawFuel(&_World.TransactOpts, entityId, unitAmount)
}

// EveworldWithdrawFuel is a paid mutator transaction binding the contract method 0x365d7949.
//
// Solidity: function eveworld__withdrawFuel(uint256 entityId, uint256 unitAmount) returns()
func (_World *WorldTransactorSession) EveworldWithdrawFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _World.Contract.EveworldWithdrawFuel(&_World.TransactOpts, entityId, unitAmount)
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

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x742d6118.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, bytes4 systemFunctionSelector) returns(bytes4 worldFunctionSelector)
func (_World *WorldTransactor) RegisterRootFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, worldFunctionSignature string, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _World.contract.Transact(opts, "registerRootFunctionSelector", systemId, worldFunctionSignature, systemFunctionSelector)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x742d6118.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, bytes4 systemFunctionSelector) returns(bytes4 worldFunctionSelector)
func (_World *WorldSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _World.Contract.RegisterRootFunctionSelector(&_World.TransactOpts, systemId, worldFunctionSignature, systemFunctionSelector)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x742d6118.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, bytes4 systemFunctionSelector) returns(bytes4 worldFunctionSelector)
func (_World *WorldTransactorSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _World.Contract.RegisterRootFunctionSelector(&_World.TransactOpts, systemId, worldFunctionSignature, systemFunctionSelector)
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

// WatchStoreSpliceDynamicData is a free log subscription operation binding the contract event 0xaa63765a776145e5e6492f471ae097dfed11cd57a61bc2679dd43180422385b4.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
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

// ParseStoreSpliceDynamicData is a log parse operation binding the contract event 0xaa63765a776145e5e6492f471ae097dfed11cd57a61bc2679dd43180422385b4.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
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
