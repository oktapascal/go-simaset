package config

import (
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/spf13/viper"
)

func SetupMailjetClient() *mailjet.Client {
	client := mailjet.NewMailjetClient(viper.GetString("MJ_APIKEY_PUBLIC"), viper.GetString("MJ_APIKEY_PRIVATE"))

	return client
}
