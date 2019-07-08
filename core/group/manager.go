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

	"github.com/vmihailenco/msgpack"
	"github.com/zvchain/zvchain/common"
	"github.com/zvchain/zvchain/middleware/types"
	"github.com/zvchain/zvchain/storage/account"
)

// Manager implements groupContextProvider in package consensus
type Manager struct {
	chain            chainReader
	checkerImpl      types.GroupCreateChecker
	storeReaderImpl  types.GroupStoreReader
	packetSenderImpl types.GroupPacketSender
	minerReaderImpl  minerReader
	poolImpl         pool
}

func (m *Manager) GetGroupStoreReader() types.GroupStoreReader {
	return m.storeReaderImpl
}

func (m *Manager) GetGroupPacketSender() types.GroupPacketSender {
	return m.packetSenderImpl
}

func (m *Manager) RegisterGroupCreateChecker(checker types.GroupCreateChecker) {
	m.checkerImpl = checker
}

func NewManager(chain chainReader, reader minerReader) Manager {
	store := NewStore(chain)
	packetSender := NewPacketSender(chain)
	gPool := newPool()
	err := gPool.initPool(chain.LatestStateDB())
	if err != nil {
		panic(fmt.Sprintf("failed to init group manager pool %v", err))
	}
	managerImpl := Manager{storeReaderImpl: store, packetSenderImpl: packetSender, poolImpl: *gPool, minerReaderImpl:reader}
	return managerImpl
}

// RegularCheck try to create group, do punishment and refresh active group
func (m *Manager) RegularCheck(db *account.AccountDB) {
	ctx := &CheckerContext{m.chain.Height()}
	m.tryCreateGroup(db, m.checkerImpl, ctx)
	m.tryDoPunish(db, m.checkerImpl, ctx)
	_ = freshActiveGroup(db)
}

func (m *Manager) tryCreateGroup(db *account.AccountDB, checker types.GroupCreateChecker, ctx types.CheckerContext) {
	createResult := checker.CheckGroupCreateResult(ctx)
	if createResult == nil {
		return
	}
	if createResult.Err() != nil {
		return
	}
	switch createResult.Code() {
	case types.CreateResultSuccess:
		_ = m.saveGroup(db, newGroup(createResult.GroupInfo()))
	case types.CreateResultMarkEvil:
		_ = markGroupFail(db, newGroup(createResult.GroupInfo()))
	case types.CreateResultFail:
		// do nothing
	}
	_ = markEvil(db, createResult.FrozenMiners())
}

func (m *Manager) tryDoPunish(db *account.AccountDB, checker types.GroupCreateChecker, ctx types.CheckerContext) {
	msg, err := checker.CheckGroupCreatePunishment(ctx)
	if err != nil {
		return
	}
	for _, p := range msg.PenaltyTarget() {
		addr := common.BytesToAddress(p)
		_, err = m.minerReaderImpl.MinerFrozen(db, addr, ctx.Height())
		if err != nil {
			//TODO panic
		}
	}
	//TODO reward
	//for _, r := range msg.RewardTarget() {
	//	addr := common.BytesToAddress(r)s

	//	m.minerReaderImpl.MinerPenalty(db, addr,ctx.Height())
	//}

}

func (m *Manager) saveGroup(db *account.AccountDB, group *Group) error {
	byteData, err := msgpack.Marshal(group)
	if err != nil {
		return err
	}

	byteHeader, err := msgpack.Marshal(group.Header().(*GroupHeader))
	if err != nil {
		return err
	}
	_ = m.poolImpl.add(db, group)
	db.SetData(common.HashToAddress(group.HeaderD.Seed()), groupDataKey, byteData)
	db.SetData(common.HashToAddress(group.HeaderD.Seed()), groupHeaderKey, byteHeader)

	return nil
}

func freshActiveGroup(db *account.AccountDB) error {

	return nil
}

func markEvil(db *account.AccountDB, frozenMiners [][]byte) error {
	if frozenMiners == nil || len(frozenMiners) == 0 {
		return nil
	}
	//TODO: call miner interface
	return nil
}

// markGroupFail mark group member should upload origin piece
func markGroupFail(db *account.AccountDB, group *Group) error {
	db.SetData(common.HashToAddress(group.HeaderD.Seed()), originPieceReqKey, []byte{1})
	return nil
}
