package opslogic

import (
	"gitee.com/i-Things/share/stores"
	"github.com/i-Things/things/service/udsvr/internal/repo/relationDB"
	"github.com/i-Things/things/service/udsvr/pb/ud"
)

func ToOpsWorkOrderPo(in *ud.OpsWorkOrder) *relationDB.UdOpsWorkOrder {
	if in == nil {
		return nil
	}
	return &relationDB.UdOpsWorkOrder{
		ID:          in.Id,
		AreaID:      stores.AreaID(in.AreaID),
		RaiseUserID: in.RaiseUserID,
		IssueDesc:   in.IssueDesc,
		Number:      in.Number,
		Type:        in.Type,
		Params:      in.Params,
		Status:      in.Status,
	}
}