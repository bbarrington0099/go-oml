package gooml

type omlTracker struct {
	activeMarkups map[string]int
}

func newOMLTracker() *omlTracker {
	return &omlTracker{
		activeMarkups: make(map[string]int),
	}
}

func (t *omlTracker) addMarkup(markup string) {
	if _, exists := t.activeMarkups[markup]; !exists {
		t.activeMarkups[markup] = 0
	}
	t.activeMarkups[markup]++
}

func (t *omlTracker) removeMarkup(markup string) {
	if count, exists := t.activeMarkups[markup]; exists {
		if count > 1 {
			t.activeMarkups[markup]--
		} else {
			delete(t.activeMarkups, markup)
		}
	}
}

func (t *omlTracker) isMarkupActive(markup string) bool {
	_, exists := t.activeMarkups[markup]
	return exists
}

func (t *omlTracker) getActiveMarkups() ([]string, int) {
	activeMarkups := make([]string, 0, len(t.activeMarkups))
	for markup := range t.activeMarkups {
		activeMarkups = append(activeMarkups, markup)
	}
	return activeMarkups, len(activeMarkups)
}

func (t *omlTracker) clear() {
	t.activeMarkups = make(map[string]int)
}