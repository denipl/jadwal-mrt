package main

import (
	"github.com/denipl/jadwal-mrt.git/module/station"
	"github.com/gin-gonic/gin"
)
func main() {
	initrouter()
}

func initrouter() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api/")
	)

	station.Initiate(api)
	router.Run(":8080")
	print("Server is running on port localhost:8080")
}