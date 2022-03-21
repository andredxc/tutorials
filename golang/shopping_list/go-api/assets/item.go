package assets

import "errors"

type Item struct {
	Description string
	Quantity    int
	debugValue  int
	//Category    GroceryCategory
}

func (i *Item) SameAs(item Item) bool {
	return i.Description == item.Description
}

func (i *Item) Add(item Item) error {
	if i.SameAs(item) {
		i.Quantity += item.Quantity
		return nil
	} else {
		return errors.New("[Add] Items are not the same type")
	}
}

func (i *Item) IsValid() bool {
	return i.Description != "" && i.Quantity > 0
}

func (i *Item) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{}, 2)
	result["description"] = i.Description
	result["quantity"] = i.Quantity
	return result
}

func (i *Item) ChangeTo(item Item) {
	i.Description = item.Description
	if item.Quantity > 0 {
		i.Quantity = item.Quantity
	}
}
