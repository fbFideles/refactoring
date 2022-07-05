package cpf_test

import (
	"refactoring/cpf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPFValidator(t *testing.T) {
	t.Run("Should return false for a CPF with invalid length", func(t *testing.T) {
		assert.Equal(t, false, cpf.Validate("000000000000000000"))
	})

	t.Run("Should return true for a valid CPF", func(t *testing.T) {
		assert.Equal(t, true, cpf.Validate("824.208.440-84"))
	})
}
