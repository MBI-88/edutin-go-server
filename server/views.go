package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

var databasePath = "./data/dataStore.json"

// Index  function. HTML response
func Index(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/index.html")
	if r.Method == "GET" {
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			// logic here
			file, err := os.OpenFile(databasePath, os.O_RDONLY, 0775)
			defer file.Close()
			if err != nil {
				fmt.Fprintf(w, "500 %s", err)
			} else {
				decoder := json.NewDecoder(file)
				var jdata []data
				if err := decoder.Decode(&jdata); err != nil {
					fmt.Fprintf(w, "500 %s", err)
				} else {
					page.Execute(w, jdata)
				}
			}
		}
	} else {
		fmt.Fprintln(w, "404 Bad request, mothod not allowed")
	}

}

// Man function. HTML response
func Man(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/man.html")
	if r.Method == "GET" {
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			file, err := os.OpenFile(databasePath, os.O_RDONLY, 0775)
			defer file.Close()
			if err != nil {
				fmt.Fprintf(w, "500 %s", err)
			} else {
				decoder := json.NewDecoder(file)
				var jdata []data
				if err := decoder.Decode(&jdata); err != nil {
					fmt.Fprintf(w, "500 %s", err)
				} else {
					var mandata []data
					for _, man := range jdata {
						if man.Cate == "man" {
							mandata = append(mandata, man)
						}
					}
					page.Execute(w, mandata)
				}
			}

		}
	} else {
		fmt.Fprintln(w, "404 Bad request, mothod not allowed")
	}

}

// Women function. HTML response
func Women(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/women.html")
	if r.Method == "GET" {
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			file, err := os.OpenFile(databasePath, os.O_RDONLY, 0775)
			defer file.Close()
			if err != nil {
				fmt.Fprintf(w, "500 %s", err)
			} else {
				decoder := json.NewDecoder(file)
				var jdata []data
				if err := decoder.Decode(&jdata); err != nil {
					fmt.Fprintf(w, "500 %s", err)
				} else {
					var womendata []data
					for _, woman := range jdata {
						if woman.Cate == "woman" {
							womendata = append(womendata, woman)
						}
					}
					page.Execute(w, womendata)
				}
			}
		}
	} else {
		fmt.Fprintln(w, "404 Bad request, mothod not allowed")
	}

}

// Children function. HTML response
func Children(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/children.html")
	if r.Method == "GET" {
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			file, err := os.OpenFile(databasePath, os.O_RDONLY, 0775)
			defer file.Close()
			if err != nil {
				fmt.Fprintf(w, "500 %s", err)
			} else {
				decoder := json.NewDecoder(file)
				var jdata []data
				if err := decoder.Decode(&jdata); err != nil {
					fmt.Fprintf(w, "500 %s", err)
				} else {
					var childrendata []data
					for _, child := range jdata {
						if child.Cate == "children" {
							childrendata = append(childrendata, child)
						}
					}
					page.Execute(w, childrendata)
				}
			}
		}
	} else {
		fmt.Fprintln(w, "404 Bad request, mothod not allowed")
	}

}

// Offerts function. HTML response
func Offerts(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/offerts.html")
	if r.Method == "GET" {
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			file, err := os.OpenFile(databasePath, os.O_RDONLY, 0775)
			defer file.Close()
			if err != nil {
				fmt.Fprintf(w, "500 %s", err)
			} else {
				decoder := json.NewDecoder(file)
				var jdata []data
				if err := decoder.Decode(&jdata); err != nil {
					fmt.Fprintf(w, "500 %s", err)
				} else {
					page.Execute(w, jdata[:5])
				}
			}
		}
	} else {
		fmt.Fprintln(w, "404 Bad request, mothod not allowed")
	}

}

// News function. HTML response
func News(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/news.html")
	if r.Method == "GET" {
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			// logic here

			page.Execute(w, nil)
		}
	} else {
		fmt.Fprintln(w, "404 Bad request, mothod not allowed")
	}

}

// Pays function. HTML response
func Pays(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/pays.html")
	switch r.Method {
	case "GET":
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			id := r.URL.Query().Get("id")
			page.Execute(w, id)
		}
	case "POST":
		file, err := os.OpenFile(databasePath, os.O_RDWR, 0775)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			id, _ := strconv.Atoi(r.FormValue("id"))
			decoder := json.NewDecoder(file)
			var jsdata []*data
			if err := decoder.Decode(&jsdata); err != nil {
				fmt.Fprintf(w, "500 %s", err)
			} else {
				for _, item := range jsdata {
					if item.ID == id {
						item.Selected = false
						item.Bought = true
						break
					}
				}
				file.Truncate(0)
				file.Seek(0, 0)
				encoder := json.NewEncoder(file)
				if err := encoder.Encode(jsdata); err != nil {
					fmt.Fprintf(w, "500 %s", err)
				} 

			}
			http.Redirect(w, r, "/car", http.StatusSeeOther)
		}
	default:
		fmt.Fprintln(w, "404 Bad request, mothod not allowed")
	}

}

// Car function. HTML response
func Car(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/carrito.html")
	switch r.Method {
	case "POST":
		if r.FormValue("delete") == "" {
			if err != nil {
				fmt.Fprintf(w, "500 %s", err)
			} else {
				file, err := os.OpenFile(databasePath, os.O_RDWR, 0775)
				defer file.Close()
				if err != nil {
					fmt.Fprintf(w, "500 %s", err)
				} else {
					decoder := json.NewDecoder(file)
					var jdata []*data
					if err := decoder.Decode(&jdata); err != nil {
						fmt.Fprintf(w, "500 %s", err)
					} else {
						var cardata []data
						id, _ := strconv.Atoi(r.FormValue("id"))
						for _, car := range jdata {
							if car.ID == id {
								car.Selected = true
								break
							}
						}
						file.Truncate(0)
						file.Seek(0, 0)
						encoder := json.NewEncoder(file)
						if err := encoder.Encode(jdata); err != nil {
							fmt.Fprintf(w, "500 %s", err)
						} else {
							for _, car := range jdata {
								if car.Selected == true {
									cardata = append(cardata, *car)
								}
							}
						}
						page.Execute(w, cardata)
					}
				}

			}
		} else {
			if err != nil {
				fmt.Fprintf(w, "500 %s", err)
			} else {
				file, err := os.OpenFile(databasePath, os.O_RDWR, 0775)
				defer file.Close()
				if err != nil {
					fmt.Fprintf(w, "500 %s", err)
				} else {
					decoder := json.NewDecoder(file)
					var jdata []*data
					if err := decoder.Decode(&jdata); err != nil {
						fmt.Fprintf(w, "500 %s", err)
					} else {
						var cardata []data
						id, _ := strconv.Atoi(r.FormValue("id"))
						if id == -1 {
							for _, car := range jdata {
								car.Selected = false

							}
						} else {
							for _, car := range jdata {
								if car.ID == id {
									car.Selected = false
									break
								}
							}

						}

						file.Truncate(0)
						file.Seek(0, 0)
						encoder := json.NewEncoder(file)
						if err := encoder.Encode(jdata); err != nil {
							fmt.Fprintf(w, "500 %s", err)
						} else {
							for _, car := range jdata {
								if car.Selected == true {
									cardata = append(cardata, *car)
								}
							}
						}
						page.Execute(w, cardata)
					}
				}
			}
		}

	case "GET":
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			file, err := os.OpenFile(databasePath, os.O_RDONLY, 0775)
			defer file.Close()
			if err != nil {
				fmt.Fprintf(w, "500 %s", err)
			} else {
				var jdata []data
				decoder := json.NewDecoder(file)
				if err := decoder.Decode(&jdata); err != nil {
					fmt.Fprintf(w, "500 %s", err)
				} else {
					var dataselected []data
					for _, d := range jdata {
						if d.Selected == true {
							dataselected = append(dataselected, d)
						}
					}
					page.Execute(w, dataselected)
				}

			}
		}

	default:
		fmt.Fprintln(w, "404 Bad request, mothod not allowed")
	}

}

// Contact function. HTML response
func Contact(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/contact.html")
	if r.Method == "GET" {
		if err != nil {
			fmt.Fprintf(w, "500 %s", err)
		} else {
			// logic here

			page.Execute(w, nil)
		}
	} else {
		fmt.Fprintln(w, "404 Bad request, mothod not allowed")
	}

}
