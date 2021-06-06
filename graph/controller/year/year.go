package year

import (
	"database/sql"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type Year struct {
	DB *sql.DB
}

func (a *Year) GetYears() []*model.Year {
	query, err := a.DB.Prepare("SELECT * FROM years")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query()
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var years []*model.Year

	for res.Next() {
		var year model.Year
		err = res.Scan(&year.ID, &year.Year)

		if err != nil {
			log.Fatal(err)
		}

		years = append(years, &year)
	}

	return years
}
