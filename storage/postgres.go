package storage

import (
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/stdapps/graphql-example/ticketing"
)

type PostgresStorage struct {
	db *sql.DB
	b  sq.StatementBuilderType
}

func OpenPostgresDB(connstr string) *sql.DB {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}
	return db
}

func NewPostgresStorage(db *sql.DB) PostgresStorage {
	return PostgresStorage{
		db: db,
		b:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// GetUsers search for users by their name with pagination
func (s PostgresStorage) GetUsers(p ticketing.Pagination, criteria ticketing.UsersSearchCriteria) ([]ticketing.User, error) {
	// treat page=0 as page=1
	if p.Page == 0 {
		p.Page = 1
	}
	offset := p.Limit * (p.Page - 1)
	query := s.b.Select("id", "name").From("users").
		Limit(uint64(p.Limit)).
		Offset(uint64(offset))

	if criteria.Name != "" {
		query = query.Where(sq.Like{"name": fmt.Sprintf("%%%s%%", criteria.Name)})
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
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

// FindUser finds user by its ID
func (s PostgresStorage) FindUser(id int) (user ticketing.User, err error) {
	query := s.b.Select("id", "name").From("users").Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return user, err
	}
	err = s.db.QueryRow(sql, args...).Scan(&user.ID, &user.Name)
	if err != nil {
		return user, err
	}
	return user, err
}

// CreateUser persist a user to postgres database
func (s PostgresStorage) CreateUser(user ticketing.User) (ticketing.User, error) {
	query := s.b.Insert("users").
		Columns("name").
		Values(user.Name).
		Suffix("returning id")

	sql, args, err := query.ToSql()
	if err != nil {
		return user, err
	}
	err = s.db.QueryRow(sql, args...).Scan(&user.ID)
	if err != nil {
		return user, err
	}
	return user, err
}
