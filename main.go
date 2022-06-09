package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type hello struct {
	app.Compo

	name string
	cols int
}

func (h *hello) Render() app.UI {
	fmt.Println(h.cols)
	h.cols = int(math.Max(float64(h.cols), 1))
	elems := make([]app.UI, h.cols)
	for i := 0; i < h.cols; i++ {
		elems[i] = app.Td().Body(app.H1().Body(
			app.Text("Hello, "),
			app.If(h.name != "",
				app.Text(h.name),
			).Else(
				app.Text("World!"),
			),
		))
	}
	hellos := app.Table().Body(elems...).Style("width", "100%")

	return app.Div().Body(
		hellos,
		app.P().Body(
			app.Input().
				Type("text").
				Value(h.name).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(h.ValueTo(&h.name)),
		),
		app.P().Body(
			app.Input().
				Type("int").
				Value(h.cols).
				Placeholder("How many times?").
				AutoFocus(false).
				OnChange(h.ValueTo(&h.cols)),
		),
	)
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &hello{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"/assets/css/murlok.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
