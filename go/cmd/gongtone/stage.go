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

	__Freqency__000000_A4 := (&models.Freqency{Name: `A4`}).Stage(stage)
	__Freqency__000001_E5 := (&models.Freqency{Name: `E5`}).Stage(stage)

	__Note__000000_A4_E5 := (&models.Note{Name: `A4+E5`}).Stage(stage)

	// Setup of values

	__Freqency__000000_A4.Name = `A4`

	__Freqency__000001_E5.Name = `E5`

	__Note__000000_A4_E5.Name = `A4+E5`
	__Note__000000_A4_E5.Start = 0.000000
	__Note__000000_A4_E5.Duration = 1.000000
	__Note__000000_A4_E5.Velocity = 1.000000

	// Setup of pointers
	__Note__000000_A4_E5.Frequencies = append(__Note__000000_A4_E5.Frequencies, __Freqency__000000_A4)
	__Note__000000_A4_E5.Frequencies = append(__Note__000000_A4_E5.Frequencies, __Freqency__000001_E5)
}
