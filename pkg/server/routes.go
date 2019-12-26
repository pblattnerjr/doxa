package server

// routes initializes all the routes for the server.
func (s *server) routes() {
	s.router.HandleFunc("/id/{library}/{topic}/{key}", s.handleID("pkg/server/templates/table.gohtml")).Methods("GET")
	s.router.HandleFunc("/id/{topic}/{key}", s.handleTK()).Queries("empty","{empty}").Methods("GET")
	s.router.HandleFunc("/id/{topic}/{key}", s.handleTK()).Methods("GET")
	s.router.HandleFunc("/id", s.handleHome()).Methods("GET")

	s.router.HandleFunc("/topic/{library}/{topic}", s.handleTopic()).Queries("empty","{empty}").Methods("GET")
	s.router.HandleFunc("/topic/{library}/{topic}", s.handleTopic()).Methods("GET")

	s.router.HandleFunc("/", s.handleHome()).Methods("GET")

	// api version 1
	s.api1.HandleFunc("/status", s.handleHomeV1())

	// api version 2
	s.api2.HandleFunc("/status", s.handleHomeV2())
}
