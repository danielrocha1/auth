package auth

import (
	"database/sql"

	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const (
	driver = "mysql"
	route  = "localhost"
	port   = "3306"
	user   = "r00t"
	pass   = "ABC-123+456"
	dbname = "info_response"
)

func Connect() *sql.DB {

	db, err := sql.Open(driver, "r00t:ABC-123+456@tcp(127.0.0.1:3306)/info_response")
	Check_error(err)
	return db
}

func Check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
} //FEITO

func Filtro_TokenOfertaT(db *sql.DB) []string {
	var token_oferta []string
	var filtro []string
	var check string

	result, err := db.Query("SELECT `TokenOferta` FROM `transportadoras`")
	Check_error(err)

	for result.Next() {
		erro := result.Scan(&check)
		token_oferta = append(token_oferta, check)
		Check_error(erro)

	}
	for item := range token_oferta {
		if token_oferta[item] != check {
			check = token_oferta[item]
			filtro = append(filtro, check)

		}
	}
	if len(filtro) == 0 {
		filtro = append(filtro, check)
	}
	return filtro
} // FEITO   ------------------------- VERIFICAR TOKEN OFERTA Transportadoras

func Filtro_TokenOfertaV(db *sql.DB) []string {
	var token_oferta []string
	var filtro []string
	var check string

	result, err := db.Query("SELECT `TokenOferta` FROM `volumes`")
	Check_error(err)

	for result.Next() {
		erro := result.Scan(&check)
		token_oferta = append(token_oferta, check)
		Check_error(erro)
	}
	for item := range token_oferta {
		if token_oferta[item] != check {
			check = token_oferta[item]
			filtro = append(filtro, token_oferta[item])
		}
	}

	return filtro
} // FEITO   ------------------------- VERIFICAR TOKEN OFERTA Volumes

func Check_TokenOferta(value []string, check string) {
	for item := range value {
		if check == value[item] {
			fmt.Println(value)
			log.Fatal("Token ja inserido no Banco de Dados!")
		}
	}
}

//================================================================  API


