package fn

import (
	"corsanhub.com/lisg/corsan/core"
	"corsanhub.com/lisg/corsan/logging"
)

var log = logging.Logger{Name: "fn.math"}

func Sum(nums ...core.MalType) interface{} {
	log.Debug("############## Adding some numbers ...")

	isFloatingOp := false
	for _, num := range nums {
		mType := core.GetType(num)
		if mType == "core.MalFloat" {
			isFloatingOp = true
			break
		}
	}

	if isFloatingOp {
		var total float64 = 0
		for _, num := range nums {
			dd := num.Value()
			total += dd.(float64)
		}
		return total
	} else {
		var total int64 = 0
		for _, num := range nums {
			dd := num.Value()
			total += dd.(int64)
		}
		return total
	}

}
