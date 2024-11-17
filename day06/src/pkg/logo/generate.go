package logo

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

const (
	width   = 300
	height  = 300
	centerX = width / 2
	centerY = height / 2
	radius  = 100
)

func GenerateLogo(filePath string) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	backgroundColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: backgroundColor}, image.Point{}, draw.Src)

	starColor := color.RGBA{R: 255, G: 215, B: 0, A: 255}

	// Расчет вершин звезды
	numPoints := 5
	points := make([][2]int, numPoints*2)

	for i := 0; i < numPoints*2; i++ {
		angle := math.Pi/2 + 2*math.Pi*float64(i)/(float64(numPoints)*2)
		r := radius
		if i%2 != 0 {
			r = radius / 2 // Сокращаем радиус для внутренних вершин
		}
		x := centerX + int(float64(r)*math.Cos(angle))
		y := centerY - int(float64(r)*math.Sin(angle))
		points[i] = [2]int{x, y}
	}

	// Соединение вершин для рисования звезды
	for i := 0; i < numPoints*2; i++ {
		next := (i + 1) % (numPoints * 2)
		drawLine(img, points[i][0], points[i][1], points[next][0], points[next][1], starColor)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// drawLine рисует линию на изображении (алгоритм Брезенхема)
func drawLine(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := -1
	if x1 < x2 {
		sx = 1
	}
	sy := -1
	if y1 < y2 {
		sy = 1
	}
	err := dx - dy

	for {
		img.Set(x1, y1, col)
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

// abs возвращает абсолютное значение числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
