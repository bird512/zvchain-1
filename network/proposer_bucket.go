package network

import (
	"fmt"
	"math"
	"sync"
)

const DEFAULT_BUCKET_GROUP_SIZE = 100
const DEFAULT_BUCKET_GROUP_NAME = "BucketGroup"

type Proposer struct {
	ID    NodeID
	Stake uint64
}

type ProposerBucket struct {
	proposers  []*Proposer
	groups     []*Group
	groupSize  int
	groupCount int
	groupName  string

	mutex sync.RWMutex
}

func newProposerBucket(groupName string, groupSize int) *ProposerBucket {
	pm := &ProposerBucket{
		proposers: make([]*Proposer, 0),
		groups:    make([]*Group, 0),
		groupSize: groupSize,
		groupName: groupName,
	}
	if pm.groupSize == 0 {
		pm.groupSize = DEFAULT_BUCKET_GROUP_SIZE
	}

	if len(pm.groupName) == 0 {
		pm.groupName = DEFAULT_BUCKET_GROUP_NAME
	}

	return pm
}

func (pb *ProposerBucket) GroupNameByIndex(index int) string {
	return fmt.Sprintf("%v_%v", pb.groupName, index)
}

func (pb *ProposerBucket) Build(proposers []*Proposer) {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()

	groupCountOld := pb.groupCount
	pb.proposers = proposers

	pb.groupCount = int(math.Ceil(float64(len(proposers)) / float64(pb.groupSize)))
	if groupCountOld > pb.groupCount {
		for i := pb.groupCount; i < groupCountOld; i++ {
			netCore.groupManager.removeGroup(pb.GroupNameByIndex(i))
		}
	}
	for i := 0; i < pb.groupCount; i++ {
		pb.buildGroup(i)
	}
}

func (pb *ProposerBucket) IsContained(proposer *Proposer) bool {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()
	for i := 0; i < len(pb.proposers); i++ {
		if pb.proposers[i].ID == proposer.ID {
			return true
		}
	}
	return false
}

func (pb *ProposerBucket) AddProposers(proposers []*Proposer) {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()
	//proposersOld := pb.proposers

	pb.proposers = append(pb.proposers, proposers...)

	groupIndex := pb.groupCount - 1
	if groupIndex < 0 {
		return
	}
	pb.buildGroup(groupIndex)
}

func (pb *ProposerBucket) buildGroup(groupIndex int) {

	groupID := pb.GroupNameByIndex(groupIndex)
	members := pb.groupMembersByIndex(groupIndex)

	netCore.groupManager.buildGroup(groupID, members)
}
func (pb *ProposerBucket) groupMembersByIndex(groupIndex int) []NodeID {

	members := make([]NodeID, 0)
	startIndex := groupIndex * pb.groupSize
	endIndex := startIndex + pb.groupSize
	if endIndex > len(pb.proposers) {
		endIndex = len(pb.proposers)
	}

	for n := startIndex; n < endIndex; n++ {
		members = append(members, pb.proposers[n].ID)
	}

	return members
}

func (pb *ProposerBucket) groupMembersHexByIndex(groupIndex int) []string {

	members := make([]string, 0)
	startIndex := groupIndex * pb.groupSize
	endIndex := startIndex + pb.groupSize
	if endIndex > len(pb.proposers) {
		endIndex = len(pb.proposers)
	}

	for n := startIndex; n < endIndex; n++ {
		members = append(members, pb.proposers[n].ID.GetHexString())
	}

	return members
}

func (pb *ProposerBucket) Broadcast(msg *MsgData, code uint32) {
	if msg == nil {
		Logger.Errorf("[proposer bucket] group broadcast,msg is nil,code:%v", code)
		return
	}
	Logger.Infof("[proposer bucket] group broadcast, code:%v", code)
	pb.mutex.Lock()
	defer pb.mutex.Unlock()

	for i := 0; i < pb.groupCount; i++ {
		groupID := pb.GroupNameByIndex(i)
		members := pb.groupMembersHexByIndex(i)

		netCore.groupManager.Broadcast(groupID, msg, members, code)
	}
}
