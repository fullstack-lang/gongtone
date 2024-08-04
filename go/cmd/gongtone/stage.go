package main

import (
	"time"

	"github.com/fullstack-lang/gongtone/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Freqency__000000_Bb4 := (&models.Freqency{Name: `Bb4`}).Stage(stage)
	__Freqency__000001_Eb4 := (&models.Freqency{Name: `Eb4`}).Stage(stage)
	__Freqency__000002_G4 := (&models.Freqency{Name: `G4`}).Stage(stage)

	__Note__000000_Sampler := (&models.Note{Name: `Sampler`}).Stage(stage)

	// Setup of values

	__Freqency__000000_Bb4.Name = `Bb4`

	__Freqency__000001_Eb4.Name = `Eb4`

	__Freqency__000002_G4.Name = `G4`

	__Note__000000_Sampler.Name = `Sampler`
	__Note__000000_Sampler.Start = 0.000000
	__Note__000000_Sampler.Duration = 4.000000
	__Note__000000_Sampler.Velocity = 1.000000
	__Note__000000_Sampler.Info = ``

	// Setup of pointers
	__Note__000000_Sampler.Frequencies = append(__Note__000000_Sampler.Frequencies, __Freqency__000001_Eb4)
	__Note__000000_Sampler.Frequencies = append(__Note__000000_Sampler.Frequencies, __Freqency__000002_G4)
	__Note__000000_Sampler.Frequencies = append(__Note__000000_Sampler.Frequencies, __Freqency__000000_Bb4)
}
