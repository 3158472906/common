package test

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestClickhouse(t *testing.T) {
	dsn := "tcp://localhost:9000?database=system&debug=true"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	res := []map[string]interface{}{}

	err =  db.Table("tables").Find(&res).Error
	if err != nil {
		log.Fatal(err)
	}

	//for _, re := range res {
	//	fmt.Println(re)
	//}

}
