package gooml

import (
	"fmt"
)

type omlTracker struct {
	activeOmls map[*oml]bool
}

func newOMLTracker() *omlTracker {
	return &omlTracker{
		activeOmls: make(map[*oml]bool),
	}
}

func (t *omlTracker) addActiveOml(activeOml *oml) (err error) {
	if active, exists := t.activeOmls[activeOml]; !exists {
		t.activeOmls[activeOml] = true
	} else if active {
		err = fmt.Errorf("oml %s is already active", activeOml)
	}
	return
}

func (t *omlTracker) closeActiveOml(activeOml *oml) (err error) {
	if _, exists := t.activeOmls[activeOml]; !exists {
		err = fmt.Errorf("oml %s is not active", activeOml)
	} else {
		t.activeOmls[activeOml] = false
	}
	return
}

func (t *omlTracker) removeActiveOml(activeOml *oml) {
	if _, exists := t.activeOmls[activeOml]; exists {
		delete(t.activeOmls, activeOml)
	}
}

func (t *omlTracker) isActiveOml(activeOml *oml) bool {
	active, exists := t.activeOmls[activeOml]
	return exists && active
}

func (t *omlTracker) getActiveOmls() ([]*oml, int) {
	activeOmlCount := len(t.activeOmls)
	if activeOmlCount == 0 {
		return nil, 0
	}

	activeOmls := make([]*oml, 0, activeOmlCount)
	for activeOml := range t.activeOmls {
		activeOmls = append(activeOmls, activeOml)
	}
	return activeOmls, activeOmlCount
}

func (t *omlTracker) clear() {
	t.activeOmls = make(map[*oml]bool)
}

func (t *omlTracker) release() {
	for activeOml := range t.activeOmls {
		t.removeActiveOml(activeOml)
	}
	t = nil
}