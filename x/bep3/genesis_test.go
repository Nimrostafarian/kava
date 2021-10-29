package bep3_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/kava-labs/kava/app"
	"github.com/kava-labs/kava/x/bep3/keeper"
	"github.com/kava-labs/kava/x/bep3/types"
)

type GenesisTestSuite struct {
	suite.Suite

	app    app.TestApp
	ctx    sdk.Context
	keeper keeper.Keeper
	addrs  []sdk.AccAddress
}

func (suite *GenesisTestSuite) SetupTest() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	tApp := app.NewTestApp()
	suite.ctx = tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	suite.keeper = tApp.GetBep3Keeper()
	suite.app = tApp

	_, addrs := app.GeneratePrivKeyAddressPairs(3)
	suite.addrs = addrs
}

func (suite *GenesisTestSuite) TestGenesisState() {

	type GenState func() app.GenesisState

	testCases := []struct {
		name       string
		genState   GenState
		expectPass bool
	}{
		{
			name: "default",
			genState: func() app.GenesisState {
				return NewBep3GenStateMulti(suite.addrs[0].String())
			},
			expectPass: true,
		},
		{
			name: "import atomic swaps and asset supplies",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				_, addrs := app.GeneratePrivKeyAddressPairs(2)
				var swaps []types.AtomicSwap
				var supplies []types.AssetSupply
				for i := 0; i < 2; i++ {
					swap, supply := loadSwapAndSupply(addrs[i].String(), i)
					swaps = append(swaps, swap)
					supplies = append(supplies, supply)
				}
				gs.AtomicSwaps = swaps
				gs.Supplies = supplies
				return app.GenesisState{"bep3": types.ModuleCdc.MustMarshalJSON(&gs)}
			},
			expectPass: true,
		},
		{
			name: "0 deputy fees",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				gs.Params.AssetParams[0].FixedFee = sdk.ZeroInt()
				return app.GenesisState{"bep3": types.ModuleCdc.MustMarshalJSON(&gs)}
			},
			expectPass: true,
		},
		{
			name: "incoming supply doesn't match amount in incoming atomic swaps",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				_, addrs := app.GeneratePrivKeyAddressPairs(1)
				swap, _ := loadSwapAndSupply(addrs[0].String(), 2)
				gs.AtomicSwaps = []types.AtomicSwap{swap}
				return app.GenesisState{"bep3": types.ModuleCdc.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "current supply above limit",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				assetParam, _ := suite.keeper.GetAsset(suite.ctx, "bnb")
				gs.Supplies = []types.AssetSupply{
					{
						IncomingSupply: c("bnb", 0),
						OutgoingSupply: c("bnb", 0),
						CurrentSupply:  c("bnb", assetParam.SupplyLimit.Limit.Add(i(1)).Int64()),
					},
				}
				return app.GenesisState{"bep3": types.ModuleCdc.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "incoming supply above limit",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				// Set up overlimit amount
				assetParam, _ := suite.keeper.GetAsset(suite.ctx, "bnb")
				overLimitAmount := assetParam.SupplyLimit.Limit.Add(i(1))

				// Set up an atomic swap with amount equal to the currently asset supply
				_, addrs := app.GeneratePrivKeyAddressPairs(2)
				timestamp := ts(0)
				randomNumber, _ := types.GenerateSecureRandomNumber()
				randomNumberHash := types.CalculateRandomHash(randomNumber[:], timestamp)
				swap := types.NewAtomicSwap(cs(c("bnb", overLimitAmount.Int64())), randomNumberHash,
					types.DefaultMinBlockLock, timestamp, suite.addrs[0].String(), addrs[1].String(), TestSenderOtherChain,
					TestRecipientOtherChain, 0, types.SWAP_STATUS_OPEN, true, types.SWAP_DIRECTION_INCOMING)
				gs.AtomicSwaps = []types.AtomicSwap{swap}

				// Set up asset supply with overlimit current supply
				gs.Supplies = []types.AssetSupply{
					{
						IncomingSupply: c("bnb", assetParam.SupplyLimit.Limit.Add(i(1)).Int64()),
						OutgoingSupply: c("bnb", 0),
						CurrentSupply:  c("bnb", 0),
					},
				}
				return app.GenesisState{"bep3": types.ModuleCdc.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "incoming supply + current supply above limit",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				// Set up overlimit amount
				assetParam, _ := suite.keeper.GetAsset(suite.ctx, "bnb")
				halfLimit := assetParam.SupplyLimit.Limit.Int64() / 2
				overHalfLimit := halfLimit + 1

				// Set up an atomic swap with amount equal to the currently asset supply
				_, addrs := app.GeneratePrivKeyAddressPairs(2)
				timestamp := ts(0)
				randomNumber, _ := types.GenerateSecureRandomNumber()
				randomNumberHash := types.CalculateRandomHash(randomNumber[:], timestamp)
				swap := types.NewAtomicSwap(cs(c("bnb", halfLimit)), randomNumberHash,
					uint64(360), timestamp, suite.addrs[0].String(), addrs[1].String(), TestSenderOtherChain,
					TestRecipientOtherChain, 0, types.SWAP_STATUS_OPEN, true, types.SWAP_DIRECTION_INCOMING)
				gs.AtomicSwaps = []types.AtomicSwap{swap}

				// Set up asset supply with overlimit current supply
				gs.Supplies = []types.AssetSupply{
					{
						IncomingSupply: c("bnb", halfLimit),
						OutgoingSupply: c("bnb", 0),
						CurrentSupply:  c("bnb", overHalfLimit),
					},
				}
				return app.GenesisState{"bep3": types.ModuleCdc.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "outgoing supply above limit",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				// Set up overlimit amount
				assetParam, _ := suite.keeper.GetAsset(suite.ctx, "bnb")
				overLimitAmount := assetParam.SupplyLimit.Limit.Add(i(1))

				// Set up an atomic swap with amount equal to the currently asset supply
				_, addrs := app.GeneratePrivKeyAddressPairs(2)
				timestamp := ts(0)
				randomNumber, _ := types.GenerateSecureRandomNumber()
				randomNumberHash := types.CalculateRandomHash(randomNumber[:], timestamp)
				swap := types.NewAtomicSwap(cs(c("bnb", overLimitAmount.Int64())), randomNumberHash,
					types.DefaultMinBlockLock, timestamp, addrs[1].String(), suite.addrs[0].String(), TestSenderOtherChain,
					TestRecipientOtherChain, 0, types.SWAP_STATUS_OPEN, true, types.SWAP_DIRECTION_OUTGOING)
				gs.AtomicSwaps = []types.AtomicSwap{swap}

				// Set up asset supply with overlimit current supply
				gs.Supplies = []types.AssetSupply{
					{
						IncomingSupply: c("bnb", 0),
						OutgoingSupply: c("bnb", 0),
						CurrentSupply:  c("bnb", assetParam.SupplyLimit.Limit.Add(i(1)).Int64()),
					},
				}
				return app.GenesisState{"bep3": types.ModuleCdc.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "asset supply denom is not a supported asset",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				gs.Supplies = []types.AssetSupply{
					{
						IncomingSupply: c("fake", 0),
						OutgoingSupply: c("fake", 0),
						CurrentSupply:  c("fake", 0),
					},
				}
				return app.GenesisState{"bep3": types.ModuleCdc.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "atomic swap asset type is unsupported",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				_, addrs := app.GeneratePrivKeyAddressPairs(2)
				timestamp := ts(0)
				randomNumber, _ := types.GenerateSecureRandomNumber()
				randomNumberHash := types.CalculateRandomHash(randomNumber[:], timestamp)
				swap := types.NewAtomicSwap(cs(c("fake", 500000)), randomNumberHash,
					uint64(360), timestamp, suite.addrs[0].String(), addrs[1].String(), TestSenderOtherChain,
					TestRecipientOtherChain, 0, types.SWAP_STATUS_OPEN, true, types.SWAP_DIRECTION_INCOMING)

				gs.AtomicSwaps = []types.AtomicSwap{swap}
				return app.GenesisState{"bep3": types.ModuleCdc.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "atomic swap status is invalid",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				_, addrs := app.GeneratePrivKeyAddressPairs(2)
				timestamp := ts(0)
				randomNumber, _ := types.GenerateSecureRandomNumber()
				randomNumberHash := types.CalculateRandomHash(randomNumber[:], timestamp)
				swap := types.NewAtomicSwap(cs(c("bnb", 5000)), randomNumberHash,
					uint64(360), timestamp, suite.addrs[0].String(), addrs[1].String(), TestSenderOtherChain,
					TestRecipientOtherChain, 0, types.SWAP_STATUS_UNSPECIFIED, true, types.SWAP_DIRECTION_INCOMING)

				gs.AtomicSwaps = []types.AtomicSwap{swap}
				return app.GenesisState{"bep3": types.ModuleCdc.LegacyAmino.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "minimum block lock cannot be > maximum block lock",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				gs.Params.AssetParams[0].MinBlockLock = 201
				gs.Params.AssetParams[0].MaxBlockLock = 200
				return app.GenesisState{"bep3": types.ModuleCdc.LegacyAmino.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "empty supported asset denom",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				gs.Params.AssetParams[0].Denom = ""
				return app.GenesisState{"bep3": types.ModuleCdc.LegacyAmino.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "negative supported asset limit",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				gs.Params.AssetParams[0].SupplyLimit.Limit = i(-100)
				return app.GenesisState{"bep3": types.ModuleCdc.LegacyAmino.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
		{
			name: "duplicate supported asset denom",
			genState: func() app.GenesisState {
				gs := baseGenState(suite.addrs[0].String())
				gs.Params.AssetParams[1].Denom = "bnb"
				return app.GenesisState{"bep3": types.ModuleCdc.LegacyAmino.MustMarshalJSON(&gs)}
			},
			expectPass: false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.expectPass {
				suite.NotPanics(func() {
					suite.app.InitializeFromGenesisStates(tc.genState())
				}, tc.name)
			} else {
				suite.Panics(func() {
					suite.app.InitializeFromGenesisStates(tc.genState())
				}, tc.name)
			}
		})

	}
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
