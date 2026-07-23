package cpuid

import "testing"

func TestParseFeatureTrimSpace(t *testing.T) {
	s := SSE2.String()
	got := ParseFeature("  " + s + " ")
	if got != SSE2 {
		t.Fatalf("got %v want SSE2 (%q)", got, s)
	}
}
