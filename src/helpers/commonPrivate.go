package helpers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	logger "github.com/IvanSkripnikov/go-logger"
)

func getIDFromRequestString(url string) (int, error) {
	vars := strings.Split(url, "/")

	return strconv.Atoi(vars[len(vars)-1])
}

func checkError(w http.ResponseWriter, err error, category string) bool {
	httpStatusCode := http.StatusOK
	if err != nil {
		logger.Errorf("Runtime error %s", err.Error())

		var data ResponseData
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			httpStatusCode = http.StatusNotFound
			data = ResponseData{
				"response": "Data not found",
			}
		} else {
			httpStatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			data = ResponseData{
				"response": "Internal error",
			}
		}

		SendResponse(w, data, category, httpStatusCode)

		return true
	}

	return false
}
