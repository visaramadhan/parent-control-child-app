package models

import "time"

type ScreenTime struct {
	ID         uint `gorm:"primaryKey"`
	ChildID    uint `gorm:"not null;uniqueIndex"`
	DailyLimit int  // dalam menit
	StartTime  time.Time
	EndTime    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Child []Child
}
