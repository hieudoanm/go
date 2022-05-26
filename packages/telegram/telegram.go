package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const TELEGRAM_BOT = "https://api.telegram.org/bot"

func SendMessage(
  apiToken string,
  webhook string,
  chatId int,
  text string,
  parseMode string,
) ([]byte, error) {
  var requestBody map[string]string = map[string]string{}
  requestBody["chat_id"] = strconv.Itoa(chatId)
  requestBody["text"] = text
  if parseMode == "" {
    requestBody["parse_mode"] = "HTML"
  } else {
    requestBody["parse_mode"] = parseMode
  }

  // Build url
  var url = fmt.Sprintf(
    "%s%s/sendMessage",
    TELEGRAM_BOT,
    apiToken,
  )

  body, getError := Post(url, requestBody)
  if getError != nil {
    return nil, getError
  }

  return body, nil
}

func SetWebhook(apiToken string, webhook string) ([]byte, error) {
  var requestBody map[string]string = map[string]string{}
  requestBody["webhook"] = webhook

  // Build url
  var url = fmt.Sprintf(
    "%s%s/setWebhook",
    TELEGRAM_BOT,
    apiToken,
  )

  body, getError := Post(url, requestBody)
  if getError != nil {
    return nil, getError
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
