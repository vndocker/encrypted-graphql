package models

import "github.com/vndocker/encrypted-graphql/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
