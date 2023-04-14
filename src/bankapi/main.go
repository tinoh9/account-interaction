package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/msft/bank"
)

type CustomAccount struct {
    *bank.Account
}

var accounts = map[float64]*CustomAccount{}

func deposit(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            err := account.Deposit(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account.Statement())
            }
        }
    }
}

func withdraw(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number")
    } else {
        if account, ok := accounts[number]; !ok {
            fmt.Fprintf(w, "Account with number %v can't be found", number)
        } else {
            err := account.Withdraw(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account.Statement())
            }
        }
    }
}

func transfer(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")
    desqs := req.URL.Query().Get("des")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number")
    } else if des, err := strconv.ParseFloat(desqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account destination number")
    } else {
        if account1, ok := accounts[number]; !ok {
            fmt.Fprintf(w, "Account with number %v can't be found", number)
        } else if account2, ok := accounts[des]; !ok {
            fmt.Fprintf(w, "Account with number %v can't be found", des)
        } else {
            err := account1.Transfer(amount, account2.Account)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account1.Statement())
            }
        }
    }
}

// Statement method converts endpoint from string format to JSON format
func (c *CustomAccount) Statement() string {
    json, err := json.Marshal(c)
    if err != nil {
        return err.Error()
    }

    return string(json)
}

func statement(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else {
    if account, ok := accounts[number]; !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            json.NewEncoder(w).Encode(bank.Statement(account))
        }
    }
}

func main() {
    accounts[101] = &CustomAccount{
        Account: *bank.Account{
            Customer: bank.Customer{
                Name: "Adam",
                Address: "Westminster, California",
                Phone: "(123) 456 789",
            },
            Number: 101,
	    },
    }

    accounts[102] = &CustomAccount{
        Account: *bank.Account{
            Customer: bank.Customer{
    			Name: "Eve",
                Address: "Los Angeles, California",
                Phone: "(000) 000 000",
            },
            Number: 102,
	    },
    }

    http.HandleFunc("/statement", statement)
    http.HandleFunc("/deposit", deposit)
    http.HandleFunc("/withdraw", withdraw)
    http.HandleFunc("/transfer", transfer)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}