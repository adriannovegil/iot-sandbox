package protector

import (
	"fmt"
	"reflect"

	"github.com/influxdata/influxql"
)

// ValidateQueryType determines if the given query should be allowed
func ValidateQueryType(stmt influxql.Statement) error {
	if _, ok := stmt.(*influxql.SelectStatement); ok {
		return nil
	}
	if _, ok := stmt.(*influxql.ShowSeriesStatement); ok {
		return nil
	}
	if _, ok := stmt.(*influxql.ShowTagValuesStatement); ok {
		return nil
	}
	if _, ok := stmt.(*influxql.ShowRetentionPoliciesStatement); ok {
		return nil
	}
	return fmt.Errorf("query type not allowed: %s", reflect.TypeOf(stmt).String())
}
