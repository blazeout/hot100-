package _7_旋转图像

// 顺时针选择90度 : 先将矩阵转置, 再左右镜像
// 顺时针旋转180度 : 先将矩阵上下镜像, 再左右镜像
// 逆时针旋转90度 : 先将矩阵转置,再上下镜像

func rotate(matrix [][]int)  {
	// 转置矩阵 : 行变成列, 列变成行
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	left := 0
	// 代表列数
	right := len(matrix[0]) - 1
	for left < right {
		for i := 0; i < len(matrix); i++ {
			temp := matrix[i][left]
			matrix[i][left] = matrix[i][right]
			matrix[i][right] = temp
		}
		left++
		right--
	}

}
