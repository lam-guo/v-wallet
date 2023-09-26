package model

import "v-factor.com/v-wallet/db"

type Models struct {
	Wallet WalletModel
}

func NewModels(db *db.Db) Models {
	return Models{
		Wallet: NewWallet(db),
	}
}
