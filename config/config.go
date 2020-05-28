package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var TimeWheelConfig = &TmWheelConfig{
	Millisecond: 20,
	Task:        map[string]int{ //每隔模块
	},
}

// LogConfig is the Configuration of log
type TmWheelConfig struct {
	Millisecond int64          `json:"millisecond"`
	Task        map[string]int `json:"task"`
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func LoadTmWheelConfig(dir string) (*TmWheelConfig, error) {
	path := filepath.Join(dir, "./configs/tmwheel.toml")
	filePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	config := new(TmWheelConfig)
	if CheckFileIsExist(filePath) { //文件存在
		if _, err := toml.DecodeFile(filePath, config); err != nil {
			return nil, err
		} else {
			TimeWheelConfig = config
		}
	} else {
		configBuf := new(bytes.Buffer)
		if err := toml.NewEncoder(configBuf).Encode(TimeWheelConfig); err != nil {
			return nil, err
		}
		err := ioutil.WriteFile(filePath, configBuf.Bytes(), 0666)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
