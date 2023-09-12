package utils

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
