package inventoryPagination

import (
	"go.uber.org/zap"
	"internship/internal/logg"
	"internship/internal/pagination"
)

type PaginateInventory interface {
	PaginateInventory(limit, offset int) (int, int, int, error)
}

type paginateItemsInventory struct {
	count paginationItems
}

func NewPaginateItemsInventory() *paginateItemsInventory {
	return &paginateItemsInventory{}
}

func (p *paginateItemsInventory) PaginateInventory(limit, offset int) (int, int, int, error) {
	logg.Logger.Info("Запуск метод пагинации.",
		zap.String("package", "inventoryPagination.PaginateInventory"))
	count := p.count.inventoryCount()

	item := pagination.NewExceedingLimitItems(count, limit, offset)

	err := item.CheckForExceedingLimit()

	if err != nil {
		return 0, 0, 0, err
	}

	logg.Logger.Info("Пагинация применена.",
		zap.String("package", "inventoryPagination.PaginateInventory"))

	return limit, offset, count, nil
}
