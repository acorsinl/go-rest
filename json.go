package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// EncodeJSON returns a []byte encoded given a json interface
func EncodeJSON(input interface{}) ([]byte, error) {
	var b []byte
	var err error

	b, err = json.Marshal(input)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// DecodeJSON returns an interface with unmarshalled data given an http request
func DecodeJSON(r *http.Request, destination interface{}) error {
	content, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, destination)
	if err != nil {
		return err
	}

	return nil
}

// WriteJSON returns the API result via HTTP
func WriteJSON(input interface{}, status int, w http.ResponseWriter) {
	output, err := EncodeJSON(input)
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(output)
}
