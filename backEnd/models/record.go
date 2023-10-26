package models

import (
	"time"
)

type Record struct {
	// pk(auto): primary key auto increment. 
	Id      int `pk(auto)`
	Time    time.Time
	Message string
}

func InsertRecord(str string){
	o:=db.New()
	m:=Record{
		Id:      0,
		Time:    time.Now(),
		Message: str,
	}

	o.Create(&m)
}