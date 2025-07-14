package user

import (
	"fmt"

	"github.com/agastiya/tiyago/models"
)

type IUserRepository interface {
	BrowseUser(params BrowseUserFilter) ([]BrowseUserWithMeta, error)
	DetailUser(id int64) (*models.User, error)
	CheckUsernameExists(username string, id int64) (bool, error)
	CheckEmailExists(email string, id int64) (bool, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

func (r *UserRepository) BrowseUser(params BrowseUserFilter) ([]BrowseUserWithMeta, error) {

	var users []BrowseUserWithMeta
	db := r.PostgreDB.Table("users").
		Select(`
			users.*, 
			COUNT(*) OVER() AS total_records,
			CASE 
				WHEN CEILING(COUNT(*) OVER() / CAST(? AS FLOAT)) = (? + 1) 
				THEN TRUE 
				ELSE FALSE 
			END AS has_reach_max`, params.PageSize, params.PageNumber).
		Where("deleted_at IS NULL")

	if params.Fullname != nil {
		db = db.Where("fullname ILIKE ?", "%"+*params.Fullname+"%")
	}

	if params.Username != nil {
		db = db.Where("username ILIKE ?", "%"+*params.Username+"%")
	}

	if params.Email != nil {
		db = db.Where("email ILIKE ?", "%"+*params.Email+"%")
	}

	if params.SortColumn != "" && params.SortOrder != "" {
		order := fmt.Sprintf("%s %s", params.SortColumn, params.SortOrder)
		db = db.Order(order)
	}

	db = db.Limit(params.PageSize).Offset(params.PageNumber)

	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) DetailUser(id int64) (*models.User, error) {

	var user models.User
	err := r.PostgreDB.
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CheckUsernameExists(username string, id int64) (bool, error) {
	var exists bool
	err := r.PostgreDB.
		Raw("SELECT EXISTS(SELECT 1 FROM users WHERE username = ? and id != ?)", username, id).
		Scan(&exists).Error
	return exists, err
}

func (r *UserRepository) CheckEmailExists(email string, id int64) (bool, error) {
	var exists bool
	err := r.PostgreDB.
		Raw("SELECT EXISTS(SELECT 1 FROM users WHERE email = ? and id != ?)", email, id).
		Scan(&exists).Error
	return exists, err
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.PostgreDB.Create(&user).Error
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.PostgreDB.Where("id = ?", user.Id).Updates(user).Error
}

func (r *UserRepository) DeleteUser(user *models.User) error {
	return r.PostgreDB.Where("id = ?", user.Id).Updates(&user).Error
}
