// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package page

import (
	"bytes"

	"github.com/error-ident/go-knc-admin/context"
	"github.com/error-ident/go-knc-admin/modules/config"
	"github.com/error-ident/go-knc-admin/modules/db"
	"github.com/error-ident/go-knc-admin/modules/logger"
	"github.com/error-ident/go-knc-admin/modules/menu"
	"github.com/error-ident/go-knc-admin/plugins/admin/models"
	"github.com/error-ident/go-knc-admin/template"
	"github.com/error-ident/go-knc-admin/template/types"
)

// SetPageContent set and return the panel of page content.
func SetPageContent(ctx *context.Context, user models.UserModel, c func(ctx interface{}) (types.Panel, error), conn db.Connection) {

	panel, err := c(ctx)

	if err != nil {
		logger.Error("SetPageContent", err)
		panel = template.WarningPanel(err.Error())
	}

	tmpl, tmplName := template.Get(config.GetTheme()).GetTemplate(ctx.IsPjax())

	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(&types.NewPageParam{
		User:         user,
		Menu:         menu.GetGlobalMenu(user, conn, ctx.Lang()).SetActiveClass(config.URLRemovePrefix(ctx.Path())),
		Panel:        panel.GetContent(config.IsProductionEnvironment()),
		Assets:       template.GetComponentAssetImportHTML(),
		TmplHeadHTML: template.Default().GetHeadHTML(),
		TmplFootJS:   template.Default().GetFootJS(),
		Iframe:       ctx.IsIframe(),
	}))
	if err != nil {
		logger.Error("SetPageContent", err)
	}
	ctx.WriteString(buf.String())
}
