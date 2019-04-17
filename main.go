package main

import (
	"log"
	"net/http"
	"encoding/json"
	"database/sql"
	"./modal"
	"strconv"
	"io"

	"github.com/gorilla/mux"
)

func main() {
	storage := initStorage()

	router := mux.NewRouter()
	router = initRoutes(router, storage)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func ListCoupons(w http.ResponseWriter, r *http.Request) {
	var coupon modal.Coupon
	var s modal.Searcher
	var coupons []modal.Coupon
	db, err := sql.Open("sqlite3", "./coupon.db")
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(r.Body).Decode(&s)
	switch {
	case err == io.EOF:
		coupons, err = coupon.GetCoupons(db)
	case err != nil:
		panic(err)
	default:
		coupons, err = coupon.GetCouponsSearched(s, db)
	}
	for i := 0; i < len(coupons); i++ {
	  json.NewEncoder(w).Encode(coupons[i])
  }
}

//Taking retrieve to mean claim/redeem
func RetrieveCoupon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var c modal.Coupon
	id, err := strconv.ParseInt(params["id"], 10, 64)
	db, err := sql.Open("sqlite3", "./coupon.db")
	coupon := c.GetCoupon(id, db)
	err = c.ClaimCoupon(id, db)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(coupon)
}

func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon modal.Coupon
	db, err := sql.Open("sqlite3", "./coupon.db")
	if err != nil {
		panic(err)
	}
	_ = json.NewDecoder(r.Body).Decode(&coupon)
	_ = coupon.CreateCoupon(db)
}

func UpdateCoupon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var coupon modal.Coupon
	id, err := strconv.ParseInt(params["id"], 10, 64)
	db, err := sql.Open("sqlite3", "./coupon.db")
	if err != nil {
		panic(err)
	}
	_ = json.NewDecoder(r.Body).Decode(&coupon)
	_ = coupon.UpdateCoupon(id, db)
}
