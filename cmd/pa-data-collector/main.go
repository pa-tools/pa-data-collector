package main

import (
	"github.com/pa-tools/pa-data-collector/internal/api"
	"github.com/pa-tools/pa-data-collector/internal/orm"
)

func main() {
	go orm.InitDb()
	api.InitApi()
}
