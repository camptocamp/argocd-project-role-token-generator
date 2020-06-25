package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	project  string
	role     string
	lifetime time.Duration
)

func init() {
	flag.StringVar(&project, "project", "", "Argo CD project which the role belongs to")
	flag.StringVar(&role, "role", "", "Argo CD project role which to create a token for")
	flag.DurationVar(&lifetime, "lifetime", 0, "Lifetime of the token")
	flag.Parse()
}

func main() {
	if project == "" {
		fmt.Println("project is missing")
		flag.Usage()
		os.Exit(2)
	}

	if role == "" {
		fmt.Println("role is missing")
		flag.Usage()
		os.Exit(2)
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Argo CD secret key?")

	if !scanner.Scan() {
		fmt.Println("Failed to read secret:", scanner.Err())
		os.Exit(1)
	}

	secret := scanner.Bytes()

	id, err := uuid.NewRandom()

	if err != nil {
		fmt.Println("Failed to generate UUID:", err)
		os.Exit(1)
	}

	subject := fmt.Sprintf("proj:%s:%s", project, role)
	now := time.Now().UTC()

	claims := jwt.StandardClaims{
		Id:        id.String(),
		Subject:   subject,
		Issuer:    "argocd",
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
	}

	if lifetime > 0 {
		expires := now.Add(lifetime)
		claims.ExpiresAt = expires.Unix()
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)

	if err != nil {
		fmt.Println("Failed to generate token:", err)
		os.Exit(1)
	}

	fmt.Println("claims:")

	fmt.Println("  id:", claims.Id)
	fmt.Println("  iat:", claims.IssuedAt)

	if lifetime > 0 {
		fmt.Println("  exp:", claims.ExpiresAt)
	}

	fmt.Println("token:", token)
}
