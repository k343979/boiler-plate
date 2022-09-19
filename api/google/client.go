// GoogleAPIクライアント用パッケージ
package google

import (
	"os"

	"golang.org/x/oauth2"
	g "golang.org/x/oauth2/google"
)

// API通信用クライアント
type Client struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	EndPoint     oauth2.Endpoint
}

// API通信用クライアント構造体を生成
// return *Client
func NewClient() *Client {
	return &Client{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		EndPoint:     g.Endpoint,
	}
}