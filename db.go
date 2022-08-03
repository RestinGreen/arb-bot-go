package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

var DB sql.DB
var tokens = make(map[string]Token)

type Token struct {
	ID			int
	Symbol	 	string
	Address 	string
	Blacklisted	bool
	decimals	int
	Minimum		string

}
func initDB(user string, pass string, host string, port string, db string) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host, port, user, pass, db)

	DB, err := sql.Open("postgres", psqlInfo)

	
	if err != nil {
		panic(err)
	}
	
	defer DB.Close()

	
	sql := `SELECT * from tokens`
	rows, err := DB.Query(sql)
	if err != nil {
		panic(err)
	}
	
	defer rows.Close()

	for rows.Next() {
		var token Token
		err = rows.Scan(&token.ID, &token.Symbol, &token.Address, &token.Blacklisted, &token.decimals, &token.Minimum)
		if err != nil {
			panic(err)
		}
		tokens[token.Address] = token
	}
}

func getTokenDataByAddress(address string) (Token, bool) {
	
	token, found := tokens[strings.ToLower(address)]
	if !found {
		var _token Token
		//TODO insert into db new token
		return _token, false
	} else {
		return token, true
	}
	
}