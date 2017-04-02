package lesspass

import (
	"math/big"
	"strings"
	"testing"
)

func TestRenderPasswordUseRemainderOfLongDivisionBetweenEntropyAndSetOfCharsLengthAsAnIndex(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = GetPasswordProfile(PasswordProfile{})
	var firstCharacter = RenderPassword([]byte(entropy), passwordProfile)[0]
	if firstCharacter != 'W' {
		t.Fatal("firstCharacter != 'W'")
	}
}

func TestRenderPasswordUseQuotientAsSecondEntropyRecursively(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = GetPasswordProfile(PasswordProfile{})
	var secondCharacter = RenderPassword([]byte(entropy), passwordProfile)[1]
	if secondCharacter != 'H' {
		t.Fatal("secondCharacter != 'H'")
	}
}

func TestRenderPasswordHasDefaultLengthOfSixteen(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = GetPasswordProfile(PasswordProfile{})
	var passwordLength = len(RenderPassword([]byte(entropy), passwordProfile))
	if passwordLength != 16 {
		t.Fatal("passwordLength != 16")
	}
}

func TestRenderPasswordCanSpecifyLength(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = GetPasswordProfile(PasswordProfile{"length": 20})
	var passwordLength = len(RenderPassword([]byte(entropy), passwordProfile))
	if passwordLength != 20 {
		t.Fatal("passwordLength != 20")
	}
}

func TestIncludeOneCharPerSetOfCharacters(t *testing.T) {
	var password = InsertStringPseudoRandomly("123456", big.NewInt(7*6+2), "uT")
	if password != "T12u3456" {
		t.Fatal("password != \"T12u3456\"")
	}
}

func TestRenderPasswordReturnAtLeastOneCharInEveryCharacterSet(t *testing.T) {
	var entropy = "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e"
	var passwordProfile = GetPasswordProfile(PasswordProfile{"length": 6})
	var generatedPassword = RenderPassword([]byte(entropy), passwordProfile)
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
