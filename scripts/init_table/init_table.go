package main

import (
	"order/global"
	"order/model/system"
)

func main() {
	global.GetDB().AutoMigrate(&system.Order{})
}
