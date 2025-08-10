package exchange

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
)

func WsReconnectWait(exchange string, reason string) {
	log.Info(fmt.Sprintf("%s reconneting...", exchange))
	log.Info(fmt.Sprintf("Reason %s", reason))
	time.Sleep(ReconnectInterval)
}
