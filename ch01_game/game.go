package game

import "strings"

func ListItems(items []string) string {
	// result := "You can see here "
	// if len(items) == 0 {
	// 	return ""
	// }
	// if len(items) == 1 {
	// 	return "You can see " + items[0] + " here."
	// }
	// if len(items) < 3 {
	// 	return result + items[0] + " and " + items[1] + "."
	// }
	// result += strings.Join(items[:len(items)-1], ", ")
	// result += ", and "
	// result += items[len(items)-1]
	// result += "."
	// return result

	switch len(items) {
	case 0:
		return ""
	case 1:
		return "You can see " + items[0] + " here."
	case 2:
		return "You can see here " + items[0] + " and " + items[1] + "."
	default:
		return "You can see here " +
			strings.Join(items[:len(items)-1], ", ") +
			", and " + items[len(items)-1] + "."
	}
}
