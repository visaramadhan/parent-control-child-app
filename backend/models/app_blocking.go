package models

import "time"

type App struct {
    ID          uint      `gorm:"primaryKey"`
    DeviceID    uint      `gorm:"not null;index"`
    AppName     string    `gorm:"size:100"`
    PackageName string    `gorm:"size:150"`
    IsBlocked   bool      `gorm:"default:false"`
    CreatedAt   time.Time
    UpdatedAt   time.Time

    Device Device
}
