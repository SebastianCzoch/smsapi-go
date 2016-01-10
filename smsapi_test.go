package smsapi

import (
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	object := New("testuser", "testpasswd")
	assert.Equal(t, &Client{"testuser", "testpasswd"}, object)
}

func TestGetPasswordHash(t *testing.T) {
	object := New("testuser", "testpasswd")
	assert.Equal(t, "59de1412ec33fd96ac4a4bfc793f1133", object.getPasswordHash())
}

func TestGetUrl(t *testing.T) {
	assert.Equal(t, "https://api.smsapi.pl/test_action", getURL("test_action"))
}

func TestGetBasicValuesForRequest(t *testing.T) {
	object := New("testuser", "testpasswd")
	excepted := url.Values{}
	excepted.Add("username", "testuser")
	excepted.Add("password", "59de1412ec33fd96ac4a4bfc793f1133")
	excepted.Add("encoding", "utf-8")
	assert.Equal(t, excepted, object.getBasicValuesForRequest())
}

func TestSendEcoSMSError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://api.smsapi.pl/sms.do", httpmock.NewStringResponder(200, `ERROR:101`))

	object := New("testuser", "testpasswd")
	err := object.SendEcoSMS("XXXXXXXXX", "test message")
	assert.Error(t, err)
}

func TestSendEcoSMSSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://api.smsapi.pl/sms.do", httpmock.NewStringResponder(200, `OK:1013211`))

	object := New("testuser", "testpasswd")
	err := object.SendEcoSMS("XXXXXXXXX", "test message")
	assert.NoError(t, err)
}
