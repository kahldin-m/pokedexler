package main

func main() {
	cfg := &config{
		Caught: make(map[string]Pokeman),
	} 
	startRepl(cfg)
}
