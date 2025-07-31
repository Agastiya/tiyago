package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/go-chi/chi/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetUrl(r *http.Request, key string) (int64, error) {

	var err error
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return 0, fmt.Errorf("parameter not valid")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parameter not valid")
	}

	return id, nil
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

func ToFuncName(name string) string {
	parts := strings.Split(name, "_")
	for i, p := range parts {
		parts[i] = cases.Title(language.English).String(p)
	}
	return strings.Join(parts, "")
}

func RenderTemplate(name string, data map[string]string) (string, error) {
	tplContent, err := os.ReadFile(fmt.Sprintf("cmd/stubs/%s.stub", name))
	if err != nil {
		return "", err
	}

	tpl, err := template.New("stub").Parse(string(tplContent))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
