package funding

// Go Example: https://gobyexample.com/structs
type Fund struct {
  balance int // private
}

// Function
// Go Example: https://gobyexample.com/functions

// Returns a pointer to FUND
// Go Example: https://gobyexample.com/pointers
func NewFund(initialBalance int) *Fund {
  return &Fund {
    balance: initialBalance
  }
}

// Methods
// Methods start with a *receiver*, in this case a Fund pointer
// Go Example: https://gobyexample.com/methods

// Returns the balance of a Fund pointer.
func (f *Fund) Balance() int {
    return f.balance
}

// Remove the amount from the balance
func (f *Fund) Withdraw(amount int) {
    f.balance -= amount
}
