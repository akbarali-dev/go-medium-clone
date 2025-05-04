package server

import (
	v1 "mediumclone/server/v1"
	"mediumclone/storage"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Strg storage.StorageI
}

func NewServer(opts *Options) *gin.Engine {
	router := gin.New()
	handler := v1.New(&v1.HandlerV1{
		Strg: opts.Strg,
	})
	router.POST("/v1/users", handler.CreateUser)
	router.DELETE("/v1/user/:id", handler.DeleteUser)
	router.GET("/v1/user/:id", handler.GetUser)
	router.PUT("/v1/user/:id", handler.UpdateUser)
	return router
}
