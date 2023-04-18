package anagram

import "testing"

func TestMultiIsAnagram(t *testing.T) {
	subtests := []struct {
		s string
		t string
		result bool
	} {
		{
			s: "anagram",
			t: "nagaram",
			result: true,
		},
		{
			s: "rat",
			t: "car",
			result: false,
		},
	}

	for _, st := range subtests {
		if res := isAnagram(st.s, st.t); res != st.result {
			t.Errorf("For %s and %s wanted (%t), got %t", st.s, st.t, res, st.result)
		}
	}
}

func TestMultiSortRunesToString(t *testing.T) {
	subtests := []struct {
		s string
		result string
	}{
		{
			s: "apple",
			result: "aelpp",
		},
		{
			s: "abcdefgasdjfhkajsdf1985",
			result: "1589aaabcdddefffghjjkss",
		},
		{
			s: "0",
			result: "0",
		},
		{	s: "",
			result: "",
		},
	}

	for _, st := range subtests {
		if s := sortRunesToString(st.s); s != st.result {
			t.Errorf("For %s wanted (%v), got %s", st.s, st.result, s)
		}
	} 
}