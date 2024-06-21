package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Album model
type Album struct {
	gorm.Model
	ID     string    `gorm:"type:uuid;primaryKey" json:"id"`
	Title  string    `gorm:"type:varchar(100);not null" json:"title"`
	Artist string    `gorm:"type:varchar(100);not null" json:"artist"`
	Date   time.Time `gorm:"type:date" json:"date"`
	Users  []*User   `gorm:"many2many:user_albums;" json:"users"`
}

// BeforeSave hook to enforce character limit on Title and Artist
func (a *Album) BeforeSave(tx *gorm.DB) (err error) {
	if len(a.Title) > 100 {
		return errors.New("title exceeds 100 characters")
	}
	if len(a.Artist) > 100 {
		return errors.New("artist exceeds 100 characters")
	}
	return nil
}

// User model
type User struct {
	gorm.Model
	Email    string   `gorm:"not null; unique" json:"email"`
	Username string   `gorm:"not null; unique" json:"username"`
	Password string   `gorm:"not null" json:"password"`
	Albums   []*Album `gorm:"many2many:user_albums;" json:"albums"`
}
