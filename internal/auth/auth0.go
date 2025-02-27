// Package auth provides authentication and authorization functionality.
package auth

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

// Authenticator handles JWT authentication using Auth0.
type Authenticator struct {
	Domain     string
	Audience   string
	middleware *jwtmiddleware.JWTMiddleware
}

// NewAuthenticator creates a new Authenticator instance configured with Auth0 settings.
func NewAuthenticator(domain, audience string) *Authenticator {
	return &Authenticator{
		Domain:   domain,
		Audience: audience,
	}
}

// GetMiddleware returns the JWT middleware handler for protecting routes.
func (a *Authenticator) GetMiddleware() func(http.Handler) http.Handler {
	return a.middleware.Handler
}

func (a *Authenticator) getPemCert(token *jwt.Token) (string, error) {
	// Implementar a lógica para obter o certificado do JWKS endpoint do Auth0
	// ...
	return "", nil
}
