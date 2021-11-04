package xtime

import (
	"time"
)

// TZ8 TimeZone 8[东8区（也就是：北京时间）]
var TZ8 *time.Location = time.FixedZone("CST", 3600*8)
