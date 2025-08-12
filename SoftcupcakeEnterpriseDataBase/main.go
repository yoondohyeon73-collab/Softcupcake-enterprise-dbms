package main

import (
	"fmt"
	"sedb/modules/parsers"
)

func main() {

	tff := "Title : \"test_table\"\nTABLE_S BEGIN \n COL1 number,COL2 text\nEND DATA_SECTION :\n Data->"

	var tokens []parsers.Tff_token

	err := parsers.ParseHeader(tff, &tokens)
	if err == 1 {
		print("ERROR!")
	}

	for i, tok := range tokens {
		fmt.Printf("Token %d: Type=%v, Value=%v\n", i, tok.Token_type, tok.Token)
	}

}
