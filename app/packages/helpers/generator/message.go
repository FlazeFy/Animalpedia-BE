package generator

import (
	"app/packages/helpers/typography"
	"fmt"
)

func GenerateQueryMsg(ctx string, total int) string {
	ctx = typography.UcFirst(ctx)

	if total > 0 {
		return ctx + " found"
	} else {
		return ctx + " not found"
	}
}

func GenerateCommandMsg(ctx string, cmd string, is_success bool) string {
	if is_success {
		return "Successfully " + cmd + " " + ctx
	} else {
		return "Failed to " + cmd + " " + ctx
	}
}

func GenerateValidatorMsg(ctx string, min, max int) string {
	if ctx != "Valid until" {
		if min != 0 && max != 0 {
			return fmt.Sprintf("%s must be between %d and %d characters", ctx, min, max)
		} else if min == 0 {
			return fmt.Sprintf("%s must be below than %d characters", ctx, max)
		} else {
			return fmt.Sprintf("%s must be more than %d characters", ctx, min)
		}
	} else {
		return fmt.Sprintf("%s must be between year %d and %d", ctx, min, max)
	}
}
