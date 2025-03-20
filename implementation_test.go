package lab2 // Змінено з "main" на "lab2"

import (
	"errors" // Додано імпорт пакету errors
	"fmt"
	"testing"
)

func TestPrefixToInfix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"+ 5 * - 4 2 ^ 3 2", "(5 + ((4 - 2) * (3 ^ 2)))", nil},
		{"* + 1 2 3", "((1 + 2) * 3)", nil},
		{"^ 2 3", "(2 ^ 3)", nil},
		{"", "", errors.New("пустий рядок")}, // Використання errors.New
		{"+ 1", "", errors.New("недостатньо операндів для оператора")},
		{"+ 1 2 3", "", errors.New("неправильний вираз")},
		{"+ 1 a", "", errors.New("недопустимий символ у виразі")},
	}

	for _, test := range tests {
		result, err := PrefixToInfix(test.input)
		if result != test.expected || (err != nil && err.Error() != test.err.Error()) {
			t.Errorf("Для виразу %s очікувався результат %s, або помилка %v, але отримано %s, %v", test.input, test.expected, test.err, result, err)
		}
	}
}

func ExamplePrefixToInfix() {
	expression := "+ 5 * - 4 2 ^ 3 2"
	result, _ := PrefixToInfix(expression)
	fmt.Println(result)
	// Output: (5 + ((4 - 2) * (3 ^ 2)))
}
