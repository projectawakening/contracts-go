package test

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/projectawakening/contracts-go/v2/bindings/world"
	"github.com/projectawakening/contracts-go/v2/bindings/worldErrors"
	"github.com/projectawakening/contracts-go/v2/utils/objectid"
	"github.com/projectawakening/contracts-go/v2/utils/txerr"
	"math/big"
	"math/rand/v2"
	"testing"
)

func TestCreateSmartCharacter(t *testing.T) {
	ctx := context.Background()

	userWallet, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	smartCharWallet := crypto.PubkeyToAddress(*userWallet.Public().(*ecdsa.PublicKey))

	errAbi, err := worldErrors.WorldErrorsMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}

	worldAbi, err := world.WorldMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}

	txErr := txerr.NewDecoder(errAbi, worldAbi)

	// `TENANT` in the .env
	objectidGen, err := objectid.NewGenerator("TEST")
	if err != nil {
		t.Fatal(err)
	}

	// foundry address.
	ethClient, err := ethclient.DialContext(ctx, "http://localhost:8546")
	if err != nil {
		t.Fatal(err)
	}

	// world address in the logs of the world-deployer.
	w, err := world.NewWorld(common.HexToAddress("0x0165878a594ca255338adfa4d48449f69242eb8f"), ethClient)
	if err != nil {
		t.Fatal(err)
	}

	// private key in the docker-compose.
	adminPrivateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		t.Fatal(err)
	}

	// chain id in the logs of the world-deployer.
	txOpts, err := bind.NewKeyedTransactorWithChainID(adminPrivateKey, big.NewInt(31337))
	if err != nil {
		t.Fatal(err)
	}

	// any int, should match `ItemId` in the `CreateCharacter` tx
	characterItemID := rand.Int64()

	smartCharID, err := objectidGen.Generate(characterItemID)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("creating character with id %s and address %s", smartCharID.String(), smartCharWallet.String())

	transaction, err := w.EvefrontierCreateCharacter(
		txOpts,
		smartCharID,
		smartCharWallet,
		big.NewInt(22), // tribe id
		world.EntityRecordParams{
			TenantId: objectidGen.GetTenant(),
			TypeId:   big.NewInt(42000000100), // `CHARACTER_TYPE_ID` in the .env
			ItemId:   big.NewInt(characterItemID),
			Volume:   big.NewInt(0),
		},
		world.EntityMetadataParams{
			Name:        "vayan",
			DappURL:     "",
			Description: "",
		},
	)
	if err != nil {
		t.Fatal(txErr.Decode(err))
	}

	_, err = bind.WaitMined(ctx, ethClient, transaction)
	if err != nil {
		t.Fatal(txErr.Decode(err))
	}

	networkNodeItemID := rand.Int64()

	networkNodeSmartID, err := objectidGen.Generate(networkNodeItemID)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("creating network node with id %s", smartCharID.String())

	transaction, err = w.EvefrontierCreateAndAnchorNetworkNode(
		txOpts,
		world.CreateAndAnchorParams{
			SmartObjectId: networkNodeSmartID,
			AssemblyType:  "NWN",
			EntityRecordParams: world.EntityRecordParams{
				TenantId: objectidGen.GetTenant(),
				TypeId:   big.NewInt(88092), // `NETWORK_NODE_TYPE_ID` in the .env
				ItemId:   big.NewInt(networkNodeItemID),
				Volume:   big.NewInt(100),
			},
			Owner: smartCharWallet,
			LocationData: world.LocationData{
				SolarSystemId: big.NewInt(0),
				X:             big.NewInt(0),
				Y:             big.NewInt(0),
				Z:             big.NewInt(0),
			},
		},
		world.FuelParams{
			FuelMaxCapacity:       big.NewInt(100),
			FuelBurnRateInSeconds: big.NewInt(3600),
		},
		big.NewInt(200),
		big.NewInt(0),
	)
	if err != nil {
		t.Fatal(txErr.Decode(err))
	}

	_, err = bind.WaitMined(ctx, ethClient, transaction)
	if err != nil {
		t.Fatal(txErr.Decode(err))
	}

	fuelSmartObjectId, err := objectidGen.Generate(77818) // one of FUEL_TYPE_ID in .env
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("depositing fuel %s", fuelSmartObjectId.String())

	transaction, err = w.EvefrontierDepositFuel(
		txOpts,
		networkNodeSmartID,
		fuelSmartObjectId,
		big.NewInt(3),
	)

	if err != nil {
		t.Fatal(txErr.Decode(err))
	}

	_, err = bind.WaitMined(ctx, ethClient, transaction)
	if err != nil {
		t.Fatal(txErr.Decode(err))
	}

	t.Logf("start burn on network node %s", networkNodeSmartID.String())

	transaction, err = w.EvefrontierStartBurn(
		txOpts,
		networkNodeSmartID,
	)

	if err != nil {
		t.Fatal(txErr.Decode(err))
	}

	_, err = bind.WaitMined(ctx, ethClient, transaction)
	if err != nil {
		t.Fatal(txErr.Decode(err))
	}

	result, err := w.EvefrontierGetCurrentFuelConsumptionStatus(
		&bind.CallOpts{Pending: false, Context: ctx},
		networkNodeSmartID,
	)
	if err != nil {
		t.Fatal(txErr.Decode(err))
	}

	t.Logf("%+v", result)

	if result.FuelAmount.Cmp(big.NewInt(0)) <= 0 {
		t.Fatal("incorrect fuel amount")
	}

}
