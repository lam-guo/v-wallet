package ctrl

import "v-factor.com/v-wallet/svc"

func NewCtrl(svcs svc.Svcs) Ctrls {
	return Ctrls{
		Wallet: NewWalletCtrl(svcs.WalletSvc),
	}
}

type Ctrls struct {
	Wallet WalletCtrl
}
