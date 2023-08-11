package gohadmin

import (
	"context"

	adminctx "github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
)

const (
	AdministratorRoleSlug    = "administrator"
	HasAcceptedTermsRoleSlug = "accepted_terms"
)

var (
	RoleSlugToID = map[string]int32{}
	RoleSlugs    = []string{AdministratorRoleSlug, HasAcceptedTermsRoleSlug}
)

func GetCurrentUser(ctx *adminctx.Context) (*models.UserModel, bool) {
	usersValue, ok := ctx.UserValue["user"]
	if !ok {
		return nil, false
	}
	usr, ok := usersValue.(models.UserModel)
	return &usr, ok
}

func GetContext(ctx *adminctx.Context) context.Context {
	return ctx.Request.Context()
}
