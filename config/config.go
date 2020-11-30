package config

var twilioBaseURL = "https://api.twilio.com/2010-04-01/Accounts/"

var (
	emailAccount,
	password,
	accountSid,
	authToken,
	twilioNumber,
	personalNumber,
	uploadURL,
	serverURL string
)

// EmailAccount .
func EmailAccount() string {
	return emailAccount
}

// Password .
func Password() string {
	return password
}

// TwilioURL .
func TwilioURL() string {
	return twilioBaseURL + accountSid + "/Messages.json"
}

// AccountSid .
func AccountSid() string {
	return accountSid
}

//AuthToken .
func AuthToken() string {
	return authToken
}

// TwilioNumber .
func TwilioNumber() string {
	return twilioNumber
}

// PersonalNumber .
func PersonalNumber() string {
	return personalNumber
}

// UploadURL .
func UploadURL() string {
	return uploadURL
}

// ServerURL .
func ServerURL() string {
	return serverURL
}
