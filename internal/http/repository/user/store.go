package user

import (
	"database/sql"

	"github.com/dealense7/market-price-go/internal/models"
	"github.com/dealense7/market-price-go/utils"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetByEmail(email string) (*models.User, error) {
	var item = models.User{}
	err := s.db.Get(&item, "SELECT * FROM users WHERE email=?", email)

	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case when no rows are found
			return nil, utils.ErrResourceNotFound
		}
		return nil, err
	}

	return &item, nil
}

func (s *Store) GetById(id int) (*models.User, error) {
	var item = models.User{}
	err := s.db.Get(&item, "SELECT * FROM users WHERE id=?", id)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *Store) Create(user models.User) error {
	_, err := s.db.NamedExec(`INSERT INTO users (first_name, last_name, email, password) VALUES (:first_name, :last_name, :email, :password)`, user)

	return err
}
