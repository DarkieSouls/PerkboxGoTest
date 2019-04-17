package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestListCoupons(t *testing.T) {
  req, err := http.NewRequest("GET", "/coupon/list", nil)
  if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(ListCoupons)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("returned wrong status code: got %v want %v", status, http.StatusOK)
  }
}
