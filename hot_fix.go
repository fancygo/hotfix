package hotfix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Car  int    `json:"car"`
}

func loadConfig(config *Config) {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("load config error: ", err)
	}
	err = json.Unmarshal(f, config)
	if err != nil {
		fmt.Println("Para config failed: ", err)
	}
}

/*
func main() {
	config := new(Config)
	loadConfig(config)
	fmt.Println(*config)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGUSR2)
	go func() {
		for {
			fmt.Println("wait sig")
			<-sig
			loadConfig(config)
			fmt.Println(*config)
		}
	}()
	time.Sleep(time.Second * 1000)
}
*/
