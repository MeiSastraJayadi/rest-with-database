package usecase

import (
	"context"
	"database/sql"
	"io"
	"net/http"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
	appmodel "github.com/MeiSastraJayadi/rest-with-datatabase/model"
	"github.com/MeiSastraJayadi/rest-with-datatabase/repository"
	"github.com/gorilla/mux"
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

func (c *CategoryUsecase) Create(w io.Reader, ctx context.Context) error {
	ctg := &appmodel.Category{}
	decodeErr := FromJSON(w, ctg)
	if decodeErr != nil {
		return decodeErr
	}

	ctx = context.WithValue(ctx, repository.ContextValue{}, ctg)
	err := c.db.CreateCategory(ctx, "category")
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryUsecase) Delete(r *http.Request) error {
	vr := mux.Vars(r)
	ctx := r.Context()
	id := vr["id"]
	err := c.db.Delete(ctx, id, "category")
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryUsecase) Update(r *http.Request) error {
	ctx := r.Context()
	data := &model.Category{}
	err := FromJSON(r.Body, data)
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, repository.ContextValue{}, data)
	err = c.db.Update(ctx, "category")
	if err != nil {
		return err
	}
	return nil
}
