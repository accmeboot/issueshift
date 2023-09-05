package routes

func (p *Provider) mapForms() {
	p.Mux.Post("/signin", p.handlers.SignInForm)
	p.Mux.Post("/signup", p.handlers.SignUpForm)
}
