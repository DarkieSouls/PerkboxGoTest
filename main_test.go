package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
	"encoding/json"
  "./modal"
  "bytes"

  "github.com/gorilla/mux"
)

func ListRouter() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/coupon/list", ListCoupons).Methods("GET")
  return router
}

func CreateRouter() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/coupon/create", CreateCoupon).Methods("POST")
  return router
}

func UpdateRouter() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/coupon/update/{id}", UpdateCoupon).Methods("POST")
  return router
}
/*
//Incomplete
func TestListCoupons(t *testing.T) {
  req, _ := http.NewRequest("GET", "/coupon/list", nil)
  response := httptest.NewRecorder()
  ListRouter().ServeHTTP(response, req)

  if status := response.Code; status != http.StatusOK {
    t.Errorf("returned wrong status code: got %v want %v", status, http.StatusOK)
  }
}*/

func TestCreateCoupon(t *testing.T) {
  coupon := &modal.Coupon{
    Name: "UnitTestsC",
    Brand: "Go",
    Value: 20,
    CreatedAt: "",
    Expiry: "",
    Redeemed: 0,
  }
  jsonCoupon, _ := json.Marshal(coupon)
  request, _ := http.NewRequest("POST", "/coupon/create", bytes.NewBuffer(jsonCoupon))
  response := httptest.NewRecorder()
  CreateRouter().ServeHTTP(response, request)
  if status := response.Code; status != http.StatusOK {
    t.Errorf("Returned wrong status code: got %v want %v", status, http.StatusOK)
  }
}

func TestUpdateCoupon(t *testing.T) {
  coupon := &modal.Coupon{
    Name: "UnitTestsU",
    Brand: "Go",
    Value: 10,
    CreatedAt: "",
    Expiry: "",
    Redeemed: 0,
  }
  jsonCoupon, _ := json.Marshal(coupon)
  request, _ := http.NewRequest("POST", "/coupon/update/1", bytes.NewBuffer(jsonCoupon))
  response := httptest.NewRecorder()
  UpdateRouter().ServeHTTP(response, request)
  if status := response.Code; status != http.StatusOK {
    t.Errorf("Returned wrong status code: got %v want %v", status, http.StatusOK)
  }
}

func TestUpdateNoCoupon(t *testing.T) {
  coupon := &modal.Coupon{
    Name: "UnitTestsU",
    Brand: "Go",
    Value: 10,
    CreatedAt: "",
    Expiry: "",
    Redeemed: 0,
  }
  jsonCoupon, _ := json.Marshal(coupon)
  request, _ := http.NewRequest("POST", "/coupon/update", bytes.NewBuffer(jsonCoupon))
  response := httptest.NewRecorder()
  UpdateRouter().ServeHTTP(response, request)
  if status := response.Code; status != http.StatusNotFound {
    t.Errorf("Returned wrong status code: got %v want %v", status, http.StatusNotFound)
  }
}
