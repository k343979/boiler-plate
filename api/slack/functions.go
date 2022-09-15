// SlackのAPI処理用パッケージ
package slack

import (
	"context"

	"github.com/slack-go/slack"
)

// メッセージ構造体
type Msg struct {
	Type string
	Text string
}

// 送信情報の構造体
type Info struct {
	Channel  string
	IsNotify bool
	Msgs     []*Msg
}

// NewMsg
// メッセージ構造体を生成
// return *Msg
func NewMsg(msgType, text string) *Msg {
	return &Msg{
		Type: msgType,
		Text: text,
	}
}

// NewInfo
// 送信情報の構造体を生成
// return *Info
func NewInfo(channel string, isNotify bool, msgs ...*Msg) *Info {
	return &Info{
		Channel:  channel,
		IsNotify: isNotify,
		Msgs:     msgs,
	}
}

// PostTestMsg
// メッセージのテスト送信
// param ctx : コンテキスト
// return エラー情報
func (c *Client) PostTestMsg(ctx context.Context) error {
	sc := slack.New(c.BotAccessToken)
	_, _, err := sc.PostMessageContext(ctx, "#app_linked", slack.MsgOptionBlocks(
		slack.NewSectionBlock(
			slack.NewTextBlockObject("mrkdwn", "<!here>", false, false),
			[]*slack.TextBlockObject{
				slack.NewTextBlockObject("plain_text", "Hello World!", false, false),
			},
			nil,
		),
	))
	if err != nil {
		return err
	}

	return nil
}

// PostMsg
// メッセージの送信
// param ctx : コンテキスト
// param channel : 送信先チャンネル
// param msg : 送信メッセージのスライス
// return エラー情報
func (c *Client) PostMsg(ctx context.Context, info *Info) error {
	sc := slack.New(c.BotAccessToken)
	// 通知メンションブロック
	var notifyBlock *slack.TextBlockObject
	// 通知する場合
	if info.IsNotify {
		notifyBlock = slack.NewTextBlockObject("mrkdwn", "<!here>", false, false)
	}

	// メッセージブロック
	msgBlock := make([]*slack.TextBlockObject, 0, len(info.Msgs))
	for _, v := range info.Msgs {
		v := v

		txtBlock := slack.NewTextBlockObject(v.Type, v.Text, false, false)
		msgBlock = append(msgBlock, txtBlock)
	}

	// メッセージ送信
	_, _, err := sc.PostMessageContext(ctx, info.Channel, slack.MsgOptionBlocks(
		slack.NewSectionBlock(
			notifyBlock,
			msgBlock,
			nil,
		),
	))
	if err != nil {
		return err
	}

	return nil
}
