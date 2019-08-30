package config

import "fmt"

func DefaultDbConfig() *DbConfig {
	return &DbConfig{
		"mysql", "qos_sell", "qos_sell", "192.168.1.235", 3306, "kepler",
	}
}

type DbConfig struct {
	Driver   string
	Name     string
	Password string
	Ip       string
	Port     int
	Database string
}

func (config *DbConfig) DateSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&autocommit=true", config.Name, config.Password, config.Ip, config.Port, config.Database)
}
