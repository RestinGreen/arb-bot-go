package main

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/log"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var tokens = make(map[string]Token)
var gqlQueryStrings = make(map[string]GQLQueryString)
var m sync.Mutex
var m2 sync.Mutex

type Token struct {
	ID          int
	Symbol      string
	Address     string
	Blacklisted bool
	decimals    int64
	Minimum     string
}

type GQLQueryString struct {
	pair        string
	queryString string
}

type TokenPairUsage struct {
	SymbolA  string
	SymbolB  string
	AddressA string
	AddressB string
}

func initDB(user string, pass string, host string, port string, db string) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// defer DB.Close()

	getAllTokens()
	getAllGQLQueryStrings()

}

func getAllTokens() {
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
		tokens[token.Address] = Token{
			ID:          token.ID,
			Symbol:      strings.ToUpper(token.Symbol),
			Address:     strings.ToLower(token.Address),
			Blacklisted: token.Blacklisted,
			decimals:    token.decimals,
			Minimum:     token.Minimum,
		}
	}
}

func getAllGQLQueryStrings() {
	sql := `SELECT token_pair, gql_query_string from graphql_query_strings`
	rows, err := DB.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var queryString GQLQueryString
		err = rows.Scan(&queryString.pair, &queryString.queryString)
		if err != nil {
			panic(err)
		}
		gqlQueryStrings[queryString.pair] = GQLQueryString{
			pair:        queryString.pair,
			queryString: queryString.queryString,
		}
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

func insertGQLQueryString(pair string, queryString string) {

	sql := `INSERT INTO graphql_query_strings (token_pair, gql_query_string) values ('` + pair + `','` + queryString + `')`

	m.Lock()
	_, err := DB.Exec(sql)
	if err != nil {
		log.Error(err.Error())
	}
	gqlQueryStrings[pair] = GQLQueryString{
		pair:        pair,
		queryString: queryString,
	}
	m.Unlock()
}

func getTokenUsage() []TokenPairUsage {
	sql := `SELECT symbol_a, address_a, symbol_b, address_b from token_usage where used >= 10`

	rows, err := DB.Query(sql)
	if err != nil {
		panic(err)
	}
	tokenUsages := make([]TokenPairUsage, 0)
	for rows.Next() {
		var tokenUsage TokenPairUsage
		err := rows.Scan(&tokenUsage.SymbolA, &tokenUsage.AddressA, &tokenUsage.SymbolB, &tokenUsage.AddressB)
		if err != nil {
			panic(err)
		}
		tokenUsages = append(tokenUsages, tokenUsage)
	}
	return tokenUsages
}

func updateTokenUsage(tokenA Token, tokenB Token) {

	a, b := sortToken(tokenA, tokenB)

	sql := `SELECT used from token_usage where address_a like '` + a.Address + `' and address_b like '` + b.Address + `'`

	rows, err := DB.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var used int
		err := rows.Scan(&used)
		if err != nil {
			panic(err)
		}
		m2.Lock()
		sql := `UPDATE token_usage set used = $1 where address_a like $2 and address_b like $3`
		_, err = DB.Exec(sql, used+1, a.Address, b.Address)
		m2.Unlock()

	} else {
		sql := `INSERT INTO token_usage (symbol_a, address_a, symbol_b, address_b, used) values ('` + a.Symbol + `','` + a.Address + `','` + b.Symbol + `','` + b.Address + `',0)`
		m2.Lock()
		_, err := DB.Exec(sql)
		if err != nil {
			log.Error(err.Error())
		}
		m2.Unlock()
	}
}
