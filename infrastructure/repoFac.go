package infrastructure

import (
	"github.com/KendoCross/kendocqrs/infrastructure/repos"
	"github.com/KendoCross/kendocqrs/interfaces"
)

func init() {
	RepoFac.ProtocolRepo = repos.NewprotocolRepo()
}

var (
	RepoFac *RepoFactory = &RepoFactory{}
)

type RepoFactory struct {
	ProtocolRepo interfaces.IProtocolRepo
}
