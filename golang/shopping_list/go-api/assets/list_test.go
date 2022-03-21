package assets

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestList_AddItem(t *testing.T) {

	// TODO: Split this test into: test add, test update item, test invalid item

	r := require.New(t)
	shoppingList := List{}
	itemsInList := 0

	// Add first element
	newItem := Item{Description: "Soap", Quantity: 1}
	err := shoppingList.AddItem(newItem)
	itemsInList++
	r.Nil(err)
	r.Equal(newItem.Description, shoppingList.items.Back().Value.(*Item).Description)
	r.Equal(newItem.Quantity, shoppingList.items.Back().Value.(*Item).Quantity)
	r.Equal(itemsInList, shoppingList.items.Len())

	// Add second element
	newItem = Item{Description: "Paper", Quantity: 1}
	err = shoppingList.AddItem(newItem)
	itemsInList++
	r.Nil(err)
	r.Equal(newItem.Description, shoppingList.items.Back().Value.(*Item).Description)
	r.Equal(newItem.Quantity, shoppingList.items.Back().Value.(*Item).Quantity)
	r.Equal(itemsInList, shoppingList.items.Len())

	// Add repeating element
	newItem = Item{Description: "Paper", Quantity: 2}
	err = shoppingList.AddItem(newItem)
	r.Nil(err)
	r.Equal(newItem.Description, shoppingList.items.Back().Value.(*Item).Description)
	r.Equal(newItem.Quantity+1, shoppingList.items.Back().Value.(*Item).Quantity)
	r.Equal(itemsInList, shoppingList.items.Len())

	// Add element with invalid quantity
	newItem = Item{Description: "Toilet Paper", Quantity: 0}
	err = shoppingList.AddItem(newItem)
	r.NotNil(err)
	r.Equal(itemsInList, shoppingList.items.Len())

	// Add element with another invalid quantity
	newItem = Item{Description: "Toilet Paper", Quantity: -1}
	err = shoppingList.AddItem(newItem)
	r.NotNil(err)
	r.Equal(itemsInList, shoppingList.items.Len())

	// Add element with invalid name
	newItem = Item{Description: "", Quantity: 2}
	err = shoppingList.AddItem(newItem)
	r.NotNil(err)
	r.Equal(itemsInList, shoppingList.items.Len())
}

func TestList_RemoveItem(t *testing.T) {

	r := require.New(t)
	shoppingList := List{}

	// Remove an item from an empty list
	item, err := shoppingList.RemoveItem(Item{Description: "Paper"})
	r.NotNil(err)
	r.Nil(item)

	// Add a few items to delete
	paperItem := Item{Description: "Paper", Quantity: 1}
	err = shoppingList.AddItem(paperItem)
	r.Nil(err)
	appleItem := Item{Description: "Apple", Quantity: 2}
	err = shoppingList.AddItem(appleItem)
	r.Nil(err)
	r.Equal(2, shoppingList.items.Len())

	// Delete an item that does not exist
	item, err = shoppingList.RemoveItem(Item{Description: "Banana"})
	r.NotNil(err)
	r.Nil(item)
	r.Equal(2, shoppingList.items.Len())

	// Delete both items from the list
	item, err = shoppingList.RemoveItem(Item{Description: "Paper"})
	r.Nil(err)
	r.Equal(1, shoppingList.items.Len())
	r.Equal(paperItem, *item)
	item, err = shoppingList.RemoveItem(Item{Description: "Apple"})
	r.Equal(appleItem, *item)
	r.Equal(0, shoppingList.items.Len())
}

func Test_findItem(t *testing.T) {

	r := require.New(t)
	shoppingList := List{}
	shoppingList.Init()

	// Create a few items
	bananaItem := Item{Description: "Banana", Quantity: 1}
	appleItem := Item{Description: "Apple", Quantity: 5}
	grapeItem := Item{Description: "Grape", Quantity: 10}

	// Try to find an item in an empty list
	itemFound := shoppingList.findItem(bananaItem)
	r.Nil(itemFound)

	// Add a few items
	shoppingList.AddItem(bananaItem)
	shoppingList.AddItem(appleItem)
	shoppingList.AddItem(grapeItem)
	r.Equal(3, shoppingList.items.Len())

	// Try to find the recently added items
	itemFound = shoppingList.findItem(bananaItem)
	r.NotNil(itemFound)
	r.Equal(bananaItem.Quantity, itemFound.Quantity)
	r.Equal(bananaItem.Description, itemFound.Description)

	itemFound = shoppingList.findItem(appleItem)
	r.NotNil(itemFound)
	r.Equal(appleItem.Quantity, itemFound.Quantity)
	r.Equal(appleItem.Description, itemFound.Description)

	itemFound = shoppingList.findItem(grapeItem)
	r.NotNil(itemFound)
	r.Equal(grapeItem.Quantity, itemFound.Quantity)
	r.Equal(grapeItem.Description, itemFound.Description)
}

func Test_ToJson(t *testing.T) {

	r := require.New(t)
	shoppingList := List{}
	shoppingList.Init()

	// Populate the list
	shoppingList.AddItem(Item{Description: "Banana", Quantity: 1})
	shoppingList.AddItem(Item{Description: "Apple", Quantity: 5})
	shoppingList.AddItem(Item{Description: "Grape", Quantity: 10})

	jsonList := shoppingList.ToJson()

	date, ok := jsonList["dateCreated"]
	r.True(ok)
	r.Equal(shoppingList.DateCreated, date)

	date, ok = jsonList["dateModified"]
	r.True(ok)
	r.Equal(shoppingList.DateModified, date)

	// Find each of the items in the shopping list in the json representation
	jsonItems := jsonList["items"].([]interface{})
	itemsFound := 0
	for e := shoppingList.items.Front(); e != nil; e = e.Next() {
		item := e.Value.(*Item)

		for _, jsonItem := range jsonItems {
			itemMap := jsonItem.(map[string]interface{})
			description, ok := itemMap["description"]
			r.True(ok)
			quantity, ok := itemMap["quantity"]
			r.True(ok)
			if description == item.Description && quantity == item.Quantity {
				itemsFound++
				break
			}
		}
	}
	r.Equal(3, itemsFound)
}

func Test_ChangeItem(t *testing.T) {

	r := require.New(t)
	shoppingList := List{}
	shoppingList.Init()

	// Populate the list
	bananaItem := Item{Description: "Banana", Quantity: 1}
	appleItem := Item{Description: "Apple", Quantity: 5}
	grapeItem := Item{Description: "Grape", Quantity: 10}
	shoppingList.AddItem(bananaItem)
	shoppingList.AddItem(appleItem)
	shoppingList.AddItem(grapeItem)

	changedBanana := Item{Description: "Not a banana"}
	changedApple := Item{Description: "Not an apple", Quantity: 1}
	changedGrape := Item{Description: "Not a grape", Quantity: -1}

	changedItem, err := shoppingList.ChangeItem(bananaItem, changedBanana)
	r.Equal(3, shoppingList.items.Len())
	r.Nil(err)
	r.Equal(changedBanana.Description, changedItem.Description)
	r.Equal(bananaItem.Quantity, changedItem.Quantity) // Zero quantity should be ignored
	changedItem, err = shoppingList.ChangeItem(appleItem, changedApple)
	r.Nil(err)
	r.Equal(changedApple.Description, changedItem.Description)
	r.Equal(changedApple.Quantity, changedItem.Quantity)
	changedItem, err = shoppingList.ChangeItem(grapeItem, changedGrape)
	r.Nil(err)
	r.Equal(changedGrape.Description, changedItem.Description)
	r.Equal(grapeItem.Quantity, changedItem.Quantity) // Negative quantity should be ignored
}

func Test_ChangeItemToDuplicate(t *testing.T) {

	r := require.New(t)
	shoppingList := List{}
	shoppingList.Init()

	// Populate the list
	grapeItem := Item{Description: "Grape", Quantity: 2}
	bananaItem := Item{Description: "Banana", Quantity: 1}
	appleItem := Item{Description: "Apple", Quantity: 5}
	shoppingList.AddItem(grapeItem)
	shoppingList.AddItem(bananaItem)
	shoppingList.AddItem(appleItem)

	changedBanana := Item{Description: "Apple"}
	changedItem, err := shoppingList.ChangeItem(bananaItem, changedBanana)
	r.Equal(2, shoppingList.items.Len())
	r.Nil(err)
	r.Equal(changedBanana.Description, changedItem.Description)
	r.Equal(bananaItem.Quantity+appleItem.Quantity, changedItem.Quantity)

	changedGrape := Item{Description: "Apple", Quantity: 6}
	changedItem, err = shoppingList.ChangeItem(grapeItem, changedGrape)
	r.Equal(1, shoppingList.items.Len())
	r.Nil(err)
	r.Equal(changedGrape.Description, changedItem.Description)
	r.Equal(changedGrape.Quantity+bananaItem.Quantity+appleItem.Quantity, changedItem.Quantity)
}

func Test_Copy(t *testing.T) {

	item1 := Item{Description: "item1", Quantity: 1, debugValue: 1}
	item2 := item1

	log.Printf("item1=%p, item2=%p\n", &item1, &item2)
	log.Printf("item1=%v, item2=%v\n", &item1, &item2)

	item2.Description = "item2"
	item2.Quantity = 2
	item2.debugValue = 2

	log.Printf("item1=%p, item2=%p\n", &item1, &item2)
	log.Printf("item1=%v, item2=%v\n", &item1, &item2)
}
