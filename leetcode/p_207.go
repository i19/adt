package leetcode

// 207. 课程表
// https://leetcode.cn/problems/course-schedule
// 卡恩算法 Kahnn 算法
// https://www.bilibili.com/video/BV1pV4y1K75T
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	m := make(map[int][]int, 0)
	for _, v := range prerequisites {
		if _, ok := m[v[0]]; !ok {
			m[v[0]] = []int{}
		}
		m[v[0]] = append(m[v[0]], v[1])
		inDegree[v[1]]++
	}

	var res []int
	for {
		cousrse := -1
		for i, v := range inDegree {
			if v == 0 {
				cousrse = i
				break
			}
		}

		if cousrse != -1 {
			inDegree[cousrse] = -1
			for _, v := range m[cousrse] {
				inDegree[v]--
			}
			res = append(res, cousrse)
		} else {
			if len(res) == numCourses {
				return true
			} else {
				return false
			}
		}
	}
}
