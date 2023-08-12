package model

import "time"

type User struct {
	UserID          int64           `json:"userId,omitempty"`
	Username        string          `json:"username,omitempty"`
	YayHash         string          `json:"yayHash,omitempty"`
	BooHash         string          `json:"booHash,omitempty"`
	MehHash         string          `json:"mehHash,omitempty"`
	LastLogin       time.Time       `json:"lastLogin"`
	PermissionLevel PermissionLevel `json:"permissionLevel,omitempty"`
}
