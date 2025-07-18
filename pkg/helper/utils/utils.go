package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/golang-jwt/jwt/v5"
)

func SetValueContext(r *http.Request) (result ContextMap, contextMap jwt.MapClaims, err error) {

	ctxValue := r.Context().Value(constant.ClaimsKey)
	if ctxValue == nil {
		err = fmt.Errorf("claims_value not found in context")
		return
	}

	contextMap, ok := ctxValue.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("claims_value is not of type map[string]any")
		return
	}

	result.Id = fmt.Sprintf("%v", contextMap["id"])
	result.Fullname = fmt.Sprintf("%v", contextMap["fullname"])
	result.Username = fmt.Sprintf("%v", contextMap["username"])
	result.Email = fmt.Sprintf("%v", contextMap["email"])

	return
}

func GetValueOfContext(key string, ctx context.Context) any {
	if ctx.Value(key) != nil {
		return ctx.Value(key)
	}
	return ""
}

func TimeNow() time.Time {
	location, _ := time.LoadLocation(constant.TimeLocation)
	return time.Now().In(location)
}
