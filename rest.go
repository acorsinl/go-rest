package rest

import "net/http"

// CheckError checks if there is an error and handles an appropiate
// HTTP output
func CheckError(err error, errorCode int, w http.ResponseWriter) {
	if err == nil {
		return
	}

	output := make(map[string]map[string]interface{})
	output["result"] = make(map[string]interface{})
	output["result"]["code"] = errorCode
	output["result"]["info"] = err.Error()

	WriteJSON(output, http.StatusInternalServerError, w)
	panic(err.Error())
}

// APIReturn builds the message to be returned by the API
func APIReturn(returnCode int, returnInfo string, w http.ResponseWriter) {
	output := make(map[string]map[string]interface{})
	output["result"] = make(map[string]interface{})
	output["result"]["code"] = returnCode
	output["result"]["info"] = returnInfo

	WriteJSON(output, returnCode, w)
}

// APISingleResult builds the message to be returned
// when there is just one single result to be shown
func APISingleResult(returnCode int, returnInfo string, data map[string]interface{}, w http.ResponseWriter) {
	output := make(map[string]map[string]interface{})
	output["result"] = make(map[string]interface{})
	output["result"]["code"] = returnCode
	output["result"]["info"] = returnInfo
	output["data"] = data

	WriteJSON(output, returnCode, w)
}
