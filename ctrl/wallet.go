package ctrl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"v-factor.com/v-wallet/svc"
)

type walletCtrl struct {
	walletSvc svc.WalletSvc
}

type WalletCtrl interface {
	GetWalletById(ctx *gin.Context)
	DeductFunds(ctx *gin.Context)
	DepositFunds(ctx *gin.Context)
}

func NewWalletCtrl(walletSvc svc.WalletSvc) WalletCtrl {
	return &walletCtrl{
		walletSvc,
	}
}

func (w *walletCtrl) GetWalletById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		// TODO
		panic(err)
	}
	wallet, err := w.walletSvc.GetWalletById(int64(id))
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, wallet)
}

func (w *walletCtrl) DeductFunds(ctx *gin.Context) {
	panic("not implement")
}

func (w *walletCtrl) DepositFunds(ctx *gin.Context) {
	panic("not implement")
}
