package lesspass

import (
	"math/big"
	"strings"
	"testing"
)

func TestGetDefaultSetOfCharacters(t *testing.T) {
	var setOfCharacters = getSetOfCharacters(nil)
	if setOfCharacters != "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 26*2+10+32 {
		t.Fatal("len(setOfCharacters) != 26 * 2 + 10 + 32")
	}
}

func TestGetDefaultSetOfCharactersConcatRulesInOrder(t *testing.T) {
	var setOfCharacters = getSetOfCharacters([]string{"lowercase", "uppercase", "numbers"})
	if setOfCharacters != "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 26*2+10 {
		t.Fatal("len(setOfCharacters) != 26 * 2 + 10")
	}
}

func TestgetSetOfCharactersOnlyLowercase(t *testing.T) {
	var setOfCharacters = getSetOfCharacters([]string{"lowercase"})
	if setOfCharacters != "abcdefghijklmnopqrstuvwxyz" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 26 {
		t.Fatal("len(setOfCharacters) != 26")
	}
}

func TestgetSetOfCharactersOnlyUppercase(t *testing.T) {
	var setOfCharacters = getSetOfCharacters([]string{"uppercase"})
	if setOfCharacters != "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 26 {
		t.Fatal("len(setOfCharacters) != 26")
	}
}

func TestgetSetOfCharactersOnlyNumbers(t *testing.T) {
	var setOfCharacters = getSetOfCharacters([]string{"numbers"})
	if setOfCharacters != "0123456789" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 10 {
		t.Fatal("len(setOfCharacters) != 10")
	}
}

func TestgetSetOfCharactersOnlySymbols(t *testing.T) {
	var setOfCharacters = getSetOfCharacters([]string{"symbols"})
	if setOfCharacters != "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 32 {
		t.Fatal("len(setOfCharacters) != 32")
	}
}

func TestGenerateOneCharPerRule(t *testing.T) {
	value, entropy := getOneCharPerRule(big.NewInt(26*26), []string{"lowercase", "uppercase"})
	if value != "aA" {
		t.Fatal("value != \"aA\"")
	}
	if len(value) != 2 {
		t.Fatal("len(value) != 2")
	}
	if entropy.Text(10) != "1" {
		t.Fatal("string(entropy) != \"1\"")
	}
}

func TestConfiguredRules(t *testing.T) {
	if strings.Join(getConfiguredRules(PasswordProfile{"uppercase": true}), ", ") != "uppercase" {
		t.Fatal("rules != \"uppercase\"")
	}
	if strings.Join(getConfiguredRules(PasswordProfile{"uppercase": true, "lowercase": true}), ", ") != "lowercase, uppercase" {
		t.Fatal("rules != \"lowercase, uppercase\"")
	}
	if strings.Join(getConfiguredRules(PasswordProfile{"lowercase": true, "symbols": false}), ", ") != "lowercase" {
		t.Fatal("rules != \"lowercase\"")
	}
	if strings.Join(getConfiguredRules(PasswordProfile{"lowercase": true, "uppercase": true, "symbols": true, "numbers": true}), ", ") != "lowercase, uppercase, numbers, symbols" {
		t.Fatal("rules != \"lowercase, uppercase, numbers, symbols\"")
	}
}
