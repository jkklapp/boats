package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/typenil/go-docker/backend/src/api"
	"gopkg.in/gin-gonic/gin.v1"
)

func sendRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var _ = Describe("Routes", func() {
	const (
		port    = "8000"
		baseURL = "http://localhost:" + port
	)

	var (
		router        *gin.Engine
		requestResult *httptest.ResponseRecorder
	)

	router = gin.Default()
	router = api.InitRoutes(router)
	go router.Run(port)

	Describe("Testing Route Endpoints", func() {
		Context(baseURL+"/api/v1/healthcheck", func() {

			requestResult = sendRequest(router, "GET", baseURL+"/api/v1/healthcheck", nil)

			It("should be HTTP Status OK", func() {
				Expect(requestResult.Code).To(Equal(http.StatusOK))
			})

			It("should return JSON containing 'status': 'ok'", func() {
				bodyStr := requestResult.Body.String()
				Expect(bodyStr).To(MatchRegexp(".*status[^a-zA-Z]+ok*"))
			})
		})
	})
})
