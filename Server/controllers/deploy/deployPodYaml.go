package deploy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"gok8s/handlers"

	"github.com/gin-gonic/gin"
)

// DeployYamlController process YAML deploy
func DeployYamlController(c *gin.Context) {
	fmt.Println("DeployYamlController start........................")

	// 1. Read request data
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to read request body",
			"details": err.Error(),
		})
		return
	}

	// 2. if yaml is empty then return error
	if len(bodyBytes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Empty YAML content",
			"details": "The YAML content cannot be empty",
		})
		return
	}

	// 3. porcess json data.
	var yamlContent []byte
	contentType := c.ContentType()

	if strings.Contains(contentType, "application/json") {
		// Decode json and get yaml content.
		jsonData := make(map[string]interface{})
		err := json.Unmarshal(bodyBytes, &jsonData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed to parse JSON data request",
				"details": err.Error(),
			})
			return
		}

		// get yaml content from json
		if item, ok := jsonData["item"].(string); ok {
			yamlContent = []byte(item)
			fmt.Println("Extracted YAML from JSON item:", item)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "JSON request missing 'item' field",
				"details": "Expected JSON object with 'item' string field containing YAML",
			})
			return
		}
	} else {
		// if no json, then assume it's raw yaml
		yamlContent = bodyBytes
	}

	// 4. check content is empty or not.
	if len(yamlContent) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Empty YAML content after processing",
			"details": "No valid YAML content found in request",
		})
		return
	}

	// 5. build a new request for deploypod handler
	req, err := http.NewRequest(
		http.MethodPost,
		"/api/deploypod",
		io.NopCloser(bytes.NewBuffer(yamlContent)),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create request",
			"details": err.Error(),
		})
		return
	}

	// 6. setup request head.
	req.Header.Set("Content-Type", "application/yaml")
	// copy all headers from original request to new request
	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	// 5. creat a new response writer
	w := httptest.NewRecorder()

	// 6. Put to DeployPodHandler process request
	handlers.DeployPodHandler(w, req)

	// 7. Process request.
	result := w.Result()
	defer result.Body.Close()

	// 8. Copy response head.
	for k, v := range result.Header {
		c.Writer.Header()[k] = v
	}

	// 9. write response code
	c.Writer.WriteHeader(result.StatusCode)

	// 10. write response body.
	responseBody, err := io.ReadAll(result.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to read response body",
			"details": err.Error(),
		})
		return
	}

	if _, err := c.Writer.Write(responseBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to write response",
			"details": err.Error(),
		})
	}
}
