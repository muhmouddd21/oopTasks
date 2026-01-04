package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Transaction struct {
	Time    time.Time
	Message string
}

type Account interface {
	GetBalance() float64

	Deposit(amount float64) error
	Withdraw(amount float64) error
}

// DTO
// because we need to show data only
type AccountStatement struct {
	AccountNumber string
	AccountHolder string
	From          time.Time
	To            time.Time
	Transactions  []Transaction
	Balance       float64
}
type BaseAccount struct {
	accountNumber string
	accountHolder string
	balance       float64
	transactions  []Transaction
}
type SavingsAccount struct {
	*BaseAccount
	interestRate   float64
	minimumBalance float64
}
type CheckingAccount struct {
	*BaseAccount
	overdraftLimit float64
	transactionFee float64
}
type BusinessAccount struct {
	*BaseAccount
	overdraftLimit float64
}

type LoanAccount struct {
	*BaseAccount
	interestRate float64
}

type Bank struct {
	bankName string
	accounts map[string]Account
	counter  int
}

func newBankAccount(accountNumber string, accountHolder string, initial float64) *BaseAccount {
	return &BaseAccount{
		accountNumber: accountNumber,
		accountHolder: accountHolder,
		balance:       initial,
		transactions: []Transaction{
			{
				Time:    time.Now(),
				Message: fmt.Sprintf("Account opened with $%.2f", initial),
			},
		},
	}
}
func NewSavingsAccount(number, holder string, initial float64) *SavingsAccount {
	return &SavingsAccount{
		BaseAccount:    newBankAccount(number, holder, initial),
		interestRate:   0.02,
		minimumBalance: 100,
	}
}
func NewCheckingAccount(number, holder string, initial float64) *CheckingAccount {
	return &CheckingAccount{
		BaseAccount:    newBankAccount(number, holder, initial),
		overdraftLimit: 500,
		transactionFee: 0.5,
	}
}
func NewBusinessAccount(number, holder string, initial float64) *BusinessAccount {
	return &BusinessAccount{
		BaseAccount:    newBankAccount(number, holder, initial),
		overdraftLimit: 2000,
	}
}

func NewLoanAccount(number, holder string, loanAmount float64) *LoanAccount {
	if loanAmount <= 0 {
		panic("loan amount must be greater than zero")
	}

	return &LoanAccount{
		BaseAccount: &BaseAccount{
			accountNumber: number,
			accountHolder: holder,
			balance:       -loanAmount,
			transactions: []Transaction{
				{
					Time:    time.Now(),
					Message: fmt.Sprintf("A new loan accound created with loan amount %v", loanAmount),
				},
			},
		},
		interestRate: 0.05,
	}
}

func NewBank(name string) *Bank {
	return &Bank{
		bankName: name,
		accounts: make(map[string]Account),
		counter:  1000,
	}
}

func (b *BaseAccount) GetBalance() float64 {
	return b.balance
}

func (b *BaseAccount) addTransaction(message string) {
	b.transactions = append(b.transactions, Transaction{
		Time:    time.Now(),
		Message: message,
	})
}

func (b *BaseAccount) DisplayTransactions() {
	for _, tx := range b.transactions {
		fmt.Printf(
			"[%s] %s\n",
			tx.Time.Format("2006-01-02 15:04:05"),
			tx.Message,
		)
	}
}
func (b *BaseAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid deposit amount")
	}
	b.balance += amount
	b.addTransaction(fmt.Sprintf("Deposit: +$%.2f", amount))
	return nil
}

func (b *BaseAccount) CanBeClosed() error {
	if b.balance != 0 {
		return errors.New("account balance must be zero to close")
	}
	return nil
}

func (s *SavingsAccount) Withdraw(amount float64) error {
	if s.balance-amount < s.minimumBalance {
		return errors.New("minimum balance violation")
	}
	s.balance -= amount
	s.addTransaction(fmt.Sprintf("Withdrawal: -$%.2f", amount))
	return nil

}
func (s *SavingsAccount) ApplyInterest() {
	interest := s.balance * s.interestRate
	s.balance += interest
	s.addTransaction(fmt.Sprintf("Interest applied: +$%.2f", interest))
}
func (c *CheckingAccount) Withdraw(amount float64) error {
	total := amount + c.transactionFee
	if c.balance-total < -c.overdraftLimit {
		return errors.New("overdraft limit exceeded")
	}

	c.balance -= total
	c.addTransaction(fmt.Sprintf(
		"Withdrawal: -$%.2f (fee $%.2f)",
		amount, c.transactionFee,
	))
	return nil
}

// Add monthly maintenance fees for checking accounts
func (c *CheckingAccount) ApplyMonthlyMaintenanceFee() {
	c.balance -= c.transactionFee
	c.addTransaction(
		fmt.Sprintf("Monthly maintenance fee: -$%.2f", c.transactionFee),
	)
}

func (b *BusinessAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid withdrawal amount")
	}

	if b.balance-amount < -b.overdraftLimit {
		return errors.New("business overdraft limit exceeded")
	}

	b.balance -= amount
	b.addTransaction(fmt.Sprintf("Withdrawal: -$%.2f", amount))
	return nil
}
func (l *LoanAccount) Withdraw(amount float64) error {
	return errors.New("withdrawals are not allowed on loan accounts")
}

func (l *LoanAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("repayment amount must be greater than zero")
	}

	l.balance += amount
	l.addTransaction(fmt.Sprintf("Loan repayment: +$%.2f", amount))
	return nil
}

func (l *LoanAccount) ApplyInterest() {
	if l.balance >= 0 {
		return // loan already settled
	}

	interest := -l.balance * l.interestRate
	l.balance -= interest

	l.addTransaction(
		fmt.Sprintf("Loan interest applied: -$%.2f", interest),
	)
}

func (b *Bank) nextAccountNumber() string {
	b.counter++
	return strconv.Itoa(b.counter)
}

func (b *Bank) CreateSavingsAccount(holder string, initialDeposit float64) Account {
	savingsAccount := NewSavingsAccount(b.nextAccountNumber(), holder, initialDeposit)
	b.accounts[savingsAccount.accountNumber] = savingsAccount
	return savingsAccount
}

func (b *Bank) CreateCheckingAccount(holder string, initialDeposit float64) Account {
	savingsAccount := NewCheckingAccount(b.nextAccountNumber(), holder, initialDeposit)
	b.accounts[savingsAccount.accountNumber] = savingsAccount
	return savingsAccount
}

func (b *Bank) CreateLoanAccount(holder string, loanAmount float64) Account {
	loan := NewLoanAccount(b.nextAccountNumber(), holder, loanAmount)
	b.accounts[loan.accountNumber] = loan
	return loan
}

func (b *Bank) FindAccount(accountNumber string) (Account, bool) {
	account, ok := b.accounts[accountNumber]
	return account, ok
}
func (b *Bank) GetTotalBankBalance() float64 {
	total := 0.0
	for _, acc := range b.accounts {
		total += acc.GetBalance()
	}
	return total
}

func (b *Bank) DisplayAllAccounts() {
	fmt.Printf("=== All Accounts in %s ===\n", b.bankName)

	for number, acc := range b.accounts {
		fmt.Printf(
			"Account Number: %s | Balance: $%.2f\n",
			number,
			acc.GetBalance(),
		)
	}
}

func (b *Bank) Transfer(from, to string, amount float64) error {
	fromAcc, ok := b.accounts[from]
	if !ok {
		return errors.New("source account not found")
	}
	toAcc, ok := b.accounts[to]
	if !ok {
		return errors.New("destination account not found")
	}
	if err := fromAcc.Withdraw(amount); err != nil {
		return err
	}

	if err := toAcc.Deposit(amount); err != nil {
		fromAcc.Deposit(amount)
		return err
	}

	return nil
}

func (b *Bank) ApplyMonthlyMaintenanceFees() {
	for _, acc := range b.accounts {
		if checking, ok := acc.(*CheckingAccount); ok {
			checking.ApplyMonthlyMaintenanceFee()
		}
	}
}

func (b *Bank) CloseAccount(accountNumber string) error {
	acc, ok := b.accounts[accountNumber]
	if !ok {
		return errors.New("account not found")
	}

	// نحتاج نوصل للـ BaseAccount
	var base *BaseAccount

	switch a := acc.(type) {
	case *SavingsAccount:
		base = a.BaseAccount
	case *CheckingAccount:
		base = a.BaseAccount
	default:
		return errors.New("unsupported account type")
	}

	// Rule check
	if err := base.CanBeClosed(); err != nil {
		return err
	}

	// إزالة الحساب من البنك
	delete(b.accounts, accountNumber)

	return nil
}

func (b *Bank) GenerateAccountStatement(
	accountNumber string,
	from, to time.Time,
) (*AccountStatement, error) {

	if accountNumber == "" {
		return nil, errors.New("account number is required")
	}

	acc, ok := b.accounts[accountNumber]
	if !ok {
		return nil, errors.New("account not found")
	}

	// نوصل للـ BaseAccount
	var base *BaseAccount
	switch a := acc.(type) {
	case *SavingsAccount:
		base = a.BaseAccount
	case *CheckingAccount:
		base = a.BaseAccount
	case *LoanAccount:
		base = a.BaseAccount
	default:
		return nil, errors.New("unsupported account type")
	}

	var txs []Transaction
	for _, tx := range base.transactions {
		if (tx.Time.After(from) || tx.Time.Equal(from)) &&
			(tx.Time.Before(to) || tx.Time.Equal(to)) {
			txs = append(txs, tx)
		}
	}

	return &AccountStatement{
		AccountNumber: base.accountNumber,
		AccountHolder: base.accountHolder,
		From:          from,
		To:            to,
		Transactions:  txs,
		Balance:       base.balance,
	}, nil
}

func main() {
	bank := NewBank("First National Bank")

	savings := bank.CreateSavingsAccount("Alice", 1000)
	checking := bank.CreateCheckingAccount("Bob", 500)

	bank.DisplayAllAccounts()
	savings.Deposit(5000.00)
	checking.Deposit(6000.00)
	fmt.Printf("Alice's balance: $ %v ", savings.GetBalance())
	err := savings.Withdraw(8000.00)
	if err != nil {
		log.Fatal("Error in withdraw ", err)
	}

	fmt.Println("Total balance:", bank.GetTotalBankBalance())

}
