// GoogleAPIのconfig設定用パッケージ
package google

import (
	"golang.org/x/oauth2"
)

// NewConfig
// oauth2のConfigを生成
// param scopes : スコープ
// return *oauth.Config
func (c *Client) NewConfig(scopes ...string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Scopes:       scopes,
		Endpoint:     c.EndPoint,
	}
}