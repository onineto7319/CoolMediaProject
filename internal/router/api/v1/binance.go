package v1

import (
	"binance_project/global"
	"binance_project/internal/dao"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	binanceInstance *binance
	store_key       = "streams=btcusdt@aggTrade"
)

type binance struct {
}

func NewBinance() *binance {
	var once sync.Once

	if binanceInstance == nil {
		once.Do(func() { binanceInstance = &binance{} })
	}
	return binanceInstance
}

func (b *binance) Get_Collection_Transactions(c *gin.Context) {
	binanceDao := dao.NewBinance(global.RedisDB)

	result, err := binanceDao.Get(store_key)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
