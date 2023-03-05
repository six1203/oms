package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"order/global"
	"order/global/logger"
)

type AMapResponse struct {
	Geocodes *[]Regeocode
	Status   string
	Info     string
}

type Regeocode struct {
	FormattedAddress string
	Country          string
	Province         string
	City             string
	District         string
	Township         *[]string
	Adcode           string
	Street           *[]string
	Number           *[]int8
	Location         string
	Level            string
}

const aMapBase = "https://restapi.amap.com"

func Geocode(addr string) (*Regeocode, error) {
	url := fmt.Sprintf("%s/v3/geocode/geo?address=%s&key=%s", aMapBase, addr, *global.AMapKey)
	logger.Info(url)
	response, err := http.Get(url)
	if err != nil {
		logger.Errorf("请求高德地图API失败：%s", err)
		return nil, err
	}
	// 读取响应
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("高德地图读取响应失败：", err)
		return nil, err
	}
	var aMapResp AMapResponse
	err = json.Unmarshal(body, &aMapResp)
	if err != nil {
		logger.Errorf("高德地图JSON解码失败：%s，%s", body, err)
		return nil, err
	}
	if aMapResp.Status != "1" {
		logger.Errorf("高德地图请求失败：%s", aMapResp.Info)
		return nil, err
	}

	var regeocode Regeocode

	regeocode = Regeocode{
		FormattedAddress: (*aMapResp.Geocodes)[0].FormattedAddress,
		Country:          (*aMapResp.Geocodes)[0].Country,
		Province:         (*aMapResp.Geocodes)[0].Province,
		City:             (*aMapResp.Geocodes)[0].City,
		District:         (*aMapResp.Geocodes)[0].District,
		Township:         (*aMapResp.Geocodes)[0].Township,
		Adcode:           (*aMapResp.Geocodes)[0].Adcode,
		Street:           (*aMapResp.Geocodes)[0].Street,
		Number:           (*aMapResp.Geocodes)[0].Number,
		Location:         (*aMapResp.Geocodes)[0].Location,
		Level:            (*aMapResp.Geocodes)[0].Level,
	}
	return &regeocode, nil
}
