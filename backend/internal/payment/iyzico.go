package payment

import (
    "bytes"
    "crypto/hmac"
    "crypto/md5"
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type IyzicoConfig struct {
    APIKey    string
    SecretKey string
    BaseURL   string
}

type Init3DSRequest struct {
    Locale            string  `json:"locale"`
    ConversationID    string  `json:"conversationId"`
    Price             string  `json:"price"`
    PaidPrice         string  `json:"paidPrice"`
    Currency          string  `json:"currency"`
    Installment       int     `json:"installment"`
    BasketID          string  `json:"basketId"`
    PaymentGroup      string  `json:"paymentGroup"`
    PaymentChannel    string  `json:"paymentChannel"`
    CallbackURL       string  `json:"callbackUrl"`
    Buyer             Buyer   `json:"buyer"`
    ShippingAddress   Address `json:"shippingAddress"`
    BillingAddress    Address `json:"billingAddress"`
    BasketItems       []BasketItem `json:"basketItems"`
}

type Buyer struct {
    ID              string `json:"id"`
    Name            string `json:"name"`
    Surname         string `json:"surname"`
    GsmNumber       string `json:"gsmNumber"`
    Email           string `json:"email"`
    IdentityNumber  string `json:"identityNumber"`
    City            string `json:"city"`
    Country         string `json:"country"`
    ZipCode         string `json:"zipCode"`
}

type Address struct {
    ContactName string `json:"contactName"`
    City        string `json:"city"`
    Country     string `json:"country"`
    Address     string `json:"address"`
    ZipCode     string `json:"zipCode"`
}

type BasketItem struct {
    Id       string `json:"id"`
    Name     string `json:"name"`
    Category1 string `json:"category1"`
    ItemType string `json:"itemType"`
    Price    string `json:"price"`
}

type Init3DSResponse struct {
    Status string `json:"status"`
    CheckoutFormContent string `json:"checkoutFormContent"`
    ErrorMessage string `json:"errorMessage"`
}

func randString(n int) string {
    b := make([]byte, n)
    rand.Read(b)
    return hex.EncodeToString(b)[:n]
}

func buildSignature(apiKey, secretKey, rnd, timestamp, body string) string {
    // iyzipay V2 imza adımları (özet):
    // auth = base64(md5(apiKey:secretKey))
    // Doğrulama için dokümantasyona bak. Aşağıda yaygın kullanım:
    h := hmac.New(sha256.New, []byte(secretKey))
    payload := fmt.Sprintf("%s\n%s\n%s\n%s", apiKey, rnd, timestamp, body)
    h.Write([]byte(payload))
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func InitCheckoutForm(cfg IyzicoConfig, req Init3DSRequest) (*Init3DSResponse, error) {
    bodyBytes, _ := json.Marshal(req)
    rnd := randString(16)
    ts := fmt.Sprintf("%d", time.Now().Unix())
    // İmzayı üret
    signature := buildSignature(cfg.APIKey, cfg.SecretKey, rnd, ts, string(bodyBytes))

    httpReq, _ := http.NewRequest(http.MethodPost, cfg.BaseURL+"/v2/checkout/threeds/initialize", bytes.NewReader(bodyBytes))
    httpReq.Header.Set("Content-Type", "application/json")
    httpReq.Header.Set("Authorization", "Iyzipay "+cfg.APIKey+":"+signature)
    httpReq.Header.Set("x-iyzi-rnd", rnd)
    httpReq.Header.Set("x-iyzi-timestamp", ts)

    client := &http.Client{Timeout: 15*time.Second}
    resp, err := client.Do(httpReq)
    if err != nil { return nil, err }
    defer resp.Body.Close()

    var out Init3DSResponse
    if err := json.NewDecoder(resp.Body).Decode(&out); err != nil { return nil, err }
    return &out, nil
}
