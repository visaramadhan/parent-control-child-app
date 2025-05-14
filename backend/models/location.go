package models

import "time"

type Location struct {
    ID        uint      `gorm:"primaryKey"`
    ChildID   uint      `gorm:"not null;index"`
    Latitude  float64
    Longitude float64
    Timestamp time.Time `gorm:"autoCreateTime"`

    Child Child
}
