package types

import (
	"fmt"

	btypes "github.com/QOSGroup/qbase/types"
	qtypes "github.com/QOSGroup/qos/types"
	"github.com/tendermint/tendermint/crypto"
)

type GenesisState struct {
	Params                 Params                           `json:"params"`
	Validators             []Validator                      `json:"validators"`          //validatorKey, validatorByOwnerKey,validatorByInactiveKey,validatorByVotePowerKey
	ValidatorsVoteInfo     []ValidatorVoteInfoState         `json:"val_votes_info"`      //validatorVoteInfoKey
	ValidatorsVoteInWindow []ValidatorVoteInWindowInfoState `json:"val_votes_in_window"` //validatorVoteInfoInWindowKey
	DelegatorsInfo         []DelegationInfoState            `json:"delegators_info"`     //DelegationByDelValKey, DelegationByValDelKey
	CurrentValidators      []Validator                      `json:"current_validators"`  // currentValidatorsAddressKey
}

func NewGenesisState(params Params,
	validators []Validator,
	validatorsVoteInfo []ValidatorVoteInfoState,
	validatorsVoteInWindow []ValidatorVoteInWindowInfoState,
	delegatorsInfo []DelegationInfoState,
	currentValidators []Validator) GenesisState {
	return GenesisState{
		Params:                 params,
		Validators:             validators,
		ValidatorsVoteInfo:     validatorsVoteInfo,
		ValidatorsVoteInWindow: validatorsVoteInWindow,
		DelegatorsInfo:         delegatorsInfo,
		CurrentValidators:      currentValidators,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}

func ValidateGenesis(genesisAccounts []*qtypes.QOSAccount, data GenesisState) error {
	err := validateValidators(genesisAccounts, data.Validators)
	if err != nil {
		return err
	}

	return nil
}

func validateValidators(genesisAccounts []*qtypes.QOSAccount, validators []Validator) (err error) {
	addrMap := make(map[string]bool, len(validators))
	for i := 0; i < len(validators); i++ {
		val := validators[i]
		strKey := string(val.ValidatorPubKey.Bytes())
		if _, ok := addrMap[strKey]; ok {
			return fmt.Errorf("duplicate validator in genesis state: Name %v, Owner %v", val.Description.Moniker, val.Owner)
		}
		if val.Status != Active {
			return fmt.Errorf("validator is bonded and jailed in genesis state: Name %v, Owner %v", val.Description.Moniker, val.Owner)
		}
		addrMap[strKey] = true

		var ownerExists bool
		for _, acc := range genesisAccounts {
			if acc.AccountAddress.EqualsTo(val.Owner) {
				ownerExists = true
			}
		}

		if !ownerExists {
			return fmt.Errorf("owner of %s not exists", val.Description.Moniker)
		}
	}
	return nil
}

type ValidatorVoteInfoState struct {
	ValidatorPubKey crypto.PubKey     `json:"validator_pub_key"`
	VoteInfo        ValidatorVoteInfo `json:"vote_info"`
}

type ValidatorVoteInWindowInfoState struct {
	ValidatorPubKey crypto.PubKey `json:"validator_pub_key"`
	Index           uint64        `json:"index"`
	Vote            bool          `json:"vote"`
}

type DelegationInfoState struct {
	DelegatorAddr   btypes.Address `json:"delegator_addr"`
	ValidatorPubKey crypto.PubKey  `json:"validator_pub_key"`
	Amount          uint64         `json:"delegate_amount"`
	IsCompound      bool           `json:"is_compound"`
}
