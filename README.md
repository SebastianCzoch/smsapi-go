# smsapi-go
[![Build Status](https://travis-ci.org/SebastianCzoch/smsapi-go.svg?branch=master)](https://travis-ci.org/SebastianCzoch/smsapi-go) [![Code Climate](https://codeclimate.com/github/SebastianCzoch/smsapi-go/badges/gpa.svg)](https://codeclimate.com/github/SebastianCzoch/smsapi-go) [![Coverage Status](https://coveralls.io/repos/SebastianCzoch/smsapi-go/badge.svg?branch=master&service=github)](https://coveralls.io/github/SebastianCzoch/smsapi-go?branch=master)  [![GoDoc](https://godoc.org/github.com/SebastianCzoch/smsapi-go?status.svg)](https://godoc.org/github.com/SebastianCzoch/smsapi-go)  [![License](https://img.shields.io/badge/licence-MIT-green.svg)](./LICENSE)



Go utils for working with [SMSAPI](http://www.smsapi.pl/) service. This is not stable release, PLEASE DO NOT USING IT ON PRODUCTION

## Examples
All examples are already in example directory

## Install

```
$ go get github.com/SebastianCzoch/smsapi-go
````

or via [Godep](https://github.com/tools/godep)
```
$ godep get github.com/SebastianCzoch/smsapi-go
```


## API

### New(username, password string) *Client
Return a pointer to new smsapi client

### (c *Client) SendEcoSMS(to, body string) error
Send Eco sms to selected phone number with specified message

## Tests

```
$ go test ./...
````

## License

[MIT](./LICENSE)

## Support

Issues for this project should be reported on GitHub issues

Staff responsible for project:

* [Sebastian Czoch <sebastian@czoch.pl>](sebastian@czoch.pl)
