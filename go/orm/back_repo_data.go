// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ToneAPIs []*ToneAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, toneDB := range backRepo.BackRepoTone.Map_ToneDBID_ToneDB {

		var toneAPI ToneAPI
		toneAPI.ID = toneDB.ID
		toneAPI.TonePointersEncoding = toneDB.TonePointersEncoding
		toneDB.CopyBasicFieldsToTone_WOP(&toneAPI.Tone_WOP)

		backRepoData.ToneAPIs = append(backRepoData.ToneAPIs, &toneAPI)
	}

}
