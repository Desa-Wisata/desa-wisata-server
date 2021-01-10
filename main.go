package main

import (
	"github.com/desa-wisata/desa-wisata-server/delivery/handler"
	"github.com/desa-wisata/desa-wisata-server/delivery/router"
	"github.com/desa-wisata/desa-wisata-server/repository"
	"github.com/desa-wisata/desa-wisata-server/usecase"
	"github.com/desa-wisata/library/database"
	"github.com/desa-wisata/library/server"
	_ "github.com/joho/godotenv/autoload" //load env
)

func main() {

	db := database.MYSQLConnection()
	defer db.Close()

	ap := repository.InitRepository(db)
	sr := usecase.InitUsecase(ap)
	ct := handler.InitController(sr)
	rt := router.InitRoute(ct)

	server.Start(rt.Mux())
}
