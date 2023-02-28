package model

import "time"

const DefaultDelCondition = "deleted_at = 0"

type BaseModel struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt int64
}
