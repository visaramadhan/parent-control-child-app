package models

import "time"

type EducationalGame struct {
    ID         uint      `gorm:"primaryKey"`
    ChildID    uint      `gorm:"not null;index"`
    GameName   string    `gorm:"size:100"`
    LaunchTime time.Time
    Duration   int       // dalam menit
    CreatedAt  time.Time

    Child Child
}
