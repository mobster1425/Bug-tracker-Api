package models

import "time"

/*
type Project struct {
	ID          uint       `json:"id" gorm:"primaryKey;unique"`
	ProjectName string     `json:"project_name" gorm:"not null;unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	CreatedBy  User         // Many-to-One relationship: Many projects belong to one user
	//Members    []Member     // One-to-Many relationship: A project can have multiple members
	//Bugs       []Bug        // One-to-Many relationship: A project can have multiple bugs
	Members []Member `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	Bugs    []Bug    `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`

}

*/

type Project struct {
	ID          uint   `json:"id" gorm:"primaryKey;unique"`
	ProjectName string `json:"project_name" gorm:"not null;unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	CreatedBy User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId    uint
	//Members    []Member     // One-to-Many relationship: A project can have multiple members
	//Bugs       []Bug        // One-to-Many relationship: A project can have multiple bugs
	Members []Member
	Bugs    []Bug
}
