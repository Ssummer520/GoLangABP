package controllers

import . "GoLangABP/demo.Web.Host/conf"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// carInventoryHandler
// @Summary 获取最新carid
// @Tags CarInventoryController 库存相关
// @Accept application/json
// @Produce application/json
// @Router /car [get]
func GetCarHandler(c *gin.Context) {
	CarLock.Lock()
	if CarInventoryModel.List.Len() == 0 {
		CarInventoryModel.SellCount++
		CarInventoryModel.TotalCount++
		CarLock.Unlock()
		c.JSON(200, gin.H{
			"car": time.Now().UnixNano(),
		})
	} else {
		carId := CarInventoryModel.List.Front()
		CarInventoryModel.SellCount++
		CarInventoryModel.List.Remove(CarInventoryModel.List.Front())
		CarLock.Unlock()
		c.JSON(200, gin.H{
			"car": carId.Value,
		})
	}
}

// carInventoryHandler
// @Summary 获取最新售卖率
// @Tags CarInventoryController 库存相关
// @Accept application/json
// @Produce application/json
// @Router /rate [get]
func GetRateHandler(c *gin.Context) {
	CarLock.RLock()
	var r = CarInventoryModel.Rate
	CarLock.RUnlock()
	c.JSON(200, gin.H{
		"rate": strconv.FormatFloat(r, 'f', 2, 32),
	})
}

// carInventoryHandler
// @Summary 获取最新库存数量
// @Tags CarInventoryController 库存相关
// @Accept application/json
// @Produce application/json
// @Router /buffer [get]
func GetBufferHandler(c *gin.Context) {
	CarLock.RLock()
	var count = CarInventoryModel.Rate * CarInventoryModel.X
	fmt.Println(count)
	CarLock.RUnlock()
	c.JSON(200, gin.H{
		"buffer": int(count),
	})
}
