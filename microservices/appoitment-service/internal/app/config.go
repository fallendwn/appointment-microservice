package app

import "os"

type Config struct {
	Port              string
	MongoURI          string
	DoctorServiceAddr string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("PORT"),
		MongoURI:          os.Getenv("MONGO_URI"),
		DoctorServiceAddr: os.Getenv("DOCTOR_SERVICE_URL"),
	}
}
