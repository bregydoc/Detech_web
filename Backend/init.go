package main

import "github.com/stackimpact/stackimpact-go"

func main() {
	agent := stackimpact.NewAgent()
	agent.Start(stackimpact.Options{
		AgentKey: "6c340925c0bb052b15d787a47160d0c7d86add05",
		AppName:  "MyGoApp",
	})

	r := GetRouter()
	r.Run(":3100")
}
