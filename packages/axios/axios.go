package axios

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
  response, httpGetError := http.Get(url)
  if httpGetError != nil {
    return nil, httpGetError
  }
  defer response.Body.Close()

  body, readBodyError := ioutil.ReadAll(response.Body)
  if readBodyError != nil {
    return nil, readBodyError
  }

  return body, nil
}

func Post(url string, requestBody map[string]string) ([]byte, error) {
  jsonData, jsonMarshalError := json.Marshal(requestBody)
  if jsonMarshalError != nil {
    return nil, jsonMarshalError
  }

  response, httpPostError := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
  if httpPostError != nil {
    return nil, httpPostError
  }
  defer response.Body.Close()

  body, readBodyError := ioutil.ReadAll(response.Body)
  if readBodyError != nil {
    return nil, readBodyError
  }

  return body, nil
}
