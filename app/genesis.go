package app

import (
	"fmt"
	bacc "github.com/QOSGroup/qbase/account"
	"github.com/QOSGroup/qbase/context"
	"github.com/QOSGroup/qos/module/mint"
	"github.com/QOSGroup/qos/module/qcp"
	"github.com/QOSGroup/qos/module/qsc"
	"github.com/QOSGroup/qos/module/stake"
	"github.com/QOSGroup/qos/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// QOS初始状态
type GenesisState struct {
	Accounts  []*types.QOSAccount `json:"accounts"`
	MintData  mint.GenesisState   `json:"mint"`
	StakeData stake.GenesisState  `json:"stake"`
	QCPData   qcp.GenesisState    `json:"qcp"`
	QSCData   qsc.GenesisState    `json:"qsc"`
}

func NewDefaultGenesisState() GenesisState {
	return GenesisState{
		MintData:  mint.DefaultGenesisState(),
		StakeData: stake.DefaultGenesisState(),
	}
}

func ValidGenesis(state GenesisState) error {
	if err := validateAccounts(state.Accounts); err != nil {
		return err
	}

	if err := stake.ValidateGenesis(state.Accounts, state.StakeData); err != nil {
		return err
	}

	return nil
}

func InitGenesis(ctx context.Context, state GenesisState) []abci.ValidatorUpdate {
	// accounts init should in the first
	initAccounts(ctx, state.Accounts)
	mint.InitGenesis(ctx, state.MintData)
	stake.InitGenesis(ctx, state.StakeData)
	qcp.InitGenesis(ctx, state.QCPData)
	qsc.InitGenesis(ctx, state.QSCData)

	return stake.GetUpdatedValidators(ctx, uint64(state.StakeData.Params.MaxValidatorCnt))
}

func initAccounts(ctx context.Context, accounts []*types.QOSAccount) {
	if len(accounts) == 0 {
		return
	}
	var appliedQOSAmount uint64

	accountMapper := ctx.Mapper(bacc.AccountMapperName).(*bacc.AccountMapper)
	mintMapper := ctx.Mapper(mint.MintMapperName).(*mint.MintMapper)
	for _, acc := range accounts {
		accountMapper.SetAccount(acc)
		appliedQOSAmount += uint64(acc.QOS.Int64())
	}

	mintMapper.SetAppliedQOSAmount(appliedQOSAmount)
}

func validateAccounts(accs []*types.QOSAccount) error {
	addrMap := make(map[string]bool, len(accs))
	for i := 0; i < len(accs); i++ {
		acc := accs[i]
		strAddr := string(acc.AccountAddress)
		if _, ok := addrMap[strAddr]; ok {
			return fmt.Errorf("Duplicate account in genesis state: Address %v", acc.AccountAddress)
		}
		addrMap[strAddr] = true
	}
	return nil
}
