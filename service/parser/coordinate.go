package parser

import (
	"fmt"

	"github.com/go-spatial/proj"
)

func InverseCoordinates(x, y float64) (float64, float64, error) {
	var lonlat = []float64{x, y}

	xy, err := proj.Inverse(proj.EPSG3857, lonlat)
	if err != nil {
		return 0, 0, fmt.Errorf("proj inverse: %w", err)
	}

	return xy[0], xy[1], nil
}
