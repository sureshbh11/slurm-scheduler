// (c) Copyright 2020 Hewlett Packard Enterprise Development LP

package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/twinj/uuid"

	hpeErrors "github.com/hpe-hcss/errors/pkg/errors"
)

// GetEnv is a helper function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// GetRequestIDFromGin gets requestID from context
func GetRequestIDFromGin(c *gin.Context) interface{} {
	requestID, ok := c.Get("RequestId")
	if !ok {
		// this is not expected
		log.Errorf("Couldn't find the request id in the context")
		requestID = "no request id"
	}
	return requestID
}

// GetProjectUUID is a function which generates the uuid based on the string given
func GetProjectUUID(s string) string {
	return uuid.NewV5(uuid.NameSpaceURL, s).String()
}

// FetchActualErrorMsgFromResp extracts a the body from the error response without
// draining the response body
func FetchActualErrorMsgFromResp(resp *http.Response) string {
	if resp.Body == nil {
		log.Debug("Response body is empty")
		return ""
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to decode response body: %v", err)
		return ""
	}
	// to make sure the resp.body is not drained out
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	respmsgstr := string(bodyBytes)
	log.Debugf("resp err msg str: %v", respmsgstr)
	return respmsgstr
}

// ParseResponse a common response parsing method
func ParseResponse(resp *http.Response, respBody interface{}) (interface{}, error) {
	log.Debugf("Status code from the server :%v", resp.StatusCode)
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		err := json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to decode response msg: %v", err)
			log.Error(errorMsg)
			return nil, hpeErrors.UnexpectedResponseErrIfNotOk([]int{http.StatusOK, http.StatusAccepted, http.StatusCreated}, resp, fmt.Errorf("an unexpected error occurred"))
		}
		return respBody, nil
	case http.StatusNoContent:
		return "No content", nil
	case http.StatusNotFound:
		respmsg := FetchActualErrorMsgFromResp(resp)
		if respmsg == "" {
			respmsg = "Resource not found"
		}
		return nil, hpeErrors.MakeErrNotFound(hpeErrors.ErrorResponse{Message: respmsg})
	case http.StatusUnauthorized:
		msg := FetchActualErrorMsgFromResp(resp)
		if msg == "" {
			msg = "Unauthorized access"
		}
		return nil, hpeErrors.MakeErrUnauthorized(msg)
	case http.StatusBadRequest:
		respmsg := FetchActualErrorMsgFromResp(resp)
		if respmsg == "" {
			respmsg = "Bad request"
		}
		return nil, hpeErrors.MakeErrBadRequest(hpeErrors.ErrorResponse{Message: respmsg})
	default:
		respmsg := FetchActualErrorMsgFromResp(resp)
		return nil, hpeErrors.MakeErrInternalError(hpeErrors.ErrorResponse{Message: respmsg})
	}
}
