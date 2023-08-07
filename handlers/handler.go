package handlers

import (
	"net/http"
	"testtask/models"
	"testtask/storage"

	"time"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Store *storage.RequestStorage
}

func NewRequestHandler(store *storage.RequestStorage) *RequestHandler {
	return &RequestHandler{
		Store: store,
	}
}

func (r *RequestHandler) Create(ctx *gin.Context) {
	var newRequest models.Request
	if err := ctx.ShouldBindJSON(&newRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.Store.AddRequest(newRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Request recorded successfully",
	})
}

func (r *RequestHandler) GetRequests(ctx *gin.Context) {
	currentTime := time.Now().UTC().Add(-5 * time.Second)
	lastFSecond := r.Store.GetRequestsSince(currentTime)

	ctx.JSON(http.StatusOK, gin.H{
		"count":    len(lastFSecond),
		"requests": lastFSecond,
	})
}
