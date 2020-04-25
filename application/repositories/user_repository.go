package repositories

import (
	"log"

	"github.com/andrewmak/poc_golang/domain"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {

	err := user.Prepare()

	if err != nil {
		log.Fatalf("error during the user validation: %v", err)
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error to persis user %v", err)
		return user, err
	}

	return user, nil
}
