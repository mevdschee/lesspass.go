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

// GetPasswordProfile returns the default password configuration.
func GetPasswordProfile(passwordProfile PasswordProfile) PasswordProfile {
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

// GeneratePassword generates v2 password.
func GeneratePassword(site, login, masterPassword string, passwordProfile PasswordProfile) string {
	passwordProfile = GetPasswordProfile(passwordProfile)
	entropy := CalcEntropy(site, login, masterPassword, passwordProfile)
	return RenderPassword(entropy, passwordProfile)
}

// CalcEntropy function
func CalcEntropy(site, login, masterPassword string, passwordProfile PasswordProfile) []byte {
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

// GetSetOfCharacters function
func GetSetOfCharacters(rules []string) string {
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

// ConsumeEntropy function
func ConsumeEntropy(generatedPassword string, quotient *big.Int, setOfCharacters string, maxLength int) (string, *big.Int) {
	if len(generatedPassword) >= maxLength {
		return generatedPassword, quotient
	}
	quotient, remainder := big.NewInt(0).DivMod(quotient, big.NewInt(int64(len(setOfCharacters))), big.NewInt(0))
	generatedPassword += string(setOfCharacters[int(remainder.Uint64())])
	return ConsumeEntropy(generatedPassword, quotient, setOfCharacters, maxLength)
}

// GetOneCharPerRule function
func GetOneCharPerRule(entropy *big.Int, rules []string) (string, *big.Int) {
	var oneCharPerRules = ""
	for _, rule := range rules {
		password, curEntropy := ConsumeEntropy("", entropy, characterSubsets[rule], 1)
		oneCharPerRules += password
		entropy = curEntropy
	}
	return oneCharPerRules, entropy
}

// InsertStringPseudoRandomly function
func InsertStringPseudoRandomly(generatedPassword string, entropy *big.Int, _string string) string {
	for i := 0; i < len(_string); i++ {
		quotient, remainder := big.NewInt(0).DivMod(entropy, big.NewInt(int64(len(generatedPassword))), big.NewInt(0))
		generatedPassword = generatedPassword[0:int(remainder.Uint64())] + string(_string[i]) + generatedPassword[int(remainder.Uint64()):]
		entropy = quotient
	}
	return generatedPassword
}

// GetConfiguredRules function
func GetConfiguredRules(passwordProfile PasswordProfile) []string {
	var rules []string
	allRules := []string{"lowercase", "uppercase", "numbers", "symbols"}
	for _, rule := range allRules {
		if passwordProfile[rule].(bool) {
			rules = append(rules, rule)
		}
	}
	return rules
}

// RenderPassword function
func RenderPassword(entropy []byte, passwordProfile PasswordProfile) string {
	var rules = GetConfiguredRules(passwordProfile)
	var setOfCharacters = GetSetOfCharacters(rules)
	var newInt, _ = big.NewInt(0).SetString(string(entropy), 16)
	password, passwordEntropy := ConsumeEntropy("", newInt, setOfCharacters, passwordProfile["length"].(int)-len(rules))
	charactersToAdd, characterEntropy := GetOneCharPerRule(passwordEntropy, rules)
	return InsertStringPseudoRandomly(password, characterEntropy, charactersToAdd)
}
