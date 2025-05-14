package models

import "time"

type Child struct {
    ID        uint      `gorm:"primaryKey"`
    ParentID  uint      `gorm:"not null;index"`
    Name      string    `gorm:"size:100;not null"`
    Birthdate time.Time
    DeviceID  *uint
    CreatedAt time.Time
    UpdatedAt time.Time
	FCMToken  string

    Parent         Parent
    Device         *Device
    ScreenTime     ScreenTime
    Locations      []Location
    GeoFences      []GeoFence
    ActivityLogs   []ActivityLog
    EducationalGames []EducationalGame
}
