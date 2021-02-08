package psql

import (
	"fmt"
	"strings"
)

const (
	openBracket = "("

	closeBracket = ")"
)

func convertPlaceholderInt(IDs []int) string {
	var conditions []string
	for _, id := range IDs {
		conditions = append(conditions, fmt.Sprintf("%d", id))
	}

	conditionS := strings.Join(conditions, ",")
	return fmt.Sprintf("%s%s%s", openBracket, conditionS, closeBracket)
}
