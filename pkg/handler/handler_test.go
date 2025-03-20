package handler

import (
	"errors"
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr error
	}{
		{
			name:        "Valid expression",
			input:       "+ 5 * - 4 2 ^ 3 2",
			expected:    "(5 + ((4 - 2) * (3 ^ 2)))",
			expectedErr: nil,
		},
		{
			name:        "Invalid expression",
			input:       "+ 1",
			expected:    "",
			expectedErr: errors.New("недостатньо операндів для оператора"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := strings.NewReader(test.input)
			output := &strings.Builder{}

			handler := ComputeHandler{
				Input:  input,
				Output: output,
			}

			err := handler.Compute()
			if (err != nil && test.expectedErr == nil) || (err == nil && test.expectedErr != nil) {
				t.Errorf("Очікувана помилка: %v, отримана помилка: %v", test.expectedErr, err)
			}

			if err == nil && strings.TrimSpace(output.String()) != strings.TrimSpace(test.expected) {
				t.Errorf("Очікуваний результат: %s, отриманий результат: %s", test.expected, output.String())
			}
		})
	}
}
