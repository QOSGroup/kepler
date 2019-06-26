package config

type MailConfig struct {
	Identity string
	Host     string
	Port     string
	Sender   string
	Password string
}

func DefaultMailConfig() *MailConfig {
	return &MailConfig{
		"", "smtp.mxhichina.com", "25","", "",
	}
}
