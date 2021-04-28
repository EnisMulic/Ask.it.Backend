package services

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	jwt "github.com/dgrijalva/jwt-go"
)


type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (as *AuthService) Register (req requests.RegisterRequest) (*responses.AuthResponse, *responses.ErrorResponse) {
	user, err := as.repo.GetByEmail(req.Email)

	if user.ID == 0 {
		err := responses.ErrorResponseModel{
			FieldName: "email",
			Message: constants.EmailIsTakenError,
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	var salt = generateRandomSalt(saltSize)
	var hash = hashPassword(req.Password, salt)

	user = domain.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		PasswordSalt: salt,
		PasswordHash: hash,
	}
	
	newUser, err := as.repo.Create(user)
	if err != nil {
		resErr := responses.ErrorResponseModel{
			FieldName: "",
			Message: "An error occurred",
		}

		errors := responses.NewErrorResponse(resErr)	


		return nil, errors
	}

	res, _ := generateAuthResponse(newUser)

	return res, nil
}

func (as *AuthService) Login (req requests.LoginRequest) (*responses.AuthResponse, *responses.ErrorResponse) {
	user, _ := as.repo.GetByEmail(req.Email)

	if user.ID == 0 {
		err := responses.ErrorResponseModel{
			FieldName: "email",
			Message: constants.EmailIsTakenError,
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	if !doPasswordsMatch(user.PasswordHash, req.Password, user.PasswordSalt) {
		resErr := responses.ErrorResponseModel{
			FieldName: "",
			Message: "Wrong Password!",
		}

		errors := responses.NewErrorResponse(resErr)

		return nil, errors
	}

	res, _ := generateAuthResponse(user)

	return res, nil
}

const saltSize = 16


func generateRandomSalt(saltSize int) string {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(salt)
}

func hashPassword(password string, salt string) string {
	var passwordBytes = []byte(password)
	var saltBytes = []byte(salt)

	var sha512Hasher = sha512.New()

	passwordBytes = append(passwordBytes, saltBytes...)

	sha512Hasher.Write(passwordBytes)

	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

func doPasswordsMatch(passwordHash string, currPassword string, salt string) bool {
	var currPasswordHash = hashPassword(currPassword, salt)

	return passwordHash == currPasswordHash
}

func generateJWT(user domain.User) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	tokenLifetime, err := strconv.Atoi(os.Getenv("JWT_TOKEN_LIFETIME_IN_MINUTES"))

	if err != nil {
		return "", err
	}

	claims["iss"] = os.Getenv("JWT_ISSUER")
	claims["aud"] = os.Getenv("JWT_AUDIENCE")
	claims["sub"] = fmt.Sprint(user.ID)
	claims["iat"] = fmt.Sprint(time.Now().Unix())
	claims["nbf"] = fmt.Sprint(time.Now().Unix())
	claims["exp"] = fmt.Sprint(time.Now().Add(time.Second * time.Duration(tokenLifetime)).Unix())
	
	claims["email"] = user.Email

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}

	return tokenString, nil
}

func generateAuthResponse(user domain.User) (*responses.AuthResponse, error) {
	token, err := generateJWT(user)

	if err != nil {
		return nil, err
	}

	return &responses.AuthResponse{
		Data: struct {
			Token string
			RefreshToken string
		} {
			Token: token,
			RefreshToken: "",
		},
	}, nil
}