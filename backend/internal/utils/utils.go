package utils

import (
	db "Backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckPasswordHash(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

type Claims struct {
	UserId int64
	Roles  []string
	jwt.RegisteredClaims
}

type Email struct {
	To      []string
	Cc      []string
	Subject string
	Body    string
}

func CreateToken(user db.GetUserByEmailRow) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		Roles:  strings.Split(string(user.Roles), ","),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "Sharobim_deacons",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))
}

func SendEmail(payload *Email) error {
	from := os.Getenv("EMAIL")
	pass := os.Getenv("EMAIL_PASSWORD")

	msg := "From: " + from + "\r\n" +
		"To: " + strings.Join(payload.To, ", ") + "\r\n" +
		"Cc: " + strings.Join(payload.Cc, ", ") + "\r\n" +
		"Subject: " + payload.Subject + "\r\n" +
		payload.Body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, payload.To, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}
