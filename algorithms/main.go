package algo

import (
	"math/big"
)

// IsPalindrome returns true if plaindrome
// note: this is just created for simple checking
// this is not implmented for the sentenecs
// as you can see the trim methods are not applied and regex too
func IsPalindrome(str string) bool {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	s := string(r)
	return str == s
}

// EdeanGCD recursive gcd method
// wikipidea: 	Euclidean algorithm section
//
//	more efficient method is the Euclidean algorithm, a variant in which the difference of the two
//
// numbers a and b is replaced by the remainder of the Euclidean division
// (also called division with remainder) of a by b.
// Denoting this remainder as a mod b, the algorithm replaces (a, b) with (b, a mod b)
// repeatedly until the pair is (d, 0), where d is the greatest common divisor.
func EdeanGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return EdeanGCD(b, a%b)
}

// EuGCD  algortihm but not o(n)
// wikipida section: Euclid's algorithm
// The method introduced by Euclid for computing greatest common divisors is based on the fact that,
//
//	given two positive integers a and b such that a > b, the common divisors of a and b are the same as the common divisors of a – b and b.
//
// So, Euclid's method for computing
// the greatest common divisor of two positive integers consists of replacing the larger number with the difference of the numbers, and repeating this until the two numbers are equal: that is their greatest common divisor.
func EuGCD(a, b int) int {

	if a == b {
		return a
	}
	if a-b < 0 {
		return EuGCD(a, b-a)
	} else {

		return EuGCD(a-b, b)
	}
}

// LCM simple lcm
func LCM(a, b int) int {
	return b * a / EdeanGCD(a, b)
}

// Factorial returns recursive fact
func Factorial(n int64) *big.Int {
	fact := new(big.Int)
	fact.MulRange(1, n)
	return fact
}

// BigFactor factorial too
func BigFactor(n int64) *big.Int {
	fact := big.NewInt(1)
	one := big.NewInt(1)
	limit := big.NewInt(n)
	if n <= 1 {
		return one
	}
	for p := big.NewInt(2); p.Cmp(limit) <= 0; p.Add(p, one) {
		fact.Mul(fact, p)
	}
	return fact
}

// ModFactorial saw the problem on lead code but still lacks a lot
func ModFactorial(n, mod *big.Int) *big.Int {
	one := big.NewInt(1)
	res := big.NewInt(1)

	for p := big.NewInt(2); p.Cmp(n) <= 0; p.Add(p, one) {
		res.Mul(res, p)
		res.Mod(res, mod)
	}
	return res
}

// isWilsion i forgot the origin
func isWilsion(p int64) bool {
	if p < 2 {
		return false
	}
	bigP := big.NewInt(p)
	one := big.NewInt(1)
	fact := ModFactorial(new(big.Int).Sub(bigP, one), bigP)
	return fact.Cmp(new(big.Int).Sub(fact, one)) == 0
}

// Gen returns the count
func Gen(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		if isWilsion(int64(i)) {
			count++
		}
	}
	return count
}

// PascalTriangle fucked up we'll do later
// func PascalTriangle(n, k int) int {
// 	if k < 0 {
// 		return 0
// 	}
// 	x := big.NewInt(int64(n))
// 	y := big.NewInt(int64(k))
// 	return int(x.Div(Factorial(x), y.Mul(Factorial(y), Factorial(x.Sub(x, y)))).Int64())
// }

// MakePascel it creates the PascalTriangle but current the PascalTriangle is fucked up and i'll solve it later
// func MakePascel(rows int) [][]int {
// 	triangle := make([][]int, rows)
// 	for i := 1; i < rows; i++ {
// 		triangle[i] = make([]int, i+1)
// 		for j := 1; j < i; j++ {
// 			triangle[i][j] = PascalTriangle(i, j)
// 		}
// 	}
// 	return triangle
// }
