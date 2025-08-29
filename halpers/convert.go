package halpers

import "database/sql"

func ConvertToInt64(data sql.NullInt64) *int64 {
	if data.Valid {
		return &data.Int64
	}
	return nil
}
