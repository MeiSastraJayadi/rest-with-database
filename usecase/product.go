package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
	"github.com/MeiSastraJayadi/rest-with-datatabase/repository"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type ProductUsecase struct {
	db     *repository.UseDB
	logger hclog.Logger
}

func NewProductUsecase(db *sql.DB, logger hclog.Logger) *ProductUsecase {
	usedb := repository.NewUseDB(logger, db)
	return &ProductUsecase{
		db:     usedb,
		logger: logger,
	}
}

func (pu *ProductUsecase) Create(r *http.Request) error {
	data := &model.Product{}
	err := FromJSON(r.Body, data)
	if err != nil {
		return err
	}
	ctx := r.Context()
	ctx = context.WithValue(ctx, repository.ContextValue{}, data)
	vr := mux.Vars(r)
	id := vr["id"]
	ctx = context.WithValue(ctx, "id", id)
	err = pu.db.CreateProduct(ctx, "product")
	if err != nil {
		return err
	}
	return nil
}

func (pu *ProductUsecase) Delete(r *http.Request) error {
	vr := mux.Vars(r)
	id := vr["id"]
	ctx := r.Context()
	ctx = context.WithValue(ctx, "id", id)
	err := pu.db.Delete(ctx, id, "product")
	if err != nil {
		return err
	}
	return nil
}

func (pu *ProductUsecase) GetAll(r *http.Request) (*model.Products, error) {
	ctx := r.Context()
	productData, err := pu.db.Fetch(ctx, "product")
	if err != nil {
		return nil, err
	}
	var allProduct model.Products
	for productData.Next() {
		var (
			id       int
			name     string
			category *int
			owner    int
			price    int
			stock    int
			address  string
			phone    string
		)
		err = productData.Scan(&id, &name, &category, &owner, &price, &stock)
		if err != nil {
			return nil, err
		}
		product := model.NewProduct(id, name, category, owner, price, stock)

		var categoryData *model.Category = nil
		if category != nil {
			ctxCat := context.WithValue(ctx, "id", strconv.FormatInt(int64(*category), 10))
			thisCategory, catErr := pu.db.FetchByID(ctxCat, "category")
			if catErr != nil {
				return nil, catErr
			}

			if thisCategory.Next() {
				err = thisCategory.Scan(&id, &name)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, errors.New(repository.NotFoundError)
			}
			categoryData = model.NewCategory(id, name)
		}

		ctxOwner := context.WithValue(ctx, "id", strconv.FormatInt(int64(owner), 10))
		thisOwner, ownErr := pu.db.FetchByID(ctxOwner, "owner")
		if ownErr != nil {
			return nil, ownErr
		}

		if thisOwner.Next() {
			err = thisOwner.Scan(&id, &name, &address, &phone)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New(repository.NotFoundError)
		}
		ownerData := model.NewOwner(id, name, address, phone)
		productJson := model.NewProductJSON(product, categoryData, ownerData)
		allProduct = append(allProduct, productJson)
	}
	if len(allProduct) == 0 {
		allProduct = make(model.Products, 0)
	}
	return &allProduct, nil
}
