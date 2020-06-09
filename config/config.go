package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
)

type Config struct {
	Name string
}

func (c Config) initConfig() error {
	// 指定配置文件位置，若不指定则默认在 conf 下 config.yaml
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")     // 默认的配置文件格式
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取指定前缀的环境变量
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("config file changed: %s", e.Name)
	})
}

func Init(cfg string) error {
	c := Config{Name: cfg}
	if err := c.initConfig(); err != nil {
		return err
	}
	// 监控配置文件并热加载
	c.watchConfig()

	return nil
}
