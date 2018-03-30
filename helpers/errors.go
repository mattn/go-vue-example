package helpers

// Errors keep a record of the validation errors
type Errors struct {
	Messages map[string]string
}

// NewErrors creates a new instance of Errors
func NewErrors() *Errors {
	errors := &Errors{
		Messages: map[string]string{},
	}

	return errors
}

// Add adds a new error given a field and a message
func (errors *Errors) Add(field string, message string) {
	errors.Messages[field] = message
}

// ValidateMinValue validates that the given value is less than minValueAllowed.
func (errors *Errors) ValidateMinValue(value int, minValueAllowed int, field string, message string) {
	if value < minValueAllowed {
		errors.Add(field, message)
	}
}

// ValidateMaxValue validates that the given value is greater than maxValueAllowed.
func (errors *Errors) ValidateMaxValue(value int, maxValueAllowed int, field string, message string) {
	if value > maxValueAllowed {
		errors.Add(field, message)
	}
}

// Clear removes all tracked errors.
func (errors *Errors) Clear() {
	errors.Messages = map[string]string{}
}

// HasMessages returns true if the estructure is not empty.
func (errors *Errors) HasMessages() bool {
	return len(errors.Messages) > 0
}

// vi:syntax=go
