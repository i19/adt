package own

import "fmt"

type hashMapChaining struct {
	size        int
	cap         int
	loadThres   float64
	extendRatio int
	buckets     [][]hashPair
}

func newHashMapChaining() *hashMapChaining {
	bucketCount := 4
	buckets := make([][]hashPair, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = []hashPair{}
	}

	return &hashMapChaining{
		size:        0,
		cap:         bucketCount,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     buckets,
	}
}

func (m *hashMapChaining) hashFunc(key int) int {
	return key % m.cap
}

func (m *hashMapChaining) loadFactor() float64 {
	return float64(m.size / m.cap)
}

func (m *hashMapChaining) get(key int) string {
	idx := m.hashFunc(key)
	bucket := m.buckets[idx]
	for _, p := range bucket {
		if p.key == key {
			return p.val
		}
	}

	return ""
}

func (m *hashMapChaining) put(key int, val string) {
	if m.loadFactor() > m.loadThres {
		m.extend()
	}

	idx := m.hashFunc(key)
	for _, p := range m.buckets[idx] {
		if p.key == key {
			return
		}
	}
	m.buckets[idx] = append(m.buckets[idx], hashPair{key: key, val: val})
	m.size++
}

func (m *hashMapChaining) remove(key int) {
	idx := m.hashFunc(key)
	for i, p := range m.buckets[idx] {
		if p.key == key {
			m.buckets[idx] = append(m.buckets[idx][:i], m.buckets[idx][i+1:]...)
			m.size--
			break
		}
	}
}

func (m *hashMapChaining) extend() {
	tmpBuckets := make([][]hashPair, len(m.buckets))
	for i := 0; i < len(m.buckets); i++ {
		tmpBuckets[i] = make([]hashPair, len(m.buckets[i]))
		copy(tmpBuckets[i], m.buckets[i])
	}

	m.cap *= m.extendRatio
	m.buckets = make([][]hashPair, m.cap)
	for i := 0; i < m.cap; i++ {
		m.buckets[i] = make([]hashPair, 0)
	}
	m.size = 0
	for _, bucket := range tmpBuckets {
		for _, p := range bucket {
			m.put(p.key, p.val)
		}
	}
}

func (m *hashMapChaining) print() {
	for i, bucket := range m.buckets {
		fmt.Printf("bucket [%d] :\n", i)
		for _, p := range bucket {
			fmt.Printf("\tkey %d -> %s\n", p.key, p.val)
		}
	}
}
