// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Apps", testApps)
	t.Run("AppInventories", testAppInventories)
	t.Run("AppPlaybooks", testAppPlaybooks)
	t.Run("Failures", testFailures)
	t.Run("Inventories", testInventories)
	t.Run("Jobs", testJobs)
	t.Run("Playbooks", testPlaybooks)
	t.Run("Runs", testRuns)
	t.Run("Statistics", testStatistics)
}

func TestDelete(t *testing.T) {
	t.Run("Apps", testAppsDelete)
	t.Run("AppInventories", testAppInventoriesDelete)
	t.Run("AppPlaybooks", testAppPlaybooksDelete)
	t.Run("Failures", testFailuresDelete)
	t.Run("Inventories", testInventoriesDelete)
	t.Run("Jobs", testJobsDelete)
	t.Run("Playbooks", testPlaybooksDelete)
	t.Run("Runs", testRunsDelete)
	t.Run("Statistics", testStatisticsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Apps", testAppsQueryDeleteAll)
	t.Run("AppInventories", testAppInventoriesQueryDeleteAll)
	t.Run("AppPlaybooks", testAppPlaybooksQueryDeleteAll)
	t.Run("Failures", testFailuresQueryDeleteAll)
	t.Run("Inventories", testInventoriesQueryDeleteAll)
	t.Run("Jobs", testJobsQueryDeleteAll)
	t.Run("Playbooks", testPlaybooksQueryDeleteAll)
	t.Run("Runs", testRunsQueryDeleteAll)
	t.Run("Statistics", testStatisticsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Apps", testAppsSliceDeleteAll)
	t.Run("AppInventories", testAppInventoriesSliceDeleteAll)
	t.Run("AppPlaybooks", testAppPlaybooksSliceDeleteAll)
	t.Run("Failures", testFailuresSliceDeleteAll)
	t.Run("Inventories", testInventoriesSliceDeleteAll)
	t.Run("Jobs", testJobsSliceDeleteAll)
	t.Run("Playbooks", testPlaybooksSliceDeleteAll)
	t.Run("Runs", testRunsSliceDeleteAll)
	t.Run("Statistics", testStatisticsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Apps", testAppsExists)
	t.Run("AppInventories", testAppInventoriesExists)
	t.Run("AppPlaybooks", testAppPlaybooksExists)
	t.Run("Failures", testFailuresExists)
	t.Run("Inventories", testInventoriesExists)
	t.Run("Jobs", testJobsExists)
	t.Run("Playbooks", testPlaybooksExists)
	t.Run("Runs", testRunsExists)
	t.Run("Statistics", testStatisticsExists)
}

func TestFind(t *testing.T) {
	t.Run("Apps", testAppsFind)
	t.Run("AppInventories", testAppInventoriesFind)
	t.Run("AppPlaybooks", testAppPlaybooksFind)
	t.Run("Failures", testFailuresFind)
	t.Run("Inventories", testInventoriesFind)
	t.Run("Jobs", testJobsFind)
	t.Run("Playbooks", testPlaybooksFind)
	t.Run("Runs", testRunsFind)
	t.Run("Statistics", testStatisticsFind)
}

func TestBind(t *testing.T) {
	t.Run("Apps", testAppsBind)
	t.Run("AppInventories", testAppInventoriesBind)
	t.Run("AppPlaybooks", testAppPlaybooksBind)
	t.Run("Failures", testFailuresBind)
	t.Run("Inventories", testInventoriesBind)
	t.Run("Jobs", testJobsBind)
	t.Run("Playbooks", testPlaybooksBind)
	t.Run("Runs", testRunsBind)
	t.Run("Statistics", testStatisticsBind)
}

func TestOne(t *testing.T) {
	t.Run("Apps", testAppsOne)
	t.Run("AppInventories", testAppInventoriesOne)
	t.Run("AppPlaybooks", testAppPlaybooksOne)
	t.Run("Failures", testFailuresOne)
	t.Run("Inventories", testInventoriesOne)
	t.Run("Jobs", testJobsOne)
	t.Run("Playbooks", testPlaybooksOne)
	t.Run("Runs", testRunsOne)
	t.Run("Statistics", testStatisticsOne)
}

func TestAll(t *testing.T) {
	t.Run("Apps", testAppsAll)
	t.Run("AppInventories", testAppInventoriesAll)
	t.Run("AppPlaybooks", testAppPlaybooksAll)
	t.Run("Failures", testFailuresAll)
	t.Run("Inventories", testInventoriesAll)
	t.Run("Jobs", testJobsAll)
	t.Run("Playbooks", testPlaybooksAll)
	t.Run("Runs", testRunsAll)
	t.Run("Statistics", testStatisticsAll)
}

func TestCount(t *testing.T) {
	t.Run("Apps", testAppsCount)
	t.Run("AppInventories", testAppInventoriesCount)
	t.Run("AppPlaybooks", testAppPlaybooksCount)
	t.Run("Failures", testFailuresCount)
	t.Run("Inventories", testInventoriesCount)
	t.Run("Jobs", testJobsCount)
	t.Run("Playbooks", testPlaybooksCount)
	t.Run("Runs", testRunsCount)
	t.Run("Statistics", testStatisticsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Apps", testAppsHooks)
	t.Run("AppInventories", testAppInventoriesHooks)
	t.Run("AppPlaybooks", testAppPlaybooksHooks)
	t.Run("Failures", testFailuresHooks)
	t.Run("Inventories", testInventoriesHooks)
	t.Run("Jobs", testJobsHooks)
	t.Run("Playbooks", testPlaybooksHooks)
	t.Run("Runs", testRunsHooks)
	t.Run("Statistics", testStatisticsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Apps", testAppsInsert)
	t.Run("Apps", testAppsInsertWhitelist)
	t.Run("AppInventories", testAppInventoriesInsert)
	t.Run("AppInventories", testAppInventoriesInsertWhitelist)
	t.Run("AppPlaybooks", testAppPlaybooksInsert)
	t.Run("AppPlaybooks", testAppPlaybooksInsertWhitelist)
	t.Run("Failures", testFailuresInsert)
	t.Run("Failures", testFailuresInsertWhitelist)
	t.Run("Inventories", testInventoriesInsert)
	t.Run("Inventories", testInventoriesInsertWhitelist)
	t.Run("Jobs", testJobsInsert)
	t.Run("Jobs", testJobsInsertWhitelist)
	t.Run("Playbooks", testPlaybooksInsert)
	t.Run("Playbooks", testPlaybooksInsertWhitelist)
	t.Run("Runs", testRunsInsert)
	t.Run("Runs", testRunsInsertWhitelist)
	t.Run("Statistics", testStatisticsInsert)
	t.Run("Statistics", testStatisticsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Apps", testAppsReload)
	t.Run("AppInventories", testAppInventoriesReload)
	t.Run("AppPlaybooks", testAppPlaybooksReload)
	t.Run("Failures", testFailuresReload)
	t.Run("Inventories", testInventoriesReload)
	t.Run("Jobs", testJobsReload)
	t.Run("Playbooks", testPlaybooksReload)
	t.Run("Runs", testRunsReload)
	t.Run("Statistics", testStatisticsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Apps", testAppsReloadAll)
	t.Run("AppInventories", testAppInventoriesReloadAll)
	t.Run("AppPlaybooks", testAppPlaybooksReloadAll)
	t.Run("Failures", testFailuresReloadAll)
	t.Run("Inventories", testInventoriesReloadAll)
	t.Run("Jobs", testJobsReloadAll)
	t.Run("Playbooks", testPlaybooksReloadAll)
	t.Run("Runs", testRunsReloadAll)
	t.Run("Statistics", testStatisticsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Apps", testAppsSelect)
	t.Run("AppInventories", testAppInventoriesSelect)
	t.Run("AppPlaybooks", testAppPlaybooksSelect)
	t.Run("Failures", testFailuresSelect)
	t.Run("Inventories", testInventoriesSelect)
	t.Run("Jobs", testJobsSelect)
	t.Run("Playbooks", testPlaybooksSelect)
	t.Run("Runs", testRunsSelect)
	t.Run("Statistics", testStatisticsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Apps", testAppsUpdate)
	t.Run("AppInventories", testAppInventoriesUpdate)
	t.Run("AppPlaybooks", testAppPlaybooksUpdate)
	t.Run("Failures", testFailuresUpdate)
	t.Run("Inventories", testInventoriesUpdate)
	t.Run("Jobs", testJobsUpdate)
	t.Run("Playbooks", testPlaybooksUpdate)
	t.Run("Runs", testRunsUpdate)
	t.Run("Statistics", testStatisticsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Apps", testAppsSliceUpdateAll)
	t.Run("AppInventories", testAppInventoriesSliceUpdateAll)
	t.Run("AppPlaybooks", testAppPlaybooksSliceUpdateAll)
	t.Run("Failures", testFailuresSliceUpdateAll)
	t.Run("Inventories", testInventoriesSliceUpdateAll)
	t.Run("Jobs", testJobsSliceUpdateAll)
	t.Run("Playbooks", testPlaybooksSliceUpdateAll)
	t.Run("Runs", testRunsSliceUpdateAll)
	t.Run("Statistics", testStatisticsSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Apps", testAppsUpsert)
	t.Run("AppInventories", testAppInventoriesUpsert)
	t.Run("AppPlaybooks", testAppPlaybooksUpsert)
	t.Run("Failures", testFailuresUpsert)
	t.Run("Inventories", testInventoriesUpsert)
	t.Run("Jobs", testJobsUpsert)
	t.Run("Playbooks", testPlaybooksUpsert)
	t.Run("Runs", testRunsUpsert)
	t.Run("Statistics", testStatisticsUpsert)
}