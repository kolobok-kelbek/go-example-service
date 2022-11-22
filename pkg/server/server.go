package server

type Server interface {
	Run()
}

type ServerImpl struct {
	router Router
}

func NewServerImpl() *ServerImpl {
	return &ServerImpl{}
}

func (server *ServerImpl) Init() {

}

func (server *ServerImpl) Run() {
}
