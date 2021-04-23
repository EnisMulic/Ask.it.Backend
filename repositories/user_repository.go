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
		query = query.Limit(search.PageNumber).Offset((search.PageNumber - 1) * search.PageSize)
	}

	query.Find(&users)

	return users
}

func (ur *UserRepository) GetById (id int) {

}

func (ur *UserRepository) Create (user domain.User) {

}

func (ur *UserRepository) Update (id int, user domain.User) {
	
}