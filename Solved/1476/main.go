package main

// SubrectangleQueries SubrectangleQueries
type SubrectangleQueries struct {
	row       []int
	axisValue [][]int
}

// rectangle = [][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}}

// Constructor Constructor
func Constructor(rectangle [][]int) SubrectangleQueries {
	var ans SubrectangleQueries
	length := len(rectangle)
	for i := 0; i < length; i++ {
		ans.row = append(ans.row, i)
		ans.axisValue = append(ans.axisValue, rectangle[i])
	}
	return ans
}

// UpdateSubrectangle UpdateSubrectangle
func (u *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {

	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			u.axisValue[i][j] = newValue
		}
	}
}

// GetValue GetValue
func (u *SubrectangleQueries) GetValue(row int, col int) int {
	ans := u.axisValue[row][col]
	return ans
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
