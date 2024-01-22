package mail

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSend(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../env/dev.env")
	if err != nil {
		return
	}

	// Test case: all parameters are valid
	t.Run("all parameters are valid", func(t *testing.T) {
		err := SendMail("thanhanphan17@gmail.com", "New Password",
			"../../templates/forgot_password_template.html", map[string]interface{}{
				"Password": "123456",
			})

		assert.NoError(t, err)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})

	// Test case: to parameter is empty
	t.Run("to parameter is empty", func(t *testing.T) {
		err := SendMail("", "Test Subject",
			"../../templates/forgot_password_template.html", map[string]interface{}{})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	// Test case: subject parameter is empty
	t.Run("subject parameter is empty", func(t *testing.T) {
		err := SendMail("test@example.com", "",
			"../../templates/forgot_password_template.html", map[string]interface{}{})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	// Test case: templatePath parameter is empty
	t.Run("templatePath parameter is empty", func(t *testing.T) {
		err := SendMail("test@example.com",
			"Test Subject", "", map[string]interface{}{})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	// Test case: all parameters are empty
	t.Run("all parameters are empty", func(t *testing.T) {
		err := SendMail("", "", "", map[string]interface{}{})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}
