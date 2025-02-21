package randomstring

import (
	"fmt"
	"testing"
)

func init() {
	Seed()
}

func TestString(t *testing.T) {
	fmt.Printf("%x\n", String(10))
}

func TestHumanFriendlyString(t *testing.T) {
	fmt.Printf("%s\n", HumanFriendlyString(7))
	fmt.Printf("%s\n", HumanFriendlyString(20))
}

func TestCookieFriendlyString(t *testing.T) {
	fmt.Printf("%s\n", CookieFriendlyString(20))
}

func TestEnglishFrequencyString(t *testing.T) {
	fmt.Printf("%s\n", EnglishFrequencyString(20))
}

func TestHumanFriendlyEnglishString(t *testing.T) {
	fmt.Printf("%s\n", HumanFriendlyEnglishString(7))
	fmt.Printf("%s\n", HumanFriendlyEnglishString(20))
}
