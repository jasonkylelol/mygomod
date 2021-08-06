package mygomod

func validQueens(maps [][]rune, row, col int) bool {
	for _, cols := range maps {
		if cols[col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if maps[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len(maps); i, j = i-1, j+1 {
		if maps[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func getNQueens(resmaps *[][]string, maps *[][]rune, row int) {
	if len(*maps) == row {
		*resmaps = append(*resmaps, convMaps(*maps))
		return
	}
	columnSize := len((*maps)[row])
	for col := 0; col < columnSize; col++ {
		if !validQueens(*maps, row, col) {
			continue
		}
		(*maps)[row][col] = 'Q'
		getNQueens(resmaps, maps, row+1)
		(*maps)[row][col] = '.'
	}
}

func initMaps(n int) [][]rune {
	maps := [][]rune{}
	for i := 0; i < n; i++ {
		row := []rune{}
		for j := 0; j < n; j++ {
			row = append(row, '.')
		}
		maps = append(maps, row)
	}
	return maps
}

func convMaps(maps [][]rune) []string {
	resmaps := []string{}
	for _, row := range maps {
		resmaps = append(resmaps, string(row))
	}
	return resmaps
}

func solveNQueens(n int) [][]string {
	var resmaps = [][]string{}
	maps := initMaps(n)
	getNQueens(&resmaps, &maps, 0)
	return resmaps
}
