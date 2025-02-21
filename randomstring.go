// Package randomstring can be used for generating different types of random strings
package randomstring

import (
	"math/rand"
	"strings"
	"time"
)

var freq = map[rune]int{
	'e': 21912,
	't': 16587,
	'a': 14810,
	'o': 14003,
	'i': 13318,
	'n': 12666,
	's': 11450,
	'r': 10977,
	'h': 10795,
	'd': 7874,
	'l': 7253,
	'u': 5246,
	'c': 4943,
	'm': 4761,
	'f': 4200,
	'y': 3853,
	'w': 3819,
	'g': 3693,
	'p': 3316,
	'b': 2715,
	'v': 2019,
	'k': 1257,
	'x': 315,
	'q': 205,
	'j': 188,
	'z': 128,
}

var freqVowel = map[rune]int{
	'e': 21912,
	'a': 14810,
	'o': 14003,
	'i': 13318,
	'u': 5246,
}

var freqCons = map[rune]int{
	't': 16587,
	'n': 12666,
	's': 11450,
	'r': 10977,
	'h': 10795,
	'd': 7874,
	'l': 7253,
	'c': 4943,
	'm': 4761,
	'f': 4200,
	'y': 3853,
	'w': 3819,
	'g': 3693,
	'p': 3316,
	'b': 2715,
	'v': 2019,
	'k': 1257,
	'x': 315,
	'q': 205,
	'j': 188,
	'z': 128,
}

// freqsum is a sum of all the frequencies in the freq map
var freqsum = func() int {
	n := 0
	for _, v := range freq {
		n += v
	}
	return n
}()

// freqsumVowel is a sum of all the frequencies in the freqVowel map
var freqsumVowel = func() int {
	n := 0
	for _, v := range freqVowel {
		n += v
	}
	return n
}()

// freqsumCons is a sum of all the frequencies in the freqCons map
var freqsumCons = func() int {
	n := 0
	for _, v := range freqCons {
		n += v
	}
	return n
}()

// PickLetter will pick a letter, weighted by the frequency table
func PickLetter() rune {
	target := rand.Intn(freqsum)
	selected := 'a'
	n := 0
	for k, v := range freq {
		n += v
		if n >= target {
			selected = k
			break
		}
	}
	return selected
}

// PickVowel will pick a vowel, weighted by the frequency table
func PickVowel() rune {
	target := rand.Intn(freqsumVowel)
	selected := 'a'
	n := 0
	for k, v := range freqVowel {
		n += v
		if n >= target {
			selected = k
			break
		}
	}
	return selected
}

// PickCons will pick a consonant, weighted by the frequency table
func PickCons() rune {
	target := rand.Intn(freqsumCons)
	selected := 't'
	n := 0
	for k, v := range freqCons {
		n += v
		if n >= target {
			selected = k
			break
		}
	}
	return selected
}

// Seed the random number generator in one of many possible ways.
func Seed() {
	rand.Seed(time.Now().UTC().UnixNano() + 1337)
}

// String generates a random string of a given length.
func String(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = byte(rand.Int63() & 0xff)
	}
	return string(b)
}

// EnglishFrequencyString returns a random string that uses the letter frequency of English,
// ref: http://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html
func EnglishFrequencyString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteRune(PickLetter())
	}
	return sb.String()
}

/*HumanFriendlyString generates a random, but human-friendly, string of
 * the given length. It should be possible to read out loud and send in an email
 * without problems. The string alternates between vowels and consontants.
 *
 * Google Translate believes the output is Samoan.
 *
 * Example output for length 7: rabunor
 */
func HumanFriendlyString(length int) string {
	const (
		someVowels     = "aeoiu"          // a selection of vowels. email+browsers didn't like "æøå" too much
		someConsonants = "bdfgklmnoprstv" // a selection of consonants
		moreLetters    = "chjqwxyz"       // the rest of the letters from a-z
	)
	vowelOffset := rand.Intn(2)
	vowelDistribution := 2
	b := make([]byte, length)
	for i := 0; i < length; i++ {
	again:
		if (i+vowelOffset)%vowelDistribution == 0 {
			b[i] = someVowels[rand.Intn(len(someVowels))]
		} else if rand.Intn(100) > 0 { // 99 of 100 times
			b[i] = someConsonants[rand.Intn(len(someConsonants))]
			// Don't repeat
			if i >= 1 && b[i] == b[i-1] {
				// Also use more vowels
				vowelDistribution = 1
				// Then try again
				goto again
			}
		} else {
			b[i] = moreLetters[rand.Intn(len(moreLetters))]
			// Don't repeat
			if i >= 1 && b[i] == b[i-1] {
				// Also use more vowels
				vowelDistribution = 1
				// Then try again
				goto again
			}
		}
		// Avoid three letters in a row
		if i >= 2 && b[i] == b[i-2] {
			// Then try again
			goto again
		}
	}
	return string(b)
}

// CookieFriendlyString generates a random, but cookie-friendly, string of
// the given length.
func CookieFriendlyString(length int) string {
	const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = allowed[rand.Intn(len(allowed))]
	}
	return string(b)
}

// PKCEFriendlyString generates a random, but PKCE code_verifier spec following,
// string of
// [A-Z] / [a-z] / [0-9] / "-" / "." / "_" / "~"
// with a minimum length of 43 characters and a maximum length of 128 characters.
//
// https://datatracker.ietf.org/doc/html/rfc7636#section-4.1
func PKCECodeVerifierSpecString(length int) string {
	const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	if length < 43 {
		length = 43
	}
	if length > 128 {
		length = 128
	}
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = allowed[rand.Intn(len(allowed))]
	}
	return string(b)
}

/*HumanFriendlyEnglishString generates a random, but human-friendly, string of
 * the given length. It should be possible to read out loud and send in an email
 * without problems. The string alternates between vowels and consontants.
 *
 * The vowels and consontants are wighted by the frequency table
 */
func HumanFriendlyEnglishString(length int) string {
	vowelOffset := rand.Intn(2)
	vowelDistribution := 2
	b := make([]byte, length)
	for i := 0; i < length; i++ {
	again:
		if (i+vowelOffset)%vowelDistribution == 0 {
			b[i] = byte(PickVowel())
		} else if rand.Intn(100) > 0 { // 99 of 100 times
			b[i] = byte(PickCons())
			// Don't repeat
			if i >= 1 && b[i] == b[i-1] {
				// Also use more vowels
				vowelDistribution = 1
				// Then try again
				goto again
			}
		} else {
			b[i] = byte(PickLetter())
			// Don't repeat
			if i >= 1 && b[i] == b[i-1] {
				// Also use more vowels
				vowelDistribution = 1
				// Then try again
				goto again
			}
		}
		// Avoid three letters in a row
		if i >= 2 && b[i] == b[i-2] {
			// Then try again
			goto again
		}
	}
	return string(b)
}
