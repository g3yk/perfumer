package model

import (
	"time"

	"gorm.io/gorm"
)

// Perfume Notes

type PerfumeNote struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Note string `gorm:"not null" json:"note"`
}

// Perfume struct
type Perfume struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"DeletedAt,omitempty" swaggerignore:"true"`

	Name   string `gorm:"not null" json:"name"`
	Amount *int   `json:"amount"`

	// Relations
	BaseNotes  []PerfumeNote `gorm:"many2many:perfume_base_notes;" json:"baseNotes"`
	HeartNotes []PerfumeNote `gorm:"many2many:perfume_heart_notes;" json:"heartNotes"`
	TopNotes   []PerfumeNote `gorm:"many2many:perfume_top_notes;" json:"topNotes"`
}

type Order struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"DeletedAt,omitempty" swaggerignore:"true"`

	Perfumes []Perfume `gorm:"many2many:order_perfumes;" json:"perfumes"`
	UserID   int       `gorm:"not null" json:"userId"`
}
