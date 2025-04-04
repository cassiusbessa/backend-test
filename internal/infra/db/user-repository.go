package db

import (
	"github.com/cassiusbessa/backend-test/internal/application/repositories"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) repositories.UserRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Save(user entities.User) error {
	model := UserToModel(user)
	return r.db.Create(&model).Error
}

func (r *UserGormRepository) FindByEmail(email string) (*entities.User, error) {
	var model UserModel
	err := r.db.Where("email = ?", email).First(&model).Error
	if err != nil {
		return nil, err
	}

	user := model.ToDomain()
	return &user, nil
}
