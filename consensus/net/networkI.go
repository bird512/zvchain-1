//   Copyright (C) 2018 ZVChain
//
//   This program is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License as published by
//   the Free Software Foundation, either version 3 of the License, or
//   (at your option) any later version.
//
//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package net implements the network message handling and transmission functions
package net

import (
	"github.com/zvchain/zvchain/common"
	"github.com/zvchain/zvchain/consensus/groupsig"
	"github.com/zvchain/zvchain/consensus/model"
	"github.com/zvchain/zvchain/middleware/types"
)

// MessageProcessor interface defines the process functions of all consensus messages
type MessageProcessor interface {
	// Whether the processor is ready for handling messages
	Ready() bool

	// Returns the miner id of the processor
	GetMinerID() groupsig.ID

	// OnMessageCast handles the message from the proposer
	// Note that, if the pre-block of the block present int the message isn't on the blockchain, it will caches the message
	// and trigger it after the pre-block added on chain
	OnMessageCast(msg *model.ConsensusCastMessage)

	// OnMessageVerify handles the verification messages from other members in the group for a specified block proposal
	// Note that, it will cache the messages if the corresponding proposal message doesn't come yet and trigger them
	// as long as the condition met
	OnMessageVerify(msg *model.ConsensusVerifyMessage)

	// OnMessageCastRewardSignReq handles reward transaction signature requests
	// It signs the message if and only if the block of the transaction already added on chain,
	// otherwise the message will be cached util the condition met
	OnMessageCastRewardSignReq(msg *model.CastRewardTransSignReqMessage)

	// OnMessageCastRewardSign receives signed messages for the reward transaction from group members
	// If threshold signature received and the group signature recovered successfully,
	// the node will submit the reward transaction to the pool
	OnMessageCastRewardSign(msg *model.CastRewardTransSignMessage)

	// OnMessageReqProposalBlock handles block body request from the verify group members
	// It only happens in the proposal role and when the group signature generated by the verify-group
	OnMessageReqProposalBlock(msg *model.ReqProposalBlock, sourceID string)

	// OnMessageResponseProposalBlock handles block body response from proposal node
	// It only happens in the verify roles and after block body request to the proposal node
	// It will add the block on chain and then broadcast
	OnMessageResponseProposalBlock(msg *model.ResponseProposalBlock) string
}

// GroupBrief represents the brief info of one group including group id and member ids
type GroupBrief struct {
	GSeed  common.Hash
	MemIds []groupsig.ID
}

// NetworkServer defines some network transmission functions for various types of data.
type NetworkServer interface {

	// SendCastVerify happens at the proposal role.
	// It send the message contains the proposed-block to all of the members of the verify-group for the verification consensus
	SendCastVerify(ccm *model.ConsensusCastMessage, gb *GroupBrief, proveHashs []common.Hash)

	// SendVerifiedCast broadcast the signed message for specified block proposal among group members
	SendVerifiedCast(cvm *model.ConsensusVerifyMessage, gSeed common.Hash)

	// BroadcastNewBlock means network-wide broadcast for the generated block.
	// Based on bandwidth and performance considerations, it only transits the block to all of the proposers and
	// the next verify-group
	BroadcastNewBlock(block *types.Block, group *GroupBrief)

	// BuildGroupNet builds the group net in local for inter-group communication
	BuildGroupNet(groupIdentifier string, mems []groupsig.ID)

	// ReleaseGroupNet releases the group net in local
	ReleaseGroupNet(groupIdentifier string)

	// SendCastRewardSignReq sends reward transaction sign request to other members of the group
	SendCastRewardSignReq(msg *model.CastRewardTransSignReqMessage)

	// SendCastRewardSign sends signed message of the reward transaction to the requester by group relaying
	SendCastRewardSign(msg *model.CastRewardTransSignMessage)

	// ReqProposalBlock request block body from the target
	ReqProposalBlock(msg *model.ReqProposalBlock, target string)

	// ResponseProposalBlock sends block body to the requester
	ResponseProposalBlock(msg *model.ResponseProposalBlock, target string)
}
