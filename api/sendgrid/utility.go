// メールユーティリティパッケージ
// メール組立に必要な処理を実装
package sendgrid

import (
	"io/ioutil"
	"os"

	"github.com/boiler-plate/tools/logger"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Build
// メール情報の組立
// param batchID : バッチID
// return メールのリクエスト内容
func (info *Info) Build(batchID string) []byte {
	// メール基本情報
	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail("yusuke", os.Getenv("SENDGRID_SENDER_EMAIL")))
	m.Subject = info.Subject

	// 送信対象をセット
	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail(info.Target.Name, info.Target.Email))
	// 送信対象者用に置換文字列をセット
	info.SetConv(p)
	// 送信対象をメール情報にセット
	m.AddPersonalizations(p)

	// Content-Typeの設定(HTML形式を優先)
	contentType := "text/plain"
	// ファイルパスの設定
	fp := info.PathPlain
	if info.PathHtml != "" {
		contentType = "text/html"
		fp = info.PathHtml
	}

	// テンプレートファイルの読み込み
	buf, err := ioutil.ReadFile(fp)
	if err != nil {
		_ = logger.Log.Errorf("テンプレートファイルの読み込みに失敗しました/%w", err)
	}

	// メール内容のセット
	content := mail.NewContent(contentType, string(buf))
	m.AddContent(content)

	// バッチIDをセット
	m.SetBatchID(batchID)

	return mail.GetRequestBody(m)
}

// SetConv
// 置換文字列のセット
// param p : 送信対象情報の設定先
func (info *Info) SetConv(p *mail.Personalization) {
	// 送信対象
	t := info.Target

	// メール本文の置換
	p.SetSubstitution("{%name%}", t.Name)
	p.SetSubstitution("{%email%}", t.Email)
	p.SetSubstitution("{%url%}", "")
	p.SetSubstitution("{%pass_reset_url%}", "")
	p.SetSubstitution("{%unsubscribe_url%}", "")
}
