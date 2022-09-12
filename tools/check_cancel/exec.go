// キャンセルの確認用パッケージ
package checkcancel

import (
	"context"

	"github.com/boiler-plate/tools/logger"
)

// CheckCancel
// キャンセルの確認
// param ctx : コンテキスト
// return エラー情報
func CheckCancel(ctx context.Context) error {
	select {
	case <-ctx.Done():
		_ = logger.Log.Errorf("context was canceled")
		return ctx.Err()
	default:
		return nil
	}
}
