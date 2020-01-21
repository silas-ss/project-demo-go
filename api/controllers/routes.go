package controllers

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/api/v1/company", s.CreateCompany).Methods("POST")
	s.Router.HandleFunc("/api/v1/company", s.FindAllCompanies).Methods("GET")
	s.Router.HandleFunc("/api/v1/company/{id}", s.FindCompanyByID).Methods("GET")
	s.Router.HandleFunc("/api/v1/company/{id}", s.UpdateCompany).Methods("PUT")
	s.Router.HandleFunc("/api/v1/company/{id}", s.DeleteCompany).Methods("DELETE")
}
