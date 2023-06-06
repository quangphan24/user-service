package user

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"user-service/model"
)

type RepoUser struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepoUser {
	return &RepoUser{db: db}
}

//go:generate mockery --name IRepoUser
type IRepoUser interface {
	CreateUser(user *model.User) error
	GetOneUserByEmail(email string) (*model.User, error)
	CreateRefreshToken(rt *model.RefreshToken) error
	GetOne(id string) (*model.User, error)
}

func (r *RepoUser) CreateRefreshToken(rt *model.RefreshToken) error {
	return r.db.Model(&model.RefreshToken{}).Create(&rt).Error
}
func (r *RepoUser) CreateUser(user *model.User) error {
	tx := r.db
	return tx.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "email"}}, UpdateAll: true}).
		Create(&user).Error
}
func (r *RepoUser) GetOneUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	tx := r.db
	if err := tx.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *RepoUser) GetOne(id string) (*model.User, error) {
	rt := &model.User{}

	tx := r.db

	tx.Where("id=?", id).Take(&rt)
	rt.ID = id
	return rt, nil
}
