package routes

import "github.com/go-chi/chi/v5"

func (p *Provider) MapEndpoints() {
	p.Mux.Group(func(r chi.Router) {
		mux := r.With(p.middlewares.Authenticated)

		mux.Get("/v1/whoami", p.handlers.WhoAmI)

		mux.Get("/v1/tasks", p.handlers.GetAllTasks)
		mux.Post("/v1/tasks", p.handlers.CreateTask)
		//mux.Get("/tasks/create/cancel", p.handlers.CreatTaskViewCancel)
	})

	p.Mux.Post("/v1/signin", p.handlers.SignIn)
	p.Mux.Post("/v1/signup", p.handlers.SignUp)
	p.Mux.Get("/v1/images/{id}", p.handlers.GetImage)
	p.Mux.Post("/v1/images", p.handlers.CreateImage)
}
