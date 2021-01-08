package hashtable

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type hashFn func(key int, size int) int

type HashTable struct {
	slots []*Node
	hash  hashFn
}

type optFn func(*opt)

type opt struct {
	hash hashFn
	size int
}

func WithHashFn(hash hashFn) optFn {
	return func(o *opt) {
		o.hash = hash
	}
}

func WithSize(size int) optFn {
	return func(o *opt) {
		o.size = size
	}
}

var defaultHashFn = func(key int, size int) int {
	return key % size
}

const (
	defaultSize = 10
)

func New(options ...optFn) *HashTable {
	opt := &opt{
		size: defaultSize,
		hash: defaultHashFn,
	}
	for _, option := range options {
		option(opt)
	}
	table := &HashTable{
		hash:  opt.hash,
		slots: make([]*Node, opt.size),
	}
	return table
}

func (ht *HashTable) Put(key int, value int) {
	position := ht.hash(key, len(ht.slots))
	head := ht.slots[position]
	newNode := &Node{Key: key, Value: value}
	if head == nil {
		ht.slots[position] = newNode
		return
	}
	if head.Value == value {
		return
	}
	tmp := head
	for tmp.Next != nil {
		if tmp.Next.Value == value {
			return
		}
		tmp = tmp.Next
	}
	tmp.Next = newNode
}

func (ht *HashTable) Get(key int) (int, bool) {
	position := ht.hash(key, len(ht.slots))
	head := ht.slots[position]
	if head == nil {
		return 0, false
	}
	tmp := head
	for tmp != nil && tmp.Key != key {
		tmp = tmp.Next
	}
	if tmp == nil {
		return 0, false
	}
	return tmp.Value, true
}