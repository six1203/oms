package main

import (
	"order/global"
	"order/model"
)

func main() {
	global.GetDB().AutoMigrate(&model.Order{})
}
