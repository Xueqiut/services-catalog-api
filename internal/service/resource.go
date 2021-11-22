package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Xueqiut/services-catalog-api/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *gin.RouterGroup, repository repository, logger *log.Logger) {
	res := resource{repository, logger}

	r.GET("/services", res.list)
	r.GET("/services/:id", res.get)
	r.POST("/services", res.post)
}

type resource struct {
	repository	repository
	logger		*log.Logger
}

func (r resource) list(c *gin.Context) {
	search := c.Query("search")
	sort := c.Query("sort")
	page := c.Query("page")
	perPage := c.Query("per_page")

	count, err := r.repository.count()
	if err != nil {
		r.logger.Fatal(err)
		c.Error(err)
		return
	}

	pages := pagination.NewPage(page, perPage, count)

	services, err := r.repository.list(search, sort, pages.Offset(), pages.Limit())
	if err != nil {
		r.logger.Fatal(err)
		c.Error(err)
		return
	}

	pages.Items = services
	c.IndentedJSON(http.StatusOK, pages)
}

func (r resource) get(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    r.logger.Fatal(err)
    c.Error(err)
		return
  }

	service, err := r.repository.get(i)
	if err != nil {
		r.logger.Fatal(err)
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, service)
}

func (r resource) post(c *gin.Context) {
	var newService Service

  // Call BindJSON to bind the received JSON to newService, which the deserializes the JSON
  if err := c.BindJSON(&newService); err != nil {
	r.logger.Fatal(err)
    c.Error(err)
    return
  }

	id, err := r.repository.create(newService)
	if err != nil {
		r.logger.Fatal(err)
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, id)
}