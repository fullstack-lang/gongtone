// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtone/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Tone":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tone Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ToneFormCallback(
			nil,
			probe,
			formGroup,
		)
		tone := new(models.Tone)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tone, formGroup, probe)
	}
	formStage.Commit()
}
