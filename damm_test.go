package damm

import "testing"

func BenchmarkDammSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DammDigit("123")
	}
}

func BenchmarkDammLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DammDigit("00123014764700968325")
	}
}

func BenchmarkDammSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			DammDigit("123")
		}
	})
}

func BenchmarkDammLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			DammDigit("00123014764700968325")
		}
	})
}

func TestDamm(t *testing.T) {
	if !DammValidate("00123014764700968325") {
		t.Errorf("Checksum failed for valid unput")
	}

	if DammValidate("1234567812345678") {
		t.Errorf("Checksum passed for invalid unput")
	}

	if DammValidate("xy-1") {
		t.Errorf("Checksum passed for invalid alphabet")
	}
}
