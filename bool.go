package intercom

// Bool is a helper method to create *bool.
// *bool is preferred to bool because it allows distinction between false and absence.
func Bool(value bool) *bool {
	return &value
}
