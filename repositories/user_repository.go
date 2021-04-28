package repositories

import (
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

type SortFilter struct {
	Column string
	Order string
}

type UserFilter struct {
}

func (ur *UserRepository) GetPaged (filter UserFilter, sorting []SortFilter, pagination PaginationFilter) ([]domain.User, int64) {
	var users []domain.User
	query := ur.db.Model(&domain.User{})

	var count int64

	var sortStr string
	for i, sort := range sorting {
		sortStr = sort.Column + " " + sort.Order
		if i < len(sorting) - 1 {
			sortStr = sortStr + ","
		}
	}
	
	if sortStr != "" {
		query = query.Order(sortStr)
	}

	query.Count(&count)

	if (PaginationFilter{} != pagination) {
		
		if pagination.PageNumber > 0 && pagination.PageSize > 0 {
			query = query.Limit(pagination.PageSize).Offset((pagination.PageNumber - 1) * pagination.PageSize)
		}
	}
	

	query.Find(&users)

	return users, count
}

func (ur *UserRepository) GetById (id uint) domain.User {
	var user domain.User
	ur.db.First(&user, id)
	return user
}

func (ur *UserRepository) GetPersonalInfo(id uint) domain.User {
	var user domain.User
	ur.db.Preload("UserQuestionRatings").Preload("UserAnswerRatings").First(&user, id)
	return user
}

func (ur *UserRepository) GetByEmail (email string) (domain.User, error) {
	var user domain.User

	if result := ur.db.Where("email = ?", email).First(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil;
}

func (ur *UserRepository) Create (user domain.User) (domain.User, error) {
	result := ur.db.Create(&user)
	return user, result.Error
}

func (ur *UserRepository) Update (user domain.User, updatedUser domain.User) (domain.User, error) {
	result := ur.db.Model(&user).Updates(domain.User{
		FirstName: updatedUser.FirstName,
		LastName: updatedUser.LastName,
		Email: updatedUser.Email,
	})

	return user, result.Error
}

func (ur *UserRepository) ChangePassword (user domain.User, updatedUser domain.User) (domain.User, error) {
	result := ur.db.Model(&user).Updates(domain.User{
		PasswordSalt: updatedUser.PasswordSalt,
		PasswordHash: updatedUser.PasswordHash,
	})

	return user, result.Error
}