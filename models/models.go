package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ExpenseCategory struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Expense struct {
	ID                uint      `gorm:"primaryKey"`
	UserID            uint      `gorm:"not null"`
	ExpenseCategoryID uint      `gorm:"not null"`
	Amount            float64   `gorm:"not null"`
	Date              time.Time `gorm:"not null"`
	Note              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"unique;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &ExpenseCategory{}, &Expense{}, &RefreshToken{})
}
