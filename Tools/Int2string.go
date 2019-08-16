package Tools

import "strconv"

func Int2String(data interface{}) string {

	var re string
	switch data.(type) {

	case int64:
		re = strconv.FormatInt(data.(int64), 10)
		break

	case string:
		re = `"` + data.(string) + `"`
		break

	case int:
		re = strconv.Itoa(data.(int))
		break

	case float64:
		re = strconv.FormatFloat(data.(float64), 'E', -1, 64)
	}

	return re

}
