// メール送信用パッケージ
package sendgrid

import (
	"context"

	checkcancel "github.com/boiler-plate/tools/check_cancel"
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

// リマインド配信用構造体
type Remind struct {
	Client *Client // API通信用クライアント
	Info   *Info   // メール基本情報
}

// お知らせ配信用構造体
type Notification struct {
	Client *Client // API通信用クライアント
	Info   *Info   // メール基本情報
}

// パスワードリセット配信用構造体
type PassReset struct {
	Client *Client
	Info   *Info
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
// param subject : メール件名
// param html : HTMLメールテンプレートのパス
// param plain : テキストメールテンプレートのパス
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

// NewRemind
// Remind構造体をMailインターフェースで生成
// param subject : メール件名
// param html : HTMLメールテンプレートのパス
// param plain : テキストメールテンプレートのパス
// return Mailインターフェース
func (t *Target) NewRemind(subject, html, plain string) Mail {
	return &Remind{
		Client: NewClient(),
		Info: &Info{
			Target:    t,
			Subject:   subject,
			PathHtml:  html,
			PathPlain: plain,
		},
	}
}

// NewNotification
// Notification構造体をMailインターフェースで生成
// param subject : メール件名
// param html : HTMLメールテンプレートのパス
// param plain : テキストメールテンプレートのパス
// return Mailインターフェース
func (t *Target) NewNotification(subject, html, plain string) Mail {
	return &Notification{
		Client: NewClient(),
		Info: &Info{
			Target:    t,
			Subject:   subject,
			PathHtml:  html,
			PathPlain: plain,
		},
	}
}

// NewPassReset
// PassReset構造体をMailインターフェースで生成
// param subject : メール件名
// param html : HTMLメールテンプレートのパス
// param plain : テキストメールテンプレートのパス
// return Mailインターフェース
func (t *Target) NewPassReset(subject, html, plain string) Mail {
	return &PassReset{
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

	// キャンセルの確認
	checkcancel.Exec(ctx)
	// バッチIDを生成
	batchID, err := c.CreateBatchID(ctx)
	if err != nil {
		return err
	}

	// キャンセルの確認
	checkcancel.Exec(ctx)
	// batchIDの有効チェック
	if err := c.ValidateBatchID(ctx, batchID); err != nil {
		return err
	}

	// キャンセルの確認
	checkcancel.Exec(ctx)

	// メール情報を組立
	reqBody := info.Build(batchID)

	// メール送信
	return c.Send(ctx, reqBody)
}

// (r *Remind) Send
// リマインド配信処理
// param ctx : コンテキスト
// return エラー情報
func (r *Remind) Send(ctx context.Context) error {
	c, info := r.Client, r.Info

	// キャンセルの確認
	checkcancel.Exec(ctx)
	// バッチIDを生成
	batchID, err := c.CreateBatchID(ctx)
	if err != nil {
		return err
	}

	// キャンセルの確認
	checkcancel.Exec(ctx)
	// batchIDの有効チェック
	if err := c.ValidateBatchID(ctx, batchID); err != nil {
		return err
	}

	// キャンセルの確認
	checkcancel.Exec(ctx)

	// メール情報を組立
	reqBody := info.Build(batchID)

	// メール送信
	return c.Send(ctx, reqBody)
}

// (n *Notification) Send
// お知らせ配信処理
// param ctx : コンテキスト
// return エラー情報
func (n *Notification) Send(ctx context.Context) error {
	c, info := n.Client, n.Info

	// キャンセルの確認
	checkcancel.Exec(ctx)
	// バッチIDを生成
	batchID, err := c.CreateBatchID(ctx)
	if err != nil {
		return err
	}

	// キャンセルの確認
	checkcancel.Exec(ctx)
	// batchIDの有効チェック
	if err := c.ValidateBatchID(ctx, batchID); err != nil {
		return err
	}

	// キャンセルの確認
	checkcancel.Exec(ctx)

	// メール情報を組立
	reqBody := info.Build(batchID)

	// メール送信
	return c.Send(ctx, reqBody)
}

// (ps *PassReset) Send
// パスワードリセット配信処理
// param ctx : コンテキスト
// return エラー情報
func (ps *PassReset) Send(ctx context.Context) error {
	c, info := ps.Client, ps.Info

	// キャンセルの確認
	checkcancel.Exec(ctx)
	// バッチIDを生成
	batchID, err := c.CreateBatchID(ctx)
	if err != nil {
		return err
	}

	// キャンセルの確認
	checkcancel.Exec(ctx)
	// batchIDの有効チェック
	if err := c.ValidateBatchID(ctx, batchID); err != nil {
		return err
	}

	// キャンセルの確認
	checkcancel.Exec(ctx)

	// メール情報を組立
	reqBody := info.Build(batchID)

	// メール送信
	return c.Send(ctx, reqBody)
}