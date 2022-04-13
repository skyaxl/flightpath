package flightpathservice

import "fmt"

type FlightPathService struct {
}

func New() *FlightPathService {
	return &FlightPathService{}
}

// getFlightMaps this method return a  map[string]string where key is origin and the value is a destination
func getFlightMaps(input [][]string) (forward map[string]string, reverse map[string]string, err error) {
	forward = make(map[string]string, len(input))
	reverse = make(map[string]string, len(input))

	for i, v := range input {
		if len(v) != 2 {
			return nil, nil, fmt.Errorf("invalid fligh index: %v, value: %v", i, v)
		}
		forward[v[0]] = v[1]
		reverse[v[1]] = v[0]
	}
	return forward, reverse, nil
}

// getLast get last airport
func getLast(key string, flights map[string]string) string {
	if _, ok := flights[key]; !ok {
		return key
	}
	return getLast(flights[key], flights)
}

func (fps *FlightPathService) CalculateFlighPath(input [][]string) (response []string, err error) {
	var (
		forwardFlights map[string]string
		reverse        map[string]string
	)

	if forwardFlights, reverse, err = getFlightMaps(input); err != nil {
		return nil, err
	}
	var firstAirport string
	var secondAirport string
	for k, v := range forwardFlights {
		if _, ok := reverse[k]; ok {
			continue
		}
		firstAirport = k
		secondAirport = v
		break
	}
	last := getLast(secondAirport, forwardFlights)

	return []string{firstAirport, last}, nil
}
