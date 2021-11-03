package main

import "fmt"

const order = 4

// DataItem 定义一个内部数据
type DataItem struct {
	val int
}

// 打印数据时候使用
func (i *DataItem) display() {
	fmt.Println(i.val)
}

// TreeNode node节点
type TreeNode234 struct {
	numItems   int
	parent     *TreeNode234
	childArray [order]*TreeNode234
	dataItems  [order - 1]*DataItem
}

// 连接一个子节点
func (r *TreeNode234) connectionChild(childNum int, child *TreeNode234) {
	r.childArray[childNum] = child
	if nil != child {
		child.parent = r
	}
}

// 删除一个子节点
func (r *TreeNode234) disConnectionChild(childNum int) *TreeNode234 {
	node := r.childArray[childNum]
	r.childArray[childNum] = nil
	return node
}

// 获取第n个位置的孩子
func (r *TreeNode234) getChild(childNum int) *TreeNode234 {
	return r.childArray[childNum]
}

// 获取父亲
func (r *TreeNode234) getParent() *TreeNode234 {
	return r.parent
}

// 是否为叶子节点
func (r *TreeNode234) isLeaf() bool {
	return r.childArray[0] == nil
}

// 是否满了
func (r *TreeNode234) isFull() bool {
	return r.numItems == (order - 1)
}

// 获取节点数量
func (r *TreeNode234) getNumItems() int {
	return r.numItems
}

// 获取节点数据
func (r *TreeNode234) getDataItem(index int) *DataItem {
	return r.dataItems[index]
}

// 查找元素
func (r *TreeNode234) findItem(key int) int {
	for i := 0; i < order-1; i++ {
		// 如果这里还没有填充跳过
		if r.dataItems[i] == nil {
			break
			// 如果数组中有一个值等于key,返回位置
		} else if r.dataItems[i].val == key {
			return i
		}
	}
	// 没有找到相等的
	return -1
}

func (r *TreeNode234) insertDataItem(newItem *DataItem) int {
	// 增加一个元素
	r.numItems++
	// 元素的key
	newVal := newItem.val
	// 从右边开始
	for j := order - 2; j >= 0; j-- {
		// 还没插入元素，因为这里如果是一个4节点的树，也就是0，1，2，3肯定会分裂，所以这里的树只可能到0，1，2
		if r.dataItems[j] == nil {
			continue
		} else {
			// 拿到最后当前数值
			v := r.dataItems[j].val
			// 还没找到正确位置，后面元素肯定比现在都大，向后移动
			if newVal < v {
				r.dataItems[j+1] = r.dataItems[j]
			} else {
				// 现在j+1已经被移动到了j+2
				// 现在j比当前元素小，插入到j+1
				r.dataItems[j+1] = newItem
				return j + 1
			}
		}
	}
	// 没找到大于newItem位置，则插入到0位置
	r.dataItems[0] = newItem
	return 0
}

// 删除最大元素
func (r *TreeNode234) removeDataItem() *DataItem {
	// 查找最大元素
	item := r.dataItems[r.numItems-1]
	// 删除引用
	r.dataItems[r.numItems-1] = nil
	// 元素个数减少
	r.numItems--
	// 返回被移除的元素
	return item
}

func (r *TreeNode234) displayNode() {
	for j := 0; j < r.numItems; j++ {
		r.dataItems[j].display()
		fmt.Println("/")
	}
}

type Tree234 struct {
	root *TreeNode234
}

func (t Tree234) find(key int) int {
	// 先拿到根节点
	currentNode := t.root
	// 先设置未找到
	childNumber := -1
	for true {
		// 从当前节点查找所有的item是否包含值
		childNumber = currentNode.findItem(key)
		// 找到了返回位置
		if childNumber != -1 {
			return childNumber
			// 如果是叶子节点，肯定不包含
		} else if currentNode.isLeaf() {
			return -1
		} else {
			// 找到下一个节点，然后在继续判断，也就是节点引用
			currentNode = t.getNextNode(currentNode, key)
		}
	}
	return childNumber
}

func (t Tree234) getNextNode(node *TreeNode234, key int) *TreeNode234 {
	i := 0
	// 所有的元素
	items := node.getNumItems()
	//  开始迭代，找到大于key的第一个节点，也就是最左边节点
	for i := 0; i < items; i++ {
		if key < node.dataItems[i].val {
			return node.getChild(i)
		}
	}
	// 走到这里说明全部走完，返回最右边的节点
	return node.getChild(i)
}

func newDataItem(key int) *DataItem {
	return &DataItem{
		val: key,
	}
}

func (t Tree234) insert(key int) {
	// 设置根节点
	currentNode := t.root
	// 创建一个数据节点
	item := newDataItem(key)
	for true {
		// 如果已经满了
		if currentNode.isFull() {
			// 分裂节点
			t.spilt(currentNode)
			// 找到父亲
			parent := currentNode.getParent()
			// 找到正确节点
			currentNode = t.getNextNode(parent, key)
		} else if currentNode.isLeaf() {
			// 叶子节点，跳出
			break
		} else {
			currentNode = t.getNextNode(currentNode, key)
		}
	}
	// 当前节点插入数据
	currentNode.insertDataItem(item)
}

func (t *Tree234) spilt(node *TreeNode234) {
	itemIndex := 0
	// 拿到最右边的数据
	itemC := node.removeDataItem()
	// 中间节点数据
	itemB := node.removeDataItem()
	// 右边数据的两个链接，这两个连接需要放到新node上
	child2 := node.disConnectionChild(2)
	child3 := node.disConnectionChild(3)
	// 创建新的节点
	newRight := newTreeNode()
	parent := node
	if node == t.root {
		// 创建新root
		t.root = newTreeNode()
		parent = t.root
		// 当前节点放到最左节点上，根节点与普通节点区别就是root设置和需要将当前节点放到最左边
		t.root.connectionChild(0, node)
	} else {
		parent = node.getParent()
	}
	// 中间数据上浮上去，这里itemIndex可能等于1或者2
	itemIndex = parent.insertDataItem(itemB)
	// 总元素数量
	items := parent.getNumItems()
	// 如果新插入的元素插入到了中间位置
	for j := items - 1; j > itemIndex; j-- {
		// 后面位置链接的数据需要向后移动
		temp := parent.disConnectionChild(j)
		parent.connectionChild(j+1, temp)
	}
	// 新节点，在其父亲的itemIndex+1处
	parent.connectionChild(itemIndex+1, newRight)
	// 新节点增加数据
	newRight.insertDataItem(itemC)
	// 连接数据
	newRight.connectionChild(0, child2)
	newRight.connectionChild(1, child3)
}

func newTreeNode() *TreeNode234 {
	return &TreeNode234{numItems: 0, parent: nil}
}

func main() {
	tree234 := Tree234{
		root: newTreeNode(),
	}

	tree234.insert(50)
	tree234.insert(40)
	tree234.insert(60)
	tree234.insert(30)
	tree234.insert(70)
}
