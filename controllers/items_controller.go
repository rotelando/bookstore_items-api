package controllers

import (
	"fmt"
	"github.com/rotelando/bookstore_items-api/domain/items"
	"github.com/rotelando/bookstore_items-api/services"
	"github.com/rotelando/bookstore_oauth-go/oauth"
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
		//TODO: Return json error to the user.
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemService.Create(item)
	if err != nil {
		//TODO: Return json error to the user
	}

	fmt.Println(result)
	//TODO: Return created item as json with HTTP Status 201 - Created
}

func (it *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
