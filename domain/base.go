package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
	DeletedAt time.Time `json:"deleted_at" valid:"-" sql:"index"`
}

func (Base *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now)

	if err != nil {
		log.Fatalf("error during obj creating %v", err)
	}

	err = scope.SetColumn("ID", uuid.NewV4().String())

	if err != nil {
		log.Fatalf("error during obj creating %v", err)
	}

	return nil
}
