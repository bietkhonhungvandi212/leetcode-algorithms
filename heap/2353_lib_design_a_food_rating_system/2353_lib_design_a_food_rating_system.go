package libdesignafoodratingsystem

import "container/heap"

type Food struct {
	name string
	rate int
}

type MaxRating []Food

func (h MaxRating) Len() int { return len(h) }
func (h MaxRating) Less(i, j int) bool {
	if h[i].rate != h[j].rate {
		return h[i].rate > h[j].rate
	}

	return h[i].name < h[j].name
}
func (h MaxRating) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxRating) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Food))
}

func (h *MaxRating) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRate    map[string]int
	cuisineHeap   map[string]*MaxRating
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRate:    map[string]int{},
		cuisineHeap:   make(map[string]*MaxRating),
	}

	for i, food := range foods {
		cuisine := cuisines[i]
		rate := ratings[i]
		fr.foodToCuisine[food] = cuisine
		fr.foodToRate[food] = rate
		if _, ok := fr.cuisineHeap[cuisine]; !ok {
			fr.cuisineHeap[cuisine] = &MaxRating{}
			heap.Init(fr.cuisineHeap[cuisine])
		}
		heap.Push(fr.cuisineHeap[cuisine], Food{name: food, rate: rate})
	}

	return fr

}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.foodToRate[food] = newRating
	cuisine := this.foodToCuisine[food]

	heap.Push(this.cuisineHeap[cuisine], Food{food, newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	h := this.cuisineHeap[cuisine]
	for h.Len() > 0 {
		food := (*h)[0]
		if this.foodToRate[food.name] == food.rate {
			return food.name
		}
		heap.Pop(h)
	}

	return ""
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
