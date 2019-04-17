package main

import (
	"database/sql"

	"./services"

	"github.com/gorilla/mux"
)

func initRoutes(router *mux.Router, db *sql.DB) *mux.Router {
	services.NewIndex(router, db)
	services.NewCreate(router, db)
	router.HandleFunc("/coupon/list", ListCoupons).Methods("GET")
	router.HandleFunc("/coupon/retrieve/{id}", RetrieveCoupon).Methods("GET")
	router.HandleFunc("/coupon/create", CreateCoupon).Methods("POST")
	router.HandleFunc("/coupon/update/{id}", UpdateCoupon).Methods("POST")
	return router
}
