package history

import (
	"fmt"
	"time"
)

func ReportDuplicate(try int, maxTries int, poolSize int64) {
	fmt.Printf("[HISTORY] DUPLICATE | TRY: %d/%d | POOL: %d | TTL: %dh\n", try, maxTries, poolSize, ttl/time.Hour)
}
