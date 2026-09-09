package handler

import (
    "encoding/json"
    "net/http"
)

type PaymentInitReq struct { BookingNo string `json:"booking_no"` }

func InitPayment(w http.ResponseWriter, r *http.Request) {
    var req PaymentInitReq
    json.NewDecoder(r.Body).Decode(&req)
    // iyzico InitCheckoutForm çağrısı burada yapılacak
    // Dönüş: checkout form html içeriği
    json.NewEncoder(w).Encode(map[string]string{
        "checkoutFormContent": "<form>iyzico checkout form embed</form>",
        "paymentPageUrl": "https://sandbox.iyzipay.com",
    })
}

func IyzicoCallback(w http.ResponseWriter, r *http.Request) {
    // iyzico webhook'u: token ile RetrieveCheckoutForm, paymentStatus kontrolü
    // Başarılıysa booking status -> paid, payment kaydı oluştur, kontenjan artır, e-posta gönder
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
}
