// file resposnsible for hashing the passwords
package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // converting string to the byte slice
	return string(bytes), err
}

// This function is typically used for authentication, where a user provides a plaintext password, and the system checks it against a stored hashed password (e.g., during login).
// bcrypt is a secure hashing algorithm that incorporates a salt and is designed to be computationally expensive to resist brute-force attacks.
// Example Workflow:
// A user enters a password, say "myPassword123".
// The system retrieves the stored hashed password (e.g., a bcrypt hash like $2a$10$...).
// CheckHashedPassword("myPassword123", storedHash) is called.
// If the plaintext password "myPassword123" matches the hash, the function returns true; otherwise, it returns false.

func CheckHashedPassword(password, hashedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil
}
