package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// func (app *application) routes() http.Handler {
// 	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
// 	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

// 	mux := pat.New()
// 	mux.Get("/", dynamicMiddleware.ThenFunc(app.homePage))
// 	mux.Get("/post/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createPostPageForm))
// 	mux.Post("/post/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createPostPage))
// 	mux.Get("/post/:id", dynamicMiddleware.ThenFunc(app.showPostPage))
// 	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
// 	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
// 	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
// 	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
// 	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))
// 	fileServer := http.FileServer(http.Dir("./ui/static/"))
// 	mux.Get("/static/", http.StripPrefix("/static", fileServer))

// 	return standardMiddleware.Then(mux)
// }

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)
	rau := alice.New(app.requireAuthenticatedUser)

	router := mux.NewRouter()
	router.Handle("/", dynamicMiddleware.Then(http.HandlerFunc(app.homePage))).Methods("GET")
	// router.HandleFunc("/", app.homePage).Methods("GET")

	router.Handle("/post/create", dynamicMiddleware.Then(rau.Then(http.HandlerFunc(app.createPostPageForm)))).Methods("GET")
	router.Handle("/post/create", dynamicMiddleware.Then(rau.Then(http.HandlerFunc(app.createPostPage)))).Methods("POST")
	router.Handle("/post/{id}", dynamicMiddleware.Then(http.HandlerFunc(app.showPostPage))).Methods("GET")

	router.Handle("/user/signup", dynamicMiddleware.Then(http.HandlerFunc(app.signupUserForm))).Methods("GET")
	router.Handle("/user/signup", dynamicMiddleware.Then(http.HandlerFunc(app.signupUser))).Methods("POST")
	router.Handle("/user/login", dynamicMiddleware.Then(http.HandlerFunc(app.loginUserForm))).Methods("GET")
	router.Handle("/user/login", dynamicMiddleware.Then(http.HandlerFunc(app.loginUser))).Methods("POST")
	router.Handle("/user/logout", dynamicMiddleware.Then(rau.Then(http.HandlerFunc(app.logoutUser)))).Methods("POST")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer)).Methods("GET")

	return standardMiddleware.Then(router)

}
