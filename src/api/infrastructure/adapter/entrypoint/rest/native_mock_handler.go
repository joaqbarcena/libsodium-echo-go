package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/model"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/resilience"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/usecase/response"
)

const ()

type nativeMockImplementation struct {
	fileResponseUseCase          response.FileResponseUseCase
	fileResponseEncryptedUseCase response.FileResponseEncryptedUseCase
}

func NewNativeMockHandler(fileResponseUseCase response.FileResponseUseCase, fileResponseEncryptedUseCase response.FileResponseEncryptedUseCase) *nativeMockImplementation {
	return &nativeMockImplementation{
		fileResponseUseCase:          fileResponseUseCase,
		fileResponseEncryptedUseCase: fileResponseEncryptedUseCase,
	}
}

func (impl *nativeMockImplementation) GetHandlerConfig() *Handler {
	return &Handler{
		Path:        "/native/mock",
		HandlerFunc: impl.Handle,
		Method:      http.MethodGet,
	}
}

func (impl *nativeMockImplementation) Handle(c *gin.Context) HandleError {
	encodedPublickKey := c.Query(publicKeyParamKey)
	var content string
	var err *resilience.ApiError = nil

	if encodedPublickKey != "" {
		fileResponseEncryptedInfo := model.FileResponseEncryptedInfo{
			EncodedPublicKey: encodedPublickKey,
		}
		content, err = impl.fileResponseEncryptedUseCase.Execute(fileResponseEncryptedInfo)
	} else {
		content, err = impl.fileResponseUseCase.Execute()
	}

	if err != nil {
		return err
	}

	c.Data(http.StatusOK, "application/json", []byte(content))

	return nil
}
