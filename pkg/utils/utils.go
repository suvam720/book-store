package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func ParseBody(c *gin.Context, x interface{}) {
	if body, err := ioutil.ReadAll(c.Request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
