package main

func main() {
	newToy := NewToy(SpiderMan{})
	// This performs the spider man dialogue.
	newToy.PerformDialogue()
	// Change the behaviour at runtime.
	newToy.SetSuperHero(SuperMan{})
	// This performs the super man dialogue.
	newToy.PerformDialogue()
}
