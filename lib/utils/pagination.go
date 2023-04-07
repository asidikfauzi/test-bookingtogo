package utils

import "strconv"

func GetPagination(pageStr, limitStr string) (int, int, int) {
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 100
	}
	offset := limit * (page - 1)
	return offset, limit, page
}
