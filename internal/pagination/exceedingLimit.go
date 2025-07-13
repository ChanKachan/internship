package pagination

import (
	"fmt"
	"go.uber.org/zap"
	"internship/internal/logg"
)

type ExceedingLimit interface {
	PaginateInventory(limit, offset int) (int, int, int, error)
}

type exceedingLimitItems struct {
	count  int
	limit  int
	offset int
}

func NewExceedingLimitItems(count, limit, offset int) *exceedingLimitItems {
	return &exceedingLimitItems{count: count, limit: limit, offset: offset}
}

func (p *exceedingLimitItems) CheckForExceedingLimit() error {
	logg.Logger.Info("Проверка допустимых значений.",
		zap.String("package", "pagination.CheckForExceedingLimit"))

	if p.offset < 0 {
		p.offset = 0
	}

	if p.limit <= 0 {
		p.limit = p.count
	}

	if p.count < p.limit+p.offset && p.count <= p.offset {
		err := fmt.Errorf("Offset или Limit привышает Total")
		logg.Logger.Error(err.Error(),
			zap.String("package", "pagination.CheckForExceedingLimit"))
		return err
	}

	logg.Logger.Info("Проверка прошла успешно.",
		zap.String("package", "pagination.CheckForExceedingLimit"))
	return nil
}
