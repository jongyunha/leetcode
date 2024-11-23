package main

const (
	stoneOfBox    = '#' // 35
	emptyOfBox    = '.' // 46
	obstacleOfBox = '*' // 42
)

func rotateTheBox(box [][]byte) [][]byte {
	rows := len(box)
	cols := len(box[0])
	rotateBox := make([][]byte, len(box[0]))

	for i := range rotateBox {
		rotateBox[i] = make([]byte, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotateBox[j][rows-1-i] = box[i][j]
		}
	}

	for row := 0; row < len(rotateBox[0]); row++ {
		for col := 0; col < len(rotateBox); col++ {
			empty, stone, obstacle := -1, -1, -1

			for i := col; i < len(rotateBox); i++ {
				curr := rotateBox[i][row]
				switch curr {
				case emptyOfBox:
					empty = i
				case stoneOfBox:
					if stone != -1 {
						stone = min(stone, i)
					} else {
						stone = i
					}
				case obstacleOfBox:
					obstacle = i
				}
				if obstacle != -1 {
					break
				}
			}

			if stone == -1 || empty == -1 {
				continue
			}

			if obstacle == -1 && stone < empty {
				tmp := rotateBox[empty][row]
				rotateBox[empty][row] = rotateBox[stone][row]
				rotateBox[stone][row] = tmp
			}

			if stone < obstacle && empty < obstacle && stone < empty {
				tmp := rotateBox[empty][row]
				rotateBox[empty][row] = rotateBox[stone][row]
				rotateBox[stone][row] = tmp
			}
		}
	}

	return rotateBox
}
