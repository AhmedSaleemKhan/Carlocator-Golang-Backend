package lib

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func NewCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			var (
				hasNumber      = false
				hasSpecialChar = false
				hasLetter      = false
				hasSuitableLen = false
			)

			password := fl.Field().String()

			if utf8.RuneCountInString(password) <= 30 && utf8.RuneCountInString(password) >= 8 {
				hasSuitableLen = true
			}

			for _, c := range password {
				switch {
				case unicode.IsNumber(c):
					hasNumber = true
				case unicode.IsPunct(c) || unicode.IsSymbol(c):
					hasSpecialChar = true
				case unicode.IsLetter(c) || c == ' ':
					hasLetter = true
				default:
					return false
				}
			}
			return hasNumber && hasSpecialChar && hasLetter && hasSuitableLen
		})
	}
}

func StaffRoleValidation(role string) error {
	availableRoles := []string{Accounting, ITAdministrator, InventoryManager, LotPorter, Receptionist, SalesRepresentative, SalesManager, ServiceTechnician}
	for _, availableRole := range availableRoles {
		if availableRole == role {
			return nil
		}
	}
	return fmt.Errorf("role doesn't exist")
}
