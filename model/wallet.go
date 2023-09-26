package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"v-factor.com/v-wallet/db"
)

type WalletModel interface {
	GetById(walletId int64) (Wallet, error)
	GetByUserId(userId int64) (Wallet, error)
	Deduct(id int64, amount int64) error
	Deposit(id int64, amount int64) error
	Create() (Wallet, error)
}

type wallet struct {
	db *mongo.Collection
}

type Wallet struct {
	Id       int64
	UserId   int64
	Balance  int64
	Status   int8
	Remark   string
	CreateAt int64
}

func NewWallet(db *db.Db) WalletModel {
	return &wallet{db: db.Db.Collection("wallet")}
}

func (w *wallet) GetById(walletId int64) (Wallet, error) {
	wa := Wallet{Id: walletId}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	err := w.db.FindOne(ctx, wa).Decode(&wa)
	return wa, err
}

func (w *wallet) GetByUserId(userId int64) (Wallet, error) {
	panic("not implemented") // TODO: Implement
}

func (w *wallet) Deduct(id int64, amount int64) error {
	panic("not implemented") // TODO: Implement
}

func (w *wallet) Deposit(id int64, amount int64) error {
	panic("not implemented") // TODO: Implement
}

func (w *wallet) Create() (Wallet, error) {
	panic("not implemented") // TODO: Implement
}
