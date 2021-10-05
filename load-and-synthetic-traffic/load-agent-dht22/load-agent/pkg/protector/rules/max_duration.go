package rules

import (
	"fmt"
	"github.com/influxdata/influxql"
	"time"
)

func validateMaxDuration(query influxql.Statement, options *Options) error {
	now := time.Now().UTC()
	if stmt, ok := query.(*influxql.SelectStatement); ok {
		nowValuer := influxql.NowValuer{Now: now}
		stmt.Condition = influxql.Reduce(stmt.Condition, &nowValuer)
		for _, d := range stmt.Dimensions {
			d.Expr = influxql.Reduce(d.Expr, &nowValuer)
		}
		// Replace time(..) references in the query
		stmt.RewriteTimeFields()
		var _, timeRange, err = influxql.ConditionExpr(stmt.Condition, &nowValuer)
		//minTime, maxTime, err := influxql.TimeRange(stmt.Condition)
		if timeRange.Max.IsZero() {
			timeRange.Max = now
		}
		if err != nil {
			return err
		}
		duration := timeRange.Max.Sub(timeRange.Min)
		if duration > options.MaxDuration {
			return fmt.Errorf("max-duration limit exceeded: (%d/%d)", duration, options.MaxDuration)
		}
		return nil
	}
	return nil
}
