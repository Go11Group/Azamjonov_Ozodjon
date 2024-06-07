package postgres

import (
	"database/sql"
	"strconv"
	"time"
	"users_and_produncts/model"
)

type Products struct {
	DB *sql.DB
}

func NewProdectsRepo(db *sql.DB) *Products {
	return &Products{DB: db}
}

func (p *Products) CreateProduct(product model.Product) error {
	tr, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `insert into products(name, price, stock_quantity) values($1, $2, $3)`
	_, err = p.DB.Exec(query, product.Name, product.Price, product.Stock_quantity)
	return err
}

func (p *Products) GetProduct(filter model.FilterProduct) (*[]model.Product, error) {
	query := `select * from products`

	isEmptyParams := true
	sch := 1
	params := []interface{}{}
	if filter.Id != nil {
		query += ` where`
		query += ` id=$` + strconv.Itoa(sch)
		params = append(params, filter.Id)
		isEmptyParams = false
		sch++
	}
	if filter.Name != nil {
		if !isEmptyParams {
			query += ` and`
		} else {
			query += ` where`
		}
		query += ` name=$` + strconv.Itoa(sch)
		params = append(params, filter.Name)
		isEmptyParams = false
		sch++
	}
	if filter.Description != nil {
		if !isEmptyParams {
			query += ` and`
		} else {
			query += ` where`
		}
		query += ` description=$` + strconv.Itoa(sch)
		params = append(params, filter.Description)
		isEmptyParams = false
		sch++
	}
	if filter.Price != nil {
		if !isEmptyParams {
			query += ` and`
		} else {
			query += ` where`
		}
		query += ` price=$` + strconv.Itoa(sch)
		params = append(params, filter.Price)
		isEmptyParams = false
		sch++
	}
	if filter.Stock_quantity != nil {
		if !isEmptyParams {
			query += ` and`
		} else {
			query += ` where`
		}
		query += ` stock_quantity=$` + strconv.Itoa(sch)
		params = append(params, filter.Stock_quantity)
		isEmptyParams = false
		sch++
	}
	if isEmptyParams {
		query += ` where deletedat is null`
	} else {
		query += ` and deletedat is null`
	}

	products := []model.Product{}
	rows, err := p.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := model.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description,
			&product.Price, &product.Stock_quantity, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &products, err
}

func (p *Products) UpdateProduct(product *model.Product) error {
	tr, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update products set name=$1, description=$2, price=$3, 
		stock_quantity= $4, updatedAt=$5 where id=$6`
	_, err = tr.Exec(query, product.Name, product.Description, product.Price,
		product.Stock_quantity, time.Now(), product.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Products) DeleteProduct(productID int) error {
	tr, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update products set deletedat=$1 where id=$2`
	_, err = tr.Exec(query, time.Now(), productID)
	if err != nil {
		return err
	}
	return nil
}
