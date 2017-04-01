package tests

import (
	"math/big"
	"strings"
	"testing"

	lesspass ".."
)

func testGetDefaultSetOfCharacters(t *testing.T) {
	var setOfCharacters = lesspass.GetSetOfCharacters(nil)
	if setOfCharacters != "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 26*2+10+32 {
		t.Fatal("len(setOfCharacters) != 26 * 2 + 10 + 32")
	}
}

func testGetDefaultSetOfCharactersConcatRulesInOrder(t *testing.T) {
	var setOfCharacters = lesspass.GetSetOfCharacters([]string{"lowercase", "uppercase", "numbers"})
	if setOfCharacters != "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 26*2+10 {
		t.Fatal("len(setOfCharacters) != 26 * 2 + 10")
	}
}

func testGetSetOfCharactersOnlyLowercase(t *testing.T) {
	var setOfCharacters = lesspass.GetSetOfCharacters([]string{"lowercase"})
	if setOfCharacters != "abcdefghijklmnopqrstuvwxyz" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 26 {
		t.Fatal("len(setOfCharacters) != 26")
	}
}

func testGetSetOfCharactersOnlyUppercase(t *testing.T) {
	var setOfCharacters = lesspass.GetSetOfCharacters([]string{"uppercase"})
	if setOfCharacters != "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 26 {
		t.Fatal("len(setOfCharacters) != 26")
	}
}

func testGetSetOfCharactersOnlyNumbers(t *testing.T) {
	var setOfCharacters = lesspass.GetSetOfCharacters([]string{"numbers"})
	if setOfCharacters != "0123456789" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 10 {
		t.Fatal("len(setOfCharacters) != 10")
	}
}

func testGetSetOfCharactersOnlySymbols(t *testing.T) {
	var setOfCharacters = lesspass.GetSetOfCharacters([]string{"symbols"})
	if setOfCharacters != "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~" {
		t.Fatal("setOfCharacters is incorrect")
	}
	if len(setOfCharacters) != 32 {
		t.Fatal("len(setOfCharacters) != 32")
	}
}

func testGenerateOneCharPerRule(t *testing.T) {
	value, entropy := lesspass.GetOneCharPerRule(big.NewInt(26*26), []string{"lowercase", "uppercase"})
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

func testConfiguredRules(t *testing.T) {
	if strings.Join(lesspass.GetConfiguredRules(lesspass.PasswordProfile{"uppercase": true}), ", ") != "uppercase" {
		t.Fatal("rules != \"uppercase\"")
	}
	if strings.Join(lesspass.GetConfiguredRules(lesspass.PasswordProfile{"uppercase": true, "lowercase": true}), ", ") != "lowercase, uppercase" {
		t.Fatal("rules != \"lowercase, uppercase\"")
	}
	if strings.Join(lesspass.GetConfiguredRules(lesspass.PasswordProfile{"lowercase": true, "symbols": false}), ", ") != "lowercase" {
		t.Fatal("rules != \"lowercase\"")
	}
	if strings.Join(lesspass.GetConfiguredRules(lesspass.PasswordProfile{"lowercase": true, "uppercase": true, "symbols": true, "numbers": true}), ", ") != "lowercase, uppercase, numbers, symbols" {
		t.Fatal("rules != \"lowercase, uppercase, numbers, symbols\"")
	}
}
