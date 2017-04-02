package lesspass

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"math/big"
	"strconv"

	"golang.org/x/crypto/pbkdf2"
)

var characterSubsets = map[string]string{
	"lowercase": "abcdefghijklmnopqrstuvwxyz",
	"uppercase": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"numbers":   "0123456789",
	"symbols":   "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~",
}

// PasswordProfile holds the lesspass settings
type PasswordProfile map[string]interface{}

func getPasswordProfile(passwordProfile PasswordProfile) PasswordProfile {
	var defaultPasswordProfile = PasswordProfile{
		"lowercase":  true,
		"uppercase":  true,
		"numbers":    true,
		"symbols":    true,
		"digest":     "sha256",
		"iterations": 100000,
		"keylen":     32,
		"length":     16,
		"counter":    1,
		"version":    2,
	}
	for key, value := range passwordProfile {
		defaultPasswordProfile[key] = value
	}
	return defaultPasswordProfile
}

// GeneratePassword generates a v2 lesspass password.
func GeneratePassword(site, login, masterPassword string, passwordProfile PasswordProfile) string {
	passwordProfile = getPasswordProfile(passwordProfile)
	entropy := calcEntropy(site, login, masterPassword, passwordProfile)
	return renderPassword(entropy, passwordProfile)
}

func calcEntropy(site, login, masterPassword string, passwordProfile PasswordProfile) []byte {
	var salt = site + login + strconv.FormatInt(int64(passwordProfile["counter"].(int)), 16)
	var digest func() hash.Hash
	switch passwordProfile["digest"] {
	case "sha256":
		digest = sha256.New
	case "sha512":
		digest = sha512.New
	}
	return []byte(hex.EncodeToString(pbkdf2.Key([]byte(masterPassword), []byte(salt), passwordProfile["iterations"].(int), passwordProfile["keylen"].(int), digest)))
}

func getSetOfCharacters(rules []string) string {
	var setOfChars = ""
	if rules == nil {
		return characterSubsets["lowercase"] + characterSubsets["uppercase"] + characterSubsets["numbers"] + characterSubsets["symbols"]
	}
	for _, val := range rules {
		v, ok := characterSubsets[val]
		if ok {
			setOfChars += v
		}
	}
	return setOfChars
}

func consumeEntropy(generatedPassword string, quotient *big.Int, setOfCharacters string, maxLength int) (string, *big.Int) {
	if len(generatedPassword) >= maxLength {
		return generatedPassword, quotient
	}
	quotient, remainder := big.NewInt(0).DivMod(quotient, big.NewInt(int64(len(setOfCharacters))), big.NewInt(0))
	generatedPassword += string(setOfCharacters[int(remainder.Uint64())])
	return consumeEntropy(generatedPassword, quotient, setOfCharacters, maxLength)
}

func getOneCharPerRule(entropy *big.Int, rules []string) (string, *big.Int) {
	var oneCharPerRules = ""
	for _, rule := range rules {
		password, curEntropy := consumeEntropy("", entropy, characterSubsets[rule], 1)
		oneCharPerRules += password
		entropy = curEntropy
	}
	return oneCharPerRules, entropy
}

func insertStringPseudoRandomly(generatedPassword string, entropy *big.Int, _string string) string {
	for i := 0; i < len(_string); i++ {
		quotient, remainder := big.NewInt(0).DivMod(entropy, big.NewInt(int64(len(generatedPassword))), big.NewInt(0))
		generatedPassword = generatedPassword[0:int(remainder.Uint64())] + string(_string[i]) + generatedPassword[int(remainder.Uint64()):]
		entropy = quotient
	}
	return generatedPassword
}

func getConfiguredRules(passwordProfile PasswordProfile) []string {
	var rules []string
	allRules := []string{"lowercase", "uppercase", "numbers", "symbols"}
	for _, rule := range allRules {
		if passwordProfile[rule] != nil && passwordProfile[rule].(bool) {
			rules = append(rules, rule)
		}
	}
	return rules
}

func renderPassword(entropy []byte, passwordProfile PasswordProfile) string {
	var rules = getConfiguredRules(passwordProfile)
	var setOfCharacters = getSetOfCharacters(rules)
	var newInt, _ = big.NewInt(0).SetString(string(entropy), 16)
	password, passwordEntropy := consumeEntropy("", newInt, setOfCharacters, passwordProfile["length"].(int)-len(rules))
	charactersToAdd, characterEntropy := getOneCharPerRule(passwordEntropy, rules)
	return insertStringPseudoRandomly(password, characterEntropy, charactersToAdd)
}
