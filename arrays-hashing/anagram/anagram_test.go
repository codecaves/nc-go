package anagram

import "testing"

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