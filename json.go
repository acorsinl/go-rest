package rest

import (
  "net/http"
)

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
