package v1

import (
	"database/sql"
	"errors"
	"log"
	"mediumclone/server/models"
	"mediumclone/storage/repo"
	// "os/user"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerV1) CreateUser(ctx *gin.Context) {
	var req models.CreateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate user ID"})
		return
	}

	user, err := h.strg.User().Create(ctx, &repo.User{
		ID:        id.String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, models.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	})

	// user = models.User{
	// 	ID:        createdUser.ID,
	// 	FirstName: createdUser.FirstName,
	// 	LastName:  createdUser.LastName,
	// 	Email:     createdUser.Email,
	// 	Password:  createdUser.Password,
	// }

}
func (h *handlerV1) UpdateUser(ctx *gin.Context) {}
func (h *handlerV1) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.strg.User().Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		}
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	ctx.JSON(http.StatusOK, models.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	})

}

func (h *handlerV1) DeleteUser(ctx *gin.Context) {}
