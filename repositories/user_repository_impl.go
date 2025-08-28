package repositories

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, filter web.UserFilter) []domain.User {
	query := "select * from users"
	var args []interface{}
	var conditions []string

	if filter.Search != nil {
		conditions = append(conditions, "(name like ?)")
		args = append(args, "%"+*filter.Search+"%")
	}
	if filter.Level != nil {
		conditions = append(conditions, "level like ?")
		args = append(args, "%"+*filter.Level+"%")
	}

	if len(conditions) > 0 {
		query += " where " + strings.Join(conditions, " and ")
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	halpers.IfPanicError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Level, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		halpers.IfPanicError(err)

		users = append(users, user)
	}
	return users
}

func (u UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (domain.User, error) {
	rows, err := tx.QueryContext(ctx, "select * from users where id= ?", id)
	halpers.IfPanicError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Level, &user.CreatedAt, &user.UpdatedAt)
		halpers.IfPanicError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (u UserRepositoryImpl) FindByName(ctx context.Context, dbOrTx QueryExecutor, name string) (*domain.User, error) {
	rows, err := dbOrTx.QueryContext(ctx, "select * from users where name=?", name)
	halpers.IfPanicError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Level, &user.CreatedAt, &user.UpdatedAt)
		halpers.IfPanicError(err)
		return &user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func (u UserRepositoryImpl) FindByEmail(ctx context.Context, dbOrTx QueryExecutor, email string) (*domain.User, error) {
	rows, err := dbOrTx.QueryContext(ctx, "select * from users where email=?", email)
	halpers.IfPanicError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Level, &user.CreatedAt, &user.UpdatedAt)
		halpers.IfPanicError(err)
		return &user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func (u UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	result, err := tx.ExecContext(ctx, "insert into users (name, email, phone, level) values (?,?,?,?)", user.Name, user.Email, user.Phone, user.Level)
	halpers.IfPanicError(err)

	id, err := result.LastInsertId()
	halpers.IfPanicError(err)

	//	find by id
	rows, err := tx.QueryContext(ctx, "select * from users where id=?", id)
	halpers.IfPanicError(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Level, &user.CreatedAt, &user.UpdatedAt)
		halpers.IfPanicError(err)
	}
	return user
}

func (u UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	result, err := tx.ExecContext(ctx, "update users set name = ?, email = ?, phone = ?, level = ? where id = ?", user.Name, user.Email, user.Phone, user.Level, user.Id)
	halpers.IfPanicError(err)

	id, err := result.LastInsertId()
	halpers.IfPanicError(err)

	//	find by id
	rows, err := tx.QueryContext(ctx, "select * from users where id=?", id)
	halpers.IfPanicError(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Level, &user.CreatedAt, &user.UpdatedAt)
		halpers.IfPanicError(err)
	}
	return user
}

func (u UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int64) {
	_, err := tx.ExecContext(ctx, "delete from users where id = ?", id)
	halpers.IfPanicError(err)
}
