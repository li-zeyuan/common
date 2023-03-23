package model

import "time"

const (
	DefaultDelCondition    = "deleted_at = 0"
	UpdatedAtDESCCondition = "updated_at desc"
	IDASCCondition         = "id asc"
)

type BaseModel struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt int64
}
