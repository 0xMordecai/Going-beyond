package conv

import "strconv"

func main() {
}
func conv(str string) (num int64, err error) {
	//num, err = strconv.ParseInt(str, 10, 64)
	num, err = strconv.ParseInt(str, 2, 64)
	return
}
