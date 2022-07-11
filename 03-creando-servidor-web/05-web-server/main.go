package main

const port string = ":8030"

func main() {
	server := NewServer(port)
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Loggin()))
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Listen()
	//
}
