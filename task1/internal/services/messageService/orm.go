package messageService

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Text string `json:"text"`
}


