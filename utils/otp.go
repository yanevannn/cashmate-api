package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP(Length int) string {
	const digits = "0123456789"
	otp := ""

	for i := 0 ; i < Length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			otp += "0" // Fallback to '0' in case of error
		}
		otp += string(digits[num.Int64()])
	}

	return otp
}