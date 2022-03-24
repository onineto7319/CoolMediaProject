package router

import (
	v1 "binance_project/internal/router/api/v1"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {

	e := gin.Default()

	binance := v1.NewBinance()

	v1 := e.Group("/api/v1")
	{
		v1.GET("/", binance.Get_Collection_Transactions)
	}

	return e
}
