package service

type Storeger interface {
	AddURL(string) string
	GetURL(string) (string, bool)
}

type ConfigerService interface {
	GetAddressServerURL() string
}

func New(db Storeger, conf ConfigerService) *NodeService {
	return &NodeService{db: db, conf: conf}
}

type NodeService struct {
	db   Storeger
	conf ConfigerService
}

func (r *NodeService) GetURL(id string) (string, bool) {
	return r.db.GetURL(id)
}

func (r *NodeService) addURL(link string) string {
	return r.db.AddURL(link)
}

func (r *NodeService) GetSmallLink(longLink string) string {
	id := r.addURL(longLink)
	return r.conf.GetAddressServerURL() + id
}
