// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetails)
	t.Run("Answers", testAnswers)
	t.Run("Owners", testOwners)
	t.Run("Records", testRecords)
}

func TestSoftDelete(t *testing.T) {}

func TestQuerySoftDeleteAll(t *testing.T) {}

func TestSliceSoftDeleteAll(t *testing.T) {}

func TestDelete(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsDelete)
	t.Run("Answers", testAnswersDelete)
	t.Run("Owners", testOwnersDelete)
	t.Run("Records", testRecordsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsQueryDeleteAll)
	t.Run("Answers", testAnswersQueryDeleteAll)
	t.Run("Owners", testOwnersQueryDeleteAll)
	t.Run("Records", testRecordsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsSliceDeleteAll)
	t.Run("Answers", testAnswersSliceDeleteAll)
	t.Run("Owners", testOwnersSliceDeleteAll)
	t.Run("Records", testRecordsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsExists)
	t.Run("Answers", testAnswersExists)
	t.Run("Owners", testOwnersExists)
	t.Run("Records", testRecordsExists)
}

func TestFind(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsFind)
	t.Run("Answers", testAnswersFind)
	t.Run("Owners", testOwnersFind)
	t.Run("Records", testRecordsFind)
}

func TestBind(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsBind)
	t.Run("Answers", testAnswersBind)
	t.Run("Owners", testOwnersBind)
	t.Run("Records", testRecordsBind)
}

func TestOne(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsOne)
	t.Run("Answers", testAnswersOne)
	t.Run("Owners", testOwnersOne)
	t.Run("Records", testRecordsOne)
}

func TestAll(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsAll)
	t.Run("Answers", testAnswersAll)
	t.Run("Owners", testOwnersAll)
	t.Run("Records", testRecordsAll)
}

func TestCount(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsCount)
	t.Run("Answers", testAnswersCount)
	t.Run("Owners", testOwnersCount)
	t.Run("Records", testRecordsCount)
}

func TestHooks(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsHooks)
	t.Run("Answers", testAnswersHooks)
	t.Run("Owners", testOwnersHooks)
	t.Run("Records", testRecordsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsInsert)
	t.Run("AnswerDetails", testAnswerDetailsInsertWhitelist)
	t.Run("Answers", testAnswersInsert)
	t.Run("Answers", testAnswersInsertWhitelist)
	t.Run("Owners", testOwnersInsert)
	t.Run("Owners", testOwnersInsertWhitelist)
	t.Run("Records", testRecordsInsert)
	t.Run("Records", testRecordsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AnswerDetailToAnswerUsingAnswer", testAnswerDetailToOneAnswerUsingAnswer)
	t.Run("AnswerToRecordUsingRecord", testAnswerToOneRecordUsingRecord)
	t.Run("AnswerToOwnerUsingOwner", testAnswerToOneOwnerUsingOwner)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("AnswerToAnswerDetailUsingAnswerDetail", testAnswerOneToOneAnswerDetailUsingAnswerDetail)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("OwnerToAnswers", testOwnerToManyAnswers)
	t.Run("RecordToAnswers", testRecordToManyAnswers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AnswerDetailToAnswerUsingAnswerDetail", testAnswerDetailToOneSetOpAnswerUsingAnswer)
	t.Run("AnswerToRecordUsingAnswers", testAnswerToOneSetOpRecordUsingRecord)
	t.Run("AnswerToOwnerUsingAnswers", testAnswerToOneSetOpOwnerUsingOwner)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("AnswerToAnswerDetailUsingAnswerDetail", testAnswerOneToOneSetOpAnswerDetailUsingAnswerDetail)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("OwnerToAnswers", testOwnerToManyAddOpAnswers)
	t.Run("RecordToAnswers", testRecordToManyAddOpAnswers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsReload)
	t.Run("Answers", testAnswersReload)
	t.Run("Owners", testOwnersReload)
	t.Run("Records", testRecordsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsReloadAll)
	t.Run("Answers", testAnswersReloadAll)
	t.Run("Owners", testOwnersReloadAll)
	t.Run("Records", testRecordsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsSelect)
	t.Run("Answers", testAnswersSelect)
	t.Run("Owners", testOwnersSelect)
	t.Run("Records", testRecordsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsUpdate)
	t.Run("Answers", testAnswersUpdate)
	t.Run("Owners", testOwnersUpdate)
	t.Run("Records", testRecordsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AnswerDetails", testAnswerDetailsSliceUpdateAll)
	t.Run("Answers", testAnswersSliceUpdateAll)
	t.Run("Owners", testOwnersSliceUpdateAll)
	t.Run("Records", testRecordsSliceUpdateAll)
}
