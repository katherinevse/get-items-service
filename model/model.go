package model

type Employee struct {
	ID         int    `json:"id"`
	UID        string `json:"uid"`
	CN         string `json:"cn"`
	Department string `json:"department"`
	Title      string `json:"title"`
	Who        string `json:"who"`
}
