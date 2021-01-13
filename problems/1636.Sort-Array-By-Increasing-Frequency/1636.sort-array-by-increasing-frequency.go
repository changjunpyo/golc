package problems

import "sort"

func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for _ , val := range nums{
		m[val]++
	}

	arr := make(Pairs, len(m))
	idx := 0
	for k , v := range m {
		arr[idx] = Pair{k,v}
		idx++
	}
	sort.Sort(arr)
	idx =0
	for _,  val := range arr{
		for j:=0; j< val.Freq; j++{
			nums[idx] = val.Num
			idx++
		}
	}

	return nums
}

type Pair struct{
	Num int
	Freq int
}

type Pairs []Pair

func (p Pairs) Len() int { return len(p)}
func (p Pairs) Less(i, j int) bool {
	if p[i].Freq == p[j].Freq{
		return p[i].Num > p[j].Num
	}
	return p[i].Freq < p[j].Freq
}
func (p Pairs) Swap(i,j int) {p[i], p[j] = p[j] , p[i]}
