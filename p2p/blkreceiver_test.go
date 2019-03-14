/*
 * @file
 * @copyright defined in aergo/LICENSE.txt
 */

package p2p

import (
	"testing"
	"time"

	"github.com/aergoio/aergo/message"
	"github.com/aergoio/aergo/p2p/p2pmocks"
	"github.com/aergoio/aergo/p2p/subproto"
	"github.com/aergoio/aergo/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBlocksChunkReceiver_StartGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputHashes := make([]message.BlockHash, len(sampleBlks))
	for i, hash := range sampleBlks {
		inputHashes[i] = hash
	}
	tests := []struct {
		name  string
		input []message.BlockHash
		ttl   time.Duration
	}{
		{"TSimple", inputHashes, time.Millisecond * 10},
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//mockContext := new(mockContext)
			mockActor := p2pmocks.NewMockActorService(ctrl)
			//mockActor.On("SendRequest", message.P2PSvc, mock.AnythingOfType("*types.GetBlock"))
			//mockActor.On("TellRequest", message.SyncerSvc, mock.AnythingOfType("*types.GetBlock"))
			mockMF := p2pmocks.NewMockMoFactory(ctrl)
			mockMo := createDummyMo(ctrl)
			mockMF.EXPECT().NewMsgBlockRequestOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockMo)
			mockPeer := p2pmocks.NewMockRemotePeer(ctrl)
			mockPeer.EXPECT().MF().Return(mockMF)
			mockPeer.EXPECT().SendMessage(mockMo).Times(1)

			expire := time.Now().Add(test.ttl)
			br := NewBlockReceiver(mockActor, mockPeer, test.input, test.ttl)

			br.StartGet()

			assert.Equal(t, len(test.input), len(br.blockHashes))
			assert.False(t, expire.After(br.timeout))
		})
	}
}

func TestBlocksChunkReceiver_ReceiveResp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	blkNo := uint64(100)
	prevHash := dummyBlockHash
	inputHashes := make([]message.BlockHash, len(sampleBlks))
	inputBlocks := make([]*types.Block, len(sampleBlks))
	for i, hash := range sampleBlks {
		inputHashes[i] = hash
		inputBlocks[i] = &types.Block{Hash: hash, Header: &types.BlockHeader{PrevBlockHash: prevHash, BlockNo: blkNo}}
		blkNo++
		prevHash = hash
	}
	tests := []struct {
		name        string
		input       []message.BlockHash
		ttl         time.Duration
		blkInterval time.Duration
		blkInput    [][]*types.Block

		// to verify
		consumed  int
		sentResp  int
		respError bool
	}{
		{"TSingleResp", inputHashes, time.Minute, 0, [][]*types.Block{inputBlocks}, 1, 1, false},
		{"TMultiResp", inputHashes, time.Minute, 0, [][]*types.Block{inputBlocks[:1], inputBlocks[1:3], inputBlocks[3:]}, 1, 1, false},
		// Fail1 remote err
		{"TRemoteFail", inputHashes, time.Minute, 0, [][]*types.Block{inputBlocks[:0]}, 1, 1, true},
		// server didn't sent last parts. and it is very similar to timeout
		//{"TNotComplete", inputHashes, time.Minute,0,[][]*types.Block{inputBlocks[:2]},1,0, false},
		// Fail2 missing some blocks
		{"TMissingBlk", inputHashes, time.Minute, 0, [][]*types.Block{inputBlocks[:1], inputBlocks[2:3], inputBlocks[3:]}, 1, 1, true},
		// Fail3 unexpected block
		{"TDupBlock", inputHashes, time.Minute, 0, [][]*types.Block{inputBlocks[:2], inputBlocks[1:3], inputBlocks[3:]}, 1, 1, true},
		// Fail4 response sent after timeout
		{"TTimeout", inputHashes, time.Millisecond * 10, time.Millisecond * 20, [][]*types.Block{inputBlocks[:1], inputBlocks[1:3], inputBlocks[3:]}, 1, 0, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//mockContext := new(mockContext)
			mockActor := p2pmocks.NewMockActorService(ctrl)
			if test.sentResp > 0 {
				mockActor.EXPECT().TellRequest(message.SyncerSvc, gomock.Any()).
					DoAndReturn(func(a string, arg *message.GetBlockChunksRsp) {
						if !((arg.Err != nil) == test.respError) {
							t.Fatalf("Wrong error (have %v)\n", arg.Err)
						}
					}).Times(test.sentResp)
			}

			mockMF := p2pmocks.NewMockMoFactory(ctrl)
			mockMo := createDummyMo(ctrl)
			mockMF.EXPECT().NewMsgBlockRequestOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockMo)
			mockPeer := p2pmocks.NewMockRemotePeer(ctrl)
			mockPeer.EXPECT().ID().Return(dummyPeerID).AnyTimes()
			mockPeer.EXPECT().MF().Return(mockMF)
			mockPeer.EXPECT().SendMessage(gomock.Any()).Times(1)
			mockPeer.EXPECT().ConsumeRequest(gomock.Any()).Times(test.consumed) //mock.AnythingOfType("p2pcommon.MsgID"))

			//expire := time.Now().Add(test.ttl)
			br := NewBlockReceiver(mockActor, mockPeer, test.input, test.ttl)
			br.StartGet()

			msg := &V030Message{subProtocol: subproto.GetBlocksResponse, id: sampleMsgID}
			for i, blks := range test.blkInput {
				if test.blkInterval > 0 {
					time.Sleep(test.blkInterval)
				}
				body := &types.GetBlockResponse{Blocks: blks, HasNext: i < len(test.blkInput)-1}
				br.ReceiveResp(msg, body)
				if br.finished {
					break
				}
			}

		})
	}
}
