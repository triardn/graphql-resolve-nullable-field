package pkg

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/graphql-go/graphql"
)

const (
	formatError = `[ResolveNullable]`
)

func resolveFieldValue(fieldValue interface{}) (result interface{}) {
	switch value := fieldValue.(type) {
	case sql.NullString:
		if value.Valid {
			result = value.String
		}
	case *sql.NullString:
		if value != nil && value.Valid {
			result = value.String
		}
	case sql.NullInt32:
		if value.Valid {
			result = value.Int32
		}
	case *sql.NullInt32:
		if value != nil && value.Valid {
			result = value.Int32
		}
	case sql.NullInt64:
		if value.Valid {
			result = value.Int64
		}
	case *sql.NullInt64:
		if value != nil && value.Valid {
			result = value.Int64
		}
	case sql.NullFloat64:
		if value.Valid {
			result = value.Float64
		}
	case *sql.NullFloat64:
		if value != nil && value.Valid {
			result = value.Float64
		}
	case sql.NullBool:
		if value.Valid {
			result = value.Bool
		}
	case *sql.NullBool:
		if value != nil && value.Valid {
			result = value.Bool
		}
	case sql.NullTime:
		if value.Valid {
			result = value.Time
		}
	case *sql.NullTime:
		if value != nil && value.Valid {
			result = value.Time
		}
	case *string:
		if value != nil {
			result = *value
		}
	case []byte:
		result = string(value)
	case *[]byte:
		if value != nil {
			result = string(*value)
		}
	default:
		if fieldValue != nil {
			result = fieldValue
		}
	}
	return
}

// ResolveNullable resolve the sql.Null* field
func ResolveNullable(p graphql.ResolveParams) (result interface{}, err error) {
	v := reflect.ValueOf(p.Source)
	if !v.IsValid() {
		errMsg := fmt.Sprintf("%s Failed when try to get value: %s", formatError, p.Info.FieldName)
		return nil, errors.New(errMsg)
	}

	fieldName := strings.Title(p.Info.FieldName)
	f := v.FieldByName(fieldName)
	if !f.IsValid() {
		errMsg := fmt.Sprintf("%s missing field: %s", formatError, fieldName)
		return nil, errors.New(errMsg)
	}

	result = resolveFieldValue(f)

	return
}
