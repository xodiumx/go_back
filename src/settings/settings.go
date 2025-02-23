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

// InitConfig инициализирует глобальную переменную конфигурации
func InitConfig() {

	// Загружает переменные из .env в окружение
	if err := godotenv.Load(); err != nil {
		log.Println("Ошибка загрузки .env файла, продолжаем использовать только переменные окружения")
	}

	// Проверка, что переменные загружены в окружение
	fmt.Println("Loaded in env ", os.Getenv("DB_HOST"))

	viper.AutomaticEnv() // Автоматически загружает переменные окружения

	// Устанавливаем значения по умолчанию
	viper.SetDefault("FROM_CONFIG_FILE", false)

	// Пробуем прочитать конфиг из файла (если есть)
	viper.SetConfigName("config")     // Имя файла без расширения
	viper.SetConfigType("yaml")       // Формат файла
	viper.AddConfigPath("./settings") // Каталог для поиска файла

	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Используется конфигурационный файл: %s", viper.ConfigFileUsed())
	} else {
		log.Println("Файл конфигурации не найден, используются только ENV переменные")
	}

	// Маппинг переменных в структуру
	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	} else {
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
