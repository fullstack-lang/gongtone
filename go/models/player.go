package models

// Player is the singloton object to get the callbacks on play/paused
type Player struct {
	Name string

	Status Status
}
