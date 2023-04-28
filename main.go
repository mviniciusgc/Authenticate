package main

import (
	"github.com/mviniciusgc/authenticate/src/server"
	"github.com/spf13/viper"
)

type mainService struct {
	Server server.Services
}

func main() {
	viper.SetDefault("KEYCLOAK_USER", "admin")
	viper.SetDefault("KEYCLOAK_PASS", "admin")
	viper.SetDefault("KEYCLOAK_DOMAIN", "http://localhost:8080")
	viper.SetDefault("KEYCLOAK_REALM", "master")
	viper.SetDefault("KEYCLOAK_MAIN_REALM", "master")
	viper.SetDefault("KEYCLOAK_CLIENT_SECRET", "ZDJCb2VgbX5U97uw3Z2G3vfvYVgV0aFs")
	viper.SetDefault("KEYCLOAK_CLIENT_ID", "authentication")
	//viper.SetDefault("PUBLIC_KEY", "MIICmzCCAYMCBgGHhwcBlDANBgkqhkiG9w0BAQsFADARMQ8wDQYDVQQDDAZtYXN0ZXIwHhcNMjMwNDE1MjIyNDU2WhcNMzMwNDE1MjIyNjM2WjARMQ8wDQYDVQQDDAZtYXN0ZXIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCdVYsX1z2TK2uaArTBLOzYRGAZEN5d+hAEYqUd+4qr1/3Fr+iclBQt2iy2W/Rczjy+KWqCS4YEkRbJAP7u7QV5JJLmcpNxmddL+Cl2vrwzDfK+viim+JmwUNXChYzZEvbQcjjNiq5HgFuX9/7AqmaEhjHMBuVhTCrVM1zclcFLLOSFLH8JOXA9OivUslgAJOkgdCPmza6cBkrGHyv3s6haQjYB8bUlCeQeL/JMMsH7mlajvcVRPasXk33VUKvV9nUESF2cVnLioj38IMv0W87VisYyVI2P9oTDpymo09IKqMVOSLmzJnYGAxVaFl0LazHcwqTSkBXBkxCYyYNygmIBAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAGA3YXibwTod/wJx/TFG4jYZQfdeA0MxmgSH8vzXWzZJk5jOJrLBB1mIyzPBTGisa3ervrSK9I8A8dkZojM8akKpOHBD8IWat6bVs12EZzJ/+C4q0N77wH8bRyVcVYhvGE9U9M5ji6rUKBw8pj/hIVn2enM5nvNUxMna2lMofkOD7P/NYj7UeVDqBYqpbouDGE8jLqLrDN6ccTgQXaNFV4gM9fxGux62lEMvZDvXaVdbvlBooo9aMbNG+EffwFJggSFi2NU9fzoo+8UCgjrb24MMAYRyvJ0H4gd9fn+C1doF/ZJBQBD2UKfngwwpdepOD2iB47QJZinlRF1xZpoNqA8=")
	viper.SetDefault("PUBLIC_KEY", "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnVWLF9c9kytrmgK0wSzs2ERgGRDeXfoQBGKlHfuKq9f9xa/onJQULdostlv0XM48vilqgkuGBJEWyQD+7u0FeSSS5nKTcZnXS/gpdr68Mw3yvr4opviZsFDVwoWM2RL20HI4zYquR4Bbl/f+wKpmhIYxzAblYUwq1TNc3JXBSyzkhSx/CTlwPTor1LJYACTpIHQj5s2unAZKxh8r97OoWkI2AfG1JQnkHi/yTDLB+5pWo73FUT2rF5N91VCr1fZ1BEhdnFZy4qI9/CDL9FvO1YrGMlSNj/aEw6cpqNPSCqjFTki5syZ2BgMVWhZdC2sx3MKk0pAVwZMQmMmDcoJiAQIDAQAB")
	var s mainService
	s.Server.InitializeServer()
}
