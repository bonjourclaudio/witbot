package main

import (
	"github.com/claudioontheweb/witbot/config"
	"github.com/claudioontheweb/witbot/pkg/web"
)

func main() {

	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	server := web.NewServer(*conf)

	err = server.Run()
	if err != nil {
		panic(err)
	}

}