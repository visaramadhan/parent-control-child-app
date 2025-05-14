package models

import "time"

type Parent struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"size:100;not null"`
    Email       string    `gorm:"size:100;uniqueIndex;not null"`
    Password    string    `gorm:"size:255;not null"`
    PhoneNumber string    `gorm:"size:20"`
    FirebaseUID string    `gorm:"size:128;uniqueIndex"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Children    []Child   `gorm:"constraint:OnDelete:CASCADE;"`
}
