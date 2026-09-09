package handler

import (
    "encoding/json"
    "net/http"
)

func AdminUpdateTour(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status":"updated"})
}
func AdminDeleteTour(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status":"deleted"})
}
func ListPayments(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]interface{}{"data":[]})
}
func RefundPayment(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status":"refund_requested"})
}
func CreateCoupon(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status":"coupon_created"})
}
