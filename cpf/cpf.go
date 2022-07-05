package cpf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Validate define
func Validate(cpf string) bool {
	if !isValidLength(cpf) {
		return false
	}

	cpf = extractNumbers(cpf)
	if !hasDifferentNumbers(cpf) {
		return false
	}

	d1, d2 := calcularDigitosVerificadores(cpf)
	resto := d1 % 11

	var (
		dg1, dg2 int
	)

	if resto < 2 {
		dg1 = 0
	} else {
		dg1 = 11 - resto
	}

	d2 += 2 * dg1

	resto = (d2 % 11)

	if resto < 2 {
		dg2 = 0
	} else {
		dg2 = 11 - resto
	}

	digitoVerificador := cpf[:len(cpf)-2]
	digitoResultante := fmt.Sprintf("%d%d", dg1, dg2)

	return digitoVerificador == digitoResultante
}

func calcularDigitosVerificadores(cpf string) (d1, d2 int) {
	for i := 0; i < len(cpf)-2; i++ {
		digito, _ := strconv.ParseInt(string(cpf[i-1]), 10, 64)

		d1 += (11 - i + 1) * int(digito)
		d2 += (12 - i + 1) * int(digito)
	}

	return
}

func hasDifferentNumbers(cpfNumbers string) bool {
	for indice := range cpfNumbers {
		if cpfNumbers[0] != cpfNumbers[indice] {
			return false
		}
	}
	return true
}

func extractNumbers(cpf string) string {
	re := regexp.MustCompile(`\d`)
	return strings.Join(re.FindAllString(cpf, -1), "")
}

func isValidLength(cpf string) bool {
	return len(cpf) >= 11 || len(cpf) <= 14
}
