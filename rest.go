package rest

import "net/http"

// APIReturn builds the message to be returned by the API
func APIReturn(returnCode int, returnInfo string, w http.ResponseWriter) {
	output := make(map[string]map[string]interface{})
	output["result"] = make(map[string]interface{})
	output["result"]["code"] = returnCode
	output["result"]["info"] = returnInfo

	WriteJSON(output, returnCode, w)
}
