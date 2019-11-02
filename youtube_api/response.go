package youtube_api

import (
  "encoding/json"
  "net/http"
  api_json "youtube_api/json"
)

type ResponseInterface interface {
  GetBody() api_json.ResultInterface
}

type Response struct {
  Response   *http.Response
  JsonStruct api_json.JsonInterface
}

func (response *Response) GetBody() api_json.ResultInterface {
  results    := (*response).Response
  jsonStruct := (*response).JsonStruct
  json.NewDecoder(results.Body).Decode(&jsonStruct)
  return jsonStruct.GetResult()
}

