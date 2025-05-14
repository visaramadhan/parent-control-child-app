package models

import "time"

type ActivityLog struct {
    ID            uint      `gorm:"primaryKey"`
    ChildID       uint      `gorm:"not null;index"`
    ActivityType  string    `gorm:"size:50"` // contoh: 'call', 'sms', 'app_usage'
    ActivityDetail string
    Timestamp     time.Time `gorm:"autoCreateTime"`

    Child Child
}
