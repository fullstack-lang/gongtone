// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongtone/go/models"
	"github.com/fullstack-lang/gongtone/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Note__dummysDeclaration__ models.Note
var __Note_time__dummyDeclaration time.Duration

var mutexNote sync.Mutex

// An NoteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNote updateNote deleteNote
type NoteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NoteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNote updateNote
type NoteInput struct {
	// The Note to submit or modify
	// in: body
	Note *orm.NoteAPI
}

// GetNotes
//
// swagger:route GET /notes notes getNotes
//
// # Get all notes
//
// Responses:
// default: genericError
//
//	200: noteDBResponse
func (controller *Controller) GetNotes(c *gin.Context) {

	// source slice
	var noteDBs []orm.NoteDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNotes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote.GetDB()

	query := db.Find(&noteDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	noteAPIs := make([]orm.NoteAPI, 0)

	// for each note, update fields from the database nullable fields
	for idx := range noteDBs {
		noteDB := &noteDBs[idx]
		_ = noteDB
		var noteAPI orm.NoteAPI

		// insertion point for updating fields
		noteAPI.ID = noteDB.ID
		noteDB.CopyBasicFieldsToNote_WOP(&noteAPI.Note_WOP)
		noteAPI.NotePointersEncoding = noteDB.NotePointersEncoding
		noteAPIs = append(noteAPIs, noteAPI)
	}

	c.JSON(http.StatusOK, noteAPIs)
}

// PostNote
//
// swagger:route POST /notes notes postNote
//
// Creates a note
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNote(c *gin.Context) {

	mutexNote.Lock()
	defer mutexNote.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNotes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote.GetDB()

	// Validate input
	var input orm.NoteAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create note
	noteDB := orm.NoteDB{}
	noteDB.NotePointersEncoding = input.NotePointersEncoding
	noteDB.CopyBasicFieldsFromNote_WOP(&input.Note_WOP)

	query := db.Create(&noteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNote.CheckoutPhaseOneInstance(&noteDB)
	note := backRepo.BackRepoNote.Map_NoteDBID_NotePtr[noteDB.ID]

	if note != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), note)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, noteDB)
}

// GetNote
//
// swagger:route GET /notes/{ID} notes getNote
//
// Gets the details for a note.
//
// Responses:
// default: genericError
//
//	200: noteDBResponse
func (controller *Controller) GetNote(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNote", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote.GetDB()

	// Get noteDB in DB
	var noteDB orm.NoteDB
	if err := db.First(&noteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var noteAPI orm.NoteAPI
	noteAPI.ID = noteDB.ID
	noteAPI.NotePointersEncoding = noteDB.NotePointersEncoding
	noteDB.CopyBasicFieldsToNote_WOP(&noteAPI.Note_WOP)

	c.JSON(http.StatusOK, noteAPI)
}

// UpdateNote
//
// swagger:route PATCH /notes/{ID} notes updateNote
//
// # Update a note
//
// Responses:
// default: genericError
//
//	200: noteDBResponse
func (controller *Controller) UpdateNote(c *gin.Context) {

	mutexNote.Lock()
	defer mutexNote.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNote", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote.GetDB()

	// Validate input
	var input orm.NoteAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var noteDB orm.NoteDB

	// fetch the note
	query := db.First(&noteDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	noteDB.CopyBasicFieldsFromNote_WOP(&input.Note_WOP)
	noteDB.NotePointersEncoding = input.NotePointersEncoding

	query = db.Model(&noteDB).Updates(noteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	noteNew := new(models.Note)
	noteDB.CopyBasicFieldsToNote(noteNew)

	// redeem pointers
	noteDB.DecodePointers(backRepo, noteNew)

	// get stage instance from DB instance, and call callback function
	noteOld := backRepo.BackRepoNote.Map_NoteDBID_NotePtr[noteDB.ID]
	if noteOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), noteOld, noteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the noteDB
	c.JSON(http.StatusOK, noteDB)
}

// DeleteNote
//
// swagger:route DELETE /notes/{ID} notes deleteNote
//
// # Delete a note
//
// default: genericError
//
//	200: noteDBResponse
func (controller *Controller) DeleteNote(c *gin.Context) {

	mutexNote.Lock()
	defer mutexNote.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNote", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote.GetDB()

	// Get model if exist
	var noteDB orm.NoteDB
	if err := db.First(&noteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&noteDB)

	// get an instance (not staged) from DB instance, and call callback function
	noteDeleted := new(models.Note)
	noteDB.CopyBasicFieldsToNote(noteDeleted)

	// get stage instance from DB instance, and call callback function
	noteStaged := backRepo.BackRepoNote.Map_NoteDBID_NotePtr[noteDB.ID]
	if noteStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), noteStaged, noteDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}