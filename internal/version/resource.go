package version

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *gin.RouterGroup, repository repository, logger *log.Logger) {
	res := resource{repository, logger}

	r.GET("/services/:id/versions", res.list)
}

type resource struct {
	repository	repository
	logger		*log.Logger
}

func (r resource) list(c *gin.Context) {
	serviceId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        r.logger.Fatal(err)
        c.Error(err)
		return
    }

	versions, err := r.repository.list(serviceId)
	if err != nil {
		r.logger.Fatal(err)
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, versions)
}