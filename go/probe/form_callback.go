// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtone/go/models"
	"github.com/fullstack-lang/gongtone/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__ToneFormCallback(
	tone *models.Tone,
	probe *Probe,
	formGroup *table.FormGroup,
) (toneFormCallback *ToneFormCallback) {
	toneFormCallback = new(ToneFormCallback)
	toneFormCallback.probe = probe
	toneFormCallback.tone = tone
	toneFormCallback.formGroup = formGroup

	toneFormCallback.CreationMode = (tone == nil)

	return
}

type ToneFormCallback struct {
	tone *models.Tone

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (toneFormCallback *ToneFormCallback) OnSave() {

	log.Println("ToneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	toneFormCallback.probe.formStage.Checkout()

	if toneFormCallback.tone == nil {
		toneFormCallback.tone = new(models.Tone).Stage(toneFormCallback.probe.stageOfInterest)
	}
	tone_ := toneFormCallback.tone
	_ = tone_

	for _, formDiv := range toneFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tone_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if toneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tone_.Unstage(toneFormCallback.probe.stageOfInterest)
	}

	toneFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tone](
		toneFormCallback.probe,
	)
	toneFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if toneFormCallback.CreationMode || toneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		toneFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(toneFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ToneFormCallback(
			nil,
			toneFormCallback.probe,
			newFormGroup,
		)
		tone := new(models.Tone)
		FillUpForm(tone, newFormGroup, toneFormCallback.probe)
		toneFormCallback.probe.formStage.Commit()
	}

	fillUpTree(toneFormCallback.probe)
}
