package validators

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateReceiptDate(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	return err == nil
}

func ValidateReceiptTime(fl validator.FieldLevel) bool {
	_, err := time.Parse("15:04", fl.Field().String())
	return err == nil
}

// ValidateAlphanumeric validates alphanumeric strings (including whitespace, hyphens, and ampersands).
func ValidateRetailerName(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString("^[\\w\\s\\-&]+$", fl.Field().String())
	return match
}

// ValidateAlphanumeric validates alphanumeric strings (including whitespace, hyphens, and ampersands).
func ValidateShortDescription(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString("^[\\w\\s\\-]+$", fl.Field().String())
	return match
}

// ValidateDecimal validates numeric strings with two decimal places.
func ValidateDecimal(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString("^\\d+\\.\\d{2}$", fl.Field().String())
	return match
}
