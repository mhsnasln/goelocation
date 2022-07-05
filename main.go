package goelocation

import "math"

// WIP
func Distance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) (float64, error) {
	rad := math.Pi / 180
	theta := lon1 - lon2
	distance := math.Sin(lat1*rad)*math.Sin(lat2*rad) + math.Cos(lat1*rad)*math.Cos(lat2*rad)*math.Cos(theta*rad)
	return (math.Acos(distance) / rad * 60 * 1.853) * 1000, nil
}
