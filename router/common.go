package router

import "strconv"

func convertStringToUint64(id string) (uint, error) {
	num, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(num), nil
}
