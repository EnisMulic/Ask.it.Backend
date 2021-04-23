package services

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	"github.com/EnisMulic/Ask.it.Backend/utils"
	jwt "github.com/dgrijalva/jwt-go"
)


type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (as *AuthService) Register (req requests.RegisterRequest) (*responses.AuthResponse, error) {
	user, err := as.repo.GetByEmail(req.Email)

	if err == nil {
		return nil, err
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
		return nil, err
	}

	return generateAuthResponse(newUser)
}

func (as *AuthService) Login (req requests.LoginRequest) (*responses.AuthResponse, error) {
	user, err := as.repo.GetByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if !doPasswordsMatch(user.PasswordHash, req.Password, user.PasswordSalt) {
		return nil, fmt.Errorf(`{"response":"Wrong Password!"}`)
	}

	return generateAuthResponse(user)
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
	secretKey, err := utils.GetEnvVariable("SECRET_KEY")

	if err != nil {
		return "", err
	}

	sk := []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(sk)
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