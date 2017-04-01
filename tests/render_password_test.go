package tests

import (
	"math/big"
	"strings"
	"testing"

	lesspass ".."
)

func testRenderPasswordUseRemainderOfLongDivisionBetweenEntropyAndSetOfCharsLengthAsAnIndex(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = lesspass.PasswordProfile{}
	var firstCharacter = lesspass.RenderPassword([]byte(entropy), passwordProfile)[0]
	if firstCharacter != 'W' {
		t.Fatal("firstCharacter != 'W'")
	}
}

func testRenderPasswordUseQuotientAsSecondEntropyRecursively(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = lesspass.PasswordProfile{}
	var secondCharacter = lesspass.RenderPassword([]byte(entropy), passwordProfile)[1]
	if secondCharacter != 'H' {
		t.Fatal("secondCharacter != 'H'")
	}
}

func testRenderPasswordHasDefaultLengthOfSixteen(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = lesspass.PasswordProfile{}
	var passwordLength = len(lesspass.RenderPassword([]byte(entropy), passwordProfile))
	if passwordLength != 16 {
		t.Fatal("passwordLength != 16")
	}
}

func testRenderPasswordCanSpecifyLength(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = lesspass.PasswordProfile{"length": 20}
	var passwordLength = len(lesspass.RenderPassword([]byte(entropy), passwordProfile))
	if passwordLength != 20 {
		t.Fatal("passwordLength != 20")
	}
}

func testIncludeOneCharPerSetOfCharacters(t *testing.T) {
	var password = lesspass.InsertStringPseudoRandomly("123456", big.NewInt(7*6+2), "uT")
	if password != "T12u3456" {
		t.Fatal("password != \"T12u3456\"")
	}
}

func testRenderPasswordReturnAtLeastOneCharInEveryCharacterSet(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = lesspass.PasswordProfile{"length": 6}
	var generatedPassword = lesspass.RenderPassword([]byte(entropy), passwordProfile)
	var passwordLength = len(generatedPassword)
	var lowercaseOk = strings.ContainsAny(generatedPassword, "abcdefghijklmnopqrstuvwxyz")
	var uppercaseOk = strings.ContainsAny(generatedPassword, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numbersOk = strings.ContainsAny(generatedPassword, "0123456789")
	var symbolsOk = strings.ContainsAny(generatedPassword, "!\"#$%&'()*+,-./:;<: ?@[\\]^_`{|}~")
	if passwordLength != 6 {
		t.Fatal("passwordLength != 6")
	}
	if lowercaseOk && uppercaseOk && numbersOk && symbolsOk != true {
		t.Fatal("there is not at least one char in every characters set")
	}
}
