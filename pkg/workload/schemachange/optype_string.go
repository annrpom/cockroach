// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

// Code generated by "stringer"; DO NOT EDIT.

package schemachange

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[insertRow-0]
	_ = x[selectStmt-1]
	_ = x[validate-2]
	_ = x[renameIndex-3]
	_ = x[renameSequence-4]
	_ = x[renameTable-5]
	_ = x[renameView-6]
	_ = x[alterDatabaseAddRegion-7]
	_ = x[alterDatabasePrimaryRegion-8]
	_ = x[alterDatabaseSurvivalGoal-9]
	_ = x[alterTableAddColumn-10]
	_ = x[alterTableAddConstraint-11]
	_ = x[alterTableAddConstraintForeignKey-12]
	_ = x[alterTableAddConstraintUnique-13]
	_ = x[alterTableAlterColumnType-14]
	_ = x[alterTableDropColumn-15]
	_ = x[alterTableDropConstraint-16]
	_ = x[alterTableDropNotNull-17]
	_ = x[alterTableDropColumnDefault-18]
	_ = x[alterTableDropStored-19]
	_ = x[alterTableLocality-20]
	_ = x[alterTableRenameColumn-21]
	_ = x[alterTableSetColumnDefault-22]
	_ = x[alterTableSetColumnNotNull-23]
	_ = x[alterTypeDropValue-24]
	_ = x[createTypeEnum-25]
	_ = x[createIndex-26]
	_ = x[createSchema-27]
	_ = x[createSequence-28]
	_ = x[createTable-29]
	_ = x[createTableAs-30]
	_ = x[createView-31]
	_ = x[dropIndex-32]
	_ = x[dropSchema-33]
	_ = x[dropSequence-34]
	_ = x[dropTable-35]
	_ = x[dropView-36]
}

func (i opType) String() string {
	switch i {
	case insertRow:
		return "insertRow"
	case selectStmt:
		return "selectStmt"
	case validate:
		return "validate"
	case renameIndex:
		return "renameIndex"
	case renameSequence:
		return "renameSequence"
	case renameTable:
		return "renameTable"
	case renameView:
		return "renameView"
	case alterDatabaseAddRegion:
		return "alterDatabaseAddRegion"
	case alterDatabasePrimaryRegion:
		return "alterDatabasePrimaryRegion"
	case alterDatabaseSurvivalGoal:
		return "alterDatabaseSurvivalGoal"
	case alterTableAddColumn:
		return "alterTableAddColumn"
	case alterTableAddConstraint:
		return "alterTableAddConstraint"
	case alterTableAddConstraintForeignKey:
		return "alterTableAddConstraintForeignKey"
	case alterTableAddConstraintUnique:
		return "alterTableAddConstraintUnique"
	case alterTableAlterColumnType:
		return "alterTableAlterColumnType"
	case alterTableDropColumn:
		return "alterTableDropColumn"
	case alterTableDropConstraint:
		return "alterTableDropConstraint"
	case alterTableDropNotNull:
		return "alterTableDropNotNull"
	case alterTableDropColumnDefault:
		return "alterTableDropColumnDefault"
	case alterTableDropStored:
		return "alterTableDropStored"
	case alterTableLocality:
		return "alterTableLocality"
	case alterTableRenameColumn:
		return "alterTableRenameColumn"
	case alterTableSetColumnDefault:
		return "alterTableSetColumnDefault"
	case alterTableSetColumnNotNull:
		return "alterTableSetColumnNotNull"
	case alterTypeDropValue:
		return "alterTypeDropValue"
	case createTypeEnum:
		return "createTypeEnum"
	case createIndex:
		return "createIndex"
	case createSchema:
		return "createSchema"
	case createSequence:
		return "createSequence"
	case createTable:
		return "createTable"
	case createTableAs:
		return "createTableAs"
	case createView:
		return "createView"
	case dropIndex:
		return "dropIndex"
	case dropSchema:
		return "dropSchema"
	case dropSequence:
		return "dropSequence"
	case dropTable:
		return "dropTable"
	case dropView:
		return "dropView"
	default:
		return "opType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
