package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	OneTime_CreditDebitCard                          string
	OneTime_EDC                                      string
	CardInstallment_OmiseInstallment_3Terms          string
	CardInstallment_OmiseInstallment_4Terms          string
	CardInstallment_OmiseInstallment_6Terms          string
	CardInstallment_OmiseInstallment_10Terms         string
	CashInstallment_CreditDebitCard_3Terms           string
	CashInstallment_CreditDebitCard_6Terms           string
	CashInstallment_CreditDebitCard_10Terms          string
	CashInstallment_CreditDebitCardRecurring_3Terms  string
	CashInstallment_CreditDebitCardRecurring_6Terms  string
	CashInstallment_CreditDebitCardRecurring_10Terms string
	CashInstallment_EDC_3Terms                       string
	CashInstallment_EDC_6Terms                       string
	CashInstallment_EDC_10Terms                      string
}

func main() {
	//var cfg Config
	//fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	//fs.StringVar(&cfg.OneTime_CreditDebitCard, "OneTime_CreditDebitCard", "", "")
	//fs.StringVar(&cfg.OneTime_EDC, "OneTime_EDC", "", "")
	//fs.StringVar(&cfg.CardInstallment_OmiseInstallment_3Terms, "CardInstallment_OmiseInstallment_3Terms", "", "")
	//fs.StringVar(&cfg.CardInstallment_OmiseInstallment_4Terms, "CardInstallment_OmiseInstallment_4Terms", "", "")
	//fs.StringVar(&cfg.CardInstallment_OmiseInstallment_6Terms, "CardInstallment_OmiseInstallment_6Terms", "", "")
	//fs.StringVar(&cfg.CardInstallment_OmiseInstallment_10Terms, "CardInstallment_OmiseInstallment_10Terms", "", "")
	//fs.StringVar(&cfg.CashInstallment_CreditDebitCard_3Terms, "CashInstallment_CreditDebitCard_3Terms", "", "")
	//fs.StringVar(&cfg.CashInstallment_CreditDebitCard_6Terms, "CashInstallment_CreditDebitCard_6Terms", "", "")
	//fs.StringVar(&cfg.CashInstallment_CreditDebitCard_10Terms, "CashInstallment_CreditDebitCard_10Terms", "", "")
	//fs.StringVar(&cfg.CashInstallment_CreditDebitCardRecurring_3Terms, "CashInstallment_CreditDebitCardRecurring_3Terms", "", "")
	//fs.StringVar(&cfg.CashInstallment_CreditDebitCardRecurring_6Terms, "CashInstallment_CreditDebitCardRecurring_6Terms", "", "")
	//fs.StringVar(&cfg.CashInstallment_CreditDebitCardRecurring_10Terms, "CashInstallment_CreditDebitCardRecurring_10Terms", "", "")
	//fs.StringVar(&cfg.CashInstallment_EDC_3Terms, "CashInstallment_EDC_3Terms", "", "")
	//fs.StringVar(&cfg.CashInstallment_EDC_6Terms, "CashInstallment_EDC_6Terms", "", "")
	//fs.StringVar(&cfg.CashInstallment_EDC_10Terms, "CashInstallment_EDC_10Terms", "", "")

	fmt.Printf("Starting server...")

	http.HandleFunc("/", helloworld)
	http.HandleFunc("/health_check", check)
	http.HandleFunc("/env", env)

	http.ListenAndServe(":3000", nil)
}

// helloworld just displays a banner message for testing
func helloworld(w http.ResponseWriter, r *http.Request) {
	firstname := os.Getenv("FIRSTNAME")
	lastname := os.Getenv("LASTNAME")
	status := http.StatusOK
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
		<head><title>Hello %s %s</title></head>
		<body><h1>Hello %s %s!</h1></body>
	</html>
	`, firstname, lastname, firstname, lastname)))
}

func check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Health check</h1>`))
}

func env(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(os.Environ(), "\n")))
}
