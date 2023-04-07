package utils

import "golang.org/x/crypto/bcrypt"

// handle errors globally
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// hash clear text password
func HashAndSalt(pwd []byte)string{
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	HandleError(err)
	return string(hashedPassword)
}
