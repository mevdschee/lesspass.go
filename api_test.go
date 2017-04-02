package lesspass

import "testing"

func TestrenderPassword(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	passwordProfile := PasswordProfile{}
	generatedPassword := GeneratePassword(site, login, masterPassword, passwordProfile)
	if generatedPassword != "WHLpUL)e00[iHR+w" {
		t.Fatal("generatedPassword != \"WHLpUL)e00[iHR+w\"")
	}
}

func TestrenderPasswordNoSymbols(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	passwordProfile := PasswordProfile{"length": 14, "counter": 2, "symbols": false}
	generatedPassword := GeneratePassword(site, login, masterPassword, passwordProfile)
	if generatedPassword != "MBAsB7b1Prt8Sl" {
		t.Fatal("generatedPassword != \"MBAsB7b1Prt8Sl\"")
	}
}

func TestrenderPasswordOnlyDigits(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	passwordProfile := PasswordProfile{"length": 6, "counter": 3, "lowercase": false, "uppercase": false, "symbols": false}
	generatedPassword := GeneratePassword(site, login, masterPassword, passwordProfile)
	if generatedPassword != "117843" {
		t.Fatal("generatedPassword != \"117843\"")
	}
}

func TestrenderPasswordNoNumbers(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	passwordProfile := PasswordProfile{"length": 14, "numbers": false}
	generatedPassword := GeneratePassword(site, login, masterPassword, passwordProfile)
	if generatedPassword != "sB>{qF}wN%/-fm" {
		t.Fatal("generatedPassword != \"sB>{qF}wN%/-fm\"")
	}
}

func TestrenderPasswordWithDefaultOptions(t *testing.T) {
	site := "example.org"
	login := "contact@example.org"
	masterPassword := "password"
	generatedPassword := GeneratePassword(site, login, masterPassword, nil)
	if generatedPassword != "WHLpUL)e00[iHR+w" {
		t.Fatal("generatedPassword != \"WHLpUL)e00[iHR+w\"")
	}
}
