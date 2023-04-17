package workflow

func getStepThatRequires(requiredStep Step, dependencies map[Step][]Step) []Step {
	results := make([]Step, 0, len(dependencies))
	for step, stepDependencies := range dependencies {
		for _, dependency := range stepDependencies {
			if dependency.GetId() == requiredStep.GetId() {
				results = append(results, step)
				// Switch to next step
				break
			}
		}
	}
	return results
}

func areRequirementsFullFilled(step Step, dependencies map[Step][]Step) bool {
	for _, dep := range dependencies[step] {
		if dep.GetStatus() != SUCCESS {
			return false
		}
	}
	return true
}
