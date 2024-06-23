package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/WatchYourPorts/internal/check"
	"github.com/aceberg/WatchYourPorts/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8853")
	viper.SetDefault("THEME", "grass")
	viper.SetDefault("COLOR", "dark")
	viper.SetDefault("TIMEOUT", 10)
	viper.SetDefault("HIST_TRIM", 90)
	viper.SetDefault("INFLUX_ENABLE", false)

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.Color, _ = viper.Get("COLOR").(string)
	config.Timeout = viper.GetInt("TIMEOUT")
	config.HistTrim = viper.GetInt("HIST_TRIM")
	config.InfluxEnable = viper.GetBool("INFLUX_ENABLE")
	config.InfluxSkipTLS = viper.GetBool("INFLUX_SKIP_TLS")
	config.InfluxAddr, _ = viper.Get("INFLUX_ADDR").(string)
	config.InfluxToken, _ = viper.Get("INFLUX_TOKEN").(string)
	config.InfluxOrg, _ = viper.Get("INFLUX_ORG").(string)
	config.InfluxBucket, _ = viper.Get("INFLUX_BUCKET").(string)

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("color", config.Color)
	viper.Set("timeout", config.Timeout)
	viper.Set("hist_trim", config.HistTrim)

	viper.Set("influx_enable", config.InfluxEnable)
	viper.Set("influx_skip_tls", config.InfluxSkipTLS)
	viper.Set("influx_addr", config.InfluxAddr)
	viper.Set("influx_token", config.InfluxToken)
	viper.Set("influx_org", config.InfluxOrg)
	viper.Set("influx_bucket", config.InfluxBucket)

	err := viper.WriteConfig()
	check.IfError(err)
}
