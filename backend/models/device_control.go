package models

import "time"

type Device struct {
    ID        uint      `gorm:"primaryKey"`
    ChildID   uint      `gorm:"not null;uniqueIndex"`
    Name      string    `gorm:"size:100"`
    OSVersion string    `gorm:"size:50"`
    Status    string    `gorm:"size:20;check:status IN ('locked','unlocked')"`
    CreatedAt time.Time
    UpdatedAt time.Time

    Child Child
    Apps  []App
}
