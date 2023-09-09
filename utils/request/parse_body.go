package request

import (
	"encoding/json"
	"net/http"
)

func ParseBody[T interface{}](r *http.Request) (T, error) {
	var body T

	err := json.NewDecoder(r.Body).Decode(&body)

	return body, err
}
