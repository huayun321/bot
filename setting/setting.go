package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")   // name of config file (without extension)
	vp.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	vp.AddConfigPath("configs/") // path to look for the config file in
	err := vp.ReadInConfig()     // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		return nil, err
	}
	return &Setting{vp: vp}, err
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	return err
}
