package main

import "testing"

func BenchmarkDamm(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Damm("5869720134")
	}
}

func BenchmarkDammParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Damm("5869720134")
		}
	})
}

func TestDamm(t *testing.T) {
	if !Damm("00123014764700968325") {
		t.Errorf("Checksum failed for valid unput")
	}

	if Damm("1234567812345678") {
		t.Errorf("Checksum passed for invalid unput")
	}
}
