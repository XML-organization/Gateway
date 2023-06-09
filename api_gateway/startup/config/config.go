package config

import "os"

type Config struct {
	Port                string
	UserHost            string
	UserPort            string
	AutentificationHost string
	AutentificationPort string
	BookingHost         string
	BookingPort         string
	AccomodationHost    string
	AccomodationPort    string
	NotificationHost    string
	NotificationPort    string
}

func NewConfig() *Config {
	return &Config{
		Port:                os.Getenv("GATEWAY_PORT"),
		UserHost:            os.Getenv("USER_SERVICE_HOST"),
		UserPort:            os.Getenv("USER_SERVICE_PORT"),
		AutentificationHost: os.Getenv("AUTENTIFICATION_SERVICE_HOST"),
		AutentificationPort: os.Getenv("AUTENTIFICATION_SERVICE_PORT"),
		BookingHost:         os.Getenv("BOOKING_SERVICE_HOST"),
		BookingPort:         os.Getenv("BOOKING_SERVICE_PORT"),
		AccomodationHost:    os.Getenv("ACCOMODATION_SERVICE_HOST"),
		AccomodationPort:    os.Getenv("ACCOMODATION_SERVICE_PORT"),
		NotificationHost:    os.Getenv("NOTIFICATION_SERVICE_HOST"),
		NotificationPort:    os.Getenv("NOTIFICATION_SERVICE_PORT"),
	}
}
