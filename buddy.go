package gobuddy

type Buddy struct {
	size int
	tree []int
}

func isPowerOfTwo(num int) bool {
	if (num & (num - 1)) == 0 {
		return true
	}
	return false
}

func fitPowerOfTwo(size int, bsize int) int {
	if size > bsize {
		return 0
	}
	for {
		if bsize/2 < size {
			return bsize
		} else {
			bsize /= 2
		}
	}
}

func leftNode(index int) int {
	return 2*index + 1
}

func rightNode(index int) int {
	return 2*index + 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewBuddySystem(zsize int) *Buddy {
	if zsize < 1 || !isPowerOfTwo(zsize) {
		return nil
	}
	nodesize := zsize

	b := new(Buddy)
	b.size = zsize
	b.tree = make([]int, 2*zsize-1)
	b.tree[0] = zsize
	for i := 1; i < 2*zsize-1; i++ {
		if isPowerOfTwo(i + 1) {
			nodesize /= 2
		}
		b.tree[i] = nodesize
	}
	return b
}

func (b *Buddy) Alloc(zsize int) int {
	if zsize < 1 {
		return -1
	}

	if !isPowerOfTwo(zsize) {
		zsize = fitPowerOfTwo(zsize, b.size)
	}

	index := 0
	nodesize := b.size
	if zsize > b.tree[index] {
		return -1
	}
	for ; nodesize != zsize; nodesize /= 2 {
		if b.tree[leftNode(index)] >= zsize {
			index = leftNode(index)
		} else {
			index = rightNode(index)
		}
	}
	b.tree[index] = 0
	offset := (index+1)*nodesize - b.size

	for index > 0 {
		index = (index - 1) / 2
		b.tree[index] = max(b.tree[rightNode(index)], b.tree[leftNode(index)])

	}
	return offset
}

func (b *Buddy) Free(offset int) {
	nodesize := 1
	index := offset + b.size - 1
	for ; b.tree[index] > 0; index = (index - 1) / 2 {
		nodesize *= 2
		if index == 0 {
			return
		}
	}
	b.tree[index] = nodesize

	for ; index > 0; nodesize *= 2 {
		index = (index - 1) / 2

		if b.tree[leftNode(index)] == nodesize && b.tree[rightNode(index)] == nodesize {
			b.tree[index] = nodesize * 2
		} else {
			b.tree[index] = max(b.tree[leftNode(index)], b.tree[rightNode(index)])
		}
	}
}
