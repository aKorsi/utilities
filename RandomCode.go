package utilities

import (
	"encoding/base32"
	"encoding/base64"
	"math/rand"
	"strings"
)

var (
	enLowerChars = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
	}
	enCapitalChars = []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z",
	}
	enDigits = []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
	faChars = []string{
		"ا",
		"آ",
		"ب",
		"پ",
		"ت",
		"ث",
		"ج",
		"چ",
		"ح",
		"خ",
		"د",
		"ذ",
		"ر",
		"ز",
		"ژ",
		"س",
		"ش",
		"ص",
		"ض",
		"ط",
		"ظ",
		"ع",
		"غ",
		"ف",
		"ق",
		"ک",
		"گ",
		"ل",
		"م",
		"ن",
		"و",
		"ه",
		"ی",
	}
	faDigits = []string{
		"۰",
		"۱",
		"۲",
		"۳",
		"۴",
		"۵",
		"۶",
		"۷",
		"۸",
		"۹",
	}
)

func MakeBase2Code(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	for i := 0; i < length; i++ {
		code = code + enDigits[rand.Intn(2)]
	}
	return code, nil
}

func MakeBase8Code(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	for i := 0; i < length; i++ {
		code = code + enDigits[rand.Intn(8)]
	}
	return code, nil
}

func MakeBase10Code(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	for i := 0; i < length; i++ {
		code = code + enDigits[rand.Intn(10)]
	}
	return code, nil
}

func MakeBase16Code(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	for i := 0; i < length; i++ {
		code = code + append(enDigits, enCapitalChars[:6]...)[rand.Intn(16)]
	}
	return code, nil
}

func MakeBase32Code(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := make([]byte, length)
	_, err := rand.Read(code)
	if err != nil {
		return "", err
	}
	return strings.Replace(base32.StdEncoding.EncodeToString(code), "=", "", -1)[:length], nil
}

func MakeBase64Code(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := make([]byte, length)
	_, err := rand.Read(code)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(code)[:length], nil
}

func MakeEnAlphaNumericCode(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	src := append(enCapitalChars, enDigits...)
	for i := 0; i < length; i++ {
		code = code + src[rand.Intn(len(src))]
	}
	return code, nil
}

func MakeEnAlphaNumericCaseSensitiveCode(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	src := append(enLowerChars, append(enCapitalChars, enDigits...)...)
	for i := 0; i < length; i++ {
		code = code + src[rand.Intn(len(src))]
	}
	return code, nil
}

func MakeEnAlphabetCode(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	src := enCapitalChars
	for i := 0; i < length; i++ {
		code = code + src[rand.Intn(len(src))]
	}
	return code, nil
}

func MakeEnAlphabetCaseSensitiveCode(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	src := append(enLowerChars, enCapitalChars...)
	for i := 0; i < length; i++ {
		code = code + src[rand.Intn(len(src))]
	}
	return code, nil
}

func MakeFaAlphabetCode(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	src := faChars
	for i := 0; i < length; i++ {
		code = code + src[rand.Intn(len(src))]
	}
	return code, nil
}

func MakeFaAlphaNumericCode(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	src := append(faChars, faDigits...)
	for i := 0; i < length; i++ {
		code = code + src[rand.Intn(len(src))]
	}
	return code, nil
}

func MakeFaNumericCode(seed int64, length int) (string, error) {
	rand.Seed(seed)
	code := ""
	src := faDigits
	for i := 0; i < length; i++ {
		code = code + src[rand.Intn(len(src))]
	}
	return code, nil
}
