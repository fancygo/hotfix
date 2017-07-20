package hotfix

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestHotfix(t *testing.T) {
	config := new(Config)
	loadConfig(config)
	fmt.Println(*config)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGUSR2)
	go func() {
		for {
			<-sig
			loadConfig(config)
			fmt.Println(*config)
		}
	}()
	time.Sleep(time.Second * 1000)
}
