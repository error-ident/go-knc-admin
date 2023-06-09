package example

import (
	"github.com/error-ident/go-knc-admin/context"
	"github.com/error-ident/go-knc-admin/modules/auth"
	"github.com/error-ident/go-knc-admin/modules/db"
	"github.com/error-ident/go-knc-admin/modules/service"
)

func (e *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), e.TestHandler)

	return app
}
