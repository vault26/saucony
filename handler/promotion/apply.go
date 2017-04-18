package promotion

import (
	"encoding/json"
	"net/http"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

type Response struct {
	Error           string  `json:"error,omitempty"`
	DiscountPercent float64 `json:"discount_percent,omitempty"`
}

func ApplyCode(w http.ResponseWriter, r *http.Request) {
	ctx := helper.GetContext(w, r)
	session := ctx["session"].(*model.Session)
	db := ctx["db"].(database.DB)

	decoder := json.NewDecoder(r.Body)
	request := struct{ Code string }{}
	decoder.Decode(&request)
	promotion, err := db.Promotion(request.Code)
	if err != nil {
		js, _ := json.Marshal(Response{Error: err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
	err = session.ApplyPromotion(w, r, promotion)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	js, _ := json.Marshal(Response{DiscountPercent: promotion.DiscountPercent})
	w.Write(js)
}
