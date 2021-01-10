package utilities

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
)

func Sha512(str string) (string, error) {
	hash := sha512.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Sha384(str string) (string, error) {
	hash := sha512.New384()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Sha512224(str string) (string, error) {
	hash := sha512.New512_224()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Sha512256(str string) (string, error) {
	hash := sha512.New512_256()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Sha1(str string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Sha256(str string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Md5(str string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Fnv32(str string) (string, error) {
	hash := fnv.New32()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Fnv32a(str string) (string, error) {
	hash := fnv.New32a()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Fnv64(str string) (string, error) {
	hash := fnv.New64()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Fnv64a(str string) (string, error) {
	hash := fnv.New64a()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Fnv128(str string) (string, error) {
	hash := fnv.New128()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Fnv128a(str string) (string, error) {
	hash := fnv.New128a()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Adler32(str string) (string, error) {
	hash := adler32.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Crc32Ieee(str string) (string, error) {
	hash := crc32.NewIEEE()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Crc64Iso(str string) (string, error) {
	hash := crc64.New(crc64.MakeTable(crc64.ISO))
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Crc64Ecma(str string) (string, error) {
	hash := crc64.New(crc64.MakeTable(crc64.ECMA))
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func HmacSha1(str, key string) (string, error) {
	hash := hmac.New(sha1.New, []byte(key))
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func HmacSha256(str, key string) (string, error) {
	hash := hmac.New(sha256.New, []byte(key))
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func HmacSha512(str, key string) (string, error) {
	hash := hmac.New(sha512.New, []byte(key))
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func HmacMd5(str, key string) (string, error) {
	hash := hmac.New(md5.New, []byte(key))
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}
