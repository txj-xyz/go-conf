package config

type Server struct {
	Port   int    `json:"port"`
	Domain string `json:"domain"`
	HTTPS  bool   `json:"https"`
}

type Logging struct {
	Enabled bool          `json:"enabled"`
	Discord DiscordModule `json:"discord"`
}

type DiscordModule struct {
	Enabled bool   `json:"enabled"`
	Webhook string `json:"webhook"`
}

type Debug struct {
	Enabled bool `json:"enabled"`
}

type ConfigBase struct {
	Server  Server  `json:"server"`
	Logging Logging `json:"logging"`
	Debug   Debug   `json:"debug"`
}
