package designafoodratingsystem

type node struct {
	name string
	rate int
}

type MaxHeap struct {
	arr []node
}

func newMaxHeap() *MaxHeap {
	return &MaxHeap{
		arr: make([]node, 0),
	}
}

func (this *MaxHeap) insert(food string, rate int) {
	this.arr = append(this.arr, node{name: food, rate: rate})
	this.bubbleUp(len(this.arr) - 1)
}

func (this *MaxHeap) peek(foodToRate map[string]int) string {
	for len(this.arr) > 0 {
		top := this.arr[0]
		if top.rate == foodToRate[top.name] {
			return top.name
		}
		this.pop()
	}
	return ""
}

func (this *MaxHeap) pop() {
	if len(this.arr) == 0 {
		return
	}
	lastIdx := len(this.arr) - 1
	this.arr[0] = this.arr[lastIdx]
	this.arr = this.arr[:lastIdx]
	if len(this.arr) > 0 {
		this.bubbleDown(0)
	}
}

func (this *MaxHeap) bubbleUp(idx int) {
	for idx > 0 {
		parentIdx := this.parent(idx)
		if parentIdx < 0 {
			break
		}
		parent := this.arr[parentIdx]
		child := this.arr[idx]
		if parent.rate > child.rate ||
			(parent.rate == child.rate && parent.name <= child.name) {
			break
		}
		this.swap(idx, parentIdx)
		idx = parentIdx
	}
}

func (this *MaxHeap) bubbleDown(idx int) {
	for {
		left := this.left(idx)
		right := this.right(idx)
		largest := idx
		if left < len(this.arr) {
			if this.arr[left].rate > this.arr[largest].rate ||
				(this.arr[left].rate == this.arr[largest].rate && this.arr[left].name < this.arr[largest].name) {
				largest = left
			}
		}
		if right < len(this.arr) {
			if this.arr[right].rate > this.arr[largest].rate ||
				(this.arr[right].rate == this.arr[largest].rate && this.arr[right].name < this.arr[largest].name) {
				largest = right
			}
		}
		if largest == idx {
			break
		}
		this.swap(idx, largest)
		idx = largest
	}
}

func (this *MaxHeap) swap(i, j int) {
	if i == j {
		return
	}
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
}

func (this *MaxHeap) left(idx int) int {
	return 2*idx + 1
}

func (this *MaxHeap) right(idx int) int {
	return 2*idx + 2
}

func (this *MaxHeap) parent(idx int) int {
	if idx <= 0 {
		return -1
	}
	return (idx - 1) / 2
}

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRate    map[string]int
	cuisineHeap   map[string]*MaxHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRate:    make(map[string]int),
		cuisineHeap:   make(map[string]*MaxHeap),
	}
	for i, food := range foods {
		cuisine := cuisines[i]
		fr.foodToCuisine[food] = cuisine
		fr.foodToRate[food] = ratings[i]
		if _, ok := fr.cuisineHeap[cuisine]; !ok {
			fr.cuisineHeap[cuisine] = newMaxHeap()
		}
		fr.cuisineHeap[cuisine].insert(food, ratings[i])
	}
	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine := this.foodToCuisine[food]
	this.foodToRate[food] = newRating
	this.cuisineHeap[cuisine].insert(food, newRating)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.cuisineHeap[cuisine].peek(this.foodToRate)
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
