package middleware

import (
	"fmt"
	"os"
	"slices"

	"github.com/MahmoodAhmed-SE/degree-progress-tracker/api/schemas"
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/database/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authenticate(next gin.HandlerFunc, allowed []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secret := []byte(os.Getenv("SECRET_KEY"))
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			token, ok := ctx.GetQuery("token")
			if !ok {
				schemas.ErrorResponse(ctx, 401, "Unauthorized")
			}
			authHeader = fmt.Sprintf("Bearer %v", token)
		}
		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v")
			}
			// if you want to add an audience
			// and issuer too
			return secret, nil
		})
		if err != nil || !token.Valid {
			schemas.ErrorResponse(ctx, 401, "Unauthorized")
			return
		}
		info := token.Claims.(jwt.MapClaims)
		currentUserRoles := make([]string, len(info["roles"].([]any)))
		for i, r := range info["roles"].([]any) {
			currentUserRoles[i] = r.(string)
		}
		userID := int(info["id"].(float64))
		user := models.User{
			ID:          userID,
			Username:    info["username"].(string),
			Phone:       info["phone"].(string),
			AuthGroupID: int64(info["auth_group_id"].(float64)),
			Roles:       currentUserRoles,
		}
		ctx.Set("user", user)
		if len(allowed) > 0 {
			for _, allowedRole := range allowed {
				if slices.Contains(currentUserRoles, allowedRole) {
					next(ctx)
					return
				}
			}
			schemas.ErrorResponse(ctx, 403, "Forbidden")
			return
		} else {
			next(ctx)
			return
		}
	}
}
