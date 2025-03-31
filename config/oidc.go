package config

type OIDCConfig struct {
	Issuer   string
	Audience string
}

var OidcConfigDev = OIDCConfig{
	Issuer:   "https://localhost:8080/auth/realms/starter",
	Audience: "dev",
}

var OidcConfigProd = OIDCConfig{
	Issuer:   "https://localhost:8080/auth/realms/starter",
	Audience: "spa",
}
