package main

import (
	"github.com/Gergenus/config"
	"github.com/Gergenus/internal/server"
	"github.com/Gergenus/pkg"
)

func main() {
	z := config.GetConfig()
	zv := pkg.NewPostgresDatabase(z)

	aniki := server.NewEchoServer(zv, z)
	aniki.InitializationRouts()
	aniki.Start()

}
