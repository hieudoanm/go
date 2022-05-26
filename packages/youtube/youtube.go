package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const YOUTUBE_V3 = "https://www.googleapis.com/youtube/v3";

type VideoCategorySnippet struct {
  Title string `json:"title"`
  Assignable bool `json:"assignable"`
  ChannelId string `json:"channelId"`
}

type VideoCategoryItem struct {
  Id string `json:"id"`
  Snippet VideoCategorySnippet `json:"snippet"`
}

type VideoCategoryResponseBody struct {
  Items []VideoCategoryItem `json:"items"`
}

type VideoCategory struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Assignable bool `json:"assignable"`
  ChannelId string `json:"channelId"`
}

func GetVideoCategories(key string, regionCode string, part string) ([]VideoCategory, error) {
  // Assign default value
  if regionCode == "" {
    regionCode = "US"
  }

  // Build query parameters
  var queryParameters []string = []string{}
  if key != "" {
    queryParameters = append(queryParameters, "key=" + key)
  }
  if regionCode != "" {
    queryParameters = append(queryParameters, "regionCode=" + regionCode)
  }
  if part != "" {
    queryParameters = append(queryParameters, "part=" + part)
  }

  var url = fmt.Sprintf(
    "%s/videoCategories?%s",
    YOUTUBE_V3,
    strings.Join(queryParameters, "&"),
  )

  body, getError := Get(url)
  if getError != nil {
    return nil, getError
  }

  var responseBody VideoCategoryResponseBody
  jsonUnmarshalError := json.Unmarshal(body, &responseBody)
  if jsonUnmarshalError != nil {
    return nil, jsonUnmarshalError
  }

  var items = responseBody.Items
  var videoCategories []VideoCategory
  for _, item := range items {
    var videoCategory VideoCategory
    videoCategory.Id = item.Id
    videoCategory.Title = item.Snippet.Title
    videoCategory.Assignable = item.Snippet.Assignable
    videoCategory.ChannelId = item.Snippet.ChannelId

    videoCategories = append(videoCategories, videoCategory)
  }

  return videoCategories, nil
}

type VideoSnippet struct {
  Title string `json:"title"`
  Description string `json:"description"`
  ChannelId string `json:"channelId"`
  ChannelTitle string `json:"channelTitle"`
  CategoryId string `json:"categoryId"`
  PublishedAt string `json:"publishedAt"`
}

type VideoItem struct {
  Id string `json:"id"`
  Snippet VideoSnippet `json:"snippet"`
}

type VideoResponseBody struct {
  Items []VideoItem `json:"items"`
}

type Video struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  ChannelId string `json:"channelId"`
  ChannelTitle string `json:"channelTitle"`
  CategoryId string `json:"categoryId"`
  PublishedAt string `json:"publishedAt"`
}

func GetVideos(key string, chart string, part string) ([]Video, error) {
  // Assign default value
  if chart == "" {
    chart = "mostPopular"
  }

  var queryParameters []string = []string{}
  if key != "" {
    queryParameters = append(queryParameters, "key=" + key)
  }
  if chart != "" {
    queryParameters = append(queryParameters, "chart=" + chart)
  }
  if part != "" {
    queryParameters = append(queryParameters, "part=" + part)
  }

  var url = fmt.Sprintf(
    "%s/videoCategories?%s",
    YOUTUBE_V3,
    strings.Join(queryParameters, "&"),
  )

  body, getError := Get(url)
  if getError != nil {
    return nil, getError
  }

  var responseBody VideoResponseBody
  jsonUnmarshalError := json.Unmarshal(body, &responseBody)
  if jsonUnmarshalError != nil {
    return nil, jsonUnmarshalError
  }

  var items = responseBody.Items
  var videos []Video
  for _, item := range items {
    var video Video
    video.Id = item.Id
    video.Title = item.Snippet.Title
    video.Description = item.Snippet.Description
    video.ChannelId = item.Snippet.ChannelId
    video.ChannelTitle = item.Snippet.ChannelTitle
    video.CategoryId = item.Snippet.CategoryId
    video.PublishedAt = item.Snippet.PublishedAt

    videos = append(videos, video)
  }

  return videos, nil
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
