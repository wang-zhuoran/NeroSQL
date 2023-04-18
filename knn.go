package nerosql

import (
	"fmt"
	"math"
	"sort"
)

// KNN算法
func knn(features [][]interface{}, label []int, k int, distanceMetric string) int {
	// 计算每个训练样本与测试样本的距离
	distances := make([]float64, len(label))
	for i := 0; i < len(label); i++ {
		distances[i] = distance(features[i], features[k], distanceMetric)
	}

	// 找出距离最近的k个样本
	indices := topk(distances, k)

	// 统计k个样本中最常见的标签
	count := make(map[int]int)
	for _, index := range indices {
		count[label[index]]++
	}

	// 返回最常见的标签
	maxCount := -1
	maxLabel := -1
	for l, c := range count {
		if c > maxCount {
			maxCount = c
			maxLabel = l
		}
	}
	return maxLabel
}

// 计算两个样本之间的距离
func distance(x []interface{}, y []interface{}, distanceMetric string) float64 {
	if distanceMetric == "euclidean" {
		// 计算欧式距离
		var d float64
		for i := 0; i < len(x); i++ {
			if x[i] == y[i] {
				continue
			}
			switch x[i].(type) {
			case float64:
				d += math.Pow(x[i].(float64)-y[i].(float64), 2)
			case string:
				d += math.Pow(float64(len([]rune(x[i].(string)))-len([]rune(y[i].(string)))), 2)
			default:
				panic(fmt.Sprintf("unsupported type %T", x[i]))
			}
		}
		return math.Sqrt(d)
	} else if distanceMetric == "manhattan" {
		// 计算曼哈顿距离
		var d float64
		for i := 0; i < len(x); i++ {
			if x[i] == y[i] {
				continue
			}
			switch x[i].(type) {
			case float64:
				d += math.Abs(x[i].(float64) - y[i].(float64))
			case string:
				d += float64(len([]rune(x[i].(string))) - len([]rune(y[i].(string))))
			default:
				panic(fmt.Sprintf("unsupported type %T", x[i]))
			}
		}
		return d
	} else {
		// 默认使用欧式距离
		var d float64
		for i := 0; i < len(x); i++ {
			if x[i] == y[i] {
				continue
			}
			switch x[i].(type) {
			case float64:
				d += math.Pow(x[i].(float64)-y[i].(float64), 2)
			case string:
				d += math.Pow(float64(len([]rune(x[i].(string)))-len([]rune(y[i].(string)))), 2)
			default:
				panic(fmt.Sprintf("unsupported type %T", x[i]))
			}
		}
		return math.Sqrt(d)
	}
}

// 返回前k个最小值的索引
func topk(arr []float64, k int) []int {
	indices := make([]int, len(arr))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool { return arr[indices[i]] < arr[indices[j]] })
	return indices[:k]
}
