package binarytree

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64 |
	~string
}

type Tree[T Ordered] struct {
	head *node[T]
}

type node[T Ordered] struct {
	left, right *node[T]
	v           T
}

func NewTree[T Ordered]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Add(v T) {
	t.head = insert(t.head, v)
}

func insert[T Ordered](n *node[T], v T) *node[T] {
	if n == nil {
		return &node[T]{v: v}
	}

	if v < n.v {
		n.left = insert(n.left, v)
	} else if v > n.v {
		n.right = insert(n.right, v)
	}
	// Если v == n.v, не добавляем дубликат

	return n
}

func (t *Tree[T]) Remove(v T) {
	t.head = deleteNode(t.head, v)
}

func deleteNode[T Ordered](n *node[T], v T) *node[T] {
	if n == nil {
		return nil
	}

	if v < n.v {
		n.left = deleteNode(n.left, v)
	} else if v > n.v {
		n.right = deleteNode(n.right, v)
	} else {
		// Нашли узел для удаления
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		// У узла есть оба потомка
		minNode := findMin(n.right)
		n.v = minNode.v
		n.right = deleteNode(n.right, minNode.v)
	}

	return n
}

func findMin[T Ordered](n *node[T]) *node[T] {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (t *Tree[T]) Values() []T {
	result := []T{}
	inOrder(t.head, &result)
	return result
}

func inOrder[T Ordered](n *node[T], result *[]T) {
	if n != nil {
		inOrder(n.left, result)
		*result = append(*result, n.v)
		inOrder(n.right, result)
	}
}