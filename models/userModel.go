package models

import "time"

/*

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;unique"`
	UserName  string    `json:"username" gorm:"not null;unique" validate:"required,min=2,max=50"`
	Email     string    `json:"email" gorm:"not null;unique" validate:"email,required"`
	Password  string    `json:"password" gorm:"not null" validate:"required"`
	Otp       string    `json:"otp"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Projects []Project `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE"`
	Notes    []Note    `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Members  []Member  `gorm:"foreignKey:MemberID;constraint:OnDelete:CASCADE"`
	Bugs     []Bug     `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE"`
}

*/

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique"`
	UserName  string `json:"username" gorm:"not null;unique" validate:"required,min=2,max=50"`
	Email     string `json:"email" gorm:"not null;unique" validate:"email,required"`
	Password  string `json:"password" gorm:"not null" validate:"required"`
	Otp       string `json:"otp"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Projects []Project
	// Notes    []Note
	// Members  []Member
	// Bugs     []Bug
}
