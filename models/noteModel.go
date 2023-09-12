package models

import "time"

/*

type Note struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique"`
	Body      string `json:"note_body" gorm:"not null" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time

	//	Author User            // Many-to-One relationship: Many notes are authored by one user
	// Bug    Bug             // Many-to-One relationship: Many notes can be associated with one bug
	Author User `gorm:"constraint:OnDelete:CASCADE"` // Add cascading delete for author (user)
	Bug    Bug  `gorm:"constraint:OnDelete:CASCADE"` // Add cascading delete for associated bug
}


*/

type Note struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique"`
	Body      string `json:"note_body" gorm:"not null" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time

	//	Author User            // Many-to-One relationship: Many notes are authored by one user
	// Bug    Bug             // Many-to-One relationship: Many notes can be associated with one bug
	Author User  `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
UserId uint
	Bug    Bug   `gorm:"foreignKey:BugId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
BugId uint 
}
