package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_time_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteTime process Invert Graphite Unit tests
func TestGraphiteTime(t *testing.T) {
	RunTest(t, graphiteGraphiteTime, "")
}

var graphiteGraphiteTime = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "time",
				Arguments:  []string{"The.time.series"},
				Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: "CLEAR",
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "The.time.series",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 1548168444.000000}, {1548168504507231.000000, 1548168504.000000}, {1548168564507231.000000, 1548168564.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
