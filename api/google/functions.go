// GoogleのAPI処理用パッケージ
package google

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// Auth
// param ctx : コンテキスト
// return レスポンス
// return エラー情報
func Auth(ctx context.Context) (*http.Response, error) {
	// クライアント生成
	c := NewClient()
	// configのセット
	conf := c.NewConfig("openid", "email", "profile")
	// ユーザーをGoogle同意ページにリダイレクトし、上記で指定したスコープの許可を求める
	url := conf.AuthCodeURL("state")
	// 交換コードを処理してトランスポートを開始
	tok, err := conf.Exchange(ctx, "authorization-code", oauth2.AccessTypeOffline)
	if err != nil {
		return nil, err
	}

	client := conf.Client(ctx, tok)
	return client.Get(url)
}