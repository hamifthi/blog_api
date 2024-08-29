package service

import (
	"errors"
	"regexp"
)

func validateBlog(title, content string, authorID uint) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	if content == "" {
		return errors.New("content cannot be empty")
	}
	if authorID == 0 {
		return errors.New("invalid author ID")
	}
	return nil
}

func validateName(name string) bool {
	if len(name) < 3 || len(name) > 50 {
		return false
	}
	return true
}

func validateUser(name, email string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile the regex pattern.
	re := regexp.MustCompile(regex)
	// Return whether the email matches the pattern.
	if email != "" && re.MatchString(email) == false {
		return errors.New("invalid email")
	}
	if validateName(name) == false {
		return errors.New("invalid name")
	}
	return nil
}
