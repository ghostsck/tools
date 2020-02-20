package authcode

func Decrypt(str string, key string) (string, bool) {
	return authcode(str, key, false, 0)
}

func Encrypt(str string, key string, expiry int64) (string, bool) {
	return authcode(str, key, true, expiry)
}
