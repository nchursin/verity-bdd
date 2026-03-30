package api

import (
	internalapi "github.com/nchursin/verity-bdd/internal/abilities/api"
)

type CallAnAPI = internalapi.CallAnAPI
type RequestBuilder = internalapi.RequestBuilder
type RequestActivity = internalapi.RequestActivity

type LastResponseStatus = internalapi.LastResponseStatus
type LastResponseBody = internalapi.LastResponseBody
type ResponseHeader = internalapi.ResponseHeader
type JSONPath = internalapi.JSONPath
type ResponseTime = internalapi.ResponseTime

var Using = internalapi.Using
var CallAnApiAt = internalapi.CallAnApiAt

var NewRequestBuilder = internalapi.NewRequestBuilder

var SendRequest = internalapi.SendRequest
var SendGetRequest = internalapi.SendGetRequest
var SendPostRequest = internalapi.SendPostRequest
var SendPutRequest = internalapi.SendPutRequest
var SendDeleteRequest = internalapi.SendDeleteRequest

var NewResponseHeader = internalapi.NewResponseHeader
var NewJSONPath = internalapi.NewJSONPath

func NewResponseBodyAsJSON[T any]() internalapi.ResponseBodyAsJSON[T] {
	return internalapi.NewResponseBodyAsJSON[T]()
}

var LastResponseStatusQ = internalapi.LastResponseStatusQ
var LastResponseBodyQ = internalapi.LastResponseBodyQ
var ResponseTimeQ = internalapi.ResponseTimeQ
