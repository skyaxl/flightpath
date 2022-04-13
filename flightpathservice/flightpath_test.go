package flightpathservice_test

import (
	"errors"
	"testing"

	"github.com/skyaxl/flightpath/flightpathservice"
	"github.com/stretchr/testify/assert"
)

func TestFlightPathService_GetAiportsPath(t *testing.T) {

	tests := []struct {
		name         string
		input        [][]string
		wantResponse []string
		wamtError    error
	}{
		{
			"Ok with one flight",
			[][]string{{"SFO", "EWR"}},
			[]string{"SFO", "EWR"},
			nil,
		},
		{
			"Ok with two flights",
			[][]string{{"ATL", "EWR"}, {"SFO", "ATL"}},
			[]string{"SFO", "EWR"},
			nil,
		},
		{
			"Ok with four flights",
			[][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			[]string{"SFO", "EWR"},
			nil,
		},
		{
			"N Ok with four flights",
			[][]string{{"IND"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			nil,
			errors.New(`invalid fligh index: 0, value: [IND]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fps := flightpathservice.New()
			gotResponse, err := fps.CalculateFlighPath(tt.input)
			assert.Equal(t, tt.wamtError, err)
			assert.Equal(t, tt.wantResponse, gotResponse)

		})
	}
}
