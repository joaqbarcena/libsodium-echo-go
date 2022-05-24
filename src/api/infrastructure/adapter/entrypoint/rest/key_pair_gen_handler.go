package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaqbarcena/libsodium-echo-go/src/api/core/usecase/keypairgen"
)

const (
	seedParamKey = "seed"
)

type keyPairGenHandler struct {
	useCase keypairgen.UseCase
}

func NewKeyPairGenHandler(keyPairGenUseCase keypairgen.UseCase) *keyPairGenHandler {
	return &keyPairGenHandler{
		useCase: keyPairGenUseCase,
	}
}

func (keyPairGenHandler *keyPairGenHandler) GetHandlerConfig() *Handler {
	return &Handler{
		Path:        "/keypairgen",
		HandlerFunc: keyPairGenHandler.Handle,
		Method:      http.MethodGet,
	}
}

func (keyPairGenHandler *keyPairGenHandler) Handle(c *gin.Context) HandleError {
	seed := c.Query(seedParamKey)

	keyPair, err := keyPairGenHandler.useCase.Execute(seed)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, keyPair)

	return nil
}
