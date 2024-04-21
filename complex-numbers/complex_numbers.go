package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	a, b float64
}

// Real is a for a + b * i statement.
func (n Number) Real() float64 {
	return n.a
}

// Imaginary is b for a + b * i statement.
func (n Number) Imaginary() float64 {
	return n.b
}

// Add is (a + i * b) + (c + i * d) = (a + c) + (b + d) * i.
func (n1 Number) Add(n2 Number) Number {
	return Number{
		a: n1.a + n2.a,
		b: n1.b + n2.b,
	}
}

// Subtract is (a + i * b) - (c + i * d) = (a - c) + (b - d) * i.
func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		a: n1.a - n2.a,
		b: n1.b - n2.b,
	}
}

// Multiply is (a + i * b) * (c + i * d) = (a * c - b * d) + (b * c + a * d) * i.
func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		a: n1.a*n2.a - n1.b*n2.b,
		b: n1.b*n2.a + n1.a*n2.b,
	}
}

// Times
func (n Number) Times(factor float64) Number {
	return Number{
		a: n.a * factor,
		b: n.b * factor,
	}
}

// Divide is (a + i * b) / (c + i * d) = (a * c + b * d)/(c^2 + d^2) + (b * c - a * d)/(c^2 + d^2) * i.
func (n1 Number) Divide(n2 Number) Number {
	return Number{
		a: (n1.a*n2.a + n1.b*n2.b) / (math.Pow(n2.a, 2) + math.Pow(n2.b, 2)),
		b: (n1.b*n2.a - n1.a*n2.b) / (math.Pow(n2.a, 2) + math.Pow(n2.b, 2)),
	}
}

// Conjugate a + b * i is a - b * i.
func (n Number) Conjugate() Number {
	return Number{
		a: n.a,
		b: -n.b,
	}
}

// Abs a + b * i is sqrt(a^2 + b^2).
func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.a, 2) + math.Pow(n.b, 2))
}

// Exp is e^(a + i * b) = e^a * e^(i * b), e^(i * b) = cos(b) + i * sin(b).
func (n Number) Exp() Number {
	aExp := math.Exp(n.a)
	return Number{
		a: aExp * math.Cos(n.b),
		b: aExp * math.Sin(n.b),
	}
}
