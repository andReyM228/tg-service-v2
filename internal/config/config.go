package config

import (
	"fmt"
	"github.com/andReyM228/lib/redis"
	"github.com/andReyM228/one/chain_client"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Config struct {
		Chain  chain_client.ClientConfig `yaml:"chain"`
		TgBot  TgBot                     `yaml:"tg-bot" validate:"required"`
		HTTP   HTTP                      `yaml:"http" validate:"required"`
		Extra  Extra                     `yaml:"extra" validate:"required"`
		Rabbit Rabbit                    `yaml:"rabbit" validate:"required"`
		Redis  redis.Config              `yaml:"redis"`
	}

	TgBot struct {
		Token string `yaml:"token" validate:"required"`
	}

	ChatGPT struct {
		Key   string `yaml:"key" validate:"required"`
		Model string `yaml:"model" validate:"required"`
	}

	Rabbit struct {
		RabbitUrl string `yaml:"url_rabbit" validate:"required"`
	}

	HTTP struct {
		Port int `yaml:"port" validate:"required"`
	}

	Extra struct {
		CarPaymentAddress string `yaml:"car_payment_address" validate:"required"`
		UrlGetAllCars     string `yaml:"url_get_all_cars" validate:"required"`
		UrlGetUserCars    string `yaml:"url_get_user_cars" validate:"required"`
		UrlBuyCar         string `yaml:"url_buy_car" validate:"required"`
		UrlSellCar        string `yaml:"url_sell_car" validate:"required"`
	}
)

func ParseConfig() (Config, error) {
	file, err := os.ReadFile("./cmd/config.yaml")
	if err != nil {
		fmt.Errorf("parseConfig: %s", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		fmt.Errorf("parseConfig: %s", err)
	}

	return cfg, nil
}
