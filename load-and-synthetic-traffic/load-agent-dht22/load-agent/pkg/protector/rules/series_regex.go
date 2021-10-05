package rules

import (
	"fmt"
	"github.com/influxdata/influxql"
	"strings"
)

const minRegexLength = 5

func validateSeries(query influxql.Statement) error {
	s := influxql.NewScanner(strings.NewReader(query.String()))
	tok, _, _ := s.ScanRegex()
	if tok != influxql.REGEX {
		return fmt.Errorf("query has no regex")
	}
	regex := query.String()
	regex = strings.TrimLeft(regex, "/")
	regex = strings.TrimRight(regex, "/")
	if len(regex) < minRegexLength || !strings.Contains(regex, ".") {
		return fmt.Errorf("regex is too short: %s", regex)
	}
	return nil
}
