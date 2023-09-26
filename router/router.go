package router

import (
	"github.com/gin-gonic/gin"
	"v-factor.com/v-wallet/ctrl"
)

func NewRouter(ctrls ctrl.Ctrls) *gin.Engine {
	r := gin.New()

	api := r.Group("/api")
	v1 := api.Group("/v1")

	wallet := v1.Group("/wallet")
	wallet.GET("", ctrls.Wallet.GetWalletById)
	return r
}
