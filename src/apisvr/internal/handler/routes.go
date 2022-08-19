// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	systemuser "github.com/i-Things/things/src/apisvr/internal/handler/system/user"
	thingsdeviceauth "github.com/i-Things/things/src/apisvr/internal/handler/things/device/auth"
	thingsdeviceinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/device/info"
	thingsdeviceinteract "github.com/i-Things/things/src/apisvr/internal/handler/things/device/interact"
	thingsdevicemsg "github.com/i-Things/things/src/apisvr/internal/handler/things/device/msg"
	thingsproductinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/product/info"
	thingsproductschema "github.com/i-Things/things/src/apisvr/internal/handler/things/product/schema"
	"github.com/i-Things/things/src/apisvr/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/core/create",
				Handler: systemuser.CoreCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/info/create",
				Handler: systemuser.InfoCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/captcha",
				Handler: systemuser.CaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/core/index",
				Handler: systemuser.CoreIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: systemuser.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/info/update",
					Handler: systemuser.InfoUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: systemuser.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info/delete",
					Handler: systemuser.InfoDeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: thingsdeviceauth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/access",
				Handler: thingsdeviceauth.AccessHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/root-check",
				Handler: thingsdeviceauth.RootCheckHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/auth"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/hub-log/index",
				Handler: thingsdevicemsg.HubLogIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sdk-log/index",
				Handler: thingsdevicemsg.SdkLogIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/schema-log/index",
				Handler: thingsdevicemsg.SchemaLogIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/schema-latest/index",
				Handler: thingsdevicemsg.SchemaLatestIndexHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/msg"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: thingsdeviceinfo.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: thingsdeviceinfo.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: thingsdeviceinfo.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/index",
				Handler: thingsdeviceinfo.IndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/read",
				Handler: thingsdeviceinfo.ReadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/info"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/send-action",
				Handler: thingsdeviceinteract.SendActionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/send-property",
				Handler: thingsdeviceinteract.SendPropertyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/interact"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: thingsproductinfo.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: thingsproductinfo.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: thingsproductinfo.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/index",
				Handler: thingsproductinfo.IndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/read",
				Handler: thingsproductinfo.ReadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/product/info"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/schema/update",
				Handler: thingsproductschema.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/schema/read",
				Handler: thingsproductschema.ReadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/product/schema"),
	)
}
