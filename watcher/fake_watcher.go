package watcher

import (
	"path"

	"github.com/golang/glog"
)

// FakeWatcher implements an in-memory Watcher.
type FakeWatcher struct {
	watches map[string]bool
	events  chan Event
}

func NewFakeWatcher() *FakeWatcher {
	return &FakeWatcher{
		make(map[string]bool),
		make(chan Event)}
}

// Add adds a watch to the FakeWatcher
func (w *FakeWatcher) Add(name string) error {
	w.watches[name] = true
	return nil
}

// Close closes down the FakeWatcher
func (w *FakeWatcher) Close() error {
	close(w.events)
	return nil
}

// Remove removes a watch from the FakeWatcher
func (w *FakeWatcher) Remove(name string) error {
	delete(w.watches, name)
	return nil
}

// Events returns the channel of messages.
func (w *FakeWatcher) Events() <-chan Event { return w.events }

// InjectCreate lets a test inject a fake creation event.
func (w *FakeWatcher) InjectCreate(name string) {
	dirname := path.Dir(name)
	if w.watches[dirname] {
		w.events <- Event{name, Create}
	} else {
		glog.Infof("not watching %s to see %s", dirname, name)
	}
}

// InjectUpdate lets a test inject a fake update event.
func (w *FakeWatcher) InjectUpdate(name string) {
	if w.watches[name] {
		w.events <- Event{name, Update}
	} else {
		glog.Infof("not watching %s", name)
	}
}

// InjectDelete lets a test inject a fake deletion event.
func (w *FakeWatcher) InjectDelete(name string) {
	if w.watches[name] {
		w.events <- Event{name, Delete}
	} else {
		glog.Infof("not watching %s", name)
	}
}