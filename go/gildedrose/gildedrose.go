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
	for i, v := range items {
		var typedItem QualityCalculator
		switch v.Name {
		case "Aged Brie":
			typedItem = &BrieCheese{}
		case "Backstage passes to a TAFKAL80ETC concert":
			typedItem = &BackstagePass{}
		case "Sulfuras, Hand of Ragnaros":
			typedItem = &Legendary{}
		case "Conjured Mana Cake":
			typedItem = &Conjured{}
		default:
			typedItem = &Ordinary{}
		}

		items[i].SellIn--
		items[i].Quality = typedItem.CalculateQuality(items[i].SellIn, items[i].Quality)
	}

}
