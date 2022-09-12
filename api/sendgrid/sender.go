// メール送信用パッケージ
package sendgrid

import (
	"context"
)

// 送信対象情報
type Target struct {
	Name  string
	Email string
}

// メール基本情報
type Info struct {
	Target    *Target // 送信対象情報
	Subject   string  // メール件名
	PathHtml  string  // HTMLメール本文のテンプレートパス
	PathPlain string  // テキストメール本文のテンプレートパス
}

// テスト配信用構造体
type Test struct {
	Client *Client // API通信用クライアント
	Info   *Info   // メール基本情報
}

// メールインターフェース
type Mail interface {
	Send(context.Context) error // メール送信処理
}

// NewTarget
// Target構造体を生成
// param name : 送信対象者名
// param email : 送信先メールアドレス
// return *Target
func NewTarget(name, email string) *Target {
	return &Target{
		Name:  name,
		Email: email,
	}
}

// NewTest
// Test構造体をMailインターフェースで生成
// return Mailインターフェース
func (t *Target) NewTest(subject, html, plain string) Mail {
	return &Test{
		Client: NewClient(),
		Info: &Info{
			Target:    t,
			Subject:   subject,
			PathHtml:  html,
			PathPlain: plain,
		},
	}
}

// (t *Test) Send
// テスト配信処理
// param ctx : コンテキスト
// return エラー情報
func (t *Test) Send(ctx context.Context) error {
	c, info := t.Client, t.Info
	// バッチIDを生成
	batchID, err := c.CreateBatchID(ctx)
	if err != nil {
		return err
	}

	// batchIDの有効チェック
	if err := c.ValidateBatchID(ctx, batchID); err != nil {
		return err
	}

	// メール情報を組立
	reqBody := info.Build(batchID)

	// メール送信
	return c.Send(ctx, reqBody)
}