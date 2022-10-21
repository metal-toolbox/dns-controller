// Package router has a router for dnscontroller
package router

import (
	"path"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.hollow.sh/toolbox/ginjwt"
	"go.uber.org/zap"
)

const (
	// V1URI is the path prefix for all v1 endpoints
	V1URI = "/api/v1"

	// RecordsURI is the path to the regular record endpoint, called by the
	// client to .
	RecordsURI = "/records"

	// RecordURI is the path to the endpoint used for
	// retrieving the stored record for an instance
	RecordURI = "/records/:record/:recordtype"

	// RecordAnswerURI is for interactions with record's answers
	RecordAnswerURI = "/records/:record/:recordtype/answers"

	// scopePrefix = "dnscontroller"
)

// Router provides a router for the v1 API
type Router struct {
	authMW *ginjwt.Middleware
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

// New builds a Router
func New(amw *ginjwt.Middleware, db *sqlx.DB, l *zap.SugaredLogger) *Router {
	return &Router{authMW: amw, db: db, logger: l}
}

// Routes will add the routes for this API version to a router group
func (r *Router) Routes(rg *gin.RouterGroup) {
	// TODO: add auth'd endpoints
	// authMw := r.AuthMW
	// rg.POST(RecordURI, authMw.AuthRequired(), authMw.RequiredScopes(upsertScopes("record")))
	rg.GET(RecordURI, r.getRecord)
	rg.POST(RecordURI, r.createRecord)
	rg.DELETE(RecordURI, r.deleteRecord)
}

// GetRecordPath returns the path used by an instance to fetch Record
func GetRecordPath() string {
	return path.Join(V1URI, RecordsURI)
}

// func upsertScopes(items ...string) []string {
// 	s := []string{"write", "create", "update"}
// 	for _, i := range items {
// 		s = append(s, fmt.Sprintf("%s:create:%s", scopePrefix, i))
// 	}

// 	for _, i := range items {
// 		s = append(s, fmt.Sprintf("%s:update:%s", scopePrefix, i))
// 	}

// 	return s
// }

// func readScopes(items ...string) []string {
// 	s := []string{"read"}
// 	for _, i := range items {
// 		s = append(s, fmt.Sprintf("%s:read:%s", scopePrefix, i))
// 	}

// 	return s
// }

// func deleteScopes(items ...string) []string {
// 	s := []string{"write", "delete"}
// 	for _, i := range items {
// 		s = append(s, fmt.Sprintf("%s:delete:%s", scopePrefix, i))
// 	}

// 	return s
// }
