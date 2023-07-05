package models

import (
	"time"
)

type Pokemon struct {
	Name string
	Type string
	CaughtDate time.Time
}