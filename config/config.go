package config

import (
	"log"
	"net"
	"os"
)

func MustGetenv(name string) string {
	val := os.Getenv(name)
	if val == "" {
		log.Fatalf("Environment variable %s is not set!", name)
	}
	return val
}

func GetenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

var DatabaseUrl = MustGetenv("DATABASE_URL")
var Port = GetenvWithDefault("PORT", "3000")
var BindHost = GetenvWithDefault("BIND_HOST", "::")
var ListenOn = net.JoinHostPort(BindHost, Port)
