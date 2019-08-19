package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zvchain/zvchain/browser/models"
)

const (
	Blockrewardtophight = "block_reward.top_block_height"
	LIMIT               = 20
)

func (storage *Storage) UpdateBatchAccount(accounts []*models.Account) bool {
	//fmt.Println("[Storage] add log ")
	if storage.db == nil {
		fmt.Println("[Storage] storage.db == nil")
		return false
	}

	for i := 0; i < len(accounts); i++ {
		if accounts[i] != nil {
			storage.UpdateObject(&accounts[i])
		}
	}
	return true
}

func (storage *Storage) AddBatchAccount(accounts []*models.Account) bool {
	//fmt.Println("[Storage] add log ")
	if storage.db == nil {
		fmt.Println("[Storage] storage.db == nil")
		return false
	}
	for i := 0; i < len(accounts); i++ {
		if accounts[i] != nil {
			storage.AddObjects(&accounts[i])
		}
	}
	return true
}

func (storage *Storage) GetAccountById(address string) []*models.Account {
	//fmt.Println("[Storage] add Verification ")
	if storage.db == nil {
		fmt.Println("[Storage] storage.db == nil")
		return nil
	}
	accounts := make([]*models.Account, 1, 1)
	storage.db.Where("address = ? ", address).Find(&accounts)
	return accounts
}

func (storage *Storage) GetAccountByMaxPrimaryId(maxid uint64) []*models.Account {
	//fmt.Println("[Storage] add Verification ")
	if storage.db == nil {
		fmt.Println("[Storage] storage.db == nil")
		return nil
	}
	accounts := make([]*models.Account, LIMIT, LIMIT)
	storage.db.Where("id > ? ", maxid).Limit(LIMIT).Find(&accounts)
	return accounts
}

func (storage *Storage) AddBlockRewardSystemconfig(sys *models.Sys) bool {
	hight := storage.TopBlockRewardHeight()
	if hight > 0 {
		storage.db.Model(&sys).UpdateColumn("value", gorm.Expr("value + ?", 1))
	} else {
		sys.Value = 1
		storage.AddObjects(&sys)
	}
	return true

}

func (storage *Storage) UpdateAccountByColumn(account *models.Account, attrs map[string]interface{}) bool {
	//fmt.Println("[Storage] add Verification ")
	if storage.db == nil {
		fmt.Println("[Storage] storage.db == nil")
		return false
	}
	storage.db.Model(&account).Updates(attrs)

	return true

}

// get topblockreward height
func (storage *Storage) TopBlockRewardHeight() uint64 {
	if storage.db == nil {
		return 0
	}
	sys := make([]models.Sys, 0, 1)
	storage.db.Limit(1).Where("variable = ?", Blockrewardtophight).Find(&sys)
	if len(sys) > 0 {
		storage.topBlockHigh = sys[0].Value
		return sys[0].Value
	}
	return 0
}
