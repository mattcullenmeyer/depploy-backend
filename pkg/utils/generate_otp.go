package utils

import "github.com/pquerna/otp/totp"

type GenerateOtpParams struct {
	Email string
}

func GenerateOtp(args GenerateOtpParams) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "depploy.io",
		AccountName: args.Email,
	})
	if err != nil {
		return "", err
	}

	return key.Secret(), nil
}
