package main

func main() {
	app := NewApp(":8000", "data.json")
	app.Serve()
}
