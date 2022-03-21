package router

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-api/assets"
	"go-api/list_repository"
	"log"
	"net/http"
	"strconv"
)

type ListRouter struct {
	ListRepository list_repository.Repository
}

func (r *ListRouter) CreateList(c echo.Context) error {

	var err error
	var userId uint

	if err, userId = readUserIdFromContext(c); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	log.Printf("[CreateList] Request from userId=%d\n", userId)

	err = r.ListRepository.CreateListForUserId(userId)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "List created")
}

func (r *ListRouter) AddItem(c echo.Context) error {

	var item *assets.Item
	var itemAdded assets.Item
	var err error
	var userId uint

	// TODO: Move this logic into ListRepository, so that the router does not need to know about any of the ListRepository specifics
	if err, userId = readUserIdFromContext(c); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	log.Printf("[AddItem] Request from userId=%d\n", userId)

	if item, err = createItemFromContext(c); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if itemAdded, err = r.ListRepository.AddItem(userId, item); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, itemAdded)
}

func (r *ListRouter) ChangeItem(c echo.Context) error {

	var oldItem, newItem *assets.Item
	var finalItem assets.Item
	var err error
	var userId uint

	if err, userId = readUserIdFromContext(c); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid userid=%d", userId))
	}

	if oldItem, err = createItemFromContext(c); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if newItem, err = createItemChangeFromContext(c); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if finalItem, err = r.ListRepository.ChangeItem(userId, *oldItem, *newItem); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, finalItem)
}

func (r *ListRouter) RemoveItem(c echo.Context) error {

	var err error
	var userId uint
	var item, removedItem *assets.Item

	if err, userId = readUserIdFromContext(c); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid userid=%d", userId))
	}

	if item, err = createItemFromContext(c); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if removedItem, err = r.ListRepository.RemoveItem(userId, item); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, removedItem)
}

func (r *ListRouter) GetList(c echo.Context) error {

	var err error
	var userId uint

	if err, userId = readUserIdFromContext(c); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	l := r.ListRepository.ListForUserId(userId)
	log.Printf("[GetList] Request from userId=%d\n", userId)

	if l == nil {
		return c.String(http.StatusOK, fmt.Sprintf("No list for userid=%d", userId))
	}

	return c.JSON(http.StatusOK, l.ToJson())
}

func readUserIdFromContext(c echo.Context) (error, uint) {

	userIdStr := c.QueryParam("userid")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	return err, uint(userId)
}

func createItemFromContext(c echo.Context) (*assets.Item, error) {

	var err error
	var quantity int

	description := c.QueryParam("desc")
	quantityStr := c.QueryParam("qtd")

	if description == "" {
		return nil, errors.New("no description provided")
	}
	if quantityStr != "" {
		if quantity, err = strconv.Atoi(quantityStr); err != nil {
			return nil, err
		}
		if quantity < 1 {
			return nil, errors.New("quantity must be greater than 0")
		}
	} else {
		quantity = 1
	}

	return &assets.Item{
		Description: description,
		Quantity:    quantity,
	}, nil
}

func createItemChangeFromContext(c echo.Context) (*assets.Item, error) {

	// TODO: Refactor this and createItemFromContext, lots of repeated code
	var err error
	var quantity int

	description := c.QueryParam("newDesc")
	quantityStr := c.QueryParam("newQtd")
	if quantityStr != "" {
		quantity, err = strconv.Atoi(quantityStr)
	} else {
		quantity = 0
	}

	if err != nil {
		return nil, err
	}

	return &assets.Item{
		Description: description,
		Quantity:    quantity,
	}, nil
}
