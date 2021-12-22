// Package restful package restful
package restful

import (
	"github.com/go-resty/resty/v2"
)

// Client Client
var Client *resty.Client

// GetClient GetClient
func GetClient() *resty.Client {
	if Client != nil {
		return Client
	}
	client := resty.New()
	Client = client
	return Client
}
