package modal

import (
	"database/sql"
	"fmt"
)

type Coupon struct {
	ID        int    `json:"ID"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	Value     int    `json:"value"`
	CreatedAt string `json:"createdAt"`
	Expiry    string `json:"expiry"`
	Redeemed	int 	 `json:"redeemed"`
}

type Searcher struct {
	Term			string	`json:"term"`
	Criteria	string	`json:"criteria"`
}

func (c *Coupon) GetCoupons(db *sql.DB) ([]Coupon, error) {
	stmt := "SELECT * FROM COUPON WHERE REDEEMED = 0"
	rows, err := db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	coupons := []Coupon{}

	for rows.Next() {
		var c Coupon
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Brand,
			&c.Value,
			&c.CreatedAt,
			&c.Expiry,
			&c.Redeemed,
		); err != nil {
			return nil, err
		}
		coupons = append(coupons, c)
	}

	return coupons, nil
}

func (c *Coupon) GetCouponsSearched(s Searcher, db *sql.DB) ([]Coupon, error) {
	stmt := fmt.Sprintf("SELECT * FROM COUPON WHERE REDEEMED = 0 AND %s = '%s'",
		s.Term,
		s.Criteria,
  )
	rows, err := db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	coupons := []Coupon{}

	for rows.Next() {
		var c Coupon
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Brand,
			&c.Value,
			&c.CreatedAt,
			&c.Expiry,
			&c.Redeemed,
		); err != nil {
			return nil, err
		}
		coupons = append(coupons, c)
	}

	return coupons, nil
}

func (c *Coupon) GetCoupon(id int64, db *sql.DB) (Coupon) {
	stmt := fmt.Sprintf("SELECT * FROM COUPON WHERE ID = %d AND REDEEMED = 0 LIMIT 1", id)
	var cp Coupon

	err := db.QueryRow(stmt).Scan(
		&cp.ID,
		&cp.Name,
		&cp.Brand,
		&cp.Value,
		&cp.CreatedAt,
		&cp.Expiry,
		&cp.Redeemed,
	)
	if err != nil {
	  panic(err)
	}
	return cp
}

func (c *Coupon) CreateCoupon(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO coupon(name, brand, value, createdAt, expiry, redeemed) VALUES('%s', '%s', %d, '%s', '%s', '%d')",
		c.Name,
		c.Brand,
		c.Value,
		c.CreatedAt,
		c.Expiry,
		c.Redeemed,
	)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}

func (c *Coupon) UpdateCoupon(id int64, db *sql.DB) error {
	stmt := fmt.Sprintf("UPDATE COUPON SET NAME = '%s', BRAND = '%s', VALUE = %d, CREATEDAT = '%s', EXPIRY = '%s', REDEEMED = %d WHERE ID = %d",
	  c.Name,
	  c.Brand,
		c.Value,
		c.CreatedAt,
		c.Expiry,
		c.Redeemed,
		id,
		)
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

func (c *Coupon) ClaimCoupon(id int64, db *sql.DB) error {
	stmt := fmt.Sprintf("UPDATE COUPON SET REDEEMED = 1 WHERE ID = %d", id)
	_, err := db.Exec(stmt)
	if err != nil {
		return(err)
	}
	return nil
}
