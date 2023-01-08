package utils

import "github.com/pquerna/otp/totp"

func GenerateOtp() (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "depploy.io",
		AccountName: "hello@depploy.io",
	})
	if err != nil {
		return "", err
	}

	return key.Secret(), nil
}
