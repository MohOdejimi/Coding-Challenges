package huffman

import (
	"container/heap"

	"compressionTool/minheap"
)

func HuffManImplementation(mh minheap.MinHeap) *minheap.TreeNode {
	for mh.Len() > 1 {
		val1 := heap.Pop(&mh)
		val2 := heap.Pop(&mh)

		val1Node := val1.(*minheap.TreeNode)
		val2Node := val2.(*minheap.TreeNode)

		mergedNode := &minheap.TreeNode{
			Freq: val1Node.Freq + val2Node.Freq,
			Left: val1Node,
			Right: val2Node,
		}

		heap.Push(&mh, mergedNode)
	}
	return heap.Pop(&mh).(*minheap.TreeNode)  
}