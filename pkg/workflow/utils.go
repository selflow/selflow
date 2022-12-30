package workflow

import "errors"

func createErrorList(size int) []error {
	return make([]error, 0, size)
}

func joinErrorList(errorLst []error) error {
	acc := ""
	for _, err := range errorLst {
		if err != nil {
			acc += err.Error() + " ; "
		}
	}
	return errors.New(acc)
}

func appendErrorList(errorLst []error, err error) []error {
	if err != nil {
		errorLst = append(errorLst, err)
	}
	return errorLst
}

func mergeStringStringStringMaps(destination map[string]map[string]string, maps ...map[string]map[string]string) map[string]map[string]string {
	for _, m := range maps {
		for key, value := range m {
			destination[key] = value
		}
	}
	return destination
}
