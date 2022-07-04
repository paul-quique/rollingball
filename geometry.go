package rollingball

import (
	"math"

	"github.com/google/uuid"
)

func Dist(p1, p2 *Point) float64 {
	x1, x2, y1, y2 := p1.X, p2.X, p1.Y, p2.Y
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func RandomID() string {
	return uuid.New().String()
}
