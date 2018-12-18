/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */

package system

import (
	"bytes"
	"encoding/gob"
	"errors"
	"math/big"
	"sort"

	"github.com/aergoio/aergo/internal/common"
	"github.com/aergoio/aergo/internal/enc"
	"github.com/aergoio/aergo/state"
	"github.com/aergoio/aergo/types"
	"github.com/mr-tron/base58"
)

var votingkey = []byte("voting")
var totalkey = []byte("totalvote")
var sortedlistkey = []byte("sortedlist")

const PeerIDLength = 39
const VotingDelay = 5

func voting(txBody *types.TxBody, sender *state.V, scs *state.ContractState, blockNo types.BlockNo) error {
	staked, err := getStaking(scs, sender.ID())
	if err != nil {
		return err
	}
	if staked.GetWhen()+VotingDelay > blockNo {
		return types.ErrLessTimeHasPassed
	}
	staked.When = blockNo
	err = setStaking(scs, sender.ID(), staked)
	if err != nil {
		return err
	}
	oldvote, err := getVote(scs, sender.ID())
	if err != nil {
		return err
	}

	voteResult, err := loadVoteResult(scs)
	if err != nil {
		return err
	}

	for offset := 0; offset < len(oldvote.Candidate); offset += PeerIDLength {
		key := oldvote.Candidate[offset : offset+PeerIDLength]
		voteResult[base58.Encode(key)] = new(big.Int).Sub(voteResult[base58.Encode(key)], oldvote.GetAmountBigInt())
	}

	if txBody.Payload[0] != 'v' { //called from unstaking
		oldvote.Amount = staked.GetAmount()
		err = setVote(scs, sender.ID(), oldvote)
		if err != nil {
			return err
		}
		for offset := 0; offset < len(oldvote.Candidate); offset += PeerIDLength {
			key := oldvote.Candidate[offset : offset+PeerIDLength]
			voteResult[base58.Encode(key)] = new(big.Int).Add(voteResult[base58.Encode(key)], staked.GetAmountBigInt())
		}
	} else {
		if staked.GetAmountBigInt().Cmp(new(big.Int).SetUint64(0)) == 0 {
			return types.ErrMustStakeBeforeVote
		}
		vote := &types.Vote{Candidate: txBody.Payload[1:], Amount: staked.GetAmount()}
		err = setVote(scs, sender.ID(), vote)
		if err != nil {
			return err
		}
		for offset := 0; offset < len(txBody.Payload[1:]); offset += PeerIDLength {
			key := txBody.Payload[offset+1 : offset+PeerIDLength+1]

			if voteResult[base58.Encode(key)] == nil {
				voteResult[base58.Encode(key)] = new(big.Int).SetUint64(0)
			}

			voteResult[base58.Encode(key)] = new(big.Int).Add(voteResult[base58.Encode(key)], staked.GetAmountBigInt())
		}
	}

	err = syncVoteResult(scs, voteResult)
	if err != nil {
		return err
	}
	return nil
}

//GetVote return amount, to, err
func GetVote(scs *state.ContractState, voter []byte) (*types.Vote, error) {
	return getVote(scs, voter)
}

func getVote(scs *state.ContractState, voter []byte) (*types.Vote, error) {
	key := append(votingkey, voter...)
	data, err := scs.GetData(key)
	if err != nil {
		return nil, err
	}
	var vote types.Vote
	if len(data) != 0 {
		dec := gob.NewDecoder(bytes.NewBuffer(data))
		err = dec.Decode(&vote)
		if err != nil {
			return nil, err
		}
	}

	return &vote, nil
}

func setVote(scs *state.ContractState, voter []byte, vote *types.Vote) error {
	var data bytes.Buffer
	key := append(votingkey, voter...)
	enc := gob.NewEncoder(&data)
	err := enc.Encode(vote)
	if err != nil {
		return err
	}
	return scs.SetData(key, data.Bytes())
}

func loadVoteResult(scs *state.ContractState) (map[string]*big.Int, error) {
	voteResult := map[string]*big.Int{}
	data, err := scs.GetData(sortedlistkey)
	if err != nil {
		return nil, err
	}
	if len(data) != 0 {
		dec := gob.NewDecoder(bytes.NewBuffer(data))
		var voteList types.VoteList
		err = dec.Decode(&voteList)
		if err != nil {
			return nil, err
		}
		for _, v := range voteList.GetVotes() {
			voteResult[base58.Encode(v.Candidate)] = v.GetAmountBigInt()
		}
	}
	return voteResult, nil
}

func InitVoteResult(scs *state.ContractState, voteResult map[string]*big.Int) error {
	if voteResult == nil {
		return errors.New("Invalid argument : voteReult should not nil")
	}
	return syncVoteResult(scs, voteResult)
}

func syncVoteResult(scs *state.ContractState, voteResult map[string]*big.Int) error {
	voteList := buildVoteList(voteResult)

	//logger.Info().Msgf("VOTE set list %v", voteList.Votes)
	data, err := common.GobEncode(voteList)
	if err != nil {
		return err
	}
	return scs.SetData(sortedlistkey, data)
}

// BuildOrderedCandidates returns a candidate list ordered by votes.xs
func BuildOrderedCandidates(vote map[string]*big.Int) []string {
	// TODO: cleanup
	l := buildVoteList(vote)
	bps := make([]string, 0, len(l.Votes))
	for _, v := range l.Votes {
		bp := enc.ToString(v.Candidate)
		bps = append(bps, bp)
	}
	return bps
}

// BuildVoteList builds and returns a voteList type obejct from vote.
func buildVoteList(vote map[string]*big.Int) *types.VoteList {
	var voteList types.VoteList
	for k, v := range vote {
		c, _ := enc.ToBytes(k)
		vote := &types.Vote{
			Candidate: c,
			Amount:    v.Bytes(),
		}
		voteList.Votes = append(voteList.Votes, vote)
	}
	sort.Sort(sort.Reverse(voteList))

	return &voteList
}

// AccountStateReader is an interface for getting a system account state.
type AccountStateReader interface {
	GetSystemAccountState() (*state.ContractState, error)
}

// GetVoteResult returns the top n voting result from the system account state.
func GetVoteResult(ar AccountStateReader, n int) (*types.VoteList, error) {
	scs, err := ar.GetSystemAccountState()
	if err != nil {
		return nil, err
	}
	return getVoteResult(scs, n)
}

func getVoteResult(scs *state.ContractState, n int) (*types.VoteList, error) {
	data, err := scs.GetData(sortedlistkey)
	if err != nil {
		return nil, err
	}

	voteList := &types.VoteList{}
	err = common.GobDecode(data, voteList)
	if err != nil {
		return nil, err
	}
	if n < len(voteList.Votes) {
		voteList.Votes = voteList.Votes[:n]
	}
	return voteList, nil
}
