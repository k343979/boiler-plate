// SlackAPIクライアント用パッケージ
package slack

import (
	"os"
)

// API通信用クライアント
type Client struct {
	BotAccessToken  string // ボット用アクセストークン
	UserAccessToken string // ユーザー用アクセストークン
}

// API通信用クライアント構造体を生成
// return *Client
func NewClient() *Client {
	return &Client{
		BotAccessToken:  os.Getenv("SLACK_BOT_ACCESS_TOKEN"),
		UserAccessToken: os.Getenv("SLACK_USER_ACCESS_TOKEN"),
	}
}
