package internal

import (
	"time"
)

type Entry struct {
	Id        uint64 `json:"id"`
	Base62_id string `json:"short_url"`
	LongUrl   string `json:"long_url"`
	Date_Created time.Time `json:"date_created"`
}
