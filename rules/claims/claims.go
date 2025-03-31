package claims

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"code.icod.de/dalu/oidc/options"
)

type Claims struct {
	Exp            time.Time `json:"exp,omitempty"`
	Iat            time.Time `json:"iat,omitempty"`
	Jti            string    `json:"jti,omitempty"`
	Iss            string    `json:"iss,omitempty"`
	Aud            []string  `json:"aud,omitempty"`
	Sub            string    `json:"sub,omitempty"`
	Typ            string    `json:"typ,omitempty"`
	Azp            string    `json:"azp,omitempty"`
	SessionState   string    `json:"session_state,omitempty"`
	AllowedOrigins []string  `json:"allowed-origins,omitempty"`
	RealmAccess    struct {
		Roles []string `json:"roles,omitempty"`
	} `json:"realm_access,omitempty"`
	ResourceAccess struct {
		Account struct {
			Roles []string `json:"roles,omitempty"`
		} `json:"account,omitempty"`
	} `json:"resource_access,omitempty"`
	Scope             string `json:"scope,omitempty"`
	Sid               string `json:"sid,omitempty"`
	EmailVerified     bool   `json:"email_verified,omitempty"`
	Name              string `json:"name,omitempty"`
	PreferredUsername string `json:"preferred_username,omitempty"`
	GivenName         string `json:"given_name,omitempty"`
	FamilyName        string `json:"family_name,omitempty"`
	Email             string `json:"email,omitempty"`
}

func FromContext(ctx context.Context) *Claims {
	v, found := ctx.Value(options.DefaultClaimsContextKeyName).(map[string]interface{})
	if found {
		claims := new(Claims)
		j, e := json.Marshal(v)
		if e != nil {
			log.Println(e.Error())
			return nil
		}
		if e := json.Unmarshal(j, claims); e != nil {
			log.Println(e.Error())
			return nil
		}
		return claims
	} else {
		return nil
	}
}

func SubFromContext(ctx context.Context) string {
	v, x := ctx.Value(options.DefaultClaimsContextKeyName).(map[string]interface{})
	if x {
		return v["sub"].(string)
	}
	return ""
}
