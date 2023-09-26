package svc

import "v-factor.com/v-wallet/model"

type WalletSvc interface {
	GetWalletById(walletId int64) (model.Wallet, error)
	GetWalletByUserId(userId int64) (model.Wallet, error)
	Deduct(walletId int64, amount int64) error
	Deposit(walletId int64, amount int64) error
}

type walletSvc struct {
	walletModel model.WalletModel
}

func NewWalletSvc(walletModel model.WalletModel) WalletSvc {
	return &walletSvc{
		walletModel: walletModel,
	}
}

func (w *walletSvc) GetWalletById(walletId int64) (model.Wallet, error) {
	return w.walletModel.GetById(walletId)
}

func (w *walletSvc) GetWalletByUserId(userId int64) (model.Wallet, error) {
	panic("not implemented") // TODO: Implement
}

func (w *walletSvc) Deduct(walletId int64, amount int64) error {
	panic("not implemented") // TODO: Implement
}

func (w *walletSvc) Deposit(walletId int64, amount int64) error {
	panic("not implemented") // TODO: Implement
}
