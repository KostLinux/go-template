package config

type New struct {
	App        *AppParams        `mapstructure:"app"`
	HTTP       *HTTPParams       `mapstructure:"http"`
	Database   *DatabaseParams   `mapstructure:"database"`
	Middleware *MiddlewareParams `mapstructure:"middleware"`
	Monitoring *MonitoringParams `mapstructure:"monitoring"`
}

type AppParams struct {
	Name    string     `mapstructure:"name"`
	Version string     `mapstructure:"version"`
	Port    int        `mapstructure:"port"`
	Env     string     `mapstructure:"env"`
	Log     *LogParams `mapstructure:"log"`
}

type LogParams struct {
	Level string `mapstructure:"level"`
}

type HTTPParams struct {
	ReadTimeout    int `mapstructure:"readTimeout"`
	WriteTimeout   int `mapstructure:"writeTimeout"`
	IdleTimeout    int `mapstructure:"idleTimeout"`
	MaxHeaderBytes int `mapstructure:"maxHeaderBytes"`
}

type DatabaseParams struct {
	Driver     string            `mapstructure:"driver"`
	Host       string            `mapstructure:"host"`
	Port       int               `mapstructure:"port"`
	User       string            `mapstructure:"user"`
	Password   string            `mapstructure:"password"`
	Name       string            `mapstructure:"name"`
	SSLMode    string            `mapstructure:"sslMode"`
	Connection *ConnectionParams `mapstructure:"connection"`
}

type ConnectionParams struct {
	MaxIdleConns    int `mapstructure:"maxIdleConns"`
	MaxOpenConns    int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`
	ConnMaxIdleTime int `mapstructure:"connMaxIdleTime"`
}

type MiddlewareParams struct {
	Cors CorsParams `mapstructure:"cors"`
	Csrf CsrfParams `mapstructure:"csrf"`
}

type CorsParams struct {
	Enabled          bool     `mapstructure:"enabled"`
	AllowOrigins     []string `mapstructure:"allowOrigins"`
	AllowMethods     []string `mapstructure:"allowMethods"`
	AllowHeaders     []string `mapstructure:"allowHeaders"`
	ExposeHeaders    []string `mapstructure:"exposeHeaders"`
	MaxAge           int      `mapstructure:"maxAge"`
	AllowCredentials bool     `mapstructure:"allowCredentials"`
}

type CsrfParams struct {
	Enabled bool   `mapstructure:"enabled"`
	Key     string `mapstructure:"key"`
	MaxAge  int    `mapstructure:"maxAge"`
	Domain  string `mapstructure:"domain"`
}

type MonitoringParams struct {
	Telemetry *TelemetryParams `mapstructure:"telemetry"`
}

type TelemetryParams struct {
	Enabled          bool              `mapstructure:"enabled"`
	OTLPEndpoint     string            `mapstructure:"otlp_endpoint"`
	OTLPHeaders      map[string]string `mapstructure:"otlp_headers"`
	OTLPCompression  string            `mapstructure:"otlp_compression"`
	OTLPQueueSize    int               `mapstructure:"otlp_queue_size"`
	OTLPMaxBatchSize int               `mapstructure:"otlp_max_batch_size"`
	OTLPBatchTimeout int               `mapstructure:"otlp_batch_timeout"`
	OTLPInsecure     bool              `mapstructure:"otlp_insecure"`
	OTLPTimeout      int               `mapstructure:"otlp_timeout"`
}
