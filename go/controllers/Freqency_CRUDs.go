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
var __Freqency__dummysDeclaration__ models.Freqency
var __Freqency_time__dummyDeclaration time.Duration

var mutexFreqency sync.Mutex

// An FreqencyID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFreqency updateFreqency deleteFreqency
type FreqencyID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FreqencyInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFreqency updateFreqency
type FreqencyInput struct {
	// The Freqency to submit or modify
	// in: body
	Freqency *orm.FreqencyAPI
}

// GetFreqencys
//
// swagger:route GET /freqencys freqencys getFreqencys
//
// # Get all freqencys
//
// Responses:
// default: genericError
//
//	200: freqencyDBResponse
func (controller *Controller) GetFreqencys(c *gin.Context) {

	// source slice
	var freqencyDBs []orm.FreqencyDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFreqencys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	query := db.Find(&freqencyDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	freqencyAPIs := make([]orm.FreqencyAPI, 0)

	// for each freqency, update fields from the database nullable fields
	for idx := range freqencyDBs {
		freqencyDB := &freqencyDBs[idx]
		_ = freqencyDB
		var freqencyAPI orm.FreqencyAPI

		// insertion point for updating fields
		freqencyAPI.ID = freqencyDB.ID
		freqencyDB.CopyBasicFieldsToFreqency_WOP(&freqencyAPI.Freqency_WOP)
		freqencyAPI.FreqencyPointersEncoding = freqencyDB.FreqencyPointersEncoding
		freqencyAPIs = append(freqencyAPIs, freqencyAPI)
	}

	c.JSON(http.StatusOK, freqencyAPIs)
}

// PostFreqency
//
// swagger:route POST /freqencys freqencys postFreqency
//
// Creates a freqency
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFreqency(c *gin.Context) {

	mutexFreqency.Lock()
	defer mutexFreqency.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFreqencys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	// Validate input
	var input orm.FreqencyAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create freqency
	freqencyDB := orm.FreqencyDB{}
	freqencyDB.FreqencyPointersEncoding = input.FreqencyPointersEncoding
	freqencyDB.CopyBasicFieldsFromFreqency_WOP(&input.Freqency_WOP)

	query := db.Create(&freqencyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFreqency.CheckoutPhaseOneInstance(&freqencyDB)
	freqency := backRepo.BackRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]

	if freqency != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), freqency)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, freqencyDB)
}

// GetFreqency
//
// swagger:route GET /freqencys/{ID} freqencys getFreqency
//
// Gets the details for a freqency.
//
// Responses:
// default: genericError
//
//	200: freqencyDBResponse
func (controller *Controller) GetFreqency(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFreqency", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	// Get freqencyDB in DB
	var freqencyDB orm.FreqencyDB
	if err := db.First(&freqencyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var freqencyAPI orm.FreqencyAPI
	freqencyAPI.ID = freqencyDB.ID
	freqencyAPI.FreqencyPointersEncoding = freqencyDB.FreqencyPointersEncoding
	freqencyDB.CopyBasicFieldsToFreqency_WOP(&freqencyAPI.Freqency_WOP)

	c.JSON(http.StatusOK, freqencyAPI)
}

// UpdateFreqency
//
// swagger:route PATCH /freqencys/{ID} freqencys updateFreqency
//
// # Update a freqency
//
// Responses:
// default: genericError
//
//	200: freqencyDBResponse
func (controller *Controller) UpdateFreqency(c *gin.Context) {

	mutexFreqency.Lock()
	defer mutexFreqency.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFreqency", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	// Validate input
	var input orm.FreqencyAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var freqencyDB orm.FreqencyDB

	// fetch the freqency
	query := db.First(&freqencyDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	freqencyDB.CopyBasicFieldsFromFreqency_WOP(&input.Freqency_WOP)
	freqencyDB.FreqencyPointersEncoding = input.FreqencyPointersEncoding

	query = db.Model(&freqencyDB).Updates(freqencyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	freqencyNew := new(models.Freqency)
	freqencyDB.CopyBasicFieldsToFreqency(freqencyNew)

	// redeem pointers
	freqencyDB.DecodePointers(backRepo, freqencyNew)

	// get stage instance from DB instance, and call callback function
	freqencyOld := backRepo.BackRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]
	if freqencyOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), freqencyOld, freqencyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the freqencyDB
	c.JSON(http.StatusOK, freqencyDB)
}

// DeleteFreqency
//
// swagger:route DELETE /freqencys/{ID} freqencys deleteFreqency
//
// # Delete a freqency
//
// default: genericError
//
//	200: freqencyDBResponse
func (controller *Controller) DeleteFreqency(c *gin.Context) {

	mutexFreqency.Lock()
	defer mutexFreqency.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFreqency", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtone/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFreqency.GetDB()

	// Get model if exist
	var freqencyDB orm.FreqencyDB
	if err := db.First(&freqencyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&freqencyDB)

	// get an instance (not staged) from DB instance, and call callback function
	freqencyDeleted := new(models.Freqency)
	freqencyDB.CopyBasicFieldsToFreqency(freqencyDeleted)

	// get stage instance from DB instance, and call callback function
	freqencyStaged := backRepo.BackRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]
	if freqencyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), freqencyStaged, freqencyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}