package services

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
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

func (as *AuthService) Register (req requests.RegisterRequest) (*responses.AuthResponse, error) {
	user, _ := as.repo.GetByEmail(req.Email)

	if user.Email == req.Email {
		return nil, constants.ErrEmailIsTaken
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
		return nil, constants.ErrGeneric
	}

	res, _ := generateAuthResponse(newUser)

	return res, nil
}

func (as *AuthService) Login (req requests.LoginRequest) (*responses.AuthResponse, error) {
	user, _ := as.repo.GetByEmail(req.Email)

	if user.ID == 0 {	
		return nil, constants.ErrEmailIsTaken
	}

	if !doPasswordsMatch(user.PasswordHash, req.Password, user.PasswordSalt) {
		// resErr := responses.ErrorResponseModel{
		// 	FieldName: "",
		// 	Message: "Wrong Password!",
		// }

		// errors := responses.NewErrorResponse(resErr)

		return nil, constants.ErrWrongPassword
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
	tokenLifetime, err := strconv.Atoi(os.Getenv("JWT_TOKEN_LIFETIME_IN_MINUTES"))

	if err != nil {
		return "", err
	}

	claims := jwt.StandardClaims{
		Issuer: os.Getenv("JWT_ISSUER"),
		Audience: os.Getenv("JWT_AUDIENCE"),
		Subject: fmt.Sprint(user.ID),
		IssuedAt: time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(tokenLifetime)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func generateAuthResponse(user domain.User) (*responses.AuthResponse, error) {
	token, err := generateJWT(user)

	if err != nil {
		return nil, err
	}

	data := responses.AuthResponseModel{
		Token: token,
		RefreshToken: "",
	}

	return &responses.AuthResponse{Data: data}, nil
}