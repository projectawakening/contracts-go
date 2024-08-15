// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package worldErrors

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

// ERC2771ForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type ERC2771ForwarderForwardRequest struct {
	From     common.Address
	To       common.Address
	Value    *big.Int
	Gas      *big.Int
	Nonce    *big.Int
	Deadline *big.Int
	Data     []byte
}

// ERC2771ForwarderForwardRequestData is an auto generated low-level Go binding around an user-defined struct.
type ERC2771ForwarderForwardRequestData struct {
	From      common.Address
	To        common.Address
	Value     *big.Int
	Gas       *big.Int
	Nonce     *big.Int
	Deadline  *big.Int
	Data      []byte
	Signature []byte
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

// WorldPosition is an auto generated low-level Go binding around an user-defined struct.
type WorldPosition struct {
	SolarSystemId *big.Int
	Position      Coord
}

// WorldErrorsMetaData contains all meta data concerning the WorldErrors contract.
var WorldErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2771ForwarderExpiredRequest\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"ERC2771ForwarderInvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2771ForwarderMismatchedValue\",\"inputs\":[{\"name\":\"requestedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2771UntrustfulTarget\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721Module_InvalidNamespace\",\"inputs\":[{\"name\":\"namespace\",\"type\":\"bytes14\",\"internalType\":\"bytes14\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EntityAlreadyAssociated\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"associatedId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EntityAlreadyRegistered\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EntityAlreadyTagged\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tagId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EntityNotAssociatedWithModule\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EntityNotRegistered\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EntityTypeAlreadyRegistered\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EntityTypeAssociationNotAllowed\",\"inputs\":[{\"name\":\"entityType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taggedEntityType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EntityTypeNotRegistered\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FieldLayoutLib_Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FieldLayoutLib_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"staticDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"computedStaticDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayoutLib_StaticLengthDoesNotFitInAWord\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayoutLib_StaticLengthIsNotZero\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayoutLib_StaticLengthIsZero\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayoutLib_TooManyDynamicFields\",\"inputs\":[{\"name\":\"numFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FieldLayoutLib_TooManyFields\",\"inputs\":[{\"name\":\"numFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFields\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FunctionSelectorNotRegistered\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"HookAlreadyRegistered\",\"inputs\":[{\"name\":\"hookId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"HookNotRegistered\",\"inputs\":[{\"name\":\"hookId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidArrayLength\",\"inputs\":[{\"name\":\"length1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length2\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidEntityId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidModuleId\",\"inputs\":[{\"name\":\"moduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Inventory_InsufficientCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"usedCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidCapacity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidDeployable\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deployableId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidItemQuantity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidItem\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidOwner\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ephemeralInvOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Inventory_InvalidQuantity\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxQuantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LocationModule_InvalidNamespace\",\"inputs\":[{\"name\":\"namespace\",\"type\":\"bytes14\",\"internalType\":\"bytes14\"}]},{\"type\":\"error\",\"name\":\"ModuleNotFound\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ModuleNotRegistered\",\"inputs\":[{\"name\":\"moduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Module_AlreadyInstalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Module_MissingDependency\",\"inputs\":[{\"name\":\"dependency\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Module_NonRootInstallNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Module_RootInstallNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PackedCounter_InvalidLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ResourceNotRegistered\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"SchemaLib_InvalidLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SchemaLib_StaticTypeAfterDynamicType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Slice_OutOfBounds\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartCharacterERC721AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartCharacterTokenCidCannotBeEmpty\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenCid\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"SmartCharacter_ERC721AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartCharacter_UndefinedClassIds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartDeployableERC721AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartDeployable_IncorrectState\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentState\",\"type\":\"uint8\",\"internalType\":\"enumState\"}]},{\"type\":\"error\",\"name\":\"SmartDeployable_NoFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartDeployable_StateTransitionPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartDeployable_TooMuchFuelDeposited\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountDeposited\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SmartStorageUnitERC721AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmartStorageUnit_UndefinedClassId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Store_IndexOutOfBounds\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessedIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidBounds\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidFieldNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidKeyNamesLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidSplice\",\"inputs\":[{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"fieldLength\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidStaticDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaDynamicLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_InvalidValueSchemaStaticLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Store_TableAlreadyExists\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Store_TableNotFound\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"tableIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"SystemAlreadyAssociatedWithModule\",\"inputs\":[{\"name\":\"moduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"SystemNotAssociatedWithModule\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_AccessDenied\",\"inputs\":[{\"name\":\"resource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"World_CallbackNotAllowed\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_DelegationNotFound\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_DirectCallToSystemForbidden\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorAlreadyExists\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_FunctionSelectorNotFound\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"World_InterfaceNotSupported\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"World_InvalidNamespace\",\"inputs\":[{\"name\":\"namespace\",\"type\":\"bytes14\",\"internalType\":\"bytes14\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceId\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_InvalidResourceType\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceAlreadyExists\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_ResourceNotFound\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"resourceIdString\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"World_SystemAlreadyExists\",\"inputs\":[{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"World_UnlimitedDelegationNotAllowed\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedForwardRequest\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HelloStore\",\"inputs\":[{\"name\":\"storeVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HelloWorld\",\"inputs\":[{\"name\":\"worldVersion\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_DeleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SetRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"PackedCounter\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Store_SpliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"function\",\"name\":\"_msgSender\",\"inputs\":[],\"outputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_msgValue\",\"inputs\":[],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"_world\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addHook\",\"inputs\":[{\"name\":\"hookId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hookType\",\"type\":\"uint8\",\"internalType\":\"enumHookType\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"anchor\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"locationData\",\"type\":\"tuple\",\"internalType\":\"structLocationTableData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"associateHooks\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hookIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"associateHook\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hookId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"associateModules\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"associateModule\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchCallFrom\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallFromData[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchCall\",\"inputs\":[{\"name\":\"systemCalls\",\"type\":\"tuple[]\",\"internalType\":\"structSystemCallData[]\",\"components\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bringOffline\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bringOnline\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callFrom\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"characterSystemId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configureInteractionHandler\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interactionParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndAnchorSmartStorageUnit\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityRecordData\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"worldPosition\",\"type\":\"tuple\",\"internalType\":\"structWorldPosition\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"position\",\"type\":\"tuple\",\"internalType\":\"structCoord\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"fuelUnitVolume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionPerMinute\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralStorageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndDepositItemsToEphemeralInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndDepositItemsToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createCharacter\",\"inputs\":[{\"name\":\"characterId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"characterAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entityRecord\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordData\",\"components\":[{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"entityRecordOffchain\",\"type\":\"tuple\",\"internalType\":\"structEntityRecordOffchainTableData\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"tokenCid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEntityRecordOffchain\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEntityRecord\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"creator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentFuelAmount\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositToEphemeralInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositToInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"destroyDeployable\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ephemeralToInventoryTransfer\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeBatch\",\"inputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structERC2771Forwarder.ForwardRequestData[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"refundReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structERC2771Forwarder.ForwardRequestData\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getDynamicFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicFieldSlice\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLayout\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFieldLength\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeySchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValueSchema\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"globalResume\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"trustedForwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialMsgSender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initModule\",\"type\":\"address\",\"internalType\":\"contractIModule\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installRootModule\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"contractIModule\"},{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"installRoot\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"install\",\"inputs\":[{\"name\":\"encodedArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"inventoryToEphemeralTransfer\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outItems\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isNonceUsed\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"popFromDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"byteLengthToPop\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pushToDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dataToPush\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDeployableToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDeployable\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"smartObjectData\",\"type\":\"tuple\",\"internalType\":\"structSmartObjectData\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"fuelUnitVolumeInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionPerMinuteInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelMaxCapacityInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerERC721Token\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerEVEModules\",\"inputs\":[{\"name\":\"moduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleName\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"},{\"name\":\"systemIds\",\"type\":\"bytes32[]\",\"internalType\":\"ResourceId[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerEVEModule\",\"inputs\":[{\"name\":\"moduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleName\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerEntities\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"entityType\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerEntityTypeAssociation\",\"inputs\":[{\"name\":\"entityType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tagEntityType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerEntityType\",\"inputs\":[{\"name\":\"entityTypeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"entityType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerEntity\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"systemFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"functionId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"delegationControlId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"initCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerNamespace\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerRootFunctionSelector\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"worldFunctionSignature\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"systemFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"},{\"name\":\"enabledHooksBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSystem\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"system\",\"type\":\"address\",\"internalType\":\"contractSystem\"},{\"name\":\"publicAccess\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerTable\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"},{\"name\":\"keySchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"valueSchema\",\"type\":\"bytes32\",\"internalType\":\"Schema\"},{\"name\":\"keyNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"fieldNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeEntityHookAssociation\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hookId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeEntityModuleAssociation\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeEntityTag\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityTagId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeHook\",\"inputs\":[{\"name\":\"hookId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hookType\",\"type\":\"uint8\",\"internalType\":\"enumHookType\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeSystemModuleAssociation\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"moduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeAccess\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"grantee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"saveLocation\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"location\",\"type\":\"tuple\",\"internalType\":\"structLocationTableData\",\"components\":[{\"name\":\"solarSystemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseURI\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"baseURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCharClassId\",\"inputs\":[{\"name\":\"classId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCid\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDappURL\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDeploybaleMetadata\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDescription\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEntityMetadata\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dappURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEphemeralInventoryCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralStorageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFuelConsumptionPerMinute\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fuelConsumptionPerMinuteInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFuelMaxCapacity\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacityInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInventoryCapacity\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storageCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadata\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structStaticDataGlobalTableData\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"baseURI\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRecord\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"staticData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encodedLengths\",\"type\":\"bytes32\",\"internalType\":\"PackedCounter\"},{\"name\":\"dynamicData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSSUClassId\",\"inputs\":[{\"name\":\"classId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaticField\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"fieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fieldLayout\",\"type\":\"bytes32\",\"internalType\":\"FieldLayout\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTrustedForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceDynamicData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dynamicFieldIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startWithinField\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deleteCount\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spliceStaticData\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"keyTuple\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"start\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storeVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"structHash\",\"inputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structERC2771Forwarder.ForwardRequest\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tagEntities\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityTagIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tagEntity\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entityTagId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferBalanceToAddress\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferBalanceToNamespace\",\"inputs\":[{\"name\":\"fromNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"toNamespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unanchor\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterDelegation\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterNamespaceDelegation\",\"inputs\":[{\"name\":\"namespaceId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterStoreHook\",\"inputs\":[{\"name\":\"tableId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractIStoreHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterSystemHook\",\"inputs\":[{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"hookAddress\",\"type\":\"address\",\"internalType\":\"contractISystemHook\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemId\",\"type\":\"bytes32\",\"internalType\":\"ResourceId\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structERC2771Forwarder.ForwardRequestData\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFromEphemeralInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ephemeralInventoryOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFromInventory\",\"inputs\":[{\"name\":\"smartObjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structInventoryItem[]\",\"components\":[{\"name\":\"inventoryItemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFuel\",\"inputs\":[{\"name\":\"entityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"worldVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"AccessControl_AccessConfigAccessDenied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControl_InvalidRoleId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControl_NoPermission\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roleId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// WorldErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use WorldErrorsMetaData.ABI instead.
var WorldErrorsABI = WorldErrorsMetaData.ABI

// WorldErrors is an auto generated Go binding around an Ethereum contract.
type WorldErrors struct {
	WorldErrorsCaller     // Read-only binding to the contract
	WorldErrorsTransactor // Write-only binding to the contract
	WorldErrorsFilterer   // Log filterer for contract events
}

// WorldErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorldErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorldErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorldErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorldErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorldErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorldErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorldErrorsSession struct {
	Contract     *WorldErrors      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorldErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorldErrorsCallerSession struct {
	Contract *WorldErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WorldErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorldErrorsTransactorSession struct {
	Contract     *WorldErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WorldErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorldErrorsRaw struct {
	Contract *WorldErrors // Generic contract binding to access the raw methods on
}

// WorldErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorldErrorsCallerRaw struct {
	Contract *WorldErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// WorldErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorldErrorsTransactorRaw struct {
	Contract *WorldErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorldErrors creates a new instance of WorldErrors, bound to a specific deployed contract.
func NewWorldErrors(address common.Address, backend bind.ContractBackend) (*WorldErrors, error) {
	contract, err := bindWorldErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorldErrors{WorldErrorsCaller: WorldErrorsCaller{contract: contract}, WorldErrorsTransactor: WorldErrorsTransactor{contract: contract}, WorldErrorsFilterer: WorldErrorsFilterer{contract: contract}}, nil
}

// NewWorldErrorsCaller creates a new read-only instance of WorldErrors, bound to a specific deployed contract.
func NewWorldErrorsCaller(address common.Address, caller bind.ContractCaller) (*WorldErrorsCaller, error) {
	contract, err := bindWorldErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsCaller{contract: contract}, nil
}

// NewWorldErrorsTransactor creates a new write-only instance of WorldErrors, bound to a specific deployed contract.
func NewWorldErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*WorldErrorsTransactor, error) {
	contract, err := bindWorldErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsTransactor{contract: contract}, nil
}

// NewWorldErrorsFilterer creates a new log filterer instance of WorldErrors, bound to a specific deployed contract.
func NewWorldErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*WorldErrorsFilterer, error) {
	contract, err := bindWorldErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsFilterer{contract: contract}, nil
}

// bindWorldErrors binds a generic wrapper to an already deployed contract.
func bindWorldErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorldErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorldErrors *WorldErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorldErrors.Contract.WorldErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorldErrors *WorldErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorldErrors.Contract.WorldErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorldErrors *WorldErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorldErrors.Contract.WorldErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorldErrors *WorldErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorldErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorldErrors *WorldErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorldErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorldErrors *WorldErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorldErrors.Contract.contract.Transact(opts, method, params...)
}

// MsgSender is a free data retrieval call binding the contract method 0x119df25f.
//
// Solidity: function _msgSender() view returns(address sender)
func (_WorldErrors *WorldErrorsCaller) MsgSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "_msgSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MsgSender is a free data retrieval call binding the contract method 0x119df25f.
//
// Solidity: function _msgSender() view returns(address sender)
func (_WorldErrors *WorldErrorsSession) MsgSender() (common.Address, error) {
	return _WorldErrors.Contract.MsgSender(&_WorldErrors.CallOpts)
}

// MsgSender is a free data retrieval call binding the contract method 0x119df25f.
//
// Solidity: function _msgSender() view returns(address sender)
func (_WorldErrors *WorldErrorsCallerSession) MsgSender() (common.Address, error) {
	return _WorldErrors.Contract.MsgSender(&_WorldErrors.CallOpts)
}

// MsgValue is a free data retrieval call binding the contract method 0x45ec9354.
//
// Solidity: function _msgValue() pure returns(uint256 value)
func (_WorldErrors *WorldErrorsCaller) MsgValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "_msgValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MsgValue is a free data retrieval call binding the contract method 0x45ec9354.
//
// Solidity: function _msgValue() pure returns(uint256 value)
func (_WorldErrors *WorldErrorsSession) MsgValue() (*big.Int, error) {
	return _WorldErrors.Contract.MsgValue(&_WorldErrors.CallOpts)
}

// MsgValue is a free data retrieval call binding the contract method 0x45ec9354.
//
// Solidity: function _msgValue() pure returns(uint256 value)
func (_WorldErrors *WorldErrorsCallerSession) MsgValue() (*big.Int, error) {
	return _WorldErrors.Contract.MsgValue(&_WorldErrors.CallOpts)
}

// World is a free data retrieval call binding the contract method 0xe1af802c.
//
// Solidity: function _world() view returns(address)
func (_WorldErrors *WorldErrorsCaller) World(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "_world")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// World is a free data retrieval call binding the contract method 0xe1af802c.
//
// Solidity: function _world() view returns(address)
func (_WorldErrors *WorldErrorsSession) World() (common.Address, error) {
	return _WorldErrors.Contract.World(&_WorldErrors.CallOpts)
}

// World is a free data retrieval call binding the contract method 0xe1af802c.
//
// Solidity: function _world() view returns(address)
func (_WorldErrors *WorldErrorsCallerSession) World() (common.Address, error) {
	return _WorldErrors.Contract.World(&_WorldErrors.CallOpts)
}

// CharacterSystemId is a free data retrieval call binding the contract method 0xc40ce8b7.
//
// Solidity: function characterSystemId() view returns(bytes32)
func (_WorldErrors *WorldErrorsCaller) CharacterSystemId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "characterSystemId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CharacterSystemId is a free data retrieval call binding the contract method 0xc40ce8b7.
//
// Solidity: function characterSystemId() view returns(bytes32)
func (_WorldErrors *WorldErrorsSession) CharacterSystemId() ([32]byte, error) {
	return _WorldErrors.Contract.CharacterSystemId(&_WorldErrors.CallOpts)
}

// CharacterSystemId is a free data retrieval call binding the contract method 0xc40ce8b7.
//
// Solidity: function characterSystemId() view returns(bytes32)
func (_WorldErrors *WorldErrorsCallerSession) CharacterSystemId() ([32]byte, error) {
	return _WorldErrors.Contract.CharacterSystemId(&_WorldErrors.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_WorldErrors *WorldErrorsCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_WorldErrors *WorldErrorsSession) Creator() (common.Address, error) {
	return _WorldErrors.Contract.Creator(&_WorldErrors.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_WorldErrors *WorldErrorsCallerSession) Creator() (common.Address, error) {
	return _WorldErrors.Contract.Creator(&_WorldErrors.CallOpts)
}

// CurrentFuelAmount is a free data retrieval call binding the contract method 0x22d3581d.
//
// Solidity: function currentFuelAmount(uint256 entityId) view returns(uint256 amount)
func (_WorldErrors *WorldErrorsCaller) CurrentFuelAmount(opts *bind.CallOpts, entityId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "currentFuelAmount", entityId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFuelAmount is a free data retrieval call binding the contract method 0x22d3581d.
//
// Solidity: function currentFuelAmount(uint256 entityId) view returns(uint256 amount)
func (_WorldErrors *WorldErrorsSession) CurrentFuelAmount(entityId *big.Int) (*big.Int, error) {
	return _WorldErrors.Contract.CurrentFuelAmount(&_WorldErrors.CallOpts, entityId)
}

// CurrentFuelAmount is a free data retrieval call binding the contract method 0x22d3581d.
//
// Solidity: function currentFuelAmount(uint256 entityId) view returns(uint256 amount)
func (_WorldErrors *WorldErrorsCallerSession) CurrentFuelAmount(entityId *big.Int) (*big.Int, error) {
	return _WorldErrors.Contract.CurrentFuelAmount(&_WorldErrors.CallOpts, entityId)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WorldErrors *WorldErrorsCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WorldErrors *WorldErrorsSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WorldErrors.Contract.Eip712Domain(&_WorldErrors.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WorldErrors *WorldErrorsCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WorldErrors.Contract.Eip712Domain(&_WorldErrors.CallOpts)
}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_WorldErrors *WorldErrorsCaller) GetDynamicField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getDynamicField", tableId, keyTuple, dynamicFieldIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_WorldErrors *WorldErrorsSession) GetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	return _WorldErrors.Contract.GetDynamicField(&_WorldErrors.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicField is a free data retrieval call binding the contract method 0x1e788977.
//
// Solidity: function getDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(bytes)
func (_WorldErrors *WorldErrorsCallerSession) GetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) ([]byte, error) {
	return _WorldErrors.Contract.GetDynamicField(&_WorldErrors.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_WorldErrors *WorldErrorsCaller) GetDynamicFieldLength(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getDynamicFieldLength", tableId, keyTuple, dynamicFieldIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_WorldErrors *WorldErrorsSession) GetDynamicFieldLength(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	return _WorldErrors.Contract.GetDynamicFieldLength(&_WorldErrors.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldLength is a free data retrieval call binding the contract method 0xdbbf0e21.
//
// Solidity: function getDynamicFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex) view returns(uint256)
func (_WorldErrors *WorldErrorsCallerSession) GetDynamicFieldLength(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8) (*big.Int, error) {
	return _WorldErrors.Contract.GetDynamicFieldLength(&_WorldErrors.CallOpts, tableId, keyTuple, dynamicFieldIndex)
}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_WorldErrors *WorldErrorsCaller) GetDynamicFieldSlice(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getDynamicFieldSlice", tableId, keyTuple, dynamicFieldIndex, start, end)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_WorldErrors *WorldErrorsSession) GetDynamicFieldSlice(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	return _WorldErrors.Contract.GetDynamicFieldSlice(&_WorldErrors.CallOpts, tableId, keyTuple, dynamicFieldIndex, start, end)
}

// GetDynamicFieldSlice is a free data retrieval call binding the contract method 0x4dc77d97.
//
// Solidity: function getDynamicFieldSlice(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 start, uint256 end) view returns(bytes data)
func (_WorldErrors *WorldErrorsCallerSession) GetDynamicFieldSlice(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, start *big.Int, end *big.Int) ([]byte, error) {
	return _WorldErrors.Contract.GetDynamicFieldSlice(&_WorldErrors.CallOpts, tableId, keyTuple, dynamicFieldIndex, start, end)
}

// GetField is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_WorldErrors *WorldErrorsCaller) GetField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getField", tableId, keyTuple, fieldIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetField is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_WorldErrors *WorldErrorsSession) GetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	return _WorldErrors.Contract.GetField(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetField is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(bytes data)
func (_WorldErrors *WorldErrorsCallerSession) GetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) ([]byte, error) {
	return _WorldErrors.Contract.GetField(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetField0 is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_WorldErrors *WorldErrorsCaller) GetField0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getField0", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetField0 is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_WorldErrors *WorldErrorsSession) GetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	return _WorldErrors.Contract.GetField0(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetField0 is a free data retrieval call binding the contract method 0x05242d2f.
//
// Solidity: function getField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes data)
func (_WorldErrors *WorldErrorsCallerSession) GetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([]byte, error) {
	return _WorldErrors.Contract.GetField0(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_WorldErrors *WorldErrorsCaller) GetFieldLayout(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getFieldLayout", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_WorldErrors *WorldErrorsSession) GetFieldLayout(tableId [32]byte) ([32]byte, error) {
	return _WorldErrors.Contract.GetFieldLayout(&_WorldErrors.CallOpts, tableId)
}

// GetFieldLayout is a free data retrieval call binding the contract method 0x3a77c2c2.
//
// Solidity: function getFieldLayout(bytes32 tableId) view returns(bytes32 fieldLayout)
func (_WorldErrors *WorldErrorsCallerSession) GetFieldLayout(tableId [32]byte) ([32]byte, error) {
	return _WorldErrors.Contract.GetFieldLayout(&_WorldErrors.CallOpts, tableId)
}

// GetFieldLength is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_WorldErrors *WorldErrorsCaller) GetFieldLength(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getFieldLength", tableId, keyTuple, fieldIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldLength is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_WorldErrors *WorldErrorsSession) GetFieldLength(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	return _WorldErrors.Contract.GetFieldLength(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetFieldLength is a free data retrieval call binding the contract method 0xa53417ed.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex) view returns(uint256)
func (_WorldErrors *WorldErrorsCallerSession) GetFieldLength(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8) (*big.Int, error) {
	return _WorldErrors.Contract.GetFieldLength(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex)
}

// GetFieldLength0 is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_WorldErrors *WorldErrorsCaller) GetFieldLength0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getFieldLength0", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldLength0 is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_WorldErrors *WorldErrorsSession) GetFieldLength0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	return _WorldErrors.Contract.GetFieldLength0(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetFieldLength0 is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(uint256)
func (_WorldErrors *WorldErrorsCallerSession) GetFieldLength0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) (*big.Int, error) {
	return _WorldErrors.Contract.GetFieldLength0(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_WorldErrors *WorldErrorsCaller) GetKeySchema(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getKeySchema", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_WorldErrors *WorldErrorsSession) GetKeySchema(tableId [32]byte) ([32]byte, error) {
	return _WorldErrors.Contract.GetKeySchema(&_WorldErrors.CallOpts, tableId)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 tableId) view returns(bytes32 keySchema)
func (_WorldErrors *WorldErrorsCallerSession) GetKeySchema(tableId [32]byte) ([32]byte, error) {
	return _WorldErrors.Contract.GetKeySchema(&_WorldErrors.CallOpts, tableId)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() pure returns(bytes16)
func (_WorldErrors *WorldErrorsCaller) GetName(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() pure returns(bytes16)
func (_WorldErrors *WorldErrorsSession) GetName() ([16]byte, error) {
	return _WorldErrors.Contract.GetName(&_WorldErrors.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() pure returns(bytes16)
func (_WorldErrors *WorldErrorsCallerSession) GetName() ([16]byte, error) {
	return _WorldErrors.Contract.GetName(&_WorldErrors.CallOpts)
}

// GetRecord is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_WorldErrors *WorldErrorsCaller) GetRecord(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getRecord", tableId, keyTuple)

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

// GetRecord is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_WorldErrors *WorldErrorsSession) GetRecord(tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _WorldErrors.Contract.GetRecord(&_WorldErrors.CallOpts, tableId, keyTuple)
}

// GetRecord is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_WorldErrors *WorldErrorsCallerSession) GetRecord(tableId [32]byte, keyTuple [][32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _WorldErrors.Contract.GetRecord(&_WorldErrors.CallOpts, tableId, keyTuple)
}

// GetRecord0 is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_WorldErrors *WorldErrorsCaller) GetRecord0(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getRecord0", tableId, keyTuple, fieldLayout)

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

// GetRecord0 is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_WorldErrors *WorldErrorsSession) GetRecord0(tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _WorldErrors.Contract.GetRecord0(&_WorldErrors.CallOpts, tableId, keyTuple, fieldLayout)
}

// GetRecord0 is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 tableId, bytes32[] keyTuple, bytes32 fieldLayout) view returns(bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_WorldErrors *WorldErrorsCallerSession) GetRecord0(tableId [32]byte, keyTuple [][32]byte, fieldLayout [32]byte) (struct {
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
}, error) {
	return _WorldErrors.Contract.GetRecord0(&_WorldErrors.CallOpts, tableId, keyTuple, fieldLayout)
}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_WorldErrors *WorldErrorsCaller) GetStaticField(opts *bind.CallOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getStaticField", tableId, keyTuple, fieldIndex, fieldLayout)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_WorldErrors *WorldErrorsSession) GetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	return _WorldErrors.Contract.GetStaticField(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetStaticField is a free data retrieval call binding the contract method 0x8c364d59.
//
// Solidity: function getStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes32 fieldLayout) view returns(bytes32)
func (_WorldErrors *WorldErrorsCallerSession) GetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, fieldLayout [32]byte) ([32]byte, error) {
	return _WorldErrors.Contract.GetStaticField(&_WorldErrors.CallOpts, tableId, keyTuple, fieldIndex, fieldLayout)
}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_WorldErrors *WorldErrorsCaller) GetValueSchema(opts *bind.CallOpts, tableId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "getValueSchema", tableId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_WorldErrors *WorldErrorsSession) GetValueSchema(tableId [32]byte) ([32]byte, error) {
	return _WorldErrors.Contract.GetValueSchema(&_WorldErrors.CallOpts, tableId)
}

// GetValueSchema is a free data retrieval call binding the contract method 0xe228a4a3.
//
// Solidity: function getValueSchema(bytes32 tableId) view returns(bytes32 valueSchema)
func (_WorldErrors *WorldErrorsCallerSession) GetValueSchema(tableId [32]byte) ([32]byte, error) {
	return _WorldErrors.Contract.GetValueSchema(&_WorldErrors.CallOpts, tableId)
}

// InitialMsgSender is a free data retrieval call binding the contract method 0x1e71d54c.
//
// Solidity: function initialMsgSender() view returns(address)
func (_WorldErrors *WorldErrorsCaller) InitialMsgSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "initialMsgSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InitialMsgSender is a free data retrieval call binding the contract method 0x1e71d54c.
//
// Solidity: function initialMsgSender() view returns(address)
func (_WorldErrors *WorldErrorsSession) InitialMsgSender() (common.Address, error) {
	return _WorldErrors.Contract.InitialMsgSender(&_WorldErrors.CallOpts)
}

// InitialMsgSender is a free data retrieval call binding the contract method 0x1e71d54c.
//
// Solidity: function initialMsgSender() view returns(address)
func (_WorldErrors *WorldErrorsCallerSession) InitialMsgSender() (common.Address, error) {
	return _WorldErrors.Contract.InitialMsgSender(&_WorldErrors.CallOpts)
}

// InstallRoot is a free data retrieval call binding the contract method 0x7c5b0335.
//
// Solidity: function installRoot(bytes ) pure returns()
func (_WorldErrors *WorldErrorsCaller) InstallRoot(opts *bind.CallOpts, arg0 []byte) error {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "installRoot", arg0)

	if err != nil {
		return err
	}

	return err

}

// InstallRoot is a free data retrieval call binding the contract method 0x7c5b0335.
//
// Solidity: function installRoot(bytes ) pure returns()
func (_WorldErrors *WorldErrorsSession) InstallRoot(arg0 []byte) error {
	return _WorldErrors.Contract.InstallRoot(&_WorldErrors.CallOpts, arg0)
}

// InstallRoot is a free data retrieval call binding the contract method 0x7c5b0335.
//
// Solidity: function installRoot(bytes ) pure returns()
func (_WorldErrors *WorldErrorsCallerSession) InstallRoot(arg0 []byte) error {
	return _WorldErrors.Contract.InstallRoot(&_WorldErrors.CallOpts, arg0)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0xcab7e8eb.
//
// Solidity: function isNonceUsed(address owner, uint256 data) view returns(bool)
func (_WorldErrors *WorldErrorsCaller) IsNonceUsed(opts *bind.CallOpts, owner common.Address, data *big.Int) (bool, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "isNonceUsed", owner, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonceUsed is a free data retrieval call binding the contract method 0xcab7e8eb.
//
// Solidity: function isNonceUsed(address owner, uint256 data) view returns(bool)
func (_WorldErrors *WorldErrorsSession) IsNonceUsed(owner common.Address, data *big.Int) (bool, error) {
	return _WorldErrors.Contract.IsNonceUsed(&_WorldErrors.CallOpts, owner, data)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0xcab7e8eb.
//
// Solidity: function isNonceUsed(address owner, uint256 data) view returns(bool)
func (_WorldErrors *WorldErrorsCallerSession) IsNonceUsed(owner common.Address, data *big.Int) (bool, error) {
	return _WorldErrors.Contract.IsNonceUsed(&_WorldErrors.CallOpts, owner, data)
}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_WorldErrors *WorldErrorsCaller) StoreVersion(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "storeVersion")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_WorldErrors *WorldErrorsSession) StoreVersion() ([32]byte, error) {
	return _WorldErrors.Contract.StoreVersion(&_WorldErrors.CallOpts)
}

// StoreVersion is a free data retrieval call binding the contract method 0xc1122229.
//
// Solidity: function storeVersion() view returns(bytes32 version)
func (_WorldErrors *WorldErrorsCallerSession) StoreVersion() ([32]byte, error) {
	return _WorldErrors.Contract.StoreVersion(&_WorldErrors.CallOpts)
}

// StructHash is a free data retrieval call binding the contract method 0xf1109062.
//
// Solidity: function structHash((address,address,uint256,uint256,uint256,uint48,bytes) request) view returns(bytes32)
func (_WorldErrors *WorldErrorsCaller) StructHash(opts *bind.CallOpts, request ERC2771ForwarderForwardRequest) ([32]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "structHash", request)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StructHash is a free data retrieval call binding the contract method 0xf1109062.
//
// Solidity: function structHash((address,address,uint256,uint256,uint256,uint48,bytes) request) view returns(bytes32)
func (_WorldErrors *WorldErrorsSession) StructHash(request ERC2771ForwarderForwardRequest) ([32]byte, error) {
	return _WorldErrors.Contract.StructHash(&_WorldErrors.CallOpts, request)
}

// StructHash is a free data retrieval call binding the contract method 0xf1109062.
//
// Solidity: function structHash((address,address,uint256,uint256,uint256,uint48,bytes) request) view returns(bytes32)
func (_WorldErrors *WorldErrorsCallerSession) StructHash(request ERC2771ForwarderForwardRequest) ([32]byte, error) {
	return _WorldErrors.Contract.StructHash(&_WorldErrors.CallOpts, request)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_WorldErrors *WorldErrorsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_WorldErrors *WorldErrorsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WorldErrors.Contract.SupportsInterface(&_WorldErrors.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_WorldErrors *WorldErrorsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WorldErrors.Contract.SupportsInterface(&_WorldErrors.CallOpts, interfaceId)
}

// Verify is a free data retrieval call binding the contract method 0x1a86b550.
//
// Solidity: function verify(address , bytes32 systemId, bytes ) view returns(bool)
func (_WorldErrors *WorldErrorsCaller) Verify(opts *bind.CallOpts, arg0 common.Address, systemId [32]byte, arg2 []byte) (bool, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "verify", arg0, systemId, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x1a86b550.
//
// Solidity: function verify(address , bytes32 systemId, bytes ) view returns(bool)
func (_WorldErrors *WorldErrorsSession) Verify(arg0 common.Address, systemId [32]byte, arg2 []byte) (bool, error) {
	return _WorldErrors.Contract.Verify(&_WorldErrors.CallOpts, arg0, systemId, arg2)
}

// Verify is a free data retrieval call binding the contract method 0x1a86b550.
//
// Solidity: function verify(address , bytes32 systemId, bytes ) view returns(bool)
func (_WorldErrors *WorldErrorsCallerSession) Verify(arg0 common.Address, systemId [32]byte, arg2 []byte) (bool, error) {
	return _WorldErrors.Contract.Verify(&_WorldErrors.CallOpts, arg0, systemId, arg2)
}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_WorldErrors *WorldErrorsCaller) WorldVersion(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WorldErrors.contract.Call(opts, &out, "worldVersion")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_WorldErrors *WorldErrorsSession) WorldVersion() ([32]byte, error) {
	return _WorldErrors.Contract.WorldVersion(&_WorldErrors.CallOpts)
}

// WorldVersion is a free data retrieval call binding the contract method 0x6951955d.
//
// Solidity: function worldVersion() view returns(bytes32)
func (_WorldErrors *WorldErrorsCallerSession) WorldVersion() ([32]byte, error) {
	return _WorldErrors.Contract.WorldVersion(&_WorldErrors.CallOpts)
}

// AddHook is a paid mutator transaction binding the contract method 0xe3c84776.
//
// Solidity: function addHook(uint256 hookId, uint8 hookType, bytes32 systemId, bytes4 functionSelector) returns()
func (_WorldErrors *WorldErrorsTransactor) AddHook(opts *bind.TransactOpts, hookId *big.Int, hookType uint8, systemId [32]byte, functionSelector [4]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "addHook", hookId, hookType, systemId, functionSelector)
}

// AddHook is a paid mutator transaction binding the contract method 0xe3c84776.
//
// Solidity: function addHook(uint256 hookId, uint8 hookType, bytes32 systemId, bytes4 functionSelector) returns()
func (_WorldErrors *WorldErrorsSession) AddHook(hookId *big.Int, hookType uint8, systemId [32]byte, functionSelector [4]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.AddHook(&_WorldErrors.TransactOpts, hookId, hookType, systemId, functionSelector)
}

// AddHook is a paid mutator transaction binding the contract method 0xe3c84776.
//
// Solidity: function addHook(uint256 hookId, uint8 hookType, bytes32 systemId, bytes4 functionSelector) returns()
func (_WorldErrors *WorldErrorsTransactorSession) AddHook(hookId *big.Int, hookType uint8, systemId [32]byte, functionSelector [4]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.AddHook(&_WorldErrors.TransactOpts, hookId, hookType, systemId, functionSelector)
}

// Anchor is a paid mutator transaction binding the contract method 0xc0395eab.
//
// Solidity: function anchor(uint256 entityId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_WorldErrors *WorldErrorsTransactor) Anchor(opts *bind.TransactOpts, entityId *big.Int, locationData LocationTableData) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "anchor", entityId, locationData)
}

// Anchor is a paid mutator transaction binding the contract method 0xc0395eab.
//
// Solidity: function anchor(uint256 entityId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_WorldErrors *WorldErrorsSession) Anchor(entityId *big.Int, locationData LocationTableData) (*types.Transaction, error) {
	return _WorldErrors.Contract.Anchor(&_WorldErrors.TransactOpts, entityId, locationData)
}

// Anchor is a paid mutator transaction binding the contract method 0xc0395eab.
//
// Solidity: function anchor(uint256 entityId, (uint256,uint256,uint256,uint256) locationData) returns()
func (_WorldErrors *WorldErrorsTransactorSession) Anchor(entityId *big.Int, locationData LocationTableData) (*types.Transaction, error) {
	return _WorldErrors.Contract.Anchor(&_WorldErrors.TransactOpts, entityId, locationData)
}

// AssociateHook is a paid mutator transaction binding the contract method 0x6092ee5a.
//
// Solidity: function associateHook(uint256 entityId, uint256 hookId) returns()
func (_WorldErrors *WorldErrorsTransactor) AssociateHook(opts *bind.TransactOpts, entityId *big.Int, hookId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "associateHook", entityId, hookId)
}

// AssociateHook is a paid mutator transaction binding the contract method 0x6092ee5a.
//
// Solidity: function associateHook(uint256 entityId, uint256 hookId) returns()
func (_WorldErrors *WorldErrorsSession) AssociateHook(entityId *big.Int, hookId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.AssociateHook(&_WorldErrors.TransactOpts, entityId, hookId)
}

// AssociateHook is a paid mutator transaction binding the contract method 0x6092ee5a.
//
// Solidity: function associateHook(uint256 entityId, uint256 hookId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) AssociateHook(entityId *big.Int, hookId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.AssociateHook(&_WorldErrors.TransactOpts, entityId, hookId)
}

// AssociateHooks is a paid mutator transaction binding the contract method 0x4938b5ce.
//
// Solidity: function associateHooks(uint256 entityId, uint256[] hookIds) returns()
func (_WorldErrors *WorldErrorsTransactor) AssociateHooks(opts *bind.TransactOpts, entityId *big.Int, hookIds []*big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "associateHooks", entityId, hookIds)
}

// AssociateHooks is a paid mutator transaction binding the contract method 0x4938b5ce.
//
// Solidity: function associateHooks(uint256 entityId, uint256[] hookIds) returns()
func (_WorldErrors *WorldErrorsSession) AssociateHooks(entityId *big.Int, hookIds []*big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.AssociateHooks(&_WorldErrors.TransactOpts, entityId, hookIds)
}

// AssociateHooks is a paid mutator transaction binding the contract method 0x4938b5ce.
//
// Solidity: function associateHooks(uint256 entityId, uint256[] hookIds) returns()
func (_WorldErrors *WorldErrorsTransactorSession) AssociateHooks(entityId *big.Int, hookIds []*big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.AssociateHooks(&_WorldErrors.TransactOpts, entityId, hookIds)
}

// AssociateModule is a paid mutator transaction binding the contract method 0x4882da35.
//
// Solidity: function associateModule(uint256 entityId, uint256 moduleId) returns()
func (_WorldErrors *WorldErrorsTransactor) AssociateModule(opts *bind.TransactOpts, entityId *big.Int, moduleId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "associateModule", entityId, moduleId)
}

// AssociateModule is a paid mutator transaction binding the contract method 0x4882da35.
//
// Solidity: function associateModule(uint256 entityId, uint256 moduleId) returns()
func (_WorldErrors *WorldErrorsSession) AssociateModule(entityId *big.Int, moduleId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.AssociateModule(&_WorldErrors.TransactOpts, entityId, moduleId)
}

// AssociateModule is a paid mutator transaction binding the contract method 0x4882da35.
//
// Solidity: function associateModule(uint256 entityId, uint256 moduleId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) AssociateModule(entityId *big.Int, moduleId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.AssociateModule(&_WorldErrors.TransactOpts, entityId, moduleId)
}

// AssociateModules is a paid mutator transaction binding the contract method 0xc43c7992.
//
// Solidity: function associateModules(uint256 entityId, uint256[] moduleIds) returns()
func (_WorldErrors *WorldErrorsTransactor) AssociateModules(opts *bind.TransactOpts, entityId *big.Int, moduleIds []*big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "associateModules", entityId, moduleIds)
}

// AssociateModules is a paid mutator transaction binding the contract method 0xc43c7992.
//
// Solidity: function associateModules(uint256 entityId, uint256[] moduleIds) returns()
func (_WorldErrors *WorldErrorsSession) AssociateModules(entityId *big.Int, moduleIds []*big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.AssociateModules(&_WorldErrors.TransactOpts, entityId, moduleIds)
}

// AssociateModules is a paid mutator transaction binding the contract method 0xc43c7992.
//
// Solidity: function associateModules(uint256 entityId, uint256[] moduleIds) returns()
func (_WorldErrors *WorldErrorsTransactorSession) AssociateModules(entityId *big.Int, moduleIds []*big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.AssociateModules(&_WorldErrors.TransactOpts, entityId, moduleIds)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_WorldErrors *WorldErrorsTransactor) BatchCall(opts *bind.TransactOpts, systemCalls []SystemCallData) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "batchCall", systemCalls)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_WorldErrors *WorldErrorsSession) BatchCall(systemCalls []SystemCallData) (*types.Transaction, error) {
	return _WorldErrors.Contract.BatchCall(&_WorldErrors.TransactOpts, systemCalls)
}

// BatchCall is a paid mutator transaction binding the contract method 0xce5e8dd9.
//
// Solidity: function batchCall((bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_WorldErrors *WorldErrorsTransactorSession) BatchCall(systemCalls []SystemCallData) (*types.Transaction, error) {
	return _WorldErrors.Contract.BatchCall(&_WorldErrors.TransactOpts, systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_WorldErrors *WorldErrorsTransactor) BatchCallFrom(opts *bind.TransactOpts, systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "batchCallFrom", systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_WorldErrors *WorldErrorsSession) BatchCallFrom(systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _WorldErrors.Contract.BatchCallFrom(&_WorldErrors.TransactOpts, systemCalls)
}

// BatchCallFrom is a paid mutator transaction binding the contract method 0x8fc8cf7e.
//
// Solidity: function batchCallFrom((address,bytes32,bytes)[] systemCalls) returns(bytes[] returnDatas)
func (_WorldErrors *WorldErrorsTransactorSession) BatchCallFrom(systemCalls []SystemCallFromData) (*types.Transaction, error) {
	return _WorldErrors.Contract.BatchCallFrom(&_WorldErrors.TransactOpts, systemCalls)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactor) BringOffline(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "bringOffline", entityId)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsSession) BringOffline(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.BringOffline(&_WorldErrors.TransactOpts, entityId)
}

// BringOffline is a paid mutator transaction binding the contract method 0xe1509e87.
//
// Solidity: function bringOffline(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) BringOffline(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.BringOffline(&_WorldErrors.TransactOpts, entityId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactor) BringOnline(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "bringOnline", entityId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsSession) BringOnline(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.BringOnline(&_WorldErrors.TransactOpts, entityId)
}

// BringOnline is a paid mutator transaction binding the contract method 0xeb5f2f58.
//
// Solidity: function bringOnline(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) BringOnline(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.BringOnline(&_WorldErrors.TransactOpts, entityId)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_WorldErrors *WorldErrorsTransactor) Call(opts *bind.TransactOpts, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "call", systemId, callData)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_WorldErrors *WorldErrorsSession) Call(systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.Call(&_WorldErrors.TransactOpts, systemId, callData)
}

// Call is a paid mutator transaction binding the contract method 0x3ae7af08.
//
// Solidity: function call(bytes32 systemId, bytes callData) payable returns(bytes)
func (_WorldErrors *WorldErrorsTransactorSession) Call(systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.Call(&_WorldErrors.TransactOpts, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_WorldErrors *WorldErrorsTransactor) CallFrom(opts *bind.TransactOpts, delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "callFrom", delegator, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_WorldErrors *WorldErrorsSession) CallFrom(delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.CallFrom(&_WorldErrors.TransactOpts, delegator, systemId, callData)
}

// CallFrom is a paid mutator transaction binding the contract method 0x894ecc58.
//
// Solidity: function callFrom(address delegator, bytes32 systemId, bytes callData) payable returns(bytes)
func (_WorldErrors *WorldErrorsTransactorSession) CallFrom(delegator common.Address, systemId [32]byte, callData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.CallFrom(&_WorldErrors.TransactOpts, delegator, systemId, callData)
}

// ConfigureInteractionHandler is a paid mutator transaction binding the contract method 0xb81db02a.
//
// Solidity: function configureInteractionHandler(uint256 smartObjectId, bytes interactionParams) returns()
func (_WorldErrors *WorldErrorsTransactor) ConfigureInteractionHandler(opts *bind.TransactOpts, smartObjectId *big.Int, interactionParams []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "configureInteractionHandler", smartObjectId, interactionParams)
}

// ConfigureInteractionHandler is a paid mutator transaction binding the contract method 0xb81db02a.
//
// Solidity: function configureInteractionHandler(uint256 smartObjectId, bytes interactionParams) returns()
func (_WorldErrors *WorldErrorsSession) ConfigureInteractionHandler(smartObjectId *big.Int, interactionParams []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.ConfigureInteractionHandler(&_WorldErrors.TransactOpts, smartObjectId, interactionParams)
}

// ConfigureInteractionHandler is a paid mutator transaction binding the contract method 0xb81db02a.
//
// Solidity: function configureInteractionHandler(uint256 smartObjectId, bytes interactionParams) returns()
func (_WorldErrors *WorldErrorsTransactorSession) ConfigureInteractionHandler(smartObjectId *big.Int, interactionParams []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.ConfigureInteractionHandler(&_WorldErrors.TransactOpts, smartObjectId, interactionParams)
}

// CreateAndAnchorSmartStorageUnit is a paid mutator transaction binding the contract method 0x1bd1eb3e.
//
// Solidity: function createAndAnchorSmartStorageUnit(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionPerMinute, uint256 fuelMaxCapacity, uint256 storageCapacity, uint256 ephemeralStorageCapacity) returns()
func (_WorldErrors *WorldErrorsTransactor) CreateAndAnchorSmartStorageUnit(opts *bind.TransactOpts, smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionPerMinute *big.Int, fuelMaxCapacity *big.Int, storageCapacity *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "createAndAnchorSmartStorageUnit", smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionPerMinute, fuelMaxCapacity, storageCapacity, ephemeralStorageCapacity)
}

// CreateAndAnchorSmartStorageUnit is a paid mutator transaction binding the contract method 0x1bd1eb3e.
//
// Solidity: function createAndAnchorSmartStorageUnit(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionPerMinute, uint256 fuelMaxCapacity, uint256 storageCapacity, uint256 ephemeralStorageCapacity) returns()
func (_WorldErrors *WorldErrorsSession) CreateAndAnchorSmartStorageUnit(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionPerMinute *big.Int, fuelMaxCapacity *big.Int, storageCapacity *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateAndAnchorSmartStorageUnit(&_WorldErrors.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionPerMinute, fuelMaxCapacity, storageCapacity, ephemeralStorageCapacity)
}

// CreateAndAnchorSmartStorageUnit is a paid mutator transaction binding the contract method 0x1bd1eb3e.
//
// Solidity: function createAndAnchorSmartStorageUnit(uint256 smartObjectId, (uint256,uint256,uint256) entityRecordData, (address,string) smartObjectData, (uint256,(uint256,uint256,uint256)) worldPosition, uint256 fuelUnitVolume, uint256 fuelConsumptionPerMinute, uint256 fuelMaxCapacity, uint256 storageCapacity, uint256 ephemeralStorageCapacity) returns()
func (_WorldErrors *WorldErrorsTransactorSession) CreateAndAnchorSmartStorageUnit(smartObjectId *big.Int, entityRecordData EntityRecordData, smartObjectData SmartObjectData, worldPosition WorldPosition, fuelUnitVolume *big.Int, fuelConsumptionPerMinute *big.Int, fuelMaxCapacity *big.Int, storageCapacity *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateAndAnchorSmartStorageUnit(&_WorldErrors.TransactOpts, smartObjectId, entityRecordData, smartObjectData, worldPosition, fuelUnitVolume, fuelConsumptionPerMinute, fuelMaxCapacity, storageCapacity, ephemeralStorageCapacity)
}

// CreateAndDepositItemsToEphemeralInventory is a paid mutator transaction binding the contract method 0xe111c1a6.
//
// Solidity: function createAndDepositItemsToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactor) CreateAndDepositItemsToEphemeralInventory(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "createAndDepositItemsToEphemeralInventory", smartObjectId, ephemeralInventoryOwner, items)
}

// CreateAndDepositItemsToEphemeralInventory is a paid mutator transaction binding the contract method 0xe111c1a6.
//
// Solidity: function createAndDepositItemsToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsSession) CreateAndDepositItemsToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateAndDepositItemsToEphemeralInventory(&_WorldErrors.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// CreateAndDepositItemsToEphemeralInventory is a paid mutator transaction binding the contract method 0xe111c1a6.
//
// Solidity: function createAndDepositItemsToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactorSession) CreateAndDepositItemsToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateAndDepositItemsToEphemeralInventory(&_WorldErrors.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// CreateAndDepositItemsToInventory is a paid mutator transaction binding the contract method 0x9e19e7aa.
//
// Solidity: function createAndDepositItemsToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactor) CreateAndDepositItemsToInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "createAndDepositItemsToInventory", smartObjectId, items)
}

// CreateAndDepositItemsToInventory is a paid mutator transaction binding the contract method 0x9e19e7aa.
//
// Solidity: function createAndDepositItemsToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsSession) CreateAndDepositItemsToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateAndDepositItemsToInventory(&_WorldErrors.TransactOpts, smartObjectId, items)
}

// CreateAndDepositItemsToInventory is a paid mutator transaction binding the contract method 0x9e19e7aa.
//
// Solidity: function createAndDepositItemsToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactorSession) CreateAndDepositItemsToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateAndDepositItemsToInventory(&_WorldErrors.TransactOpts, smartObjectId, items)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x777dd579.
//
// Solidity: function createCharacter(uint256 characterId, address characterAddress, (uint256,uint256,uint256) entityRecord, (string,string,string) entityRecordOffchain, string tokenCid) returns()
func (_WorldErrors *WorldErrorsTransactor) CreateCharacter(opts *bind.TransactOpts, characterId *big.Int, characterAddress common.Address, entityRecord EntityRecordData, entityRecordOffchain EntityRecordOffchainTableData, tokenCid string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "createCharacter", characterId, characterAddress, entityRecord, entityRecordOffchain, tokenCid)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x777dd579.
//
// Solidity: function createCharacter(uint256 characterId, address characterAddress, (uint256,uint256,uint256) entityRecord, (string,string,string) entityRecordOffchain, string tokenCid) returns()
func (_WorldErrors *WorldErrorsSession) CreateCharacter(characterId *big.Int, characterAddress common.Address, entityRecord EntityRecordData, entityRecordOffchain EntityRecordOffchainTableData, tokenCid string) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateCharacter(&_WorldErrors.TransactOpts, characterId, characterAddress, entityRecord, entityRecordOffchain, tokenCid)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0x777dd579.
//
// Solidity: function createCharacter(uint256 characterId, address characterAddress, (uint256,uint256,uint256) entityRecord, (string,string,string) entityRecordOffchain, string tokenCid) returns()
func (_WorldErrors *WorldErrorsTransactorSession) CreateCharacter(characterId *big.Int, characterAddress common.Address, entityRecord EntityRecordData, entityRecordOffchain EntityRecordOffchainTableData, tokenCid string) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateCharacter(&_WorldErrors.TransactOpts, characterId, characterAddress, entityRecord, entityRecordOffchain, tokenCid)
}

// CreateEntityRecord is a paid mutator transaction binding the contract method 0x2c63f58f.
//
// Solidity: function createEntityRecord(uint256 entityId, uint256 itemId, uint256 typeId, uint256 volume) returns()
func (_WorldErrors *WorldErrorsTransactor) CreateEntityRecord(opts *bind.TransactOpts, entityId *big.Int, itemId *big.Int, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "createEntityRecord", entityId, itemId, typeId, volume)
}

// CreateEntityRecord is a paid mutator transaction binding the contract method 0x2c63f58f.
//
// Solidity: function createEntityRecord(uint256 entityId, uint256 itemId, uint256 typeId, uint256 volume) returns()
func (_WorldErrors *WorldErrorsSession) CreateEntityRecord(entityId *big.Int, itemId *big.Int, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateEntityRecord(&_WorldErrors.TransactOpts, entityId, itemId, typeId, volume)
}

// CreateEntityRecord is a paid mutator transaction binding the contract method 0x2c63f58f.
//
// Solidity: function createEntityRecord(uint256 entityId, uint256 itemId, uint256 typeId, uint256 volume) returns()
func (_WorldErrors *WorldErrorsTransactorSession) CreateEntityRecord(entityId *big.Int, itemId *big.Int, typeId *big.Int, volume *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateEntityRecord(&_WorldErrors.TransactOpts, entityId, itemId, typeId, volume)
}

// CreateEntityRecordOffchain is a paid mutator transaction binding the contract method 0xbf662710.
//
// Solidity: function createEntityRecordOffchain(uint256 entityId, string name, string dappURL, string description) returns()
func (_WorldErrors *WorldErrorsTransactor) CreateEntityRecordOffchain(opts *bind.TransactOpts, entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "createEntityRecordOffchain", entityId, name, dappURL, description)
}

// CreateEntityRecordOffchain is a paid mutator transaction binding the contract method 0xbf662710.
//
// Solidity: function createEntityRecordOffchain(uint256 entityId, string name, string dappURL, string description) returns()
func (_WorldErrors *WorldErrorsSession) CreateEntityRecordOffchain(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateEntityRecordOffchain(&_WorldErrors.TransactOpts, entityId, name, dappURL, description)
}

// CreateEntityRecordOffchain is a paid mutator transaction binding the contract method 0xbf662710.
//
// Solidity: function createEntityRecordOffchain(uint256 entityId, string name, string dappURL, string description) returns()
func (_WorldErrors *WorldErrorsTransactorSession) CreateEntityRecordOffchain(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _WorldErrors.Contract.CreateEntityRecordOffchain(&_WorldErrors.TransactOpts, entityId, name, dappURL, description)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_WorldErrors *WorldErrorsTransactor) DeleteRecord(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "deleteRecord", tableId, keyTuple)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_WorldErrors *WorldErrorsSession) DeleteRecord(tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.DeleteRecord(&_WorldErrors.TransactOpts, tableId, keyTuple)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 tableId, bytes32[] keyTuple) returns()
func (_WorldErrors *WorldErrorsTransactorSession) DeleteRecord(tableId [32]byte, keyTuple [][32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.DeleteRecord(&_WorldErrors.TransactOpts, tableId, keyTuple)
}

// DepositFuel is a paid mutator transaction binding the contract method 0xc3e9a45f.
//
// Solidity: function depositFuel(uint256 entityId, uint256 unitAmount) returns()
func (_WorldErrors *WorldErrorsTransactor) DepositFuel(opts *bind.TransactOpts, entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "depositFuel", entityId, unitAmount)
}

// DepositFuel is a paid mutator transaction binding the contract method 0xc3e9a45f.
//
// Solidity: function depositFuel(uint256 entityId, uint256 unitAmount) returns()
func (_WorldErrors *WorldErrorsSession) DepositFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.DepositFuel(&_WorldErrors.TransactOpts, entityId, unitAmount)
}

// DepositFuel is a paid mutator transaction binding the contract method 0xc3e9a45f.
//
// Solidity: function depositFuel(uint256 entityId, uint256 unitAmount) returns()
func (_WorldErrors *WorldErrorsTransactorSession) DepositFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.DepositFuel(&_WorldErrors.TransactOpts, entityId, unitAmount)
}

// DepositToEphemeralInventory is a paid mutator transaction binding the contract method 0xaff42af2.
//
// Solidity: function depositToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactor) DepositToEphemeralInventory(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "depositToEphemeralInventory", smartObjectId, ephemeralInventoryOwner, items)
}

// DepositToEphemeralInventory is a paid mutator transaction binding the contract method 0xaff42af2.
//
// Solidity: function depositToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsSession) DepositToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.DepositToEphemeralInventory(&_WorldErrors.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// DepositToEphemeralInventory is a paid mutator transaction binding the contract method 0xaff42af2.
//
// Solidity: function depositToEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactorSession) DepositToEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.DepositToEphemeralInventory(&_WorldErrors.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// DepositToInventory is a paid mutator transaction binding the contract method 0x15e306bb.
//
// Solidity: function depositToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactor) DepositToInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "depositToInventory", smartObjectId, items)
}

// DepositToInventory is a paid mutator transaction binding the contract method 0x15e306bb.
//
// Solidity: function depositToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsSession) DepositToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.DepositToInventory(&_WorldErrors.TransactOpts, smartObjectId, items)
}

// DepositToInventory is a paid mutator transaction binding the contract method 0x15e306bb.
//
// Solidity: function depositToInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactorSession) DepositToInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.DepositToInventory(&_WorldErrors.TransactOpts, smartObjectId, items)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactor) DestroyDeployable(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "destroyDeployable", entityId)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsSession) DestroyDeployable(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.DestroyDeployable(&_WorldErrors.TransactOpts, entityId)
}

// DestroyDeployable is a paid mutator transaction binding the contract method 0x80a63ec4.
//
// Solidity: function destroyDeployable(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) DestroyDeployable(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.DestroyDeployable(&_WorldErrors.TransactOpts, entityId)
}

// EphemeralToInventoryTransfer is a paid mutator transaction binding the contract method 0x6f89f6c3.
//
// Solidity: function ephemeralToInventoryTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactor) EphemeralToInventoryTransfer(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "ephemeralToInventoryTransfer", smartObjectId, items)
}

// EphemeralToInventoryTransfer is a paid mutator transaction binding the contract method 0x6f89f6c3.
//
// Solidity: function ephemeralToInventoryTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsSession) EphemeralToInventoryTransfer(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.EphemeralToInventoryTransfer(&_WorldErrors.TransactOpts, smartObjectId, items)
}

// EphemeralToInventoryTransfer is a paid mutator transaction binding the contract method 0x6f89f6c3.
//
// Solidity: function ephemeralToInventoryTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactorSession) EphemeralToInventoryTransfer(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.EphemeralToInventoryTransfer(&_WorldErrors.TransactOpts, smartObjectId, items)
}

// Execute is a paid mutator transaction binding the contract method 0xccff1ca1.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) payable returns()
func (_WorldErrors *WorldErrorsTransactor) Execute(opts *bind.TransactOpts, request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "execute", request)
}

// Execute is a paid mutator transaction binding the contract method 0xccff1ca1.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) payable returns()
func (_WorldErrors *WorldErrorsSession) Execute(request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _WorldErrors.Contract.Execute(&_WorldErrors.TransactOpts, request)
}

// Execute is a paid mutator transaction binding the contract method 0xccff1ca1.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) payable returns()
func (_WorldErrors *WorldErrorsTransactorSession) Execute(request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _WorldErrors.Contract.Execute(&_WorldErrors.TransactOpts, request)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xa6b4806a.
//
// Solidity: function executeBatch((address,address,uint256,uint256,uint256,uint48,bytes,bytes)[] requests, address refundReceiver) payable returns()
func (_WorldErrors *WorldErrorsTransactor) ExecuteBatch(opts *bind.TransactOpts, requests []ERC2771ForwarderForwardRequestData, refundReceiver common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "executeBatch", requests, refundReceiver)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xa6b4806a.
//
// Solidity: function executeBatch((address,address,uint256,uint256,uint256,uint48,bytes,bytes)[] requests, address refundReceiver) payable returns()
func (_WorldErrors *WorldErrorsSession) ExecuteBatch(requests []ERC2771ForwarderForwardRequestData, refundReceiver common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.ExecuteBatch(&_WorldErrors.TransactOpts, requests, refundReceiver)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xa6b4806a.
//
// Solidity: function executeBatch((address,address,uint256,uint256,uint256,uint48,bytes,bytes)[] requests, address refundReceiver) payable returns()
func (_WorldErrors *WorldErrorsTransactorSession) ExecuteBatch(requests []ERC2771ForwarderForwardRequestData, refundReceiver common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.ExecuteBatch(&_WorldErrors.TransactOpts, requests, refundReceiver)
}

// GlobalPause is a paid mutator transaction binding the contract method 0xf12d54d8.
//
// Solidity: function globalPause() returns()
func (_WorldErrors *WorldErrorsTransactor) GlobalPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "globalPause")
}

// GlobalPause is a paid mutator transaction binding the contract method 0xf12d54d8.
//
// Solidity: function globalPause() returns()
func (_WorldErrors *WorldErrorsSession) GlobalPause() (*types.Transaction, error) {
	return _WorldErrors.Contract.GlobalPause(&_WorldErrors.TransactOpts)
}

// GlobalPause is a paid mutator transaction binding the contract method 0xf12d54d8.
//
// Solidity: function globalPause() returns()
func (_WorldErrors *WorldErrorsTransactorSession) GlobalPause() (*types.Transaction, error) {
	return _WorldErrors.Contract.GlobalPause(&_WorldErrors.TransactOpts)
}

// GlobalResume is a paid mutator transaction binding the contract method 0x59c7d378.
//
// Solidity: function globalResume() returns()
func (_WorldErrors *WorldErrorsTransactor) GlobalResume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "globalResume")
}

// GlobalResume is a paid mutator transaction binding the contract method 0x59c7d378.
//
// Solidity: function globalResume() returns()
func (_WorldErrors *WorldErrorsSession) GlobalResume() (*types.Transaction, error) {
	return _WorldErrors.Contract.GlobalResume(&_WorldErrors.TransactOpts)
}

// GlobalResume is a paid mutator transaction binding the contract method 0x59c7d378.
//
// Solidity: function globalResume() returns()
func (_WorldErrors *WorldErrorsTransactorSession) GlobalResume() (*types.Transaction, error) {
	return _WorldErrors.Contract.GlobalResume(&_WorldErrors.TransactOpts)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_WorldErrors *WorldErrorsTransactor) GrantAccess(opts *bind.TransactOpts, resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "grantAccess", resourceId, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_WorldErrors *WorldErrorsSession) GrantAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.GrantAccess(&_WorldErrors.TransactOpts, resourceId, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x40554c3a.
//
// Solidity: function grantAccess(bytes32 resourceId, address grantee) returns()
func (_WorldErrors *WorldErrorsTransactorSession) GrantAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.GrantAccess(&_WorldErrors.TransactOpts, resourceId, grantee)
}

// InitDelegation is a paid mutator transaction binding the contract method 0x78fcacbd.
//
// Solidity: function initDelegation(bytes32 namespaceId, address trustedForwarder) returns()
func (_WorldErrors *WorldErrorsTransactor) InitDelegation(opts *bind.TransactOpts, namespaceId [32]byte, trustedForwarder common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "initDelegation", namespaceId, trustedForwarder)
}

// InitDelegation is a paid mutator transaction binding the contract method 0x78fcacbd.
//
// Solidity: function initDelegation(bytes32 namespaceId, address trustedForwarder) returns()
func (_WorldErrors *WorldErrorsSession) InitDelegation(namespaceId [32]byte, trustedForwarder common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.InitDelegation(&_WorldErrors.TransactOpts, namespaceId, trustedForwarder)
}

// InitDelegation is a paid mutator transaction binding the contract method 0x78fcacbd.
//
// Solidity: function initDelegation(bytes32 namespaceId, address trustedForwarder) returns()
func (_WorldErrors *WorldErrorsTransactorSession) InitDelegation(namespaceId [32]byte, trustedForwarder common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.InitDelegation(&_WorldErrors.TransactOpts, namespaceId, trustedForwarder)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_WorldErrors *WorldErrorsTransactor) Initialize(opts *bind.TransactOpts, initModule common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "initialize", initModule)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_WorldErrors *WorldErrorsSession) Initialize(initModule common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.Initialize(&_WorldErrors.TransactOpts, initModule)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initModule) returns()
func (_WorldErrors *WorldErrorsTransactorSession) Initialize(initModule common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.Initialize(&_WorldErrors.TransactOpts, initModule)
}

// Install is a paid mutator transaction binding the contract method 0x13861fb5.
//
// Solidity: function install(bytes encodedArgs) returns()
func (_WorldErrors *WorldErrorsTransactor) Install(opts *bind.TransactOpts, encodedArgs []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "install", encodedArgs)
}

// Install is a paid mutator transaction binding the contract method 0x13861fb5.
//
// Solidity: function install(bytes encodedArgs) returns()
func (_WorldErrors *WorldErrorsSession) Install(encodedArgs []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.Install(&_WorldErrors.TransactOpts, encodedArgs)
}

// Install is a paid mutator transaction binding the contract method 0x13861fb5.
//
// Solidity: function install(bytes encodedArgs) returns()
func (_WorldErrors *WorldErrorsTransactorSession) Install(encodedArgs []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.Install(&_WorldErrors.TransactOpts, encodedArgs)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_WorldErrors *WorldErrorsTransactor) InstallModule(opts *bind.TransactOpts, module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "installModule", module, encodedArgs)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_WorldErrors *WorldErrorsSession) InstallModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.InstallModule(&_WorldErrors.TransactOpts, module, encodedArgs)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes encodedArgs) returns()
func (_WorldErrors *WorldErrorsTransactorSession) InstallModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.InstallModule(&_WorldErrors.TransactOpts, module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_WorldErrors *WorldErrorsTransactor) InstallRootModule(opts *bind.TransactOpts, module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "installRootModule", module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_WorldErrors *WorldErrorsSession) InstallRootModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.InstallRootModule(&_WorldErrors.TransactOpts, module, encodedArgs)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes encodedArgs) returns()
func (_WorldErrors *WorldErrorsTransactorSession) InstallRootModule(module common.Address, encodedArgs []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.InstallRootModule(&_WorldErrors.TransactOpts, module, encodedArgs)
}

// InventoryToEphemeralTransfer is a paid mutator transaction binding the contract method 0x181c3771.
//
// Solidity: function inventoryToEphemeralTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_WorldErrors *WorldErrorsTransactor) InventoryToEphemeralTransfer(opts *bind.TransactOpts, smartObjectId *big.Int, outItems []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "inventoryToEphemeralTransfer", smartObjectId, outItems)
}

// InventoryToEphemeralTransfer is a paid mutator transaction binding the contract method 0x181c3771.
//
// Solidity: function inventoryToEphemeralTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_WorldErrors *WorldErrorsSession) InventoryToEphemeralTransfer(smartObjectId *big.Int, outItems []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.InventoryToEphemeralTransfer(&_WorldErrors.TransactOpts, smartObjectId, outItems)
}

// InventoryToEphemeralTransfer is a paid mutator transaction binding the contract method 0x181c3771.
//
// Solidity: function inventoryToEphemeralTransfer(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] outItems) returns()
func (_WorldErrors *WorldErrorsTransactorSession) InventoryToEphemeralTransfer(smartObjectId *big.Int, outItems []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.InventoryToEphemeralTransfer(&_WorldErrors.TransactOpts, smartObjectId, outItems)
}

// IsTrustedForwarder is a paid mutator transaction binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) returns(bool)
func (_WorldErrors *WorldErrorsTransactor) IsTrustedForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "isTrustedForwarder", forwarder)
}

// IsTrustedForwarder is a paid mutator transaction binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) returns(bool)
func (_WorldErrors *WorldErrorsSession) IsTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.IsTrustedForwarder(&_WorldErrors.TransactOpts, forwarder)
}

// IsTrustedForwarder is a paid mutator transaction binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) returns(bool)
func (_WorldErrors *WorldErrorsTransactorSession) IsTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.IsTrustedForwarder(&_WorldErrors.TransactOpts, forwarder)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_WorldErrors *WorldErrorsTransactor) PopFromDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "popFromDynamicField", tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_WorldErrors *WorldErrorsSession) PopFromDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.PopFromDynamicField(&_WorldErrors.TransactOpts, tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PopFromDynamicField is a paid mutator transaction binding the contract method 0xd9c03a04.
//
// Solidity: function popFromDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint256 byteLengthToPop) returns()
func (_WorldErrors *WorldErrorsTransactorSession) PopFromDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.PopFromDynamicField(&_WorldErrors.TransactOpts, tableId, keyTuple, dynamicFieldIndex, byteLengthToPop)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_WorldErrors *WorldErrorsTransactor) PushToDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "pushToDynamicField", tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_WorldErrors *WorldErrorsSession) PushToDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.PushToDynamicField(&_WorldErrors.TransactOpts, tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// PushToDynamicField is a paid mutator transaction binding the contract method 0x150f3262.
//
// Solidity: function pushToDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes dataToPush) returns()
func (_WorldErrors *WorldErrorsTransactorSession) PushToDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.PushToDynamicField(&_WorldErrors.TransactOpts, tableId, keyTuple, dynamicFieldIndex, dataToPush)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterDelegation(opts *bind.TransactOpts, delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerDelegation", delegatee, delegationControlId, initCallData)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_WorldErrors *WorldErrorsSession) RegisterDelegation(delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterDelegation(&_WorldErrors.TransactOpts, delegatee, delegationControlId, initCallData)
}

// RegisterDelegation is a paid mutator transaction binding the contract method 0x1d2257ba.
//
// Solidity: function registerDelegation(address delegatee, bytes32 delegationControlId, bytes initCallData) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterDelegation(delegatee common.Address, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterDelegation(&_WorldErrors.TransactOpts, delegatee, delegationControlId, initCallData)
}

// RegisterDeployable is a paid mutator transaction binding the contract method 0x922bd0ca.
//
// Solidity: function registerDeployable(uint256 entityId, (address,string) smartObjectData, uint256 fuelUnitVolumeInWei, uint256 fuelConsumptionPerMinuteInWei, uint256 fuelMaxCapacityInWei) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterDeployable(opts *bind.TransactOpts, entityId *big.Int, smartObjectData SmartObjectData, fuelUnitVolumeInWei *big.Int, fuelConsumptionPerMinuteInWei *big.Int, fuelMaxCapacityInWei *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerDeployable", entityId, smartObjectData, fuelUnitVolumeInWei, fuelConsumptionPerMinuteInWei, fuelMaxCapacityInWei)
}

// RegisterDeployable is a paid mutator transaction binding the contract method 0x922bd0ca.
//
// Solidity: function registerDeployable(uint256 entityId, (address,string) smartObjectData, uint256 fuelUnitVolumeInWei, uint256 fuelConsumptionPerMinuteInWei, uint256 fuelMaxCapacityInWei) returns()
func (_WorldErrors *WorldErrorsSession) RegisterDeployable(entityId *big.Int, smartObjectData SmartObjectData, fuelUnitVolumeInWei *big.Int, fuelConsumptionPerMinuteInWei *big.Int, fuelMaxCapacityInWei *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterDeployable(&_WorldErrors.TransactOpts, entityId, smartObjectData, fuelUnitVolumeInWei, fuelConsumptionPerMinuteInWei, fuelMaxCapacityInWei)
}

// RegisterDeployable is a paid mutator transaction binding the contract method 0x922bd0ca.
//
// Solidity: function registerDeployable(uint256 entityId, (address,string) smartObjectData, uint256 fuelUnitVolumeInWei, uint256 fuelConsumptionPerMinuteInWei, uint256 fuelMaxCapacityInWei) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterDeployable(entityId *big.Int, smartObjectData SmartObjectData, fuelUnitVolumeInWei *big.Int, fuelConsumptionPerMinuteInWei *big.Int, fuelMaxCapacityInWei *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterDeployable(&_WorldErrors.TransactOpts, entityId, smartObjectData, fuelUnitVolumeInWei, fuelConsumptionPerMinuteInWei, fuelMaxCapacityInWei)
}

// RegisterDeployableToken is a paid mutator transaction binding the contract method 0x7b0d0e3c.
//
// Solidity: function registerDeployableToken(address tokenAddress) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterDeployableToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerDeployableToken", tokenAddress)
}

// RegisterDeployableToken is a paid mutator transaction binding the contract method 0x7b0d0e3c.
//
// Solidity: function registerDeployableToken(address tokenAddress) returns()
func (_WorldErrors *WorldErrorsSession) RegisterDeployableToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterDeployableToken(&_WorldErrors.TransactOpts, tokenAddress)
}

// RegisterDeployableToken is a paid mutator transaction binding the contract method 0x7b0d0e3c.
//
// Solidity: function registerDeployableToken(address tokenAddress) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterDeployableToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterDeployableToken(&_WorldErrors.TransactOpts, tokenAddress)
}

// RegisterERC721Token is a paid mutator transaction binding the contract method 0x2c3309d1.
//
// Solidity: function registerERC721Token(address tokenAddress) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterERC721Token(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerERC721Token", tokenAddress)
}

// RegisterERC721Token is a paid mutator transaction binding the contract method 0x2c3309d1.
//
// Solidity: function registerERC721Token(address tokenAddress) returns()
func (_WorldErrors *WorldErrorsSession) RegisterERC721Token(tokenAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterERC721Token(&_WorldErrors.TransactOpts, tokenAddress)
}

// RegisterERC721Token is a paid mutator transaction binding the contract method 0x2c3309d1.
//
// Solidity: function registerERC721Token(address tokenAddress) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterERC721Token(tokenAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterERC721Token(&_WorldErrors.TransactOpts, tokenAddress)
}

// RegisterEVEModule is a paid mutator transaction binding the contract method 0xbf1e3c14.
//
// Solidity: function registerEVEModule(uint256 moduleId, bytes16 moduleName, bytes32 systemId) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterEVEModule(opts *bind.TransactOpts, moduleId *big.Int, moduleName [16]byte, systemId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerEVEModule", moduleId, moduleName, systemId)
}

// RegisterEVEModule is a paid mutator transaction binding the contract method 0xbf1e3c14.
//
// Solidity: function registerEVEModule(uint256 moduleId, bytes16 moduleName, bytes32 systemId) returns()
func (_WorldErrors *WorldErrorsSession) RegisterEVEModule(moduleId *big.Int, moduleName [16]byte, systemId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEVEModule(&_WorldErrors.TransactOpts, moduleId, moduleName, systemId)
}

// RegisterEVEModule is a paid mutator transaction binding the contract method 0xbf1e3c14.
//
// Solidity: function registerEVEModule(uint256 moduleId, bytes16 moduleName, bytes32 systemId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterEVEModule(moduleId *big.Int, moduleName [16]byte, systemId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEVEModule(&_WorldErrors.TransactOpts, moduleId, moduleName, systemId)
}

// RegisterEVEModules is a paid mutator transaction binding the contract method 0x6629499f.
//
// Solidity: function registerEVEModules(uint256 moduleId, bytes16 moduleName, bytes32[] systemIds) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterEVEModules(opts *bind.TransactOpts, moduleId *big.Int, moduleName [16]byte, systemIds [][32]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerEVEModules", moduleId, moduleName, systemIds)
}

// RegisterEVEModules is a paid mutator transaction binding the contract method 0x6629499f.
//
// Solidity: function registerEVEModules(uint256 moduleId, bytes16 moduleName, bytes32[] systemIds) returns()
func (_WorldErrors *WorldErrorsSession) RegisterEVEModules(moduleId *big.Int, moduleName [16]byte, systemIds [][32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEVEModules(&_WorldErrors.TransactOpts, moduleId, moduleName, systemIds)
}

// RegisterEVEModules is a paid mutator transaction binding the contract method 0x6629499f.
//
// Solidity: function registerEVEModules(uint256 moduleId, bytes16 moduleName, bytes32[] systemIds) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterEVEModules(moduleId *big.Int, moduleName [16]byte, systemIds [][32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEVEModules(&_WorldErrors.TransactOpts, moduleId, moduleName, systemIds)
}

// RegisterEntities is a paid mutator transaction binding the contract method 0x2d42dda9.
//
// Solidity: function registerEntities(uint256[] entityId, uint8[] entityType) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterEntities(opts *bind.TransactOpts, entityId []*big.Int, entityType []uint8) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerEntities", entityId, entityType)
}

// RegisterEntities is a paid mutator transaction binding the contract method 0x2d42dda9.
//
// Solidity: function registerEntities(uint256[] entityId, uint8[] entityType) returns()
func (_WorldErrors *WorldErrorsSession) RegisterEntities(entityId []*big.Int, entityType []uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEntities(&_WorldErrors.TransactOpts, entityId, entityType)
}

// RegisterEntities is a paid mutator transaction binding the contract method 0x2d42dda9.
//
// Solidity: function registerEntities(uint256[] entityId, uint8[] entityType) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterEntities(entityId []*big.Int, entityType []uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEntities(&_WorldErrors.TransactOpts, entityId, entityType)
}

// RegisterEntity is a paid mutator transaction binding the contract method 0x71e86ed6.
//
// Solidity: function registerEntity(uint256 entityId, uint8 entityType) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterEntity(opts *bind.TransactOpts, entityId *big.Int, entityType uint8) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerEntity", entityId, entityType)
}

// RegisterEntity is a paid mutator transaction binding the contract method 0x71e86ed6.
//
// Solidity: function registerEntity(uint256 entityId, uint8 entityType) returns()
func (_WorldErrors *WorldErrorsSession) RegisterEntity(entityId *big.Int, entityType uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEntity(&_WorldErrors.TransactOpts, entityId, entityType)
}

// RegisterEntity is a paid mutator transaction binding the contract method 0x71e86ed6.
//
// Solidity: function registerEntity(uint256 entityId, uint8 entityType) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterEntity(entityId *big.Int, entityType uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEntity(&_WorldErrors.TransactOpts, entityId, entityType)
}

// RegisterEntityType is a paid mutator transaction binding the contract method 0x3a25de29.
//
// Solidity: function registerEntityType(uint8 entityTypeId, bytes32 entityType) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterEntityType(opts *bind.TransactOpts, entityTypeId uint8, entityType [32]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerEntityType", entityTypeId, entityType)
}

// RegisterEntityType is a paid mutator transaction binding the contract method 0x3a25de29.
//
// Solidity: function registerEntityType(uint8 entityTypeId, bytes32 entityType) returns()
func (_WorldErrors *WorldErrorsSession) RegisterEntityType(entityTypeId uint8, entityType [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEntityType(&_WorldErrors.TransactOpts, entityTypeId, entityType)
}

// RegisterEntityType is a paid mutator transaction binding the contract method 0x3a25de29.
//
// Solidity: function registerEntityType(uint8 entityTypeId, bytes32 entityType) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterEntityType(entityTypeId uint8, entityType [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEntityType(&_WorldErrors.TransactOpts, entityTypeId, entityType)
}

// RegisterEntityTypeAssociation is a paid mutator transaction binding the contract method 0x4a33bfcd.
//
// Solidity: function registerEntityTypeAssociation(uint8 entityType, uint8 tagEntityType) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterEntityTypeAssociation(opts *bind.TransactOpts, entityType uint8, tagEntityType uint8) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerEntityTypeAssociation", entityType, tagEntityType)
}

// RegisterEntityTypeAssociation is a paid mutator transaction binding the contract method 0x4a33bfcd.
//
// Solidity: function registerEntityTypeAssociation(uint8 entityType, uint8 tagEntityType) returns()
func (_WorldErrors *WorldErrorsSession) RegisterEntityTypeAssociation(entityType uint8, tagEntityType uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEntityTypeAssociation(&_WorldErrors.TransactOpts, entityType, tagEntityType)
}

// RegisterEntityTypeAssociation is a paid mutator transaction binding the contract method 0x4a33bfcd.
//
// Solidity: function registerEntityTypeAssociation(uint8 entityType, uint8 tagEntityType) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterEntityTypeAssociation(entityType uint8, tagEntityType uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterEntityTypeAssociation(&_WorldErrors.TransactOpts, entityType, tagEntityType)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_WorldErrors *WorldErrorsTransactor) RegisterFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerFunctionSelector", systemId, systemFunctionSignature)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_WorldErrors *WorldErrorsSession) RegisterFunctionSelector(systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterFunctionSelector(&_WorldErrors.TransactOpts, systemId, systemFunctionSignature)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0x26d98102.
//
// Solidity: function registerFunctionSelector(bytes32 systemId, string systemFunctionSignature) returns(bytes4 worldFunctionSelector)
func (_WorldErrors *WorldErrorsTransactorSession) RegisterFunctionSelector(systemId [32]byte, systemFunctionSignature string) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterFunctionSelector(&_WorldErrors.TransactOpts, systemId, systemFunctionSignature)
}

// RegisterHook is a paid mutator transaction binding the contract method 0x23e30ed6.
//
// Solidity: function registerHook(bytes32 systemId, bytes4 functionId) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterHook(opts *bind.TransactOpts, systemId [32]byte, functionId [4]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerHook", systemId, functionId)
}

// RegisterHook is a paid mutator transaction binding the contract method 0x23e30ed6.
//
// Solidity: function registerHook(bytes32 systemId, bytes4 functionId) returns()
func (_WorldErrors *WorldErrorsSession) RegisterHook(systemId [32]byte, functionId [4]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterHook(&_WorldErrors.TransactOpts, systemId, functionId)
}

// RegisterHook is a paid mutator transaction binding the contract method 0x23e30ed6.
//
// Solidity: function registerHook(bytes32 systemId, bytes4 functionId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterHook(systemId [32]byte, functionId [4]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterHook(&_WorldErrors.TransactOpts, systemId, functionId)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterNamespace(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerNamespace", namespaceId)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_WorldErrors *WorldErrorsSession) RegisterNamespace(namespaceId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterNamespace(&_WorldErrors.TransactOpts, namespaceId)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xb29e4089.
//
// Solidity: function registerNamespace(bytes32 namespaceId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterNamespace(namespaceId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterNamespace(&_WorldErrors.TransactOpts, namespaceId)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterNamespaceDelegation(opts *bind.TransactOpts, namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerNamespaceDelegation", namespaceId, delegationControlId, initCallData)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_WorldErrors *WorldErrorsSession) RegisterNamespaceDelegation(namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterNamespaceDelegation(&_WorldErrors.TransactOpts, namespaceId, delegationControlId, initCallData)
}

// RegisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xbfdfaff7.
//
// Solidity: function registerNamespaceDelegation(bytes32 namespaceId, bytes32 delegationControlId, bytes initCallData) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterNamespaceDelegation(namespaceId [32]byte, delegationControlId [32]byte, initCallData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterNamespaceDelegation(&_WorldErrors.TransactOpts, namespaceId, delegationControlId, initCallData)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x742d6118.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, bytes4 systemFunctionSelector) returns(bytes4 worldFunctionSelector)
func (_WorldErrors *WorldErrorsTransactor) RegisterRootFunctionSelector(opts *bind.TransactOpts, systemId [32]byte, worldFunctionSignature string, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerRootFunctionSelector", systemId, worldFunctionSignature, systemFunctionSelector)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x742d6118.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, bytes4 systemFunctionSelector) returns(bytes4 worldFunctionSelector)
func (_WorldErrors *WorldErrorsSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterRootFunctionSelector(&_WorldErrors.TransactOpts, systemId, worldFunctionSignature, systemFunctionSelector)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x742d6118.
//
// Solidity: function registerRootFunctionSelector(bytes32 systemId, string worldFunctionSignature, bytes4 systemFunctionSelector) returns(bytes4 worldFunctionSelector)
func (_WorldErrors *WorldErrorsTransactorSession) RegisterRootFunctionSelector(systemId [32]byte, worldFunctionSignature string, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterRootFunctionSelector(&_WorldErrors.TransactOpts, systemId, worldFunctionSignature, systemFunctionSelector)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterStoreHook(opts *bind.TransactOpts, tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerStoreHook", tableId, hookAddress, enabledHooksBitmap)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_WorldErrors *WorldErrorsSession) RegisterStoreHook(tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterStoreHook(&_WorldErrors.TransactOpts, tableId, hookAddress, enabledHooksBitmap)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x530f4b60.
//
// Solidity: function registerStoreHook(bytes32 tableId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterStoreHook(tableId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterStoreHook(&_WorldErrors.TransactOpts, tableId, hookAddress, enabledHooksBitmap)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterSystem(opts *bind.TransactOpts, systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerSystem", systemId, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_WorldErrors *WorldErrorsSession) RegisterSystem(systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterSystem(&_WorldErrors.TransactOpts, systemId, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x3350b6a9.
//
// Solidity: function registerSystem(bytes32 systemId, address system, bool publicAccess) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterSystem(systemId [32]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterSystem(&_WorldErrors.TransactOpts, systemId, system, publicAccess)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterSystemHook(opts *bind.TransactOpts, systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerSystemHook", systemId, hookAddress, enabledHooksBitmap)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_WorldErrors *WorldErrorsSession) RegisterSystemHook(systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterSystemHook(&_WorldErrors.TransactOpts, systemId, hookAddress, enabledHooksBitmap)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0xd5f8337f.
//
// Solidity: function registerSystemHook(bytes32 systemId, address hookAddress, uint8 enabledHooksBitmap) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterSystemHook(systemId [32]byte, hookAddress common.Address, enabledHooksBitmap uint8) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterSystemHook(&_WorldErrors.TransactOpts, systemId, hookAddress, enabledHooksBitmap)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_WorldErrors *WorldErrorsTransactor) RegisterTable(opts *bind.TransactOpts, tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "registerTable", tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_WorldErrors *WorldErrorsSession) RegisterTable(tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterTable(&_WorldErrors.TransactOpts, tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x0ba51f49.
//
// Solidity: function registerTable(bytes32 tableId, bytes32 fieldLayout, bytes32 keySchema, bytes32 valueSchema, string[] keyNames, string[] fieldNames) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RegisterTable(tableId [32]byte, fieldLayout [32]byte, keySchema [32]byte, valueSchema [32]byte, keyNames []string, fieldNames []string) (*types.Transaction, error) {
	return _WorldErrors.Contract.RegisterTable(&_WorldErrors.TransactOpts, tableId, fieldLayout, keySchema, valueSchema, keyNames, fieldNames)
}

// RemoveEntityHookAssociation is a paid mutator transaction binding the contract method 0x9f8605e0.
//
// Solidity: function removeEntityHookAssociation(uint256 entityId, uint256 hookId) returns()
func (_WorldErrors *WorldErrorsTransactor) RemoveEntityHookAssociation(opts *bind.TransactOpts, entityId *big.Int, hookId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "removeEntityHookAssociation", entityId, hookId)
}

// RemoveEntityHookAssociation is a paid mutator transaction binding the contract method 0x9f8605e0.
//
// Solidity: function removeEntityHookAssociation(uint256 entityId, uint256 hookId) returns()
func (_WorldErrors *WorldErrorsSession) RemoveEntityHookAssociation(entityId *big.Int, hookId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveEntityHookAssociation(&_WorldErrors.TransactOpts, entityId, hookId)
}

// RemoveEntityHookAssociation is a paid mutator transaction binding the contract method 0x9f8605e0.
//
// Solidity: function removeEntityHookAssociation(uint256 entityId, uint256 hookId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RemoveEntityHookAssociation(entityId *big.Int, hookId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveEntityHookAssociation(&_WorldErrors.TransactOpts, entityId, hookId)
}

// RemoveEntityModuleAssociation is a paid mutator transaction binding the contract method 0xcbe46b64.
//
// Solidity: function removeEntityModuleAssociation(uint256 entityId, uint256 moduleId) returns()
func (_WorldErrors *WorldErrorsTransactor) RemoveEntityModuleAssociation(opts *bind.TransactOpts, entityId *big.Int, moduleId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "removeEntityModuleAssociation", entityId, moduleId)
}

// RemoveEntityModuleAssociation is a paid mutator transaction binding the contract method 0xcbe46b64.
//
// Solidity: function removeEntityModuleAssociation(uint256 entityId, uint256 moduleId) returns()
func (_WorldErrors *WorldErrorsSession) RemoveEntityModuleAssociation(entityId *big.Int, moduleId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveEntityModuleAssociation(&_WorldErrors.TransactOpts, entityId, moduleId)
}

// RemoveEntityModuleAssociation is a paid mutator transaction binding the contract method 0xcbe46b64.
//
// Solidity: function removeEntityModuleAssociation(uint256 entityId, uint256 moduleId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RemoveEntityModuleAssociation(entityId *big.Int, moduleId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveEntityModuleAssociation(&_WorldErrors.TransactOpts, entityId, moduleId)
}

// RemoveEntityTag is a paid mutator transaction binding the contract method 0x2390ddc0.
//
// Solidity: function removeEntityTag(uint256 entityId, uint256 entityTagId) returns()
func (_WorldErrors *WorldErrorsTransactor) RemoveEntityTag(opts *bind.TransactOpts, entityId *big.Int, entityTagId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "removeEntityTag", entityId, entityTagId)
}

// RemoveEntityTag is a paid mutator transaction binding the contract method 0x2390ddc0.
//
// Solidity: function removeEntityTag(uint256 entityId, uint256 entityTagId) returns()
func (_WorldErrors *WorldErrorsSession) RemoveEntityTag(entityId *big.Int, entityTagId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveEntityTag(&_WorldErrors.TransactOpts, entityId, entityTagId)
}

// RemoveEntityTag is a paid mutator transaction binding the contract method 0x2390ddc0.
//
// Solidity: function removeEntityTag(uint256 entityId, uint256 entityTagId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RemoveEntityTag(entityId *big.Int, entityTagId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveEntityTag(&_WorldErrors.TransactOpts, entityId, entityTagId)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x7b8b7e17.
//
// Solidity: function removeHook(uint256 hookId, uint8 hookType, bytes32 systemId, bytes4 functionSelector) returns()
func (_WorldErrors *WorldErrorsTransactor) RemoveHook(opts *bind.TransactOpts, hookId *big.Int, hookType uint8, systemId [32]byte, functionSelector [4]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "removeHook", hookId, hookType, systemId, functionSelector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x7b8b7e17.
//
// Solidity: function removeHook(uint256 hookId, uint8 hookType, bytes32 systemId, bytes4 functionSelector) returns()
func (_WorldErrors *WorldErrorsSession) RemoveHook(hookId *big.Int, hookType uint8, systemId [32]byte, functionSelector [4]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveHook(&_WorldErrors.TransactOpts, hookId, hookType, systemId, functionSelector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x7b8b7e17.
//
// Solidity: function removeHook(uint256 hookId, uint8 hookType, bytes32 systemId, bytes4 functionSelector) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RemoveHook(hookId *big.Int, hookType uint8, systemId [32]byte, functionSelector [4]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveHook(&_WorldErrors.TransactOpts, hookId, hookType, systemId, functionSelector)
}

// RemoveSystemModuleAssociation is a paid mutator transaction binding the contract method 0x7e6dd49a.
//
// Solidity: function removeSystemModuleAssociation(bytes32 systemId, uint256 moduleId) returns()
func (_WorldErrors *WorldErrorsTransactor) RemoveSystemModuleAssociation(opts *bind.TransactOpts, systemId [32]byte, moduleId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "removeSystemModuleAssociation", systemId, moduleId)
}

// RemoveSystemModuleAssociation is a paid mutator transaction binding the contract method 0x7e6dd49a.
//
// Solidity: function removeSystemModuleAssociation(bytes32 systemId, uint256 moduleId) returns()
func (_WorldErrors *WorldErrorsSession) RemoveSystemModuleAssociation(systemId [32]byte, moduleId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveSystemModuleAssociation(&_WorldErrors.TransactOpts, systemId, moduleId)
}

// RemoveSystemModuleAssociation is a paid mutator transaction binding the contract method 0x7e6dd49a.
//
// Solidity: function removeSystemModuleAssociation(bytes32 systemId, uint256 moduleId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RemoveSystemModuleAssociation(systemId [32]byte, moduleId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.RemoveSystemModuleAssociation(&_WorldErrors.TransactOpts, systemId, moduleId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_WorldErrors *WorldErrorsTransactor) RenounceOwnership(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "renounceOwnership", namespaceId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_WorldErrors *WorldErrorsSession) RenounceOwnership(namespaceId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RenounceOwnership(&_WorldErrors.TransactOpts, namespaceId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x219adc2e.
//
// Solidity: function renounceOwnership(bytes32 namespaceId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RenounceOwnership(namespaceId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.RenounceOwnership(&_WorldErrors.TransactOpts, namespaceId)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_WorldErrors *WorldErrorsTransactor) RevokeAccess(opts *bind.TransactOpts, resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "revokeAccess", resourceId, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_WorldErrors *WorldErrorsSession) RevokeAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.RevokeAccess(&_WorldErrors.TransactOpts, resourceId, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x8d53b208.
//
// Solidity: function revokeAccess(bytes32 resourceId, address grantee) returns()
func (_WorldErrors *WorldErrorsTransactorSession) RevokeAccess(resourceId [32]byte, grantee common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.RevokeAccess(&_WorldErrors.TransactOpts, resourceId, grantee)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 entityId, (uint256,uint256,uint256,uint256) location) returns()
func (_WorldErrors *WorldErrorsTransactor) SaveLocation(opts *bind.TransactOpts, entityId *big.Int, location LocationTableData) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "saveLocation", entityId, location)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 entityId, (uint256,uint256,uint256,uint256) location) returns()
func (_WorldErrors *WorldErrorsSession) SaveLocation(entityId *big.Int, location LocationTableData) (*types.Transaction, error) {
	return _WorldErrors.Contract.SaveLocation(&_WorldErrors.TransactOpts, entityId, location)
}

// SaveLocation is a paid mutator transaction binding the contract method 0x2f525c1c.
//
// Solidity: function saveLocation(uint256 entityId, (uint256,uint256,uint256,uint256) location) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SaveLocation(entityId *big.Int, location LocationTableData) (*types.Transaction, error) {
	return _WorldErrors.Contract.SaveLocation(&_WorldErrors.TransactOpts, entityId, location)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x8bf3d594.
//
// Solidity: function setBaseURI(bytes32 systemId, string baseURI) returns()
func (_WorldErrors *WorldErrorsTransactor) SetBaseURI(opts *bind.TransactOpts, systemId [32]byte, baseURI string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setBaseURI", systemId, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x8bf3d594.
//
// Solidity: function setBaseURI(bytes32 systemId, string baseURI) returns()
func (_WorldErrors *WorldErrorsSession) SetBaseURI(systemId [32]byte, baseURI string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetBaseURI(&_WorldErrors.TransactOpts, systemId, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x8bf3d594.
//
// Solidity: function setBaseURI(bytes32 systemId, string baseURI) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetBaseURI(systemId [32]byte, baseURI string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetBaseURI(&_WorldErrors.TransactOpts, systemId, baseURI)
}

// SetCharClassId is a paid mutator transaction binding the contract method 0xf83a5887.
//
// Solidity: function setCharClassId(uint256 classId) returns()
func (_WorldErrors *WorldErrorsTransactor) SetCharClassId(opts *bind.TransactOpts, classId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setCharClassId", classId)
}

// SetCharClassId is a paid mutator transaction binding the contract method 0xf83a5887.
//
// Solidity: function setCharClassId(uint256 classId) returns()
func (_WorldErrors *WorldErrorsSession) SetCharClassId(classId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetCharClassId(&_WorldErrors.TransactOpts, classId)
}

// SetCharClassId is a paid mutator transaction binding the contract method 0xf83a5887.
//
// Solidity: function setCharClassId(uint256 classId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetCharClassId(classId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetCharClassId(&_WorldErrors.TransactOpts, classId)
}

// SetCid is a paid mutator transaction binding the contract method 0x1cc5fe59.
//
// Solidity: function setCid(uint256 entityId, string cid) returns()
func (_WorldErrors *WorldErrorsTransactor) SetCid(opts *bind.TransactOpts, entityId *big.Int, cid string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setCid", entityId, cid)
}

// SetCid is a paid mutator transaction binding the contract method 0x1cc5fe59.
//
// Solidity: function setCid(uint256 entityId, string cid) returns()
func (_WorldErrors *WorldErrorsSession) SetCid(entityId *big.Int, cid string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetCid(&_WorldErrors.TransactOpts, entityId, cid)
}

// SetCid is a paid mutator transaction binding the contract method 0x1cc5fe59.
//
// Solidity: function setCid(uint256 entityId, string cid) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetCid(entityId *big.Int, cid string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetCid(&_WorldErrors.TransactOpts, entityId, cid)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 entityId, string dappURL) returns()
func (_WorldErrors *WorldErrorsTransactor) SetDappURL(opts *bind.TransactOpts, entityId *big.Int, dappURL string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setDappURL", entityId, dappURL)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 entityId, string dappURL) returns()
func (_WorldErrors *WorldErrorsSession) SetDappURL(entityId *big.Int, dappURL string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetDappURL(&_WorldErrors.TransactOpts, entityId, dappURL)
}

// SetDappURL is a paid mutator transaction binding the contract method 0x63b6b498.
//
// Solidity: function setDappURL(uint256 entityId, string dappURL) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetDappURL(entityId *big.Int, dappURL string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetDappURL(&_WorldErrors.TransactOpts, entityId, dappURL)
}

// SetDeploybaleMetadata is a paid mutator transaction binding the contract method 0x10f38d54.
//
// Solidity: function setDeploybaleMetadata(uint256 smartObjectId, string name, string dappURL, string description) returns()
func (_WorldErrors *WorldErrorsTransactor) SetDeploybaleMetadata(opts *bind.TransactOpts, smartObjectId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setDeploybaleMetadata", smartObjectId, name, dappURL, description)
}

// SetDeploybaleMetadata is a paid mutator transaction binding the contract method 0x10f38d54.
//
// Solidity: function setDeploybaleMetadata(uint256 smartObjectId, string name, string dappURL, string description) returns()
func (_WorldErrors *WorldErrorsSession) SetDeploybaleMetadata(smartObjectId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetDeploybaleMetadata(&_WorldErrors.TransactOpts, smartObjectId, name, dappURL, description)
}

// SetDeploybaleMetadata is a paid mutator transaction binding the contract method 0x10f38d54.
//
// Solidity: function setDeploybaleMetadata(uint256 smartObjectId, string name, string dappURL, string description) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetDeploybaleMetadata(smartObjectId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetDeploybaleMetadata(&_WorldErrors.TransactOpts, smartObjectId, name, dappURL, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 entityId, string description) returns()
func (_WorldErrors *WorldErrorsTransactor) SetDescription(opts *bind.TransactOpts, entityId *big.Int, description string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setDescription", entityId, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 entityId, string description) returns()
func (_WorldErrors *WorldErrorsSession) SetDescription(entityId *big.Int, description string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetDescription(&_WorldErrors.TransactOpts, entityId, description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x2a6446ca.
//
// Solidity: function setDescription(uint256 entityId, string description) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetDescription(entityId *big.Int, description string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetDescription(&_WorldErrors.TransactOpts, entityId, description)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_WorldErrors *WorldErrorsTransactor) SetDynamicField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setDynamicField", tableId, keyTuple, dynamicFieldIndex, data)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_WorldErrors *WorldErrorsSession) SetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetDynamicField(&_WorldErrors.TransactOpts, tableId, keyTuple, dynamicFieldIndex, data)
}

// SetDynamicField is a paid mutator transaction binding the contract method 0xef6ea862.
//
// Solidity: function setDynamicField(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, bytes data) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetDynamicField(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetDynamicField(&_WorldErrors.TransactOpts, tableId, keyTuple, dynamicFieldIndex, data)
}

// SetEntityMetadata is a paid mutator transaction binding the contract method 0x5f9c496a.
//
// Solidity: function setEntityMetadata(uint256 entityId, string name, string dappURL, string description) returns()
func (_WorldErrors *WorldErrorsTransactor) SetEntityMetadata(opts *bind.TransactOpts, entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setEntityMetadata", entityId, name, dappURL, description)
}

// SetEntityMetadata is a paid mutator transaction binding the contract method 0x5f9c496a.
//
// Solidity: function setEntityMetadata(uint256 entityId, string name, string dappURL, string description) returns()
func (_WorldErrors *WorldErrorsSession) SetEntityMetadata(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetEntityMetadata(&_WorldErrors.TransactOpts, entityId, name, dappURL, description)
}

// SetEntityMetadata is a paid mutator transaction binding the contract method 0x5f9c496a.
//
// Solidity: function setEntityMetadata(uint256 entityId, string name, string dappURL, string description) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetEntityMetadata(entityId *big.Int, name string, dappURL string, description string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetEntityMetadata(&_WorldErrors.TransactOpts, entityId, name, dappURL, description)
}

// SetEphemeralInventoryCapacity is a paid mutator transaction binding the contract method 0x20305602.
//
// Solidity: function setEphemeralInventoryCapacity(uint256 smartObjectId, uint256 ephemeralStorageCapacity) returns()
func (_WorldErrors *WorldErrorsTransactor) SetEphemeralInventoryCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setEphemeralInventoryCapacity", smartObjectId, ephemeralStorageCapacity)
}

// SetEphemeralInventoryCapacity is a paid mutator transaction binding the contract method 0x20305602.
//
// Solidity: function setEphemeralInventoryCapacity(uint256 smartObjectId, uint256 ephemeralStorageCapacity) returns()
func (_WorldErrors *WorldErrorsSession) SetEphemeralInventoryCapacity(smartObjectId *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetEphemeralInventoryCapacity(&_WorldErrors.TransactOpts, smartObjectId, ephemeralStorageCapacity)
}

// SetEphemeralInventoryCapacity is a paid mutator transaction binding the contract method 0x20305602.
//
// Solidity: function setEphemeralInventoryCapacity(uint256 smartObjectId, uint256 ephemeralStorageCapacity) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetEphemeralInventoryCapacity(smartObjectId *big.Int, ephemeralStorageCapacity *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetEphemeralInventoryCapacity(&_WorldErrors.TransactOpts, smartObjectId, ephemeralStorageCapacity)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_WorldErrors *WorldErrorsTransactor) SetField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setField", tableId, keyTuple, fieldIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_WorldErrors *WorldErrorsSession) SetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetField(&_WorldErrors.TransactOpts, tableId, keyTuple, fieldIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetField(&_WorldErrors.TransactOpts, tableId, keyTuple, fieldIndex, data)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_WorldErrors *WorldErrorsTransactor) SetField0(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setField0", tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_WorldErrors *WorldErrorsSession) SetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetField0(&_WorldErrors.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetField0 is a paid mutator transaction binding the contract method 0x3708196e.
//
// Solidity: function setField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetField0(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetField0(&_WorldErrors.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetFuelConsumptionPerMinute is a paid mutator transaction binding the contract method 0xeb37e8f0.
//
// Solidity: function setFuelConsumptionPerMinute(uint256 entityId, uint256 fuelConsumptionPerMinuteInWei) returns()
func (_WorldErrors *WorldErrorsTransactor) SetFuelConsumptionPerMinute(opts *bind.TransactOpts, entityId *big.Int, fuelConsumptionPerMinuteInWei *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setFuelConsumptionPerMinute", entityId, fuelConsumptionPerMinuteInWei)
}

// SetFuelConsumptionPerMinute is a paid mutator transaction binding the contract method 0xeb37e8f0.
//
// Solidity: function setFuelConsumptionPerMinute(uint256 entityId, uint256 fuelConsumptionPerMinuteInWei) returns()
func (_WorldErrors *WorldErrorsSession) SetFuelConsumptionPerMinute(entityId *big.Int, fuelConsumptionPerMinuteInWei *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetFuelConsumptionPerMinute(&_WorldErrors.TransactOpts, entityId, fuelConsumptionPerMinuteInWei)
}

// SetFuelConsumptionPerMinute is a paid mutator transaction binding the contract method 0xeb37e8f0.
//
// Solidity: function setFuelConsumptionPerMinute(uint256 entityId, uint256 fuelConsumptionPerMinuteInWei) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetFuelConsumptionPerMinute(entityId *big.Int, fuelConsumptionPerMinuteInWei *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetFuelConsumptionPerMinute(&_WorldErrors.TransactOpts, entityId, fuelConsumptionPerMinuteInWei)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 entityId, uint256 capacityInWei) returns()
func (_WorldErrors *WorldErrorsTransactor) SetFuelMaxCapacity(opts *bind.TransactOpts, entityId *big.Int, capacityInWei *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setFuelMaxCapacity", entityId, capacityInWei)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 entityId, uint256 capacityInWei) returns()
func (_WorldErrors *WorldErrorsSession) SetFuelMaxCapacity(entityId *big.Int, capacityInWei *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetFuelMaxCapacity(&_WorldErrors.TransactOpts, entityId, capacityInWei)
}

// SetFuelMaxCapacity is a paid mutator transaction binding the contract method 0xb25f99bf.
//
// Solidity: function setFuelMaxCapacity(uint256 entityId, uint256 capacityInWei) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetFuelMaxCapacity(entityId *big.Int, capacityInWei *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetFuelMaxCapacity(&_WorldErrors.TransactOpts, entityId, capacityInWei)
}

// SetInventoryCapacity is a paid mutator transaction binding the contract method 0x21b01b1d.
//
// Solidity: function setInventoryCapacity(uint256 smartObjectId, uint256 storageCapacity) returns()
func (_WorldErrors *WorldErrorsTransactor) SetInventoryCapacity(opts *bind.TransactOpts, smartObjectId *big.Int, storageCapacity *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setInventoryCapacity", smartObjectId, storageCapacity)
}

// SetInventoryCapacity is a paid mutator transaction binding the contract method 0x21b01b1d.
//
// Solidity: function setInventoryCapacity(uint256 smartObjectId, uint256 storageCapacity) returns()
func (_WorldErrors *WorldErrorsSession) SetInventoryCapacity(smartObjectId *big.Int, storageCapacity *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetInventoryCapacity(&_WorldErrors.TransactOpts, smartObjectId, storageCapacity)
}

// SetInventoryCapacity is a paid mutator transaction binding the contract method 0x21b01b1d.
//
// Solidity: function setInventoryCapacity(uint256 smartObjectId, uint256 storageCapacity) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetInventoryCapacity(smartObjectId *big.Int, storageCapacity *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetInventoryCapacity(&_WorldErrors.TransactOpts, smartObjectId, storageCapacity)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x7eaf1400.
//
// Solidity: function setMetadata(bytes32 systemId, (string,string,string) data) returns()
func (_WorldErrors *WorldErrorsTransactor) SetMetadata(opts *bind.TransactOpts, systemId [32]byte, data StaticDataGlobalTableData) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setMetadata", systemId, data)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x7eaf1400.
//
// Solidity: function setMetadata(bytes32 systemId, (string,string,string) data) returns()
func (_WorldErrors *WorldErrorsSession) SetMetadata(systemId [32]byte, data StaticDataGlobalTableData) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetMetadata(&_WorldErrors.TransactOpts, systemId, data)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x7eaf1400.
//
// Solidity: function setMetadata(bytes32 systemId, (string,string,string) data) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetMetadata(systemId [32]byte, data StaticDataGlobalTableData) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetMetadata(&_WorldErrors.TransactOpts, systemId, data)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 systemId, string name) returns()
func (_WorldErrors *WorldErrorsTransactor) SetName(opts *bind.TransactOpts, systemId [32]byte, name string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setName", systemId, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 systemId, string name) returns()
func (_WorldErrors *WorldErrorsSession) SetName(systemId [32]byte, name string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetName(&_WorldErrors.TransactOpts, systemId, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 systemId, string name) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetName(systemId [32]byte, name string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetName(&_WorldErrors.TransactOpts, systemId, name)
}

// SetName0 is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 entityId, string name) returns()
func (_WorldErrors *WorldErrorsTransactor) SetName0(opts *bind.TransactOpts, entityId *big.Int, name string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setName0", entityId, name)
}

// SetName0 is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 entityId, string name) returns()
func (_WorldErrors *WorldErrorsSession) SetName0(entityId *big.Int, name string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetName0(&_WorldErrors.TransactOpts, entityId, name)
}

// SetName0 is a paid mutator transaction binding the contract method 0xfe55932a.
//
// Solidity: function setName(uint256 entityId, string name) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetName0(entityId *big.Int, name string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetName0(&_WorldErrors.TransactOpts, entityId, name)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_WorldErrors *WorldErrorsTransactor) SetRecord(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setRecord", tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_WorldErrors *WorldErrorsSession) SetRecord(tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetRecord(&_WorldErrors.TransactOpts, tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetRecord is a paid mutator transaction binding the contract method 0x298314fb.
//
// Solidity: function setRecord(bytes32 tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetRecord(tableId [32]byte, keyTuple [][32]byte, staticData []byte, encodedLengths [32]byte, dynamicData []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetRecord(&_WorldErrors.TransactOpts, tableId, keyTuple, staticData, encodedLengths, dynamicData)
}

// SetSSUClassId is a paid mutator transaction binding the contract method 0xe2afb03b.
//
// Solidity: function setSSUClassId(uint256 classId) returns()
func (_WorldErrors *WorldErrorsTransactor) SetSSUClassId(opts *bind.TransactOpts, classId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setSSUClassId", classId)
}

// SetSSUClassId is a paid mutator transaction binding the contract method 0xe2afb03b.
//
// Solidity: function setSSUClassId(uint256 classId) returns()
func (_WorldErrors *WorldErrorsSession) SetSSUClassId(classId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetSSUClassId(&_WorldErrors.TransactOpts, classId)
}

// SetSSUClassId is a paid mutator transaction binding the contract method 0xe2afb03b.
//
// Solidity: function setSSUClassId(uint256 classId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetSSUClassId(classId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetSSUClassId(&_WorldErrors.TransactOpts, classId)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_WorldErrors *WorldErrorsTransactor) SetStaticField(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setStaticField", tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_WorldErrors *WorldErrorsSession) SetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetStaticField(&_WorldErrors.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetStaticField is a paid mutator transaction binding the contract method 0x390baae0.
//
// Solidity: function setStaticField(bytes32 tableId, bytes32[] keyTuple, uint8 fieldIndex, bytes data, bytes32 fieldLayout) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetStaticField(tableId [32]byte, keyTuple [][32]byte, fieldIndex uint8, data []byte, fieldLayout [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetStaticField(&_WorldErrors.TransactOpts, tableId, keyTuple, fieldIndex, data, fieldLayout)
}

// SetSymbol is a paid mutator transaction binding the contract method 0x0b1cb716.
//
// Solidity: function setSymbol(bytes32 systemId, string symbol) returns()
func (_WorldErrors *WorldErrorsTransactor) SetSymbol(opts *bind.TransactOpts, systemId [32]byte, symbol string) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setSymbol", systemId, symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0x0b1cb716.
//
// Solidity: function setSymbol(bytes32 systemId, string symbol) returns()
func (_WorldErrors *WorldErrorsSession) SetSymbol(systemId [32]byte, symbol string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetSymbol(&_WorldErrors.TransactOpts, systemId, symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0x0b1cb716.
//
// Solidity: function setSymbol(bytes32 systemId, string symbol) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetSymbol(systemId [32]byte, symbol string) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetSymbol(&_WorldErrors.TransactOpts, systemId, symbol)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address forwarder) returns()
func (_WorldErrors *WorldErrorsTransactor) SetTrustedForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "setTrustedForwarder", forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address forwarder) returns()
func (_WorldErrors *WorldErrorsSession) SetTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetTrustedForwarder(&_WorldErrors.TransactOpts, forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address forwarder) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SetTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.SetTrustedForwarder(&_WorldErrors.TransactOpts, forwarder)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_WorldErrors *WorldErrorsTransactor) SpliceDynamicData(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "spliceDynamicData", tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_WorldErrors *WorldErrorsSession) SpliceDynamicData(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SpliceDynamicData(&_WorldErrors.TransactOpts, tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceDynamicData is a paid mutator transaction binding the contract method 0xc0a2895a.
//
// Solidity: function spliceDynamicData(bytes32 tableId, bytes32[] keyTuple, uint8 dynamicFieldIndex, uint40 startWithinField, uint40 deleteCount, bytes data) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SpliceDynamicData(tableId [32]byte, keyTuple [][32]byte, dynamicFieldIndex uint8, startWithinField *big.Int, deleteCount *big.Int, data []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SpliceDynamicData(&_WorldErrors.TransactOpts, tableId, keyTuple, dynamicFieldIndex, startWithinField, deleteCount, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_WorldErrors *WorldErrorsTransactor) SpliceStaticData(opts *bind.TransactOpts, tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "spliceStaticData", tableId, keyTuple, start, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_WorldErrors *WorldErrorsSession) SpliceStaticData(tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SpliceStaticData(&_WorldErrors.TransactOpts, tableId, keyTuple, start, data)
}

// SpliceStaticData is a paid mutator transaction binding the contract method 0xb047c1eb.
//
// Solidity: function spliceStaticData(bytes32 tableId, bytes32[] keyTuple, uint48 start, bytes data) returns()
func (_WorldErrors *WorldErrorsTransactorSession) SpliceStaticData(tableId [32]byte, keyTuple [][32]byte, start *big.Int, data []byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.SpliceStaticData(&_WorldErrors.TransactOpts, tableId, keyTuple, start, data)
}

// TagEntities is a paid mutator transaction binding the contract method 0xd4ae42b9.
//
// Solidity: function tagEntities(uint256 entityId, uint256[] entityTagIds) returns()
func (_WorldErrors *WorldErrorsTransactor) TagEntities(opts *bind.TransactOpts, entityId *big.Int, entityTagIds []*big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "tagEntities", entityId, entityTagIds)
}

// TagEntities is a paid mutator transaction binding the contract method 0xd4ae42b9.
//
// Solidity: function tagEntities(uint256 entityId, uint256[] entityTagIds) returns()
func (_WorldErrors *WorldErrorsSession) TagEntities(entityId *big.Int, entityTagIds []*big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.TagEntities(&_WorldErrors.TransactOpts, entityId, entityTagIds)
}

// TagEntities is a paid mutator transaction binding the contract method 0xd4ae42b9.
//
// Solidity: function tagEntities(uint256 entityId, uint256[] entityTagIds) returns()
func (_WorldErrors *WorldErrorsTransactorSession) TagEntities(entityId *big.Int, entityTagIds []*big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.TagEntities(&_WorldErrors.TransactOpts, entityId, entityTagIds)
}

// TagEntity is a paid mutator transaction binding the contract method 0x05819208.
//
// Solidity: function tagEntity(uint256 entityId, uint256 entityTagId) returns()
func (_WorldErrors *WorldErrorsTransactor) TagEntity(opts *bind.TransactOpts, entityId *big.Int, entityTagId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "tagEntity", entityId, entityTagId)
}

// TagEntity is a paid mutator transaction binding the contract method 0x05819208.
//
// Solidity: function tagEntity(uint256 entityId, uint256 entityTagId) returns()
func (_WorldErrors *WorldErrorsSession) TagEntity(entityId *big.Int, entityTagId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.TagEntity(&_WorldErrors.TransactOpts, entityId, entityTagId)
}

// TagEntity is a paid mutator transaction binding the contract method 0x05819208.
//
// Solidity: function tagEntity(uint256 entityId, uint256 entityTagId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) TagEntity(entityId *big.Int, entityTagId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.TagEntity(&_WorldErrors.TransactOpts, entityId, entityTagId)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_WorldErrors *WorldErrorsTransactor) TransferBalanceToAddress(opts *bind.TransactOpts, fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "transferBalanceToAddress", fromNamespaceId, toAddress, amount)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_WorldErrors *WorldErrorsSession) TransferBalanceToAddress(fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.TransferBalanceToAddress(&_WorldErrors.TransactOpts, fromNamespaceId, toAddress, amount)
}

// TransferBalanceToAddress is a paid mutator transaction binding the contract method 0x45afd199.
//
// Solidity: function transferBalanceToAddress(bytes32 fromNamespaceId, address toAddress, uint256 amount) returns()
func (_WorldErrors *WorldErrorsTransactorSession) TransferBalanceToAddress(fromNamespaceId [32]byte, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.TransferBalanceToAddress(&_WorldErrors.TransactOpts, fromNamespaceId, toAddress, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_WorldErrors *WorldErrorsTransactor) TransferBalanceToNamespace(opts *bind.TransactOpts, fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "transferBalanceToNamespace", fromNamespaceId, toNamespaceId, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_WorldErrors *WorldErrorsSession) TransferBalanceToNamespace(fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.TransferBalanceToNamespace(&_WorldErrors.TransactOpts, fromNamespaceId, toNamespaceId, amount)
}

// TransferBalanceToNamespace is a paid mutator transaction binding the contract method 0xc9c85a60.
//
// Solidity: function transferBalanceToNamespace(bytes32 fromNamespaceId, bytes32 toNamespaceId, uint256 amount) returns()
func (_WorldErrors *WorldErrorsTransactorSession) TransferBalanceToNamespace(fromNamespaceId [32]byte, toNamespaceId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.TransferBalanceToNamespace(&_WorldErrors.TransactOpts, fromNamespaceId, toNamespaceId, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_WorldErrors *WorldErrorsTransactor) TransferOwnership(opts *bind.TransactOpts, namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "transferOwnership", namespaceId, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_WorldErrors *WorldErrorsSession) TransferOwnership(namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.TransferOwnership(&_WorldErrors.TransactOpts, namespaceId, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xef5d6bbb.
//
// Solidity: function transferOwnership(bytes32 namespaceId, address newOwner) returns()
func (_WorldErrors *WorldErrorsTransactorSession) TransferOwnership(namespaceId [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.TransferOwnership(&_WorldErrors.TransactOpts, namespaceId, newOwner)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactor) Unanchor(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "unanchor", entityId)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsSession) Unanchor(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.Unanchor(&_WorldErrors.TransactOpts, entityId)
}

// Unanchor is a paid mutator transaction binding the contract method 0x2d910c34.
//
// Solidity: function unanchor(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) Unanchor(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.Unanchor(&_WorldErrors.TransactOpts, entityId)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_WorldErrors *WorldErrorsTransactor) UnregisterDelegation(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "unregisterDelegation", delegatee)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_WorldErrors *WorldErrorsSession) UnregisterDelegation(delegatee common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.UnregisterDelegation(&_WorldErrors.TransactOpts, delegatee)
}

// UnregisterDelegation is a paid mutator transaction binding the contract method 0xcdc938c5.
//
// Solidity: function unregisterDelegation(address delegatee) returns()
func (_WorldErrors *WorldErrorsTransactorSession) UnregisterDelegation(delegatee common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.UnregisterDelegation(&_WorldErrors.TransactOpts, delegatee)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_WorldErrors *WorldErrorsTransactor) UnregisterNamespaceDelegation(opts *bind.TransactOpts, namespaceId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "unregisterNamespaceDelegation", namespaceId)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_WorldErrors *WorldErrorsSession) UnregisterNamespaceDelegation(namespaceId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.UnregisterNamespaceDelegation(&_WorldErrors.TransactOpts, namespaceId)
}

// UnregisterNamespaceDelegation is a paid mutator transaction binding the contract method 0xaa66e9c8.
//
// Solidity: function unregisterNamespaceDelegation(bytes32 namespaceId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) UnregisterNamespaceDelegation(namespaceId [32]byte) (*types.Transaction, error) {
	return _WorldErrors.Contract.UnregisterNamespaceDelegation(&_WorldErrors.TransactOpts, namespaceId)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_WorldErrors *WorldErrorsTransactor) UnregisterStoreHook(opts *bind.TransactOpts, tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "unregisterStoreHook", tableId, hookAddress)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_WorldErrors *WorldErrorsSession) UnregisterStoreHook(tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.UnregisterStoreHook(&_WorldErrors.TransactOpts, tableId, hookAddress)
}

// UnregisterStoreHook is a paid mutator transaction binding the contract method 0x05609129.
//
// Solidity: function unregisterStoreHook(bytes32 tableId, address hookAddress) returns()
func (_WorldErrors *WorldErrorsTransactorSession) UnregisterStoreHook(tableId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.UnregisterStoreHook(&_WorldErrors.TransactOpts, tableId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_WorldErrors *WorldErrorsTransactor) UnregisterSystemHook(opts *bind.TransactOpts, systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "unregisterSystemHook", systemId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_WorldErrors *WorldErrorsSession) UnregisterSystemHook(systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.UnregisterSystemHook(&_WorldErrors.TransactOpts, systemId, hookAddress)
}

// UnregisterSystemHook is a paid mutator transaction binding the contract method 0xa92813ad.
//
// Solidity: function unregisterSystemHook(bytes32 systemId, address hookAddress) returns()
func (_WorldErrors *WorldErrorsTransactorSession) UnregisterSystemHook(systemId [32]byte, hookAddress common.Address) (*types.Transaction, error) {
	return _WorldErrors.Contract.UnregisterSystemHook(&_WorldErrors.TransactOpts, systemId, hookAddress)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactor) UpdateFuel(opts *bind.TransactOpts, entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "updateFuel", entityId)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsSession) UpdateFuel(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.UpdateFuel(&_WorldErrors.TransactOpts, entityId)
}

// UpdateFuel is a paid mutator transaction binding the contract method 0x265f0d9a.
//
// Solidity: function updateFuel(uint256 entityId) returns()
func (_WorldErrors *WorldErrorsTransactorSession) UpdateFuel(entityId *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.UpdateFuel(&_WorldErrors.TransactOpts, entityId)
}

// Verify0 is a paid mutator transaction binding the contract method 0xafeb5022.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) returns(bool)
func (_WorldErrors *WorldErrorsTransactor) Verify0(opts *bind.TransactOpts, request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "verify0", request)
}

// Verify0 is a paid mutator transaction binding the contract method 0xafeb5022.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) returns(bool)
func (_WorldErrors *WorldErrorsSession) Verify0(request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _WorldErrors.Contract.Verify0(&_WorldErrors.TransactOpts, request)
}

// Verify0 is a paid mutator transaction binding the contract method 0xafeb5022.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) returns(bool)
func (_WorldErrors *WorldErrorsTransactorSession) Verify0(request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _WorldErrors.Contract.Verify0(&_WorldErrors.TransactOpts, request)
}

// WithdrawFromEphemeralInventory is a paid mutator transaction binding the contract method 0x2e098e36.
//
// Solidity: function withdrawFromEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactor) WithdrawFromEphemeralInventory(opts *bind.TransactOpts, smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "withdrawFromEphemeralInventory", smartObjectId, ephemeralInventoryOwner, items)
}

// WithdrawFromEphemeralInventory is a paid mutator transaction binding the contract method 0x2e098e36.
//
// Solidity: function withdrawFromEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsSession) WithdrawFromEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.WithdrawFromEphemeralInventory(&_WorldErrors.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// WithdrawFromEphemeralInventory is a paid mutator transaction binding the contract method 0x2e098e36.
//
// Solidity: function withdrawFromEphemeralInventory(uint256 smartObjectId, address ephemeralInventoryOwner, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactorSession) WithdrawFromEphemeralInventory(smartObjectId *big.Int, ephemeralInventoryOwner common.Address, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.WithdrawFromEphemeralInventory(&_WorldErrors.TransactOpts, smartObjectId, ephemeralInventoryOwner, items)
}

// WithdrawFromInventory is a paid mutator transaction binding the contract method 0x8f7513ff.
//
// Solidity: function withdrawFromInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactor) WithdrawFromInventory(opts *bind.TransactOpts, smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "withdrawFromInventory", smartObjectId, items)
}

// WithdrawFromInventory is a paid mutator transaction binding the contract method 0x8f7513ff.
//
// Solidity: function withdrawFromInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsSession) WithdrawFromInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.WithdrawFromInventory(&_WorldErrors.TransactOpts, smartObjectId, items)
}

// WithdrawFromInventory is a paid mutator transaction binding the contract method 0x8f7513ff.
//
// Solidity: function withdrawFromInventory(uint256 smartObjectId, (uint256,address,uint256,uint256,uint256,uint256)[] items) returns()
func (_WorldErrors *WorldErrorsTransactorSession) WithdrawFromInventory(smartObjectId *big.Int, items []InventoryItem) (*types.Transaction, error) {
	return _WorldErrors.Contract.WithdrawFromInventory(&_WorldErrors.TransactOpts, smartObjectId, items)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 entityId, uint256 unitAmount) returns()
func (_WorldErrors *WorldErrorsTransactor) WithdrawFuel(opts *bind.TransactOpts, entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.contract.Transact(opts, "withdrawFuel", entityId, unitAmount)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 entityId, uint256 unitAmount) returns()
func (_WorldErrors *WorldErrorsSession) WithdrawFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.WithdrawFuel(&_WorldErrors.TransactOpts, entityId, unitAmount)
}

// WithdrawFuel is a paid mutator transaction binding the contract method 0xe19a0384.
//
// Solidity: function withdrawFuel(uint256 entityId, uint256 unitAmount) returns()
func (_WorldErrors *WorldErrorsTransactorSession) WithdrawFuel(entityId *big.Int, unitAmount *big.Int) (*types.Transaction, error) {
	return _WorldErrors.Contract.WithdrawFuel(&_WorldErrors.TransactOpts, entityId, unitAmount)
}

// WorldErrorsEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the WorldErrors contract.
type WorldErrorsEIP712DomainChangedIterator struct {
	Event *WorldErrorsEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *WorldErrorsEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldErrorsEIP712DomainChanged)
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
		it.Event = new(WorldErrorsEIP712DomainChanged)
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
func (it *WorldErrorsEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldErrorsEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldErrorsEIP712DomainChanged represents a EIP712DomainChanged event raised by the WorldErrors contract.
type WorldErrorsEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WorldErrors *WorldErrorsFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*WorldErrorsEIP712DomainChangedIterator, error) {

	logs, sub, err := _WorldErrors.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &WorldErrorsEIP712DomainChangedIterator{contract: _WorldErrors.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WorldErrors *WorldErrorsFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *WorldErrorsEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _WorldErrors.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldErrorsEIP712DomainChanged)
				if err := _WorldErrors.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WorldErrors *WorldErrorsFilterer) ParseEIP712DomainChanged(log types.Log) (*WorldErrorsEIP712DomainChanged, error) {
	event := new(WorldErrorsEIP712DomainChanged)
	if err := _WorldErrors.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldErrorsExecutedForwardRequestIterator is returned from FilterExecutedForwardRequest and is used to iterate over the raw logs and unpacked data for ExecutedForwardRequest events raised by the WorldErrors contract.
type WorldErrorsExecutedForwardRequestIterator struct {
	Event *WorldErrorsExecutedForwardRequest // Event containing the contract specifics and raw log

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
func (it *WorldErrorsExecutedForwardRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldErrorsExecutedForwardRequest)
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
		it.Event = new(WorldErrorsExecutedForwardRequest)
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
func (it *WorldErrorsExecutedForwardRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldErrorsExecutedForwardRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldErrorsExecutedForwardRequest represents a ExecutedForwardRequest event raised by the WorldErrors contract.
type WorldErrorsExecutedForwardRequest struct {
	Signer  common.Address
	Nonce   *big.Int
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutedForwardRequest is a free log retrieval operation binding the contract event 0x842fb24a83793558587a3dab2be7674da4a51d09c5542d6dd354e5d0ea70813c.
//
// Solidity: event ExecutedForwardRequest(address indexed signer, uint256 nonce, bool success)
func (_WorldErrors *WorldErrorsFilterer) FilterExecutedForwardRequest(opts *bind.FilterOpts, signer []common.Address) (*WorldErrorsExecutedForwardRequestIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _WorldErrors.contract.FilterLogs(opts, "ExecutedForwardRequest", signerRule)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsExecutedForwardRequestIterator{contract: _WorldErrors.contract, event: "ExecutedForwardRequest", logs: logs, sub: sub}, nil
}

// WatchExecutedForwardRequest is a free log subscription operation binding the contract event 0x842fb24a83793558587a3dab2be7674da4a51d09c5542d6dd354e5d0ea70813c.
//
// Solidity: event ExecutedForwardRequest(address indexed signer, uint256 nonce, bool success)
func (_WorldErrors *WorldErrorsFilterer) WatchExecutedForwardRequest(opts *bind.WatchOpts, sink chan<- *WorldErrorsExecutedForwardRequest, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _WorldErrors.contract.WatchLogs(opts, "ExecutedForwardRequest", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldErrorsExecutedForwardRequest)
				if err := _WorldErrors.contract.UnpackLog(event, "ExecutedForwardRequest", log); err != nil {
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

// ParseExecutedForwardRequest is a log parse operation binding the contract event 0x842fb24a83793558587a3dab2be7674da4a51d09c5542d6dd354e5d0ea70813c.
//
// Solidity: event ExecutedForwardRequest(address indexed signer, uint256 nonce, bool success)
func (_WorldErrors *WorldErrorsFilterer) ParseExecutedForwardRequest(log types.Log) (*WorldErrorsExecutedForwardRequest, error) {
	event := new(WorldErrorsExecutedForwardRequest)
	if err := _WorldErrors.contract.UnpackLog(event, "ExecutedForwardRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldErrorsHelloStoreIterator is returned from FilterHelloStore and is used to iterate over the raw logs and unpacked data for HelloStore events raised by the WorldErrors contract.
type WorldErrorsHelloStoreIterator struct {
	Event *WorldErrorsHelloStore // Event containing the contract specifics and raw log

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
func (it *WorldErrorsHelloStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldErrorsHelloStore)
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
		it.Event = new(WorldErrorsHelloStore)
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
func (it *WorldErrorsHelloStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldErrorsHelloStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldErrorsHelloStore represents a HelloStore event raised by the WorldErrors contract.
type WorldErrorsHelloStore struct {
	StoreVersion [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHelloStore is a free log retrieval operation binding the contract event 0xc7f5fdc8526b76f54916701bc910876243ffff2a40b0bb8d59eea8151c52c005.
//
// Solidity: event HelloStore(bytes32 indexed storeVersion)
func (_WorldErrors *WorldErrorsFilterer) FilterHelloStore(opts *bind.FilterOpts, storeVersion [][32]byte) (*WorldErrorsHelloStoreIterator, error) {

	var storeVersionRule []interface{}
	for _, storeVersionItem := range storeVersion {
		storeVersionRule = append(storeVersionRule, storeVersionItem)
	}

	logs, sub, err := _WorldErrors.contract.FilterLogs(opts, "HelloStore", storeVersionRule)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsHelloStoreIterator{contract: _WorldErrors.contract, event: "HelloStore", logs: logs, sub: sub}, nil
}

// WatchHelloStore is a free log subscription operation binding the contract event 0xc7f5fdc8526b76f54916701bc910876243ffff2a40b0bb8d59eea8151c52c005.
//
// Solidity: event HelloStore(bytes32 indexed storeVersion)
func (_WorldErrors *WorldErrorsFilterer) WatchHelloStore(opts *bind.WatchOpts, sink chan<- *WorldErrorsHelloStore, storeVersion [][32]byte) (event.Subscription, error) {

	var storeVersionRule []interface{}
	for _, storeVersionItem := range storeVersion {
		storeVersionRule = append(storeVersionRule, storeVersionItem)
	}

	logs, sub, err := _WorldErrors.contract.WatchLogs(opts, "HelloStore", storeVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldErrorsHelloStore)
				if err := _WorldErrors.contract.UnpackLog(event, "HelloStore", log); err != nil {
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
func (_WorldErrors *WorldErrorsFilterer) ParseHelloStore(log types.Log) (*WorldErrorsHelloStore, error) {
	event := new(WorldErrorsHelloStore)
	if err := _WorldErrors.contract.UnpackLog(event, "HelloStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldErrorsHelloWorldIterator is returned from FilterHelloWorld and is used to iterate over the raw logs and unpacked data for HelloWorld events raised by the WorldErrors contract.
type WorldErrorsHelloWorldIterator struct {
	Event *WorldErrorsHelloWorld // Event containing the contract specifics and raw log

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
func (it *WorldErrorsHelloWorldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldErrorsHelloWorld)
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
		it.Event = new(WorldErrorsHelloWorld)
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
func (it *WorldErrorsHelloWorldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldErrorsHelloWorldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldErrorsHelloWorld represents a HelloWorld event raised by the WorldErrors contract.
type WorldErrorsHelloWorld struct {
	WorldVersion [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHelloWorld is a free log retrieval operation binding the contract event 0x7f8f36afe3fb61c459c1a54a60b8a477eab02cc58e49f547561a40906239cb82.
//
// Solidity: event HelloWorld(bytes32 indexed worldVersion)
func (_WorldErrors *WorldErrorsFilterer) FilterHelloWorld(opts *bind.FilterOpts, worldVersion [][32]byte) (*WorldErrorsHelloWorldIterator, error) {

	var worldVersionRule []interface{}
	for _, worldVersionItem := range worldVersion {
		worldVersionRule = append(worldVersionRule, worldVersionItem)
	}

	logs, sub, err := _WorldErrors.contract.FilterLogs(opts, "HelloWorld", worldVersionRule)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsHelloWorldIterator{contract: _WorldErrors.contract, event: "HelloWorld", logs: logs, sub: sub}, nil
}

// WatchHelloWorld is a free log subscription operation binding the contract event 0x7f8f36afe3fb61c459c1a54a60b8a477eab02cc58e49f547561a40906239cb82.
//
// Solidity: event HelloWorld(bytes32 indexed worldVersion)
func (_WorldErrors *WorldErrorsFilterer) WatchHelloWorld(opts *bind.WatchOpts, sink chan<- *WorldErrorsHelloWorld, worldVersion [][32]byte) (event.Subscription, error) {

	var worldVersionRule []interface{}
	for _, worldVersionItem := range worldVersion {
		worldVersionRule = append(worldVersionRule, worldVersionItem)
	}

	logs, sub, err := _WorldErrors.contract.WatchLogs(opts, "HelloWorld", worldVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldErrorsHelloWorld)
				if err := _WorldErrors.contract.UnpackLog(event, "HelloWorld", log); err != nil {
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
func (_WorldErrors *WorldErrorsFilterer) ParseHelloWorld(log types.Log) (*WorldErrorsHelloWorld, error) {
	event := new(WorldErrorsHelloWorld)
	if err := _WorldErrors.contract.UnpackLog(event, "HelloWorld", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldErrorsStoreDeleteRecordIterator is returned from FilterStoreDeleteRecord and is used to iterate over the raw logs and unpacked data for StoreDeleteRecord events raised by the WorldErrors contract.
type WorldErrorsStoreDeleteRecordIterator struct {
	Event *WorldErrorsStoreDeleteRecord // Event containing the contract specifics and raw log

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
func (it *WorldErrorsStoreDeleteRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldErrorsStoreDeleteRecord)
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
		it.Event = new(WorldErrorsStoreDeleteRecord)
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
func (it *WorldErrorsStoreDeleteRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldErrorsStoreDeleteRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldErrorsStoreDeleteRecord represents a StoreDeleteRecord event raised by the WorldErrors contract.
type WorldErrorsStoreDeleteRecord struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreDeleteRecord is a free log retrieval operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_WorldErrors *WorldErrorsFilterer) FilterStoreDeleteRecord(opts *bind.FilterOpts, tableId [][32]byte) (*WorldErrorsStoreDeleteRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _WorldErrors.contract.FilterLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsStoreDeleteRecordIterator{contract: _WorldErrors.contract, event: "Store_DeleteRecord", logs: logs, sub: sub}, nil
}

// WatchStoreDeleteRecord is a free log subscription operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_WorldErrors *WorldErrorsFilterer) WatchStoreDeleteRecord(opts *bind.WatchOpts, sink chan<- *WorldErrorsStoreDeleteRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _WorldErrors.contract.WatchLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldErrorsStoreDeleteRecord)
				if err := _WorldErrors.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
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
func (_WorldErrors *WorldErrorsFilterer) ParseStoreDeleteRecord(log types.Log) (*WorldErrorsStoreDeleteRecord, error) {
	event := new(WorldErrorsStoreDeleteRecord)
	if err := _WorldErrors.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldErrorsStoreSetRecordIterator is returned from FilterStoreSetRecord and is used to iterate over the raw logs and unpacked data for StoreSetRecord events raised by the WorldErrors contract.
type WorldErrorsStoreSetRecordIterator struct {
	Event *WorldErrorsStoreSetRecord // Event containing the contract specifics and raw log

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
func (it *WorldErrorsStoreSetRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldErrorsStoreSetRecord)
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
		it.Event = new(WorldErrorsStoreSetRecord)
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
func (it *WorldErrorsStoreSetRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldErrorsStoreSetRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldErrorsStoreSetRecord represents a StoreSetRecord event raised by the WorldErrors contract.
type WorldErrorsStoreSetRecord struct {
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
func (_WorldErrors *WorldErrorsFilterer) FilterStoreSetRecord(opts *bind.FilterOpts, tableId [][32]byte) (*WorldErrorsStoreSetRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _WorldErrors.contract.FilterLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsStoreSetRecordIterator{contract: _WorldErrors.contract, event: "Store_SetRecord", logs: logs, sub: sub}, nil
}

// WatchStoreSetRecord is a free log subscription operation binding the contract event 0x8dbb3a9672eebfd3773e72dd9c102393436816d832c7ba9e1e1ac8fcadcac7a9.
//
// Solidity: event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_WorldErrors *WorldErrorsFilterer) WatchStoreSetRecord(opts *bind.WatchOpts, sink chan<- *WorldErrorsStoreSetRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _WorldErrors.contract.WatchLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldErrorsStoreSetRecord)
				if err := _WorldErrors.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
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
func (_WorldErrors *WorldErrorsFilterer) ParseStoreSetRecord(log types.Log) (*WorldErrorsStoreSetRecord, error) {
	event := new(WorldErrorsStoreSetRecord)
	if err := _WorldErrors.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldErrorsStoreSpliceDynamicDataIterator is returned from FilterStoreSpliceDynamicData and is used to iterate over the raw logs and unpacked data for StoreSpliceDynamicData events raised by the WorldErrors contract.
type WorldErrorsStoreSpliceDynamicDataIterator struct {
	Event *WorldErrorsStoreSpliceDynamicData // Event containing the contract specifics and raw log

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
func (it *WorldErrorsStoreSpliceDynamicDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldErrorsStoreSpliceDynamicData)
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
		it.Event = new(WorldErrorsStoreSpliceDynamicData)
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
func (it *WorldErrorsStoreSpliceDynamicDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldErrorsStoreSpliceDynamicDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldErrorsStoreSpliceDynamicData represents a StoreSpliceDynamicData event raised by the WorldErrors contract.
type WorldErrorsStoreSpliceDynamicData struct {
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
func (_WorldErrors *WorldErrorsFilterer) FilterStoreSpliceDynamicData(opts *bind.FilterOpts, tableId [][32]byte) (*WorldErrorsStoreSpliceDynamicDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _WorldErrors.contract.FilterLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsStoreSpliceDynamicDataIterator{contract: _WorldErrors.contract, event: "Store_SpliceDynamicData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceDynamicData is a free log subscription operation binding the contract event 0xaa63765a776145e5e6492f471ae097dfed11cd57a61bc2679dd43180422385b4.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
func (_WorldErrors *WorldErrorsFilterer) WatchStoreSpliceDynamicData(opts *bind.WatchOpts, sink chan<- *WorldErrorsStoreSpliceDynamicData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _WorldErrors.contract.WatchLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldErrorsStoreSpliceDynamicData)
				if err := _WorldErrors.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
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
func (_WorldErrors *WorldErrorsFilterer) ParseStoreSpliceDynamicData(log types.Log) (*WorldErrorsStoreSpliceDynamicData, error) {
	event := new(WorldErrorsStoreSpliceDynamicData)
	if err := _WorldErrors.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldErrorsStoreSpliceStaticDataIterator is returned from FilterStoreSpliceStaticData and is used to iterate over the raw logs and unpacked data for StoreSpliceStaticData events raised by the WorldErrors contract.
type WorldErrorsStoreSpliceStaticDataIterator struct {
	Event *WorldErrorsStoreSpliceStaticData // Event containing the contract specifics and raw log

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
func (it *WorldErrorsStoreSpliceStaticDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldErrorsStoreSpliceStaticData)
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
		it.Event = new(WorldErrorsStoreSpliceStaticData)
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
func (it *WorldErrorsStoreSpliceStaticDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldErrorsStoreSpliceStaticDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldErrorsStoreSpliceStaticData represents a StoreSpliceStaticData event raised by the WorldErrors contract.
type WorldErrorsStoreSpliceStaticData struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Start    *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreSpliceStaticData is a free log retrieval operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_WorldErrors *WorldErrorsFilterer) FilterStoreSpliceStaticData(opts *bind.FilterOpts, tableId [][32]byte) (*WorldErrorsStoreSpliceStaticDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _WorldErrors.contract.FilterLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &WorldErrorsStoreSpliceStaticDataIterator{contract: _WorldErrors.contract, event: "Store_SpliceStaticData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceStaticData is a free log subscription operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_WorldErrors *WorldErrorsFilterer) WatchStoreSpliceStaticData(opts *bind.WatchOpts, sink chan<- *WorldErrorsStoreSpliceStaticData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _WorldErrors.contract.WatchLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldErrorsStoreSpliceStaticData)
				if err := _WorldErrors.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
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
func (_WorldErrors *WorldErrorsFilterer) ParseStoreSpliceStaticData(log types.Log) (*WorldErrorsStoreSpliceStaticData, error) {
	event := new(WorldErrorsStoreSpliceStaticData)
	if err := _WorldErrors.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
