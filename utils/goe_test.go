package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDistanceInKmBetweenEarthCoordinatesLessThan1km(t *testing.T) {
	latitude := -34.634162
	longitude := -58.439050

	latitude2 := -34.634621
	longitude2 := -58.436132

	result := DistanceInKmBetweenEarthCoordinates(latitude, longitude, latitude2, longitude2)

	assert.Equal(t, 0.27179160552228193, result)
}


func TestDistanceInKmBetweenEarthCoordinatesMoreThan1km(t *testing.T) {
	latitude := -34.634162
	longitude := -58.439050

	latitude2 := -34.622863
	longitude2 := -58.441961

	result := DistanceInKmBetweenEarthCoordinates(latitude, longitude, latitude2, longitude2)

	assert.Equal(t, 1.2842516181403352, result)
}
