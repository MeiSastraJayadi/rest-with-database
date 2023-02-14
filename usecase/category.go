package usecase

import (
	"database/sql"
	"io"
	"net/http"

	appmodel "github.com/MeiSastraJayadi/rest-with-datatabase/model"
	"github.com/MeiSastraJayadi/rest-with-datatabase/repository"
	"github.com/hashicorp/go-hclog"
)

type CategoryUsecase struct {
	db     *repository.UseDB
	logger hclog.Logger
}

func NewCategoryUsecase(db *sql.DB, logger hclog.Logger) *CategoryUsecase {
	usedb := repository.NewUseDB(logger, db)
	return &CategoryUsecase{
		db:     usedb,
		logger: logger,
	}
}

func (c *CategoryUsecase) Fetch(w io.Writer, r *http.Request) (*appmodel.Categories, error) {
	ctx := r.Context()
	row, err := c.db.Fetch(ctx, "category")
	defer row.Close()
	if err != nil {
		return nil, err
	}

	var categories appmodel.Categories

	for row.Next() {
		var (
			id   int
			name string
		)

		err = row.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		category := appmodel.NewCategory(id, name)
		categories = append(categories, category)
	}

	return &categories, nil
}
