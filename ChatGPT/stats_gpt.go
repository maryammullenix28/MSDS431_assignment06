package main

import (
	"fmt"
	"math"
	"time"
)

func mean(data []float64) float64 {
	total := 0.0
	for _, val := range data {
		total += val
	}
	return total / float64(len(data))
}

func sumSquares(data []float64) float64 {
	m := mean(data)
	total := 0.0
	for _, val := range data {
		total += math.Pow(val-m, 2)
	}
	return total
}

func linearRegression(x []float64, y []float64) (yIntercept float64, slope float64, rSquared float64, adjustedRSquared float64, fStat float64, prob float64) {
	if len(x) != len(y) {
		panic("Data sets must have the same number of observations.")
	}

	n := float64(len(x))
	sumX := mean(x)
	sumY := mean(y)
	sumX2 := sumSquares(x)
	sumY2 := sumSquares(y)
	sumXY := 0.0
	for i := 0; i < len(x); i++ {
		sumXY += x[i] * y[i]
	}
	covXY := sumXY - n*sumX*sumY
	varXY := (sumX2 - n*math.Pow(sumX, 2)) * (sumY2 - n*math.Pow(sumY, 2))
	varXY = math.Sqrt(varXY)
	r := covXY / varXY

	slope = r * (sumY2 - n*math.Pow(sumY, 2)) / (sumX2 - n*math.Pow(sumX, 2))
	yIntercept = sumY - slope*sumX

	ssRes := 0.0
	for i := 0; i < len(x); i++ {
		ssRes += math.Pow(y[i]-yIntercept-slope*x[i], 2)
	}
	ssTot := 0.0
	for i := 0; i < len(x); i++ {
		ssTot += math.Pow(y[i]-sumY, 2)
	}
	rSquared = 1 - ssRes/ssTot
	adjustedRSquared = 1 - (1-rSquared)*(n-1)/(n-2)

	fStat = (rSquared / (1 - rSquared)) * ((n - 2) / 1)
	prob = 1 - math.Pow(math.Pow(1-rSquared, n-2)*(1-rSquared+(fStat/(n-2))), n-2)

	return yIntercept, slope, rSquared, adjustedRSquared, fStat, prob
}

func main() {
	variable := "Dep. Variable"
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	x2 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y2 := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	x3 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y3 := []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
	x4 := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y4 := []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

	yIntercept1, slope1, rSquared1, adjustedRSquared1, fStat1, prob1 := linearRegression(x1, y1)
	yIntercept2, slope2, rSquared2, adjustedRSquared2, fStat2, prob2 := linearRegression(x2, y2)
	yIntercept3, slope3, rSquared3, adjustedRSquared3, fStat3, prob3 := linearRegression(x3, y3)
	yIntercept4, slope4, rSquared4, adjustedRSquared4, fStat4, prob4 := linearRegression(x4, y4)

	fmt.Println("\t\t\tLinear Regression Results")
	fmt.Println("===============================================================================")
	fmt.Printf("Dep. Variable:               %s\t\t\tR-squared:\t\t%.3f\n", variable, rSquared1)
	fmt.Printf("Date:                        %s\tAdjusted R-squared:\t%.3f\n", time.Now().Format("Mon, 02 Jan 2006"), adjustedRSquared1)
	fmt.Printf("Time:                        %s\t\tF-Statistic:\t\t%.2f\n", time.Now().Format("15:04:05"), fStat1)
	fmt.Printf("Number of Observations:      %d\t\t\tProb (F-Statistic):\t%.5f\n", len(x1), prob1)
	fmt.Printf("y-intercept/const            %.4f\t\tSlope:\t\t\t%.4f\n", yIntercept1, slope1)
	fmt.Println("===============================================================================")
	fmt.Println()

	fmt.Println("\t\t\tLinear Regression Results")
	fmt.Println("===============================================================================")
	fmt.Printf("Dep. Variable:               %s\t\t\tR-squared:\t\t%.3f\n", variable, rSquared2)
	fmt.Printf("Date:                        %s\tAdjusted R-squared:\t%.3f\n", time.Now().Format("Mon, 02 Jan 2006"), adjustedRSquared2)
	fmt.Printf("Time:                        %s\t\tF-Statistic:\t\t%.2f\n", time.Now().Format("15:04:05"), fStat2)
	fmt.Printf("Number of Observations:      %d\t\t\tProb (F-Statistic):\t%.5f\n", len(x2), prob2)
	fmt.Printf("y-intercept/const            %.4f\t\tSlope:\t\t\t%.4f\n", yIntercept2, slope2)
	fmt.Println("===============================================================================")
	fmt.Println()

	fmt.Println("\t\t\tLinear Regression Results")
	fmt.Println("===============================================================================")
	fmt.Printf("Dep. Variable:               %s\t\t\tR-squared:\t\t%.3f\n", variable, rSquared3)
	fmt.Printf("Date:                        %s\tAdjusted R-squared:\t%.3f\n", time.Now().Format("Mon, 02 Jan 2006"), adjustedRSquared3)
	fmt.Printf("Time:                        %s\t\tF-Statistic:\t\t%.2f\n", time.Now().Format("15:04:05"), fStat3)
	fmt.Printf("Number of Observations:      %d\t\t\tProb (F-Statistic):\t%.5f\n", len(x3), prob3)
	fmt.Printf("y-intercept/const            %.4f\t\tSlope:\t\t\t%.4f\n", yIntercept3, slope3)
	fmt.Println("===============================================================================")
	fmt.Println()

	fmt.Println("\t\t\tLinear Regression Results")
	fmt.Println("===============================================================================")
	fmt.Printf("Dep. Variable:               %s\t\t\tR-squared:\t\t%.3f\n", variable, rSquared4)
	fmt.Printf("Date:                        %s\tAdjusted R-squared:\t%.3f\n", time.Now().Format("Mon, 02 Jan 2006"), adjustedRSquared4)
	fmt.Printf("Time:                        %s\t\tF-Statistic:\t\t%.2f\n", time.Now().Format("15:04:05"), fStat4)
	fmt.Printf("Number of Observations:      %d\t\t\tProb (F-Statistic):\t%.5f\n", len(x4), prob4)
	fmt.Printf("y-intercept/const            %.4f\t\tSlope:\t\t\t%.4f\n", yIntercept4, slope4)
	fmt.Println("===============================================================================")
	fmt.Println()
}
