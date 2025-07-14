package workwx

import (
	"github.com/spf13/viper"
)

type WeChatWorkConfig struct {
	CorpID     string
	CorpSecret string
	AgentID    int64
}

func LoadConfig() (*WeChatWorkConfig, error) {
	// 使用 viper 读取配置
	viper.SetConfigName("config") // 配置文件名 (不带扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的扩展名是 yaml
	viper.AddConfigPath(".")      // 查找配置文件所在的路径

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &WeChatWorkConfig{
		CorpID:     viper.GetString("wechat_work.corp_id"),
		CorpSecret: viper.GetString("wechat_work.corp_secret"),
		AgentID:    viper.GetInt64("wechat_work.agent_id"),
	}, nil
}
