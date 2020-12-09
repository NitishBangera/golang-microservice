package model

// Config of the application defined globally
type Config struct {
	values map[string]string
}

// New method creates a Handler object.
func New(values map[string]string) *Config {
	configs := make(map[string]string, len(values))

	for k, v := range values {
		configs[k] = v
	}
	return &Config{values: configs}
}

func (config Config) GetValue(key string) string {
	return config.values[key]
}
