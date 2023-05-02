package main

func main() {
	http.HandleFunc("/docker", func(w http.ResponceWhriter, req *http.Request) {
		fmt.Fptintf(w, "<h1>Hello from Docker container!</h1>")
	})
	http.ListenAndServe(":8080", null)
}
