package minheap


type TreeNode struct {
	Freq int
	Char rune
	Right *TreeNode
	Left *TreeNode
}

type MinHeap []*TreeNode 

func (mh MinHeap) Len() int {
		return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Freq < mh[j].Freq
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(*TreeNode))
}

func (mh *MinHeap) Pop() any {
	currentHeap := *mh 
	n := len(currentHeap) 

	if n == 0 {
		return nil
	}

	poppedElement := currentHeap[n - 1]
	*mh =  currentHeap[:n-1]
	return poppedElement
}