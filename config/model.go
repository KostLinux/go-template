package config

type New struct {
	App        AppConfig        `mapstructure:"app"`
	HTTP       HttpConfig       `mapstructure:"http"`
	Database   DatabaseConfig   `mapstructure:"database"`
	Middleware MiddlewareConfig `mapstructure:"middleware"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslMode"`
}

type HttpConfig struct {
	ReadTimeout    int `mapstructure:"readTimeout"`
	WriteTimeout   int `mapstructure:"writeTimeout"`
	IdleTimeout    int `mapstructure:"idleTimeout"`
	MaxHeaderBytes int `mapstructure:"maxHeaderBytes"`
}

type AppConfig struct {
	Name    string    `mapstructure:"name"`
	Version string    `mapstructure:"version"`
	Port    int       `mapstructure:"port"`
	Env     string    `mapstructure:"env"`
	Log     LogConfig `mapstructure:"log"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type MiddlewareConfig struct {
	Cors CorsConfig `mapstructure:"cors"`
	Csrf CsrfConfig `mapstructure:"csrf"`
}

type CorsConfig struct {
	Enabled          bool     `mapstructure:"enabled"`
	AllowOrigins     []string `mapstructure:"allowOrigins"`
	AllowMethods     []string `mapstructure:"allowMethods"`
	AllowHeaders     []string `mapstructure:"allowHeaders"`
	ExposeHeaders    []string `mapstructure:"exposeHeaders"`
	MaxAge           int      `mapstructure:"maxAge"`
	AllowCredentials bool     `mapstructure:"allowCredentials"`
}

type CsrfConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Key     string `mapstructure:"key"`
	MaxAge  int    `mapstructure:"maxAge"`
	Domain  string `mapstructure:"domain"`
}
