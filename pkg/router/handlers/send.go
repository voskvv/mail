package handlers

import (
	"net/http"

	mttypes "git.containerum.net/ch/json-types/mail-templater"
	m "git.containerum.net/ch/mail-templater/pkg/router/middleware"
	"github.com/gin-gonic/gin"

	"git.containerum.net/ch/json-types/errors"
)

func SimpleSendHandler(ctx *gin.Context) {
	var request mttypes.SimpleSendRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(err)
		//	ctx.AbortWithStatusJSON(http.StatusBadRequest, ParseBindErorrs(err))
		return
	}

	svc := ctx.MustGet(m.MTServices).(*m.Services)

	_, tv, err := svc.TemplateStorage.GetLatestVersionTemplate(request.Template)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(errors.ErrorWithHTTPStatus(err))
		return
	}

	info, err := svc.UserManagerClient.UserInfoByID(ctx, request.UserID)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	recipient := &mttypes.Recipient{
		ID:        request.UserID,
		Name:      info.Login,
		Email:     info.Login,
		Variables: request.Variables,
	}
	status, err := svc.UpstreamSimple.SimpleSend(ctx, request.Template, tv, recipient)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, mttypes.SimpleSendResponse{
		UserID: status.RecipientID,
	})
}
