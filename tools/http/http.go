package http

import (
	"encoding/json"
	"github.com/kaantecik/key-value-store/internal/logging"
	"io/ioutil"
	"net/http"
)

// CheckError returns 403 if err is not nil.
func CheckError(writer http.ResponseWriter, err error) {
	if err != nil {
		logging.ErrorLogger.Error(err)
		msg := map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		SendResponse(writer, http.StatusForbidden, msg)
	}
}

// ParseRequestBody function gets body of request and returns parsed target.
func ParseRequestBody(writer http.ResponseWriter, request *http.Request, target interface{}) {
	data, err := ioutil.ReadAll(request.Body)
	CheckError(writer, err)

	err = json.Unmarshal(data, &target)
	CheckError(writer, err)
}

// SendResponse returns api response.
func SendResponse(writer http.ResponseWriter, statusCode int, msg interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	conv, e := json.Marshal(msg)
	if e != nil {
		logging.ErrorLogger.Error(e)
	}
	writer.WriteHeader(statusCode)
	_, e = writer.Write(conv)
	if e != nil {
		logging.ErrorLogger.Panic(e)
	}
}
