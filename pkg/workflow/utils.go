package workflow

func mergeStringStringStringMaps(destination map[string]map[string]string, maps ...map[string]map[string]string) map[string]map[string]string {
	for _, m := range maps {
		for key, value := range m {
			destination[key] = value
		}
	}
	return destination
}
