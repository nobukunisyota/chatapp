package main

func main() {
	router := getRouter()
	router.Logger.Fatal(router.Start(":8001"))
}
