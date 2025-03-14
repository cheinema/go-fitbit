package fitbit

import "strings"

// putDateTimeParam puts date and time information to a map of query parameters.
// Argument dt must be in one of the following formats: yyyy-MM-dd or yyyy-MM-ddTHH:mm:ss.
// The shorter format leads to a parameter with key "date" and this string.
// The longer format leads to two parameters with keys "date" and "time",
// each with the parts before and after the separator 'T'.
func putDateTimeParam(params map[string]string, dt string) {
	parts := strings.Split(dt, "T")
	params["date"] = parts[0]
	if len(parts) > 1 {
		params["time"] = parts[1]
	}
}
