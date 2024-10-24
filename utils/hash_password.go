package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the password using bcrypt.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPassword compares the hashed password with the plain-text password.
func CheckPassword(hashedPassword string, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
