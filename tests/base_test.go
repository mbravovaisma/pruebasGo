package tests

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/mercadolibre/pruebas/tests/mocks"

	"github.com/mercadolibre/pruebas/models"

	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
	"github.com/mercadolibre/pruebas/app"
	"github.com/mercadolibre/pruebas/services"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	app.ConfigureRouter()
}

func teardown() {
	rest.StopMockupServer()
}

func performRequest(method, target, body string, header []string) *httptest.ResponseRecorder {
	//var headers http.Header
	payload := strings.NewReader(body)
	req := httptest.NewRequest(method, target, payload)
	res := httptest.NewRecorder()
	if len(header) > 0 {
		for _, v := range header {
			h := strings.Split(v, ":")
			req.Header.Add(h[0], h[1])
		}
	}

	app.ServeHTTP(res, req)
	return res
}

func TestSyncAuditsCallingServiceOK(t *testing.T) {
	rest.StartMockupServer()

	defer rest.FlushMockups()

	mocks.MockAuditSuccess()

	user := models.SampleUser()
	currentUserData, _ := user.UserToMap()

	err := services.SaveAuditSync("get", "myuser", "users", "123456", currentUserData, nil, "123", []string{"save-audit-sync", "other-tag"})

	assert.Nil(t, err)
}
