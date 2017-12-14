package main

import "math"

func pearsonCorrelation(xArr, yArr map[string]RecommenderItem) float64 {
	var sum_xy, sum_x, sum_y, sum_x2, sum_y2, n float64
	if len(xArr) < len(yArr) {
		n = float64(len(xArr))
	} else {
		n = float64(len(yArr))
		swArr := xArr
		xArr = yArr
		yArr = swArr
	}
	if n == 0 {
		return 0
	}
	for key, x := range xArr {
		y := yArr[key]
		sum_xy += x.rating * y.rating
		sum_x += x.rating
		sum_y += y.rating
		sum_x2 += math.Pow(x.rating, 2)
		sum_y2 += math.Pow(y.rating, 2)
	}

	denom_value := math.Sqrt((n*sum_x2 - math.Pow(sum_x, 2)) * (n*sum_y2 - math.Pow(sum_y, 2)))
	return ((n * sum_xy) - (sum_x * sum_y)) / denom_value
}
