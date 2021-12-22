package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/model"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/resilience"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/usecase/encryption"
)

const (
	publicKeyParamKey = "public_key"
	textToEncryptKey  = "text"
)

type implementation struct {
	encryptionUseCase encryption.UseCase
}

func NewEncryptionHandler(encryptionUseCase encryption.UseCase) *implementation {
	return &implementation{
		encryptionUseCase: encryptionUseCase,
	}
}

func (impl *implementation) GetHandlerConfig() *Handler {
	return &Handler{
		Path:        "/encrypt",
		HandlerFunc: impl.Handle,
		Method:      http.MethodGet,
	}
}

func (impl *implementation) Handle(c *gin.Context) HandleError {
	encodedPK := c.Query(publicKeyParamKey)
	textToEncrypt := c.Query(textToEncryptKey)

	if encodedPK == "" {
		return resilience.NewApiError(http.StatusBadRequest, "pk is mandatory and shouldn't be empty")
	}

	if textToEncrypt == "" {
		return resilience.NewApiError(http.StatusBadRequest, "text is mandatory and shouldn't be empty")
	}

	encryptInfo := model.EncryptInfo{
		EncodedPublicKey: encodedPK,
		TextToEncrypt:    textToEncrypt,
	}

	encryptedText, err := impl.encryptionUseCase.Execute(encryptInfo)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"encrypted": encryptedText,
	})

	return nil
}
