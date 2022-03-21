package list_repository

import (
	"errors"
	"go-api/assets"
)

type Repository struct {
	Lists map[uint]*assets.List
}

func (r *Repository) Init() {
	r.Lists = make(map[uint]*assets.List)
}

func (r *Repository) ListForUserId(userId uint) *assets.List {
	if r.Lists == nil {
		r.Init()
	}
	list, _ := r.Lists[userId]
	return list
}

func (r *Repository) CreateListForUserId(userId uint) error {

	if r.Lists == nil {
		r.Init()
	}

	_, ok := r.Lists[userId]
	if ok {
		return errors.New("user already has a list")
	}

	newList := assets.List{}
	newList.Init()
	r.Lists[userId] = &newList
	return nil
}

func (r *Repository) ChangeItem(userId uint, oldItem, newItem assets.Item) (assets.Item, error) {

	l := r.ListForUserId(userId)

	if l == nil {
		return assets.Item{}, errors.New("list is empty")
	}
	item, err := l.ChangeItem(oldItem, newItem)

	if err != nil {
		return assets.Item{}, err
	}

	return item, nil
}

func (r *Repository) RemoveItem(userId uint, item *assets.Item) (*assets.Item, error) {

	l := r.ListForUserId(userId)

	if l == nil {
		return &assets.Item{}, errors.New("list does not exist")
	}

	removedItem, err := l.RemoveItem(*item)

	if err != nil {
		return nil, err
	}

	return removedItem, nil
}

func (r *Repository) AddItem(userId uint, item *assets.Item) (assets.Item, error) {

	l := r.ListForUserId(userId)

	if l == nil {
		return assets.Item{}, errors.New("list is empty")
	}

	err := l.AddItem(*item)

	if err != nil {
		return assets.Item{}, err
	}

	return *item, nil
}
