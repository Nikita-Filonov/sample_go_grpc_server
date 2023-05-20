package config

type App struct {
	Port int32
}

type Logger struct {
	IsDevMode bool   `yaml:"is_dev_mode"`
	Level     string `yaml:"level"`
}

type Database struct {
	File string `yaml:"file"`
}

type Config struct {
	App      `yaml:"app" validate:"required"`
	Logger   `yaml:"logger" validate:"required"`
	Database `yaml:"database" validate:"required"`
}
