package funding

import "testing"

// Doucmentation: https://golang.org/pkg/testing/
func BenchmarkFund(b *testing.B) {
  // Create a fund with many dollars as the number of the current iteration
  // Go Example: https://gobyexample.com/variables
  // var fund Fund = NewFund(b.N)
  fund := NewFund(b.N)

  // Burn all the found until are gone
  for i := 0; i < b.N; i++ {
    fund.Withdraw(1)
  }

  if fund.Balance() != 0 {
     b.Error("Balance wasn't zero:", fund.Balance())
  }
}
