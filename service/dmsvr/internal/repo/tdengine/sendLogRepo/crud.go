package sendLogRepo

import (
	"context"
	"database/sql"
	"fmt"
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/stores"
	sq "github.com/Masterminds/squirrel"
	"github.com/i-Things/things/service/dmsvr/internal/domain/deviceLog"
)

func (s *SendLogRepo) fillFilter(sql sq.SelectBuilder, filter deviceLog.SendFilter) sq.SelectBuilder {
	if len(filter.ProductID) != 0 {
		sql = sql.Where("`product_id`=?", filter.ProductID)
	}
	if len(filter.DeviceName) != 0 {
		sql = sql.Where("`device_name`=?", filter.DeviceName)
	}
	if len(filter.TenantCode) != 0 {
		sql = sql.Where("`tenant_code`=?", filter.TenantCode)
	}
	if filter.ProjectID != 0 {
		sql = sql.Where("`project_id` =?", filter.ProjectID)
	}
	if len(filter.AreaIDs) != 0 {
		sql = sql.Where(fmt.Sprintf("`area_id` in (%v)", stores.ArrayToSql(filter.AreaIDs)))
	}
	return sql
}

func (s *SendLogRepo) GetCountLog(ctx context.Context, filter deviceLog.SendFilter, page def.PageInfo2) (int64, error) {
	sqSql := sq.Select("Count(1)").From(s.GetLogStableName())
	sqSql = s.fillFilter(sqSql, filter)
	sqSql = page.FmtWhere(sqSql)
	sqlStr, value, err := sqSql.ToSql()
	if err != nil {
		return 0, err
	}
	row := s.t.QueryRowContext(ctx, sqlStr, value...)
	if err != nil {
		return 0, err
	}
	var (
		total int64
	)

	err = row.Scan(&total)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return total, nil
}

func (s *SendLogRepo) GetDeviceLog(ctx context.Context, filter deviceLog.SendFilter, page def.PageInfo2) (
	[]*deviceLog.Send, error) {
	sql := sq.Select("*").From(s.GetLogStableName()).OrderBy("`ts` desc")
	sql = s.fillFilter(sql, filter)
	sql = page.FmtSql(sql)
	sqlStr, value, err := sql.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := s.t.QueryContext(ctx, sqlStr, value...)
	if err != nil {
		return nil, err
	}
	var datas []map[string]any
	stores.Scan(rows, &datas)
	retLogs := make([]*deviceLog.Send, 0, len(datas))
	for _, v := range datas {
		retLogs = append(retLogs, ToDeviceLog(v))
	}
	return retLogs, nil
}

func (s *SendLogRepo) Insert(ctx context.Context, data *deviceLog.Send) error {
	sql := fmt.Sprintf(" %s (`ts`, `user_id` ,`action` ,`data_id` ,`trace_id` ,`content`,`result_code`) values (?,?,?,?,?,?,?);",
		s.GetLogTableName(data.ProductID, data.DeviceName))
	s.t.AsyncInsert(sql, data.Timestamp, data.UserID, data.Action, data.DataID, data.TraceID, data.Content, data.ResultCode)
	return nil
}