package rbac

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

// CheckRBAC uses OPA to check if the user has the necessary role for the requested resource
func CheckRBAC(opaPolicy, userRole, action, resource_type string) (bool, error) {
	ctx := rego.New(
		rego.Query("data.main.allow"),
		rego.Module("main.rego", opaPolicy),
		rego.Input(map[string]interface{}{
			"action":        action,
			"resource_type": resource_type,
			"role":          userRole,
		}),
	)

	result, err := ctx.Eval(context.Background())
	if err != nil {
		fmt.Println("OPA error:", err)
		return false, err
	}

	if len(result) > 0 && result[0].Expressions[0].Value == true {
		return true, nil
	}

	return false, nil
}
