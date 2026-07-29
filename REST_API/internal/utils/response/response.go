package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Error  string
	Status string
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func CustomResponse(err error) Response {

	var Response = Response{
		Error:  err.Error(),
		Status: "Error",
	}
	if err.Error() == "EOF" {
		Response.Error = "Body Not Provided"
	}
	return Response
}

func ValidatorResponse(err validator.ValidationErrors) Response {
	Response := Response{
		Status: "Error",
	}

	var messages []string

	for _, e := range err {
		messages = append(messages, fmt.Sprintf(
			"Field: %s, Tag: %s, Value: %v",
			e.Field(),
			e.Tag(),
			e.Value(),
		))
	}

	Response.Error = strings.Join(messages, ", ")

	return Response
}
