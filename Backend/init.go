package main

func main() {
	r := GetRouter()
	r.Run(":3100")
}
