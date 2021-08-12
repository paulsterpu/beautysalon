package salon

func supportsAction(actions []Action, action Action) bool {
	for _, a := range actions {
		if a == action {
			return true
		}
	}

	return false
}

func getActionIndex(actions []Action, action Action) int {
	for i, a := range actions {
		if a == action {
			return i
		}
	}

	return -1
}