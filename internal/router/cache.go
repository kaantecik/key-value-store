package router

import (
	"net/http"

	"github.com/kaantecik/key-value-store/internal/config"
	"github.com/kaantecik/key-value-store/internal/entities"
	middleware "github.com/kaantecik/key-value-store/internal/middleware/logging"
	httpTool "github.com/kaantecik/key-value-store/tools/http"
)

type setCacheBody struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type getCacheBody struct {
	Key string `json:"key"`
}

func Set(c *entities.Cache) http.Handler {
	return middleware.HttpLogger(set(c))
}

func Get(c *entities.Cache) http.Handler {
	return middleware.HttpLogger(get(c))
}

func Flush(c *entities.Cache) http.Handler {
	return middleware.HttpLogger(flush(c))
}

func set(c *entities.Cache) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			var body setCacheBody

			httpTool.ParseRequestBody(writer, request, &body)

			if body.Key != "" && body.Value != nil {
				c.Set(body.Key, body.Value)

				httpTool.SendResponse(writer, http.StatusCreated, map[string]interface{}{
					"status":  true,
					"message": config.MessageCreated,
				})
			} else {
				httpTool.SendResponse(writer, http.StatusForbidden, map[string]interface{}{
					"status":  false,
					"message": config.MessageInvalidBody,
				})
			}
		} else {
			httpTool.SendResponse(writer, http.StatusMethodNotAllowed, map[string]interface{}{
				"status":  false,
				"message": config.MessageMethodNotAllowed,
			})
		}
	}
}

func get(c *entities.Cache) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			var body getCacheBody

			httpTool.ParseRequestBody(writer, request, &body)
			key := body.Key

			if key != "" {
				item, found := c.Get(body.Key)
				if found {
					httpTool.SendResponse(writer, http.StatusOK, map[string]interface{}{
						"status": found,
						"item":   item,
					})
				} else {
					httpTool.SendResponse(writer, http.StatusNotFound, map[string]interface{}{
						"status":  found,
						"message": config.MessageMethodNotFound,
					})
				}
			} else {
				httpTool.SendResponse(writer, http.StatusForbidden, map[string]interface{}{
					"status":  false,
					"message": config.MessageInvalidBody,
				})
			}

		} else {
			httpTool.SendResponse(writer, http.StatusMethodNotAllowed, map[string]interface{}{
				"status":  false,
				"message": config.MessageMethodNotAllowed,
			})
		}
	}

}

func flush(c *entities.Cache) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "DELETE" {
			c.Flush()
			httpTool.SendResponse(writer, http.StatusOK, map[string]interface{}{
				"status":  true,
				"message": config.MessageCacheFlushed,
			})
		} else {
			httpTool.SendResponse(writer, http.StatusMethodNotAllowed, map[string]interface{}{
				"status":  false,
				"message": config.MessageMethodNotAllowed,
			})
		}
	}
}
