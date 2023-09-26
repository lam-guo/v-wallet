package main

import (
	"fmt"
	"os"

	"v-factor.com/v-wallet/ctrl"
	"v-factor.com/v-wallet/db"
	"v-factor.com/v-wallet/model"
	"v-factor.com/v-wallet/router"
	"v-factor.com/v-wallet/svc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbUri := os.Getenv("DATABASE_URI")
	if dbUri == "" {
		dbUri = "mongodb://admin:123456@localhost:27017"
	}
	database := os.Getenv("DATABSE")
	if database == "" {
		database = "wallet"
	}
	db, err := db.NewDB(dbUri, database)
	if err != nil {
		panic(fmt.Sprintf("new Db err :%s", err))
	}
	models := model.NewModels(db)
	svcs := svc.NewSvcs(models)
	ctrls := ctrl.NewCtrl(svcs)

	r := router.NewRouter(ctrls)

	r.Run(":" + port)
}
