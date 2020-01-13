package pkg

import (
	"github.com/graphql-go/graphql"
)

// NullableStringField field for nullable string
var NullableStringField = &graphql.Field{
	Type:    graphql.String,
	Resolve: ResolveNullable,
}

// NullableInt64Field field for nullable int64
var NullableInt64Field = &graphql.Field{
	Type:    graphql.Int,
	Resolve: ResolveNullable,
}

// NullableFloat64Field field for nullable float64
var NullableFloat64Field = &graphql.Field{
	Type:    graphql.Float,
	Resolve: ResolveNullable,
}

// NullableBoolField field for nullable bool
var NullableBoolField = &graphql.Field{
	Type:    graphql.Boolean,
	Resolve: ResolveNullable,
}

// NullableDateTimeField field for nullable datetime
var NullableDateTimeField = &graphql.Field{
	Type:    graphql.DateTime,
	Resolve: ResolveNullable,
}
