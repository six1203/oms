package model

import (
	"time"
)

type Common struct {
	Id        int64     `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null;index:created_at"`
	UpdatedAt time.Time `gorm:"not null;index:updated_at"`
	IsDeleted bool      `gorm:"not null;default:false"`
}

/* 在GORM不支持数组类型变量数据的存取，可以通过GORM的自定义数据类型实现。
 	GORM的自定义的数据类型必须实现Scanner/Valuer接口。
	Scanner接口的Scan方法，是从数据库读取数据到Go变量时需要进行的解析处理，与解码的过程类型。
	Valuer接口的Value方法，是将Go变量存到数据库时进行编码处理。*/

//type JSON json.RawMessage
//
//func (j *JSON) Scan(val interface{}) error {
//	//s := val.([]uint8)
//	//ss := strings.Split(string(s), ",")
//	//*r = ss
//	//return nil
//
//	bytes, ok := val.([]byte)
//	if !ok {
//		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
//	}
//
//	result := json.RawMessage{}
//	err := json.Unmarshal(bytes, &result)
//	*j = JSON(result)
//	return err
//}
//
//func (j JSON) Value() (driver.Value, error) {
//	if len(j) == 0 {
//		return nil, nil
//	}
//	return json.RawMessage(j).MarshalJSON()
//}
