package utils

import "strconv"

func GetPagination(pageStr, limitStr string) (int, int, int) {
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1 // set a default value if the conversion fails
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // set a default value if the conversion fails
	}
	offset := limit * (page - 1)
	return offset, limit, page
}
