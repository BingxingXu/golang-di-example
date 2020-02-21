package app

// AppService is service
type AppService struct {
	AppQuerySet AppQuerySet
}

func ProviderAppService(p AppQuerySet) AppService {
	return AppService{AppQuerySet: p}
}

func (s *AppService) FindAll() []App {
	var ret []App
	s.AppQuerySet.All(&ret)
	return ret
}
