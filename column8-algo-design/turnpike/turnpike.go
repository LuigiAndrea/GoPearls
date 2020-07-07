package turnpike

import "github.com/LuigiAndrea/GoPearls/utilities"

type turnpike []float64

func (tp *turnpike) getCostBetweenStation(fromStation, toStation int) (cost float64) {
	size := len(*tp)
	if tp == nil || size < 2 || fromStation > size || toStation > size || fromStation < 0 || toStation < 0 || fromStation == toStation {
		return -1
	}

	costOfTravel := utilities.CalculateCummulativeArray(*tp)

	if fromStation == 0 {
		cost = costOfTravel[toStation-1].Value
	} else if toStation == 0 {
		cost = costOfTravel[fromStation-1].Value
	} else if fromStation > toStation {
		cost = costOfTravel.GetDistance(toStation-1, fromStation-1)
	} else {
		cost = costOfTravel.GetDistance(fromStation-1, toStation-1)
	}

	return
}

//Assume costOfTravel has positive values
func newTurnpike(costOfTravel []float64) turnpike {
	return turnpike(costOfTravel)
}
