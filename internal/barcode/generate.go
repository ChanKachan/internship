package barcode

import (
	"fmt"
	"math/rand"
	"strconv"
)

type code string

const (
	RUSSIA code = "460"
)

func Generate() string {
	prefix := RUSSIA
	randomPart := fmt.Sprintf("%09d", rand.Intn(1000000000))
	code12 := string(prefix) + randomPart
	checkDigit := calculateEAN13CheckDigit(code12)

	return code12 + strconv.Itoa(checkDigit)
}

func calculateEAN13CheckDigit(code12 string) int {
	sum := 0

	for i, c := range code12 {
		digit := int(c - '0')
		if i%2 == 0 {
			sum += digit * 1
		} else {
			sum += digit * 3
		}
	}

	return (10 - (sum % 10)) % 10
}
