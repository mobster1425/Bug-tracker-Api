package models

import "time"

/*
type Member struct {
	ID        uint      `json:"id" gorm:"primaryKey;unique"`
	JoinedAt  time.Time

//	Project Project        // Many-to-One relationship: Many members belong to one project
//	Member  User           // Many-to-One relationship: Many members are users

Project   Project    `gorm:"constraint:OnDelete:CASCADE"` // Add cascading delete for project
	Member    User       `gorm:"constraint:OnDelete:CASCADE"` // Add cascading delete for user

	CreatedAt time.Time
	UpdatedAt time.Time
}



*/

type Member struct {
	ID       uint `json:"id" gorm:"primaryKey;unique"`
	JoinedAt time.Time

	//	Project Project        // Many-to-One relationship: Many members belong to one project
	//	Member  User           // Many-to-One relationship: Many members are users

	Project   Project `gorm:"foreignKey:ProjectId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	ProjectId uint
	Member    User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId    uint

	CreatedAt time.Time
	UpdatedAt time.Time
}
