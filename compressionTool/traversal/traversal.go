package traversal


import(
	"compressionTool/minheap"
)

type Stack struct {
	Node *minheap.TreeNode 
	CurrentCode string 
}

type StackSlice []*Stack 

var StckSlice StackSlice

func PreOrderTraversal(stckSlice StackSlice, codes map[rune]string) {
	for len(stckSlice) != 0 {
		poppedNode := stckSlice[len(stckSlice) - 1]
		stckSlice = stckSlice[:len(stckSlice) - 1]

		if poppedNode.Node.Right == nil && poppedNode.Node.Left == nil {
			codes[poppedNode.Node.Char] = poppedNode.CurrentCode
		}

		if poppedNode.Node.Right != nil {
			rightChild := &Stack{
				Node: poppedNode.Node.Right,
				CurrentCode: poppedNode.CurrentCode + "1",
			}
			stckSlice = append(stckSlice, rightChild)
		} 

		if poppedNode.Node.Left != nil {
			leftChild := &Stack{
				Node: poppedNode.Node.Left,
				CurrentCode: poppedNode.CurrentCode + "0",
			}
			stckSlice = append(stckSlice, leftChild)
		}
	}
}