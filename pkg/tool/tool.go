package tool

import (
	cryptoRand "crypto/rand"
	"math/big"
)

func GenerateRandomString(length int) string {
	koreanBase := []rune("가나다라마바사아자차카타파하")
	numbers := []rune("0123456789")
	allRunes := append(koreanBase, numbers...)

	result := make([]rune, length)

	for i := range result {
		n, err := cryptoRand.Int(
			cryptoRand.Reader,
			big.NewInt(int64(len(allRunes))),
		)
		if err != nil {
			panic(err)
		}
		result[i] = allRunes[n.Int64()]
	}

	return string(result)
}
