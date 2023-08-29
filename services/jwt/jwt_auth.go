package jwt

import (
	"client_administration/constants"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type JWTData struct {
   jwt.StandardClaims
   CustomClaims map[string]string `json:"custom_claims"`
}



func GenerateJWTAccessToken(userId string, email string, password string, role string) (string, error) {
	// tokenLifeTime := 30 * 24 * 60 * 60 // 30 days in seconds
	secretKey := os.Getenv("SECRETKEY")

    // prepare claims for token
    claims := JWTData{
        StandardClaims: jwt.StandardClaims{
           // set token lifetime in timestamp
         //   ExpiresAt: time.Now().Add(time.Duration(tokenLifeTime)).Unix(),
        },
        // add custom claims like user_id or email, 
        CustomClaims: map[string]string{
            "user_id": userId,
            "email": email,
            "password": password,
            "role": role,
        },
     }
    
     // generate a string using claims and HS256 algorithm
     tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
     // sign the generated key using secretKey
     token, err := tokenString.SignedString([]byte(secretKey))

     return token, err
}

func VeryfiToken(authHeader string)(interface{}, interface{}, error){
   
      token := strings.Split(authHeader, " ") 
      if len(token) == 2 && strings.ToLower(token[0]) == "bearer" {
         token = token[1:]
      }
      
      claims := &JWTData{}
      secretKey := os.Getenv("SECRETKEY")
      stringtoken := strings.Join(token, "")


		tokenREs, err := jwt.ParseWithClaims(stringtoken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})      

      userData := constants.UserLoginLocalStorage {
         Email: claims.CustomClaims["email"],
         Id: claims.CustomClaims["user_id"],
         Password: claims.CustomClaims["password"],
         Role: claims.CustomClaims["role"],
      }

      return tokenREs, userData, err
}