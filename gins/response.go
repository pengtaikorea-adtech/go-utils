package gins

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ctxRespType     = ".gins.context.response.type"
	ctxRespBody     = ".gins.context.response.body"
	ctxRespParamKey = ".gins.context.response.params"
	ctxStatusKey    = ".gins.context.response.status"

	respTypeJSON     = "json"
	respTypeText     = "text"
	respTypeView     = "view"
	respTypeRedirect = "redirect"
	respTypeError    = "error"
)

// ErrorResponseAlreadySet - iff 1 response handler
var ErrorResponseAlreadySet = errors.New("gins: only 1 response per a handler")

// ErrorResponseNone - any response set for the handler
var ErrorResponseNone = errors.New("gins: response not set for the handler")

func testResponseHadSet(ctx *gin.Context) bool {
	_, exists := ctx.Get(ctxRespType)
	return exists
}

// ResponseJSON - prepare data to responsing json
func ResponseJSON(ctx *gin.Context, body gin.H) {
	if !testResponseHadSet(ctx) {
		// set type json
		ctx.Set(ctxRespType, respTypeJSON)
		ctx.Set(ctxRespBody, body)
	}
}

// ResponseView - prepare data to render view
func ResponseView(ctx *gin.Context, viewpath string, params gin.H) {
	if !testResponseHadSet(ctx) {
		// set type view
		ctx.Set(ctxRespType, respTypeView)
		ctx.Set(ctxRespParamKey, params)
		ctx.Set(ctxRespBody, viewpath)
	}
}

// ResponseText - prepare data to response static
func ResponseText(ctx *gin.Context, path string) {
	if !testResponseHadSet(ctx) {
		ctx.Set(ctxRespType, respTypeText)
		ctx.Set(ctxRespBody, path)
	}
}

// ResponseRedirect - redirecting response
func ResponseRedirect(ctx *gin.Context, to string, isTemp bool) {
	if !testResponseHadSet(ctx) {
		var status int

		// set type view
		ctx.Set(ctxRespType, respTypeRedirect)
		ctx.Set(ctxStatusKey, status)
		ctx.Set(ctxRespParamKey, to)
	}
}

// ResponseError - error response data
func ResponseError(ctx *gin.Context, status int, err error) {
	// If error occurred, overwrite any other response
	if status < 300 || 600 <= status {
		status = http.StatusInternalServerError
	}
	ctx.Set(ctxRespType, respTypeError)
	ctx.Set(ctxStatusKey, status)
	if err != nil {
		ctx.Set(ctxRespBody, err)
	}
}

// HandlerWrap - complete response handler decorator
func HandlerWrap(handle gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// run core function first
		handle(ctx)

		if testResponseHadSet(ctx) {
			switch ctx.GetString(ctxRespType) {
			// response error
			case respTypeError:
				ctx.AbortWithError(
					ctx.GetInt(ctxStatusKey),
					ctx.MustGet(ctxRespBody).(error))
			// redirect url
			case respTypeRedirect:
				ctx.Redirect(ctx.GetInt(ctxStatusKey), ctx.GetString(ctxRespBody))

			// response json
			case respTypeJSON:
				ctx.JSON(http.StatusOK, ctx.MustGet(ctxRespBody).(gin.H))
			case respTypeView:
				ctx.HTML(http.StatusOK, ctx.GetString(ctxRespBody), ctx.MustGet(ctxRespParamKey).(gin.H))
			// response View
			case respTypeText:
				ctx.String(http.StatusOK, ctx.GetString(ctxRespBody))
			}
		}
	}
}
