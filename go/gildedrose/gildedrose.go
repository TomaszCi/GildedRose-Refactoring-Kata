package gildedrose

import "fmt"

type Item struct {
	Name            string
	SellIn, Quality int
}

type QualityCalculator interface {
	CalculateQuality(sellIn int, quality int) int
}

func max(l, r int) int {
	if l >= r {
		return l
	}

	return r
}

func min(l, r int) int {
	if l <= r {
		return l
	}

	return r
}

type BackstagePass struct{}

func (b *BackstagePass) CalculateQuality(sellIn int, quality int) int {
	if sellIn < 0 {
		return 0
	}

	if quality == 50 {
		return quality
	}

	if sellIn < 10 && sellIn >= 5 {
		return min(quality+2, 50)
	}

	if sellIn < 5 && sellIn >= 0 {
		return min(quality+3, 50)
	}

	return min(quality+1, 50)
}

type BrieCheese struct{}

func (b *BrieCheese) CalculateQuality(sellIn int, quality int) int {
	if quality == 50 {
		return quality
	}

	if sellIn < 0 {
		return min(quality+2, 50)
	}

	return min(quality+1, 50)
}

type Legendary struct{}

func (l *Legendary) CalculateQuality(sellIn int, quality int) int {
	return quality
}

type Ordinary struct{}

func (o *Ordinary) CalculateQuality(sellIn int, quality int) int {
	if quality == 0 {
		return 0
	}

	if sellIn < 0 {
		return max(quality-2, 0)
	}

	return max(quality-1, 0)
}

type Conjured struct{}

func (c *Conjured) CalculateQuality(sellIn int, quality int) int {
	if sellIn <= 0 {
		return max(quality-4, 0)
	}

	return max(quality-2, 0)
}

func (i *Item) String() string {
	return fmt.Sprintf("%s: %d, %d", i.Name, i.SellIn, i.Quality)
}

func UpdateQuality(items []*Item) {
	// for i := 0; i < len(items); i++ {

	// 	if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
	// 		if items[i].Quality > 0 {
	// 			if items[i].Name != "Sulfuras, Hand of Ragnaros" {
	// 				items[i].Quality = items[i].Quality - 1
	// 			}
	// 		}
	// 	} else {
	// 		if items[i].Quality < 50 {
	// 			items[i].Quality = items[i].Quality + 1
	// 			if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
	// 				if items[i].SellIn < 11 {
	// 					if items[i].Quality < 50 {
	// 						items[i].Quality = items[i].Quality + 1
	// 					}
	// 				}
	// 				if items[i].SellIn < 6 {
	// 					if items[i].Quality < 50 {
	// 						items[i].Quality = items[i].Quality + 1
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}

	// 	if items[i].Name != "Sulfuras, Hand of Ragnaros" {
	// 		items[i].SellIn = items[i].SellIn - 1
	// 	}

	// 	if items[i].SellIn < 0 {
	// 		if items[i].Name != "Aged Brie" {
	// 			if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
	// 				if items[i].Quality > 0 {
	// 					if items[i].Name != "Sulfuras, Hand of Ragnaros" {
	// 						items[i].Quality = items[i].Quality - 1
	// 					}
	// 				}
	// 			} else {
	// 				items[i].Quality = items[i].Quality - items[i].Quality
	// 			}
	// 		} else {
	// 			if items[i].Quality < 50 {
	// 				items[i].Quality = items[i].Quality + 1
	// 			}
	// 		}
	// 	}
	// }

	typedItems := make([]QualityCalculator, len(items))
	for i, v := range items {
		switch v.Name {
		case "Aged Brie":
			typedItems[i] = &BrieCheese{}
		case "Backstage passes to a TAFKAL80ETC concert":
			typedItems[i] = &BackstagePass{}
		case "Sulfuras, Hand of Ragnaros":
			typedItems[i] = &Legendary{}
		case "Conjured Mana Cake":
			typedItems[i] = &Conjured{}
		default:
			typedItems[i] = &Ordinary{}
		}
	}

	for i := 0; i < len(items); i++ {
		items[i].SellIn--
		items[i].Quality = typedItems[i].CalculateQuality(items[i].SellIn, items[i].Quality)
	}

}
