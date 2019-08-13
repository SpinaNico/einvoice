package share

import "fmt"

// IotaString2zerofill Generate number stirn with suffix, example:
// IotaString("MP", 1, 10)
// result:
// MP01
// MP02
// MP03
func IotaString2zerofill(suffix string, init int, end int) map[string]bool {
	response := make(map[string]bool)
	for i := init; i < end; i++ {
		response[fmt.Sprintf("%s%02d", suffix, i)] = true
	}
	return response
}
