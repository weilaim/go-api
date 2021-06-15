package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKdy 每日排行
	DailyRankKdy = "rank:daily"
)

// VideoViewKey 点击视频的key
//view:video：1 -》 100
//view:video: 2 -> 150
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}