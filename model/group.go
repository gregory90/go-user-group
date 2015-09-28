package model

import (
	"time"
)

type Group struct {
	UID       string    `json:"uid,omitempty"`
	Name      string    `json:"name,omitempty"`
	UserUID   string    `json:"-"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
