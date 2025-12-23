package viperx

import (
	"fmt"
	"github.com/spf13/viper"
)

// 负责读取配置文件的核心逻辑

type ViperSetting struct {
	*viper.Viper
}

func (s *ViperSetting) ReadSection(k string, v interface{}) error {
	err := s.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

func NewViperSetting(configPath string) *ViperSetting {
	vp := viper.New()
	vp.SetConfigFile(configPath) // 指定配置文件路径
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Read confg err:%v", err))
	}
	return &ViperSetting{
		Viper: vp,
	}
}
