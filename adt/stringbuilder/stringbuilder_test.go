package stringbuilder

import "testing"

func TestStringBuilder(t *testing.T) {
	s1, s2, s3 := "science ", "is ", "amazing"
	ra := s1 + s2 + s3

	sb := New()
	sb.Append(s1)
	sb.Println()
	sb.Append(s2)
	sb.Println()
	sb.Append(s3)
	sb.Println()

	if sb.ToString() != ra {
		t.Errorf("sb: %v, want: %v, got: %v\n", sb.ToString(), ra, sb.ToString())
	}
}
