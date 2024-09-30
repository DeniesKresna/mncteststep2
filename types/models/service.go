package models

type Response struct {
	ResponseCode  string      `json:"responseCode"`
	ResponseDesc  string      `json:"responseDesc"`
	ResponseData  interface{} `json:"responseData"`
	ResponseTrace string      `json:"responseTrace"`
}
