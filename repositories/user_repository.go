package repositories

import (
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetPaged (search requests.UserSearchRequest) []domain.User {
	var users []domain.User
	query := ur.db
	if (requests.UserSearchRequest{} != search) && search.PageNumber > 0 && search.PageSize > 0 {
		query = query.Limit(search.PageSize).Offset((search.PageNumber - 1) * search.PageSize)
	}

	query.Find(&users)

	return users
}

func (ur *UserRepository) GetById (id uint) domain.User {
	var user domain.User
	ur.db.First(&user, id)
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