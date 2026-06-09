package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6},
	}

	var itemsAfterUpdateDay1 = []*gildedrose.Item{
		{"+5 Dexterity Vest", 9, 19},
		{"Aged Brie", 1, 1},
		{"Elixir of the Mongoose", 4, 6},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		{"Conjured Mana Cake", 2, 4},
	}

	var itemsAfterUpdateDay10 = []*gildedrose.Item{
		{"+5 Dexterity Vest", 0, 10},
		{"Aged Brie", -8, 18},
		{"Elixir of the Mongoose", -5, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 35},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -5, 0},
		{"Conjured Mana Cake", -7, 0},
	}

	for i := 0; i < 10; i++ {
		gildedrose.UpdateQuality(items)

		if i == 0 {
			if !compareItems(items, itemsAfterUpdateDay1) {
				t.Errorf("items don't match after day 1. should be: %v, got: %v", itemsAfterUpdateDay1, items)
				t.FailNow()
			}
		}
		if i == 9 {
			if !compareItems(items, itemsAfterUpdateDay10) {
				t.Errorf("items don't match after day 10. should be: %v, got: %v", itemsAfterUpdateDay10, items)
				t.FailNow()
			}
		}
	}

}

func compareItems(l, r []*gildedrose.Item) bool {
	if len(r) != len(l) {
		return false
	}

	for i := range l {
		if l[i].Name != r[i].Name {
			return false
		}

		// if l[i].SellIn != r[i].SellIn {
		// 	return false
		// }

		if l[i].Quality != r[i].Quality {
			return false
		}
	}

	return true
}
