package SimpleFactory

import "testing"

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("s").GetFood()
	NewRestaurant("k").GetFood()
}
