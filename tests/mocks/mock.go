package mocks

import (
	"net/http"

	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
)

//MockAuditSuccess mock success requests
func MockAuditSuccess() {
	httpHeaders := make(http.Header)
	mock := &rest.Mock{
		URL:          "http://localhost:8080/audits/localApp/audits-test",
		RespHTTPCode: http.StatusCreated,
		ReqHeaders:   httpHeaders,
		ReqBody:      "",
	}
	rest.AddMockups(mock)

	//You need the mock by the name of the application too, otherwise the tests in the release process pipeline will fail.
	//Take a look at the Audits SDK documentation.
	mock1 := &rest.Mock{
		URL:          "http://localhost:8080/audits/pruebas/audits-test",
		HTTPMethod:   http.MethodPost,
		RespHTTPCode: http.StatusCreated,
		ReqHeaders:   httpHeaders,
		ReqBody:      "",
	}
	rest.AddMockups(mock1)
}

/*
//MockAuditFail mock fail requests
func MockAuditFail() {
	httpHeaders := make(http.Header)
	mock := &rest.Mock{
		URL:          "http://localhost:8080/audits/localApp/audits-test",
		HTTPMethod:   http.MethodPost,
		RespHTTPCode: http.StatusBadRequest,
		ReqHeaders:   httpHeaders,
		ReqBody:      "",
	}
	rest.AddMockups(mock)

	//You need the mock by the name of the application too, otherwise the tests in the release process pipeline will fail.
	//Take a look at the Audits SDK documentation.
	mock1 := &rest.Mock{
		URL:          "http://localhost:8080/audits/sample-audits-go/audits-test",
		HTTPMethod:   http.MethodPost,
		RespHTTPCode: http.StatusBadRequest,
		ReqHeaders:   httpHeaders,
		ReqBody:      "",
	}
	rest.AddMockups(mock1)
}
*/
