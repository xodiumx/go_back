package settings

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	FromConfigFile bool `mapstructure:"FROM_CONFIG_FILE"`
	DB             struct {
		Host     string `mapstructure:"DB_HOST"`
		Port     int    `mapstructure:"DB_PORT"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASS"`
		Name     string `mapstructure:"DB_NAME"`
	} `mapstructure:",squash"`
}

var Conf Config

// InitConfig initialize global config
func InitConfig() {

	// Load env vars from .env to environment
	if err := godotenv.Load(); err != nil {
		log.Println("Error in load .env file")
	}

	// Check env var in environment
	fmt.Println("Loaded in env ", os.Getenv("DB_HOST"))

	viper.AutomaticEnv() // Use environment for settings

	// Can set default value
	viper.SetDefault("FROM_CONFIG_FILE", false)

	// Trying load config from config file
	viper.SetConfigName("config")     // Name
	viper.SetConfigType("yaml")       // Ext
	viper.AddConfigPath("./settings") // Path

	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Сonfiguration file is set: %s", viper.ConfigFileUsed())
	} else {
		log.Println("Config file not found")
	}

	// Deserialize environment in config
	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("Deserialize failed: %v", err)
	} else {
		// Check Conf
		fmt.Println("Config type ", Conf.FromConfigFile)
		fmt.Println("Host ", Conf.DB.Host)
		fmt.Println("Port ", Conf.DB.Port)
		fmt.Println("Name ", Conf.DB.Name)
		fmt.Println("User ", Conf.DB.User)
		fmt.Println("Password ", Conf.DB.Password)
	}
}

func init() {
	InitConfig()
}
