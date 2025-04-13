package controllers

import (
	"net/http"
	helper "superapps/helpers"
	service "superapps/services"
)

func TransactionListPayment(w http.ResponseWriter, r *http.Request) {

	result, err := service.TransactionListPayment()

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]any{})
		return
	}

	helper.Logger("info", "Get All Transaction Payment success")
	helper.Response(w, http.StatusOK, false, "Successfully", result["data"])
}
