package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message`
}

type Map map[string]interface{}

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) error {
	if data == nil {
		w.Header().Set("Content-Type", "application:json; charset=utf-8")
		w.WriteHeader(statusCode)
		return nil
	}
	// try to fit the interface in a Json format
	j, err := json.Marshal(data)
	if err != nil {
		w.Header().Set("Content-Type", "application:json; charset=utf-8")
		w.WriteHeader(statusCode)
		return err
	}

	w.Write(j)
	return nil
}

func HTTPError(w http.ResponseWriter, r *http.Request, statusCode int, message string) error {
	msg := ErrorMessage{
		Message: message,
	}
	fmt.Println("Message:", message)
	return JSON(w, r, statusCode, msg)
}
