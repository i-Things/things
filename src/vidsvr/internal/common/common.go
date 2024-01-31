package common

import (
	"gitee.com/i-Things/core/shared/def"
	"github.com/i-Things/things/src/vidsvr/pb/vid"
)

func ToPageInfo(info *vid.PageInfo, defaultOrders ...def.OrderBy) *def.PageInfo {
	if info == nil {
		return nil
	}

	var orders = defaultOrders
	if infoOrders := info.GetOrders(); len(infoOrders) > 0 {
		orders = make([]def.OrderBy, 0, len(infoOrders))
		for _, infoOd := range infoOrders {
			if infoOd.GetFiled() != "" {
				orders = append(orders, def.OrderBy{infoOd.GetFiled(), infoOd.GetSort()})
			}
		}
	}

	return &def.PageInfo{
		Page:   info.GetPage(),
		Size:   info.GetSize(),
		Orders: orders,
	}
}

func ToPageInfoWithDefault(info *vid.PageInfo, defau *def.PageInfo) *def.PageInfo {
	if page := ToPageInfo(info); page == nil {
		return defau
	} else {
		if page.Page == 0 {
			page.Page = defau.Page
		}
		if page.Size == 0 {
			page.Size = defau.Size
		}
		if len(page.Orders) == 0 {
			page.Orders = defau.Orders
		}
		return page
	}
}
