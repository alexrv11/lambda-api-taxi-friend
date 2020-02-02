package utils

import "golang.org/x/crypto/bcrypt"

func CreatePassword(password string) string {
	pass := []byte(password)

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, _ := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)


	return string(hash)
}

func CompareHashAndPassword(hashedPassword, password []byte) bool {

	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return false
	}

	return true

}
