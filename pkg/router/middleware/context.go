package middleware

import (
	"git.containerum.net/ch/mail-templater/pkg/clients"
	"git.containerum.net/ch/mail-templater/pkg/storages"
	"git.containerum.net/ch/mail-templater/pkg/upstreams"
	"github.com/gin-gonic/gin"
)

const (
	//MTServices is key for services
	MTServices = "mt-service"
)

// Services is a collection of dependencies to perform server operations
type Services struct {
	MessagesStorage   storages.MessagesStorage
	TemplateStorage   storages.TemplateStorage
	Upstream          upstreams.Upstream
	UpstreamSimple    upstreams.Upstream
	UserManagerClient clients.UserManagerClient
}

// RegisterServices adds services to context
func RegisterServices(svc *Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(MTServices, svc)
	}
}
