package routes

import "github.com/go-chi/chi/v5"

func (p *Provider) mapViews() {
	// Auth protected routes
	p.Mux.Get("/", p.handlers.HomeView)

	p.Mux.Group(func(r chi.Router) {
		mux := r.With(p.middlewares.Authenticated)

		mux.Get("/", p.handlers.HomeView)
		mux.Get("/logout", p.handlers.Logout)
		mux.Get("/tasks/create", p.handlers.CreateTaskView)
		mux.Get("/tasks/create/cancel", p.handlers.CreatTaskViewCancel)
	})

	p.Mux.Get("/signin", p.handlers.SignInView)
	p.Mux.Get("/signup", p.handlers.SignUpView)
	p.Mux.Get("/error", p.handlers.ErrorView)
	p.Mux.Get("/images/{id}", p.handlers.GetImage)
}
