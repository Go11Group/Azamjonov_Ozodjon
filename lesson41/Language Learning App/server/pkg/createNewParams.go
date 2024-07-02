package pkg

import "fmt"

func CreateNewParams(param string, paramCount int) string {
	return param + fmt.Sprintf("$%d", paramCount)
}
