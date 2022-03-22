package crud

import (
	"context"
	"interface_project/ent"
)

type Crud struct {
	Ctx    *context.Context
	Client *ent.Client
}
