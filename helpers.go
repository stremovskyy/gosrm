package gosrm

import (
	"net/url"
	"strconv"
	"strings"

	geo "github.com/paulmach/go.geo"
)

func StringRef(s string) *string {
	return &s
}

// addOptionalBoolParam adds a boolean parameter to the URL values if it's not nil.
func addOptionalBoolParam(values url.Values, key string, param *bool) {
	values.Add(key, strconv.FormatBool(param != nil && *param))
}

// addOptionalStringParam adds a string parameter to the URL values if it's not nil.
func addOptionalStringParam(values url.Values, key string, param *string) {
	if param != nil {
		values.Add(key, *param)
	}
}

func formatCoordinate(coord geo.Point) string {
	return strconv.FormatFloat(coord.Lng(), 'f', 9, 64) + "," + strconv.FormatFloat(coord.Lat(), 'f', 9, 64)
}

func addOptionalIndexListParam(values url.Values, key string, list *[]int) {
	if list != nil && len(*list) > 0 {
		var indexList []string
		for _, index := range *list {
			indexList = append(indexList, strconv.Itoa(index))
		}

		values.Add(key, strings.Join(indexList, ";"))
	}
}

func addOptionalFloatParam(values url.Values, key string, param *float64) {
	if param != nil {
		values.Add(key, strconv.FormatFloat(*param, 'f', 4, 64))
	}
}

func BoolRef(val bool) *bool {
	return &val
}
