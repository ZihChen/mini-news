package logwriter

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type logWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *logWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func RequestLog(c *gin.Context) {
	now := time.Now()
	logWriter := &logWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}

	c.Writer = logWriter
	requestBody := getRequestBody(c)

	c.Next()

	latency := time.Since(now)
	responseBody := logWriter.body.String()

	logContext := make(map[string]interface{})
	logContext["request_uri"] = c.Request.RequestURI
	logContext["request_method"] = c.Request.Method
	logContext["refer_service_name"] = c.Request.Referer()
	logContext["refer_request_host"] = c.ClientIP()
	logContext["request_body"] = requestBody
	logContext["request_time"] = now.String()
	logContext["response_body"] = responseBody
	logContext["time_used"] = fmt.Sprintf("%v", latency)
	logContext["header"] = c.Request.Header

	log.Println(logContext)
}

func getRequestBody(c *gin.Context) interface{} {
	switch c.Request.Method {
	case http.MethodGet:
		return c.Request.URL.Query()

	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		var bodyBytes []byte
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return nil
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		return string(bodyBytes)
	}

	return nil
}