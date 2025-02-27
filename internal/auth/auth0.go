// internal/auth/auth0.go
package auth

import (
	"errors"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

type Authenticator struct {
	Domain   string
	Audience string
}

func NewAuthenticator(domain, audience string) *Authenticator {
	return &Authenticator{
		Domain:   domain,
		Audience: audience,
	}
}

func (a *Authenticator) GetMiddleware() func(next http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verificar o iss claim
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(a.Domain, false)
			if !checkIss {
				return token, errors.New("invalid issuer")
			}

			// Verificar o audience claim
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(a.Audience, false)
			if !checkAud {
				return token, errors.New("invalid audience")
			}

			cert, err := a.getPemCert(token)
			if err != nil {
				return nil, err
			}

			result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			if err != nil {
				return nil, err
			}

			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extrair o token do cabeçalho Authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			// Verificar o formato do token
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Authorization header must be 'Bearer {token}'", http.StatusUnauthorized)
				return
			}

			// Continuar com o middleware jwt
			jwtMiddleware.Handler(next).ServeHTTP(w, r)
		})
	}
}

func (a *Authenticator) getPemCert(token *jwt.Token) (string, error) {
	// Implementar a lógica para obter o certificado do JWKS endpoint do Auth0
	// ...
	return "", nil
}
