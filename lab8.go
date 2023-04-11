func staticHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "lab8/"+r.URL.Path[1:])
}
func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        name := r.Form.Get("name")
        fmt.Fprintf(w, "Hello, %s!", name)
    } else {
        http.ServeFile(w, r, "lab8/form.html")
    }
}
func main() {
    http.HandleFunc("/", staticHandler)
    http.HandleFunc("/form", formHandler)

    http.ListenAndServe(":8080", nil)
}
