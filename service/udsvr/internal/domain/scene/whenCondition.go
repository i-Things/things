// Package scene 触发条件
package scene

import (
	"context"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
)

type Conditions struct {
	Terms []Term       `json:"terms,omitempty"`
	Type  TermCondType `json:"type"`
}

type TermCondType string

const (
	TermConditionTypeOr  TermCondType = "or"
	TermConditionTypeAnd TermCondType = "and"
)

type TermColumnType = string

const (
	TermColumnTypeProperty TermColumnType = "property"
	TermColumnTypeEvent    TermColumnType = "event"
	//TermColumnTypeReportTime TermColumnType = "reportTime"
	TermColumnTypeTime    TermColumnType = "time"
	TermColumnTypeWeather TermColumnType = "weather"
)

type Term struct {
	Order      int64          `json:"order"`
	ColumnType TermColumnType `json:"columnType"`         //字段类型 property:属性 weather:天气
	Property   *TermProperty  `json:"property,omitempty"` //属性类型
	Weather    *TermWeather   `json:"weather,omitempty"`
	Time       *TermTime      `json:"time,omitempty"`
	Status     Status         `json:"status"` // 状态（1启用 2禁用 3异常）
	IsAbnormal bool           `json:"isAbnormal"`
	Reason     string         `json:"reason"` //异常情况的描述说明
}

func (t *Conditions) Validate(repo CheckRepo) error {
	if t == nil {
		return nil
	}
	if len(t.Terms) > 1 && !utils.SliceIn(t.Type, TermConditionTypeOr, TermConditionTypeAnd) {
		return errors.Parameter.AddMsg("填写多个条件需要填写条件类型:" + string(t.Type))
	}
	for _, v := range t.Terms {
		v.Status = StatusNormal
		v.Reason = ""

		err := v.Validate(repo)
		if err != nil {
			if !repo.UpdateType {
				return err
			}
			v.IsAbnormal = true
			v.Reason = err.Error()
			repo.Info.IsAbnormal = true
			repo.Info.Reason = err.Error()
			return nil
		}
	}
	return nil
}

func (t *Term) Validate(repo CheckRepo) error {
	if t == nil {
		return nil
	}
	if !utils.SliceIn(t.ColumnType, TermColumnTypeProperty, TermColumnTypeEvent, TermColumnTypeTime, TermColumnTypeWeather) {
		return errors.Parameter.AddMsg("条件类型不支持:" + string(t.ColumnType))
	}
	switch t.ColumnType {
	case TermColumnTypeProperty:
		if t.Property == nil {
			return errors.Parameter.AddMsg("需要填写条件属性详情")
		}
		if err := t.Property.Validate(repo); err != nil {
			return err
		}
	case TermColumnTypeWeather:
		if t.Weather == nil {
			return errors.Parameter.AddMsg("需要填写条件天气详情")
		}
		if err := t.Weather.Validate(repo); err != nil {
			return err
		}
	case TermColumnTypeTime:
		if t.Time == nil {
			return errors.Parameter.AddMsg("需要填写条件时间详情")
		}
		if err := t.Time.Validate(repo); err != nil {
			return err
		}
	}
	return nil
}

// 判断条件是否成立
func (t *Conditions) IsHit(ctx context.Context, repo CheckRepo) bool {
	if len(t.Terms) == 0 {
		return true
	}
	var finalIsHit = false
	for _, v := range t.Terms {
		isHit := v.IsHit(ctx, repo)
		if isHit && t.Type == TermConditionTypeOr {
			return true
		}
		//if v.Status == StatusAbnormal {
		//	return false
		//}
		//如果没有命中又是or条件,或者命中了但是是and条件,则需要继续判断
		finalIsHit = isHit //如果是or,每个都返回false那就是false
	}
	return finalIsHit
}

func (t *Term) IsHit(ctx context.Context, repo CheckRepo) bool {
	var isHit bool
	switch t.ColumnType {
	case TermColumnTypeProperty:
		isHit = t.Property.IsHit(ctx, t.ColumnType, repo, t)
	case TermColumnTypeWeather:
		isHit = t.Weather.IsHit(ctx, repo)
	}
	return isHit
}
