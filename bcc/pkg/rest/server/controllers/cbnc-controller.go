package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/abcd/bcc/pkg/rest/server/models"
	"github.com/unnagirirao/abcd/bcc/pkg/rest/server/services"
	"net/http"
	"strconv"
)

type CbncController struct {
	cbncService *services.CbncService
}

func NewCbncController() (*CbncController, error) {
	cbncService, err := services.NewCbncService()
	if err != nil {
		return nil, err
	}
	return &CbncController{
		cbncService: cbncService,
	}, nil
}

func (cbncController *CbncController) CreateCbnc(context *gin.Context) {
	// validate input
	var input models.Cbnc
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger cbnc creation
	if _, err := cbncController.cbncService.CreateCbnc(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Cbnc created successfully"})
}

func (cbncController *CbncController) UpdateCbnc(context *gin.Context) {
	// validate input
	var input models.Cbnc
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger cbnc update
	if _, err := cbncController.cbncService.UpdateCbnc(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cbnc updated successfully"})
}

func (cbncController *CbncController) FetchCbnc(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger cbnc fetching
	cbnc, err := cbncController.cbncService.GetCbnc(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, cbnc)
}

func (cbncController *CbncController) DeleteCbnc(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger cbnc deletion
	if err := cbncController.cbncService.DeleteCbnc(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Cbnc deleted successfully",
	})
}

func (cbncController *CbncController) ListCbncs(context *gin.Context) {
	// trigger all cbncs fetching
	cbncs, err := cbncController.cbncService.ListCbncs()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, cbncs)
}

func (*CbncController) PatchCbnc(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*CbncController) OptionsCbnc(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*CbncController) HeadCbnc(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
