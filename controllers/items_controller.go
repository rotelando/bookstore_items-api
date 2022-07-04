package controllers

import (
	"encoding/json"
	"github.com/rotelando/bookstore_items-api/domain/items"
	"github.com/rotelando/bookstore_items-api/services"
	"github.com/rotelando/bookstore_items-api/utils/errors"
	"github.com/rotelando/bookstore_items-api/utils/http_utils"
	"github.com/rotelando/bookstore_oauth-go/oauth"
	"io/ioutil"
	"net/http"
)

var (
	ItemController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (it *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.SendError(w, (*errors.RestErr)(err))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := errors.NewBadRequestError("Invalid request body")
		http_utils.SendError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := errors.NewBadRequestError("Invalid json body")
		http_utils.SendError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetClientId(r)

	result, e := services.ItemService.Create(itemRequest)
	if err != nil {
		http_utils.SendError(w, e)
		return
	}

	http_utils.SendJson(w, http.StatusCreated, result)
}

func (it *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
