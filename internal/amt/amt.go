package amt

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const offset = 48

type Amount struct {
	intPart  int
	fracPart []byte
}

func New(amtStr string) (Amount, error) {
	firstIdx := strings.Index(amtStr, ".")
	if firstIdx == -1 {
		return new(amtStr, "0")
	}
	lastIdx := strings.LastIndex(amtStr, ".")
	if firstIdx != lastIdx {
		return Amount{}, errors.New("")
	}
	if firstIdx == 0 {
		return new("0", amtStr[firstIdx+1:])
	}
	return new(amtStr[:firstIdx], amtStr[lastIdx+1:])
}

func (a Amount) Plus(oa Amount) Amount {
	aFracPart := make([]byte, len(a.fracPart))
	oaFracPart := make([]byte, len(oa.fracPart))
	copy(aFracPart, a.fracPart)
	copy(oaFracPart, oa.fracPart)
	if len(oaFracPart) > len(aFracPart) {
		aFracPart, oaFracPart = oaFracPart, aFracPart
	}

	diff := len(aFracPart) - len(oaFracPart)
	carry := 0
	for idx := diff; idx < len(aFracPart); idx++ {
		sum := int(aFracPart[idx]+oaFracPart[idx-diff]) - 2*offset + carry
		aFracPart[idx] = byte(sum%10 + offset)
		carry = sum / 10
	}

	return Amount{
		intPart:  a.intPart + oa.intPart + carry,
		fracPart: bytes.TrimLeft(aFracPart, "0"),
	}
}

func (a Amount) IsZero() bool {
	return a.intPart == 0 && len(a.fracPart) == 0
}

func (a Amount) IsNegative() bool {
	return a.intPart < 0
}

func (a Amount) String() string {
	n := len(a.fracPart)
	if n == 0 {
		return fmt.Sprintf("%d", a.intPart)
	}
	fracPart := make([]byte, n)
	for idx := 0; idx < n; idx++ {
		fracPart[idx] = a.fracPart[n-1-idx]
	}
	return fmt.Sprintf("%d.%s", a.intPart, fracPart)
}

func (a Amount) ToFloat() (float64, error) {
	af, err := strconv.ParseFloat(a.String(), 64)
	if err != nil {
		return 0, err
	}
	return af, nil
}

func new(intPartStr, fracPartStr string) (Amount, error) {
	intPart, err := strconv.Atoi(intPartStr)
	if err != nil {
		return Amount{}, errors.New("")
	}
	fracPartInt, err := strconv.Atoi(fracPartStr)
	if err != nil {
		return Amount{}, errors.New("")
	}
	if fracPartInt < 0 {
		return Amount{}, errors.New("")
	}
	n := len(fracPartStr)
	fracPart := make([]byte, n)
	for idx := 0; idx < n; idx++ {
		fracPart[idx] = fracPartStr[n-1-idx]
	}
	return Amount{
		intPart:  intPart,
		fracPart: bytes.TrimLeft(fracPart, "0"),
	}, nil
}
