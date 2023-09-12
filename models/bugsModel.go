package models

import "time"

/*
type Bug struct {
	ID          uint   `json:"id" gorm:"primaryKey;unique"`
	Title       string `json:"bug_title" gorm:"not null" validate:"required"`
	Description string `json:"bug_description" gorm:"not null" validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ReopenedAt  time.Time
	IsResolved  bool

	ClosedAt time.Time
	Priority Priority `gorm:"type:enum('low', 'medium', 'high');default:'low'"`
	Project  Project  // Many-to-One relationship: Many bugs belong to one project
	//Notes      []Note     // One-to-Many relationship: A bug can have multiple notes
	//ClosedBy   User       // Many-to-One relationship: Many bugs can be closed by one user
	//ReopenedBy User       // Many-to-One relationship: Many bugs can be reopened by one user
	//CreatedBy  User       // Many-to-One relationship: Many bugs are created by one user
	//UpdatedBy  User       // Many-to-One relationship: Many bugs are updated by one user

	Notes      []Note `gorm:"foreignKey:BugID;constraint:OnDelete:CASCADE"`
	ClosedBy   User   `gorm:"foreignKey:ClosedByID;constraint:OnDelete:CASCADE"`
	ReopenedBy User   `gorm:"foreignKey:ReopenedByID;constraint:OnDelete:CASCADE"`
	CreatedBy  User   `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE"`
	UpdatedBy  User   `gorm:"foreignKey:UpdatedByID;constraint:OnDelete:CASCADE"`
}

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)


*/

/*
type Bug struct {
    ID          uint      `json:"id" gorm:"primaryKey;unique"`
    Title       string    `json:"bug_title" gorm:"not null" validate:"required"`
    Description string    `json:"bug_description" gorm:"not null" validate:"required"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    ReopenedAt  time.Time
    IsResolved  bool
    ClosedAt    time.Time
    Priority    Priority   `gorm:"type:enum('low', 'medium', 'high');default:'low'"`

    // Define foreign keys with proper references
    ProjectID   uint       `gorm:"not null"`
    ClosedByID  uint
    ReopenedByID uint
    CreatedByID uint
    UpdatedByID uint

    // Define relationships
    Project     Project    `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
    Notes       []Note     `gorm:"foreignKey:BugID;constraint:OnDelete:CASCADE"`
    ClosedBy    User       `gorm:"foreignKey:ClosedByID;constraint:OnDelete:CASCADE"`
    ReopenedBy  User       `gorm:"foreignKey:ReopenedByID;constraint:OnDelete:CASCADE"`
    CreatedBy   User       `gorm:"foreignKey:CreatedByID;constraint:OnDelete:CASCADE"`
    UpdatedBy   User       `gorm:"foreignKey:UpdatedByID;constraint:OnDelete:CASCADE"`
}


type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

*/

type Bug struct {
	ID          uint   `json:"id" gorm:"primaryKey;unique"`
	Title       string `json:"bug_title" gorm:"not null" validate:"required"`
	Description string `json:"bug_description" gorm:"not null" validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ReopenedAt  *time.Time // Make it a pointer so that it can be nullable
	IsResolved  bool

	ClosedAt *time.Time
	// Priority Priority `gorm:"type:enum('low', 'medium', 'high');default:'low'"`
	Priority Priority `gorm:"default:'low'"`

	Project   Project `gorm:"foreignKey:ProjectId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	ProjectId uint

	Notes        []Note
	ClosedBy     User  `gorm:"foreignKey:ClosedById;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	ClosedById   *uint // Make it a pointer so that it can be nullable
	ReopenedBy   User  `gorm:"foreignKey:ReopenedById;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	ReopenedById *uint // Make it a pointer so that it can be nullable
	CreatedBy    User  `gorm:"foreignKey:CreatedById;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CreatedById  *uint // Make it a pointer so that it can be nullable
	UpdatedBy    User  `gorm:"foreignKey:UpdatedById;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UpdatedById  *uint // Make it a pointer so that it can be nullable
}

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)
