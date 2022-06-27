// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByUsername(ctx context.Context, username string) (*Users, error)
		Update(ctx context.Context, newData *Users) error
		TransUpdate(ctx context.Context, newData *Users, session sqlx.Session) error
		Delete(ctx context.Context, id int64) error
		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		TransInsert(ctx context.Context, data *Users, session sqlx.Session) (sql.Result, error) //事务insert
	}

	defaultUsersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Users struct {
		Id            int64          `db:"id"`            // 自增ID
		Username      string         `db:"username"`      // 用户名
		Password      string         `db:"password"`      // 登录密码
		Encrypt       string         `db:"encrypt"`       // 密码加密因子
		Email         string         `db:"email"`         // 电子邮件
		Mobile        string         `db:"mobile"`        // 手机号码
		FirstName     string         `db:"first_name"`    // first_name
		LastName      string         `db:"last_name"`     // last_name
		Birthday      sql.NullTime   `db:"birthday"`      // 生日日期
		RegisterIp    string         `db:"register_ip"`   // 注册IP
		Status        int64          `db:"status"`        // 状态 -1删除；0禁用；1正常
		Code          string         `db:"code"`          // 邀请码
		Avatar        sql.NullString `db:"avatar"`        // 头像
		Nickname      sql.NullString `db:"nickname"`      // 昵称
		GoogleSecret  sql.NullString `db:"google_secret"` // 谷歌验证密钥
		CreatedAt     sql.NullTime   `db:"created_at"`
		UpdatedAt     sql.NullTime   `db:"updated_at"`
		IsMainAccount int64          `db:"is_main_account"` // 是否是主账号 0否 1是
	}
)

func newUsersModel(conn sqlx.SqlConn) *defaultUsersModel {
	return &defaultUsersModel{
		conn:  conn,
		table: "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	var resp Users
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByUsername(ctx context.Context, username string) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Encrypt, data.Email, data.Mobile, data.FirstName, data.LastName, data.Birthday, data.RegisterIp, data.Status, data.Code, data.Avatar, data.Nickname, data.GoogleSecret, data.CreatedAt, data.UpdatedAt, data.IsMainAccount)
	return ret, err
}

//TransInsert 重写insert方法 使其满足事务使用同一个会话
func (m *defaultUsersModel) TransInsert(ctx context.Context, data *Users, session sqlx.Session) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.Username, data.Password, data.Encrypt, data.Email, data.Mobile, data.FirstName, data.LastName, data.Birthday, data.RegisterIp, data.Status, data.Code, data.Avatar, data.Nickname, data.GoogleSecret, data.CreatedAt, data.UpdatedAt, data.IsMainAccount)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Username, newData.Password, newData.Encrypt, newData.Email, newData.Mobile, newData.FirstName, newData.LastName, newData.Birthday, newData.RegisterIp, newData.Status, newData.Code, newData.Avatar, newData.Nickname, newData.GoogleSecret, newData.CreatedAt, newData.UpdatedAt, newData.IsMainAccount, newData.Id)
	return err
}

func (m *defaultUsersModel) TransUpdate(ctx context.Context, newData *Users, session sqlx.Session) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
	_, err := session.ExecCtx(ctx, query, newData.Username, newData.Password, newData.Encrypt, newData.Email, newData.Mobile, newData.FirstName, newData.LastName, newData.Birthday, newData.RegisterIp, newData.Status, newData.Code, newData.Avatar, newData.Nickname, newData.GoogleSecret, newData.CreatedAt, newData.UpdatedAt, newData.IsMainAccount, newData.Id)
	return err
}

// TransCtx 将此方法暴露给logic 层 这样就不用在model层 调用事务
func (m *defaultUsersModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}