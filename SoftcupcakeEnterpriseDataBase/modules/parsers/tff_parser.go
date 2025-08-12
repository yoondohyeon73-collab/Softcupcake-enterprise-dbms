package parsers

import (
	"strconv"
	"strings"
)

type Tff_tokenT int

const (
	Tff_none Tff_tokenT = iota
	// 선언 키워드
	Tff_title       // 타이틀
	Tff_TableS      // 테이블 구조 키워드
	Tff_Cnumber     // 열 타입 숫자
	Tff_Ctext       // 열 타입 문자/문자열
	Tff_ColumnName  // 열 이름
	Tff_dataSection // 데이터 저장섹션 표시

	// 속성 토큰 추가
	Tff_Notnull // NOTNULL
	Tff_Key     // KEY

	// 토큰 타입
	Tff_string  // 문자열/문자
	Tff_float64 // float64 타입 숫자

	// 특수 토큰
	Tff_begin     // 시작 키워드 (BEGIN)
	Tff_end       // 종료 키워드 (END)
	Tff_dataStart // Data-> (데이터 시작 표시)
	Tff_dataEnd   // ->End (데이터 종료 표시)
	Tff_comma     // , (콤마)
)

type Tff_token struct {
	Token      interface{}
	Token_type Tff_tokenT
}

// ParseHeader는 DATA_SECTION 전까지 파싱 (타이틀 + 테이블 구조)
func ParseHeader(input string, tokens *[]Tff_token) int {
	lines := strings.Split(input, "\n")
	inTable := false

	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "DATA_SECTION") {
			*tokens = append(*tokens, Tff_token{Token: "DATA_SECTION", Token_type: Tff_dataSection})
			return 0
		}

		if strings.HasPrefix(line, "Title") {
			*tokens = append(*tokens, Tff_token{Token: "Title", Token_type: Tff_title})
			parts := strings.SplitN(line, ":", 2)
			if len(parts) < 2 {
				return 1
			}
			val := strings.TrimSpace(parts[1])
			*tokens = append(*tokens, Tff_token{Token: val, Token_type: Tff_string})
			continue
		}

		if strings.HasPrefix(line, "TABLE_S") && strings.Contains(line, "BEGIN") {
			*tokens = append(*tokens, Tff_token{Token: "TABLE_S", Token_type: Tff_TableS})
			*tokens = append(*tokens, Tff_token{Token: "BEGIN", Token_type: Tff_begin})
			inTable = true
			continue
		}

		if inTable && strings.HasPrefix(line, "END") {
			*tokens = append(*tokens, Tff_token{Token: "END", Token_type: Tff_end})
			inTable = false
			continue
		}

		if inTable {
			cols := strings.Split(line, ",")
			for _, col := range cols {
				col = strings.TrimSpace(col)
				if col == "" {
					continue
				}
				parts := strings.Fields(col)
				if len(parts) < 2 {
					return 1
				}

				colType := parts[0]
				colName := parts[1]

				switch strings.ToLower(colType) {
				case "number":
					*tokens = append(*tokens, Tff_token{Token: colType, Token_type: Tff_Cnumber})
				case "text":
					*tokens = append(*tokens, Tff_token{Token: colType, Token_type: Tff_Ctext})
				default:
					*tokens = append(*tokens, Tff_token{Token: colType, Token_type: Tff_none})
				}

				*tokens = append(*tokens, Tff_token{Token: colName, Token_type: Tff_ColumnName})

				for _, attr := range parts[2:] {
					switch strings.ToUpper(attr) {
					case "NOTNULL":
						*tokens = append(*tokens, Tff_token{Token: "NOTNULL", Token_type: Tff_Notnull})
					case "KEY":
						*tokens = append(*tokens, Tff_token{Token: "KEY", Token_type: Tff_Key})
					}
				}
			}
		}
	}

	return 0
}

// ParseDataSection는 DATA_SECTION 이후 전체 데이터 섹션 파싱
func ParseDataSection(input string, tokens *[]Tff_token) int {
	lines := strings.Split(input, "\n")

	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "{") && strings.HasSuffix(line, "}") {
			*tokens = append(*tokens, Tff_token{Token: "{", Token_type: Tff_begin})
			dataStr := strings.Trim(line, "{} ")
			items := strings.Split(dataStr, ",")
			for i, it := range items {
				val := strings.TrimSpace(it)
				if f, err := strconv.ParseFloat(val, 64); err == nil {
					*tokens = append(*tokens, Tff_token{Token: f, Token_type: Tff_float64})
				} else {
					*tokens = append(*tokens, Tff_token{Token: val, Token_type: Tff_string})
				}
				if i < len(items)-1 {
					*tokens = append(*tokens, Tff_token{Token: ",", Token_type: Tff_comma})
				}
			}
			*tokens = append(*tokens, Tff_token{Token: "}", Token_type: Tff_end})
		}
	}

	return 0
}

// ParseDataLine은 Data-> [데이터...] ->End 한줄만 파싱
func ParseDataLine(line string, tokens *[]Tff_token) int {
	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, "Data->") || !strings.HasSuffix(line, "->End") {
		return 1
	}

	*tokens = append(*tokens, Tff_token{Token: "Data->", Token_type: Tff_dataStart})

	body := strings.TrimPrefix(line, "Data->")
	body = strings.TrimSuffix(body, "->End")
	body = strings.TrimSpace(body)
	body = strings.Trim(body, "[]")
	items := strings.Split(body, ",")
	for i, it := range items {
		val := strings.TrimSpace(it)
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			*tokens = append(*tokens, Tff_token{Token: f, Token_type: Tff_float64})
		} else {
			*tokens = append(*tokens, Tff_token{Token: val, Token_type: Tff_string})
		}
		if i < len(items)-1 {
			*tokens = append(*tokens, Tff_token{Token: ",", Token_type: Tff_comma})
		}
	}

	*tokens = append(*tokens, Tff_token{Token: "->End", Token_type: Tff_dataEnd})
	return 0
}
