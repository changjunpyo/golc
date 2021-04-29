package problems

type MovingAverage struct {
	size  int
	array []int
	total int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	a := make([]int, 0)
	mv := MovingAverage{size, a, 0}
	return mv
}

func (this *MovingAverage) Next(val int) float64 {
	this.total += val
	if len(this.array) == this.size {
		this.total -= this.array[0]
		this.array = this.array[1:]
	}
	this.array = append(this.array, val)
	return float64(this.total) / float64(len(this.array))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
