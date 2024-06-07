package main

import (
	"ginChart/router"
	"ginChart/utils"
)

func main()  {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8081")
}
