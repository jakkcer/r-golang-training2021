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

	if len(runes) == 0 {
		return ""
	}

	sep := rng.Intn(10)
	for i := 0; i < sep; i++ {
		x := rng.Intn(len(runes))
		runes = append(runes, 0)
		copy(runes[x+1:], runes[x:])
		switch rng.Intn(3) {
		case 0:
			runes[x] = ' '
		case 1:
			runes[x] = ','
		case 2:
			runes[x] = '.'
		}
	}

	return string(runes)
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"owo", true},
		{"Some men interpret nine memos", true},
		{"わたしまけましたわ", true},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf(`IsPalindrome(%q) = %v`, test.input, got)
		}
	}
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
