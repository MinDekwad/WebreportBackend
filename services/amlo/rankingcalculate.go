package amlo

import (
	"context"
	"go-api-report2/ent/predicate"
	"go-api-report2/ent/ranking"
)

func (r AmloService) CalculateRank(date, provincename string) (interface{}, error) {

	//var date string
	// if date == "" {
	// 	tn := time.Now()
	// 	date = tn.Format("2006-01-02")
	// }

	options := []predicate.Ranking{
		ranking.Or(
			ranking.NextDateCalRank(date),
			ranking.StatusRanking("New"),
		),
	}

	// if occuname == "NoOther" {
	// 	options = append(options, ranking.OccupationNameNEQ("Other"))
	// }
	if provincename != "" {
		options = append(options, ranking.ProvinceNameTH(provincename))
	}
	// if manual.Watchlist{
	// 	options = append(options , ranking.OccupationNameNEQ("Other"))
	// }

	count, err := r.connection.Ranking.Query().
		Where(options...).
		Count(context.Background())
	if err != nil {
		return nil, err
	}
	//fmt.Println(count)

	/*
		page 1 = offset 0  => sql limit 0,2000
		page 2 = offset 2000 => sql limit 2000,2000
		page 3 = offset 4000 => sql limit 4000.2000
	*/

	totalPage, perPage := r.getPage(count)
	for page := 1; page <= totalPage; page++ {
		var offset int
		offset = (page - 1) * perPage

		rankings, err := r.connection.Ranking.Query().
			Where(options...).
			Limit(perPage).
			Offset(offset).
			All(context.Background())
		if err != nil {
			return nil, err
		}

		if err := r.CalculateRankDetail(rankings, date); err != nil {
			r.logger.Errorln("[Service] Calcualte detail error", err)
			return nil, err
		}
	}

	return nil, nil
}
