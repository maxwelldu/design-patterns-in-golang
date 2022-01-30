package FactoryMethod

import "testing"

func TestFactoryMethod(t *testing.T) {
	kfcFactory := KFCFactory{}
	kfcFactory.NewRestaurant().GetFood()

	subwayFactory := SubWayFactory{}
	subwayFactory.NewRestaurant().GetFood()
}