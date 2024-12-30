# Pointers

* [pointers](https://gobyexample.com/pointers)

```go
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
```

(note, you do not need to deference the pointer `(*w).balance`)
