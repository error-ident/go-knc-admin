package chi

import (
	"net/http"
	"testing"

	"github.com/error-ident/go-knc-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestChi(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}