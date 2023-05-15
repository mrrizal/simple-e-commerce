package models

import "time"

type SignUpResponseOk struct {
	Token string `json:"token"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ProductList struct {
	Results []Product `json:"results"`
}

type Rekening struct {
	NoRekening string  `json:"no_rekening"`
	NasabahID  int     `json:"nasabah_id"`
	Saldo      float64 `json:"saldo"`
}

type TransaksiRequest struct {
	NoRekening   string  `json:"no_rekening"`
	Nominal      float64 `json:"nominal"`
	CurrentSaldo float64 `json:"current_saldo"`
	Type         string  `json:"type"`
}

type TransaksiResponseOk struct {
	Saldo float64 `json:"saldo"`
}

type MutasiTransaksi struct {
	Waktu         time.Time `json:"waktu"`
	KodeTransaksi string    `json:"kode_transaksi"`
	Nominal       float64   `json:"nominal"`
}

type MutasiResp struct {
	Results []MutasiTransaksi `json:"results"`
}
