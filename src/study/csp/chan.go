package main

func main() {
}

func channel()  {

	done := make(chan struct{})

	go func() {
		done <- struct{}{}
	}()

	<-done
}