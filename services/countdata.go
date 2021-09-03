package services

import (
	"context"
	"go-api-report2/ent/ranking"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func (r ReportService) CountData(name string) (int, error) {
	count, err := r.connection.Ranking.Query().
		Where(
			ranking.ProvinceNameTH(name),
		).
		Count(context.Background())
	if err != nil {
		return 0, err
	}
	return count, err
}
