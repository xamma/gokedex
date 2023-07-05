package models

import (
	"time"
)

type Pokemon struct {
    Name       string    `json:"name"`
    Type       string    `json:"type"`
    CaughtDate time.Time `json:"caught_date"`
}