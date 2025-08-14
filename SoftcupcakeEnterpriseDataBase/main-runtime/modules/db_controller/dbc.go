package dbcontroller

import (
	"fmt"
	dbinfo "sedb/db_info"
	"sedb/modules/parsers"
)

/*
이게 사양표에 적혀있는 VM입니다
*/

// 에러 출력기
func print_error(err string) int {
	fmt.Println(err)
	return 1
}

func Cmd_exec(script string, db_info dbinfo.DB_info) int {

	//토큰 단위로 자르기
	var script_tokens []parsers.SC_token
	parsers.Parsing_script(script, &script_tokens)

	//에러 체크
	var err_buffer string
	err := parsers.Error_checker(script_tokens, &err_buffer)
	if err == 1 {
		return print_error(err_buffer)
	}

	//첫번째 토큰이 뭔지 확인
	switch script_tokens[0].Token_type {
	//테이블 생성
	case parsers.SC_createTable:

	//데이터 가져오기
	case parsers.SC_get:

	//데이터 업데이트
	case parsers.SC_update:

	//데이터 삭제
	case parsers.SC_delete:

	//데이터 추가
	case parsers.SC_add:

	//뭔가 잘못됨
	default:
		return print_error("Something's wrong!")
	}

	return 0
}
