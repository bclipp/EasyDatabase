package github.com/bclipp/EasyDatabase

import "fmt"

// update_tables is used for handling the update process
// table field
// Params:
//       table: table to generate the update query for.
//return:
//       the error
func updateTableQuery(table string, row Row) string {
	stateFips := row.StateFips
	stateCode := row.StateCode
	blockPop := row.BlockPop
	blockID := row.BlockID
	tableID := row.ID

	return fmt.Sprintf(
		"UPDATE %s SET state_fips = %d, state_code = '%s', block_pop = %d, block_id = %d WHERE ID = %d;",
		table, stateFips, stateCode, blockPop, blockID, tableID)
}

// select_table is used for generating a query for selecting a table
// table field
// Params:
//       table: table to generate the update query for.
//		 limit: < 0 will cause assume you don't want a limit
//return:
//       the error
func selectTableQuery(table string, limit int) string {
	if limit < 0 {
		return fmt.Sprintf(
			"SELECT id,latitude,longitude FROM %s;",
			table)
	} else {
		return fmt.Sprintf(
			"SELECT id,latitude, longitude FROM %s LIMIT %d;",
			table, limit)
	}
}
