package assets

import (
	"container/list"
	"errors"
	"time"
)

type List struct {
	items        *list.List
	DateCreated  time.Time
	DateModified time.Time
}

func (s *List) Init() {
	s.items = list.New()
	s.items.Init()
	dateNow := time.Now()
	s.DateCreated = dateNow
	s.DateModified = dateNow
}

func (s *List) AddItem(item Item) error {

	var err error

	if s.items == nil && item.IsValid() {
		s.Init()
	}

	// TODO: This should work without itemCopy, check it
	if item.IsValid() {
		duplicateItem := s.findItem(item)
		if duplicateItem != nil {
			if err = duplicateItem.Add(item); err != nil {
				s.updateModifiedTime()
			}
		} else {
			itemCopy := item
			s.items.PushBack(&itemCopy)
		}
	} else {
		return errors.New("[AddItem] Item is not valid")
	}

	return err
}

func (s *List) ChangeItem(oldItem, newItem Item) (Item, error) {

	itemToChange := s.findItem(oldItem)
	if itemToChange == nil {
		return Item{}, errors.New("item is not in the list")
	}

	itemToChange.Description = newItem.Description
	if newItem.Quantity > 0 {
		itemToChange.Quantity = newItem.Quantity
	}

	// The item has changed (not just quantity, so there might be duplicates)
	if !oldItem.SameAs(newItem) {
		itemToChange = s.findAndAddDuplicates(itemToChange)
	}
	s.updateModifiedTime()
	return *itemToChange, nil
}

func (s *List) findAndAddDuplicates(item *Item) *Item {

	var duplicateElements []*list.Element

	for e := s.items.Front(); e != nil; e = e.Next() {
		listItem := listValueToItem(e)
		if listItem.SameAs(*item) {
			duplicateElements = append(duplicateElements, e)
		}
	}

	if len(duplicateElements) > 0 {
		resultItem := listValueToItem(duplicateElements[0])
		for _, duplicateElement := range duplicateElements[1:] {
			duplicateItem := listValueToItem(duplicateElement)
			resultItem.Add(*duplicateItem)
			s.items.Remove(duplicateElement)
		}
		return resultItem
	}

	return nil
}

func (s *List) updateModifiedTime() {
	s.DateModified = time.Now()
}

func listValueToItem(e *list.Element) *Item {
	item, ok := e.Value.(*Item)
	if !ok {
		panic("list item is not a *Item")
	}
	return item
}

func (s *List) RemoveItem(item Item) (*Item, error) {

	if s.items == nil || s.items.Len() == 0 {
		return nil, errors.New("[RemoveItem] List has no items")
	}

	for e := s.items.Front(); e != nil; e = e.Next() {
		oldItem := listValueToItem(e)
		if oldItem.SameAs(item) {
			s.items.Remove(e)
			s.updateModifiedTime()
			return oldItem, nil
		}
	}
	return nil, errors.New("[RemoveItem] Item not found")
}

func (s *List) findItem(item Item) *Item {

	for e := s.items.Front(); e != nil; e = e.Next() {
		if listItem := listValueToItem(e); listItem.SameAs(item) {
			return listItem
		}
	}
	return nil
}

func (s *List) ToJson() map[string]interface{} {

	result := make(map[string]interface{})
	result["dateCreated"] = s.DateCreated
	result["dateModified"] = s.DateModified

	items := make([]interface{}, s.items.Len())
	i := 0
	for e := s.items.Front(); e != nil; e = e.Next() {
		// TODO: Handle possible panic (return or defer?)
		// A return err does not seem to make sense because it is an internal error
		item := listValueToItem(e)
		items[i] = item.ToJsonMap()
		i++
	}
	result["items"] = items

	return result
}
