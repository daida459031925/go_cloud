package connect

import (
	"context"
	"time"
)

//  示例
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
//	defer cancel()
//
//	for {
//		select {
//		case <- ctx.Done():
//			fmt.Println("timeout")
//			return
//		default:
//			fmt.Println("waiting...")
//			time.Sleep(time.Second)
//		}
//	}
//}

// 执行超时，取消执行
func GetContext(i ...time.Duration) (context.Context, context.CancelFunc) {
	var t time.Duration = 5
	if i != nil && len(i) > 0 && i[0] > 0 {
		t = i[0]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*t)
	return ctx, cancel
}
