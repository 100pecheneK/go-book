package handlers

import "t/internal/utils"

func Handle(h int) int {
	a := utils.Sub(h)
	return utils.Sum(a, h)
}
