// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package querytest

import (
	"database/sql"
)

type Bar struct {
	ID                     int32
	NobodyWouldBelieveThis sql.NullInt32
	ParentID               sql.NullInt32
}
