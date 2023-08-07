package models

import "time"

type Request struct {
	Time time.Time `json:"time"`
}
