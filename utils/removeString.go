package utils

func RemoveStrings(slice []string, targets ...string) []string {
	// Convert targets to a set (map) for fast lookup
	targetSet := make(map[string]struct{}, len(targets))
	for _, target := range targets {
		targetSet[target] = struct{}{}
	}

	var result []string
	for _, str := range slice {
		if _, found := targetSet[str]; !found {
			result = append(result, str) // Keep only if NOT in targetSet
		}
	}
	return result
}
