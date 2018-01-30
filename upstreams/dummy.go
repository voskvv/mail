package upstreams

import (
	"context"

	"git.containerum.net/ch/json-types/mail-templater"
	"github.com/sirupsen/logrus"
)

type dummyUpstream struct {
	log *logrus.Entry
}

// NewDummyUpstream returns a new dummy email upstream.
// It actually does nothing, only logs actions
func NewDummyUpstream() Upstream {
	return &dummyUpstream{
		log: logrus.WithField("component", "dummy_upstream"),
	}
}

func (du *dummyUpstream) Send(ctx context.Context, templateName string, tsv *mail.TemplateStorageValue, request *mail.SendRequest) (resp *mail.SendResponse, err error) {
	resp = &mail.SendResponse{}
	for _, recipient := range request.Message.Recipients {
		du.log.WithField("template", templateName).WithFields(recipient.Variables).Infoln("Sending email to", recipient.Email)
		resp.Statuses = append(resp.Statuses, mail.SendStatus{
			RecipientID:  recipient.ID,
			TemplateName: templateName,
			Status:       "OK",
		})
	}
	return
}

func (du *dummyUpstream) SimpleSend(ctx context.Context, templateName string, tsv *mail.TemplateStorageValue, recipient *mail.Recipient) (status *mail.SendStatus, err error) {
	du.log.WithField("template", templateName).WithFields(recipient.Variables).Infoln("Sending email to", recipient.Email)
	status = &mail.SendStatus{
		RecipientID:  recipient.ID,
		TemplateName: templateName,
		Status:       "OK",
	}
	return
}