package smsapi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiProtocol     = "https"
	apiAddress      = "api.smsapi.pl"
	urlSchema       = "%s://%s/%s"
	sendSMSAction   = "sms.do"
	ecoSMS          = "Eco"
	defaultEncoding = "utf-8"
)

// Client is a basic struct for this library
type Client struct {
	username string
	password string
}

// New is a function which returns new Client struct
func New(username, password string) *Client {
	return &Client{username: username, password: password}
}

// SendEcoSMS is a method which is sending single SMS
func (c *Client) SendEcoSMS(to, body string) error {
	endpointURL := getURL(sendSMSAction)
	v := c.getBasicValuesForRequest()
	v.Add("to", to)
	v.Add("message", body)
	v.Add("from", ecoSMS)

	res, err := http.PostForm(endpointURL, v)
	if err != nil {
		return err
	}

	resInString, err := getResponseBodyAsString(res)
	if err != nil {
		return err
	}

	if strings.Contains(resInString, "ERROR") {
		return fmt.Errorf("API returns error: %s", resInString)
	}

	return nil
}

func (c *Client) getPasswordHash() string {
	hasher := md5.New()
	hasher.Write([]byte(c.password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getURL(action string) string {
	return fmt.Sprintf(urlSchema, apiProtocol, apiAddress, action)
}

func (c *Client) getBasicValuesForRequest() url.Values {
	v := url.Values{}
	v.Set("username", c.username)
	v.Set("password", c.getPasswordHash())
	v.Set("encoding", defaultEncoding)

	return v
}

func getResponseBodyAsString(response *http.Response) (string, error) {
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
