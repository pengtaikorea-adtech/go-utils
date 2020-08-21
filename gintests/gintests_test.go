package gintests

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pengtaikorea-adtech/go-utils/gins"
)

// building engine here
func buildEngine(grp gins.RouteGroup) *gin.Engine {
	engine := gin.New()
	gins.RegisterGroup(engine, grp)
	return engine
}

// responseOK gin.handler
func responseOK(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

// OK, it only contains "/ok" OK
var singleOKRoute = gins.RouteGroup{
	Path: "/",
	Entities: []gins.RouteEntity{
		{http.MethodGet, "/ok", responseOK},
	},
}

func TestGinEngineBuilder(t *testing.T) {
	buildEngine(singleOKRoute)
}

func TestGinTester(t *testing.T) {
	testEngine := buildEngine(singleOKRoute)
	assummings := []Assume{
		AssertCase(http.MethodGet, "/ok", nil, ExpectStatusOK, ExpectStatusExists),
		AssertCase(http.MethodGet, "/no", nil, ExpectStatusNotExists),
	}

	TestRequest(t, testEngine, assummings...)

}
