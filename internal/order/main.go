package main

import (
	"github.com/leebrouse/Gorder/common/config"
	"github.com/spf13/viper"
	"log"
)

// init order
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	log.Printf("%v", viper.Get("order"))
}
