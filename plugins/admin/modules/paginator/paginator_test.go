package paginator

import (
	"testing"

	"github.com/error-ident/go-knc-admin/modules/config"
	"github.com/error-ident/go-knc-admin/plugins/admin/modules/parameter"
	_ "github.com/error-ident/knc-themes/sword"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})
	Get(Config{
		Size:         105,
		Param:        parameter.BaseParam().SetPage("7"),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}
