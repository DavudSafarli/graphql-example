package storage

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/stdapps/graphql-example/ticketing"
)

type PostgresStorage struct {
	db *sql.DB
	b  sq.StatementBuilderType
}

func NewPostgresStorage(db *sql.DB) PostgresStorage {
	return PostgresStorage{
		db: db,
		b:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// GetUsers ...
func (s PostgresStorage) GetUsers(p ticketing.Pagination, criteria ticketing.UsersSearchCriteria) ([]ticketing.User, error) {
	offset := p.Limit * (p.Page - 1)
	query := s.b.Select("id", "name").From("users").
		Limit(uint64(p.Limit)).
		Offset(uint64(offset))

	if criteria.Name != "" {
		query = query.Where(sq.Like{"name": criteria.Name})
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.Query(sql, args)
	defer rows.Close()

	users := []ticketing.User{}
	for rows.Next() {
		user := ticketing.User{}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users, nil
}

// FindUser ...
func (s PostgresStorage) FindUser(id int) (user ticketing.User, err error) {
	query := s.b.Select("id", "name").From("users").Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return user, err
	}
	if err != nil {
		return user, err
	}
	err = s.db.QueryRow(sql, args...).Scan(&user.ID, &user.Name)
	if err != nil {
		return user, err
	}
	return user, err
}

// CreateUser ...
func (s PostgresStorage) CreateUser(user ticketing.User) (ticketing.User, error) {
	query := s.b.Insert("users").
		Columns("name").
		Values(user.Name).
		Suffix("returning id")

	sql, args, err := query.ToSql()
	log.Printf(sql, args...)
	if err != nil {
		return user, err
	}
	if err != nil {
		return user, err
	}
	err = s.db.QueryRow(sql, args...).Scan(&user.ID)
	if err != nil {
		return user, err
	}
	return user, err
}
