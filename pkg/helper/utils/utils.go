package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/golang-jwt/jwt/v5"
)

// GetUserClaimsFromContext extracts the JWT claims from the HTTP request context,
// parses them into a ContextMap struct, and returns both the structured result
// and the raw jwt.MapClaims.
//
// Returns an error if the claims are missing or have an unexpected type
func GetUserClaimsFromContext(r *http.Request) (result ContextMap, contextMap jwt.MapClaims, err error) {
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

// MapClaimsToContextMap converts a map of JWT claims (typically from jwt.MapClaims)
// into a ContextMap struct, extracting fields like Id, Fullname, Username, and Email.
func MapClaimsToContextMap(contextMap jwt.MapClaims) (result ContextMap) {
	result.Id = fmt.Sprintf("%v", contextMap["id"])
	result.Fullname = fmt.Sprintf("%v", contextMap["fullname"])
	result.Username = fmt.Sprintf("%v", contextMap["username"])
	result.Email = fmt.Sprintf("%v", contextMap["email"])
	return
}

func TimeNow() time.Time {
	location, _ := time.LoadLocation(constant.TimeLocation)
	return time.Now().In(location)
}

func StringToInt64(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return value
}
