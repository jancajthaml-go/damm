package damm

import "testing"

func TestDamm(t *testing.T) {
	expectSignature := func(a string, b string) {
		c, err := Generate(a)
		if err != nil {
			t.Errorf("%v", err)
		} else if b != c {
			t.Errorf("for input " + a + " expected signature " + b + " but got " + c + " instead")
		}

		if !Validate(c) {
			t.Errorf("Unable to validate signature that was generated")
		}
	}

	if !Validate("00123014764700968325") {
		t.Errorf("Checksum failed for valid input")
	}

	if Validate("1234567812345678") {
		t.Errorf("Checksum passed for invalid input")
	}

	if Validate("xy-1") {
		t.Errorf("Checksum passed for invalid alphabet")
	}

	expectSignature("0012301476470096832", "00123014764700968325")
}


func BenchmarkDigitSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Digit("123")
	}
}

func BenchmarkDigitLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Digit("00123014764700968325")
	}
}

func BenchmarkDigitSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digit("123")
		}
	})
}

func BenchmarkDigitLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digit("00123014764700968325")
		}
	})
}

func BenchmarkValidateSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Validate("123")
	}
}

func BenchmarkValidateLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Validate("00123014764700968325")
	}
}

func BenchmarkValidateSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Validate("123")
		}
	})
}

func BenchmarkValidateLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Validate("00123014764700968325")
		}
	})
}

func BenchmarkGenerateSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate("123")
	}
}

func BenchmarkGenerateLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate("00123014764700968325")
	}
}

func BenchmarkGenerateSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Generate("123")
		}
	})
}

func BenchmarkGenerateLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Generate("00123014764700968325")
		}
	})
}
