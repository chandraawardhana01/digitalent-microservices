package handler

import (
	"net/http"

	"github.com/chandraawardhana01/digitalent-microservices/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}
