package models

import "time"

type Notification struct {
    ID        uint      `gorm:"primaryKey"`
    ParentID  uint      `gorm:"not null;index"`
    Title     string    `gorm:"size:100"`
    Message   string
    IsRead    bool      `gorm:"default:false"`
    CreatedAt time.Time

    Parent Parent
}
