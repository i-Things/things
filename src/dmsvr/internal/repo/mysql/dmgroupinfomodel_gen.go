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
	dmGroupInfoFieldNames          = builder.RawFieldNames(&DmGroupInfo{})
	dmGroupInfoRows                = strings.Join(dmGroupInfoFieldNames, ",")
	dmGroupInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(dmGroupInfoFieldNames, "`createdTime`", "`deletedTime`", "`updatedTime`"), ",")
	dmGroupInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(dmGroupInfoFieldNames, "`groupID`", "`createdTime`", "`deletedTime`", "`updatedTime`"), "=?,") + "=?"
)

type (
	dmGroupInfoModel interface {
		Insert(ctx context.Context, data *DmGroupInfo) (sql.Result, error)
		FindOne(ctx context.Context, groupID int64) (*DmGroupInfo, error)
		FindOneByGroupName(ctx context.Context, groupName string) (*DmGroupInfo, error)
		Update(ctx context.Context, data *DmGroupInfo) error
		Delete(ctx context.Context, groupID int64) error
	}

	defaultDmGroupInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DmGroupInfo struct {
		GroupID     int64        `db:"groupID"`     // 分组ID
		ParentID    int64        `db:"parentID"`    // 父组ID 0-根组
		ProjectID   int64        `db:"projectID"`   // 项目ID(雪花ID)
		GroupName   string       `db:"groupName"`   // 分组名称
		Desc        string       `db:"desc"`        // 描述
		Tags        string       `db:"tags"`        // 分组标签
		CreatedTime time.Time    `db:"createdTime"` // 创建时间
		UpdatedTime time.Time    `db:"updatedTime"` // 更新时间
		DeletedTime sql.NullTime `db:"deletedTime"` // 删除时间
	}
)

func newDmGroupInfoModel(conn sqlx.SqlConn) *defaultDmGroupInfoModel {
	return &defaultDmGroupInfoModel{
		conn:  conn,
		table: "`dm_group_info`",
	}
}

func (m *defaultDmGroupInfoModel) Delete(ctx context.Context, groupID int64) error {
	query := fmt.Sprintf("delete from %s where `groupID` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, groupID)
	return err
}

func (m *defaultDmGroupInfoModel) FindOne(ctx context.Context, groupID int64) (*DmGroupInfo, error) {
	query := fmt.Sprintf("select %s from %s where `groupID` = ? limit 1", dmGroupInfoRows, m.table)
	var resp DmGroupInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, groupID)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDmGroupInfoModel) FindOneByGroupName(ctx context.Context, groupName string) (*DmGroupInfo, error) {
	var resp DmGroupInfo
	query := fmt.Sprintf("select %s from %s where `groupName` = ? limit 1", dmGroupInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, groupName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDmGroupInfoModel) Insert(ctx context.Context, data *DmGroupInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, dmGroupInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.GroupID, data.ParentID, data.ProjectID, data.GroupName, data.Desc, data.Tags)
	return ret, err
}

func (m *defaultDmGroupInfoModel) Update(ctx context.Context, newData *DmGroupInfo) error {
	query := fmt.Sprintf("update %s set %s where `groupID` = ?", m.table, dmGroupInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.ParentID, newData.ProjectID, newData.GroupName, newData.Desc, newData.Tags, newData.GroupID)
	return err
}

func (m *defaultDmGroupInfoModel) tableName() string {
	return m.table
}
