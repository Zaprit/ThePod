package model

type PermissionLevel uint

const (
	Banned        PermissionLevel = -3
	Restricted                    = -2
	Silenced                      = -1
	Default                       = 0
	Moderator                     = 1
	Administrator                 = 2
)
