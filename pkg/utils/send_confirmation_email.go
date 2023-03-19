package utils

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendConfirmationEmailParams struct {
	Otp   string
	Email string
}

// https://github.com/sendgrid/sendgrid-go/blob/main/use-cases/substitutions-with-mailer-helper.md
func SendConfirmationEmail(args SendConfirmationEmailParams) error {
	m := mail.NewV3Mail()

	fromName := "Depploy"
	fromAddress := "hello@depploy.io"
	e := mail.NewEmail(fromName, fromAddress)
	m.SetFrom(e)

	templateId, err := GetParameter("SendGridEmailVerificationTemplate")
	if err != nil {
		return err
	}

	m.SetTemplateID(templateId)

	p := mail.NewPersonalization()
	toName := "" // TODO: Fill in something for name
	toAddress := args.Email
	tos := []*mail.Email{
		mail.NewEmail(toName, toAddress),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("otp", args.Otp)

	m.AddPersonalizations(p)

	apiKey, err := GetParameter("SendGridApiKey")
	if err != nil {
		return err
	}

	request := sendgrid.GetRequest(apiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	response, err := sendgrid.API(request)
	fmt.Println(response.Body)
	if err != nil {
		return err
	}

	return nil
}
