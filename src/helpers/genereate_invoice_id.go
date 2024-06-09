package helpers

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func leftPadLen(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	retStr := strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func GenerateInvoiceID(companyID, year, month string, billingCounter int) string {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	randUniq := rand.Intn(y1.Intn(9999))
	generateInvoice := leftPadLen(companyID, "0", 4) + year + month + leftPadLen(strconv.Itoa(billingCounter), "0", 6) + leftPadLen(strconv.Itoa(randUniq), "0", 4)
	return generateInvoice
}
