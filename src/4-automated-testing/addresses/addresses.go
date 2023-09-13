package addresses

import "strings"

func ValidateAddressType(address string) string {
	validTypes := []string{"avenue", "street", "road", "highway"}

	lowerCaseAddress := strings.ToLower(address)
	splitAddress := strings.Split(lowerCaseAddress, " ")
	addressLastWord := splitAddress[len(splitAddress)-1]

	for _, addressType := range validTypes {
		if addressType == addressLastWord {
			return addressType
		}
	}

	return "Invalid address!"
}
