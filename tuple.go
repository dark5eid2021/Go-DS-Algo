package main

func powerSeries(a int) (int, int) {
	return a * a, a * a * a

}

func powerSeriesN(a int) (square int, cube int) {
	square = a * a

	cube = square * a

	return square, cube, nil

}
