package models

import "time"

type GeoFence struct {
    ID             uint      `gorm:"primaryKey"`
    ChildID        uint      `gorm:"not null;index"`
    CenterLatitude  float64
    CenterLongitude float64
    Radius         int       // dalam meter
    CreatedAt      time.Time
    UpdatedAt      time.Time

    Child Child
}
