package palindrome

import (
	"math/rand"
	"testing"
	"time"
)

// randomPalindromeは、擬似乱数生成器rngから長さと内容が計算された
// 回文を返します。
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // 24までのランダムな長さ
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // '\u0999'までのランダムなルーン
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// 擬似乱数生成器を初期化する。
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func randomNonPalindrome(rng *rand.Rand) string {
	return "abc" + randomPalindrome(rng) + "xyz"
}

func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 500; i++ {
		np := randomNonPalindrome(rng)
		if IsPalindrome(np) {
			t.Errorf("IsPalindrome(%q) = true", np)
		}
	}
}
