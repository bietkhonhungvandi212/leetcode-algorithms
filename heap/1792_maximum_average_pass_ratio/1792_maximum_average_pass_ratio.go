package maximumaveragepassratio

//Parent
//Left
//Right
//Size

//Insert
//Extract Max
//BuildHeap

type Node struct {
	Gain        float64
	Pass, Total int
}

type Heap struct {
	nodes []Node
}

func (heap *Heap) Insert(node Node) {
	heap.nodes = append(heap.nodes, node)
	heap.bubbleUp(len(heap.nodes) - 1)
}

func (heap *Heap) ExtractMax() Node {
	if heap.Size() == 0 {
		return Node{}
	}
	max := heap.nodes[0]
	heap.Swap(0, heap.Size()-1)
	heap.nodes = heap.nodes[:heap.Size()-1]

	if heap.Size() > 0 {
		heap.bubbleDown(0)
	}

	return max
}

func (heap *Heap) Average() float64 {
	sum := 0.0
	for _, node := range heap.nodes {
		sum += float64(node.Pass) / float64(node.Total)
	}

	if heap.Size() == 0 {
		return 0.0
	}

	return sum / float64(heap.Size())
}

/** HELPER FUNCTION **/
func (heap *Heap) Parent(i int) int {
	if i <= 0 {
		return -1
	}

	return (i - 1) / 2
}

func (heap *Heap) Left(i int) int {
	return 2*i + 1
}

func (heap *Heap) Right(i int) int {
	return 2*i + 2
}

func (heap *Heap) Size() int {
	return len(heap.nodes)
}

func (heap *Heap) Swap(i, j int) {
	heap.nodes[i], heap.nodes[j] = heap.nodes[j], heap.nodes[i]
}

func (heap *Heap) bubbleUp(index int) {
	for index > 0 {
		p := heap.Parent(index)
		if heap.nodes[p].Gain >= heap.nodes[index].Gain {
			break
		}

		heap.Swap(p, index)
		index = p
	}
}

func (heap *Heap) bubbleDown(index int) {
	for {
		left := heap.Left(index)
		right := heap.Right(index)
		largest := index
		if left < heap.Size() && heap.nodes[left].Gain > heap.nodes[largest].Gain {
			largest = left
		}

		if right < heap.Size() && heap.nodes[right].Gain > heap.nodes[largest].Gain {
			largest = right
		}

		if largest == index {
			break
		}

		heap.Swap(index, largest)
		index = largest
	}
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	cntAssign := extraStudents
	heap := Heap{nodes: []Node{}}
	for _, class := range classes {
		pass, total := class[0], class[1]
		gain := calculateGain(pass, total)
		node := Node{Gain: gain, Pass: pass, Total: total}

		heap.Insert(node)
	}

	for cntAssign > 0 {
		node := heap.ExtractMax()
		node.Pass++
		node.Total++
		node.Gain = calculateGain(node.Pass, node.Total)
		heap.Insert(node)
		cntAssign--
	}

	return heap.Average()
}

func calculateGain(a, b int) float64 {
	pass, total := float64(a), float64(b)
	ratio := (total - pass) / ((total + 1) * total)
	return ratio
}
