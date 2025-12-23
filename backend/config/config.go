package config

import "smartPet/backend/pkg/viperx"

type MainWindowConf struct {
	Mask   string `yaml:"mask" mapstructure:"mask"`
	X      int    `yaml:"x" mapstructure:"x"`
	Y      int    `yaml:"y" mapstructure:"y"`
	Width  int    `yaml:"width" mapstructure:"width"`
	Height int    `yaml:"height" mapstructure:"height"`
}

func NewMainWindowMaskConf(v *viperx.ViperSetting) *MainWindowConf {
	var m = &MainWindowConf{}
	err := v.ReadSection("MainWindow", m)
	if err != nil {
		panic(err)
	}
	return m
}
