package hackathon

import (
	"net/http"
)

// App struct represents the application
type App struct {
	Name string `json:""` // Name of the app.
}

// New creates a new instance of the App with given name.
func New() *App {
	return &App{}
}

func (app *App) HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the home page
	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (app *App) Sort(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the sort page
	if r.Method == "GET" {
		http.ServeFile(w, r, "sort.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/sort", http.StatusSeeOther)
	}
}

func (app *App) HandleSearch(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the search page
	if r.Method == "GET" {
		http.ServeFile(w, r, "search.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
	}
}

func (app *App) HandleFilter(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the filter page
	if r.Method == "GET" {
		http.ServeFile(w, r, "filter.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/filter", http.StatusSeeOther)
	}
}

func (app *App) Add(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the add page
	if r.Method == "GET" {
		http.ServeFile(w, r, "add.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/add", http.StatusSeeOther)
	}
}
func (app *App) Edit(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the edit page
	if r.Method == "GET" {
		http.ServeFile(w, r, "edit.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/edit", http.StatusSeeOther)
	}
}

func (app *App) Delete(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the delete page
	if r.Method == "GET" {
		http.ServeFile(w, r, "delete.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/delete", http.StatusSeeOther)
	}
}

func (app *App) Update(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the update page
	if r.Method == "GET" {
		http.ServeFile(w, r, "update.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/update", http.StatusSeeOther)
	}
}

func (app *App) Get(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the get page
	if r.Method == "GET" {
		http.ServeFile(w, r, "get.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/get", http.StatusSeeOther)
	}
}

func (app *App) GetAll(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the get all page
	if r.Method == "GET" {
		http.ServeFile(w, r, "get_all.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/get_all", http.StatusSeeOther)
	}
}

func (app *App) GetAllSorted(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the get all sorted page
	if r.Method == "GET" {
		http.ServeFile(w, r, "get_all_sorted.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/get_all_sorted", http.StatusSeeOther)
	}
}

func (app *App) GetAllFiltered(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the get all filtered page
	if r.Method == "GET" {
		http.ServeFile(w, r, "get_all_filtered.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/get_all_filtered", http.StatusSeeOther)
	}
}

func (app *App) GetAllSearched(w http.ResponseWriter, r *http.Request) {
	// Manage the access request to the get all searched page
	if r.Method == "GET" {
		http.ServeFile(w, r, "get_all_searched.html")
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/get_all_searched", http.StatusSeeOther)
	}
}
