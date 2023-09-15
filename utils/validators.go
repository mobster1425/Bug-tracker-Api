package utils

import (
	// "strings"
)

func ProjectNameError(name string) string {
	if name == "" || len(name) > 60 {
		return "Project name length must not be more than 60."
	}
	return ""
}

func ProjectMembersError(members []string) string {
	if len(members) == 0 {
		return "Members field must not be empty."
	}

	// Check for duplicate member IDs
	seen := make(map[string]bool)
	for _, member := range members {
		if seen[member] {
			return "Members field must not have already-added/duplicate IDs."
		}
		seen[member] = true

		// Check for valid UUID format (length 36)
		//	if len(member) != 36 {
		//		return "Members array must contain valid UUIDs."
		//	}
	}

	return ""
}

func ValidateBugFields(title, description, priority string) map[string]string {
	errors := make(map[string]string)
	validPriorities := []string{"low", "medium", "high"}

	if len(title) < 3 || len(title) > 60 {
		errors["title"] = "Title must be in the range of 3-60 characters length."
	}

	if description == "" {
		errors["description"] = "Description field must not be empty."
	}

	if !Contains(validPriorities, priority) {
		errors["priority"] = "Priority can only be - low, medium or high."
	}

	return errors
}

// contains checks if a string exists in a slice
func Contains(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func ContainsUint(slice []uint, value uint) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

/*
func ValidateSignup(username, password, email, confirmPassword string) (map[string]string, bool) {
	errors := make(map[string]string)
	valid := true

	if username == "" {
		errors["username"] = "Username is required."
		valid = false
	} else if len(username) < 3 || len(username) > 20 {
		errors["username"] = "Username must be between 3 and 20 characters."
		valid = false
	}

	if password == "" {
		errors["password"] = "Password is required."
		valid = false
	} else if len(password) < 6 {
		errors["password"] = "Password must be at least 6 characters long."
		valid = false
	}

	if confirmPassword == "" {
		errors["confirmPassword"] = "Confirm Password is required."
		valid = false
	} else if password != confirmPassword {
		errors["confirmPassword"] = "Passwords do not match."
		valid = false
	}

	if email == "" {
		errors["email"] = "Email is required."
		valid = false
	} else if !isValidEmail(email) {
		errors["email"] = "Invalid email address."
		valid = false
	}

	return errors, valid
}


*/

type AuthErrors struct {
	Username string
	Password string
	Email    string
}

func ValidateSignup(username string, password string, confirmpassword string, email string) (AuthErrors, bool) {
	errors := AuthErrors{}

	if len(username) == 0 || username == "" {
		errors.Username = "Username must not be null."
	}

	if len(password) == 0 || len(password) < 6 || password == "" {
		errors.Password = "Password must be at least 6 characters long."
	}

	if password != confirmpassword {
		errors.Password = "Passwords do not match."
	}
	

	if len(email) == 0 || email == "" {
		errors.Email = "Email must not be null"
	}

	valid := len(errors.Username) == 0 && len(errors.Password) == 0 && len(errors.Email) == 0

	return errors, valid
}




type AuthErrors1 struct {
	
	Password string
	Email    string
}


// ValidateLogin validates login data
func ValidateLogin(email string, password string) (AuthErrors1, bool) {
	

	errors := AuthErrors1{}


	if len(password) == 0 || len(password) < 6 || password == "" {
		errors.Password = "Password must be at least 6 characters long and password must not be empty."
	}

	if len(email) == 0 || email == "" {
		errors.Email = "Email must not be null"
	}

	valid :=  len(errors.Password) == 0 && len(errors.Email) == 0

	return errors, valid
}
