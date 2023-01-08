package utils

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendConfirmationEmailParams struct {
	Otp      string
	Username string
	Email    string
}

// https://github.com/sendgrid/sendgrid-go/blob/main/use-cases/substitutions-with-mailer-helper.md
func SendConfirmationEmail(args SendConfirmationEmailParams) error {
	m := mail.NewV3Mail()

	address := "hello@depploy.io"
	name := "Depploy"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	templateId, err := GetParameter("SendGridEmailVerificationTemplate")
	if err != nil {
		return err
	}

	m.SetTemplateID(templateId)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(args.Username, args.Email),
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
