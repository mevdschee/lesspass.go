package tests

import "testing"

import lesspass ".."

func TestRenderPassword(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	passwordProfile := lesspass.PasswordProfile{}
	generatedPassword := lesspass.GeneratePassword(site, login, masterPassword, passwordProfile)
	if generatedPassword != "WHLpUL)e00[iHR+w" {
		t.Fatal("generatedPassword != \"WHLpUL)e00[iHR+w\"")
	}
}

func TestRenderPasswordNoSymbols(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	passwordProfile := lesspass.PasswordProfile{"length": 14, "counter": 2, "symbols": false}
	generatedPassword := lesspass.GeneratePassword(site, login, masterPassword, passwordProfile)
	if generatedPassword != "MBAsB7b1Prt8Sl" {
		t.Fatal("generatedPassword != \"MBAsB7b1Prt8Sl\"")
	}
}

func TestRenderPasswordOnlyDigits(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	passwordProfile := lesspass.PasswordProfile{"length": 6, "counter": 3, "lowercase": false, "uppercase": false, "symbols": false}
	generatedPassword := lesspass.GeneratePassword(site, login, masterPassword, passwordProfile)
	if generatedPassword != "117843" {
		t.Fatal("generatedPassword != \"117843\"")
	}
}

func TestRenderPasswordNoNumbers(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	passwordProfile := lesspass.PasswordProfile{"length": 14, "numbers": false}
	generatedPassword := lesspass.GeneratePassword(site, login, masterPassword, passwordProfile)
	if generatedPassword != "sB>{qF}wN%/-fm" {
		t.Fatal("generatedPassword != \"sB>{qF}wN%/-fm\"")
	}
}

func TestRenderPasswordWithDefaultOptions(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	generatedPassword := lesspass.GeneratePassword(site, login, masterPassword, nil)
	if generatedPassword != "WHLpUL)e00[iHR+w" {
		t.Fatal("generatedPassword != \"WHLpUL)e00[iHR+w\"")
	}
}
