package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rx "go.hollow.sh/dnscontroller/pkg/api/v1/record"
)

func (r *Router) deleteRecord(c *gin.Context) {
	record, err := rx.NewRecord(c)
	if err != nil {
		badRequestResponse(c, rx.ErrorInvalidRecord.Error(), err)
	}

	if err := record.Delete(c.Request.Context(), r.db); err != nil {
		badRequestResponse(c, "failed to delete record", err)
		return
	}

	deletedResponse(c)
}

func (r *Router) createRecord(c *gin.Context) {
	record, err := rx.NewRecord(c)
	if err != nil {
		badRequestResponse(c, rx.ErrorInvalidRecord.Error(), err)
	}

	err = record.FindOrCreate(c.Request.Context(), r.db)
	if err != nil {
		badRequestResponse(c, "invalid record", err)
		return
	}

	createdResponse(c)
}

func (r *Router) getRecord(c *gin.Context) {
	record, err := rx.NewRecord(c)
	if err != nil {
		badRequestResponse(c, "could not create record", err)
	}

	err = record.Find(c.Request.Context(), r.db)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, record)
}
