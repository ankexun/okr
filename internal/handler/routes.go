// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	login "okr/internal/handler/login"
	"okr/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{},
	)

	server.AddRoutes(
		[]rest.Route{},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: login.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/okr/v3"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/userInfo",
				Handler: login.UserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/okr/v3"),
	)
}
