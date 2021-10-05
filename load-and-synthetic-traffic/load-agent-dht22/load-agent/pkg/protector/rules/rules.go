package rules

import (
	"inditex.com/protector/pkg/protector"
	"time"

	"github.com/influxdata/influxql"
)

// Options for applying rules
type Options struct {
	MaxBuckets  int
	SlowQuery   int64
	MaxDuration time.Duration
}

// RunRules runs the validation rules against the statement
func RunRules(query influxql.Statement, options *Options) error {
	if err := protector.ValidateQueryType(query); err != nil {
		return err
	}

	if err := validateDataPoints(query, options); err != nil {
		return err
	}

	if err := validateWhereClause(query); err != nil {
		return err
	}

	if err := validateSeries(query); err != nil {
		return err
	}

	if err := validateMaxDuration(query, options); err != nil {
		return err
	}
	return nil
}
