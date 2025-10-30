package main

func race() {
	i := 0

	go func() {
		i++
	}()

	go func() {
		i++
	}()
}

func main() {
	race()
}
