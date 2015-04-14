// Package gotwilio is a library for interacting with http://www.twilio.com/ API.
package gotwilio

import (
	"net/http"
	"net/url"
	"strings"
)

// Twilio stores basic information important for connecting to the
// twilio.com REST api such as AccountSid and AuthToken.
type Twilio struct {
	AccountSid string
	AuthToken  string
	BaseUrl    string
	HTTPClient *http.Client
}

// Exception is a representation of a twilio exception.
type Exception struct {
	Status   int    `json:"status"`    // HTTP specific error code
	Message  string `json:"message"`   // HTTP error message
	Code     int    `json:"code"`      // Twilio specific error code
	MoreInfo string `json:"more_info"` // Additional info from Twilio
}

// Create a new Twilio struct.
func NewTwilioClient(accountSid, authToken string, HTTPClient *http.Client) *Twilio {
	twilioUrl := "https://api.twilio.com/2010-04-01" // Should this be moved into a constant?
	return &Twilio{accountSid, authToken, twilioUrl, HTTPClient}
}

func (twilio *Twilio) post(formValues url.Values, twilioUrl string) (*http.Response, error) {
	resp, err := http.NewRequest("POST", twilioUrl, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}
	resp.SetBasicAuth(twilio.AccountSid, twilio.AuthToken)
	resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := twilio.HTTPClient

	return client.Do(resp)
}

func (twilio *Twilio) GetMessageResponse(messageSid string) (msgResponse *MessageResponse, exc *Exception, err error) {
	msgResponse, exc, err = twilio.getMessage(messageSid)
	return
}

func (twilio *Twilio) get(formValues url.Values, twilioUrl string) (*http.Response, error) {
	resp, err := http.NewRequest("GET", twilioUrl, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}
	resp.SetBasicAuth(twilio.AccountSid, twilio.AuthToken)
	resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := twilio.HTTPClient

	return client.Do(resp)
}
