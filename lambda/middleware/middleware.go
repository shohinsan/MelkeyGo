package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/aws/aws-lambda-go/events"
)

func ValidateJWTMiddleware(next func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// Extract the token from headers
		tokenString := extractTokenFromHeaders(request.Headers)
		if tokenString == "" {
			return events.APIGatewayProxyResponse{
				Body:       "Missing Auth Token",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}

		// Parse and validate the token
		claims, err := parseToken(tokenString)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       "User unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}

		// Check token expiration
		exp, ok := claims["exp"].(float64)
		if !ok {
			return events.APIGatewayProxyResponse{
				Body:       "Invalid expiration claim",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}
		if int64(exp) < time.Now().Unix() {
			return events.APIGatewayProxyResponse{
				Body:       "Token expired",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}

		// Call the next handler
		return next(request)
	}
}

func parseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := "secret"
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	if !token.Valid {
		return nil, errors.New("token is not valid - unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims of unauthorized ")
	}

	return claims, nil
}

func extractTokenFromHeaders(headers map[string]string) string {
	authHeader, ok := headers["Authorization"]
	if !ok {
		return ""
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}

	return splitToken[1]

}
