package repos

import "github.com/KendoCross/kendocqrs/interfaces"

type protocolRepo struct {
}

func NewprotocolRepo() interfaces.IProtocolRepo {
	return new(protocolRepo)
}
func (r *protocolRepo) Add(p *interfaces.ProtocolOrm) (int64, error) {
	return writeEngine.Table("Protocol").
		Insert(p)
}
