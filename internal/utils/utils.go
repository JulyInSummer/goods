package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func SlErr(err error) slog.Attr {
	return slog.Attr{
		Key:   "Error",
		Value: slog.StringValue(err.Error()),
	}
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func ParseValidationErrors(err error) map[string]string {
	parsedErrs := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			parsedErrs[strings.ToLower(e.Field())] = "required field"
		case "gt":
			parsedErrs[strings.ToLower(e.Field())] = "value should be greater than 0"
		default:
			parsedErrs[strings.ToLower(e.Field())] = "validation error on this field"
		}
	}

	return parsedErrs
}

func MakeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, err)
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func QueryParamToInt(request *http.Request, key string, defValue int) (error, *int) {
	value := request.URL.Query().Get(key)

	if value == "" {
		return nil, &defValue
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return err, nil
	}

	return nil, &valueInt
}
