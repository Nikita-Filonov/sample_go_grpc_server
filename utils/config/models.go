package config

type App struct {
	Port int32
}

type Logger struct {
	IsDevMode bool   `yaml:"isDevMode"`
	Level     string `yaml:"level"`
}

type Database struct {
	File string `yaml:"file"`
}

type Config struct {
	App      App      `yaml:"app" validate:"required"`
	Logger   Logger   `yaml:"logger" validate:"required"`
	Database Database `yaml:"database" validate:"required"`
}
