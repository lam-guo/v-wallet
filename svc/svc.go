package svc

import "v-factor.com/v-wallet/model"

type Svcs struct {
	WalletSvc WalletSvc
}

func NewSvcs(models model.Models) Svcs {
	return Svcs{
		WalletSvc: NewWalletSvc(models.Wallet),
	}
}
