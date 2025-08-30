package halpers

import "golang.org/x/crypto/bcrypt"

func HasPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	IfPanicError(err)

	return string(bytes)
}
