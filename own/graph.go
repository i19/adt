package own

// 邻接表 Graph
type GraphList struct {
	vertex []int
	edge   map[int][]int
}

func (g *GraphList) bfs() []int {
	var res []int
	isPrinted := make(map[int]bool)
	for _, v := range g.vertex {
		isPrinted[v] = false
	}

	queue := []int{g.vertex[0]}
	isPrinted[g.vertex[0]] = true
	for len(queue) != 0 {
		vertex := queue[0]
		for _, v := range g.edge[vertex] {
			if !isPrinted[v] {
				isPrinted[v] = true
				queue = append(queue, v)
			}
		}
		queue = queue[1:]
		res = append(res, vertex)
	}
	return res
}

// 用循环代替递归
// 递归是栈结构，后进先出，如果使用循环，则需要一个 栈 数据结构，可以使用数组实现
func (g *GraphList) dfs() []int {
	var res []int
	isPrinted := make(map[int]bool)
	for _, v := range g.vertex {
		isPrinted[v] = false
	}

	stack := []int{g.vertex[0]}
	isPrinted[g.vertex[0]] = true
	for len(stack) != 0 {
		v := stack[len(stack)-1]
		// 栈特性，后进先出
		stack = stack[:len(stack)-1]
		res = append(res, v)
		for _, connectVertex := range g.edge[v] {
			if !isPrinted[connectVertex] {
				// 栈特性，后进先出
				isPrinted[connectVertex] = true
				stack = append(stack, connectVertex)
			}
		}
	}

	return res

}
