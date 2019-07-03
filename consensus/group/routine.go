//   Copyright (C) 2019 ZVChain
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

package group

import (
	"fmt"
	"github.com/zvchain/zvchain/common"
	"github.com/zvchain/zvchain/consensus/groupsig"
	"github.com/zvchain/zvchain/consensus/model"
	"github.com/zvchain/zvchain/middleware/types"
)

const (
	groupMemberMin        = 80
	groupMemberMax        = 100
	threshold             = 0.51
	recvPieceMinRatio     = 0.8
	memberMaxJoinGroupNum = 5
)

func memberCountLegal(n int) bool {
	return n >= groupMemberMin && n <= groupMemberMax
}

func memberCount(totalN int) int {
	if totalN >= groupMemberMax {
		return groupMemberMax
	} else if totalN < groupMemberMin {
		return 0
	}
	return totalN
}

func candidateEnough(n int) bool {
	return n >= groupMemberMin
}

type minerReader interface {
	GetLatestVerifyMiner(id groupsig.ID) *model.MinerDO
	GetCanJoinGroupMinersAt(h uint64) []*model.MinerDO
}

type currentMinerInfoReader interface {
	MInfo() *model.SelfMinerDO
}

type createRoutine struct {
	*createChecker
	packetSender types.GroupPacketSender
}

// UpdateEra updates the era info base on current block header
func (routine *createRoutine) UpdateContext(bh *types.BlockHeader) {
	curEra := routine.currEra()
	sh := seedHeight(bh.Height)
	if curEra.seedHeight == sh {
		return
	}
	seedBH := routine.chain.QueryBlockHeaderByHeight(sh)
	routine.ctx = newCreateContext(newEra(sh, seedBH))
}

func (routine *createRoutine) selectCandidates() error {
	// Already selected
	if routine.ctx.cands != nil {
		return nil
	}

	routine.ctx.cands = make(candidates, 0)

	h := routine.currEra().seedHeight
	bh := routine.currEra().seedBlock

	allVerifiers := routine.minerReader.GetCanJoinGroupMinersAt(h)
	if !candidateEnough(len(allVerifiers)) {
		return fmt.Errorf("not enought candidate in all:%v", len(allVerifiers))
	}

	availCandidates := make([]*model.MinerDO, 0)
	groups := routine.storeReader.GetAvailableGroupInfos(h)
	memJoinedCountMap := make(map[string]int)
	for _, g := range groups {
		for _, mem := range g.Members() {
			hex := common.ToHex(mem.ID())
			if c, ok := memJoinedCountMap[hex]; !ok {
				memJoinedCountMap[hex] = 1
			} else {
				memJoinedCountMap[hex] = c + 1
			}
		}
	}

	for _, verifier := range allVerifiers {
		if c, ok := memJoinedCountMap[verifier.ID.GetHexString()]; !ok || c < memberMaxJoinGroupNum {
			availCandidates = append(availCandidates, verifier)
		}
	}
	memberCnt := memberCount(len(availCandidates))
	if !candidateEnough(memberCnt) {
		return fmt.Errorf("not enough candiates in availables:%v", len(availCandidates))
	}

	selector := newCandidateSelector(availCandidates, bh.Random)
	selectedCandidates := selector.algSatoshi(memberCnt)

	routine.ctx.cands = selectedCandidates
	return nil
}

func (routine *createRoutine) generateSharePiece(miner *model.SelfMinerDO) {
	rand := miner.GenSecretForGroup(routine.currEra().Seed())
	seck := *groupsig.NewSeckeyFromRand(rand.Deri(0))
	pk := *groupsig.NewPubkeyFromSeckey(seck)

}

func (routine *createRoutine) CheckAndSendEncryptedPieces(bh *types.BlockHeader) error {
	var err error
	era := routine.currEra()
	if !era.seedExist() {
		return fmt.Errorf("seed not exists:%v", era.seedHeight)
	}
	if !era.encPieceRange.inRange(bh.Height) {
		return fmt.Errorf("height not in the encrypted-piece round")
	}

	mInfo := routine.currentMiner.MInfo()
	if routine.hasSentEncryptedPiece(mInfo.ID) {
		return fmt.Errorf("has sent encrypted pieces")
	}
	if !mInfo.CanJoinGroup(era.seedHeight) {
		return fmt.Errorf("current miner cann't join group")
	}

}
