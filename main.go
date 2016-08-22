package madek

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

var ErrInvalidAuthentication = errors.New("Invalid Authentication")
var ErrAccessForbidden = errors.New("Access Forbidden")
var ErrRequestFailed = errors.New("Request Failed")
var ErrNotFound = errors.New("Not Found")

type Client struct {
	Address  string
	Username string
	Password string
}

func NewClient(address, username, password string) *Client {
	return &Client{
		Address:  address,
		Username: username,
		Password: password,
	}
}

func (c *Client) fetch(path string) (string, error) {
	println(path)

	res, str, err := gorequest.New().Get(path).
		SetBasicAuth(c.Username, c.Password).
		Set("Accept", "application/json-roa+json").
		End()

	if len(err) > 0 {
		return "", err[0]
	}

	if res.StatusCode == http.StatusUnauthorized {
		return "", ErrInvalidAuthentication
	}

	if res.StatusCode == http.StatusForbidden {
		return "", ErrAccessForbidden
	}

	if res.StatusCode == http.StatusNotFound {
		return "", ErrNotFound
	}

	if res.StatusCode != http.StatusOK {
		return "", ErrRequestFailed
	}

	return str, nil
}

func (c *Client) url(format string, args ...interface{}) string {
	args = append([]interface{}{c.Address}, args...)
	return fmt.Sprintf("%s"+format, args...)
}
