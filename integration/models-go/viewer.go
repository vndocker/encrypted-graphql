package models

import "github.com/LaBanHSPO/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
