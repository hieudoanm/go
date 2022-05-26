package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const NEWS_V2 = "https://newsapi.org/v2";


type SourceType struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
  Url string `json:"url"`
  Category string `json:"category"`
  Language string `json:"language"`
  Country string `json:"country"`
}

type SourcesResponseBody struct {
  Status string `json:"status"`
  Message string `json:"message"`
  Sources []SourceType `json:"sources"`
}

func GetSources(
  apiKey string,
  category string,
  language string,
  country string,
) ([]SourceType, error) {
  // Build query parameters
  var queryParameters []string = []string{}
  if apiKey != "" {
    queryParameters = append(queryParameters, "apiKey=" + apiKey)
  }
  if category != "" {
    queryParameters = append(queryParameters, "category=" + category)
  }
  if language != "" {
    queryParameters = append(queryParameters, "language=" + language)
  }
  if country != "" {
    queryParameters = append(queryParameters, "country=" + country)
  }

  // Build url
  var url = fmt.Sprintf(
    "%s/top-headlines/sources?%s",
    NEWS_V2,
    strings.Join(queryParameters, "&"),
  )

  body, getError := Get(url)
  if getError != nil {
    return nil, getError
  }

  var sourcesResponseBody SourcesResponseBody
  jsonUnmarshalError := json.Unmarshal(body, &sourcesResponseBody)
  if jsonUnmarshalError != nil {
    return nil, jsonUnmarshalError
  }

  if sourcesResponseBody.Status != "ok" {
    return []SourceType{}, errors.New(sourcesResponseBody.Message)
  }

  return sourcesResponseBody.Sources, nil
}

type Article struct {
  Source SourceType `json:"source"`
  Author string `json:"author"`
  Title string `json:"title"`
  Description string `json:"description"`
  Url string `json:"url"`
  UrlToImage string `json:"urlToImage"`
  PublishedAt string `json:"publishedAt"`
  Content string `json:"content"`
}

type TopHeadlinesResponseBody struct {
  Status string `json:"status"`
  Message string `json:"message"`
  Articles []Article `json:"articles"`
}

func GetTopHeadlines(
  apiKey string,
  sources []string,
  q string,
  category string,
  country string,
  pageSize int,
  page int,
) ([]Article, error) {
  // Assign default value
  if len(sources) == 0 {
    sources = []string{"cnn"}
  }

  // Build query parameters
  var queryParameters []string = []string{}
  if apiKey != "" {
    queryParameters = append(queryParameters, "apiKey=" + apiKey)
  }
  if len(sources) == 0 {
    queryParameters = append(queryParameters, "sources=" + strings.Join(sources, ","))
  }
  if q != "" {
    queryParameters = append(queryParameters, "q=" + q)
  }
  if category != "" {
    queryParameters = append(queryParameters, "category=" + category)
  }
  if country != "" {
    queryParameters = append(queryParameters, "country=" + country)
  }
  if pageSize > 0 {
    queryParameters = append(queryParameters, "pageSize=" + strconv.Itoa(pageSize))
  }
  if page > 0 {
    queryParameters = append(queryParameters, "page=" + strconv.Itoa(page))
  }

  // Build url
  var url = fmt.Sprintf(
    "%s/top-headlines?%s",
    NEWS_V2,
    strings.Join(queryParameters, "&"),
  )

  body, getError := Get(url)
  if getError != nil {
    return nil, getError
  }

  var topHeadlinesResponseBody TopHeadlinesResponseBody
  jsonUnmarshalError := json.Unmarshal(body, &topHeadlinesResponseBody)
  if jsonUnmarshalError != nil {
    return nil, jsonUnmarshalError
  }

  if topHeadlinesResponseBody.Status != "ok" {
    return []Article{}, errors.New(topHeadlinesResponseBody.Message)
  }

  return topHeadlinesResponseBody.Articles, nil
}

func Get(url string) ([]byte, error){
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
