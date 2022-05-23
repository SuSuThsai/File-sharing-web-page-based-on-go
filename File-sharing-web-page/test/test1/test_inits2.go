package main

import (
	"File-sharing-web-page/ALiServer/ALiInits"
	"File-sharing-web-page/GINServer/DataBase"
	"File-sharing-web-page/GINServer/GINInits"
	"File-sharing-web-page/GINServer/GINModels/KEYPassTime"
	"File-sharing-web-page/GINServer/Routers"
)

func main() {
	ALiInits.InitsConfig()
	ALiInits.InitsDriver()
	GINInits.InitsConfig()
	DataBase.Inits()
	go KEYPassTime.SPYOn()
	Routers.InitUserRouter()
}
