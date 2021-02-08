package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	cfgFile = ""
)

// splash print plain text message to console
func splash() {
	fmt.Println(`
╱╱╱╱╱╭┳╮╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╭╮
╱╱╱╱╱┃┃┃╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱┃┃
╭━━┳━╯┃╰━┳━━┳╮╭┳━━┳━━╮╱╱╭━━┳━━┳╮╭┳━━┫┃╭━━╮
┃╭╮┃╭╮┃╭╮┃╭╮┃┃┃┃━━┫┃━╋━━┫━━┫╭╮┃╰╯┃╭╮┃┃┃┃━┫
┃╭╮┃╰╯┃┃┃┃╰╯┃╰╯┣━━┃┃━╋━━╋━━┃╭╮┃┃┃┃╰╯┃╰┫┃━┫
╰╯╰┻━━┻╯╰┻━━┻━━┻━━┻━━╯╱╱╰━━┻╯╰┻┻┻┫╭━┻━┻━━╯
╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱┃┃
╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╰╯	`)
}

func initializeConfig() {
	viper.SetConfigType("toml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// search config in home directory with name "config" (without extension)
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
	}

	//read env
	viper.AutomaticEnv()

	// if a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Config application:", err)
	}

	log.Println("using config file:", viper.ConfigFileUsed())
}
