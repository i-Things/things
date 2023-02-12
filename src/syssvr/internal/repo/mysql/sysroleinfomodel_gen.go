// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysRoleInfoFieldNames          = builder.RawFieldNames(&SysRoleInfo{})
	sysRoleInfoRows                = strings.Join(sysRoleInfoFieldNames, ",")
	sysRoleInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleInfoFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), ",")
	sysRoleInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleInfoFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), "=?,") + "=?"
)

type (
	sysRoleInfoModel interface {
		Insert(ctx context.Context, data *SysRoleInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysRoleInfo, error)
		FindOneByName(ctx context.Context, name string) (*SysRoleInfo, error)
		Update(ctx context.Context, data *SysRoleInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysRoleInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysRoleInfo struct {
		Id          int64        `db:"id"`          // id编号
		Name        string       `db:"name"`        // 角色名称
		Remark      string       `db:"remark"`      // 备注
		CreatedTime time.Time    `db:"createdTime"` // 创建时间
		UpdatedTime time.Time    `db:"updatedTime"` // 更新时间
		DeletedTime sql.NullTime `db:"deletedTime"`
		Status      int64        `db:"status"` // 状态  1:启用,2:禁用
	}
)

func newSysRoleInfoModel(conn sqlx.SqlConn) *defaultSysRoleInfoModel {
	return &defaultSysRoleInfoModel{
		conn:  conn,
		table: "`sys_role_info`",
	}
}

func (m *defaultSysRoleInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysRoleInfoModel) FindOne(ctx context.Context, id int64) (*SysRoleInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleInfoRows, m.table)
	var resp SysRoleInfo
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

func (m *defaultSysRoleInfoModel) FindOneByName(ctx context.Context, name string) (*SysRoleInfo, error) {
	var resp SysRoleInfo
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", sysRoleInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysRoleInfoModel) Insert(ctx context.Context, data *SysRoleInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, sysRoleInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Remark, data.Status)
	return ret, err
}

func (m *defaultSysRoleInfoModel) Update(ctx context.Context, newData *SysRoleInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.Remark, newData.Status, newData.Id)
	return err
}

func (m *defaultSysRoleInfoModel) tableName() string {
	return m.table
}