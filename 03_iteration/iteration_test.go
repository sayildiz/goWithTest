package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := RepeatBuilder("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeatBuilder(b *testing.B) {
	for b.Loop() {
		RepeatBuilder("a")
	}
}

func BenchmarkRepeatConcat(b *testing.B) {
	for b.Loop() {
		RepeatConcat("a")
	}
}
